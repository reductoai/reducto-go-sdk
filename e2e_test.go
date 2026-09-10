//go:build e2e

package reducto

// End-to-end tests against the live Reducto API. They check that every endpoint answers with
// the expected shape, not the quality of the output. They mirror tests/e2e in the Python SDK.
//
// Run with: REDUCTO_API_KEY=... go test -tags e2e -v -timeout 20m ./...

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

const (
	documentURL        = "https://ci.reducto.ai/onepager.pdf"
	missingDocumentURL = "https://ci.reducto.ai/does-not-exist-e2e.pdf"
)

var trivialSchema = map[string]any{
	"type": "object",
	"properties": map[string]any{
		"title": map[string]any{"type": "string", "description": "The title of the document."},
	},
	"required": []string{"title"},
}

var classificationSchema = []ClassificationCategory{
	{Category: "invoice", Criteria: []string{"an invoice or bill"}},
	{Category: "report", Criteria: []string{"a narrative report"}},
}

func e2eClient(t *testing.T) (*Client, context.Context) {
	t.Helper()
	if os.Getenv("REDUCTO_API_KEY") == "" {
		t.Fatal("REDUCTO_API_KEY environment variable is required for E2E tests")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	t.Cleanup(cancel)
	return New(""), ctx
}

func e2eWait(t *testing.T, c *Client, ctx context.Context, jobID string) *EnhancedAsyncJobResponse {
	t.Helper()
	job, err := c.WaitForJob(ctx, jobID, &WaitOptions{Interval: 2 * time.Second, Timeout: 2 * time.Minute})
	var failed *JobFailedError
	if errors.As(err, &failed) {
		return failed.Job
	}
	if err != nil {
		t.Fatalf("job %s: %v", jobID, err)
	}
	return job
}

func e2eTempPDF(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "minimal.pdf")
	if err := os.WriteFile(path, []byte("%PDF-1.4 minimal test file"), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

func e2eParse(t *testing.T, c *Client, ctx context.Context, req *SyncParseConfig) *ParseResponse {
	t.Helper()
	out, err := c.Parse(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	if out.ParseResponse == nil {
		t.Fatalf("expected a ParseResponse, got %+v", out)
	}
	return out.ParseResponse
}

func e2eExtract(t *testing.T, c *Client, ctx context.Context, req *SyncExtractConfig) *V3ExtractResponse {
	t.Helper()
	out, err := c.Extract(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	if out.V3ExtractResponse == nil {
		t.Fatalf("expected a V3ExtractResponse, got %+v", out)
	}
	return out.V3ExtractResponse
}

func e2eSyncURL(v any) string {
	if m, ok := v.(map[string]any); ok && m["type"] == "url" {
		u, _ := m["url"].(string)
		return u
	}
	return ""
}

func TestE2EParse(t *testing.T) {
	t.Parallel()
	c, ctx := e2eClient(t)

	// One parse call that sends every newer parse option.
	full := e2eParse(t, c, ctx, &SyncParseConfig{
		Input: DocumentInputFromString(documentURL),
		Enhance: &Enhance{
			AdvancedChartAgent: Ptr(true),
			Agentic:            []AgenticConfig{AgenticConfigFromTableAgentic(TableAgentic{Mode: TableAgenticModeDefault})},
		},
		Settings: &Settings{
			EmbedPDFMetadata:          Ptr(true),
			EmbedPDFMetadataDPI:       Ptr[int64](72),
			ExtractDocumentProperties: Ptr(true),
			ForceURLResult:            Ptr(true),
			TenantThrottling:          &TenantThrottling{TenantID: "sdk-e2e", MaxShare: Ptr(0.5)},
		},
		Spreadsheet: &Spreadsheet{MaxCellCount: Ptr[int64](100_000)},
	})

	t.Run("sync", func(t *testing.T) {
		r := e2eParse(t, c, ctx, &SyncParseConfig{Input: DocumentInputFromString(documentURL)})
		if r.JobID == "" || r.Duration < 0 {
			t.Errorf("job_id=%q duration=%v", r.JobID, r.Duration)
		}
		if r.Result.FullResult == nil {
			t.Fatalf("result = %+v", r.Result)
		}
		if len(r.Result.FullResult.Chunks) == 0 {
			t.Error("no chunks")
		}
	})
	t.Run("response_type", func(t *testing.T) {
		if full.ResponseType != "parse" {
			t.Errorf("response_type = %q", full.ResponseType)
		}
	})
	t.Run("force_url_result", func(t *testing.T) {
		if full.Result.UrlResult == nil || !strings.HasPrefix(full.Result.UrlResult.URL, "https://") {
			t.Errorf("result = %+v", full.Result)
		}
	})
	t.Run("document_properties", func(t *testing.T) {
		if full.DocumentProperties == nil || full.DocumentProperties.Title == nil || *full.DocumentProperties.Title == "" {
			t.Errorf("document_properties = %+v", full.DocumentProperties)
		}
	})
	t.Run("embed_pdf_metadata", func(t *testing.T) {
		if full.PDFURL == nil || *full.PDFURL == "" {
			t.Error("pdf_url missing")
		}
	})
	t.Run("usage_page_billing_breakdown", func(t *testing.T) {
		kinds, ok := full.Usage.PageBillingBreakdown["1"]
		if !ok {
			t.Fatalf("page_billing_breakdown = %+v", full.Usage.PageBillingBreakdown)
		}
		found := false
		for _, k := range kinds {
			if k == "page" {
				found = true
			}
		}
		if !found {
			t.Errorf("page 1 kinds = %v", kinds)
		}
	})
}

func TestE2EParseAsync(t *testing.T) {
	t.Parallel()
	c, ctx := e2eClient(t)

	t.Run("returns_job_id", func(t *testing.T) {
		r, err := c.ParseAsync(ctx, &AsyncParseConfig{Input: DocumentInputFromString(documentURL)})
		if err != nil || r.JobID == "" {
			t.Fatalf("err=%v r=%+v", err, r)
		}
	})
	t.Run("job_completes", func(t *testing.T) {
		r, err := c.ParseAsync(ctx, &AsyncParseConfig{Input: DocumentInputFromString(documentURL)})
		if err != nil {
			t.Fatal(err)
		}
		job := e2eWait(t, c, ctx, r.JobID)
		if job.Status == JobStatusFailed {
			t.Fatalf("parse async job failed: %v", deref(job.Reason))
		}
		if job.Result == nil {
			t.Error("result missing")
		}
	})
	t.Run("queue_priority_standard", func(t *testing.T) {
		r, err := c.ParseAsync(ctx, &AsyncParseConfig{Input: DocumentInputFromString(documentURL), QueuePriority: QueuePriorityStandard})
		if err != nil || r.JobID == "" {
			t.Fatalf("err=%v r=%+v", err, r)
		}
	})
}

func TestE2EExtract(t *testing.T) {
	t.Parallel()
	c, ctx := e2eClient(t)

	// One extract call that sends every newer extract option.
	full := e2eExtract(t, c, ctx, &SyncExtractConfig{
		Input:        DocumentInputFromString(documentURL),
		Instructions: &Instructions{Schema: trivialSchema},
		Settings: &ExtractSettings{
			ForceURLResult: Ptr(true),
			PageRange:      &PageSelection{PageRange: &PageRange{Start: Ptr[int64](1), End: Ptr[int64](1)}},
			Citations:      &Citations{Enabled: Ptr(true), ParentBlock: CitationParentBlockModeBboxOnly},
		},
	})

	t.Run("sync", func(t *testing.T) {
		r := e2eExtract(t, c, ctx, &SyncExtractConfig{
			Input:        DocumentInputFromString(documentURL),
			Instructions: &Instructions{Schema: trivialSchema},
		})
		items, ok := r.Result.([]any)
		if !ok || len(items) == 0 {
			t.Errorf("result = %#v", r.Result)
		}
	})
	t.Run("response_type", func(t *testing.T) {
		if full.ResponseType != "v3_extract" {
			t.Errorf("response_type = %q", full.ResponseType)
		}
	})
	t.Run("force_url_result", func(t *testing.T) {
		if u := e2eSyncURL(full.Result); !strings.HasPrefix(u, "https://") {
			t.Errorf("result = %#v", full.Result)
		}
	})
	t.Run("confidence_fields_present", func(t *testing.T) {
		switch full.Confidence {
		case "", "high", "low":
		default:
			t.Errorf("confidence = %q", full.Confidence)
		}
	})
	t.Run("extract_as", func(t *testing.T) {
		type doc struct {
			Title string `json:"title"`
		}
		out, err := ExtractAs[doc](ctx, c, trivialSchema, &SyncExtractConfig{Input: DocumentInputFromString(documentURL)})
		if err != nil {
			t.Fatal(err)
		}
		if len(out.Result) == 0 || out.Result[0].Title == "" {
			t.Errorf("result = %+v", out.Result)
		}
	})
}

func TestE2EExtractAsync(t *testing.T) {
	t.Parallel()
	c, ctx := e2eClient(t)
	req := func() *AsyncExtractConfig {
		return &AsyncExtractConfig{Input: DocumentInputFromString(documentURL), Instructions: &Instructions{Schema: trivialSchema}}
	}

	t.Run("returns_job_id", func(t *testing.T) {
		r, err := c.ExtractAsync(ctx, req())
		if err != nil || r.JobID == "" {
			t.Fatalf("err=%v r=%+v", err, r)
		}
	})
	t.Run("job_completes", func(t *testing.T) {
		r, err := c.ExtractAsync(ctx, req())
		if err != nil {
			t.Fatal(err)
		}
		job := e2eWait(t, c, ctx, r.JobID)
		if job.Status == JobStatusFailed {
			t.Fatalf("extract async job failed: %v", deref(job.Reason))
		}
		if job.Result == nil {
			t.Fatal("result missing")
		}
		if _, err := ValidateExtract[map[string]any](job.Result); err != nil {
			t.Error(err)
		}
	})
	t.Run("queue_priority_standard", func(t *testing.T) {
		body := req()
		body.QueuePriority = QueuePriorityStandard
		r, err := c.ExtractAsync(ctx, body)
		if err != nil || r.JobID == "" {
			t.Fatalf("err=%v r=%+v", err, r)
		}
	})
}

func TestE2ESplit(t *testing.T) {
	t.Parallel()
	c, ctx := e2eClient(t)

	r, err := c.Split(ctx, &SyncSplitConfig{
		Input: DocumentInputFromString(documentURL),
		SplitDescription: []SplitCategory{
			{Name: "invoice", Description: "An invoice or bill"},
			{Name: "other", Description: "Anything else"},
		},
		Settings: &SplitSettings{
			AllowPageOverlap: Ptr(true),
			AutoPartition:    Ptr(true),
			DeepSplit:        Ptr(false),
			ForceURLResult:   Ptr(true),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Run("response_type", func(t *testing.T) {
		if r.ResponseType != "split" {
			t.Errorf("response_type = %q", r.ResponseType)
		}
	})
	t.Run("job_id_and_duration", func(t *testing.T) {
		if r.JobID == nil || *r.JobID == "" || r.Duration == nil || *r.Duration < 0 {
			t.Errorf("job_id=%v duration=%v", deref(r.JobID), deref(r.Duration))
		}
	})
	t.Run("force_url_result", func(t *testing.T) {
		if r.Result.UrlResult == nil || !strings.HasPrefix(r.Result.UrlResult.URL, "https://") {
			t.Errorf("result = %+v", r.Result)
		}
	})
}

func TestE2EClassify(t *testing.T) {
	t.Parallel()
	c, ctx := e2eClient(t)

	t.Run("category_groups_and_model", func(t *testing.T) {
		r, err := c.Classify(ctx, &ClassifyConfig{
			Input:                DocumentInputFromString(documentURL),
			ClassificationSchema: classificationSchema,
			CategoryGroups:       map[string][]string{"financial": {"invoice"}},
			Model:                ClassifyConfigModelDefault,
			Priority:             Ptr(true),
		})
		if err != nil {
			t.Fatal(err)
		}
		if r.ResponseType != "classify" {
			t.Errorf("response_type = %q", r.ResponseType)
		}
		if r.Result.ClassifyResponseCategory == nil {
			t.Fatalf("result = %+v", r.Result)
		}
		if cat := r.Result.ClassifyResponseCategory.Category; cat != "invoice" && cat != "report" {
			t.Errorf("category = %q", cat)
		}
		if g := r.ExtraMetadata["grouping"]; g != "financial" && g != "ungrouped" {
			t.Errorf("grouping = %q (%v)", g, r.ExtraMetadata)
		}
		if r.Usage == nil || r.Usage.NumCategories != int64(len(classificationSchema)) {
			t.Errorf("usage = %+v", r.Usage)
		}
	})
	t.Run("force_url_result", func(t *testing.T) {
		r, err := c.Classify(ctx, &ClassifyConfig{
			Input:                DocumentInputFromString(documentURL),
			ClassificationSchema: classificationSchema,
			ForceURLResult:       Ptr(true),
		})
		if err != nil {
			t.Fatal(err)
		}
		if r.Result.UrlResult == nil || !strings.HasPrefix(r.Result.UrlResult.URL, "https://") {
			t.Errorf("result = %+v", r.Result)
		}
	})
}

func TestE2EUpload(t *testing.T) {
	t.Parallel()
	c, ctx := e2eClient(t)

	t.Run("presign_returns_file_id", func(t *testing.T) {
		up, err := c.PresignUpload(ctx, nil)
		if err != nil {
			t.Fatal(err)
		}
		if up.FileID == "" || up.PresignedURL == nil || *up.PresignedURL == "" {
			t.Errorf("up = %+v", up)
		}
	})
	t.Run("presign_with_extension", func(t *testing.T) {
		up, err := c.PresignUpload(ctx, &UploadOptions{Extension: "pdf"})
		if err != nil {
			t.Fatal(err)
		}
		if !strings.HasPrefix(up.FileID, "reducto://") || up.PresignedURL == nil {
			t.Errorf("up = %+v", up)
		}
	})
	t.Run("upload_file", func(t *testing.T) {
		up, err := c.UploadFile(ctx, e2eTempPDF(t))
		if err != nil {
			t.Fatal(err)
		}
		if !strings.HasPrefix(up.FileID, "reducto://") {
			t.Errorf("file_id = %q", up.FileID)
		}
	})
	t.Run("delete_uploaded_file", func(t *testing.T) {
		up, err := c.UploadFile(ctx, e2eTempPDF(t))
		if err != nil {
			t.Fatal(err)
		}
		deleted, err := c.DeleteUpload(ctx, up.FileID)
		if err != nil {
			t.Fatal(err)
		}
		if deleted.FileID != up.FileID {
			t.Errorf("deleted %q, uploaded %q", deleted.FileID, up.FileID)
		}
		_, err = c.DeleteUpload(ctx, up.FileID)
		var apiErr *APIError
		if !errors.As(err, &apiErr) || apiErr.StatusCode != 404 {
			t.Errorf("second delete: %v", err)
		}
	})
}

func TestE2EJob(t *testing.T) {
	t.Parallel()
	c, ctx := e2eClient(t)
	start := func(t *testing.T, input string) string {
		t.Helper()
		r, err := c.ParseAsync(ctx, &AsyncParseConfig{Input: DocumentInputFromString(input)})
		if err != nil {
			t.Fatal(err)
		}
		return r.JobID
	}

	t.Run("get_with_parse_job", func(t *testing.T) {
		job, err := c.GetJob(ctx, start(t, documentURL))
		if err != nil {
			t.Fatal(err)
		}
		switch job.Status {
		case JobStatusPending, JobStatusCompleted, JobStatusIdle:
		default:
			t.Errorf("status = %q", job.Status)
		}
	})
	t.Run("get_completed", func(t *testing.T) {
		job := e2eWait(t, c, ctx, start(t, documentURL))
		if job.Status == JobStatusFailed {
			t.Fatalf("job failed: %v", deref(job.Reason))
		}
		if job.Result == nil || job.Result.ParseResponse == nil {
			t.Errorf("result = %+v", job.Result)
		}
	})
	t.Run("delete", func(t *testing.T) {
		jobID := start(t, documentURL)
		if job := e2eWait(t, c, ctx, jobID); job.Status != JobStatusCompleted {
			t.Fatalf("status = %q", job.Status)
		}
		deleted, err := c.DeleteJob(ctx, jobID, nil)
		if err != nil {
			t.Fatal(err)
		}
		if deleted.JobID != jobID {
			t.Errorf("deleted %q, started %q", deleted.JobID, jobID)
		}
		_, err = c.GetJob(ctx, jobID)
		var apiErr *APIError
		if !errors.As(err, &apiErr) || (apiErr.StatusCode != 409 && apiErr.StatusCode != 410) {
			t.Errorf("get after delete: %v", err)
		}
	})
	t.Run("delete_with_include_persisted", func(t *testing.T) {
		jobID := start(t, documentURL)
		e2eWait(t, c, ctx, jobID)
		deleted, err := c.DeleteJob(ctx, jobID, &DeleteJobParams{IncludePersisted: Ptr(true)})
		if err != nil {
			t.Fatal(err)
		}
		if deleted.JobID != jobID {
			t.Errorf("deleted %q, started %q", deleted.JobID, jobID)
		}
	})
	t.Run("failed_job_has_error_detail", func(t *testing.T) {
		job := e2eWait(t, c, ctx, start(t, missingDocumentURL))
		if job.Status != JobStatusFailed {
			t.Fatalf("status = %q", job.Status)
		}
		if job.Error == nil {
			t.Fatal("error missing")
		}
		if job.Error.Name != ErrorCodeInvalidConfig || job.Error.Code != 400 || job.Error.Message == "" {
			t.Errorf("error = %+v", job.Error)
		}
	})
	t.Run("list_and_cancel", func(t *testing.T) {
		jobID := start(t, documentURL)
		if _, err := c.CancelJob(ctx, jobID); err != nil {
			t.Fatal(err)
		}
		list, err := c.ListJobs(ctx, &ListJobsParams{Limit: Ptr[int64](2), ExcludeConfigs: Ptr(true)})
		if err != nil {
			t.Fatal(err)
		}
		if len(list.Jobs) == 0 {
			t.Error("no jobs listed")
		}
	})
}

func TestE2EMisc(t *testing.T) {
	t.Parallel()
	c, ctx := e2eClient(t)

	t.Run("version", func(t *testing.T) {
		v, err := c.Version(ctx)
		if err != nil || v == "" {
			t.Fatalf("err=%v v=%q", err, v)
		}
	})
	t.Run("configure_webhook", func(t *testing.T) {
		u, err := c.ConfigureWebhook(ctx)
		if err != nil || !strings.HasPrefix(u, "https://") {
			t.Fatalf("err=%v url=%q", err, u)
		}
	})
}

func deref[T any](p *T) any {
	if p == nil {
		return nil
	}
	return *p
}
