// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reducto

import (
	"context"
	"net/http"

	"github.com/reductoai/reducto-go-sdk/internal/apijson"
	"github.com/reductoai/reducto-go-sdk/internal/param"
	"github.com/reductoai/reducto-go-sdk/internal/requestconfig"
	"github.com/reductoai/reducto-go-sdk/option"
	"github.com/reductoai/reducto-go-sdk/shared"
)

// PipelineService contains methods and other services that help with interacting
// with the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPipelineService] method instead.
type PipelineService struct {
	Options []option.RequestOption
}

// NewPipelineService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPipelineService(opts ...option.RequestOption) (r *PipelineService) {
	r = &PipelineService{}
	r.Options = opts
	return
}

// Pipeline
func (r *PipelineService) Run(ctx context.Context, body PipelineRunParams, opts ...option.RequestOption) (res *shared.PipelineResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "pipeline"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Pipeline Async
func (r *PipelineService) RunJob(ctx context.Context, body PipelineRunJobParams, opts ...option.RequestOption) (res *PipelineRunJobResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "pipeline_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PipelineRunJobResponse struct {
	JobID string                     `json:"job_id,required"`
	JSON  pipelineRunJobResponseJSON `json:"-"`
}

// pipelineRunJobResponseJSON contains the JSON metadata for the struct
// [PipelineRunJobResponse]
type pipelineRunJobResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineRunJobResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineRunJobResponseJSON) RawJSON() string {
	return r.raw
}

type PipelineRunParams struct {
	// The URL of the document to be processed. You can provide one of the
	// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
	// prefixed URL obtained from the /upload endpoint after directly uploading a
	// document
	DocumentURL param.Field[PipelineRunParamsDocumentURLUnion] `json:"document_url,required"`
	// The ID of the pipeline to use for the document.
	PipelineID param.Field[string] `json:"pipeline_id,required"`
}

func (r PipelineRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The URL of the document to be processed. You can provide one of the
// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
// prefixed URL obtained from the /upload endpoint after directly uploading a
// document
//
// Satisfied by [shared.UnionString], [shared.UploadParam].
type PipelineRunParamsDocumentURLUnion interface {
	ImplementsPipelineRunParamsDocumentURLUnion()
}

type PipelineRunJobParams struct {
	// The URL of the document to be processed. You can provide one of the
	// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
	// prefixed URL obtained from the /upload endpoint after directly uploading a
	// document
	DocumentURL param.Field[PipelineRunJobParamsDocumentURLUnion] `json:"document_url,required"`
	// The ID of the pipeline to use for the document.
	PipelineID param.Field[string] `json:"pipeline_id,required"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool]                         `json:"priority"`
	Webhook  param.Field[shared.WebhookConfigNewParam] `json:"webhook"`
}

func (r PipelineRunJobParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The URL of the document to be processed. You can provide one of the
// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
// prefixed URL obtained from the /upload endpoint after directly uploading a
// document
//
// Satisfied by [shared.UnionString], [shared.UploadParam].
type PipelineRunJobParamsDocumentURLUnion interface {
	ImplementsPipelineRunJobParamsDocumentURLUnion()
}
