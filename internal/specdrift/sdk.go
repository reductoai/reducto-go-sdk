package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// The SDK side is read from source with go/ast. Reflection cannot see enum constants or which
// path a method calls, and this way the checker does not import the package it audits.
type sdk struct {
	fset  *token.FileSet
	types map[string]ast.Expr
	enums map[string][]string
	eps   []endpoint
}

type endpoint struct {
	method   string
	path     string
	source   string
	body     ast.Expr
	query    ast.Expr
	response ast.Expr
}

func loadSDK(root string) (*sdk, error) {
	files, err := filepath.Glob(filepath.Join(root, "*.go"))
	if err != nil {
		return nil, err
	}
	k := &sdk{fset: token.NewFileSet(), types: map[string]ast.Expr{}, enums: map[string][]string{}}
	seen := map[[2]string]bool{}
	sort.Strings(files)
	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") {
			continue
		}
		file, err := parser.ParseFile(k.fset, f, nil, parser.SkipObjectResolution)
		if err != nil {
			return nil, err
		}
		for _, d := range file.Decls {
			switch d := d.(type) {
			case *ast.GenDecl:
				k.collectDecl(d)
			case *ast.FuncDecl:
				k.collectEndpoints(root, d, seen)
			}
		}
	}
	return k, nil
}

func (k *sdk) collectDecl(d *ast.GenDecl) {
	for _, s := range d.Specs {
		switch s := s.(type) {
		case *ast.TypeSpec:
			k.types[s.Name.Name] = s.Type
		case *ast.ValueSpec:
			id, ok := s.Type.(*ast.Ident)
			if !ok || len(s.Values) == 0 {
				continue
			}
			for _, v := range s.Values {
				lit, ok := v.(*ast.BasicLit)
				if !ok || lit.Kind != token.STRING {
					continue
				}
				val, err := strconv.Unquote(lit.Value)
				if err == nil {
					k.enums[id.Name] = append(k.enums[id.Name], val)
				}
			}
		}
	}
}

func receiverIsClient(fd *ast.FuncDecl) bool {
	if fd.Recv == nil || len(fd.Recv.List) == 0 {
		return false
	}
	star, ok := fd.Recv.List[0].Type.(*ast.StarExpr)
	if !ok {
		return false
	}
	id, ok := star.X.(*ast.Ident)
	return ok && id.Name == "Client"
}

// collectEndpoints anchors a Client method to an endpoint by finding
// c.do(ctx, "METHOD", path, query, body, &out, opts) in its body.
func (k *sdk) collectEndpoints(root string, fd *ast.FuncDecl, seen map[[2]string]bool) {
	if !receiverIsClient(fd) || fd.Body == nil {
		return
	}
	params := map[string]ast.Expr{}
	for _, p := range fd.Type.Params.List {
		for _, n := range p.Names {
			params[n.Name] = p.Type
		}
	}
	ast.Inspect(fd.Body, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok || len(call.Args) < 6 {
			return true
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok || sel.Sel.Name != "do" {
			return true
		}
		if recv, ok := sel.X.(*ast.Ident); !ok || recv.Name != "c" {
			return true
		}
		method, ok := stringLit(call.Args[1])
		if !ok {
			return true
		}
		path := pathString(call.Args[2])
		key := [2]string{strings.ToLower(method), normalizePath(path)}
		if seen[key] {
			return true
		}
		seen[key] = true
		rel, _ := filepath.Rel(root, k.fset.Position(call.Pos()).Filename)
		ep := endpoint{
			method:   strings.ToLower(method),
			path:     path,
			source:   fmt.Sprintf("%s:%d", rel, k.fset.Position(call.Pos()).Line),
			query:    argType(call.Args[3], params),
			body:     argType(call.Args[4], params),
			response: findVarType(fd.Body, "out"),
		}
		if ov, ok := overrides[key]; ok {
			ov(&ep, params)
		}
		k.eps = append(k.eps, ep)
		return true
	})
}

