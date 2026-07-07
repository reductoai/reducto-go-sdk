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
	"github.com/reductoai/reducto-go-sdk/shared"
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
func (r *ExtractService) Run(ctx context.Context, body ExtractRunParams, opts ...option.RequestOption) (res *ExtractRunResponseUnion, err error) {
	var env apijson.UnionUnmarshaler[ExtractRunResponseUnion]
	opts = slices.Concat(r.Options, opts)
	path := "extract"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Value
	return res, nil
}

// Extract Async
func (r *ExtractService) RunJob(ctx context.Context, body ExtractRunJobParams, opts ...option.RequestOption) (res *shared.AsyncExtractResponse, err error) {
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
// [shared.UploadParam].
type AsyncExtractConfigInputUnionParam interface {
	ImplementsAsyncExtractConfigInputUnionParam()
}

type AsyncExtractConfigInputArrayParam []string

func (r AsyncExtractConfigInputArrayParam) ImplementsAsyncExtractConfigInputUnionParam() {}

type ExtractSettingsParam struct {
	// If True, use array extraction.
	ArrayExtract param.Field[bool] `json:"array_extract"`
	// The citations to use for the extraction.
	Citations param.Field[ExtractSettingsCitationsParam] `json:"citations"`
	// If True, use Deep Extract, an agentic extraction mode that iteratively refines
	// its output to achieve near-perfect accuracy. Best for complex documents where
	// accuracy is critical.
	DeepExtract param.Field[bool] `json:"deep_extract"`
	// Force the endpoint result to be returned in URL form.
	ForceURLResult param.Field[bool] `json:"force_url_result"`
	// If True, include images in the extraction.
	IncludeImages param.Field[bool] `json:"include_images"`
	// If True, jobs will be processed with a higher throughput and priority at a
	// higher cost. Defaults to False.
	OptimizeForLatency param.Field[bool] `json:"optimize_for_latency"`
	// The page range to extract from (1-indexed). By default, the entire document is
	// used. For spreadsheets, you can also provide a list of sheet names.
	PageRange param.Field[ExtractSettingsPageRangeUnionParam] `json:"page_range"`
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
	// How much of the source parse block to embed on each citation's parentBlock.
	// 'full' (default) embeds the verbatim source-block HTML in parentBlock.content.
	// 'bbox_only' suppresses parentBlock.content (returned as an empty string) while
	// keeping parentBlock.bbox and all citation-level fields — this can drastically
	// shrink responses on table-heavy schemas where the same source block is cited
	// many times.
	ParentBlock param.Field[ExtractSettingsCitationsParentBlock] `json:"parent_block"`
}

func (r ExtractSettingsCitationsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How much of the source parse block to embed on each citation's parentBlock.
// 'full' (default) embeds the verbatim source-block HTML in parentBlock.content.
// 'bbox_only' suppresses parentBlock.content (returned as an empty string) while
// keeping parentBlock.bbox and all citation-level fields — this can drastically
// shrink responses on table-heavy schemas where the same source block is cited
// many times.
type ExtractSettingsCitationsParentBlock string

const (
	ExtractSettingsCitationsParentBlockFull     ExtractSettingsCitationsParentBlock = "full"
	ExtractSettingsCitationsParentBlockBboxOnly ExtractSettingsCitationsParentBlock = "bbox_only"
)

func (r ExtractSettingsCitationsParentBlock) IsKnown() bool {
	switch r {
	case ExtractSettingsCitationsParentBlockFull, ExtractSettingsCitationsParentBlockBboxOnly:
		return true
	}
	return false
}

// The page range to extract from (1-indexed). By default, the entire document is
// used. For spreadsheets, you can also provide a list of sheet names.
//
// Satisfied by [shared.PageRangeParam], [ExtractSettingsPageRangeArrayParam],
// [ExtractSettingsPageRangeArrayParam], [ExtractSettingsPageRangeArrayParam].
type ExtractSettingsPageRangeUnionParam interface {
	ImplementsExtractSettingsPageRangeUnionParam()
}

type ExtractSettingsPageRangeArrayParam []shared.PageRangeParam

func (r ExtractSettingsPageRangeArrayParam) ImplementsExtractSettingsPageRangeUnionParam() {}

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

type V3Extract map[string]interface{}

func (r V3Extract) ImplementsPipelineResponseResultExtractUnion() {}

func (r V3Extract) ImplementsPipelineResponseResultExtractArrayResultUnion() {}

func (r V3Extract) ImplementsExtractRunResponseUnion() {}

func (r V3Extract) ImplementsJobGetResponseAsyncJobResponseResultUnion() {}

func (r V3Extract) ImplementsJobGetResponseEnhancedAsyncJobResponseResultUnion() {}

// Union satisfied by [V3Extract] or [shared.AsyncExtractResponse].
type ExtractRunResponseUnion interface {
	ImplementsExtractRunResponseUnion()
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
			Type:       reflect.TypeOf(shared.AsyncExtractResponse{}),
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
// [ExtractRunParamsBodySyncExtractConfigInputArray], [shared.UploadParam].
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
