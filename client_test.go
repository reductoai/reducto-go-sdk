package reducto

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func newTestClient(t *testing.T, h http.HandlerFunc) *Client {
	t.Helper()
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)
	c := New("test-key", WithBaseURL(srv.URL), WithMaxRetries(0))
	c.retryBase = time.Millisecond
	return c
}

func TestParseSyncDecodesFullResult(t *testing.T) {
	var gotBody map[string]any
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Errorf("auth header = %q", r.Header.Get("Authorization"))
		}
		if r.URL.Path != "/parse" {
			t.Errorf("path = %s", r.URL.Path)
		}
		_ = json.NewDecoder(r.Body).Decode(&gotBody)
		io.WriteString(w, `{
			"response_type": "parse", "job_id": "j1", "duration": 1.5,
			"usage": {"num_pages": 2},
			"result": {"type": "full", "chunks": [{"content": "hi", "embed": "hi", "blocks": []}]}
		}`)
	})
	out, err := c.Parse(context.Background(), &SyncParseConfig{
		Input:    DocumentInputFromString("https://example.com/a.pdf"),
		Settings: &Settings{PersistResults: Ptr(false), OCRSystem: SettingsOCRSystemStandard},
	})
	if err != nil {
		t.Fatal(err)
	}
	if gotBody["input"] != "https://example.com/a.pdf" {
		t.Errorf("input sent as %v", gotBody["input"])
	}
	if settings := gotBody["settings"].(map[string]any); settings["persist_results"] != false {
		t.Errorf("explicit false was dropped: %v", settings)
	}
	if out.ParseResponse == nil || out.AsyncParseResponse != nil {
		t.Fatalf("wrong variant: %+v", out)
	}
	full := out.ParseResponse.Result.FullResult
	if full == nil || len(full.Chunks) != 1 || full.Chunks[0].Content != "hi" {
		t.Fatalf("result = %+v", out.ParseResponse.Result)
	}
}

