// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"

	"github.com/reductoai/reducto-go-sdk"
	"github.com/reductoai/reducto-go-sdk/internal/apijson"
	"github.com/reductoai/reducto-go-sdk/internal/param"
	"github.com/tidwall/gjson"
)

type AsyncEditResponse struct {
	JobID string                `json:"job_id" api:"required"`
	JSON  asyncEditResponseJSON `json:"-"`
}

// asyncEditResponseJSON contains the JSON metadata for the struct
// [AsyncEditResponse]
type asyncEditResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AsyncEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asyncEditResponseJSON) RawJSON() string {
	return r.raw
}

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

func (r AsyncExtractResponse) ImplementsExtractRunResponseUnion() {}

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

func (r AsyncParseResponse) ImplementsParseRunResponse() {}

type AsyncPipelineResponse struct {
	JobID string                    `json:"job_id" api:"required"`
	JSON  asyncPipelineResponseJSON `json:"-"`
}

// asyncPipelineResponseJSON contains the JSON metadata for the struct
// [AsyncPipelineResponse]
type asyncPipelineResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AsyncPipelineResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asyncPipelineResponseJSON) RawJSON() string {
	return r.raw
}

type AsyncSplitResponse struct {
	JobID string                 `json:"job_id" api:"required"`
	JSON  asyncSplitResponseJSON `json:"-"`
}

