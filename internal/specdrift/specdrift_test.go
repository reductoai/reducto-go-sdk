package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const fixtureSpec = `{
  "info": {"version": "test"},
  "paths": {
    "/foo": {"post": {
      "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/FooConfig"}}}},
      "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/FooResponse"}}}}}
    }},
    "/foo/{foo_id}": {"get": {
      "parameters": [{"name": "limit", "in": "query", "schema": {"type": "integer"}}],
      "responses": {"200": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/FooResponse"}}}}}
    }},
    "/bar": {"delete": {"responses": {"200": {"content": {"application/json": {"schema": {}}}}}}}
  },
  "components": {"schemas": {
    "FooConfig": {"type": "object", "required": ["a", "kind"], "properties": {
      "a": {"type": "string"},
      "b": {"type": "string", "enum": ["x", "y"]},
      "c": {"anyOf": [{"type": "integer"}, {"type": "null"}]},
      "kind": {"const": "foo"},
      "when": {"type": "string", "format": "date-time"}
    }},
    "FooResponse": {"type": "object", "required": ["id", "n"], "properties": {
      "id": {"type": "string"},
      "n": {"type": "number"},
      "items": {"type": "array", "items": {"$ref": "#/components/schemas/FooResponse"}}
    }}
  }}
}`

var fixtureSDK = strings.ReplaceAll(`package reducto

import (
	"context"
	"encoding/json"
	"net/url"
)

type Time struct{}

type FooB string

const (
	FooBX FooB = "x"
	FooBZ FooB = "z"
)

type FooConfig struct {
	A    string ~json:"a"~
	B    FooB   ~json:"b,omitempty"~
	C    *int64 ~json:"c,omitempty"~
	Kind string ~json:"kind"~
	When *Time  ~json:"when,omitempty"~
	Old  bool   ~json:"old,omitempty"~
}

type FooResponse struct {
	ID    string          ~json:"id"~
	N     int64           ~json:"n,omitempty"~
	Items []FooResponse   ~json:"items,omitempty"~
	Raw   json.RawMessage ~json:"raw,omitempty"~
}

type GetFooParams struct {
	Limit  *int64
	Cursor string
}

func (p *GetFooParams) query() url.Values { return nil }

func (c *Client) Foo(ctx context.Context, req *FooConfig, opts ...Option) (*FooResponse, error) {
	var out FooResponse
	if err := c.do(ctx, "POST", "/foo", nil, req, &out, opts); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetFoo(ctx context.Context, fooID string, params *GetFooParams, opts ...Option) (*FooResponse, error) {
	var out FooResponse
	if err := c.do(ctx, "GET", "/foo/"+url.PathEscape(fooID)+"", params.query(), nil, &out, opts); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) Gone(ctx context.Context, opts ...Option) (string, error) {
	var out string
	if err := c.do(ctx, "GET", "/gone", nil, nil, &out, opts); err != nil {
		return "", err
	}
	return out, nil
}
`, "~", "`")

func fixture(t *testing.T) (*spec, *sdk) {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "api.go"), []byte(fixtureSDK), 0o644); err != nil {
		t.Fatal(err)
	}
	var doc jsonObj
	if err := json.Unmarshal([]byte(fixtureSpec), &doc); err != nil {
		t.Fatal(err)
	}
	k, err := loadSDK(dir)
	if err != nil {
		t.Fatal(err)
	}
	return newSpec(doc), k
}

func TestRunFindsEachDriftKind(t *testing.T) {
	sp, k := fixture(t)
	if len(k.eps) != 3 {
		t.Fatalf("endpoints = %d", len(k.eps))
	}
	got := map[string]bool{}
	for _, d := range run(sp, k) {
		got[d.Endpoint+" ["+d.Kind+"] "+d.Location+": "+d.Detail] = true
	}
	want := []string{
		"POST /foo [enum] request.b: sdk lacks values [y]",
		"POST /foo [enum] request.b: spec lacks values [z]",
		"POST /foo [extra] request.old: sdk has field, spec lacks it (boolean)",
		"POST /foo [extra] response.raw: sdk has field, spec lacks it (any)",
		"POST /foo [required] response.n: spec required, sdk optional",
		"GET /foo/{foo_id} [extra] query.cursor: sdk has field, spec lacks it (string)",
		"GET /gone [endpoint] : sdk calls this endpoint (api.go:59) but spec lacks it",
		"DELETE /bar [endpoint] : spec defines this endpoint but sdk has no method for it",
	}
	for _, w := range want {
		if !got[w] {
			t.Errorf("missing drift %q", w)
		}
	}
	for g := range got {
		if strings.Contains(g, "request.kind") || strings.Contains(g, "request.c") || strings.Contains(g, "request.when") {
			t.Errorf("unexpected drift %q", g)
		}
	}
	if len(got) != len(want) {
		t.Errorf("got %d drifts, want %d:\n%s", len(got), len(want), strings.Join(sortedKeys(got), "\n"))
	}
}

func TestAllowlistMatchingAndStale(t *testing.T) {
	drifts := []drift{
		{"POST /foo", "request.old", "extra", "sdk has field, spec lacks it (boolean)"},
		{"POST /foo", "response.n", "required", "spec required, sdk optional"},
	}
	detail := "no such detail"
	list := []allowed{
		{Endpoint: "POST /foo", Location: "request.old", Kind: "extra", Reason: "kept for old clients"},
		{Endpoint: "POST /foo", Location: "response.n", Kind: "required", Detail: &detail, Reason: "wrong detail"},
	}
	active, ok, stale := splitAllowed(drifts, list)
	if len(ok) != 1 || len(active) != 1 || len(stale) != 1 {
		t.Fatalf("active=%d ok=%d stale=%d", len(active), len(ok), len(stale))
	}
	if active[0].Location != "response.n" || stale[0].Location != "response.n" {
		t.Errorf("active=%+v stale=%+v", active, stale)
	}
}

func TestSnake(t *testing.T) {
	for in, want := range map[string]string{
		"Limit": "limit", "IncludePersisted": "include_persisted", "JobID": "job_id", "URLPath": "url_path", "Page2Size": "page2_size",
	} {
		if got := snake(in); got != want {
			t.Errorf("snake(%q) = %q, want %q", in, got, want)
		}
	}
}
