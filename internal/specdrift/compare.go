package main

import (
	"fmt"
	"sort"
	"strings"
)

var driftKinds = []string{"endpoint", "missing", "extra", "type", "enum", "required"}

type drift struct {
	Endpoint string `json:"endpoint"`
	Location string `json:"location"`
	Kind     string `json:"kind"`
	Detail   string `json:"detail"`
}

type comparator struct {
	endpoint  string
	seenPairs map[[2]string]bool
	drifts    []drift
}

func (c *comparator) report(loc, kind, detail string) {
	c.drifts = append(c.drifts, drift{c.endpoint, loc, kind, detail})
}

// compare walks both shapes in parallel. Null members are dropped on both sides first: Go
// models "absent" and "null" with one pointer, so the SDK never carries a null variant.
func (c *comparator) compare(spec, sdk *shape, loc string) {
	if spec.kind == "any" || sdk.kind == "any" {
		return
	}
	spec, sdk = stripNull(spec), stripNull(sdk)
	if spec.kind == "union" || sdk.kind == "union" {
		c.compareUnion(spec, sdk, loc)
		return
	}
	if spec.kind == "object" && spec.label != "" && sdk.label != "" {
		key := [2]string{spec.label, sdk.label}
		if c.seenPairs[key] {
			return
		}
		c.seenPairs[key] = true
	}
	// Go has no literal types. A one-value enum in the spec, such as a discriminator, is read
	// into a plain field on the sdk side.
	if spec.kind == "enum" && len(spec.values) == 1 && sdk.kind == "primitive" {
		return
	}
	if spec.kind != sdk.kind {
		c.report(loc, "type", fmt.Sprintf("spec %s, sdk %s", describe(spec), describe(sdk)))
		return
	}
	switch spec.kind {
	case "primitive":
		if spec.typ != sdk.typ && !(spec.typ == "number" && sdk.typ == "integer") {
			c.report(loc, "type", fmt.Sprintf("spec %s, sdk %s", spec.typ, sdk.typ))
		}
	case "enum":
		var missing, extra []string
		for v := range spec.values {
			if !sdk.values[v] {
				missing = append(missing, v)
			}
		}
		for v := range sdk.values {
			if !spec.values[v] {
				extra = append(extra, v)
			}
		}
		sort.Strings(missing)
		sort.Strings(extra)
		if len(missing) > 0 {
			c.report(loc, "enum", "sdk lacks values ["+strings.Join(missing, ", ")+"]")
		}
		if len(extra) > 0 {
			c.report(loc, "enum", "spec lacks values ["+strings.Join(extra, ", ")+"]")
		}
	case "array":
		c.compare(spec.items, sdk.items, loc+"[]")
	case "map":
		c.compare(spec.items, sdk.items, loc+"{}")
	case "object":
		c.compareObject(spec, sdk, loc)
	}
}

func (c *comparator) compareObject(spec, sdk *shape, loc string) {
	for _, name := range sortedKeys(spec.props) {
		if _, ok := sdk.props[name]; !ok {
			c.report(loc+"."+name, "missing", fmt.Sprintf("spec has field, sdk lacks it (%s)", describe(spec.props[name])))
		}
	}
	for _, name := range sortedKeys(sdk.props) {
		if _, ok := spec.props[name]; !ok {
			c.report(loc+"."+name, "extra", fmt.Sprintf("sdk has field, spec lacks it (%s)", describe(sdk.props[name])))
		}
	}
	for _, name := range sortedKeys(spec.props) {
		sdkProp, ok := sdk.props[name]
		if !ok {
			continue
		}
		child := loc + "." + name
		c.compareRequired(spec, sdk, name, child)
		c.compare(spec.props[name], sdkProp, child)
	}
}

func (c *comparator) compareRequired(spec, sdk *shape, name, loc string) {
	specReq, sdkReq := spec.required[name], sdk.required[name]
	if specReq == sdkReq {
		return
	}
	// A pointer field covers "required but nullable"; a field with a server default is always
	// present, so the sdk may treat it as required.
	if specReq && isNullable(spec.props[name]) {
		return
	}
	if sdkReq && spec.props[name].hasDefault {
		return
	}
	c.report(loc, "required", fmt.Sprintf("spec %s, sdk %s", reqWord(specReq), reqWord(sdkReq)))
}

func reqWord(required bool) string {
	if required {
		return "required"
	}
	return "optional"
}

func (c *comparator) compareUnion(spec, sdk *shape, loc string) {
	specMembers, sdkMembers := membersOf(spec), membersOf(sdk)
	pairs := assignMembers(specMembers, sdkMembers)
	matchedSpec, matchedSDK := map[int]bool{}, map[int]bool{}
	for _, p := range pairs {
		sm := specMembers[p.spec]
		name := sm.label
		if name == "" {
			name = fmt.Sprint(p.spec)
		}
		c.compare(sm, sdkMembers[p.sdk], fmt.Sprintf("%s<%s>", loc, name))
		matchedSpec[p.spec], matchedSDK[p.sdk] = true, true
	}
	for i, sm := range specMembers {
		if !matchedSpec[i] {
			c.report(loc, "missing", fmt.Sprintf("spec union member %s has no sdk counterpart", describe(sm)))
		}
	}
	for i, km := range sdkMembers {
		if !matchedSDK[i] {
			c.report(loc, "extra", fmt.Sprintf("sdk union member %s has no spec counterpart", describe(km)))
		}
	}
}

// ---- allowlist ----

type allowed struct {
	Endpoint string  `json:"endpoint"`
	Location string  `json:"location"`
	Kind     string  `json:"kind"`
	Detail   *string `json:"detail,omitempty"`
	Reason   string  `json:"reason"`
}

func (a allowed) matches(d drift) bool {
	if a.Endpoint != d.Endpoint || a.Location != d.Location || a.Kind != d.Kind {
		return false
	}
	return a.Detail == nil || *a.Detail == d.Detail
}

func splitAllowed(drifts []drift, allowlist []allowed) (active, ok []drift, stale []allowed) {
	used := map[int]bool{}
	for _, d := range drifts {
		hit := -1
		for i, a := range allowlist {
			if a.matches(d) {
				hit = i
				break
			}
		}
		if hit < 0 {
			active = append(active, d)
		} else {
			ok = append(ok, d)
			used[hit] = true
		}
	}
	for i, a := range allowlist {
		if !used[i] {
			stale = append(stale, a)
		}
	}
	return active, ok, stale
}