// asyncSplitResponseJSON contains the JSON metadata for the struct
// [AsyncSplitResponse]
type asyncSplitResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AsyncSplitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asyncSplitResponseJSON) RawJSON() string {
	return r.raw
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

// Response from classify job - returned when polling /job/{job_id}
type ClassifyResponse struct {
	JobID  string                 `json:"job_id" api:"required"`
	Result ClassifyResponseResult `json:"result" api:"required"`
	// The duration of the classify request in seconds.
	Duration float64 `json:"duration" api:"nullable"`
	// Overall confidence breakdown for classification response.
	ResponseConfidence ClassifyResponseResponseConfidence `json:"response_confidence" api:"nullable"`
	ResponseType       ClassifyResponseResponseType       `json:"response_type"`
	Usage              ClassifyResponseUsage              `json:"usage" api:"nullable"`
	JSON               classifyResponseJSON               `json:"-"`
}

// classifyResponseJSON contains the JSON metadata for the struct
// [ClassifyResponse]
type classifyResponseJSON struct {
	JobID              apijson.Field
	Result             apijson.Field
	Duration           apijson.Field
	ResponseConfidence apijson.Field
	ResponseType       apijson.Field
	Usage              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ClassifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r classifyResponseJSON) RawJSON() string {
	return r.raw
}

func (r ClassifyResponse) ImplementsJobGetResponseAsyncJobResponseResultUnion() {}

func (r ClassifyResponse) ImplementsJobGetResponseEnhancedAsyncJobResponseResultUnion() {}

type ClassifyResponseResult struct {
	Category string                     `json:"category" api:"required"`
	JSON     classifyResponseResultJSON `json:"-"`
}

// classifyResponseResultJSON contains the JSON metadata for the struct
// [ClassifyResponseResult]
type classifyResponseResultJSON struct {
	Category    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClassifyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r classifyResponseResultJSON) RawJSON() string {
	return r.raw
}

// Overall confidence breakdown for classification response.
type ClassifyResponseResponseConfidence struct {
	Categories []ClassifyResponseResponseConfidenceCategory `json:"categories" api:"required"`
	JSON       classifyResponseResponseConfidenceJSON       `json:"-"`
}

// classifyResponseResponseConfidenceJSON contains the JSON metadata for the struct
// [ClassifyResponseResponseConfidence]
type classifyResponseResponseConfidenceJSON struct {
	Categories  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClassifyResponseResponseConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r classifyResponseResponseConfidenceJSON) RawJSON() string {
	return r.raw
}

// Confidence result for a category.
type ClassifyResponseResponseConfidenceCategory struct {
	Category           string                                                           `json:"category" api:"required"`
	Confidence         float64                                                          `json:"confidence" api:"required"`
	CriteriaConfidence []ClassifyResponseResponseConfidenceCategoriesCriteriaConfidence `json:"criteria_confidence" api:"required"`
	JSON               classifyResponseResponseConfidenceCategoryJSON                   `json:"-"`
}

// classifyResponseResponseConfidenceCategoryJSON contains the JSON metadata for
// the struct [ClassifyResponseResponseConfidenceCategory]
type classifyResponseResponseConfidenceCategoryJSON struct {
	Category           apijson.Field
	Confidence         apijson.Field
	CriteriaConfidence apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ClassifyResponseResponseConfidenceCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r classifyResponseResponseConfidenceCategoryJSON) RawJSON() string {
	return r.raw
}

// Confidence result for a single criterion.
type ClassifyResponseResponseConfidenceCategoriesCriteriaConfidence struct {
	Confidence ClassifyResponseResponseConfidenceCategoriesCriteriaConfidenceConfidence `json:"confidence" api:"required"`
	Criterion  string                                                                   `json:"criterion" api:"required"`
	JSON       classifyResponseResponseConfidenceCategoriesCriteriaConfidenceJSON       `json:"-"`
}

// classifyResponseResponseConfidenceCategoriesCriteriaConfidenceJSON contains the
// JSON metadata for the struct
// [ClassifyResponseResponseConfidenceCategoriesCriteriaConfidence]
type classifyResponseResponseConfidenceCategoriesCriteriaConfidenceJSON struct {
	Confidence  apijson.Field
	Criterion   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClassifyResponseResponseConfidenceCategoriesCriteriaConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r classifyResponseResponseConfidenceCategoriesCriteriaConfidenceJSON) RawJSON() string {
	return r.raw
}

type ClassifyResponseResponseConfidenceCategoriesCriteriaConfidenceConfidence string

const (
	ClassifyResponseResponseConfidenceCategoriesCriteriaConfidenceConfidenceHigh ClassifyResponseResponseConfidenceCategoriesCriteriaConfidenceConfidence = "high"
	ClassifyResponseResponseConfidenceCategoriesCriteriaConfidenceConfidenceLow  ClassifyResponseResponseConfidenceCategoriesCriteriaConfidenceConfidence = "low"
)

func (r ClassifyResponseResponseConfidenceCategoriesCriteriaConfidenceConfidence) IsKnown() bool {
	switch r {
	case ClassifyResponseResponseConfidenceCategoriesCriteriaConfidenceConfidenceHigh, ClassifyResponseResponseConfidenceCategoriesCriteriaConfidenceConfidenceLow:
		return true
	}
	return false
}

type ClassifyResponseResponseType string

const (
	ClassifyResponseResponseTypeClassify ClassifyResponseResponseType = "classify"
)

func (r ClassifyResponseResponseType) IsKnown() bool {
	switch r {
	case ClassifyResponseResponseTypeClassify:
		return true
	}
	return false
}

type ClassifyResponseUsage struct {
	NumCategories int64                     `json:"num_categories" api:"required"`
	NumPages      int64                     `json:"num_pages" api:"required"`
	Credits       float64                   `json:"credits" api:"nullable"`
	JSON          classifyResponseUsageJSON `json:"-"`
}

// classifyResponseUsageJSON contains the JSON metadata for the struct
// [ClassifyResponseUsage]
type classifyResponseUsageJSON struct {
	NumCategories apijson.Field
	NumPages      apijson.Field
	Credits       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ClassifyResponseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r classifyResponseUsageJSON) RawJSON() string {
	return r.raw
}

type DirectWebhookConfigParam struct {
	URL  param.Field[string]                  `json:"url" api:"required"`
	Mode param.Field[DirectWebhookConfigMode] `json:"mode"`
}

func (r DirectWebhookConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectWebhookConfigParam) ImplementsAsyncConfigV3WebhookUnionParam() {}

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
	FormSchema   []reducto.EditWidget     `json:"form_schema" api:"nullable"`
	ResponseType EditResponseResponseType `json:"response_type"`
	// Usage information for the edit operation, including number of pages and credits
	// charged.
	Usage reducto.ParseUsage `json:"usage" api:"nullable"`
	JSON  editResponseJSON   `json:"-"`
}

// editResponseJSON contains the JSON metadata for the struct [EditResponse]
type editResponseJSON struct {
	DocumentURL  apijson.Field
	FormSchema   apijson.Field
	ResponseType apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editResponseJSON) RawJSON() string {
	return r.raw
}

func (r EditResponse) ImplementsJobGetResponseAsyncJobResponseResultUnion() {}

func (r EditResponse) ImplementsJobGetResponseEnhancedAsyncJobResponseResultUnion() {}

type EditResponseResponseType string

const (
	EditResponseResponseTypeEdit EditResponseResponseType = "edit"
)

