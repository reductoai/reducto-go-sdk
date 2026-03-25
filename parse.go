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

// ParseService contains methods and other services that help with interacting with
// the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewParseService] method instead.
type ParseService struct {
	Options []option.RequestOption
}

// NewParseService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewParseService(opts ...option.RequestOption) (r *ParseService) {
	r = &ParseService{}
	r.Options = opts
	return
}

// Parse
func (r *ParseService) Run(ctx context.Context, body ParseRunParams, opts ...option.RequestOption) (res *ParseRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "parse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Async Parse
func (r *ParseService) RunJob(ctx context.Context, body ParseRunJobParams, opts ...option.RequestOption) (res *AsyncParseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "parse_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AsyncConfigV3Param struct {
	// JSON metadata included in webhook request body. Defaults to None.
	Metadata param.Field[interface{}] `json:"metadata"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool] `json:"priority"`
	// The webhook configuration for the asynchronous processing.
	Webhook param.Field[AsyncConfigV3WebhookUnionParam] `json:"webhook"`
}

func (r AsyncConfigV3Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The webhook configuration for the asynchronous processing.
type AsyncConfigV3WebhookParam struct {
	Channels param.Field[interface{}]              `json:"channels"`
	Mode     param.Field[AsyncConfigV3WebhookMode] `json:"mode"`
	URL      param.Field[string]                   `json:"url"`
}

func (r AsyncConfigV3WebhookParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AsyncConfigV3WebhookParam) implementsAsyncConfigV3WebhookUnionParam() {}

// The webhook configuration for the asynchronous processing.
//
// Satisfied by [AsyncConfigV3WebhookSvixWebhookConfigParam],
// [AsyncConfigV3WebhookDirectWebhookConfigParam], [AsyncConfigV3WebhookParam].
type AsyncConfigV3WebhookUnionParam interface {
	implementsAsyncConfigV3WebhookUnionParam()
}

type AsyncConfigV3WebhookSvixWebhookConfigParam struct {
	// A list of Svix channels the message will be delivered down, omit to send to all
	// channels.
	Channels param.Field[[]string]                                  `json:"channels"`
	Mode     param.Field[AsyncConfigV3WebhookSvixWebhookConfigMode] `json:"mode"`
}

func (r AsyncConfigV3WebhookSvixWebhookConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AsyncConfigV3WebhookSvixWebhookConfigParam) implementsAsyncConfigV3WebhookUnionParam() {}

type AsyncConfigV3WebhookSvixWebhookConfigMode string

const (
	AsyncConfigV3WebhookSvixWebhookConfigModeSvix AsyncConfigV3WebhookSvixWebhookConfigMode = "svix"
)

func (r AsyncConfigV3WebhookSvixWebhookConfigMode) IsKnown() bool {
	switch r {
	case AsyncConfigV3WebhookSvixWebhookConfigModeSvix:
		return true
	}
	return false
}

type AsyncConfigV3WebhookDirectWebhookConfigParam struct {
	URL  param.Field[string]                                      `json:"url" api:"required"`
	Mode param.Field[AsyncConfigV3WebhookDirectWebhookConfigMode] `json:"mode"`
}

func (r AsyncConfigV3WebhookDirectWebhookConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AsyncConfigV3WebhookDirectWebhookConfigParam) implementsAsyncConfigV3WebhookUnionParam() {}

type AsyncConfigV3WebhookDirectWebhookConfigMode string

const (
	AsyncConfigV3WebhookDirectWebhookConfigModeDirect AsyncConfigV3WebhookDirectWebhookConfigMode = "direct"
)

func (r AsyncConfigV3WebhookDirectWebhookConfigMode) IsKnown() bool {
	switch r {
	case AsyncConfigV3WebhookDirectWebhookConfigModeDirect:
		return true
	}
	return false
}

type AsyncConfigV3WebhookMode string

const (
	AsyncConfigV3WebhookModeSvix   AsyncConfigV3WebhookMode = "svix"
	AsyncConfigV3WebhookModeDirect AsyncConfigV3WebhookMode = "direct"
)

func (r AsyncConfigV3WebhookMode) IsKnown() bool {
	switch r {
	case AsyncConfigV3WebhookModeSvix, AsyncConfigV3WebhookModeDirect:
		return true
	}
	return false
}

type AsyncParseConfigParam struct {
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input param.Field[AsyncParseConfigInputUnionParam] `json:"input" api:"required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async      param.Field[AsyncConfigV3Param] `json:"async"`
	Enhance    param.Field[EnhanceParam]       `json:"enhance"`
	Formatting param.Field[FormattingParam]    `json:"formatting"`
	// Queue priority. 'batch' for non-urgent work that processes when spare GPU
	// capacity is available.
	QueuePriority param.Field[AsyncParseConfigQueuePriority] `json:"queue_priority"`
	Retrieval     param.Field[RetrievalParam]                `json:"retrieval"`
	Settings      param.Field[SettingsParam]                 `json:"settings"`
	Spreadsheet   param.Field[SpreadsheetParam]              `json:"spreadsheet"`
}

func (r AsyncParseConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AsyncParseConfigParam) implementsParseRunParamsBodyUnion() {}

// For parse/split/extract pipelines, the URL of the document to be processed. You
// can provide one of the following: 1. A publicly available URL 2. A presigned S3
// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
// directly uploading a document 4. A jobid:// prefixed URL obtained from a
// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
// API only)
//
//	For edit pipelines, this should be a string containing the edit instructions
//
// Satisfied by [shared.UnionString], [AsyncParseConfigInputArrayParam],
// [shared.UploadParam].
type AsyncParseConfigInputUnionParam interface {
	ImplementsAsyncParseConfigInputUnionParam()
}

type AsyncParseConfigInputArrayParam []string

func (r AsyncParseConfigInputArrayParam) ImplementsAsyncParseConfigInputUnionParam() {}

// Queue priority. 'batch' for non-urgent work that processes when spare GPU
// capacity is available.
type AsyncParseConfigQueuePriority string

const (
	AsyncParseConfigQueuePriorityAuto  AsyncParseConfigQueuePriority = "auto"
	AsyncParseConfigQueuePriorityBatch AsyncParseConfigQueuePriority = "batch"
)

func (r AsyncParseConfigQueuePriority) IsKnown() bool {
	switch r {
	case AsyncParseConfigQueuePriorityAuto, AsyncParseConfigQueuePriorityBatch:
		return true
	}
	return false
}

type AsyncParseResponse struct {
	JobID string                 `json:"job_id" api:"required"`
	JSON  asyncParseResponseJSON `json:"-"`
}

// asyncParseResponseJSON contains the JSON metadata for the struct
// [AsyncParseResponse]
type asyncParseResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AsyncParseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asyncParseResponseJSON) RawJSON() string {
	return r.raw
}

