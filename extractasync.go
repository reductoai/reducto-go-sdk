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

// ExtractAsyncService contains methods and other services that help with
// interacting with the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtractAsyncService] method instead.
type ExtractAsyncService struct {
	Options []option.RequestOption
}

// NewExtractAsyncService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExtractAsyncService(opts ...option.RequestOption) (r *ExtractAsyncService) {
	r = &ExtractAsyncService{}
	r.Options = opts
	return
}

// Extract Async
func (r *ExtractAsyncService) New(ctx context.Context, body ExtractAsyncNewParams, opts ...option.RequestOption) (res *AsyncExtractResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "extract_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AsyncExtractConfigParam struct {
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input param.Field[AsyncExtractConfigInputUnionParam] `json:"input" api:"required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async param.Field[AsyncConfigV3Param] `json:"async"`
	// The instructions to use for the extraction.
	Instructions param.Field[InstructionsParam] `json:"instructions"`
	// The configuration options for parsing the document. If you are passing in a
	// jobid:// URL for the file, then this configuration will be ignored.
	Parsing param.Field[ParseOptionsParam] `json:"parsing"`
	// The settings to use for the extraction.
	Settings param.Field[ExtractSettingsParam] `json:"settings"`
}

func (r AsyncExtractConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AsyncExtractConfigParam) implementsExtractNewParamsBodyUnion() {}

// For parse/split/extract pipelines, the URL of the document to be processed. You
// can provide one of the following: 1. A publicly available URL 2. A presigned S3
// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
// directly uploading a document 4. A jobid:// prefixed URL obtained from a
// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
// API only)
//
//	For edit pipelines, this should be a string containing the edit instructions
//
// Satisfied by [shared.UnionString], [AsyncExtractConfigInputArrayParam],
// [UploadResponseParam].
type AsyncExtractConfigInputUnionParam interface {
	ImplementsAsyncExtractConfigInputUnionParam()
}

type AsyncExtractConfigInputArrayParam []string

func (r AsyncExtractConfigInputArrayParam) ImplementsAsyncExtractConfigInputUnionParam() {}

type AsyncExtractResponse struct {
	JobID string                   `json:"job_id" api:"required"`
	JSON  asyncExtractResponseJSON `json:"-"`
}

// asyncExtractResponseJSON contains the JSON metadata for the struct
// [AsyncExtractResponse]
type asyncExtractResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AsyncExtractResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asyncExtractResponseJSON) RawJSON() string {
	return r.raw
}

func (r AsyncExtractResponse) implementsExtractNewResponse() {}

type ExtractSettingsParam struct {
	// If True, use array extraction.
	ArrayExtract param.Field[bool] `json:"array_extract"`
	// The citations to use for the extraction.
	Citations param.Field[ExtractSettingsCitationsParam] `json:"citations"`
	// If True, include images in the extraction.
	IncludeImages param.Field[bool] `json:"include_images"`
	// If True, jobs will be processed with a higher throughput and priority at a
	// higher cost. Defaults to False.
	OptimizeForLatency param.Field[bool] `json:"optimize_for_latency"`
}

func (r ExtractSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The citations to use for the extraction.
type ExtractSettingsCitationsParam struct {
	// If True, include citations in the extraction.
	Enabled param.Field[bool] `json:"enabled"`
	// If True, enable numeric citation confidence scores. Defaults to True.
	NumericalConfidence param.Field[bool] `json:"numerical_confidence"`
}

func (r ExtractSettingsCitationsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstructionsParam struct {
	// The JSON schema to use for the extraction.
	Schema param.Field[interface{}] `json:"schema"`
	// The system prompt to use for the extraction.
	SystemPrompt param.Field[string] `json:"system_prompt"`
}

func (r InstructionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ParseOptionsParam struct {
	Enhance     param.Field[EnhanceParam]     `json:"enhance"`
	Formatting  param.Field[FormattingParam]  `json:"formatting"`
	Retrieval   param.Field[RetrievalParam]   `json:"retrieval"`
	Settings    param.Field[SettingsParam]    `json:"settings"`
	Spreadsheet param.Field[SpreadsheetParam] `json:"spreadsheet"`
}

func (r ParseOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractAsyncNewParams struct {
	AsyncExtractConfig AsyncExtractConfigParam `json:"async_extract_config" api:"required"`
}

func (r ExtractAsyncNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AsyncExtractConfig)
}
