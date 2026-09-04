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
func (r *ExtractService) Run(ctx context.Context, body ExtractRunParams, opts ...option.RequestOption) (res *ExtractRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "extract"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
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
	// Queue priority. 'batch' places the job in a lower-priority queue for non-urgent
	// bulk work. 'auto' (alias: 'standard') uses the default queue.
	QueuePriority param.Field[AsyncExtractConfigQueuePriority] `json:"queue_priority"`
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

// Queue priority. 'batch' places the job in a lower-priority queue for non-urgent
// bulk work. 'auto' (alias: 'standard') uses the default queue.
type AsyncExtractConfigQueuePriority string

const (
	AsyncExtractConfigQueuePriorityAuto     AsyncExtractConfigQueuePriority = "auto"
	AsyncExtractConfigQueuePriorityStandard AsyncExtractConfigQueuePriority = "standard"
	AsyncExtractConfigQueuePriorityBatch    AsyncExtractConfigQueuePriority = "batch"
)

func (r AsyncExtractConfigQueuePriority) IsKnown() bool {
	switch r {
	case AsyncExtractConfigQueuePriorityAuto, AsyncExtractConfigQueuePriorityStandard, AsyncExtractConfigQueuePriorityBatch:
		return true
	}
	return false
}

type ExtractSettingsParam struct {
	// Deprecated: prefer deep_extract, which supersedes array extraction for complex
	// and long (array-heavy) extractions via an agentic loop (at higher cost and
	// latency). If True, use array extraction.
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

type ExtractUsage struct {
	NumFields   int64                   `json:"num_fields" api:"required"`
	NumPages    int64                   `json:"num_pages" api:"required"`
	Credits     float64                 `json:"credits" api:"nullable"`
	ExtractMode ExtractUsageExtractMode `json:"extract_mode" api:"nullable"`
	// Raw extract quantities for accounts on the new pricing model.
	//
	// `extract_fields` is reported but not billed at launch. The add-on quantities
	// (`ocr_pages`, `charts`, `prompted_blocks`) come from the parse bundled into the
	// extract job; its page cost is covered by `extract_pages` but its add-ons are
	// billed separately. `tier` is "Batch" when the job ran on the batch queue, which
	// takes the batch discount on the rate card.
	UsageBreakdown ExtractUsageUsageBreakdown `json:"usage_breakdown" api:"nullable"`
	JSON           extractUsageJSON           `json:"-"`
}

// extractUsageJSON contains the JSON metadata for the struct [ExtractUsage]
type extractUsageJSON struct {
	NumFields      apijson.Field
	NumPages       apijson.Field
	Credits        apijson.Field
	ExtractMode    apijson.Field
	UsageBreakdown apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
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

// Raw extract quantities for accounts on the new pricing model.
//
// `extract_fields` is reported but not billed at launch. The add-on quantities
// (`ocr_pages`, `charts`, `prompted_blocks`) come from the parse bundled into the
// extract job; its page cost is covered by `extract_pages` but its add-ons are
// billed separately. `tier` is "Batch" when the job ran on the batch queue, which
// takes the batch discount on the rate card.
type ExtractUsageUsageBreakdown struct {
	ExtractModel   ExtractUsageUsageBreakdownExtractModel `json:"extract_model" api:"required"`
	Charts         int64                                  `json:"charts"`
	ExtractFields  int64                                  `json:"extract_fields"`
	ExtractPages   int64                                  `json:"extract_pages"`
	OcrPages       int64                                  `json:"ocr_pages"`
	PromptedBlocks int64                                  `json:"prompted_blocks"`
	Tier           ExtractUsageUsageBreakdownTier         `json:"tier"`
	JSON           extractUsageUsageBreakdownJSON         `json:"-"`
}

// extractUsageUsageBreakdownJSON contains the JSON metadata for the struct
// [ExtractUsageUsageBreakdown]
type extractUsageUsageBreakdownJSON struct {
	ExtractModel   apijson.Field
	Charts         apijson.Field
	ExtractFields  apijson.Field
	ExtractPages   apijson.Field
	OcrPages       apijson.Field
	PromptedBlocks apijson.Field
	Tier           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ExtractUsageUsageBreakdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extractUsageUsageBreakdownJSON) RawJSON() string {
	return r.raw
}

type ExtractUsageUsageBreakdownExtractModel string

const (
	ExtractUsageUsageBreakdownExtractModelExtract     ExtractUsageUsageBreakdownExtractModel = "Extract"
	ExtractUsageUsageBreakdownExtractModelDeepExtract ExtractUsageUsageBreakdownExtractModel = "Deep Extract"
)

func (r ExtractUsageUsageBreakdownExtractModel) IsKnown() bool {
	switch r {
	case ExtractUsageUsageBreakdownExtractModelExtract, ExtractUsageUsageBreakdownExtractModelDeepExtract:
		return true
	}
	return false
}

type ExtractUsageUsageBreakdownTier string

const (
	ExtractUsageUsageBreakdownTierDefault ExtractUsageUsageBreakdownTier = "Default"
	ExtractUsageUsageBreakdownTierBatch   ExtractUsageUsageBreakdownTier = "Batch"
)

func (r ExtractUsageUsageBreakdownTier) IsKnown() bool {
	switch r {
	case ExtractUsageUsageBreakdownTierDefault, ExtractUsageUsageBreakdownTierBatch:
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
	// Optional document-level deep extract confidence label.
	Confidence V3ExtractConfidence `json:"confidence" api:"nullable"`
	// Optional explanation for the document-level confidence label.
	ConfidenceReason string                `json:"confidence_reason" api:"nullable"`
	JobID            string                `json:"job_id" api:"nullable"`
	ResponseType     V3ExtractResponseType `json:"response_type"`
	// The link to the studio pipeline for the document.
	StudioLink string        `json:"studio_link" api:"nullable"`
	JSON       v3ExtractJSON `json:"-"`
}

// v3ExtractJSON contains the JSON metadata for the struct [V3Extract]
type v3ExtractJSON struct {
	Result           apijson.Field
	Usage            apijson.Field
	Confidence       apijson.Field
	ConfidenceReason apijson.Field
	JobID            apijson.Field
	ResponseType     apijson.Field
	StudioLink       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *V3Extract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3ExtractJSON) RawJSON() string {
	return r.raw
}

func (r V3Extract) ImplementsPipelineResponseResultExtractUnion() {}

func (r V3Extract) ImplementsPipelineResponseResultExtractArrayResult() {}

func (r V3Extract) ImplementsExtractRunResponse() {}

func (r V3Extract) ImplementsJobGetResponseAsyncJobResponseResult() {}

func (r V3Extract) ImplementsJobGetResponseEnhancedAsyncJobResponseResult() {}

// Optional document-level deep extract confidence label.
type V3ExtractConfidence string

const (
	V3ExtractConfidenceHigh V3ExtractConfidence = "high"
	V3ExtractConfidenceLow  V3ExtractConfidence = "low"
)

func (r V3ExtractConfidence) IsKnown() bool {
	switch r {
	case V3ExtractConfidenceHigh, V3ExtractConfidenceLow:
		return true
	}
	return false
}

type V3ExtractResponseType string

const (
	V3ExtractResponseTypeV3Extract V3ExtractResponseType = "v3_extract"
)

func (r V3ExtractResponseType) IsKnown() bool {
	switch r {
	case V3ExtractResponseTypeV3Extract:
		return true
	}
	return false
}

type ExtractRunResponse struct {
	// Optional document-level deep extract confidence label.
	Confidence ExtractRunResponseConfidence `json:"confidence" api:"nullable"`
	// Optional explanation for the document-level confidence label.
	ConfidenceReason string                         `json:"confidence_reason" api:"nullable"`
	JobID            string                         `json:"job_id" api:"nullable"`
	ResponseType     ExtractRunResponseResponseType `json:"response_type"`
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
	Confidence       apijson.Field
	ConfidenceReason apijson.Field
	JobID            apijson.Field
	ResponseType     apijson.Field
	Result           apijson.Field
	StudioLink       apijson.Field
	Usage            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
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
// Possible runtime types of the union are [V3Extract],
// [shared.AsyncExtractResponse].
func (r ExtractRunResponse) AsUnion() ExtractRunResponseUnion {
	return r.union
}

// Union satisfied by [V3Extract] or [shared.AsyncExtractResponse].
type ExtractRunResponseUnion interface {
	ImplementsExtractRunResponse()
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

// Optional document-level deep extract confidence label.
type ExtractRunResponseConfidence string

const (
	ExtractRunResponseConfidenceHigh ExtractRunResponseConfidence = "high"
	ExtractRunResponseConfidenceLow  ExtractRunResponseConfidence = "low"
)

func (r ExtractRunResponseConfidence) IsKnown() bool {
	switch r {
	case ExtractRunResponseConfidenceHigh, ExtractRunResponseConfidenceLow:
		return true
	}
	return false
}

type ExtractRunResponseResponseType string

const (
	ExtractRunResponseResponseTypeV3Extract ExtractRunResponseResponseType = "v3_extract"
)

func (r ExtractRunResponseResponseType) IsKnown() bool {
	switch r {
	case ExtractRunResponseResponseTypeV3Extract:
		return true
	}
	return false
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
	// Queue priority. 'batch' places the job in a lower-priority queue for non-urgent
	// bulk work. 'auto' (alias: 'standard') uses the default queue.
	QueuePriority param.Field[ExtractRunParamsBodyQueuePriority] `json:"queue_priority"`
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

// Queue priority. 'batch' places the job in a lower-priority queue for non-urgent
// bulk work. 'auto' (alias: 'standard') uses the default queue.
type ExtractRunParamsBodyQueuePriority string

const (
	ExtractRunParamsBodyQueuePriorityAuto     ExtractRunParamsBodyQueuePriority = "auto"
	ExtractRunParamsBodyQueuePriorityStandard ExtractRunParamsBodyQueuePriority = "standard"
	ExtractRunParamsBodyQueuePriorityBatch    ExtractRunParamsBodyQueuePriority = "batch"
)

func (r ExtractRunParamsBodyQueuePriority) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyQueuePriorityAuto, ExtractRunParamsBodyQueuePriorityStandard, ExtractRunParamsBodyQueuePriorityBatch:
		return true
	}
	return false
}

type ExtractRunJobParams struct {
	AsyncExtractConfig AsyncExtractConfigParam `json:"async_extract_config" api:"required"`
}

func (r ExtractRunJobParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AsyncExtractConfig)
}
