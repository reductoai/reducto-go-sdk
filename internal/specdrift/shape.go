package main

import (
	"fmt"
	"sort"
	"strings"
)

// One shape model for both sides. Spec schemas and Go types are normalized into it and then
// walked in parallel.
type shape struct {
	kind       string // object, enum, primitive, array, map, union, any
	typ        string // primitive: string, integer, number, boolean, null, file
	props      map[string]*shape
	required   map[string]bool
	values     map[string]bool
	items      *shape
	members    []*shape
	label      string
	hasDefault bool
}

var (
	anyShape  = &shape{kind: "any"}
	nullShape = prim("null")
)

func prim(t string) *shape { return &shape{kind: "primitive", typ: t} }

func object(label string) *shape {
	return &shape{kind: "object", props: map[string]*shape{}, required: map[string]bool{}, label: label}
}

func enum(label string, values ...string) *shape {
	s := &shape{kind: "enum", values: map[string]bool{}, label: label}
	for _, v := range values {
		s.values[v] = true
	}
	return s
}

func union(members []*shape) *shape {
	var flat []*shape
	for _, m := range members {
		if m.kind == "union" {
			flat = append(flat, m.members...)
		} else {
			flat = append(flat, m)
		}
	}
	var dedup []*shape
	for _, m := range flat {
		dup := false
		for _, d := range dedup {
			if sameSignature(m, d) {
				dup = true
				break
			}
		}
		if !dup {
			dedup = append(dedup, m)
		}
	}
	if len(dedup) == 1 {
		return dedup[0]
	}
	return &shape{kind: "union", members: dedup}
}

func sameSignature(a, b *shape) bool {
	if a.kind != b.kind {
		return false
	}
	switch a.kind {
	case "primitive":
		return a.typ == b.typ
	case "object":
		if len(a.props) != len(b.props) {
			return false
		}
		for k, av := range a.props {
			bv, ok := b.props[k]
			if !ok || !sameSignature(av, bv) {
				return false
			}
		}
		return true
	case "enum":
		return setEqual(a.values, b.values)
	case "array", "map":
		return a.items != nil && b.items != nil && sameSignature(a.items, b.items)
	case "union":
		if len(a.members) != len(b.members) {
			return false
		}
		for _, x := range a.members {
			found := false
			for _, y := range b.members {
				if sameSignature(x, y) {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
		return true
	}
	return true
}

func setEqual(a, b map[string]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if !b[k] {
			return false
		}
	}
	return true
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func stripNull(s *shape) *shape {
	if s.kind != "union" {
		return s
	}
	var rest []*shape
	for _, m := range s.members {
		if !(m.kind == "primitive" && m.typ == "null") {
			rest = append(rest, m)
		}
	}
	if len(rest) == 0 {
		return s
	}
	return union(rest)
}

func isNullable(s *shape) bool {
	if s.kind == "primitive" && s.typ == "null" {
		return true
	}
	if s.kind == "union" {
		for _, m := range s.members {
			if isNullable(m) {
				return true
			}
		}
	}
	return false
}

func describe(s *shape) string {
	switch s.kind {
	case "primitive":
		return s.typ
	case "enum":
		return "enum[" + strings.Join(sortedKeys(s.values), ", ") + "]"
	case "object":
		if len(s.props) <= 6 {
			return "object{" + strings.Join(sortedKeys(s.props), ", ") + "}"
		}
		return fmt.Sprintf("object(%d fields)", len(s.props))
	case "array":
		return "array<" + describeItems(s) + ">"
	case "map":
		return "map<" + describeItems(s) + ">"
	case "union":
		parts := make([]string, len(s.members))
		for i, m := range s.members {
			parts[i] = describe(m)
		}
		return strings.Join(parts, " | ")
	}
	return "any"
}

func describeItems(s *shape) string {
	if s.items == nil {
		return "any"
	}
	return describe(s.items)
}

// ---- union member matching ----

const similarityDepth = 2

type pair struct{ spec, sdk int }

// assignMembers greedily pairs union members one-to-one by structural similarity.
func assignMembers(specMembers, sdkMembers []*shape) []pair {
	type scored struct {
		score     float64
		spec, sdk int
	}
	var all []scored
	for si, sm := range specMembers {
		for ki, km := range sdkMembers {
			if score := similarity(sm, km, similarityDepth); score > 0 {
				all = append(all, scored{score, si, ki})
			}
		}
	}
	sort.SliceStable(all, func(i, j int) bool {
		if all[i].score != all[j].score {
			return all[i].score > all[j].score
		}
		if all[i].spec != all[j].spec {
			return all[i].spec < all[j].spec
		}
		return all[i].sdk < all[j].sdk
	})
	usedSpec, usedSDK := map[int]bool{}, map[int]bool{}
	var pairs []pair
	for _, s := range all {
		if usedSpec[s.spec] || usedSDK[s.sdk] {
			continue
		}
		usedSpec[s.spec], usedSDK[s.sdk] = true, true
		pairs = append(pairs, pair{s.spec, s.sdk})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].spec != pairs[j].spec {
			return pairs[i].spec < pairs[j].spec
		}
		return pairs[i].sdk < pairs[j].sdk
	})
	return pairs
}

func similarity(a, b *shape, depth int) float64 {
	if a.kind == "any" || b.kind == "any" {
		return 0.1
	}
	if a.kind == "union" || b.kind == "union" {
		am, bm := membersOf(a), membersOf(b)
		best := 0.0
		for _, x := range am {
			for _, y := range bm {
				if s := similarity(x, y, depth); s > best {
					best = s
				}
			}
		}
		return best
	}
	if a.kind != b.kind {
		return 0
	}
	switch a.kind {
	case "primitive":
		if a.typ == b.typ {
			return 1
		}
		return 0
	case "enum":
		if len(a.values) == 0 || len(b.values) == 0 {
			return 0.5
		}
		inter, uni := 0, len(a.values)
		for v := range b.values {
			if a.values[v] {
				inter++
			} else {
				uni++
			}
		}
		if inter == 0 {
			return 0.2
		}
		return 0.5 + 0.5*float64(inter)/float64(uni)
	case "object":
		if len(a.props) == 0 && len(b.props) == 0 {
			return 1
		}
		shared, uni := 0, len(a.props)
		for k := range b.props {
			if _, ok := a.props[k]; ok {
				shared++
			} else {
				uni++
			}
		}
		nameScore := float64(shared) / float64(max(uni, 1))
		if depth <= 0 || shared == 0 {
			return 0.5 + 0.5*nameScore
		}
		childScore := 0.0
		for k, av := range a.props {
			if bv, ok := b.props[k]; ok {
				childScore += similarity(av, bv, depth-1)
			}
		}
		childScore /= float64(shared)
		return 0.5 + 0.3*nameScore + 0.2*childScore
	case "array", "map":
		if a.items == nil || b.items == nil {
			return 0.5
		}
		return 0.5 + 0.5*similarity(a.items, b.items, depth)
	}
	return 0.5
}

func membersOf(s *shape) []*shape {
	if s.kind == "union" {
		return s.members
	}
	return []*shape{s}
}