// overrides cover the methods whose request shape is not visible in the c.do call.
var overrides = map[[2]string]func(*endpoint, map[string]ast.Expr){
	{"post", "/upload"}: func(ep *endpoint, params map[string]ast.Expr) {
		// The body is a multipart file stream built by hand; the query is built from UploadOptions.
		ep.body = &ast.Ident{Name: "uploadBody"}
		ep.query = &ast.Ident{Name: "UploadOptions"}
	},
}

// syntheticTypes are shapes for things that have no Go type of their own.
var syntheticTypes = map[string]*shape{
	"uploadBody": func() *shape {
		o := object("")
		o.props["file"] = prim("file")
		return o
	}(),
}

func stringLit(e ast.Expr) (string, bool) {
	lit, ok := e.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return "", false
	}
	s, err := strconv.Unquote(lit.Value)
	return s, err == nil
}

// pathString turns "/job/"+url.PathEscape(jobID)+"" into /job/{jobID}.
func pathString(e ast.Expr) string {
	switch e := e.(type) {
	case *ast.BasicLit:
		s, _ := stringLit(e)
		return s
	case *ast.BinaryExpr:
		return pathString(e.X) + pathString(e.Y)
	case *ast.CallExpr:
		if len(e.Args) == 1 {
			if id, ok := e.Args[0].(*ast.Ident); ok {
				return "{" + id.Name + "}"
			}
		}
		return "{?}"
	case *ast.Ident:
		return "{" + e.Name + "}"
	}
	return "{?}"
}

// normalizePath makes /job/{jobID} and /job/{job_id} compare equal.
func normalizePath(p string) string {
	var b strings.Builder
	inBrace := false
	for _, r := range p {
		switch {
		case r == '{':
			inBrace = true
			b.WriteString("{}")
		case r == '}':
			inBrace = false
		case !inBrace:
			b.WriteRune(r)
		}
	}
	return b.String()
}

// argType maps a c.do argument (nil, req, params.query()) to the parameter type behind it.
func argType(e ast.Expr, params map[string]ast.Expr) ast.Expr {
	switch e := e.(type) {
	case *ast.Ident:
		if e.Name == "nil" {
			return nil
		}
		return params[e.Name]
	case *ast.CallExpr:
		if sel, ok := e.Fun.(*ast.SelectorExpr); ok {
			if id, ok := sel.X.(*ast.Ident); ok {
				return params[id.Name]
			}
		}
	}
	return nil
}

func findVarType(body *ast.BlockStmt, name string) ast.Expr {
	var found ast.Expr
	ast.Inspect(body, func(n ast.Node) bool {
		if found != nil {
			return false
		}
		ds, ok := n.(*ast.DeclStmt)
		if !ok {
			return true
		}
		gd, ok := ds.Decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.VAR {
			return true
		}
		for _, s := range gd.Specs {
			vs, ok := s.(*ast.ValueSpec)
			if !ok || vs.Type == nil {
				continue
			}
			for _, n := range vs.Names {
				if n.Name == name {
					found = vs.Type
					return false
				}
			}
		}
		return true
	})
	return found
}

// ---- shapes ----

func (k *sdk) optShape(e ast.Expr) *shape {
	if e == nil {
		return nil
	}
	return k.shapeOf(e, nil)
}

func (k *sdk) shapeOf(e ast.Expr, seen map[string]bool) *shape {
	switch e := e.(type) {
	case *ast.StarExpr:
		return k.shapeOf(e.X, seen)
	case *ast.ArrayType:
		if id, ok := e.Elt.(*ast.Ident); ok && id.Name == "byte" {
			return anyShape
		}
		return &shape{kind: "array", items: k.shapeOf(e.Elt, seen)}
	case *ast.MapType:
		return &shape{kind: "map", items: k.shapeOf(e.Value, seen)}
	case *ast.InterfaceType:
		return anyShape
	case *ast.SelectorExpr:
		switch exprString(e) {
		case "time.Time":
			return prim("string")
		}
		return anyShape
	case *ast.StructType:
		return k.structShape(e, "", seen)
	case *ast.Ident:
		return k.namedShape(e.Name, seen)
	}
	return &shape{kind: "any", label: exprString(e)}
}

