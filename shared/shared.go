// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"

	"github.com/reductoai/reducto-go-sdk/internal/apijson"
	"github.com/reductoai/reducto-go-sdk/internal/param"
	"github.com/tidwall/gjson"
)

type BoundingBox struct {
	Height float64 `json:"height" api:"required"`
	Left   float64 `json:"left" api:"required"`
	// The page number of the bounding box (1-indexed).
	Page  int64   `json:"page" api:"required"`
	Top   float64 `json:"top" api:"required"`
	Width float64 `json:"width" api:"required"`
	// The page number in the original document of the bounding box (1-indexed).
	OriginalPage int64           `json:"original_page"`
	JSON         boundingBoxJSON `json:"-"`
}

// boundingBoxJSON contains the JSON metadata for the struct [BoundingBox]
type boundingBoxJSON struct {
	Height       apijson.Field
	Left         apijson.Field
	Page         apijson.Field
	Top          apijson.Field
	Width        apijson.Field
	OriginalPage apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BoundingBox) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r boundingBoxJSON) RawJSON() string {
	return r.raw
}

type BoundingBoxParam struct {
	Height param.Field[float64] `json:"height" api:"required"`
	Left   param.Field[float64] `json:"left" api:"required"`
	// The page number of the bounding box (1-indexed).
	Page  param.Field[int64]   `json:"page" api:"required"`
	Top   param.Field[float64] `json:"top" api:"required"`
	Width param.Field[float64] `json:"width" api:"required"`
	// The page number in the original document of the bounding box (1-indexed).
	OriginalPage param.Field[int64] `json:"original_page"`
}

