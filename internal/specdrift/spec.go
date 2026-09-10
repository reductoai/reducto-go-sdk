package main

import (
	"fmt"
	"strings"
)

var httpVerbs = []string{"get", "post", "put", "patch", "delete"}

var successCodes = []string{"200", "201", "202"}

type jsonObj = map[string]any

type spec struct {
	doc     jsonObj
	schemas jsonObj
	version string
}

func newSpec(doc jsonObj) *spec {
	s := &spec{doc: doc, schemas: jsonObj{}, version: "unknown"}
	if c, ok := doc["components"].(jsonObj); ok {
		if sch, ok := c["schemas"].(jsonObj); ok {
			s.schemas = sch
		}
	}
	if info, ok := doc["info"].(jsonObj); ok {
		if v, ok := info["version"]; ok {
			s.version = fmt.Sprint(v)
		}
	}
	return s
}

func refName(node jsonObj) string {
	ref, _ := node["$ref"].(string)
	if ref == "" {
		return ""
	}
	return ref[strings.LastIndex(ref, "/")+1:]
}

func (s *spec) resolve(node jsonObj) jsonObj {
	for {
		name := refName(node)
		if name == "" {
			return node
		}
		next, ok := s.schemas[name].(jsonObj)
		if !ok {
			return jsonObj{}
		}
		node = next
	}
}

func (s *spec) shape(node jsonObj, seen map[string]bool) *shape {
	sh := s.shapeOf(node, seen)
	if _, ok := node["default"]; ok && !sh.hasDefault {
		cp := *sh
		cp.hasDefault = true
		sh = &cp
	}
	return sh
}

func (s *spec) shapeOf(node jsonObj, seen map[string]bool) *shape {
	label := refName(node)
	if label != "" {
		if seen[label] {
			return &shape{kind: "any", label: label}
		}
		seen = with(seen, label)
		node = s.resolve(node)
	}

	variants, _ := node["anyOf"].([]any)
	if variants == nil {
		variants, _ = node["oneOf"].([]any)
	}
	if len(variants) > 0 {
		members := make([]*shape, 0, len(variants))
		for _, v := range variants {
			if vo, ok := v.(jsonObj); ok {
				members = append(members, s.shape(vo, seen))
			}
		}
		return union(members)
	}
	if allOf, ok := node["allOf"].([]any); ok {
		merged := jsonObj{"type": "object", "properties": jsonObj{}, "required": []any{}}
		for _, part := range allOf {
			po, _ := part.(jsonObj)
			po = s.resolve(po)
			if props, ok := po["properties"].(jsonObj); ok {
				for k, v := range props {
					merged["properties"].(jsonObj)[k] = v
				}
			}
			if req, ok := po["required"].([]any); ok {
				merged["required"] = append(merged["required"].([]any), req...)
			}
		}
		node = merged
	}

	if c, ok := node["const"]; ok {
		return enum(label, fmt.Sprint(c))
	}
	if vals, ok := node["enum"].([]any); ok {
		e := enum(label)
		for _, v := range vals {
			e.values[fmt.Sprint(v)] = true
		}
		return e
	}

	if types, ok := node["type"].([]any); ok {
		members := make([]*shape, 0, len(types))
		for _, t := range types {
			cp := jsonObj{}
			for k, v := range node {
				cp[k] = v
			}
			cp["type"] = t
			members = append(members, s.shape(cp, seen))
		}
		return union(members)
	}
	t, _ := node["type"].(string)
	_, hasProps := node["properties"]
	switch {
	case t == "array":
		items, _ := node["items"].(jsonObj)
		if items == nil {
			items = jsonObj{}
		}
		return &shape{kind: "array", items: s.shape(items, seen), label: label}
	case t == "object" || hasProps:
		if props, ok := node["properties"].(jsonObj); ok {
			o := object(label)
			for k, v := range props {
				vo, _ := v.(jsonObj)
				if vo == nil {
					vo = jsonObj{}
				}
				o.props[k] = s.shape(vo, seen)
			}
			if req, ok := node["required"].([]any); ok {
				for _, r := range req {
					o.required[fmt.Sprint(r)] = true
				}
			}
			return o
		}
		if extra, ok := node["additionalProperties"].(jsonObj); ok {
			return &shape{kind: "map", items: s.shape(extra, seen), label: label}
		}
		return &shape{kind: "map", items: anyShape, label: label}
	case t == "string" && node["format"] == "binary":
		return prim("file")
	case t == "string", t == "integer", t == "number", t == "boolean", t == "null":
		return prim(t)
	}
	return &shape{kind: "any", label: label}
}

func with(seen map[string]bool, k string) map[string]bool {
	cp := make(map[string]bool, len(seen)+1)
	for x := range seen {
		cp[x] = true
	}
	cp[k] = true
	return cp
}

func (s *spec) operation(method, path string) jsonObj {
	paths, _ := s.doc["paths"].(jsonObj)
	item, _ := paths[path].(jsonObj)
	op, _ := item[method].(jsonObj)
	return op
}

func content(node jsonObj, key string) jsonObj {
	sub, _ := node[key].(jsonObj)
	c, _ := sub["content"].(jsonObj)
	return c
}

func mediaSchema(c jsonObj, ct string) jsonObj {
	m, _ := c[ct].(jsonObj)
	sch, _ := m["schema"].(jsonObj)
	return sch
}

func (s *spec) requestBody(op jsonObj) *shape {
	c := content(op, "requestBody")
	for _, ct := range []string{"application/json", "multipart/form-data", "application/x-www-form-urlencoded"} {
		if sch := mediaSchema(c, ct); sch != nil {
			return s.shape(sch, nil)
		}
	}
	return nil
}

func (s *spec) query(op jsonObj) *shape {
	params, _ := op["parameters"].([]any)
	var o *shape
	for _, p := range params {
		po, _ := p.(jsonObj)
		if po["in"] != "query" {
			continue
		}
		if o == nil {
			o = object("")
		}
		name, _ := po["name"].(string)
		sch, _ := po["schema"].(jsonObj)
		if sch == nil {
			sch = jsonObj{}
		}
		o.props[name] = s.shape(sch, nil)
		if req, _ := po["required"].(bool); req {
			o.required[name] = true
		}
	}
	return o
}

func (s *spec) response(op jsonObj) *shape {
	responses, _ := op["responses"].(jsonObj)
	for _, code := range successCodes {
		if sch := mediaSchema(content(responses, code), "application/json"); sch != nil {
			return s.shape(sch, nil)
		}
	}
	return nil
}

func (s *spec) endpoints() map[[2]string]bool {
	out := map[[2]string]bool{}
	paths, _ := s.doc["paths"].(jsonObj)
	for path, item := range paths {
		io, _ := item.(jsonObj)
		for _, verb := range httpVerbs {
			if _, ok := io[verb]; ok {
				out[[2]string{verb, path}] = true
			}
		}
	}
	return out
}