func (k *sdk) namedShape(name string, seen map[string]bool) *shape {
	switch name {
	case "string":
		return prim("string")
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return prim("integer")
	case "float32", "float64":
		return prim("number")
	case "bool":
		return prim("boolean")
	case "any":
		return anyShape
	case "Time":
		return prim("string")
	}
	if s, ok := syntheticTypes[name]; ok {
		return s
	}
	if vals, ok := k.enums[name]; ok {
		return enum(name, vals...)
	}
	decl, ok := k.types[name]
	if !ok {
		return &shape{kind: "any", label: name}
	}
	if seen[name] {
		return &shape{kind: "any", label: name}
	}
	seen = with(seen, name)
	if st, ok := decl.(*ast.StructType); ok {
		if isUnionStruct(st) {
			var members []*shape
			for _, f := range st.Fields.List {
				if len(f.Names) == 1 && f.Names[0].Name == "Unknown" {
					continue
				}
				members = append(members, k.shapeOf(f.Type, seen))
			}
			return union(members)
		}
		return k.structShape(st, name, seen)
	}
	s := k.shapeOf(decl, seen)
	if s.label == "" {
		cp := *s
		cp.label = name
		s = &cp
	}
	return s
}

// A union is a struct of untagged variant fields plus Unknown json.RawMessage.
func isUnionStruct(st *ast.StructType) bool {
	hasUnknown := false
	for _, f := range st.Fields.List {
		if f.Tag != nil {
			return false
		}
		if len(f.Names) == 1 && f.Names[0].Name == "Unknown" && exprString(f.Type) == "json.RawMessage" {
			hasUnknown = true
		}
	}
	return hasUnknown
}

func (k *sdk) structShape(st *ast.StructType, label string, seen map[string]bool) *shape {
	o := object(label)
	for _, f := range st.Fields.List {
		if len(f.Names) == 0 {
			continue
		}
		name := f.Names[0].Name
		if !ast.IsExported(name) {
			continue
		}
		wire, required := name, true
		if f.Tag != nil {
			raw, _ := strconv.Unquote(f.Tag.Value)
			tag := reflect.StructTag(raw).Get("json")
			if tag == "-" {
				continue
			}
			parts := strings.Split(tag, ",")
			if parts[0] != "" {
				wire = parts[0]
			}
			for _, opt := range parts[1:] {
				if opt == "omitempty" || opt == "omitzero" {
					required = false
				}
			}
		} else {
			// Query parameter structs carry no tags: the wire name is snake_case and a
			// pointer marks an optional parameter.
			wire = snake(name)
			_, isPtr := f.Type.(*ast.StarExpr)
			required = !isPtr
		}
		o.props[wire] = k.shapeOf(f.Type, seen)
		if required {
			o.required[wire] = true
		}
	}
	return o
}

func snake(name string) string {
	rs := []rune(name)
	var b strings.Builder
	for i, r := range rs {
		if unicode.IsUpper(r) && i > 0 {
			prev := rs[i-1]
			nextLower := i+1 < len(rs) && unicode.IsLower(rs[i+1])
			if unicode.IsLower(prev) || unicode.IsDigit(prev) || (unicode.IsUpper(prev) && nextLower) {
				b.WriteByte('_')
			}
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}

func exprString(e ast.Expr) string {
	switch e := e.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.SelectorExpr:
		return exprString(e.X) + "." + e.Sel.Name
	case *ast.StarExpr:
		return "*" + exprString(e.X)
	case *ast.ArrayType:
		return "[]" + exprString(e.Elt)
	case *ast.MapType:
		return "map[" + exprString(e.Key) + "]" + exprString(e.Value)
	}
	return fmt.Sprintf("%T", e)
}
