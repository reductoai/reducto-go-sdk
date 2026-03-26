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
	opts = slices.Concat(r.Options, opts)
	path := "pipeline"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Pipeline Async
func (r *PipelineService) RunJob(ctx context.Context, body PipelineRunJobParams, opts ...option.RequestOption) (res *shared.AsyncPipelineResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pipeline_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Settings for pipeline execution that override pipeline defaults.
type PipelineSettingsParam struct {
	// Password to decrypt password-protected documents.
	DocumentPassword param.Field[string] `json:"document_password"`
}

func (r PipelineSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineRunParams struct {
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input param.Field[PipelineRunParamsInputUnion] `json:"input" api:"required"`
	// The ID of the pipeline to use for the document.
	PipelineID param.Field[string] `json:"pipeline_id" api:"required"`
	// Settings for pipeline execution that override pipeline defaults.
	Settings param.Field[PipelineSettingsParam] `json:"settings"`
}

func (r PipelineRunParams) MarshalJSON() (data []byte, err error) {
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
// Satisfied by [shared.UnionString], [PipelineRunParamsInputArray],
// [shared.UploadParam].
type PipelineRunParamsInputUnion interface {
	ImplementsPipelineRunParamsInputUnion()
}

type PipelineRunParamsInputArray []string

func (r PipelineRunParamsInputArray) ImplementsPipelineRunParamsInputUnion() {}

type PipelineRunJobParams struct {
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input param.Field[PipelineRunJobParamsInputUnion] `json:"input" api:"required"`
	// The ID of the pipeline to use for the document.
	PipelineID param.Field[string] `json:"pipeline_id" api:"required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async param.Field[AsyncConfigV3Param] `json:"async"`
	// Settings for pipeline execution that override pipeline defaults.
	Settings param.Field[PipelineSettingsParam] `json:"settings"`
}

func (r PipelineRunJobParams) MarshalJSON() (data []byte, err error) {
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
// Satisfied by [shared.UnionString], [PipelineRunJobParamsInputArray],
// [shared.UploadParam].
type PipelineRunJobParamsInputUnion interface {
	ImplementsPipelineRunJobParamsInputUnion()
}

type PipelineRunJobParamsInputArray []string

func (r PipelineRunJobParamsInputArray) ImplementsPipelineRunJobParamsInputUnion() {}
