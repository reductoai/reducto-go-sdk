// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reducto

import (
	"context"
	"net/http"
	"reflect"
	"slices"

	"github.com/reductoai/reducto-go-sdk/internal/apijson"
	"github.com/reductoai/reducto-go-sdk/internal/param"
	"github.com/reductoai/reducto-go-sdk/internal/requestconfig"
	"github.com/reductoai/reducto-go-sdk/option"
	"github.com/tidwall/gjson"
)

// ExtractService contains methods and other services that help with interacting
// with the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtractService] method instead.
type ExtractService struct {
	Options []option.RequestOption
}

// NewExtractService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewExtractService(opts ...option.RequestOption) (r *ExtractService) {
	r = &ExtractService{}
	r.Options = opts
	return
}

// Extract
func (r *ExtractService) Run(ctx context.Context, body ExtractRunParams, opts ...option.RequestOption) (res *ExtractRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "extract"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Extract Async
func (r *ExtractService) RunJob(ctx context.Context, body ExtractRunJobParams, opts ...option.RequestOption) (res *AsyncExtractResponse, err error) {
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

func (r AsyncExtractConfigParam) implementsExtractRunParamsBodyUnion() {}

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

func (r AsyncExtractResponse) implementsExtractRunResponse() {}

type ExtractSettingsParam struct {
	// If True, use array extraction.
	ArrayExtract param.Field[bool] `json:"array_extract"`
	// The citations to use for the extraction.
	Citations param.Field[ExtractSettingsCitationsParam] `json:"citations"`
	// If True, use Deep Extract, an agentic extraction mode that iteratively refines
	// its output to achieve near-perfect accuracy. Best for complex documents where
	// accuracy is critical.
	DeepExtract param.Field[bool] `json:"deep_extract"`
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

type ExtractUsage struct {
	NumFields   int64                   `json:"num_fields" api:"required"`
	NumPages    int64                   `json:"num_pages" api:"required"`
	Credits     float64                 `json:"credits" api:"nullable"`
	ExtractMode ExtractUsageExtractMode `json:"extract_mode" api:"nullable"`
	JSON        extractUsageJSON        `json:"-"`
}

// extractUsageJSON contains the JSON metadata for the struct [ExtractUsage]
type extractUsageJSON struct {
	NumFields   apijson.Field
	NumPages    apijson.Field
	Credits     apijson.Field
	ExtractMode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtractUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extractUsageJSON) RawJSON() string {
	return r.raw
}

type ExtractUsageExtractMode string

const (
	ExtractUsageExtractModeSuperAgent       ExtractUsageExtractMode = "super_agent"
	ExtractUsageExtractModeExtract          ExtractUsageExtractMode = "extract"
	ExtractUsageExtractModeSpreadsheetAgent ExtractUsageExtractMode = "spreadsheet_agent"
)

func (r ExtractUsageExtractMode) IsKnown() bool {
	switch r {
	case ExtractUsageExtractModeSuperAgent, ExtractUsageExtractModeExtract, ExtractUsageExtractModeSpreadsheetAgent:
		return true
	}
	return false
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

type V3Extract struct {
	// The extracted response in your provided schema. This is a list of dictionaries.
	// If disable_chunking is True (default), then it will be a list of length one.
	Result []interface{} `json:"result" api:"required"`
	Usage  ExtractUsage  `json:"usage" api:"required"`
	JobID  string        `json:"job_id" api:"nullable"`
	// The link to the studio pipeline for the document.
	StudioLink string        `json:"studio_link" api:"nullable"`
	JSON       v3ExtractJSON `json:"-"`
}

// v3ExtractJSON contains the JSON metadata for the struct [V3Extract]
type v3ExtractJSON struct {
	Result      apijson.Field
	Usage       apijson.Field
	JobID       apijson.Field
	StudioLink  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3Extract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3ExtractJSON) RawJSON() string {
	return r.raw
}

func (r V3Extract) implementsExtractRunResponse() {}

func (r V3Extract) implementsPipelineResponseResultExtractUnion() {}

func (r V3Extract) implementsPipelineResponseResultExtractExtractVariant0Result() {}

func (r V3Extract) implementsJobGetResponseAsyncJobResponseResult() {}

func (r V3Extract) implementsJobGetResponseEnhancedAsyncJobResponseResult() {}

type ExtractRunResponse struct {
	JobID string `json:"job_id" api:"nullable"`
	// This field can have the runtime type of [[]interface{}].
	Result interface{} `json:"result"`
	// The link to the studio pipeline for the document.
	StudioLink string                 `json:"studio_link" api:"nullable"`
	Usage      ExtractUsage           `json:"usage"`
	JSON       extractRunResponseJSON `json:"-"`
	union      ExtractRunResponseUnion
}

// extractRunResponseJSON contains the JSON metadata for the struct
// [ExtractRunResponse]
type extractRunResponseJSON struct {
	JobID       apijson.Field
	Result      apijson.Field
	StudioLink  apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r extractRunResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ExtractRunResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ExtractRunResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ExtractRunResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [V3Extract], [AsyncExtractResponse].
func (r ExtractRunResponse) AsUnion() ExtractRunResponseUnion {
	return r.union
}

// Union satisfied by [V3Extract] or [AsyncExtractResponse].
type ExtractRunResponseUnion interface {
	implementsExtractRunResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ExtractRunResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V3Extract{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AsyncExtractResponse{}),
		},
	)
}

type ExtractRunParams struct {
	Body ExtractRunParamsBodyUnion `json:"body" api:"required"`
}

func (r ExtractRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ExtractRunParamsBody struct {
	Input param.Field[interface{}] `json:"input" api:"required"`
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

func (r ExtractRunParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBody) implementsExtractRunParamsBodyUnion() {}

// Satisfied by [ExtractRunParamsBodySyncExtractConfig], [AsyncExtractConfigParam],
// [ExtractRunParamsBody].
type ExtractRunParamsBodyUnion interface {
	implementsExtractRunParamsBodyUnion()
}

type ExtractRunParamsBodySyncExtractConfig struct {
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input param.Field[ExtractRunParamsBodySyncExtractConfigInputUnion] `json:"input" api:"required"`
	// The instructions to use for the extraction.
	Instructions param.Field[InstructionsParam] `json:"instructions"`
	// The configuration options for parsing the document. If you are passing in a
	// jobid:// URL for the file, then this configuration will be ignored.
	Parsing param.Field[ParseOptionsParam] `json:"parsing"`
	// The settings to use for the extraction.
	Settings param.Field[ExtractSettingsParam] `json:"settings"`
}

func (r ExtractRunParamsBodySyncExtractConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBodySyncExtractConfig) implementsExtractRunParamsBodyUnion() {}

// For parse/split/extract pipelines, the URL of the document to be processed. You
// can provide one of the following: 1. A publicly available URL 2. A presigned S3
// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
// directly uploading a document 4. A jobid:// prefixed URL obtained from a
// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
// API only)
//
//	For edit pipelines, this should be a string containing the edit instructions
//
// Satisfied by [shared.UnionString],
// [ExtractRunParamsBodySyncExtractConfigInputArray], [UploadResponseParam].
type ExtractRunParamsBodySyncExtractConfigInputUnion interface {
	ImplementsExtractRunParamsBodySyncExtractConfigInputUnion()
}

type ExtractRunParamsBodySyncExtractConfigInputArray []string

func (r ExtractRunParamsBodySyncExtractConfigInputArray) ImplementsExtractRunParamsBodySyncExtractConfigInputUnion() {
}

type ExtractRunJobParams struct {
	AsyncExtractConfig AsyncExtractConfigParam `json:"async_extract_config" api:"required"`
}

func (r ExtractRunJobParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AsyncExtractConfig)
}
