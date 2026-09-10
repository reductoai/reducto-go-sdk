package reducto

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"
	"time"
)

// Run with: REDUCTO_SMOKE=1 REDUCTO_API_KEY=... go test -run Smoke -v -timeout 20m
func TestSmoke(t *testing.T) {
	if os.Getenv("REDUCTO_SMOKE") == "" {
		t.Skip("set REDUCTO_SMOKE=1 to hit the live API")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()
	c := New("")
	wait := &WaitOptions{Interval: 2 * time.Second, Timeout: 5 * time.Minute}
	const publicPDF = "https://pdfobject.com/pdf/sample.pdf"

	must := func(name string, err error) {
		t.Helper()
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}
	}
	report := func(name string, err error) bool {
		t.Helper()
		if err != nil {
			t.Logf("%-18s FAIL  %v", name, err)
			return false
		}
		t.Logf("%-18s ok", name)
		return true
	}

	// Version
	v, err := c.Version(ctx)
	must("Version", err)
	t.Logf("%-18s ok    %s", "Version", v)

	// Upload
	up, err := c.UploadFile(ctx, "testdata/sample.pdf")
	must("Upload", err)
	t.Logf("%-18s ok    %s", "Upload", up.FileID)

	// Parse (sync, uploaded file)
	po, err := c.Parse(ctx, &SyncParseConfig{Input: up.Input()})
	must("Parse", err)
	if po.ParseResponse == nil {
		t.Fatalf("Parse: got async response %+v", po.AsyncParseResponse)
	}
	pr := po.ParseResponse
	t.Logf("%-18s ok    job=%s pages=%d chunks=%d url=%v", "Parse", pr.JobID, pr.Usage.NumPages, chunkCount(pr.Result), pr.Result.UrlResult != nil)
	parsedJobID := pr.JobID

	// Parse (sync, public URL, with options)
	po2, err := c.Parse(ctx, &SyncParseConfig{
		Input:      DocumentInputFromString(publicPDF),
		Retrieval:  &Retrieval{Chunking: &Chunking{ChunkMode: ChunkingChunkModePage}},
		Formatting: &Formatting{TableOutputFormat: FormattingTableOutputFormatMd},
		Settings:   &Settings{ReturnOCRData: Ptr(true), PageRange: &PageSelection{PageRange: &PageRange{Start: Ptr[int64](1), End: Ptr[int64](1)}}},
	})
	if report("Parse(url+opts)", err) && po2.ParseResponse != nil && po2.ParseResponse.Result.FullResult != nil {
		t.Logf("%-18s       ocr=%v", "", po2.ParseResponse.Result.FullResult.OCR != nil)
	}

	// ParseAsync -> GetJob/WaitForJob
	pa, err := c.ParseAsync(ctx, &AsyncParseConfig{Input: up.Input()})
	must("ParseAsync", err)
	job, err := c.WaitForJob(ctx, pa.JobID, wait)
	must("WaitForJob", err)
	if job.Result == nil || job.Result.ParseResponse == nil {
		t.Fatalf("WaitForJob: unexpected result %s", pretty(job))
	}
	t.Logf("%-18s ok    job=%s status=%s type=%s dur=%v", "ParseAsync+Wait", pa.JobID, job.Status, job.Type, deref(job.Duration))

	// Extract (sync) using jobid:// so no re-parse
	schema := map[string]any{"type": "object", "properties": map[string]any{
		"title":      map[string]any{"type": "string", "description": "document title"},
		"word_count": map[string]any{"type": "integer"},
	}}
	eo, err := c.Extract(ctx, &SyncExtractConfig{
		Input:        DocumentInputFromString("jobid://" + parsedJobID),
		Instructions: &Instructions{Schema: schema},
	})
	if report("Extract", err) {
		switch {
		case eo.V3ExtractResponse != nil:
			t.Logf("%-18s       variant=v3_extract result=%s", "", pretty(eo.V3ExtractResponse.Result))
		default:
			t.Errorf("Extract: decoded as async %+v", eo.AsyncExtractResponse)
		}
	}

	// ExtractAsync
	ea, err := c.ExtractAsync(ctx, &AsyncExtractConfig{Input: up.Input(), Instructions: &Instructions{Schema: schema}})
	if report("ExtractAsync", err) {
		j, err := c.WaitForJob(ctx, ea.JobID, wait)
		if report("ExtractAsync+Wait", err) {
			t.Logf("%-18s       variant=%s", "", jobResultVariant(j.Result))
		}
	}

	// Split (sync + async)
	cats := []SplitCategory{
		{Name: "cover", Description: "The first page or a title page"},
		{Name: "body", Description: "Any other content"},
	}
	so, err := c.Split(ctx, &SyncSplitConfig{Input: DocumentInputFromString("jobid://" + parsedJobID), SplitDescription: cats})
	if report("Split", err) {
		t.Logf("%-18s       result=%s", "", pretty(so.Result))
	}
	sa, err := c.SplitAsync(ctx, &AsyncSplitConfig{Input: up.Input(), SplitDescription: cats})
	if report("SplitAsync", err) {
		j, err := c.WaitForJob(ctx, sa.JobID, wait)
		if report("SplitAsync+Wait", err) {
			t.Logf("%-18s       variant=%s", "", jobResultVariant(j.Result))
		}
	}

	// Classify
	co, err := c.Classify(ctx, &ClassifyConfig{
		Input: up.Input(),
		ClassificationSchema: []ClassificationCategory{
			{Category: "invoice", Criteria: []string{"Contains line items and a total amount"}},
			{Category: "other", Criteria: []string{"Anything that is not an invoice"}},
		},
	})
	if report("Classify", err) {
		t.Logf("%-18s       result=%s", "", pretty(co.Result))
	}

	// Edit (sync + async) — needs a form-like document; expected to work or return a clean APIError
	ed, err := c.Edit(ctx, &EditConfig{
		DocumentURL:      SingleDocumentInputFromUploadResponse(*up),
		EditInstructions: "Write 'SMOKE TEST' in the top margin of the first page.",
	})
	if report("Edit", err) {
		t.Logf("%-18s       %s", "", pretty(ed))
	}
	edA, err := c.EditAsync(ctx, &AsyncEditConfig{
		DocumentURL:      SingleDocumentInputFromUploadResponse(*up),
		EditInstructions: "Write 'SMOKE TEST' in the top margin of the first page.",
	})
	if report("EditAsync", err) {
		_, err := c.WaitForJob(ctx, edA.JobID, wait)
		report("EditAsync+Wait", err)
	}

	// Pipeline — needs a saved pipeline id; we only assert the error is a typed APIError
	_, err = c.Pipeline(ctx, &V3PipelineConfig{Input: up.Input(), PipelineID: "smoke-nonexistent"})
	logAPIError(t, "Pipeline", err)
	_, err = c.PipelineAsync(ctx, &V3AsyncPipelineConfig{Input: up.Input(), PipelineID: "smoke-nonexistent"})
	logAPIError(t, "PipelineAsync", err)

	// CancelJob — start a job and cancel it right away
	toCancel, err := c.ParseAsync(ctx, &AsyncParseConfig{Input: DocumentInputFromString(publicPDF)})
	must("ParseAsync(cancel)", err)
	raw, err := c.CancelJob(ctx, toCancel.JobID)
	if report("CancelJob", err) {
		t.Logf("%-18s       body=%s", "", raw)
	}

	// ListJobs (two pages, small limit)
	jl, err := c.ListJobs(ctx, &ListJobsParams{Limit: Ptr[int64](2), ExcludeConfigs: Ptr(true)})
	if report("ListJobs", err) {
		t.Logf("%-18s       n=%d next=%v first=%s/%s", "", len(jl.Jobs), jl.NextCursor != nil, jl.Jobs[0].Status, jl.Jobs[0].Type)
		if jl.NextCursor != nil {
			_, err := c.ListJobs(ctx, &ListJobsParams{Cursor: jl.NextCursor, Limit: Ptr[int64](2)})
			report("ListJobs(cursor)", err)
		}
	}

	// GetJob on a bogus id -> clean 404-ish APIError
	_, err = c.GetJob(ctx, "00000000-0000-0000-0000-000000000000")
	logAPIError(t, "GetJob(bogus)", err)

	// ConfigureWebhook
	portal, err := c.ConfigureWebhook(ctx)
	if report("ConfigureWebhook", err) {
		t.Logf("%-18s       url=%.60s…", "", portal)
	}

	// DeleteJob, DeleteUpload
	dj, err := c.DeleteJob(ctx, pa.JobID, &DeleteJobParams{IncludePersisted: Ptr(true)})
	if report("DeleteJob", err) {
		t.Logf("%-18s       job=%s", "", dj.JobID)
	}
	du, err := c.DeleteUpload(ctx, up.FileID)
	if report("DeleteUpload", err) {
		t.Logf("%-18s       %s", "", pretty(du))
	}
}

func logAPIError(t *testing.T, name string, err error) {
	t.Helper()
	var apiErr *APIError
	switch {
	case err == nil:
		t.Logf("%-18s ok    (unexpectedly succeeded)", name)
	case errors.As(err, &apiErr):
		t.Logf("%-18s ok    typed APIError status=%d code=%s msg=%.80s", name, apiErr.StatusCode, apiErr.Code(), apiErr.Message)
	default:
		t.Errorf("%-18s FAIL  untyped error: %v", name, err)
	}
}

func chunkCount(r ParseResult) int {
	if r.FullResult == nil {
		return -1
	}
	return len(r.FullResult.Chunks)
}

func jobResultVariant(r *JobResult) string {
	if r == nil {
		return "<nil>"
	}
	b, _ := json.Marshal(r)
	var m struct {
		ResponseType string `json:"response_type"`
	}
	_ = json.Unmarshal(b, &m)
	return m.ResponseType
}

func deref[T any](p *T) any {
	if p == nil {
		return nil
	}
	return *p
}

func pretty(v any) string {
	b, _ := json.Marshal(v)
	if len(b) > 300 {
		return string(b[:300]) + "…"
	}
	return string(b)
}
