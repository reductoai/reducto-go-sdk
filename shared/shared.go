// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"

	"github.com/reductoai/reducto-go-sdk/internal/apijson"
	"github.com/reductoai/reducto-go-sdk/internal/param"
	"github.com/tidwall/gjson"
)

type AdvancedProcessingOptionsParam struct {
	// If True, add page markers to the output (e.g. [[PAGE 1 BEGINS HERE]] and
	// [[PAGE 1 ENDS HERE]] added as blocks to the content). Defaults to False.
	AddPageMarkers param.Field[bool] `json:"add_page_markers"`
	// A flag to indicate if the hierarchy of the document should be continued from
	// chunk to chunk.
	ContinueHierarchy param.Field[bool] `json:"continue_hierarchy"`
	// Password to decrypt password-protected documents.
	DocumentPassword param.Field[string] `json:"document_password"`
	// Add <u> tags around underlined text, <s> tags around strikethrough text, and
	// <change> tags to surround both underlines and strikethroughs for change
	// detection. Defaults to False.
	EnableChangeTracking param.Field[bool] `json:"enable_change_tracking"`
	// Skip hidden rows and cols in Excel files. Defaults to False.
	ExcludeHiddenRowsCols param.Field[bool] `json:"exclude_hidden_rows_cols"`
	// Skip hidden sheets in Excel files. Defaults to False.
	ExcludeHiddenSheets param.Field[bool] `json:"exclude_hidden_sheets"`
	// If True, filter out line numbers from the output. Defaults to False.
	FilterLineNumbers param.Field[bool] `json:"filter_line_numbers"`
	// Force the URL to be downloaded as a specific file extension (e.g. .png).
	ForceFileExtension param.Field[string] `json:"force_file_extension"`
	// If line breaks should be preserved in the text.
	KeepLineBreaks param.Field[bool] `json:"keep_line_breaks"`
	// The configuration options for large table chunking (currently only supported on
	// spreadsheet and CSV files).
	LargeTableChunking param.Field[AdvancedProcessingOptionsLargeTableChunkingParam] `json:"large_table_chunking"`
	// A flag to indicate if consecutive tables with the same number of columns should
	// be merged across breaks and spaces.
	MergeTables param.Field[bool] `json:"merge_tables"`
	// The OCR system to use. Highres is recommended for documents with English
	// characters.
	OcrSystem param.Field[AdvancedProcessingOptionsOcrSystem] `json:"ocr_system"`
	// The page range to process (1-indexed). By default, the entire document is
	// processed.
	PageRange param.Field[AdvancedProcessingOptionsPageRangeUnionParam] `json:"page_range"`
	// If True, persist the results indefinitely. Defaults to False.
	PersistResults param.Field[bool] `json:"persist_results"`
	// If True, pull in PDF comments from the document. Defaults to False.
	ReadComments param.Field[bool] `json:"read_comments"`
	// If True, remove text formatting from the output (e.g. hyphens for list items).
	// Defaults to False.
	RemoveTextFormatting param.Field[bool] `json:"remove_text_formatting"`
	// If True, return OCR data in the result. Defaults to False.
	ReturnOcrData param.Field[bool] `json:"return_ocr_data"`
	// In a spreadsheet with different tables inside, we enable splitting up the tables
	// by default. Disabling will register as one large table.
	SpreadsheetTableClustering param.Field[AdvancedProcessingOptionsSpreadsheetTableClustering] `json:"spreadsheet_table_clustering"`
	// The mode to use for table output. Dynamic returns md for simpler tables and html
	// for more complex tables.
	TableOutputFormat param.Field[AdvancedProcessingOptionsTableOutputFormat] `json:"table_output_format"`
}