func (r EditResponseResponseType) IsKnown() bool {
	switch r {
	case EditResponseResponseTypeEdit:
		return true
	}
	return false
}

type ExtractResponse map[string]interface{}

func (r ExtractResponse) ImplementsPipelineResponseResultExtractUnion() {}

func (r ExtractResponse) ImplementsPipelineResponseResultExtractArrayResultUnion() {}

func (r ExtractResponse) ImplementsJobGetResponseAsyncJobResponseResultUnion() {}

func (r ExtractResponse) ImplementsJobGetResponseEnhancedAsyncJobResponseResultUnion() {}

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

func (r PageRangeParam) ImplementsClassifyRunParamsPageRangeUnion() {}

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
	Usage  reducto.ParseUsage  `json:"usage" api:"required"`
	// Which pipeline produced this response. `lite` means Reducto Flash Lite served
	// the request; `base` is the standard pipeline. Optional / nullable for forward
	// compatibility — older API instances or persisted responses written before this
	// field existed will leave it `None`; treat `None` as `base`.
	ParseMode ParseResponseParseMode `json:"parse_mode" api:"nullable"`
	// The storage URL of the converted PDF file.
	PdfURL       string                    `json:"pdf_url" api:"nullable"`
	ResponseType ParseResponseResponseType `json:"response_type"`
	// The link to the studio pipeline for the document.
	StudioLink string            `json:"studio_link" api:"nullable"`
	JSON       parseResponseJSON `json:"-"`
}

// parseResponseJSON contains the JSON metadata for the struct [ParseResponse]
type parseResponseJSON struct {
	Duration     apijson.Field
	JobID        apijson.Field
	Result       apijson.Field
	Usage        apijson.Field
	ParseMode    apijson.Field
	PdfURL       apijson.Field
	ResponseType apijson.Field
	StudioLink   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ParseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseJSON) RawJSON() string {
	return r.raw
}

func (r ParseResponse) ImplementsPipelineResponseResultParseUnion() {}

func (r ParseResponse) ImplementsParseRunResponse() {}

func (r ParseResponse) ImplementsJobGetResponseAsyncJobResponseResultUnion() {}

func (r ParseResponse) ImplementsJobGetResponseEnhancedAsyncJobResponseResultUnion() {}

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
	Bbox reducto.BoundingBox `json:"bbox" api:"required"`
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
	Bbox reducto.BoundingBox `json:"bbox" api:"required"`
	Text string              `json:"text" api:"required"`
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
	Bbox reducto.BoundingBox `json:"bbox" api:"required"`
	Text string              `json:"text" api:"required"`
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

// Which pipeline produced this response. `lite` means Reducto Flash Lite served
// the request; `base` is the standard pipeline. Optional / nullable for forward
// compatibility — older API instances or persisted responses written before this
// field existed will leave it `None`; treat `None` as `base`.
type ParseResponseParseMode string

const (
	ParseResponseParseModeBase ParseResponseParseMode = "base"
	ParseResponseParseModeLite ParseResponseParseMode = "lite"
)

func (r ParseResponseParseMode) IsKnown() bool {
	switch r {
	case ParseResponseParseModeBase, ParseResponseParseModeLite:
		return true
	}
	return false
}

type ParseResponseResponseType string

const (
	ParseResponseResponseTypeParse ParseResponseResponseType = "parse"
)

func (r ParseResponseResponseType) IsKnown() bool {
	switch r {
	case ParseResponseResponseTypeParse:
		return true
	}
	return false
}

type PipelineResponse struct {
	JobID        string                       `json:"job_id" api:"required"`
	Result       PipelineResponseResult       `json:"result" api:"required"`
	Usage        reducto.ParseUsage           `json:"usage" api:"required"`
	ResponseType PipelineResponseResponseType `json:"response_type"`
	JSON         pipelineResponseJSON         `json:"-"`
}

// pipelineResponseJSON contains the JSON metadata for the struct
// [PipelineResponse]
type pipelineResponseJSON struct {
	JobID        apijson.Field
	Result       apijson.Field
	Usage        apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PipelineResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineResponseJSON) RawJSON() string {
	return r.raw
}

func (r PipelineResponse) ImplementsJobGetResponseAsyncJobResponseResultUnion() {}

func (r PipelineResponse) ImplementsJobGetResponseEnhancedAsyncJobResponseResultUnion() {}

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
// [reducto.V3Extract].
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
			Type:       reflect.TypeOf(reducto.V3Extract{}),
		},
	)
}

