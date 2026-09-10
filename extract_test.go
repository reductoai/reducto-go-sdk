package reducto

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

type invoice struct {
	Title     string `json:"title"`
	PageCount int    `json:"page_count"`
}

var invoiceSchema = map[string]any{
	"type":       "object",
	"properties": map[string]any{"title": map[string]any{"type": "string"}, "page_count": map[string]any{"type": "integer"}},
	"required":   []string{"title"},
}

const extractBody = `{"response_type":"v3_extract","job_id":"j1","usage":{"num_pages":1,"num_fields":2},"result":%s}`

const legacyExtractBody = `{"response_type":"extract","job_id":"j1","usage":{"num_pages":1,"num_fields":2},"result":%s}`

func extractServer(t *testing.T, result string, gotBody *map[string]any) *Client {
	return newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if gotBody != nil {
			_ = json.NewDecoder(r.Body).Decode(gotBody)
		}
		io.WriteString(w, strings.Replace(extractBody, "%s", result, 1))
	})
}

func TestExtractAsSendsSchemaAndDecodes(t *testing.T) {
	var gotBody map[string]any
	c := extractServer(t, `[{"title":"Invoice","page_count":1}]`, &gotBody)
	out, err := ExtractAs[invoice](context.Background(), c, invoiceSchema, &SyncExtractConfig{
		Input:        DocumentInputFromString("https://x/y.pdf"),
		Instructions: &Instructions{SystemPrompt: Ptr("Be terse")},
	})
	if err != nil {
		t.Fatal(err)
	}
	ins := gotBody["instructions"].(map[string]any)
	if ins["system_prompt"] != "Be terse" || ins["schema"].(map[string]any)["type"] != "object" {
		t.Errorf("instructions sent as %v", ins)
	}
	if len(out.Result) != 1 || out.Result[0] != (invoice{"Invoice", 1}) {
		t.Errorf("result = %+v", out.Result)
	}
	if out.Usage().NumFields != 2 || out.V3ExtractResponse == nil || out.ExtractResponse != nil {
		t.Errorf("response = %+v", out)
	}
}

func TestExtractAsDecodesLegacySyncShape(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, strings.Replace(legacyExtractBody, "%s", `[{"title":"Legacy","page_count":3}]`, 1))
	})
	out, err := ExtractAs[invoice](context.Background(), c, invoiceSchema, &SyncExtractConfig{Input: DocumentInputFromString("https://x/y.pdf")})
	if err != nil {
		t.Fatal(err)
	}
	if len(out.Result) != 1 || out.Result[0] != (invoice{"Legacy", 3}) || out.ExtractResponse == nil || out.V3ExtractResponse != nil {
		t.Errorf("out = %+v", out)
	}
}

func TestExtractOutputUnknownResponseTypeIsNotAsync(t *testing.T) {
	var out ExtractOutput
	if err := json.Unmarshal([]byte(`{"response_type":"v4_extract","job_id":"j1","result":[]}`), &out); err != nil {
		t.Fatal(err)
	}
	if out.AsyncExtractResponse != nil || out.Unknown == nil {
		t.Fatalf("out = %+v", out)
	}
	if err := json.Unmarshal([]byte(`{"job_id":"j1"}`), &out); err != nil {
		t.Fatal(err)
	}
	if out.AsyncExtractResponse == nil || out.AsyncExtractResponse.JobID != "j1" {
		t.Fatalf("out = %+v", out)
	}
	_, err := ValidateExtract[invoice](&ExtractOutput{Unknown: json.RawMessage(`{}`)})
	if err == nil {
		t.Error("expected an error for an unknown shape")
	}
}

func TestExtractAsDecodeErrorKeepsResponse(t *testing.T) {
	c := extractServer(t, `[{"title":"ok"},{"title":1}]`, nil)
	_, err := ExtractAs[invoice](context.Background(), c, invoiceSchema, &SyncExtractConfig{Input: DocumentInputFromString("https://x/y.pdf")})
	var te *TypedExtractError
	if !errors.As(err, &te) {
		t.Fatalf("err = %v", err)
	}
	if !strings.Contains(te.Reason, "item 1 does not match reducto.invoice") || te.Err == nil {
		t.Errorf("reason = %q err = %v", te.Reason, te.Err)
	}
	resp, ok := te.Response.(*ExtractOutput)
	if !ok || resp.V3ExtractResponse == nil {
		t.Fatalf("response = %#v", te.Response)
	}
	if items, _ := resp.V3ExtractResponse.Result.([]any); len(items) != 2 {
		t.Errorf("result = %#v", resp.V3ExtractResponse.Result)
	}
}

func TestExtractAsURLResultAndQueuedJob(t *testing.T) {
	c := extractServer(t, `{"type":"url","url":"https://s3/x.json","result_id":"r1"}`, nil)
	_, err := ExtractAs[invoice](context.Background(), c, invoiceSchema, &SyncExtractConfig{Input: DocumentInputFromString("https://x/y.pdf")})
	if err == nil || !strings.Contains(err.Error(), "https://s3/x.json") {
		t.Errorf("url result: %v", err)
	}

	c = newTestClient(t, func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, `{"job_id":"j9"}`) })
	_, err = ExtractAs[invoice](context.Background(), c, invoiceSchema, &SyncExtractConfig{Input: DocumentInputFromString("https://x/y.pdf")})
	if err == nil || !strings.Contains(err.Error(), "queued as job j9") {
		t.Errorf("queued: %v", err)
	}
}

func TestValidateExtractInputs(t *testing.T) {
	var job JobResult
	if err := json.Unmarshal([]byte(strings.Replace(extractBody, "%s", `[{"title":"A"}]`, 1)), &job); err != nil {
		t.Fatal(err)
	}
	out, err := ValidateExtract[invoice](&job)
	if err != nil || out.Result[0].Title != "A" || out.V3ExtractResponse == nil {
		t.Fatalf("job result: %v %+v", err, out)
	}

	var legacyJob JobResult
	if err := json.Unmarshal([]byte(strings.Replace(legacyExtractBody, "%s", `[{"title":"B"}]`, 1)), &legacyJob); err != nil {
		t.Fatal(err)
	}
	out, err = ValidateExtract[invoice](&legacyJob)
	if err != nil || out.Result[0].Title != "B" || out.ExtractResponse == nil || out.Usage().NumFields != 2 {
		t.Fatalf("legacy job result: %v %+v", err, out)
	}

	v3 := V3ExtractResponse{Result: map[string]any{"title": "One"}}
	if out, err = ValidateExtract[invoice](v3); err != nil || len(out.Result) != 1 || out.Result[0].Title != "One" {
		t.Errorf("v3 single object: %v %+v", err, out)
	}

	for name, in := range map[string]any{
		"nil":   nil,
		"parse": &JobResult{ParseResponse: &ParseResponse{}},
		"other": "nope",
	} {
		if _, err := ValidateExtract[invoice](in); err == nil {
			t.Errorf("%s: expected an error", name)
		}
	}
}
