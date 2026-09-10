package reducto

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

var (
	_ = json.RawMessage{}
	_ = strconv.Itoa
	_ = url.PathEscape
)

// CancelJob: Cancel a job that has not finished. A job that is already done or failed is not
// changed.
//
// jobID: The job ID returned by an async operation.
//
// POST /cancel/{job_id}
func (c *Client) CancelJob(ctx context.Context, jobID string, opts ...Option) (json.RawMessage, error) {
	var out json.RawMessage
	if err := c.do(ctx, "POST", "/cancel/"+url.PathEscape(jobID)+"", nil, nil, &out, opts); err != nil {
		return nil, err
	}
	return out, nil
}

// Classify a document into one of the categories you define. Each category has a name and the
// criteria that a matching document meets. The response is the chosen category with a confidence
// for each criterion.
//
// Only the first pages are read, up to the page range you set (at most 10 pages).
//
// POST /classify
func (c *Client) Classify(ctx context.Context, req *ClassifyConfig, opts ...Option) (*ClassifyResponse, error) {
	var out ClassifyResponse
	if err := c.do(ctx, "POST", "/classify", nil, req, &out, opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// ConfigureWebhook: Create a URL of the Svix webhook portal for your organization. Open the URL in
// a browser to add and manage webhook endpoints.
//
// POST /configure_webhook
func (c *Client) ConfigureWebhook(ctx context.Context, opts ...Option) (string, error) {
	var out string
	if err := c.do(ctx, "POST", "/configure_webhook", nil, nil, &out, opts); err != nil {
		return "", err
	}
	return out, nil
}

// DeleteJobParams are the query parameters for Client.DeleteJob.
type DeleteJobParams struct {
	// IncludePersisted: Also delete long-retention persisted artifacts for this job. Default: false.
	IncludePersisted *bool
}

func (p *DeleteJobParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.IncludePersisted != nil {
		q.Set("include_persisted", strconv.FormatBool(*p.IncludePersisted))
	}
	return q
}

// DeleteJob: Asynchronously delete a job's stored artifacts.
//
// Tags the job with the deletion marker. Retrieval returns 409 until artifact cleanup finishes,
// then 410 once the deletion has completed.
//
// jobID: The job ID returned by an async operation.
//
// DELETE /job/{job_id}
func (c *Client) DeleteJob(ctx context.Context, jobID string, params *DeleteJobParams, opts ...Option) (*DeleteJobResponse, error) {
	var out DeleteJobResponse
	if err := c.do(ctx, "DELETE", "/job/"+url.PathEscape(jobID)+"", params.query(), nil, &out, opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteUpload: Delete a previously uploaded file.
//
// Removes the stored file for the given `reducto://` file ID. Deletion is immediate and permanent,
// and returns 404 if the file does not exist or was not uploaded by your organization. Jobs that
// already ran against the file are unaffected; delete those separately with `DELETE
// /job/{job_id}`.
//
// DELETE /upload/{file_id}
func (c *Client) DeleteUpload(ctx context.Context, fileID string, opts ...Option) (*DeleteUploadResponse, error) {
	var out DeleteUploadResponse
	if err := c.do(ctx, "DELETE", "/upload/"+url.PathEscape(fileID)+"", nil, nil, &out, opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// Edit changes a document with natural language instructions: fill form fields, add text, sign, or
// change content. The result is a URL of the edited document.
//
// The call waits until the document is processed. Use EditAsync to get a job ID at once, then
// collect the result with GetJob.
//
// POST /edit
func (c *Client) Edit(ctx context.Context, req *EditConfig, opts ...Option) (*EditResponse, error) {
	var out EditResponse
	if err := c.do(ctx, "POST", "/edit", nil, req, &out, opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// EditAsync: Start an edit job and return at once with a job ID. Collect the result with the job
// lookup, or set a webhook to be notified when the job is done.
//
// POST /edit_async
func (c *Client) EditAsync(ctx context.Context, req *AsyncEditConfig, opts ...Option) (*AsyncEditResponse, error) {
	var out AsyncEditResponse
	if err := c.do(ctx, "POST", "/edit_async", nil, req, &out, opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// Extract returns structured JSON from a document. Give it a JSON schema or a natural language
// description of the fields you want. Pass a jobid:// URL from an earlier parse to skip parsing
// and only pay for extraction.
//
// The call waits until the document is processed. Use ExtractAsync to get a job ID at once, then
// collect the result with GetJob.
//
// POST /extract
func (c *Client) Extract(ctx context.Context, req *SyncExtractConfig, opts ...Option) (*ExtractOutput, error) {
	var out ExtractOutput
	if err := c.do(ctx, "POST", "/extract", nil, req, &out, opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// ExtractAsync: Start an extract job and return at once with a job ID. Collect the result with the
// job lookup, or set a webhook in the async option to be notified when the job is done.
//
// POST /extract_async
func (c *Client) ExtractAsync(ctx context.Context, req *AsyncExtractConfig, opts ...Option) (*AsyncExtractResponse, error) {
	var out AsyncExtractResponse
	if err := c.do(ctx, "POST", "/extract_async", nil, req, &out, opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetJob: Get the status of a job and, when it is done, its result. The result type matches the
// operation that started the job: parse, extract, split, edit, pipeline or classify.
//
// Returns 409 while a delete is in progress and 410 after the job is deleted.
//
// jobID: The job ID returned by an async operation.
//
// GET /job/{job_id}
func (c *Client) GetJob(ctx context.Context, jobID string, opts ...Option) (*EnhancedAsyncJobResponse, error) {
	var out EnhancedAsyncJobResponse
	if err := c.do(ctx, "GET", "/job/"+url.PathEscape(jobID)+"", nil, nil, &out, opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListJobsParams are the query parameters for Client.ListJobs.
type ListJobsParams struct {
	// ExcludeConfigs: Exclude raw_config from response to reduce size. Default: false.
	ExcludeConfigs *bool
	// Cursor for pagination. Use the next_cursor from the previous response to fetch the next page.
	Cursor *string
	// Limit: Maximum number of jobs to return per page. Defaults to 100, max 500. Default: 100.
	Limit *int64
}

func (p *ListJobsParams) query() url.Values {
	q := url.Values{}
	if p == nil {
		return q
	}
	if p.ExcludeConfigs != nil {
		q.Set("exclude_configs", strconv.FormatBool(*p.ExcludeConfigs))
	}
	if p.Cursor != nil {
		q.Set("cursor", *p.Cursor)
	}
	if p.Limit != nil {
		q.Set("limit", strconv.FormatInt(*p.Limit, 10))
	}
	return q
}

// ListJobs: List your recent jobs, newest first. The response is one page and a cursor. Pass the
// cursor back to get the next page.
//
// GET /jobs
func (c *Client) ListJobs(ctx context.Context, params *ListJobsParams, opts ...Option) (*JobsResponse, error) {
	var out JobsResponse
	if err := c.do(ctx, "GET", "/jobs", params.query(), nil, &out, opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// Parse waits until the document is processed and returns the full result: text, tables, figures
// and their position on each page, plus markdown and retrieval chunks.
//
// Use ParseAsync to get a job ID at once, then collect the result with GetJob.
//
// POST /parse
func (c *Client) Parse(ctx context.Context, req *SyncParseConfig, opts ...Option) (*ParseOutput, error) {
	var out ParseOutput
	if err := c.do(ctx, "POST", "/parse", nil, req, &out, opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// ParseAsync: Start a parse job and return at once with a job ID. Collect the result with the job
// lookup, or set a webhook in the async option to be notified when the job is done.
//
// POST /parse_async
func (c *Client) ParseAsync(ctx context.Context, req *AsyncParseConfig, opts ...Option) (*AsyncParseResponse, error) {
	var out AsyncParseResponse
	if err := c.do(ctx, "POST", "/parse_async", nil, req, &out, opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// Pipeline runs a pipeline that you saved in Reducto Studio. The pipeline ID selects the steps and
// their settings, so the request only needs the document.
//
// The call waits until the pipeline is done. Use PipelineAsync to get a job ID at once, then
// collect the result with GetJob.
//
// POST /pipeline
func (c *Client) Pipeline(ctx context.Context, req *V3PipelineConfig, opts ...Option) (*PipelineResponse, error) {
	var out PipelineResponse
	if err := c.do(ctx, "POST", "/pipeline", nil, req, &out, opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// PipelineAsync: Start a saved pipeline as a job and return at once with a job ID. Collect the
// result with the job lookup, or set a webhook in the async option to be notified when the job is
// done.
//
// POST /pipeline_async
func (c *Client) PipelineAsync(ctx context.Context, req *V3AsyncPipelineConfig, opts ...Option) (*AsyncPipelineResponse, error) {
	var out AsyncPipelineResponse
	if err := c.do(ctx, "POST", "/pipeline_async", nil, req, &out, opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// Split finds named sections in a document. Describe each section you expect, and the API returns
// the page ranges that match each description.
//
// The call waits until the document is processed. Use SplitAsync to get a job ID at once, then
// collect the result with GetJob.
//
// POST /split
func (c *Client) Split(ctx context.Context, req *SyncSplitConfig, opts ...Option) (*SplitResponse, error) {
	var out SplitResponse
	if err := c.do(ctx, "POST", "/split", nil, req, &out, opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// SplitAsync: Start a split job and return at once with a job ID. Collect the result with the job
// lookup, or set a webhook in the async option to be notified when the job is done.
//
// POST /split_async
func (c *Client) SplitAsync(ctx context.Context, req *AsyncSplitConfig, opts ...Option) (*AsyncSplitResponse, error) {
	var out AsyncSplitResponse
	if err := c.do(ctx, "POST", "/split_async", nil, req, &out, opts); err != nil {
		return nil, err
	}
	return &out, nil
}

// Version: Get the version string of the Reducto API.
//
// GET /version
func (c *Client) Version(ctx context.Context, opts ...Option) (string, error) {
	var out string
	if err := c.do(ctx, "GET", "/version", nil, nil, &out, opts); err != nil {
		return "", err
	}
	return out, nil
}