func TestParseSyncDecodesAsyncFallback(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"job_id": "j2"}`)
	})
	out, err := c.Parse(context.Background(), &SyncParseConfig{Input: DocumentInputFromString("x")})
	if err != nil {
		t.Fatal(err)
	}
	if out.AsyncParseResponse == nil || out.AsyncParseResponse.JobID != "j2" {
		t.Fatalf("got %+v", out)
	}
}

func TestGetJobDiscriminatedResult(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/job/j%2F3" {
			t.Errorf("path = %s", r.RequestURI)
		}
		io.WriteString(w, `{
			"status": "Completed", "type": "Extract", "num_pages": 3,
			"result": {"response_type": "v3_extract", "job_id": "j3", "usage": {"num_pages": 3, "credits": 1}, "result": {"a": 1}}
		}`)
	})
	job, err := c.GetJob(context.Background(), "j/3")
	if err != nil {
		t.Fatal(err)
	}
	if job.Status != JobStatusCompleted || job.Type != JobTypeExtract {
		t.Fatalf("job = %+v", job)
	}
	if job.Result == nil || job.Result.V3ExtractResponse == nil {
		t.Fatalf("result = %+v", job.Result)
	}
}

func TestConstFieldsAreSetOnMarshal(t *testing.T) {
	b, err := json.Marshal(Enhance{Agentic: []AgenticConfig{
		AgenticConfigFromTableAgentic(TableAgentic{Mode: TableAgenticModeMax}),
		AgenticConfigFromTextAgentic(TextAgentic{}),
	}})
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	if !strings.Contains(s, `"scope":"table"`) || !strings.Contains(s, `"scope":"text"`) {
		t.Fatalf("consts missing: %s", s)
	}
	var back Enhance
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if back.Agentic[0].TableAgentic == nil || back.Agentic[1].TextAgentic == nil {
		t.Fatalf("round trip lost variants: %s", s)
	}
}

func TestPageSelectionVariants(t *testing.T) {
	cases := map[string]func(p PageSelection) bool{
		`{"start": 1, "end": 3}`:       func(p PageSelection) bool { return p.PageRange != nil && *p.PageRange.End == 3 },
		`[{"start": 1}, {"start": 5}]`: func(p PageSelection) bool { return len(p.PageRangeArray) == 2 },
		`[1, 2, 3]`:                    func(p PageSelection) bool { return len(p.IntArray) == 3 },
		`["Sheet1"]`:                   func(p PageSelection) bool { return len(p.StringArray) == 1 },
	}
	for in, ok := range cases {
		var p PageSelection
		if err := json.Unmarshal([]byte(in), &p); err != nil {
			t.Fatalf("%s: %v", in, err)
		}
		if !ok(p) {
			t.Errorf("%s decoded to %+v", in, p)
		}
	}
}

func TestSplitResultRequiredKeysDisambiguate(t *testing.T) {
	var u SplitResponseResult
	if err := json.Unmarshal([]byte(`{"splits": [{"name": "a", "pages": [{"page_number": 1, "evidence": "title page"}]}]}`), &u); err != nil {
		t.Fatal(err)
	}
	if u.DeepSplitResult == nil || u.SplitResult != nil {
		t.Fatalf("expected DeepSplitResult, got %+v", u)
	}
}

func TestListJobsQueryAndStringResponse(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/jobs":
			if q := r.URL.Query(); q.Get("limit") != "5" || q.Get("exclude_configs") != "true" || q.Has("cursor") {
				t.Errorf("query = %s", r.URL.RawQuery)
			}
			io.WriteString(w, `{"jobs": [{"job_id": "a", "status": "InProgress", "type": "Parse", "created_at": "2024-01-01T00:00:00"}], "next_cursor": null}`)
		case "/version":
			io.WriteString(w, `"v1.2.3"`)
		}
	})
	jobs, err := c.ListJobs(context.Background(), &ListJobsParams{Limit: Ptr[int64](5), ExcludeConfigs: Ptr(true)})
	if err != nil {
		t.Fatal(err)
	}
	if len(jobs.Jobs) != 1 || jobs.Jobs[0].Status != JobListStatusInProgress || jobs.NextCursor != nil {
		t.Fatalf("jobs = %+v", jobs)
	}
	v, err := c.Version(context.Background())
	if err != nil || v != "v1.2.3" {
		t.Fatalf("version = %q, %v", v, err)
	}
}

func TestAPIErrorShapes(t *testing.T) {
	bodies := []string{
		`{"error": {"code": 422, "name": "INVALID_CONFIG", "message": "bad input"}}`,
		`{"detail": {"code": 422, "name": "INVALID_CONFIG", "message": "bad input"}}`,
	}
	for _, body := range bodies {
		c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(422)
			io.WriteString(w, body)
		})
		_, err := c.Version(context.Background())
		var apiErr *APIError
		if !errors.As(err, &apiErr) || apiErr.StatusCode != 422 || apiErr.Code() != ErrorCodeInvalidConfig {
			t.Fatalf("%s -> %v", body, err)
		}
	}
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(422)
		io.WriteString(w, `{"detail": [{"loc": ["body", 0], "msg": "field required", "type": "missing"}]}`)
	})
	_, err := c.Version(context.Background())
	var apiErr *APIError
	if !errors.As(err, &apiErr) || len(apiErr.Validation) != 1 || apiErr.Validation[0].Loc[1].Int == nil {
		t.Fatalf("validation error = %v", err)
	}
}

func TestRetriesOn5xxThenSucceeds(t *testing.T) {
	var calls atomic.Int32
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if calls.Add(1) < 3 {
			w.WriteHeader(503)
			return
		}
		io.WriteString(w, `"ok"`)
	})
	c.maxRetries = 3
	if _, err := c.Version(context.Background()); err != nil {
		t.Fatal(err)
	}
	if calls.Load() != 3 {
		t.Fatalf("calls = %d", calls.Load())
	}
}

func TestWithOptionsOverridesRetries(t *testing.T) {
	var calls atomic.Int32
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		w.WriteHeader(503)
	})
	c.maxRetries = 3
	if _, err := c.WithOptions(WithMaxRetries(0)).Version(context.Background()); err == nil {
		t.Fatal("expected error")
	}
	if calls.Load() != 1 {
		t.Fatalf("calls = %d", calls.Load())
	}
	if c.maxRetries != 3 {
		t.Fatalf("original mutated: maxRetries = %d", c.maxRetries)
	}
}

func TestUploadMultipart(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("extension") != "pdf" {
			t.Errorf("extension = %q", r.URL.Query().Get("extension"))
		}
		f, hdr, err := r.FormFile("file")
		if err != nil {
			t.Fatal(err)
		}
		data, _ := io.ReadAll(f)
		if hdr.Filename != "doc.pdf" || string(data) != "%PDF" {
			t.Errorf("file = %s %q", hdr.Filename, data)
		}
		io.WriteString(w, `{"file_id": "reducto://abc"}`)
	})
	up, err := c.Upload(context.Background(), strings.NewReader("%PDF"), "doc.pdf", &UploadOptions{Extension: ".pdf"})
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(SyncParseConfig{Input: up.Input()})
	if !strings.Contains(string(b), `"input":{"file_id":"reducto://abc"}`) {
		t.Fatalf("body = %s", b)
	}
}

func TestWaitForJob(t *testing.T) {
	var calls atomic.Int32
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if calls.Add(1) < 3 {
			io.WriteString(w, `{"status": "Pending", "progress": 0.5}`)
			return
		}
		io.WriteString(w, `{"status": "Failed", "error": {"code": 500, "name": "DOCUMENT_CORRUPT", "message": "nope"}}`)
	})
	job, err := c.WaitForJob(context.Background(), "j", &WaitOptions{Interval: time.Millisecond})
	var failed *JobFailedError
	if !errors.As(err, &failed) || job.Error.Name != ErrorCodeDocumentCorrupt {
		t.Fatalf("err = %v, job = %+v", err, job)
	}
	if !strings.Contains(err.Error(), "DOCUMENT_CORRUPT") {
		t.Fatal(err)
	}
}

func TestClientHeaders(t *testing.T) {
	var got http.Header
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Clone()
		io.WriteString(w, `"1.0"`)
	})
	call := func(opts ...Option) http.Header {
		if _, err := c.Version(context.Background(), opts...); err != nil {
			t.Fatal(err)
		}
		return got
	}

	h := call()
	if ua := h.Get("User-Agent"); ua != "Reducto/Go "+Version {
		t.Errorf("default User-Agent = %q", ua)
	}
	want := map[string]string{
		"X-Reducto-Lang":            "go",
		"X-Reducto-Package-Version": Version,
		"X-Reducto-OS":              platformOS(),
		"X-Reducto-Arch":            platformArch(),
		"X-Reducto-Runtime":         "go",
		"X-Reducto-Runtime-Version": strings.TrimPrefix(runtime.Version(), "go"),
		"X-Reducto-Async":           "false",
		"X-Reducto-Retry-Count":     "0",
		"X-Reducto-Read-Timeout":    "3600",
	}
	for k, v := range want {
		if h.Get(k) != v {
			t.Errorf("%s = %q, want %q", k, h.Get(k), v)
		}
	}
	if h.Get("X-Reducto-Client") != "" {
		t.Errorf("X-Reducto-Client sent without WithAppInfo: %q", h.Get("X-Reducto-Client"))
	}
	if platformOS() == "MacOS" && runtime.GOOS != "darwin" || platformArch() == "x64" && runtime.GOARCH != "amd64" {
		t.Errorf("platform mapping wrong: %s %s", platformOS(), platformArch())
	}

	h = call(WithAppInfo("reducto-cli", "9.9"), WithTimeout(0))
	if ua := h.Get("User-Agent"); ua != "reducto-cli/9.9 Reducto/Go "+Version {
		t.Errorf("app User-Agent = %q", ua)
	}
	if h.Get("X-Reducto-Client") != "reducto-cli" || h.Get("X-Reducto-Client-Version") != "9.9" || h.Get("X-Reducto-Lang") != "go" {
		t.Errorf("app client headers = %v", h)
	}
	if h.Get("X-Reducto-Read-Timeout") != "" {
		t.Errorf("read timeout sent with timeout disabled: %q", h.Get("X-Reducto-Read-Timeout"))
	}

	h = call(WithAppInfo("reducto-cli", "9.9"), WithClientInfo(false))
	if h.Get("X-Reducto-Lang") != "" || h.Get("X-Reducto-Client") != "" || h.Get("X-Reducto-OS") != "" {
		t.Errorf("client info still sent: %v", h)
	}
	if h.Get("User-Agent") != "reducto-cli/9.9 Reducto/Go "+Version {
		t.Errorf("User-Agent dropped: %q", h.Get("User-Agent"))
	}

	h = call(WithUserAgent("custom/1"), WithHeader("X-Reducto-OS", "plan9"))
	if h.Get("User-Agent") != "custom/1" || h.Get("X-Reducto-OS") != "plan9" {
		t.Errorf("overrides lost: %v", h)
	}
}

func TestRetryCountHeader(t *testing.T) {
	var counts []string
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		counts = append(counts, r.Header.Get("X-Reducto-Retry-Count"))
		if len(counts) < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		io.WriteString(w, `"1.0"`)
	})
	if _, err := c.Version(context.Background(), WithMaxRetries(2)); err != nil {
		t.Fatal(err)
	}
	if strings.Join(counts, ",") != "0,1,2" {
		t.Errorf("retry counts = %v", counts)
	}
}

func TestResponseIntoCarriesBody(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"job_id": "j1", "brand_new_field": 7}`)
	})
	var resp Response
	out, err := c.ParseAsync(context.Background(), &AsyncParseConfig{Input: DocumentInputFromString("x")}, WithResponseInto(&resp))
	if err != nil || out.JobID != "j1" {
		t.Fatal(err, out)
	}
	var extra struct {
		BrandNew int `json:"brand_new_field"`
	}
	if json.Unmarshal(resp.Body, &extra) != nil || extra.BrandNew != 7 {
		t.Fatalf("body = %s", resp.Body)
	}
}