func (r AsyncParseResponse) implementsParseRunResponse() {}

type EnhanceParam struct {
	// Agentic uses vision language models to enhance the accuracy of the output of
	// different types of extraction. This will incur a cost and latency increase.
	Agentic param.Field[[]EnhanceAgenticUnionParam] `json:"agentic"`
	// If True, use an advanced vision language model to improve reading order
	// accuracy, with a small increase in latency. Defaults to False.
	IntelligentOrdering param.Field[bool] `json:"intelligent_ordering"`
	// If True, summarize figures using a small vision language model. Defaults to
	// True.
	SummarizeFigures param.Field[bool] `json:"summarize_figures"`
}

func (r EnhanceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnhanceAgenticParam struct {
	Scope param.Field[EnhanceAgenticScope] `json:"scope" api:"required"`
	// If True, use the advanced chart agent. Defaults to False.
	AdvancedChartAgent param.Field[bool] `json:"advanced_chart_agent"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
	// If True, return overlays for the figure. This is so you can use the overlays to
	// double check the quality of the extraction
	ReturnOverlays param.Field[bool] `json:"return_overlays"`
}

func (r EnhanceAgenticParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnhanceAgenticParam) implementsEnhanceAgenticUnionParam() {}

// Satisfied by [EnhanceAgenticTableAgenticParam],
// [EnhanceAgenticFigureAgenticParam], [EnhanceAgenticTextAgenticParam],
// [EnhanceAgenticParam].
type EnhanceAgenticUnionParam interface {
	implementsEnhanceAgenticUnionParam()
}

type EnhanceAgenticTableAgenticParam struct {
	Scope param.Field[EnhanceAgenticTableAgenticScope] `json:"scope" api:"required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r EnhanceAgenticTableAgenticParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnhanceAgenticTableAgenticParam) implementsEnhanceAgenticUnionParam() {}

type EnhanceAgenticTableAgenticScope string

const (
	EnhanceAgenticTableAgenticScopeTable EnhanceAgenticTableAgenticScope = "table"
)

func (r EnhanceAgenticTableAgenticScope) IsKnown() bool {
	switch r {
	case EnhanceAgenticTableAgenticScopeTable:
		return true
	}
	return false
}

type EnhanceAgenticFigureAgenticParam struct {
	Scope param.Field[EnhanceAgenticFigureAgenticScope] `json:"scope" api:"required"`
	// If True, use the advanced chart agent. Defaults to False.
	AdvancedChartAgent param.Field[bool] `json:"advanced_chart_agent"`
	// Custom prompt for figure agentic.
	Prompt param.Field[string] `json:"prompt"`
	// If True, return overlays for the figure. This is so you can use the overlays to
	// double check the quality of the extraction
	ReturnOverlays param.Field[bool] `json:"return_overlays"`
}

func (r EnhanceAgenticFigureAgenticParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnhanceAgenticFigureAgenticParam) implementsEnhanceAgenticUnionParam() {}

type EnhanceAgenticFigureAgenticScope string

const (
	EnhanceAgenticFigureAgenticScopeFigure EnhanceAgenticFigureAgenticScope = "figure"
)

func (r EnhanceAgenticFigureAgenticScope) IsKnown() bool {
	switch r {
	case EnhanceAgenticFigureAgenticScopeFigure:
		return true
	}
	return false
}

type EnhanceAgenticTextAgenticParam struct {
	Scope param.Field[EnhanceAgenticTextAgenticScope] `json:"scope" api:"required"`
	// Custom instructions for agentic text. Note: This only applies to form regions
	// (key-value).
	Prompt param.Field[string] `json:"prompt"`
}

func (r EnhanceAgenticTextAgenticParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EnhanceAgenticTextAgenticParam) implementsEnhanceAgenticUnionParam() {}

type EnhanceAgenticTextAgenticScope string

const (
	EnhanceAgenticTextAgenticScopeText EnhanceAgenticTextAgenticScope = "text"
)

func (r EnhanceAgenticTextAgenticScope) IsKnown() bool {
	switch r {
	case EnhanceAgenticTextAgenticScopeText:
		return true
	}
	return false
}

type EnhanceAgenticScope string

const (
	EnhanceAgenticScopeTable  EnhanceAgenticScope = "table"
	EnhanceAgenticScopeFigure EnhanceAgenticScope = "figure"
	EnhanceAgenticScopeText   EnhanceAgenticScope = "text"
)

func (r EnhanceAgenticScope) IsKnown() bool {
	switch r {
	case EnhanceAgenticScopeTable, EnhanceAgenticScopeFigure, EnhanceAgenticScopeText:
		return true
	}
	return false
}

type FormattingParam struct {
	// If True, add page markers to the output. Defaults to False. Useful for
	// extracting data with page specific information.
	AddPageMarkers param.Field[bool] `json:"add_page_markers"`
	// A list of formatting to include in the output.
	Include param.Field[[]FormattingInclude] `json:"include"`
	// A flag to indicate if consecutive tables with the same number of columns should
	// be merged. Defaults to False.
	MergeTables param.Field[bool] `json:"merge_tables"`
	// The mode to use for table output. Defaults to dynamic, which returns md for
	// simpler tables and html for more complex tables.
	TableOutputFormat param.Field[FormattingTableOutputFormat] `json:"table_output_format"`
}

func (r FormattingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FormattingInclude string

const (
	FormattingIncludeChangeTracking   FormattingInclude = "change_tracking"
	FormattingIncludeHighlight        FormattingInclude = "highlight"
	FormattingIncludeComments         FormattingInclude = "comments"
	FormattingIncludeHyperlinks       FormattingInclude = "hyperlinks"
	FormattingIncludeSignatures       FormattingInclude = "signatures"
	FormattingIncludeIgnoreWatermarks FormattingInclude = "ignore_watermarks"
)

func (r FormattingInclude) IsKnown() bool {
	switch r {
	case FormattingIncludeChangeTracking, FormattingIncludeHighlight, FormattingIncludeComments, FormattingIncludeHyperlinks, FormattingIncludeSignatures, FormattingIncludeIgnoreWatermarks:
		return true
	}
	return false
}

// The mode to use for table output. Defaults to dynamic, which returns md for
// simpler tables and html for more complex tables.
type FormattingTableOutputFormat string

const (
	FormattingTableOutputFormatHTML     FormattingTableOutputFormat = "html"
	FormattingTableOutputFormatJson     FormattingTableOutputFormat = "json"
	FormattingTableOutputFormatMd       FormattingTableOutputFormat = "md"
	FormattingTableOutputFormatJsonbbox FormattingTableOutputFormat = "jsonbbox"
	FormattingTableOutputFormatDynamic  FormattingTableOutputFormat = "dynamic"
	FormattingTableOutputFormatCsv      FormattingTableOutputFormat = "csv"
)

func (r FormattingTableOutputFormat) IsKnown() bool {
	switch r {
	case FormattingTableOutputFormatHTML, FormattingTableOutputFormatJson, FormattingTableOutputFormatMd, FormattingTableOutputFormatJsonbbox, FormattingTableOutputFormatDynamic, FormattingTableOutputFormatCsv:
		return true
	}
	return false
}

type ParseResponse struct {
	// The duration of the parse request in seconds.
	Duration float64 `json:"duration" api:"required"`
	JobID    string  `json:"job_id" api:"required"`
	// The response from the document processing service. Note that there can be two
	// types of responses, Full Result and URL Result. This is due to limitations on
	// the max return size on HTTPS. If the response is too large, it will be returned
	// as a presigned URL in the URL response. You should handle this in your
	// application.
	Result ParseResponseResult `json:"result" api:"required"`
	Usage  ParseUsage          `json:"usage" api:"required"`
	// The storage URL of the converted PDF file.
	PdfURL string `json:"pdf_url" api:"nullable"`
	// The link to the studio pipeline for the document.
	StudioLink string            `json:"studio_link" api:"nullable"`
	JSON       parseResponseJSON `json:"-"`
}

// parseResponseJSON contains the JSON metadata for the struct [ParseResponse]
type parseResponseJSON struct {
	Duration    apijson.Field
	JobID       apijson.Field
	Result      apijson.Field
	Usage       apijson.Field
	PdfURL      apijson.Field
	StudioLink  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseJSON) RawJSON() string {
	return r.raw
}

func (r ParseResponse) implementsParseRunResponse() {}

func (r ParseResponse) implementsPipelineResponseResultParseUnion() {}

func (r ParseResponse) implementsJobGetResponseAsyncJobResponseResult() {}

func (r ParseResponse) implementsJobGetResponseEnhancedAsyncJobResponseResult() {}

// The response from the document processing service. Note that there can be two
// types of responses, Full Result and URL Result. This is due to limitations on
// the max return size on HTTPS. If the response is too large, it will be returned
// as a presigned URL in the URL response. You should handle this in your
// application.
type ParseResponseResult struct {
	// type = 'full'
	Type ParseResponseResultType `json:"type" api:"required"`
	// This field can have the runtime type of [[]ParseResponseResultFullResultChunk].
	Chunks interface{} `json:"chunks"`
	// This field can have the runtime type of [interface{}].
	Custom interface{} `json:"custom"`
	// This field can have the runtime type of [ParseResponseResultFullResultOcr].
	Ocr      interface{}             `json:"ocr"`
	ResultID string                  `json:"result_id"`
	URL      string                  `json:"url"`
	JSON     parseResponseResultJSON `json:"-"`
	union    ParseResponseResultUnion
}

// parseResponseResultJSON contains the JSON metadata for the struct
// [ParseResponseResult]
type parseResponseResultJSON struct {
	Type        apijson.Field
	Chunks      apijson.Field
	Custom      apijson.Field
	Ocr         apijson.Field
	ResultID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r parseResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *ParseResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = ParseResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ParseResponseResultUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ParseResponseResultFullResult],
// [ParseResponseResultURLResult].
func (r ParseResponseResult) AsUnion() ParseResponseResultUnion {
	return r.union
}

// The response from the document processing service. Note that there can be two
// types of responses, Full Result and URL Result. This is due to limitations on
// the max return size on HTTPS. If the response is too large, it will be returned
// as a presigned URL in the URL response. You should handle this in your
// application.
//
// Union satisfied by [ParseResponseResultFullResult] or
// [ParseResponseResultURLResult].
type ParseResponseResultUnion interface {
	implementsParseResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ParseResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParseResponseResultFullResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParseResponseResultURLResult{}),
		},
	)
}

