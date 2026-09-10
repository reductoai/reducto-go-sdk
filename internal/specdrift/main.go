// Command specdrift reports drift between the SDK types and the Reducto OpenAPI spec.
//
// Schema names in the spec do not match SDK type names, so nothing is matched by name. Each
// Client method is anchored to an endpoint through its c.do call, then both sides are
// normalized into one shape model and walked in parallel, comparing JSON property names,
// types, enum values and required-ness.
//
// spec/drift-allowlist.json lists known, intentional deviations with a reason. Matching items
// are reported as allowed and do not fail the check. Entries that match nothing are flagged as
// stale.
//
// The check runs against the committed snapshot spec/openapi.json by default, so it is
// reproducible and needs no network. -live checks against the public URL. -update-snapshot
// fetches the live spec, rewrites the snapshot, then checks. Refreshing the snapshot is a
// manual step; commit it together with the SDK change.
//
// Usage:
//
//	go run ./internal/specdrift [-spec URL|PATH | -live] [-update-snapshot] [-json] [-warn-only] [-ignore KIND]...
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const liveSpec = "https://reducto.ai/openapi.json"

type ignoreFlag []string

func (f *ignoreFlag) String() string     { return strings.Join(*f, ",") }
func (f *ignoreFlag) Set(v string) error { *f = append(*f, v); return nil }

func main() {
	root := repoRoot()
	snapshot := filepath.Join(root, "spec", "openapi.json")
	allowlistPath := filepath.Join(root, "spec", "drift-allowlist.json")

	specFlag := flag.String("spec", snapshot, "OpenAPI spec URL or file path (default: snapshot)")
	live := flag.Bool("live", false, "check against "+liveSpec)
	update := flag.Bool("update-snapshot", false, "fetch the live spec, rewrite spec/openapi.json, then check against it")
	asJSON := flag.Bool("json", false, "emit JSON instead of text")
	warnOnly := flag.Bool("warn-only", false, "exit 0 even when drift is found")
	var ignore ignoreFlag
	flag.Var(&ignore, "ignore", "drift kind to ignore ("+strings.Join(driftKinds, ", ")+"); repeatable")
	flag.Parse()

	source := *specFlag
	if *live {
		source = liveSpec
	}
	if *update {
		raw, err := loadRaw(liveSpec)
		check(err)
		changed, err := writeSnapshot(raw, snapshot)
		check(err)
		state := "unchanged"
		if changed {
			state = "updated"
		}
		doc, err := parseDoc(raw)
		check(err)
		fmt.Fprintf(os.Stderr, "Snapshot %s %s (version %s)\n", display(root, snapshot), state, newSpec(doc).version)
		source = snapshot
	}

	raw, err := loadRaw(source)
	check(err)
	doc, err := parseDoc(raw)
	check(err)
	sp := newSpec(doc)
	k, err := loadSDK(root)
	check(err)
	allowlist, err := loadAllowlist(allowlistPath)
	check(err)

	var found []drift
	for _, d := range run(sp, k) {
		if !contains(ignore, d.Kind) {
			found = append(found, d)
		}
	}
	drifts, ok, stale := splitAllowed(found, allowlist)

	if *asJSON {
		payload := map[string]any{
			"spec":         source,
			"spec_version": sp.version,
			"drifts":       nonNil(drifts),
			"allowed":      nonNil(ok),
		}
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		check(enc.Encode(payload))
	} else {
		printReport(root, source, sp, k, drifts, ok, stale)
	}
	if len(drifts) > 0 && !*warnOnly {
		os.Exit(1)
	}
}

func run(sp *spec, k *sdk) []drift {
	var drifts []drift
	covered := map[[2]string]bool{}
	seenPairs := map[[2]string]bool{}
	specPaths := map[string]string{}
	for ep := range sp.endpoints() {
		specPaths[normalizePath(ep[1])] = ep[1]
	}

	eps := append([]endpoint(nil), k.eps...)
	sort.Slice(eps, func(i, j int) bool {
		if eps[i].path != eps[j].path {
			return eps[i].path < eps[j].path
		}
		return eps[i].method < eps[j].method
	})
	for _, ep := range eps {
		specPath, known := specPaths[normalizePath(ep.path)]
		if !known {
			specPath = ep.path
		}
		label := strings.ToUpper(ep.method) + " " + specPath
		op := sp.operation(ep.method, specPath)
		if op == nil {
			drifts = append(drifts, drift{label, "", "endpoint", fmt.Sprintf("sdk calls this endpoint (%s) but spec lacks it", ep.source)})
			continue
		}
		covered[[2]string{ep.method, specPath}] = true

		sections := []struct {
			name string
			spec *shape
			sdk  *shape
		}{
			{"request", sp.requestBody(op), k.optShape(ep.body)},
			{"query", sp.query(op), k.optShape(ep.query)},
			{"response", sp.response(op), k.optShape(ep.response)},
		}
		for _, s := range sections {
			switch {
			case s.spec == nil && s.sdk == nil:
			case s.spec == nil:
				drifts = append(drifts, drift{label, s.name, "extra", fmt.Sprintf("sdk sends %s but spec defines none", s.name)})
			case s.sdk == nil:
				if s.spec.kind == "any" || (s.spec.kind == "object" && len(s.spec.props) == 0) {
					continue
				}
				drifts = append(drifts, drift{label, s.name, "missing", fmt.Sprintf("spec defines %s (%s) but sdk has none", s.name, describe(s.spec))})
			default:
				c := &comparator{endpoint: label, seenPairs: seenPairs}
				c.compare(s.spec, s.sdk, s.name)
				drifts = append(drifts, c.drifts...)
			}
		}
	}

	var uncovered [][2]string
	for ep := range sp.endpoints() {
		if !covered[ep] {
			uncovered = append(uncovered, ep)
		}
	}
	sort.Slice(uncovered, func(i, j int) bool {
		if uncovered[i][1] != uncovered[j][1] {
			return uncovered[i][1] < uncovered[j][1]
		}
		return uncovered[i][0] < uncovered[j][0]
	})
	for _, ep := range uncovered {
		drifts = append(drifts, drift{strings.ToUpper(ep[0]) + " " + ep[1], "", "endpoint", "spec defines this endpoint but sdk has no method for it"})
	}
	return drifts
}