func TestShouldRetryHeader(t *testing.T) {
	var calls atomic.Int32
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		switch calls.Add(1) {
		case 1:
			w.Header().Set("x-should-retry", "true")
			w.WriteHeader(http.StatusBadRequest)
		case 2:
			w.Header().Set("x-should-retry", "false")
			w.WriteHeader(http.StatusServiceUnavailable)
		default:
			io.WriteString(w, `"1.0"`)
		}
	})
	_, err := c.Version(context.Background(), WithMaxRetries(5))
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.StatusCode != 503 || calls.Load() != 2 {
		t.Fatalf("err = %v, calls = %d", err, calls.Load())
	}
}

func TestRetryAfterMs(t *testing.T) {
	e := &APIError{Header: http.Header{"Retry-After-Ms": {"250"}, "Retry-After": {"30"}}}
	if d, ok := e.RetryAfter(); !ok || d != 250*time.Millisecond {
		t.Fatalf("d = %v ok = %v", d, ok)
	}
	e = &APIError{Header: http.Header{"Retry-After": {"1.5"}}}
	if d, ok := e.RetryAfter(); !ok || d != 1500*time.Millisecond {
		t.Fatalf("d = %v ok = %v", d, ok)
	}
}

func TestUnionKeepsUnknownShape(t *testing.T) {
	var r ParseResult
	if err := json.Unmarshal([]byte(`{"type": "chunked_v2", "parts": 3}`), &r); err != nil {
		t.Fatal(err)
	}
	if r.FullResult != nil || r.UrlResult != nil || !strings.Contains(string(r.Unknown), "chunked_v2") {
		t.Fatalf("r = %+v", r)
	}
	b, _ := json.Marshal(r)
	if !strings.Contains(string(b), "chunked_v2") {
		t.Fatalf("round trip lost the raw JSON: %s", b)
	}
	var sel PageSelection
	if err := json.Unmarshal([]byte(`"1-3"`), &sel); err != nil || string(sel.Unknown) != `"1-3"` {
		t.Fatalf("string into a union with no string variant: %v %+v", err, sel)
	}
}