type ParseResponseResultFullResult struct {
	Chunks []ParseResponseResultFullResultChunk `json:"chunks" api:"required"`
	// type = 'full'
	Type   ParseResponseResultFullResultType `json:"type" api:"required"`
	Custom interface{}                       `json:"custom"`
	Ocr    ParseResponseResultFullResultOcr  `json:"ocr" api:"nullable"`
	JSON   parseResponseResultFullResultJSON `json:"-"`
}

// parseResponseResultFullResultJSON contains the JSON metadata for the struct
// [ParseResponseResultFullResult]
type parseResponseResultFullResultJSON struct {
	Chunks      apijson.Field
	Type        apijson.Field
	Custom      apijson.Field
	Ocr         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParseResponseResultFullResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseResultFullResultJSON) RawJSON() string {
	return r.raw
}

func (r ParseResponseResultFullResult) implementsParseResponseResult() {}

type ParseResponseResultFullResultChunk struct {
	Blocks []ParseResponseResultFullResultChunksBlock `json:"blocks" api:"required"`
	// The content of the chunk extracted from the document.
	Content string `json:"content" api:"required"`
	// Chunk content optimized for embedding and retrieval.
	Embed string `json:"embed" api:"required"`
	// The enriched content of the chunk extracted from the document.
	Enriched string `json:"enriched" api:"required,nullable"`
	// Whether the enrichment was successful.
	EnrichmentSuccess bool                                   `json:"enrichment_success"`
	JSON              parseResponseResultFullResultChunkJSON `json:"-"`
}