// ---- report ----

func printReport(root, source string, sp *spec, k *sdk, drifts, ok []drift, stale []allowed) {
	fmt.Printf("Spec: %s (version %s)\n", display(root, source), sp.version)
	fmt.Printf("Checked %d sdk endpoints.\n", len(k.eps))
	if len(ok) > 0 {
		fmt.Printf("%d allowed drift item(s), see spec/drift-allowlist.json:\n", len(ok))
		for _, d := range ok {
			fmt.Printf("  %s [%s]%s: %s\n", d.Endpoint, d.Kind, locSuffix(d.Location), d.Detail)
		}
	}
	for _, a := range stale {
		fmt.Fprintf(os.Stderr, "Stale allowlist entry matches nothing: %s [%s] %s\n", a.Endpoint, a.Kind, a.Location)
	}
	if len(drifts) == 0 {
		fmt.Println("No drift found.")
		return
	}
	var order []string
	byEndpoint := map[string][]drift{}
	for _, d := range drifts {
		if _, ok := byEndpoint[d.Endpoint]; !ok {
			order = append(order, d.Endpoint)
		}
		byEndpoint[d.Endpoint] = append(byEndpoint[d.Endpoint], d)
	}
	counts := map[string]int{}
	for _, ep := range order {
		fmt.Printf("\n%s\n", ep)
		for _, d := range byEndpoint[ep] {
			fmt.Printf("  [%s]%s: %s\n", d.Kind, locSuffix(d.Location), d.Detail)
			counts[d.Kind]++
		}
	}
	var parts []string
	for _, k := range sortedKeys(counts) {
		parts = append(parts, fmt.Sprintf("%s=%d", k, counts[k]))
	}
	fmt.Printf("\n%d drift item(s): %s\n", len(drifts), strings.Join(parts, ", "))
}

func locSuffix(loc string) string {
	if loc == "" {
		return ""
	}
	return " " + loc
}

// ---- io ----

func repoRoot() string {
	dir, err := os.Getwd()
	check(err)
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			check(fmt.Errorf("go.mod not found above %s", dir))
		}
		dir = parent
	}
}

func display(root, source string) string {
	if rel, err := filepath.Rel(root, source); err == nil && !strings.HasPrefix(rel, "..") {
		return rel
	}
	return source
}

func loadRaw(source string) ([]byte, error) {
	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		client := &http.Client{Timeout: 30 * time.Second}
		resp, err := client.Get(source)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("GET %s: %s", source, resp.Status)
		}
		return io.ReadAll(resp.Body)
	}
	return os.ReadFile(source)
}

func parseDoc(raw []byte) (jsonObj, error) {
	var doc jsonObj
	if err := json.Unmarshal(raw, &doc); err != nil {
		return nil, fmt.Errorf("parse spec: %w", err)
	}
	return doc, nil
}

// writeSnapshot pretty-prints the document without reordering keys.
func writeSnapshot(raw []byte, path string) (bool, error) {
	var buf bytes.Buffer
	if err := json.Indent(&buf, raw, "", "  "); err != nil {
		return false, err
	}
	buf.WriteByte('\n')
	if old, err := os.ReadFile(path); err == nil && bytes.Equal(old, buf.Bytes()) {
		return false, nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return false, err
	}
	return true, os.WriteFile(path, buf.Bytes(), 0o644)
}

func loadAllowlist(path string) ([]allowed, error) {
	raw, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var out []allowed
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	return out, nil
}

func contains(xs []string, x string) bool {
	for _, v := range xs {
		if v == x {
			return true
		}
	}
	return false
}

func nonNil(ds []drift) []drift {
	if ds == nil {
		return []drift{}
	}
	return ds
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "specdrift:", err)
		os.Exit(2)
	}
}