func (r BoundingBoxParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChunkingParam struct {
	// Choose how to partition chunks. Variable mode chunks by character length and
	// visual context. Section mode chunks by section headers. Page mode chunks
	// according to pages. Page sections mode chunks first by page, then by sections
	// within each page. Disabled returns one single chunk.
	ChunkMode param.Field[ChunkingChunkMode] `json:"chunk_mode"`
	// Number of characters of overlap to include from adjacent chunks. Defaults to 0.
	ChunkOverlap param.Field[int64] `json:"chunk_overlap"`
	// The approximate size of chunks (in characters) that the document will be split
	// into. Defaults to null, in which case the chunk size is variable between 250 -
	// 1500 characters.
	ChunkSize param.Field[int64] `json:"chunk_size"`
}

func (r ChunkingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Choose how to partition chunks. Variable mode chunks by character length and
// visual context. Section mode chunks by section headers. Page mode chunks
// according to pages. Page sections mode chunks first by page, then by sections
// within each page. Disabled returns one single chunk.
type ChunkingChunkMode string

const (
	ChunkingChunkModeVariable     ChunkingChunkMode = "variable"
	ChunkingChunkModeSection      ChunkingChunkMode = "section"
	ChunkingChunkModePage         ChunkingChunkMode = "page"
	ChunkingChunkModeDisabled     ChunkingChunkMode = "disabled"
	ChunkingChunkModeBlock        ChunkingChunkMode = "block"
	ChunkingChunkModePageSections ChunkingChunkMode = "page_sections"
)

func (r ChunkingChunkMode) IsKnown() bool {
	switch r {
	case ChunkingChunkModeVariable, ChunkingChunkModeSection, ChunkingChunkModePage, ChunkingChunkModeDisabled, ChunkingChunkModeBlock, ChunkingChunkModePageSections:
		return true
	}
	return false
}

type ConfigV3AsyncConfigParam struct {
	// JSON metadata included in webhook request body. Defaults to None.
	Metadata param.Field[interface{}] `json:"metadata"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool] `json:"priority"`
	// The webhook configuration for the asynchronous processing.
	Webhook param.Field[ConfigV3AsyncConfigWebhookUnionParam] `json:"webhook"`
}

func (r ConfigV3AsyncConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The webhook configuration for the asynchronous processing.
type ConfigV3AsyncConfigWebhookParam struct {
	Channels param.Field[interface{}]                    `json:"channels"`
	Mode     param.Field[ConfigV3AsyncConfigWebhookMode] `json:"mode"`
	URL      param.Field[string]                         `json:"url"`
}

func (r ConfigV3AsyncConfigWebhookParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigV3AsyncConfigWebhookParam) ImplementsConfigV3AsyncConfigWebhookUnionParam() {}

// The webhook configuration for the asynchronous processing.
//
// Satisfied by [shared.SvixWebhookConfigParam], [shared.DirectWebhookConfigParam],
// [ConfigV3AsyncConfigWebhookParam].
type ConfigV3AsyncConfigWebhookUnionParam interface {
	ImplementsConfigV3AsyncConfigWebhookUnionParam()
}

type ConfigV3AsyncConfigWebhookMode string

const (
	ConfigV3AsyncConfigWebhookModeSvix   ConfigV3AsyncConfigWebhookMode = "svix"
	ConfigV3AsyncConfigWebhookModeDirect ConfigV3AsyncConfigWebhookMode = "direct"
)

func (r ConfigV3AsyncConfigWebhookMode) IsKnown() bool {
	switch r {
	case ConfigV3AsyncConfigWebhookModeSvix, ConfigV3AsyncConfigWebhookModeDirect:
		return true
	}
	return false
}

type DirectWebhookConfigParam struct {
	URL  param.Field[string]                  `json:"url" api:"required"`
	Mode param.Field[DirectWebhookConfigMode] `json:"mode"`
}

func (r DirectWebhookConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectWebhookConfigParam) ImplementsConfigV3AsyncConfigWebhookUnionParam() {}

type DirectWebhookConfigMode string

const (
	DirectWebhookConfigModeDirect DirectWebhookConfigMode = "direct"
)

func (r DirectWebhookConfigMode) IsKnown() bool {
	switch r {
	case DirectWebhookConfigModeDirect:
		return true
	}
	return false
}

type EditResponse struct {
	// Presigned URL to download the edited document.
	DocumentURL string `json:"document_url" api:"required"`
	// Form schema for PDF forms. List of widgets with their types, descriptions, and
	// bounding boxes.
	FormSchema []EditResponseFormSchema `json:"form_schema" api:"nullable"`
	// Usage information for the edit operation, including number of pages and credits
	// charged.
	Usage ParseUsage       `json:"usage" api:"nullable"`
	JSON  editResponseJSON `json:"-"`
}

// editResponseJSON contains the JSON metadata for the struct [EditResponse]
type editResponseJSON struct {
	DocumentURL apijson.Field
	FormSchema  apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editResponseJSON) RawJSON() string {
	return r.raw
}

func (r EditResponse) ImplementsJobGetResponseAsyncJobResponseResult() {}

func (r EditResponse) ImplementsJobGetResponseEnhancedAsyncJobResponseResult() {}

type EditResponseFormSchema struct {
	// Bounding box coordinates of the widget
	Bbox BoundingBox `json:"bbox" api:"required"`
	// Description of the widget extracted from the document
	Description string `json:"description" api:"required"`
	// Type of the form widget
	Type EditResponseFormSchemaType `json:"type" api:"required"`
	// If True (default), the system will attempt to fill this widget. If False, the
	// widget will be created but intentionally left unfilled.
	Fill bool `json:"fill"`
	// If provided, this value will be used directly instead of attempting to
	// intelligently determine the field value.
	Value string                     `json:"value" api:"nullable"`
	JSON  editResponseFormSchemaJSON `json:"-"`
}

// editResponseFormSchemaJSON contains the JSON metadata for the struct
// [EditResponseFormSchema]
type editResponseFormSchemaJSON struct {
	Bbox        apijson.Field
	Description apijson.Field
	Type        apijson.Field
	Fill        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EditResponseFormSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editResponseFormSchemaJSON) RawJSON() string {
	return r.raw
}

// Type of the form widget
type EditResponseFormSchemaType string

const (
	EditResponseFormSchemaTypeText     EditResponseFormSchemaType = "text"
	EditResponseFormSchemaTypeCheckbox EditResponseFormSchemaType = "checkbox"
	EditResponseFormSchemaTypeRadio    EditResponseFormSchemaType = "radio"
	EditResponseFormSchemaTypeDropdown EditResponseFormSchemaType = "dropdown"
	EditResponseFormSchemaTypeBarcode  EditResponseFormSchemaType = "barcode"
)

func (r EditResponseFormSchemaType) IsKnown() bool {
	switch r {
	case EditResponseFormSchemaTypeText, EditResponseFormSchemaTypeCheckbox, EditResponseFormSchemaTypeRadio, EditResponseFormSchemaTypeDropdown, EditResponseFormSchemaTypeBarcode:
		return true
	}
	return false
}

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

func (r EnhanceAgenticParam) ImplementsEnhanceAgenticUnionParam() {}

// Satisfied by [shared.TableAgenticParam], [shared.FigureAgenticParam],
// [shared.TextAgenticParam], [EnhanceAgenticParam].
type EnhanceAgenticUnionParam interface {
	ImplementsEnhanceAgenticUnionParam()
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

type ExtractResponse struct {
	// The citations corresponding to the extracted response.
	Citations []interface{} `json:"citations" api:"required,nullable"`
	// The extracted response in your provided schema. This is a list of dictionaries.
	// If disable_chunking is True (default), then it will be a list of length one.
	Result []interface{} `json:"result" api:"required"`
	Usage  ExtractUsage  `json:"usage" api:"required"`
	JobID  string        `json:"job_id" api:"nullable"`
	// The link to the studio pipeline for the document.
	StudioLink string              `json:"studio_link" api:"nullable"`
	JSON       extractResponseJSON `json:"-"`
}

// extractResponseJSON contains the JSON metadata for the struct [ExtractResponse]
type extractResponseJSON struct {
	Citations   apijson.Field
	Result      apijson.Field
	Usage       apijson.Field
	JobID       apijson.Field
	StudioLink  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtractResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extractResponseJSON) RawJSON() string {
	return r.raw
}

func (r ExtractResponse) ImplementsPipelineResponseResultExtractUnion() {}

func (r ExtractResponse) ImplementsPipelineResponseResultExtractArrayResult() {}

func (r ExtractResponse) ImplementsJobGetResponseAsyncJobResponseResult() {}

func (r ExtractResponse) ImplementsJobGetResponseEnhancedAsyncJobResponseResult() {}

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

type FigureAgenticParam struct {
	Scope param.Field[FigureAgenticScope] `json:"scope" api:"required"`
	// If True, use the advanced chart agent. Defaults to False.
	AdvancedChartAgent param.Field[bool] `json:"advanced_chart_agent"`
	// Custom prompt for figure agentic.
	Prompt param.Field[string] `json:"prompt"`
	// If True, return overlays for the figure. This is so you can use the overlays to
	// double check the quality of the extraction
	ReturnOverlays param.Field[bool] `json:"return_overlays"`
}

func (r FigureAgenticParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FigureAgenticParam) ImplementsEnhanceAgenticUnionParam() {}

type FigureAgenticScope string

const (
	FigureAgenticScopeFigure FigureAgenticScope = "figure"
)

func (r FigureAgenticScope) IsKnown() bool {
	switch r {
	case FigureAgenticScopeFigure:
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

type PageRangeParam struct {
	// The page number to stop processing at (1-indexed).
	End param.Field[int64] `json:"end"`
	// The page number to start processing from (1-indexed).
	Start param.Field[int64] `json:"start"`
}

func (r PageRangeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PageRangeParam) ImplementsSettingsPageRangeUnionParam() {}

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

func (r ParseResponse) ImplementsPipelineResponseResultParseUnion() {}

func (r ParseResponse) ImplementsJobGetResponseAsyncJobResponseResult() {}

func (r ParseResponse) ImplementsJobGetResponseEnhancedAsyncJobResponseResult() {}

func (r ParseResponse) ImplementsParseRunResponse() {}

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

type ParseUsage struct {
	NumPages        int64              `json:"num_pages" api:"required"`
	CreditBreakdown map[string]float64 `json:"credit_breakdown" api:"nullable"`
	Credits         float64            `json:"credits" api:"nullable"`
	JSON            parseUsageJSON     `json:"-"`
}

// parseUsageJSON contains the JSON metadata for the struct [ParseUsage]
type parseUsageJSON struct {
	NumPages        apijson.Field
	CreditBreakdown apijson.Field
	Credits         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ParseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseUsageJSON) RawJSON() string {
	return r.raw
}

type PipelineResponse struct {
	JobID  string                 `json:"job_id" api:"required"`
	Result PipelineResponseResult `json:"result" api:"required"`
	Usage  ParseUsage             `json:"usage" api:"required"`
	JSON   pipelineResponseJSON   `json:"-"`
}

// pipelineResponseJSON contains the JSON metadata for the struct
// [PipelineResponse]
type pipelineResponseJSON struct {
	JobID       apijson.Field
	Result      apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineResponseJSON) RawJSON() string {
	return r.raw
}

func (r PipelineResponse) ImplementsJobGetResponseAsyncJobResponseResult() {}

func (r PipelineResponse) ImplementsJobGetResponseEnhancedAsyncJobResponseResult() {}

type PipelineResponseResult struct {
	Extract PipelineResponseResultExtractUnion `json:"extract" api:"required,nullable"`
	Parse   PipelineResponseResultParseUnion   `json:"parse" api:"required,nullable"`
	Split   SplitResponse                      `json:"split" api:"required,nullable"`
	Edit    EditResponse                       `json:"edit" api:"nullable"`
	JSON    pipelineResponseResultJSON         `json:"-"`
}

// pipelineResponseResultJSON contains the JSON metadata for the struct
// [PipelineResponseResult]
type pipelineResponseResultJSON struct {
	Extract     apijson.Field
	Parse       apijson.Field
	Split       apijson.Field
	Edit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineResponseResultJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [PipelineResponseResultExtractArray], [ExtractResponse] or
// [V3ExtractResponse].
type PipelineResponseResultExtractUnion interface {
	ImplementsPipelineResponseResultExtractUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PipelineResponseResultExtractUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineResponseResultExtractArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExtractResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V3ExtractResponse{}),
		},
	)
}

type PipelineResponseResultExtractArray []PipelineResponseResultExtractArrayItem

func (r PipelineResponseResultExtractArray) ImplementsPipelineResponseResultExtractUnion() {}

// This is the response format for Extract -> Split Pipelines
type PipelineResponseResultExtractArrayItem struct {
	PageRange []int64                                    `json:"page_range" api:"required"`
	Result    PipelineResponseResultExtractArrayResult   `json:"result" api:"required"`
	SplitName string                                     `json:"split_name" api:"required"`
	Partition string                                     `json:"partition" api:"nullable"`
	JSON      pipelineResponseResultExtractArrayItemJSON `json:"-"`
}

// pipelineResponseResultExtractArrayItemJSON contains the JSON metadata for the
// struct [PipelineResponseResultExtractArrayItem]
type pipelineResponseResultExtractArrayItemJSON struct {
	PageRange   apijson.Field
	Result      apijson.Field
	SplitName   apijson.Field
	Partition   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineResponseResultExtractArrayItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineResponseResultExtractArrayItemJSON) RawJSON() string {
	return r.raw
}

type PipelineResponseResultExtractArrayResult struct {
	// This field can have the runtime type of [[]interface{}].
	Result interface{}  `json:"result" api:"required"`
	Usage  ExtractUsage `json:"usage" api:"required"`
	// This field can have the runtime type of [[]interface{}].
	Citations interface{} `json:"citations"`
	JobID     string      `json:"job_id" api:"nullable"`
	// The link to the studio pipeline for the document.
	StudioLink string                                       `json:"studio_link" api:"nullable"`
	JSON       pipelineResponseResultExtractArrayResultJSON `json:"-"`
	union      PipelineResponseResultExtractArrayResultUnion
}

// pipelineResponseResultExtractArrayResultJSON contains the JSON metadata for the
// struct [PipelineResponseResultExtractArrayResult]
type pipelineResponseResultExtractArrayResultJSON struct {
	Result      apijson.Field
	Usage       apijson.Field
	Citations   apijson.Field
	JobID       apijson.Field
	StudioLink  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r pipelineResponseResultExtractArrayResultJSON) RawJSON() string {
	return r.raw
}

func (r *PipelineResponseResultExtractArrayResult) UnmarshalJSON(data []byte) (err error) {
	*r = PipelineResponseResultExtractArrayResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PipelineResponseResultExtractArrayResultUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [ExtractResponse], [V3ExtractResponse].
func (r PipelineResponseResultExtractArrayResult) AsUnion() PipelineResponseResultExtractArrayResultUnion {
	return r.union
}

// Union satisfied by [ExtractResponse] or [V3ExtractResponse].
type PipelineResponseResultExtractArrayResultUnion interface {
	ImplementsPipelineResponseResultExtractArrayResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PipelineResponseResultExtractArrayResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExtractResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V3ExtractResponse{}),
		},
	)
}

// Union satisfied by [ParseResponse] or [PipelineResponseResultParseArray].
type PipelineResponseResultParseUnion interface {
	ImplementsPipelineResponseResultParseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PipelineResponseResultParseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParseResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineResponseResultParseArray{}),
		},
	)
}

type PipelineResponseResultParseArray []ParseResponse

func (r PipelineResponseResultParseArray) ImplementsPipelineResponseResultParseUnion() {}

type RetrievalParam struct {
	Chunking param.Field[ChunkingParam] `json:"chunking"`
	// If True, use embedding optimized mode. Defaults to False.
	EmbeddingOptimized param.Field[bool] `json:"embedding_optimized"`
	// A list of block types to filter out from 'content' and 'embed' fields. By
	// default, no blocks are filtered.
	FilterBlocks param.Field[[]RetrievalFilterBlock] `json:"filter_blocks"`
}

func (r RetrievalParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
// Satisfied by [shared.PageRangeParam], [shared.SettingsPageRangeArrayParam],
// [shared.SettingsPageRangeArrayParam], [shared.SettingsPageRangeArrayParam].
type SettingsPageRangeUnionParam interface {
	ImplementsSettingsPageRangeUnionParam()
}

type SettingsPageRangeArrayParam []PageRangeParam

func (r SettingsPageRangeArrayParam) ImplementsSettingsPageRangeUnionParam() {}

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

type SplitCategoryParam struct {
	Description  param.Field[string] `json:"description" api:"required"`
	Name         param.Field[string] `json:"name" api:"required"`
	PartitionKey param.Field[string] `json:"partition_key"`
}

func (r SplitCategoryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SplitLargeTablesParam struct {
	// If True, split large tables into smaller tables. Defaults to True.
	Enabled param.Field[bool] `json:"enabled"`
	// The size of the tables to split into. Defaults to 50. Use 'row' and 'column' to
	// independently specify the number of rows and columns to include when splitting.
	// If you only want to split by rows or columns, set the other value to None.
	Size param.Field[SplitLargeTablesSizeUnionParam] `json:"size"`
}

func (r SplitLargeTablesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The size of the tables to split into. Defaults to 50. Use 'row' and 'column' to
// independently specify the number of rows and columns to include when splitting.
// If you only want to split by rows or columns, set the other value to None.
//
// Satisfied by [shared.UnionInt],
// [shared.SplitLargeTablesSizeSplitLargeTableSizesParam].
type SplitLargeTablesSizeUnionParam interface {
	ImplementsSplitLargeTablesSizeUnionParam()
}

type SplitLargeTablesSizeSplitLargeTableSizesParam struct {
	// The number of columns to include in each chunk when splitting large tables. Does
	// not chunk columns if set to None.
	Column param.Field[int64] `json:"column"`
	// The number of rows to include in each chunk when splitting large tables. Does
	// not chunk rows if set to None.
	Row param.Field[int64] `json:"row"`
}

func (r SplitLargeTablesSizeSplitLargeTableSizesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SplitLargeTablesSizeSplitLargeTableSizesParam) ImplementsSplitLargeTablesSizeUnionParam() {}

type SplitResponse struct {
	// The split result.
	Result SplitResponseResult `json:"result" api:"required"`
	Usage  ParseUsage          `json:"usage" api:"required"`
	JSON   splitResponseJSON   `json:"-"`
}

// splitResponseJSON contains the JSON metadata for the struct [SplitResponse]
type splitResponseJSON struct {
	Result      apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitResponseJSON) RawJSON() string {
	return r.raw
}

func (r SplitResponse) ImplementsJobGetResponseAsyncJobResponseResult() {}

func (r SplitResponse) ImplementsJobGetResponseEnhancedAsyncJobResponseResult() {}

// The split result.
type SplitResponseResult struct {
	// This field can have the runtime type of [[]SplitResponseResultSplitResultSplit],
	// [[]SplitResponseResultDeepSplitResultSplit].
	Splits interface{} `json:"splits" api:"required"`
	// This field can have the runtime type of [map[string][]int64].
	SectionMapping interface{}             `json:"section_mapping"`
	JSON           splitResponseResultJSON `json:"-"`
	union          SplitResponseResultUnion
}

// splitResponseResultJSON contains the JSON metadata for the struct
// [SplitResponseResult]
type splitResponseResultJSON struct {
	Splits         apijson.Field
	SectionMapping apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r splitResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *SplitResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = SplitResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SplitResponseResultUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [SplitResponseResultSplitResult],
// [SplitResponseResultDeepSplitResult].
func (r SplitResponseResult) AsUnion() SplitResponseResultUnion {
	return r.union
}

// The split result.
//
// Union satisfied by [SplitResponseResultSplitResult] or
// [SplitResponseResultDeepSplitResult].
type SplitResponseResultUnion interface {
	implementsSplitResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SplitResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplitResponseResultSplitResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplitResponseResultDeepSplitResult{}),
		},
	)
}

type SplitResponseResultSplitResult struct {
	SectionMapping map[string][]int64                    `json:"section_mapping" api:"required,nullable"`
	Splits         []SplitResponseResultSplitResultSplit `json:"splits" api:"required"`
	JSON           splitResponseResultSplitResultJSON    `json:"-"`
}

// splitResponseResultSplitResultJSON contains the JSON metadata for the struct
// [SplitResponseResultSplitResult]
type splitResponseResultSplitResultJSON struct {
	SectionMapping apijson.Field
	Splits         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SplitResponseResultSplitResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitResponseResultSplitResultJSON) RawJSON() string {
	return r.raw
}

func (r SplitResponseResultSplitResult) implementsSplitResponseResult() {}

type SplitResponseResultSplitResultSplit struct {
	Name       string                                          `json:"name" api:"required"`
	Pages      []int64                                         `json:"pages" api:"required"`
	Conf       SplitResponseResultSplitResultSplitsConf        `json:"conf"`
	Partitions []SplitResponseResultSplitResultSplitsPartition `json:"partitions" api:"nullable"`
	JSON       splitResponseResultSplitResultSplitJSON         `json:"-"`
}

// splitResponseResultSplitResultSplitJSON contains the JSON metadata for the
// struct [SplitResponseResultSplitResultSplit]
type splitResponseResultSplitResultSplitJSON struct {
	Name        apijson.Field
	Pages       apijson.Field
	Conf        apijson.Field
	Partitions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitResponseResultSplitResultSplit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitResponseResultSplitResultSplitJSON) RawJSON() string {
	return r.raw
}

type SplitResponseResultSplitResultSplitsConf string

const (
	SplitResponseResultSplitResultSplitsConfHigh SplitResponseResultSplitResultSplitsConf = "high"
	SplitResponseResultSplitResultSplitsConfLow  SplitResponseResultSplitResultSplitsConf = "low"
)

func (r SplitResponseResultSplitResultSplitsConf) IsKnown() bool {
	switch r {
	case SplitResponseResultSplitResultSplitsConfHigh, SplitResponseResultSplitResultSplitsConfLow:
		return true
	}
	return false
}

type SplitResponseResultSplitResultSplitsPartition struct {
	Name  string                                             `json:"name" api:"required"`
	Pages []int64                                            `json:"pages" api:"required"`
	Conf  SplitResponseResultSplitResultSplitsPartitionsConf `json:"conf"`
	JSON  splitResponseResultSplitResultSplitsPartitionJSON  `json:"-"`
}

// splitResponseResultSplitResultSplitsPartitionJSON contains the JSON metadata for
// the struct [SplitResponseResultSplitResultSplitsPartition]
type splitResponseResultSplitResultSplitsPartitionJSON struct {
	Name        apijson.Field
	Pages       apijson.Field
	Conf        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitResponseResultSplitResultSplitsPartition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitResponseResultSplitResultSplitsPartitionJSON) RawJSON() string {
	return r.raw
}

type SplitResponseResultSplitResultSplitsPartitionsConf string

const (
	SplitResponseResultSplitResultSplitsPartitionsConfHigh SplitResponseResultSplitResultSplitsPartitionsConf = "high"
	SplitResponseResultSplitResultSplitsPartitionsConfLow  SplitResponseResultSplitResultSplitsPartitionsConf = "low"
)

func (r SplitResponseResultSplitResultSplitsPartitionsConf) IsKnown() bool {
	switch r {
	case SplitResponseResultSplitResultSplitsPartitionsConfHigh, SplitResponseResultSplitResultSplitsPartitionsConfLow:
		return true
	}
	return false
}

type SplitResponseResultDeepSplitResult struct {
	Splits []SplitResponseResultDeepSplitResultSplit `json:"splits" api:"required"`
	JSON   splitResponseResultDeepSplitResultJSON    `json:"-"`
}

// splitResponseResultDeepSplitResultJSON contains the JSON metadata for the struct
// [SplitResponseResultDeepSplitResult]
type splitResponseResultDeepSplitResultJSON struct {
	Splits      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitResponseResultDeepSplitResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitResponseResultDeepSplitResultJSON) RawJSON() string {
	return r.raw
}

func (r SplitResponseResultDeepSplitResult) implementsSplitResponseResult() {}

type SplitResponseResultDeepSplitResultSplit struct {
	Name       string                                              `json:"name" api:"required"`
	Pages      []SplitResponseResultDeepSplitResultSplitsPage      `json:"pages" api:"required"`
	Partitions []SplitResponseResultDeepSplitResultSplitsPartition `json:"partitions" api:"nullable"`
	JSON       splitResponseResultDeepSplitResultSplitJSON         `json:"-"`
}

// splitResponseResultDeepSplitResultSplitJSON contains the JSON metadata for the
// struct [SplitResponseResultDeepSplitResultSplit]
type splitResponseResultDeepSplitResultSplitJSON struct {
	Name        apijson.Field
	Pages       apijson.Field
	Partitions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitResponseResultDeepSplitResultSplit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitResponseResultDeepSplitResultSplitJSON) RawJSON() string {
	return r.raw
}

type SplitResponseResultDeepSplitResultSplitsPage struct {
	Evidence   string                                                  `json:"evidence" api:"required"`
	PageNumber int64                                                   `json:"page_number" api:"required"`
	Confidence SplitResponseResultDeepSplitResultSplitsPagesConfidence `json:"confidence" api:"nullable"`
	JSON       splitResponseResultDeepSplitResultSplitsPageJSON        `json:"-"`
}

// splitResponseResultDeepSplitResultSplitsPageJSON contains the JSON metadata for
// the struct [SplitResponseResultDeepSplitResultSplitsPage]
type splitResponseResultDeepSplitResultSplitsPageJSON struct {
	Evidence    apijson.Field
	PageNumber  apijson.Field
	Confidence  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitResponseResultDeepSplitResultSplitsPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitResponseResultDeepSplitResultSplitsPageJSON) RawJSON() string {
	return r.raw
}

type SplitResponseResultDeepSplitResultSplitsPagesConfidence string

const (
	SplitResponseResultDeepSplitResultSplitsPagesConfidenceHigh   SplitResponseResultDeepSplitResultSplitsPagesConfidence = "high"
	SplitResponseResultDeepSplitResultSplitsPagesConfidenceMedium SplitResponseResultDeepSplitResultSplitsPagesConfidence = "medium"
	SplitResponseResultDeepSplitResultSplitsPagesConfidenceLow    SplitResponseResultDeepSplitResultSplitsPagesConfidence = "low"
)

func (r SplitResponseResultDeepSplitResultSplitsPagesConfidence) IsKnown() bool {
	switch r {
	case SplitResponseResultDeepSplitResultSplitsPagesConfidenceHigh, SplitResponseResultDeepSplitResultSplitsPagesConfidenceMedium, SplitResponseResultDeepSplitResultSplitsPagesConfidenceLow:
		return true
	}
	return false
}

type SplitResponseResultDeepSplitResultSplitsPartition struct {
	Name  string                                                   `json:"name" api:"required"`
	Pages []SplitResponseResultDeepSplitResultSplitsPartitionsPage `json:"pages" api:"required"`
	JSON  splitResponseResultDeepSplitResultSplitsPartitionJSON    `json:"-"`
}

// splitResponseResultDeepSplitResultSplitsPartitionJSON contains the JSON metadata
// for the struct [SplitResponseResultDeepSplitResultSplitsPartition]
type splitResponseResultDeepSplitResultSplitsPartitionJSON struct {
	Name        apijson.Field
	Pages       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitResponseResultDeepSplitResultSplitsPartition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitResponseResultDeepSplitResultSplitsPartitionJSON) RawJSON() string {
	return r.raw
}

type SplitResponseResultDeepSplitResultSplitsPartitionsPage struct {
	Evidence   string                                                            `json:"evidence" api:"required"`
	PageNumber int64                                                             `json:"page_number" api:"required"`
	Confidence SplitResponseResultDeepSplitResultSplitsPartitionsPagesConfidence `json:"confidence" api:"nullable"`
	JSON       splitResponseResultDeepSplitResultSplitsPartitionsPageJSON        `json:"-"`
}

// splitResponseResultDeepSplitResultSplitsPartitionsPageJSON contains the JSON
// metadata for the struct [SplitResponseResultDeepSplitResultSplitsPartitionsPage]
type splitResponseResultDeepSplitResultSplitsPartitionsPageJSON struct {
	Evidence    apijson.Field
	PageNumber  apijson.Field
	Confidence  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitResponseResultDeepSplitResultSplitsPartitionsPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitResponseResultDeepSplitResultSplitsPartitionsPageJSON) RawJSON() string {
	return r.raw
}

type SplitResponseResultDeepSplitResultSplitsPartitionsPagesConfidence string

const (
	SplitResponseResultDeepSplitResultSplitsPartitionsPagesConfidenceHigh   SplitResponseResultDeepSplitResultSplitsPartitionsPagesConfidence = "high"
	SplitResponseResultDeepSplitResultSplitsPartitionsPagesConfidenceMedium SplitResponseResultDeepSplitResultSplitsPartitionsPagesConfidence = "medium"
	SplitResponseResultDeepSplitResultSplitsPartitionsPagesConfidenceLow    SplitResponseResultDeepSplitResultSplitsPartitionsPagesConfidence = "low"
)

func (r SplitResponseResultDeepSplitResultSplitsPartitionsPagesConfidence) IsKnown() bool {
	switch r {
	case SplitResponseResultDeepSplitResultSplitsPartitionsPagesConfidenceHigh, SplitResponseResultDeepSplitResultSplitsPartitionsPagesConfidenceMedium, SplitResponseResultDeepSplitResultSplitsPartitionsPagesConfidenceLow:
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
	Include          param.Field[[]SpreadsheetInclude]  `json:"include"`
	SplitLargeTables param.Field[SplitLargeTablesParam] `json:"split_large_tables"`
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

type SvixWebhookConfigParam struct {
	// A list of Svix channels the message will be delivered down, omit to send to all
	// channels.
	Channels param.Field[[]string]              `json:"channels"`
	Mode     param.Field[SvixWebhookConfigMode] `json:"mode"`
}

func (r SvixWebhookConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SvixWebhookConfigParam) ImplementsConfigV3AsyncConfigWebhookUnionParam() {}

type SvixWebhookConfigMode string

const (
	SvixWebhookConfigModeSvix SvixWebhookConfigMode = "svix"
)

func (r SvixWebhookConfigMode) IsKnown() bool {
	switch r {
	case SvixWebhookConfigModeSvix:
		return true
	}
	return false
}

type TableAgenticParam struct {
	Scope param.Field[TableAgenticScope] `json:"scope" api:"required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r TableAgenticParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TableAgenticParam) ImplementsEnhanceAgenticUnionParam() {}

type TableAgenticScope string

const (
	TableAgenticScopeTable TableAgenticScope = "table"
)

func (r TableAgenticScope) IsKnown() bool {
	switch r {
	case TableAgenticScopeTable:
		return true
	}
	return false
}

type TextAgenticParam struct {
	Scope param.Field[TextAgenticScope] `json:"scope" api:"required"`
	// Custom instructions for agentic text. Note: This only applies to form regions
	// (key-value).
	Prompt param.Field[string] `json:"prompt"`
}

func (r TextAgenticParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TextAgenticParam) ImplementsEnhanceAgenticUnionParam() {}

type TextAgenticScope string

const (
	TextAgenticScopeText TextAgenticScope = "text"
)

func (r TextAgenticScope) IsKnown() bool {
	switch r {
	case TextAgenticScopeText:
		return true
	}
	return false
}

type Upload struct {
	FileID       string     `json:"file_id" api:"required"`
	PresignedURL string     `json:"presigned_url" api:"nullable"`
	JSON         uploadJSON `json:"-"`
}

// uploadJSON contains the JSON metadata for the struct [Upload]
type uploadJSON struct {
	FileID       apijson.Field
	PresignedURL apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Upload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uploadJSON) RawJSON() string {
	return r.raw
}

type UploadParam struct {
	FileID       param.Field[string] `json:"file_id" api:"required"`
	PresignedURL param.Field[string] `json:"presigned_url"`
}

func (r UploadParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UploadParam) ImplementsSplitRunParamsInputUnion() {}

func (r UploadParam) ImplementsSplitRunJobParamsInputUnion() {}

func (r UploadParam) ImplementsParseRunParamsBodySyncParseConfigInputUnion() {}

func (r UploadParam) ImplementsParseRunParamsBodyAsyncParseConfigInputUnion() {}

func (r UploadParam) ImplementsParseRunJobParamsInputUnion() {}

func (r UploadParam) ImplementsExtractRunParamsBodySyncExtractConfigInputUnion() {}

func (r UploadParam) ImplementsExtractRunParamsBodyAsyncExtractConfigInputUnion() {}

func (r UploadParam) ImplementsExtractRunJobParamsInputUnion() {}

func (r UploadParam) ImplementsEditRunParamsDocumentURLUnion() {}

func (r UploadParam) ImplementsEditRunJobParamsDocumentURLUnion() {}

func (r UploadParam) ImplementsPipelineRunParamsInputUnion() {}

func (r UploadParam) ImplementsPipelineRunJobParamsInputUnion() {}

type V3ExtractResponse struct {
	// The extracted response in your provided schema. This is a list of dictionaries.
	// If disable_chunking is True (default), then it will be a list of length one.
	Result []interface{} `json:"result" api:"required"`
	Usage  ExtractUsage  `json:"usage" api:"required"`
	JobID  string        `json:"job_id" api:"nullable"`
	// The link to the studio pipeline for the document.
	StudioLink string                `json:"studio_link" api:"nullable"`
	JSON       v3ExtractResponseJSON `json:"-"`
}

// v3ExtractResponseJSON contains the JSON metadata for the struct
// [V3ExtractResponse]
type v3ExtractResponseJSON struct {
	Result      apijson.Field
	Usage       apijson.Field
	JobID       apijson.Field
	StudioLink  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3ExtractResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3ExtractResponseJSON) RawJSON() string {
	return r.raw
}

func (r V3ExtractResponse) ImplementsPipelineResponseResultExtractUnion() {}

func (r V3ExtractResponse) ImplementsPipelineResponseResultExtractArrayResult() {}

func (r V3ExtractResponse) ImplementsJobGetResponseAsyncJobResponseResult() {}

func (r V3ExtractResponse) ImplementsJobGetResponseEnhancedAsyncJobResponseResult() {}

func (r V3ExtractResponse) ImplementsExtractRunResponse() {}

type WebhookConfigNewParam struct {
	// A list of Svix channels the message will be delivered down, omit to send to all
	// channels.
	Channels param.Field[[]string] `json:"channels"`
	// JSON metadata included in webhook request body
	Metadata param.Field[interface{}] `json:"metadata"`
	// The mode to use for webhook delivery. Defaults to 'disabled'. We recommend using
	// 'svix' for production environments.
	Mode param.Field[WebhookConfigNewMode] `json:"mode"`
	// The URL to send the webhook to (if using direct webhoook).
	URL param.Field[string] `json:"url"`
}

func (r WebhookConfigNewParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The mode to use for webhook delivery. Defaults to 'disabled'. We recommend using
// 'svix' for production environments.
type WebhookConfigNewMode string

const (
	WebhookConfigNewModeDisabled WebhookConfigNewMode = "disabled"
	WebhookConfigNewModeSvix     WebhookConfigNewMode = "svix"
	WebhookConfigNewModeDirect   WebhookConfigNewMode = "direct"
)

func (r WebhookConfigNewMode) IsKnown() bool {
	switch r {
	case WebhookConfigNewModeDisabled, WebhookConfigNewModeSvix, WebhookConfigNewModeDirect:
		return true
	}
	return false
}