// parseResponseResultFullResultChunkJSON contains the JSON metadata for the struct
// [ParseResponseResultFullResultChunk]
type parseResponseResultFullResultChunkJSON struct {
	Blocks            apijson.Field
	Content           apijson.Field
	Embed             apijson.Field
	Enriched          apijson.Field
	EnrichmentSuccess apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ParseResponseResultFullResultChunk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseResultFullResultChunkJSON) RawJSON() string {
	return r.raw
}

type ParseResponseResultFullResultChunksBlock struct {
	// The bounding box of the block extracted from the document.
	Bbox BoundingBox `json:"bbox" api:"required"`
	// The content of the block extracted from the document.
	Content string `json:"content" api:"required"`
	// The type of block extracted from the document.
	Type ParseResponseResultFullResultChunksBlocksType `json:"type" api:"required"`
	// (Experimental) The URL/link to chart data JSON for figure blocks processed by
	// chart agent.
	ChartData []string `json:"chart_data" api:"nullable"`
	// The confidence for the block. It is either low or high and takes into account
	// factors like OCR and table structure
	Confidence string `json:"confidence" api:"nullable"`
	// Extra metadata fields for the block. Fields like 'is_chart' will only appear
	// when set to True.
	Extra map[string]interface{} `json:"extra" api:"nullable"`
	// Granular confidence scores for the block. It is a dictionary of confidence
	// scores for the block. The confidence scores will not be None if the user has
	// enabled numeric confidence scores.
	GranularConfidence ParseResponseResultFullResultChunksBlocksGranularConfidence `json:"granular_confidence" api:"nullable"`
	// (Experimental) The URL of the image associated with the block.
	ImageURL string                                       `json:"image_url" api:"nullable"`
	JSON     parseResponseResultFullResultChunksBlockJSON `json:"-"`
}

// parseResponseResultFullResultChunksBlockJSON contains the JSON metadata for the
// struct [ParseResponseResultFullResultChunksBlock]
type parseResponseResultFullResultChunksBlockJSON struct {
	Bbox               apijson.Field
	Content            apijson.Field
	Type               apijson.Field
	ChartData          apijson.Field
	Confidence         apijson.Field
	Extra              apijson.Field
	GranularConfidence apijson.Field
	ImageURL           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ParseResponseResultFullResultChunksBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseResultFullResultChunksBlockJSON) RawJSON() string {
	return r.raw
}

// The type of block extracted from the document.
type ParseResponseResultFullResultChunksBlocksType string

const (
	ParseResponseResultFullResultChunksBlocksTypeHeader        ParseResponseResultFullResultChunksBlocksType = "Header"
	ParseResponseResultFullResultChunksBlocksTypeFooter        ParseResponseResultFullResultChunksBlocksType = "Footer"
	ParseResponseResultFullResultChunksBlocksTypeTitle         ParseResponseResultFullResultChunksBlocksType = "Title"
	ParseResponseResultFullResultChunksBlocksTypeSectionHeader ParseResponseResultFullResultChunksBlocksType = "Section Header"
	ParseResponseResultFullResultChunksBlocksTypePageNumber    ParseResponseResultFullResultChunksBlocksType = "Page Number"
	ParseResponseResultFullResultChunksBlocksTypeListItem      ParseResponseResultFullResultChunksBlocksType = "List Item"
	ParseResponseResultFullResultChunksBlocksTypeFigure        ParseResponseResultFullResultChunksBlocksType = "Figure"
	ParseResponseResultFullResultChunksBlocksTypeTable         ParseResponseResultFullResultChunksBlocksType = "Table"
	ParseResponseResultFullResultChunksBlocksTypeKeyValue      ParseResponseResultFullResultChunksBlocksType = "Key Value"
	ParseResponseResultFullResultChunksBlocksTypeText          ParseResponseResultFullResultChunksBlocksType = "Text"
	ParseResponseResultFullResultChunksBlocksTypeComment       ParseResponseResultFullResultChunksBlocksType = "Comment"
	ParseResponseResultFullResultChunksBlocksTypeSignature     ParseResponseResultFullResultChunksBlocksType = "Signature"
)