type PipelineResponseResultExtractArray []PipelineResponseResultExtractArrayItem

func (r PipelineResponseResultExtractArray) ImplementsPipelineResponseResultExtractUnion() {}

// This is the response format for Extract -> Split Pipelines
type PipelineResponseResultExtractArrayItem struct {
	PageRange []int64                                       `json:"page_range" api:"required"`
	Result    PipelineResponseResultExtractArrayResultUnion `json:"result" api:"required"`
	SplitName string                                        `json:"split_name" api:"required"`
	Partition string                                        `json:"partition" api:"nullable"`
	JSON      pipelineResponseResultExtractArrayItemJSON    `json:"-"`
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

// Union satisfied by [ExtractResponse] or [reducto.V3Extract].
type PipelineResponseResultExtractArrayResultUnion interface {
	ImplementsPipelineResponseResultExtractArrayResultUnion()
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
			Type:       reflect.TypeOf(reducto.V3Extract{}),
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

type PipelineResponseResponseType string

const (
	PipelineResponseResponseTypePipeline PipelineResponseResponseType = "pipeline"
)

func (r PipelineResponseResponseType) IsKnown() bool {
	switch r {
	case PipelineResponseResponseTypePipeline:
		return true
	}
	return false
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
	Result       SplitResponseResult       `json:"result" api:"required"`
	Usage        reducto.ParseUsage        `json:"usage" api:"required"`
	ResponseType SplitResponseResponseType `json:"response_type"`
	JSON         splitResponseJSON         `json:"-"`
}

// splitResponseJSON contains the JSON metadata for the struct [SplitResponse]
type splitResponseJSON struct {
	Result       apijson.Field
	Usage        apijson.Field
	ResponseType apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SplitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitResponseJSON) RawJSON() string {
	return r.raw
}

func (r SplitResponse) ImplementsJobGetResponseAsyncJobResponseResultUnion() {}

func (r SplitResponse) ImplementsJobGetResponseEnhancedAsyncJobResponseResultUnion() {}

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
	Pages      []reducto.DeepSplitPageEvidence                     `json:"pages" api:"required"`
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

type SplitResponseResultDeepSplitResultSplitsPartition struct {
	Name  string                                                `json:"name" api:"required"`
	Pages []reducto.DeepSplitPageEvidence                       `json:"pages" api:"required"`
	JSON  splitResponseResultDeepSplitResultSplitsPartitionJSON `json:"-"`
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

type SplitResponseResponseType string

const (
	SplitResponseResponseTypeSplit SplitResponseResponseType = "split"
)

func (r SplitResponseResponseType) IsKnown() bool {
	switch r {
	case SplitResponseResponseTypeSplit:
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

func (r SvixWebhookConfigParam) ImplementsAsyncConfigV3WebhookUnionParam() {}

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
	// Routing mode for table agentic: 'default' runs enrichment on all tables, 'auto'
	// uses the router to skip tables where enrichment is unlikely to help.
	Mode param.Field[TableAgenticMode] `json:"mode"`
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

// Routing mode for table agentic: 'default' runs enrichment on all tables, 'auto'
// uses the router to skip tables where enrichment is unlikely to help.
type TableAgenticMode string

const (
	TableAgenticModeDefault TableAgenticMode = "default"
	TableAgenticModeAuto    TableAgenticMode = "auto"
)

func (r TableAgenticMode) IsKnown() bool {
	switch r {
	case TableAgenticModeDefault, TableAgenticModeAuto:
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

func (r UploadParam) ImplementsAsyncParseConfigInputUnionParam() {}

func (r UploadParam) ImplementsParseRunParamsBodySyncParseConfigInputUnion() {}

func (r UploadParam) ImplementsAsyncExtractConfigInputUnionParam() {}

func (r UploadParam) ImplementsExtractRunParamsBodySyncExtractConfigInputUnion() {}

func (r UploadParam) ImplementsSplitRunParamsInputUnion() {}

func (r UploadParam) ImplementsSplitRunJobParamsInputUnion() {}

func (r UploadParam) ImplementsEditRunParamsDocumentURLUnion() {}

func (r UploadParam) ImplementsEditRunJobParamsDocumentURLUnion() {}

func (r UploadParam) ImplementsPipelineRunParamsInputUnion() {}

func (r UploadParam) ImplementsPipelineRunJobParamsInputUnion() {}

func (r UploadParam) ImplementsClassifyRunParamsInputUnion() {}

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