func (r AdvancedProcessingOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration options for large table chunking (currently only supported on
// spreadsheet and CSV files).
type AdvancedProcessingOptionsLargeTableChunkingParam struct {
	// If large tables should be chunked into smaller tables, currently only supported
	// on spreadsheet and CSV files.
	Enabled param.Field[bool] `json:"enabled"`
	// The max row/column size for a table to be chunked. Defaults to 50. Header
	// rows/columns are persisted based on heuristics.
	Size param.Field[int64] `json:"size"`
}

func (r AdvancedProcessingOptionsLargeTableChunkingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The OCR system to use. Highres is recommended for documents with English
// characters.
type AdvancedProcessingOptionsOcrSystem string

const (
	AdvancedProcessingOptionsOcrSystemHighres      AdvancedProcessingOptionsOcrSystem = "highres"
	AdvancedProcessingOptionsOcrSystemMultilingual AdvancedProcessingOptionsOcrSystem = "multilingual"
	AdvancedProcessingOptionsOcrSystemCombined     AdvancedProcessingOptionsOcrSystem = "combined"
)

func (r AdvancedProcessingOptionsOcrSystem) IsKnown() bool {
	switch r {
	case AdvancedProcessingOptionsOcrSystemHighres, AdvancedProcessingOptionsOcrSystemMultilingual, AdvancedProcessingOptionsOcrSystemCombined:
		return true
	}
	return false
}

// The page range to process (1-indexed). By default, the entire document is
// processed.
//
// Satisfied by [shared.PageRangeParam],
// [shared.AdvancedProcessingOptionsPageRangeArrayParam],
// [shared.AdvancedProcessingOptionsPageRangeArrayParam].
type AdvancedProcessingOptionsPageRangeUnionParam interface {
	ImplementsAdvancedProcessingOptionsPageRangeUnionParam()
}

type AdvancedProcessingOptionsPageRangeArrayParam []PageRangeParam

func (r AdvancedProcessingOptionsPageRangeArrayParam) ImplementsAdvancedProcessingOptionsPageRangeUnionParam() {
}

// In a spreadsheet with different tables inside, we enable splitting up the tables
// by default. Disabling will register as one large table.
type AdvancedProcessingOptionsSpreadsheetTableClustering string

const (
	AdvancedProcessingOptionsSpreadsheetTableClusteringDefault  AdvancedProcessingOptionsSpreadsheetTableClustering = "default"
	AdvancedProcessingOptionsSpreadsheetTableClusteringDisabled AdvancedProcessingOptionsSpreadsheetTableClustering = "disabled"
)

func (r AdvancedProcessingOptionsSpreadsheetTableClustering) IsKnown() bool {
	switch r {
	case AdvancedProcessingOptionsSpreadsheetTableClusteringDefault, AdvancedProcessingOptionsSpreadsheetTableClusteringDisabled:
		return true
	}
	return false
}

// The mode to use for table output. Dynamic returns md for simpler tables and html
// for more complex tables.
type AdvancedProcessingOptionsTableOutputFormat string

const (
	AdvancedProcessingOptionsTableOutputFormatHTML     AdvancedProcessingOptionsTableOutputFormat = "html"
	AdvancedProcessingOptionsTableOutputFormatJson     AdvancedProcessingOptionsTableOutputFormat = "json"
	AdvancedProcessingOptionsTableOutputFormatMd       AdvancedProcessingOptionsTableOutputFormat = "md"
	AdvancedProcessingOptionsTableOutputFormatJsonbbox AdvancedProcessingOptionsTableOutputFormat = "jsonbbox"
	AdvancedProcessingOptionsTableOutputFormatDynamic  AdvancedProcessingOptionsTableOutputFormat = "dynamic"
	AdvancedProcessingOptionsTableOutputFormatAIJson   AdvancedProcessingOptionsTableOutputFormat = "ai_json"
	AdvancedProcessingOptionsTableOutputFormatCsv      AdvancedProcessingOptionsTableOutputFormat = "csv"
)

func (r AdvancedProcessingOptionsTableOutputFormat) IsKnown() bool {
	switch r {
	case AdvancedProcessingOptionsTableOutputFormatHTML, AdvancedProcessingOptionsTableOutputFormatJson, AdvancedProcessingOptionsTableOutputFormatMd, AdvancedProcessingOptionsTableOutputFormatJsonbbox, AdvancedProcessingOptionsTableOutputFormatDynamic, AdvancedProcessingOptionsTableOutputFormatAIJson, AdvancedProcessingOptionsTableOutputFormatCsv:
		return true
	}
	return false
}

type ArrayExtractConfigParam struct {
	// Array extraction allows you to extract long lists of information from lengthy
	// documents. It makes parallel calls on overlapping sections of the document.
	Enabled param.Field[bool] `json:"enabled"`
	// The array extraction version to use.
	Mode param.Field[ArrayExtractConfigMode] `json:"mode"`
	// Length of each segment, in pages, for parallel calls with array extraction.
	PagesPerSegment param.Field[int64] `json:"pages_per_segment"`
	// Number of items to extract in each stream call. Lower numbers will increase
	// quality but be much slower. 50 works well for most documents with tables.
	StreamingExtractItemDensity param.Field[int64] `json:"streaming_extract_item_density"`
}

func (r ArrayExtractConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The array extraction version to use.
type ArrayExtractConfigMode string

const (
	ArrayExtractConfigModeAuto      ArrayExtractConfigMode = "auto"
	ArrayExtractConfigModeLegacy    ArrayExtractConfigMode = "legacy"
	ArrayExtractConfigModeStreaming ArrayExtractConfigMode = "streaming"
	ArrayExtractConfigModeNoOverlap ArrayExtractConfigMode = "no_overlap"
)

func (r ArrayExtractConfigMode) IsKnown() bool {
	switch r {
	case ArrayExtractConfigModeAuto, ArrayExtractConfigModeLegacy, ArrayExtractConfigModeStreaming, ArrayExtractConfigModeNoOverlap:
		return true
	}
	return false
}

type BaseProcessingOptionsParam struct {
	// The configuration options for chunking. Chunking is commonly used for RAG
	// usecases.
	Chunking param.Field[BaseProcessingOptionsChunkingParam] `json:"chunking"`
	// The mode to use for extraction. Metadata/hybrid are only recommended with high
	// quality metadata embeddings.
	ExtractionMode param.Field[BaseProcessingOptionsExtractionMode] `json:"extraction_mode"`
	// The configuration options for figure summarization.
	FigureSummary param.Field[BaseProcessingOptionsFigureSummaryParam] `json:"figure_summary"`
	// A list of block types to filter from chunk content. Pass blocks to filter them
	// from content. By default, no blocks are filtered.
	FilterBlocks param.Field[[]BaseProcessingOptionsFilterBlock] `json:"filter_blocks"`
	// Force the result to be returned in URL form (by default only used for very large
	// responses).
	ForceURLResult param.Field[bool] `json:"force_url_result"`
	// The mode to use for OCR. Agentic mode adds an extra pass, correcting any
	// table/text mistakes at a small cost.
	OcrMode param.Field[BaseProcessingOptionsOcrMode] `json:"ocr_mode"`
	// The configuration options for table summarization.
	TableSummary param.Field[BaseProcessingOptionsTableSummaryParam] `json:"table_summary"`
}

func (r BaseProcessingOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration options for chunking. Chunking is commonly used for RAG
// usecases.
type BaseProcessingOptionsChunkingParam struct {
	// Choose how to partition chunks. Variable mode chunks by character length and
	// visual context. Section mode chunks by section headers. Page mode chunks
	// according to pages. Page sections mode chunks first by page, then by sections
	// within each page. Disabled returns one single chunk.
	ChunkMode param.Field[BaseProcessingOptionsChunkingChunkMode] `json:"chunk_mode"`
	// The approximate size of chunks (in characters) that the document will be split
	// into. Defaults to None, in which case the chunk size is variable between 250 -
	// 1500 characters.
	ChunkSize param.Field[int64] `json:"chunk_size"`
}

func (r BaseProcessingOptionsChunkingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Choose how to partition chunks. Variable mode chunks by character length and
// visual context. Section mode chunks by section headers. Page mode chunks
// according to pages. Page sections mode chunks first by page, then by sections
// within each page. Disabled returns one single chunk.
type BaseProcessingOptionsChunkingChunkMode string

const (
	BaseProcessingOptionsChunkingChunkModeVariable     BaseProcessingOptionsChunkingChunkMode = "variable"
	BaseProcessingOptionsChunkingChunkModeSection      BaseProcessingOptionsChunkingChunkMode = "section"
	BaseProcessingOptionsChunkingChunkModePage         BaseProcessingOptionsChunkingChunkMode = "page"
	BaseProcessingOptionsChunkingChunkModeBlock        BaseProcessingOptionsChunkingChunkMode = "block"
	BaseProcessingOptionsChunkingChunkModeDisabled     BaseProcessingOptionsChunkingChunkMode = "disabled"
	BaseProcessingOptionsChunkingChunkModePageSections BaseProcessingOptionsChunkingChunkMode = "page_sections"
)

func (r BaseProcessingOptionsChunkingChunkMode) IsKnown() bool {
	switch r {
	case BaseProcessingOptionsChunkingChunkModeVariable, BaseProcessingOptionsChunkingChunkModeSection, BaseProcessingOptionsChunkingChunkModePage, BaseProcessingOptionsChunkingChunkModeBlock, BaseProcessingOptionsChunkingChunkModeDisabled, BaseProcessingOptionsChunkingChunkModePageSections:
		return true
	}
	return false
}

// The mode to use for extraction. Metadata/hybrid are only recommended with high
// quality metadata embeddings.
type BaseProcessingOptionsExtractionMode string

const (
	BaseProcessingOptionsExtractionModeOcr      BaseProcessingOptionsExtractionMode = "ocr"
	BaseProcessingOptionsExtractionModeMetadata BaseProcessingOptionsExtractionMode = "metadata"
	BaseProcessingOptionsExtractionModeHybrid   BaseProcessingOptionsExtractionMode = "hybrid"
)

func (r BaseProcessingOptionsExtractionMode) IsKnown() bool {
	switch r {
	case BaseProcessingOptionsExtractionModeOcr, BaseProcessingOptionsExtractionModeMetadata, BaseProcessingOptionsExtractionModeHybrid:
		return true
	}
	return false
}

// The configuration options for figure summarization.
type BaseProcessingOptionsFigureSummaryParam struct {
	// If figure summarization should be performed.
	Enabled param.Field[bool] `json:"enabled"`
	// If the figure summary prompt should override our default prompt.
	Override param.Field[bool] `json:"override"`
	// Add information to the prompt for figure summarization. Note any visual cues
	// that should be incorporated. Example: 'When provided a diagram, extract all of
	// the figure content verbatim.'
	Prompt param.Field[string] `json:"prompt"`
}

func (r BaseProcessingOptionsFigureSummaryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BaseProcessingOptionsFilterBlock string

const (
	BaseProcessingOptionsFilterBlockHeader        BaseProcessingOptionsFilterBlock = "Header"
	BaseProcessingOptionsFilterBlockFooter        BaseProcessingOptionsFilterBlock = "Footer"
	BaseProcessingOptionsFilterBlockTitle         BaseProcessingOptionsFilterBlock = "Title"
	BaseProcessingOptionsFilterBlockSectionHeader BaseProcessingOptionsFilterBlock = "Section Header"
	BaseProcessingOptionsFilterBlockPageNumber    BaseProcessingOptionsFilterBlock = "Page Number"
	BaseProcessingOptionsFilterBlockListItem      BaseProcessingOptionsFilterBlock = "List Item"
	BaseProcessingOptionsFilterBlockFigure        BaseProcessingOptionsFilterBlock = "Figure"
	BaseProcessingOptionsFilterBlockTable         BaseProcessingOptionsFilterBlock = "Table"
	BaseProcessingOptionsFilterBlockKeyValue      BaseProcessingOptionsFilterBlock = "Key Value"
	BaseProcessingOptionsFilterBlockText          BaseProcessingOptionsFilterBlock = "Text"
	BaseProcessingOptionsFilterBlockComment       BaseProcessingOptionsFilterBlock = "Comment"
)

func (r BaseProcessingOptionsFilterBlock) IsKnown() bool {
	switch r {
	case BaseProcessingOptionsFilterBlockHeader, BaseProcessingOptionsFilterBlockFooter, BaseProcessingOptionsFilterBlockTitle, BaseProcessingOptionsFilterBlockSectionHeader, BaseProcessingOptionsFilterBlockPageNumber, BaseProcessingOptionsFilterBlockListItem, BaseProcessingOptionsFilterBlockFigure, BaseProcessingOptionsFilterBlockTable, BaseProcessingOptionsFilterBlockKeyValue, BaseProcessingOptionsFilterBlockText, BaseProcessingOptionsFilterBlockComment:
		return true
	}
	return false
}

// The mode to use for OCR. Agentic mode adds an extra pass, correcting any
// table/text mistakes at a small cost.
type BaseProcessingOptionsOcrMode string

const (
	BaseProcessingOptionsOcrModeStandard BaseProcessingOptionsOcrMode = "standard"
	BaseProcessingOptionsOcrModeAgentic  BaseProcessingOptionsOcrMode = "agentic"
)

func (r BaseProcessingOptionsOcrMode) IsKnown() bool {
	switch r {
	case BaseProcessingOptionsOcrModeStandard, BaseProcessingOptionsOcrModeAgentic:
		return true
	}
	return false
}

// The configuration options for table summarization.
type BaseProcessingOptionsTableSummaryParam struct {
	// If table summarization should be performed.
	Enabled param.Field[bool] `json:"enabled"`
	// Add information to the prompt for table summarization.
	Prompt param.Field[string] `json:"prompt"`
}

func (r BaseProcessingOptionsTableSummaryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BoundingBox struct {
	Height float64 `json:"height,required"`
	Left   float64 `json:"left,required"`
	// The page number of the bounding box (1-indexed).
	Page  int64   `json:"page,required"`
	Top   float64 `json:"top,required"`
	Width float64 `json:"width,required"`
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

type ExperimentalProcessingOptionsParam struct {
	// You probably shouldn't use this. If True, filter out boxes with width greater
	// than 50% of the document width. Defaults to False. You probably don't want to
	// use this.
	DangerFilterWideBoxes param.Field[bool] `json:"danger_filter_wide_boxes"`
	// If extracted OCR text metadata should be embedded back into the returned PDF,
	// overwriting any existing text. Defaults to False.
	EmbedTextMetadataPdf param.Field[bool] `json:"embed_text_metadata_pdf"`
	// Use an experimental checkbox detection model to add checkboxes to the output,
	// defaults to False
	EnableCheckboxes param.Field[bool] `json:"enable_checkboxes"`
	// Use an experimental equation detection model to add equations to the output,
	// defaults to False
	EnableEquations param.Field[bool] `json:"enable_equations"`
	// Add <sub> tag around subscripts and <sup> tag around superscripts, defaults to
	// False
	EnableScripts param.Field[bool] `json:"enable_scripts"`
	// The configuration options for enrichment.
	Enrich param.Field[ExperimentalProcessingOptionsEnrichParam] `json:"enrich"`
	// The layout model to use for the document. This will be deprecated in the future.
	LayoutModel param.Field[ExperimentalProcessingOptionsLayoutModel] `json:"layout_model"`
	// Instead of using LibreOffice, when enabled, this flag uses a Windows VM to
	// convert files. This is slower but more accurate.
	NativeOfficeConversion param.Field[bool] `json:"native_office_conversion"`
	// If figure images should be returned in the result. Defaults to False.
	ReturnFigureImages param.Field[bool] `json:"return_figure_images"`
	// If table images should be returned in the result. Defaults to False.
	ReturnTableImages param.Field[bool] `json:"return_table_images"`
	// Use an orientation model to detect and rotate figures as needed, defaults to
	// False
	RotateFigures param.Field[bool] `json:"rotate_figures"`
	// Use an orientation model to detect and rotate pages as needed, defaults to True
	RotatePages param.Field[bool]      `json:"rotate_pages"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r ExperimentalProcessingOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration options for enrichment.
type ExperimentalProcessingOptionsEnrichParam struct {
	// If enabled, a large language/vision model will be used to postprocess the
	// extracted content. Note: enabling enrich requires tables be outputted in
	// markdown format. Defaults to False.
	Enabled param.Field[bool] `json:"enabled"`
	// The mode to use for enrichment. Defaults to standard
	Mode param.Field[ExperimentalProcessingOptionsEnrichMode] `json:"mode"`
	// Add information to the prompt for enrichment.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ExperimentalProcessingOptionsEnrichParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The mode to use for enrichment. Defaults to standard
type ExperimentalProcessingOptionsEnrichMode string

const (
	ExperimentalProcessingOptionsEnrichModeStandard ExperimentalProcessingOptionsEnrichMode = "standard"
	ExperimentalProcessingOptionsEnrichModePage     ExperimentalProcessingOptionsEnrichMode = "page"
	ExperimentalProcessingOptionsEnrichModeTable    ExperimentalProcessingOptionsEnrichMode = "table"
)

func (r ExperimentalProcessingOptionsEnrichMode) IsKnown() bool {
	switch r {
	case ExperimentalProcessingOptionsEnrichModeStandard, ExperimentalProcessingOptionsEnrichModePage, ExperimentalProcessingOptionsEnrichModeTable:
		return true
	}
	return false
}

// The layout model to use for the document. This will be deprecated in the future.
type ExperimentalProcessingOptionsLayoutModel string

const (
	ExperimentalProcessingOptionsLayoutModelDefault ExperimentalProcessingOptionsLayoutModel = "default"
	ExperimentalProcessingOptionsLayoutModelBeta    ExperimentalProcessingOptionsLayoutModel = "beta"
)

func (r ExperimentalProcessingOptionsLayoutModel) IsKnown() bool {
	switch r {
	case ExperimentalProcessingOptionsLayoutModelDefault, ExperimentalProcessingOptionsLayoutModelBeta:
		return true
	}
	return false
}

type ExtractResponse struct {
	// The citations corresponding to the extracted response.
	Citations []interface{} `json:"citations,required,nullable"`
	// The extracted response in your provided schema. This is a list of dictionaries.
	// If disable_chunking is True (default), then it will be a list of length one.
	Result []interface{}        `json:"result,required"`
	Usage  ExtractResponseUsage `json:"usage,required"`
	JSON   extractResponseJSON  `json:"-"`
}

// extractResponseJSON contains the JSON metadata for the struct [ExtractResponse]
type extractResponseJSON struct {
	Citations   apijson.Field
	Result      apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtractResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extractResponseJSON) RawJSON() string {
	return r.raw
}

func (r ExtractResponse) ImplementsJobGetResponseResult() {}

type ExtractResponseUsage struct {
	NumFields int64                    `json:"num_fields,required"`
	NumPages  int64                    `json:"num_pages,required"`
	Credits   float64                  `json:"credits,nullable"`
	JSON      extractResponseUsageJSON `json:"-"`
}

// extractResponseUsageJSON contains the JSON metadata for the struct
// [ExtractResponseUsage]
type extractResponseUsageJSON struct {
	NumFields   apijson.Field
	NumPages    apijson.Field
	Credits     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtractResponseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extractResponseUsageJSON) RawJSON() string {
	return r.raw
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

func (r PageRangeParam) ImplementsAdvancedProcessingOptionsPageRangeUnionParam() {}

type ParseResponse struct {
	// The duration of the parse request in seconds.
	Duration float64 `json:"duration,required"`
	JobID    string  `json:"job_id,required"`
	// The response from the document processing service. Note that there can be two
	// types of responses, Full Result and URL Result. This is due to limitations on
	// the max return size on HTTPS. If the response is too large, it will be returned
	// as a presigned URL in the URL response. You should handle this in your
	// application.
	Result ParseResponseResult `json:"result,required"`
	Usage  ParseUsage          `json:"usage,required"`
	// The storage URL of the converted PDF file.
	PdfURL string            `json:"pdf_url,nullable"`
	JSON   parseResponseJSON `json:"-"`
}

// parseResponseJSON contains the JSON metadata for the struct [ParseResponse]
type parseResponseJSON struct {
	Duration    apijson.Field
	JobID       apijson.Field
	Result      apijson.Field
	Usage       apijson.Field
	PdfURL      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseJSON) RawJSON() string {
	return r.raw
}

func (r ParseResponse) ImplementsJobGetResponseResult() {}

// The response from the document processing service. Note that there can be two
// types of responses, Full Result and URL Result. This is due to limitations on
// the max return size on HTTPS. If the response is too large, it will be returned
// as a presigned URL in the URL response. You should handle this in your
// application.
type ParseResponseResult struct {
	// type = 'full'
	Type ParseResponseResultType `json:"type,required"`
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
	Chunks []ParseResponseResultFullResultChunk `json:"chunks,required"`
	// type = 'full'
	Type   ParseResponseResultFullResultType `json:"type,required"`
	Custom interface{}                       `json:"custom"`
	Ocr    ParseResponseResultFullResultOcr  `json:"ocr,nullable"`
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
	Blocks []ParseResponseResultFullResultChunksBlock `json:"blocks,required"`
	// The content of the chunk extracted from the document.
	Content string `json:"content,required"`
	// Chunk content optimized for embedding and retrieval.
	Embed string `json:"embed,required"`
	// The enriched content of the chunk extracted from the document.
	Enriched string `json:"enriched,required,nullable"`
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
	Bbox BoundingBox `json:"bbox,required"`
	// The content of the block extracted from the document.
	Content string `json:"content,required"`
	// The type of block extracted from the document.
	Type ParseResponseResultFullResultChunksBlocksType `json:"type,required"`
	// The confidence for the block. It is either low or high and takes into account
	// factors like OCR and table structure
	Confidence string `json:"confidence,nullable"`
	// (Experimental) The URL of the image associated with the block.
	ImageURL string                                       `json:"image_url,nullable"`
	JSON     parseResponseResultFullResultChunksBlockJSON `json:"-"`
}

// parseResponseResultFullResultChunksBlockJSON contains the JSON metadata for the
// struct [ParseResponseResultFullResultChunksBlock]
type parseResponseResultFullResultChunksBlockJSON struct {
	Bbox        apijson.Field
	Content     apijson.Field
	Type        apijson.Field
	Confidence  apijson.Field
	ImageURL    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
)

func (r ParseResponseResultFullResultChunksBlocksType) IsKnown() bool {
	switch r {
	case ParseResponseResultFullResultChunksBlocksTypeHeader, ParseResponseResultFullResultChunksBlocksTypeFooter, ParseResponseResultFullResultChunksBlocksTypeTitle, ParseResponseResultFullResultChunksBlocksTypeSectionHeader, ParseResponseResultFullResultChunksBlocksTypePageNumber, ParseResponseResultFullResultChunksBlocksTypeListItem, ParseResponseResultFullResultChunksBlocksTypeFigure, ParseResponseResultFullResultChunksBlocksTypeTable, ParseResponseResultFullResultChunksBlocksTypeKeyValue, ParseResponseResultFullResultChunksBlocksTypeText, ParseResponseResultFullResultChunksBlocksTypeComment:
		return true
	}
	return false
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
	Lines []ParseResponseResultFullResultOcrLine `json:"lines,required"`
	Words []ParseResponseResultFullResultOcrWord `json:"words,required"`
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
	Bbox BoundingBox                              `json:"bbox,required"`
	Text string                                   `json:"text,required"`
	JSON parseResponseResultFullResultOcrLineJSON `json:"-"`
}

// parseResponseResultFullResultOcrLineJSON contains the JSON metadata for the
// struct [ParseResponseResultFullResultOcrLine]
type parseResponseResultFullResultOcrLineJSON struct {
	Bbox        apijson.Field
	Text        apijson.Field
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
	Bbox BoundingBox                              `json:"bbox,required"`
	Text string                                   `json:"text,required"`
	JSON parseResponseResultFullResultOcrWordJSON `json:"-"`
}

// parseResponseResultFullResultOcrWordJSON contains the JSON metadata for the
// struct [ParseResponseResultFullResultOcrWord]
type parseResponseResultFullResultOcrWordJSON struct {
	Bbox        apijson.Field
	Text        apijson.Field
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
	ResultID string `json:"result_id,required"`
	// type = 'url'
	Type ParseResponseResultURLResultType `json:"type,required"`
	URL  string                           `json:"url,required"`
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
	NumPages int64          `json:"num_pages,required"`
	Credits  float64        `json:"credits,nullable"`
	JSON     parseUsageJSON `json:"-"`
}

// parseUsageJSON contains the JSON metadata for the struct [ParseUsage]
type parseUsageJSON struct {
	NumPages    apijson.Field
	Credits     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseUsageJSON) RawJSON() string {
	return r.raw
}

type SplitCategoryParam struct {
	Description  param.Field[string] `json:"description,required"`
	Name         param.Field[string] `json:"name,required"`
	PartitionKey param.Field[string] `json:"partition_key"`
}

func (r SplitCategoryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SplitResponse struct {
	// The split result.
	Result SplitResponseResult `json:"result,required"`
	Usage  ParseUsage          `json:"usage,required"`
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

func (r SplitResponse) ImplementsJobGetResponseResult() {}

// The split result.
type SplitResponseResult struct {
	SectionMapping map[string][]int64      `json:"section_mapping,required,nullable"`
	Splits         []interface{}           `json:"splits,required"`
	JSON           splitResponseResultJSON `json:"-"`
}

// splitResponseResultJSON contains the JSON metadata for the struct
// [SplitResponseResult]
type splitResponseResultJSON struct {
	SectionMapping apijson.Field
	Splits         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SplitResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitResponseResultJSON) RawJSON() string {
	return r.raw
}

type Upload struct {
	FileID       string     `json:"file_id,required"`
	PresignedURL string     `json:"presigned_url,nullable"`
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
	FileID       param.Field[string] `json:"file_id,required"`
	PresignedURL param.Field[string] `json:"presigned_url"`
}

func (r UploadParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UploadParam) ImplementsSplitRunParamsDocumentURLUnion() {}

func (r UploadParam) ImplementsSplitRunJobParamsDocumentURLUnion() {}

func (r UploadParam) ImplementsParseConfigDocumentURLUnionParam() {}

func (r UploadParam) ImplementsParseRunJobParamsDocumentURLUnion() {}

func (r UploadParam) ImplementsExtractConfigDocumentURLUnionParam() {}

func (r UploadParam) ImplementsExtractRunJobParamsDocumentURLUnion() {}

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