func (r ParseResponseResultFullResultChunksBlocksType) IsKnown() bool {
	switch r {
	case ParseResponseResultFullResultChunksBlocksTypeHeader, ParseResponseResultFullResultChunksBlocksTypeFooter, ParseResponseResultFullResultChunksBlocksTypeTitle, ParseResponseResultFullResultChunksBlocksTypeSectionHeader, ParseResponseResultFullResultChunksBlocksTypePageNumber, ParseResponseResultFullResultChunksBlocksTypeListItem, ParseResponseResultFullResultChunksBlocksTypeFigure, ParseResponseResultFullResultChunksBlocksTypeTable, ParseResponseResultFullResultChunksBlocksTypeKeyValue, ParseResponseResultFullResultChunksBlocksTypeText, ParseResponseResultFullResultChunksBlocksTypeComment, ParseResponseResultFullResultChunksBlocksTypeSignature:
		return true
	}
	return false
}

// Granular confidence scores for the block. It is a dictionary of confidence
// scores for the block. The confidence scores will not be None if the user has
// enabled numeric confidence scores.
type ParseResponseResultFullResultChunksBlocksGranularConfidence struct {
	ExtractConfidence float64                                                         `json:"extract_confidence" api:"nullable"`
	ParseConfidence   float64                                                         `json:"parse_confidence" api:"nullable"`
	JSON              parseResponseResultFullResultChunksBlocksGranularConfidenceJSON `json:"-"`
}

// parseResponseResultFullResultChunksBlocksGranularConfidenceJSON contains the
// JSON metadata for the struct
// [ParseResponseResultFullResultChunksBlocksGranularConfidence]
type parseResponseResultFullResultChunksBlocksGranularConfidenceJSON struct {
	ExtractConfidence apijson.Field
	ParseConfidence   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ParseResponseResultFullResultChunksBlocksGranularConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseResultFullResultChunksBlocksGranularConfidenceJSON) RawJSON() string {
	return r.raw
}

// type = 'full'
type ParseResponseResultFullResultType string

const (
	ParseResponseResultFullResultTypeFull ParseResponseResultFullResultType = "full"
)

func (r ParseResponseResultFullResultType) IsKnown() bool {
	switch r {
	case ParseResponseResultFullResultTypeFull:
		return true
	}
	return false
}

type ParseResponseResultFullResultOcr struct {
	Lines []ParseResponseResultFullResultOcrLine `json:"lines" api:"required"`
	Words []ParseResponseResultFullResultOcrWord `json:"words" api:"required"`
	JSON  parseResponseResultFullResultOcrJSON   `json:"-"`
}

// parseResponseResultFullResultOcrJSON contains the JSON metadata for the struct
// [ParseResponseResultFullResultOcr]
type parseResponseResultFullResultOcrJSON struct {
	Lines       apijson.Field
	Words       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParseResponseResultFullResultOcr) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseResultFullResultOcrJSON) RawJSON() string {
	return r.raw
}

type ParseResponseResultFullResultOcrLine struct {
	Bbox BoundingBox `json:"bbox" api:"required"`
	Text string      `json:"text" api:"required"`
	// The index of the chunk that the line belongs to.
	ChunkIndex int64 `json:"chunk_index" api:"nullable"`
	// OCR confidence score between 0 and 1, where 1 indicates highest confidence
	Confidence float64 `json:"confidence" api:"nullable"`
	// The rotation angle in degrees, from 0 to 360, counterclockwise.
	Rotation int64                                    `json:"rotation" api:"nullable"`
	JSON     parseResponseResultFullResultOcrLineJSON `json:"-"`
}

// parseResponseResultFullResultOcrLineJSON contains the JSON metadata for the
// struct [ParseResponseResultFullResultOcrLine]
type parseResponseResultFullResultOcrLineJSON struct {
	Bbox        apijson.Field
	Text        apijson.Field
	ChunkIndex  apijson.Field
	Confidence  apijson.Field
	Rotation    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParseResponseResultFullResultOcrLine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseResultFullResultOcrLineJSON) RawJSON() string {
	return r.raw
}

type ParseResponseResultFullResultOcrWord struct {
	Bbox BoundingBox `json:"bbox" api:"required"`
	Text string      `json:"text" api:"required"`
	// The index of the chunk that the word belongs to.
	ChunkIndex int64 `json:"chunk_index" api:"nullable"`
	// OCR confidence score between 0 and 1, where 1 indicates highest confidence
	Confidence float64 `json:"confidence" api:"nullable"`
	// The rotation angle in degrees, from 0 to 360, counterclockwise.
	Rotation int64                                    `json:"rotation" api:"nullable"`
	JSON     parseResponseResultFullResultOcrWordJSON `json:"-"`
}

// parseResponseResultFullResultOcrWordJSON contains the JSON metadata for the
// struct [ParseResponseResultFullResultOcrWord]
type parseResponseResultFullResultOcrWordJSON struct {
	Bbox        apijson.Field
	Text        apijson.Field
	ChunkIndex  apijson.Field
	Confidence  apijson.Field
	Rotation    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParseResponseResultFullResultOcrWord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseResultFullResultOcrWordJSON) RawJSON() string {
	return r.raw
}

type ParseResponseResultURLResult struct {
	ResultID string `json:"result_id" api:"required"`
	// type = 'url'
	Type ParseResponseResultURLResultType `json:"type" api:"required"`
	URL  string                           `json:"url" api:"required"`
	JSON parseResponseResultURLResultJSON `json:"-"`
}

// parseResponseResultURLResultJSON contains the JSON metadata for the struct
// [ParseResponseResultURLResult]
type parseResponseResultURLResultJSON struct {
	ResultID    apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParseResponseResultURLResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseResultURLResultJSON) RawJSON() string {
	return r.raw
}

func (r ParseResponseResultURLResult) implementsParseResponseResult() {}

// type = 'url'
type ParseResponseResultURLResultType string

const (
	ParseResponseResultURLResultTypeURL ParseResponseResultURLResultType = "url"
)

