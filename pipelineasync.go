// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reducto

import (
	"context"
	"net/http"
	"slices"

	"github.com/reductoai/reducto-go-sdk/internal/apijson"
	"github.com/reductoai/reducto-go-sdk/internal/param"
	"github.com/reductoai/reducto-go-sdk/internal/requestconfig"
	"github.com/reductoai/reducto-go-sdk/option"
)

// PipelineAsyncService contains methods and other services that help with
// interacting with the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPipelineAsyncService] method instead.
type PipelineAsyncService struct {
	Options []option.RequestOption
}

// NewPipelineAsyncService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPipelineAsyncService(opts ...option.RequestOption) (r *PipelineAsyncService) {
	r = &PipelineAsyncService{}
	r.Options = opts
	return
}

// Pipeline Async
func (r *PipelineAsyncService) New(ctx context.Context, body PipelineAsyncNewParams, opts ...option.RequestOption) (res *PipelineAsyncNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pipeline_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PipelineAsyncNewResponse struct {
	JobID string                       `json:"job_id" api:"required"`
	JSON  pipelineAsyncNewResponseJSON `json:"-"`
}

// pipelineAsyncNewResponseJSON contains the JSON metadata for the struct
// [PipelineAsyncNewResponse]
type pipelineAsyncNewResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineAsyncNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineAsyncNewResponseJSON) RawJSON() string {
	return r.raw
}

type PipelineAsyncNewParams struct {
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input param.Field[PipelineAsyncNewParamsInputUnion] `json:"input" api:"required"`
	// The ID of the pipeline to use for the document.
	PipelineID param.Field[string] `json:"pipeline_id" api:"required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async param.Field[AsyncConfigV3Param] `json:"async"`
	// Settings for pipeline execution that override pipeline defaults.
	Settings param.Field[PipelineSettingsParam] `json:"settings"`
}

func (r PipelineAsyncNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For parse/split/extract pipelines, the URL of the document to be processed. You
// can provide one of the following: 1. A publicly available URL 2. A presigned S3
// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
// directly uploading a document 4. A jobid:// prefixed URL obtained from a
// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
// API only)
//
//	For edit pipelines, this should be a string containing the edit instructions
//
// Satisfied by [shared.UnionString], [PipelineAsyncNewParamsInputArray],
// [UploadResponseParam].
type PipelineAsyncNewParamsInputUnion interface {
	ImplementsPipelineAsyncNewParamsInputUnion()
}

type PipelineAsyncNewParamsInputArray []string

func (r PipelineAsyncNewParamsInputArray) ImplementsPipelineAsyncNewParamsInputUnion() {}