func (r ParseResponseResultURLResultType) IsKnown() bool {
	switch r {
	case ParseResponseResultURLResultTypeURL:
		return true
	}
	return false
}

// type = 'full'
type ParseResponseResultType string

const (
	ParseResponseResultTypeFull ParseResponseResultType = "full"
	ParseResponseResultTypeURL  ParseResponseResultType = "url"
)

func (r ParseResponseResultType) IsKnown() bool {
	switch r {
	case ParseResponseResultTypeFull, ParseResponseResultTypeURL:
		return true
	}
	return false
}

type RetrievalParam struct {
	Chunking param.Field[RetrievalChunkingParam] `json:"chunking"`
	// If True, use embedding optimized mode. Defaults to False.
	EmbeddingOptimized param.Field[bool] `json:"embedding_optimized"`
	// A list of block types to filter out from 'content' and 'embed' fields. By
	// default, no blocks are filtered.
	FilterBlocks param.Field[[]RetrievalFilterBlock] `json:"filter_blocks"`
}

func (r RetrievalParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RetrievalChunkingParam struct {
	// Choose how to partition chunks. Variable mode chunks by character length and
	// visual context. Section mode chunks by section headers. Page mode chunks
	// according to pages. Page sections mode chunks first by page, then by sections
	// within each page. Disabled returns one single chunk.
	ChunkMode param.Field[RetrievalChunkingChunkMode] `json:"chunk_mode"`
	// Number of characters of overlap to include from adjacent chunks. Defaults to 0.
	ChunkOverlap param.Field[int64] `json:"chunk_overlap"`
	// The approximate size of chunks (in characters) that the document will be split
	// into. Defaults to null, in which case the chunk size is variable between 250 -
	// 1500 characters.
	ChunkSize param.Field[int64] `json:"chunk_size"`
}

func (r RetrievalChunkingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Choose how to partition chunks. Variable mode chunks by character length and
// visual context. Section mode chunks by section headers. Page mode chunks
// according to pages. Page sections mode chunks first by page, then by sections
// within each page. Disabled returns one single chunk.
type RetrievalChunkingChunkMode string

const (
	RetrievalChunkingChunkModeVariable     RetrievalChunkingChunkMode = "variable"
	RetrievalChunkingChunkModeSection      RetrievalChunkingChunkMode = "section"
	RetrievalChunkingChunkModePage         RetrievalChunkingChunkMode = "page"
	RetrievalChunkingChunkModeDisabled     RetrievalChunkingChunkMode = "disabled"
	RetrievalChunkingChunkModeBlock        RetrievalChunkingChunkMode = "block"
	RetrievalChunkingChunkModePageSections RetrievalChunkingChunkMode = "page_sections"
)

func (r RetrievalChunkingChunkMode) IsKnown() bool {
	switch r {
	case RetrievalChunkingChunkModeVariable, RetrievalChunkingChunkModeSection, RetrievalChunkingChunkModePage, RetrievalChunkingChunkModeDisabled, RetrievalChunkingChunkModeBlock, RetrievalChunkingChunkModePageSections:
		return true
	}
	return false
}

type RetrievalFilterBlock string

const (
	RetrievalFilterBlockHeader        RetrievalFilterBlock = "Header"
	RetrievalFilterBlockFooter        RetrievalFilterBlock = "Footer"
	RetrievalFilterBlockTitle         RetrievalFilterBlock = "Title"
	RetrievalFilterBlockSectionHeader RetrievalFilterBlock = "Section Header"
	RetrievalFilterBlockPageNumber    RetrievalFilterBlock = "Page Number"
	RetrievalFilterBlockListItem      RetrievalFilterBlock = "List Item"
	RetrievalFilterBlockFigure        RetrievalFilterBlock = "Figure"
	RetrievalFilterBlockTable         RetrievalFilterBlock = "Table"
	RetrievalFilterBlockKeyValue      RetrievalFilterBlock = "Key Value"
	RetrievalFilterBlockText          RetrievalFilterBlock = "Text"
	RetrievalFilterBlockComment       RetrievalFilterBlock = "Comment"
	RetrievalFilterBlockSignature     RetrievalFilterBlock = "Signature"
)

func (r RetrievalFilterBlock) IsKnown() bool {
	switch r {
	case RetrievalFilterBlockHeader, RetrievalFilterBlockFooter, RetrievalFilterBlockTitle, RetrievalFilterBlockSectionHeader, RetrievalFilterBlockPageNumber, RetrievalFilterBlockListItem, RetrievalFilterBlockFigure, RetrievalFilterBlockTable, RetrievalFilterBlockKeyValue, RetrievalFilterBlockText, RetrievalFilterBlockComment, RetrievalFilterBlockSignature:
		return true
	}
	return false
}

type SettingsParam struct {
	// Password to decrypt password-protected documents.
	DocumentPassword param.Field[string] `json:"document_password"`
	// If True, embed OCR metadata into the returned PDF. Defaults to False.
	EmbedPdfMetadata param.Field[bool] `json:"embed_pdf_metadata"`
	// The mode to use for text extraction from PDFs. OCR mode uses optical character
	// recognition only. Hybrid mode combines OCR with embedded PDF text for best
	// accuracy (default).
	ExtractionMode param.Field[SettingsExtractionMode] `json:"extraction_mode"`
	// Force the URL to be downloaded as a specific file extension (e.g. `.png`).
	ForceFileExtension param.Field[string] `json:"force_file_extension"`
	// Force the result to be returned in URL form.
	ForceURLResult param.Field[bool] `json:"force_url_result"`
	// Standard is our best multilingual OCR system. Legacy only supports germanic
	// languages and is available for backwards compatibility.
	OcrSystem param.Field[SettingsOcrSystem] `json:"ocr_system"`
	// The page range to process (1-indexed). By default, the entire document is
	// processed. For spreadsheets, you can also provide a list of sheet names.
	PageRange param.Field[SettingsPageRangeUnionParam] `json:"page_range"`
	// If True, persist the results indefinitely. Defaults to False.
	PersistResults param.Field[bool] `json:"persist_results"`
	// Whether to return images for the specified block types. 'page' returns full page
	// images. By default, no images are returned.
	ReturnImages param.Field[[]SettingsReturnImage] `json:"return_images"`
	// If True, return OCR data in the result. Defaults to False.
	ReturnOcrData param.Field[bool] `json:"return_ocr_data"`
	// The timeout for the job in seconds.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r SettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The mode to use for text extraction from PDFs. OCR mode uses optical character
// recognition only. Hybrid mode combines OCR with embedded PDF text for best
// accuracy (default).
type SettingsExtractionMode string

const (
	SettingsExtractionModeOcr    SettingsExtractionMode = "ocr"
	SettingsExtractionModeHybrid SettingsExtractionMode = "hybrid"
)

func (r SettingsExtractionMode) IsKnown() bool {
	switch r {
	case SettingsExtractionModeOcr, SettingsExtractionModeHybrid:
		return true
	}
	return false
}

// Standard is our best multilingual OCR system. Legacy only supports germanic
// languages and is available for backwards compatibility.
type SettingsOcrSystem string

const (
	SettingsOcrSystemStandard SettingsOcrSystem = "standard"
	SettingsOcrSystemLegacy   SettingsOcrSystem = "legacy"
)

func (r SettingsOcrSystem) IsKnown() bool {
	switch r {
	case SettingsOcrSystemStandard, SettingsOcrSystemLegacy:
		return true
	}
	return false
}

// The page range to process (1-indexed). By default, the entire document is
// processed. For spreadsheets, you can also provide a list of sheet names.
//
// Satisfied by [PageRangeParam], [SettingsPageRangeArrayParam],
// [SettingsPageRangeArrayParam], [SettingsPageRangeArrayParam].
type SettingsPageRangeUnionParam interface {
	implementsSettingsPageRangeUnionParam()
}

type SettingsPageRangeArrayParam []PageRangeParam

func (r SettingsPageRangeArrayParam) implementsSettingsPageRangeUnionParam() {}

type SettingsReturnImage string

const (
	SettingsReturnImageFigure SettingsReturnImage = "figure"
	SettingsReturnImageTable  SettingsReturnImage = "table"
	SettingsReturnImagePage   SettingsReturnImage = "page"
)

func (r SettingsReturnImage) IsKnown() bool {
	switch r {
	case SettingsReturnImageFigure, SettingsReturnImageTable, SettingsReturnImagePage:
		return true
	}
	return false
}

type SpreadsheetParam struct {
	// In a spreadsheet with different tables inside, we enable splitting up the tables
	// by default. Accurate mode applies more powerful models for superior accuracy, at
	// 5× the default per-cell rate. Disabling will register as one large table.
	Clustering param.Field[SpreadsheetClustering] `json:"clustering"`
	// Whether to exclude hidden sheets, rows, or columns in the output.
	Exclude param.Field[[]SpreadsheetExclude] `json:"exclude"`
	// Whether to include cell color, formula, and dropdown information in the output.
	Include          param.Field[[]SpreadsheetInclude]             `json:"include"`
	SplitLargeTables param.Field[SpreadsheetSplitLargeTablesParam] `json:"split_large_tables"`
}

func (r SpreadsheetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// In a spreadsheet with different tables inside, we enable splitting up the tables
// by default. Accurate mode applies more powerful models for superior accuracy, at
// 5× the default per-cell rate. Disabling will register as one large table.
type SpreadsheetClustering string

const (
	SpreadsheetClusteringAccurate SpreadsheetClustering = "accurate"
	SpreadsheetClusteringFast     SpreadsheetClustering = "fast"
	SpreadsheetClusteringDisabled SpreadsheetClustering = "disabled"
)

func (r SpreadsheetClustering) IsKnown() bool {
	switch r {
	case SpreadsheetClusteringAccurate, SpreadsheetClusteringFast, SpreadsheetClusteringDisabled:
		return true
	}
	return false
}

type SpreadsheetExclude string

const (
	SpreadsheetExcludeHiddenSheets      SpreadsheetExclude = "hidden_sheets"
	SpreadsheetExcludeHiddenRows        SpreadsheetExclude = "hidden_rows"
	SpreadsheetExcludeHiddenCols        SpreadsheetExclude = "hidden_cols"
	SpreadsheetExcludeStyling           SpreadsheetExclude = "styling"
	SpreadsheetExcludeSpreadsheetImages SpreadsheetExclude = "spreadsheet_images"
)

func (r SpreadsheetExclude) IsKnown() bool {
	switch r {
	case SpreadsheetExcludeHiddenSheets, SpreadsheetExcludeHiddenRows, SpreadsheetExcludeHiddenCols, SpreadsheetExcludeStyling, SpreadsheetExcludeSpreadsheetImages:
		return true
	}
	return false
}

type SpreadsheetInclude string

const (
	SpreadsheetIncludeCellColors SpreadsheetInclude = "cell_colors"
	SpreadsheetIncludeFormula    SpreadsheetInclude = "formula"
	SpreadsheetIncludeDropdowns  SpreadsheetInclude = "dropdowns"
)

func (r SpreadsheetInclude) IsKnown() bool {
	switch r {
	case SpreadsheetIncludeCellColors, SpreadsheetIncludeFormula, SpreadsheetIncludeDropdowns:
		return true
	}
	return false
}

type SpreadsheetSplitLargeTablesParam struct {
	// If True, split large tables into smaller tables. Defaults to True.
	Enabled param.Field[bool] `json:"enabled"`
	// The size of the tables to split into. Defaults to 50. Use 'row' and 'column' to
	// independently specify the number of rows and columns to include when splitting.
	// If you only want to split by rows or columns, set the other value to None.
	Size param.Field[SpreadsheetSplitLargeTablesSizeUnionParam] `json:"size"`
}

func (r SpreadsheetSplitLargeTablesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The size of the tables to split into. Defaults to 50. Use 'row' and 'column' to
// independently specify the number of rows and columns to include when splitting.
// If you only want to split by rows or columns, set the other value to None.
//
// Satisfied by [shared.UnionInt],
// [SpreadsheetSplitLargeTablesSizeSplitLargeTableSizesParam].
type SpreadsheetSplitLargeTablesSizeUnionParam interface {
	ImplementsSpreadsheetSplitLargeTablesSizeUnionParam()
}

type SpreadsheetSplitLargeTablesSizeSplitLargeTableSizesParam struct {
	// The number of columns to include in each chunk when splitting large tables. Does
	// not chunk columns if set to None.
	Column param.Field[int64] `json:"column"`
	// The number of rows to include in each chunk when splitting large tables. Does
	// not chunk rows if set to None.
	Row param.Field[int64] `json:"row"`
}

func (r SpreadsheetSplitLargeTablesSizeSplitLargeTableSizesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SpreadsheetSplitLargeTablesSizeSplitLargeTableSizesParam) ImplementsSpreadsheetSplitLargeTablesSizeUnionParam() {
}

type ParseRunResponse struct {
	JobID string `json:"job_id" api:"required"`
	// The duration of the parse request in seconds.
	Duration float64 `json:"duration"`
	// The storage URL of the converted PDF file.
	PdfURL string `json:"pdf_url" api:"nullable"`
	// This field can have the runtime type of [ParseResponseResult].
	Result interface{} `json:"result"`
	// The link to the studio pipeline for the document.
	StudioLink string               `json:"studio_link" api:"nullable"`
	Usage      ParseUsage           `json:"usage"`
	JSON       parseRunResponseJSON `json:"-"`
	union      ParseRunResponseUnion
}

// parseRunResponseJSON contains the JSON metadata for the struct
// [ParseRunResponse]
type parseRunResponseJSON struct {
	JobID       apijson.Field
	Duration    apijson.Field
	PdfURL      apijson.Field
	Result      apijson.Field
	StudioLink  apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r parseRunResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ParseRunResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ParseRunResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ParseRunResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ParseResponse], [AsyncParseResponse].
func (r ParseRunResponse) AsUnion() ParseRunResponseUnion {
	return r.union
}

// Union satisfied by [ParseResponse] or [AsyncParseResponse].
type ParseRunResponseUnion interface {
	implementsParseRunResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ParseRunResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParseResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AsyncParseResponse{}),
		},
	)
}

type ParseRunParams struct {
	Body ParseRunParamsBodyUnion `json:"body" api:"required"`
}

func (r ParseRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ParseRunParamsBody struct {
	Input param.Field[interface{}] `json:"input" api:"required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async      param.Field[AsyncConfigV3Param] `json:"async"`
	Enhance    param.Field[EnhanceParam]       `json:"enhance"`
	Formatting param.Field[FormattingParam]    `json:"formatting"`
	// Queue priority. 'batch' for non-urgent work that processes when spare GPU
	// capacity is available.
	QueuePriority param.Field[ParseRunParamsBodyQueuePriority] `json:"queue_priority"`
	Retrieval     param.Field[RetrievalParam]                  `json:"retrieval"`
	Settings      param.Field[SettingsParam]                   `json:"settings"`
	Spreadsheet   param.Field[SpreadsheetParam]                `json:"spreadsheet"`
}

func (r ParseRunParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBody) implementsParseRunParamsBodyUnion() {}

// Satisfied by [ParseRunParamsBodySyncParseConfig], [AsyncParseConfigParam],
// [ParseRunParamsBody].
type ParseRunParamsBodyUnion interface {
	implementsParseRunParamsBodyUnion()
}

type ParseRunParamsBodySyncParseConfig struct {
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input       param.Field[ParseRunParamsBodySyncParseConfigInputUnion] `json:"input" api:"required"`
	Enhance     param.Field[EnhanceParam]                                `json:"enhance"`
	Formatting  param.Field[FormattingParam]                             `json:"formatting"`
	Retrieval   param.Field[RetrievalParam]                              `json:"retrieval"`
	Settings    param.Field[SettingsParam]                               `json:"settings"`
	Spreadsheet param.Field[SpreadsheetParam]                            `json:"spreadsheet"`
}

func (r ParseRunParamsBodySyncParseConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBodySyncParseConfig) implementsParseRunParamsBodyUnion() {}

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
// [ParseRunParamsBodySyncParseConfigInputArray], [shared.UploadParam].
type ParseRunParamsBodySyncParseConfigInputUnion interface {
	ImplementsParseRunParamsBodySyncParseConfigInputUnion()
}

type ParseRunParamsBodySyncParseConfigInputArray []string

func (r ParseRunParamsBodySyncParseConfigInputArray) ImplementsParseRunParamsBodySyncParseConfigInputUnion() {
}

// Queue priority. 'batch' for non-urgent work that processes when spare GPU
// capacity is available.
type ParseRunParamsBodyQueuePriority string

const (
	ParseRunParamsBodyQueuePriorityAuto  ParseRunParamsBodyQueuePriority = "auto"
	ParseRunParamsBodyQueuePriorityBatch ParseRunParamsBodyQueuePriority = "batch"
)

func (r ParseRunParamsBodyQueuePriority) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyQueuePriorityAuto, ParseRunParamsBodyQueuePriorityBatch:
		return true
	}
	return false
}

type ParseRunJobParams struct {
	AsyncParseConfig AsyncParseConfigParam `json:"async_parse_config" api:"required"`
}

func (r ParseRunJobParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AsyncParseConfig)
}
