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

// SplitService contains methods and other services that help with interacting with
// the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSplitService] method instead.
type SplitService struct {
	Options []option.RequestOption
}

// NewSplitService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSplitService(opts ...option.RequestOption) (r *SplitService) {
	r = &SplitService{}
	r.Options = opts
	return
}

// Split
func (r *SplitService) Run(ctx context.Context, body SplitRunParams, opts ...option.RequestOption) (res *shared.SplitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "split"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Split Async
func (r *SplitService) RunJob(ctx context.Context, body SplitRunJobParams, opts ...option.RequestOption) (res *SplitRunJobResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "split_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SplitRunJobResponse struct {
	JobID string                  `json:"job_id,required"`
	JSON  splitRunJobResponseJSON `json:"-"`
}

// splitRunJobResponseJSON contains the JSON metadata for the struct
// [SplitRunJobResponse]
type splitRunJobResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitRunJobResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitRunJobResponseJSON) RawJSON() string {
	return r.raw
}

type SplitRunParams struct {
	// The URL of the document to be processed. You can provide one of the
	// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
	// prefixed URL obtained from the /upload endpoint after directly uploading a
	// document 4. A jobid:// prefixed URL obtained from a previous /parse invocation
	Input param.Field[SplitRunParamsInputUnion] `json:"input,required"`
	// The configuration options for processing the document.
	SplitDescription param.Field[[]shared.SplitCategoryParam] `json:"split_description,required"`
	// The configuration options for parsing the document. If you are passing in a
	// jobid:// URL for the file, then this configuration will be ignored.
	Parsing param.Field[SplitRunParamsParsing] `json:"parsing"`
	// The settings for split processing.
	Settings param.Field[SplitRunParamsSettings] `json:"settings"`
	// The prompt that describes rules for splitting the document.
	SplitRules param.Field[string] `json:"split_rules"`
}

func (r SplitRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The URL of the document to be processed. You can provide one of the
// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
// prefixed URL obtained from the /upload endpoint after directly uploading a
// document 4. A jobid:// prefixed URL obtained from a previous /parse invocation
//
// Satisfied by [shared.UnionString], [shared.UploadParam].
type SplitRunParamsInputUnion interface {
	ImplementsSplitRunParamsInputUnion()
}

// The configuration options for parsing the document. If you are passing in a
// jobid:// URL for the file, then this configuration will be ignored.
type SplitRunParamsParsing struct {
	Enhance     param.Field[SplitRunParamsParsingEnhance]     `json:"enhance"`
	Formatting  param.Field[SplitRunParamsParsingFormatting]  `json:"formatting"`
	Retrieval   param.Field[SplitRunParamsParsingRetrieval]   `json:"retrieval"`
	Settings    param.Field[SplitRunParamsParsingSettings]    `json:"settings"`
	Spreadsheet param.Field[SplitRunParamsParsingSpreadsheet] `json:"spreadsheet"`
}

func (r SplitRunParamsParsing) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SplitRunParamsParsingEnhance struct {
	// Agentic uses vision language models to enhance the accuracy of the output of
	// different types of extraction. This will incur a cost and latency increase.
	Agentic param.Field[[]SplitRunParamsParsingEnhanceAgenticUnion] `json:"agentic"`
	// If True, summarize figures using a small vision language model. Defaults to
	// True.
	SummarizeFigures param.Field[bool] `json:"summarize_figures"`
}

func (r SplitRunParamsParsingEnhance) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SplitRunParamsParsingEnhanceAgentic struct {
	Scope param.Field[SplitRunParamsParsingEnhanceAgenticScope] `json:"scope,required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r SplitRunParamsParsingEnhanceAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SplitRunParamsParsingEnhanceAgentic) implementsSplitRunParamsParsingEnhanceAgenticUnion() {}

// Satisfied by [SplitRunParamsParsingEnhanceAgenticTableAgentic],
// [SplitRunParamsParsingEnhanceAgenticFigureAgentic],
// [SplitRunParamsParsingEnhanceAgenticTextAgentic],
// [SplitRunParamsParsingEnhanceAgentic].
type SplitRunParamsParsingEnhanceAgenticUnion interface {
	implementsSplitRunParamsParsingEnhanceAgenticUnion()
}

type SplitRunParamsParsingEnhanceAgenticTableAgentic struct {
	Scope param.Field[SplitRunParamsParsingEnhanceAgenticTableAgenticScope] `json:"scope,required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r SplitRunParamsParsingEnhanceAgenticTableAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SplitRunParamsParsingEnhanceAgenticTableAgentic) implementsSplitRunParamsParsingEnhanceAgenticUnion() {
}

type SplitRunParamsParsingEnhanceAgenticTableAgenticScope string

const (
	SplitRunParamsParsingEnhanceAgenticTableAgenticScopeTable SplitRunParamsParsingEnhanceAgenticTableAgenticScope = "table"
)

func (r SplitRunParamsParsingEnhanceAgenticTableAgenticScope) IsKnown() bool {
	switch r {
	case SplitRunParamsParsingEnhanceAgenticTableAgenticScopeTable:
		return true
	}
	return false
}

type SplitRunParamsParsingEnhanceAgenticFigureAgentic struct {
	Scope param.Field[SplitRunParamsParsingEnhanceAgenticFigureAgenticScope] `json:"scope,required"`
	// Custom prompt for figure agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r SplitRunParamsParsingEnhanceAgenticFigureAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SplitRunParamsParsingEnhanceAgenticFigureAgentic) implementsSplitRunParamsParsingEnhanceAgenticUnion() {
}

type SplitRunParamsParsingEnhanceAgenticFigureAgenticScope string

const (
	SplitRunParamsParsingEnhanceAgenticFigureAgenticScopeFigure SplitRunParamsParsingEnhanceAgenticFigureAgenticScope = "figure"
)

func (r SplitRunParamsParsingEnhanceAgenticFigureAgenticScope) IsKnown() bool {
	switch r {
	case SplitRunParamsParsingEnhanceAgenticFigureAgenticScopeFigure:
		return true
	}
	return false
}

type SplitRunParamsParsingEnhanceAgenticTextAgentic struct {
	Scope param.Field[SplitRunParamsParsingEnhanceAgenticTextAgenticScope] `json:"scope,required"`
}

func (r SplitRunParamsParsingEnhanceAgenticTextAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SplitRunParamsParsingEnhanceAgenticTextAgentic) implementsSplitRunParamsParsingEnhanceAgenticUnion() {
}

type SplitRunParamsParsingEnhanceAgenticTextAgenticScope string

const (
	SplitRunParamsParsingEnhanceAgenticTextAgenticScopeText SplitRunParamsParsingEnhanceAgenticTextAgenticScope = "text"
)

func (r SplitRunParamsParsingEnhanceAgenticTextAgenticScope) IsKnown() bool {
	switch r {
	case SplitRunParamsParsingEnhanceAgenticTextAgenticScopeText:
		return true
	}
	return false
}

type SplitRunParamsParsingEnhanceAgenticScope string

const (
	SplitRunParamsParsingEnhanceAgenticScopeTable  SplitRunParamsParsingEnhanceAgenticScope = "table"
	SplitRunParamsParsingEnhanceAgenticScopeFigure SplitRunParamsParsingEnhanceAgenticScope = "figure"
	SplitRunParamsParsingEnhanceAgenticScopeText   SplitRunParamsParsingEnhanceAgenticScope = "text"
)

func (r SplitRunParamsParsingEnhanceAgenticScope) IsKnown() bool {
	switch r {
	case SplitRunParamsParsingEnhanceAgenticScopeTable, SplitRunParamsParsingEnhanceAgenticScopeFigure, SplitRunParamsParsingEnhanceAgenticScopeText:
		return true
	}
	return false
}

type SplitRunParamsParsingFormatting struct {
	// If True, add page markers to the output. Defaults to False. Useful for
	// extracting data with page specific information.
	AddPageMarkers param.Field[bool] `json:"add_page_markers"`
	// A list of formatting to include in the output. [insert description of each
	// option here later]
	Include param.Field[[]SplitRunParamsParsingFormattingInclude] `json:"include"`
	// A flag to indicate if consecutive tables with the same number of columns should
	// be merged. Defaults to False.
	MergeTables param.Field[bool] `json:"merge_tables"`
	// The mode to use for table output. Defaults to dynamic, which returns md for
	// simpler tables and html for more complex tables.
	TableOutputFormat param.Field[SplitRunParamsParsingFormattingTableOutputFormat] `json:"table_output_format"`
}

func (r SplitRunParamsParsingFormatting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SplitRunParamsParsingFormattingInclude string

const (
	SplitRunParamsParsingFormattingIncludeChangeTracking SplitRunParamsParsingFormattingInclude = "change_tracking"
	SplitRunParamsParsingFormattingIncludeHighlight      SplitRunParamsParsingFormattingInclude = "highlight"
	SplitRunParamsParsingFormattingIncludeComments       SplitRunParamsParsingFormattingInclude = "comments"
)

func (r SplitRunParamsParsingFormattingInclude) IsKnown() bool {
	switch r {
	case SplitRunParamsParsingFormattingIncludeChangeTracking, SplitRunParamsParsingFormattingIncludeHighlight, SplitRunParamsParsingFormattingIncludeComments:
		return true
	}
	return false
}

// The mode to use for table output. Defaults to dynamic, which returns md for
// simpler tables and html for more complex tables.
type SplitRunParamsParsingFormattingTableOutputFormat string

const (
	SplitRunParamsParsingFormattingTableOutputFormatHTML     SplitRunParamsParsingFormattingTableOutputFormat = "html"
	SplitRunParamsParsingFormattingTableOutputFormatJson     SplitRunParamsParsingFormattingTableOutputFormat = "json"
	SplitRunParamsParsingFormattingTableOutputFormatMd       SplitRunParamsParsingFormattingTableOutputFormat = "md"
	SplitRunParamsParsingFormattingTableOutputFormatJsonbbox SplitRunParamsParsingFormattingTableOutputFormat = "jsonbbox"
	SplitRunParamsParsingFormattingTableOutputFormatDynamic  SplitRunParamsParsingFormattingTableOutputFormat = "dynamic"
	SplitRunParamsParsingFormattingTableOutputFormatCsv      SplitRunParamsParsingFormattingTableOutputFormat = "csv"
)

func (r SplitRunParamsParsingFormattingTableOutputFormat) IsKnown() bool {
	switch r {
	case SplitRunParamsParsingFormattingTableOutputFormatHTML, SplitRunParamsParsingFormattingTableOutputFormatJson, SplitRunParamsParsingFormattingTableOutputFormatMd, SplitRunParamsParsingFormattingTableOutputFormatJsonbbox, SplitRunParamsParsingFormattingTableOutputFormatDynamic, SplitRunParamsParsingFormattingTableOutputFormatCsv:
		return true
	}
	return false
}

type SplitRunParamsParsingRetrieval struct {
	Chunking param.Field[SplitRunParamsParsingRetrievalChunking] `json:"chunking"`
	// If True, use embedding optimized mode. Defaults to False.
	EmbeddingOptimized param.Field[bool] `json:"embedding_optimized"`
	// A list of block types to filter out from 'content' and 'embed' fields. By
	// default, no blocks are filtered.
	FilterBlocks param.Field[[]SplitRunParamsParsingRetrievalFilterBlock] `json:"filter_blocks"`
}

func (r SplitRunParamsParsingRetrieval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SplitRunParamsParsingRetrievalChunking struct {
	// Choose how to partition chunks. Variable mode chunks by character length and
	// visual context. Section mode chunks by section headers. Page mode chunks
	// according to pages. Page sections mode chunks first by page, then by sections
	// within each page. Disabled returns one single chunk.
	ChunkMode param.Field[SplitRunParamsParsingRetrievalChunkingChunkMode] `json:"chunk_mode"`
	// The approximate size of chunks (in characters) that the document will be split
	// into. Defaults to null, in which case the chunk size is variable between 250 -
	// 1500 characters.
	ChunkSize param.Field[int64] `json:"chunk_size"`
}

func (r SplitRunParamsParsingRetrievalChunking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Choose how to partition chunks. Variable mode chunks by character length and
// visual context. Section mode chunks by section headers. Page mode chunks
// according to pages. Page sections mode chunks first by page, then by sections
// within each page. Disabled returns one single chunk.
type SplitRunParamsParsingRetrievalChunkingChunkMode string

const (
	SplitRunParamsParsingRetrievalChunkingChunkModeVariable     SplitRunParamsParsingRetrievalChunkingChunkMode = "variable"
	SplitRunParamsParsingRetrievalChunkingChunkModeSection      SplitRunParamsParsingRetrievalChunkingChunkMode = "section"
	SplitRunParamsParsingRetrievalChunkingChunkModePage         SplitRunParamsParsingRetrievalChunkingChunkMode = "page"
	SplitRunParamsParsingRetrievalChunkingChunkModeDisabled     SplitRunParamsParsingRetrievalChunkingChunkMode = "disabled"
	SplitRunParamsParsingRetrievalChunkingChunkModeBlock        SplitRunParamsParsingRetrievalChunkingChunkMode = "block"
	SplitRunParamsParsingRetrievalChunkingChunkModePageSections SplitRunParamsParsingRetrievalChunkingChunkMode = "page_sections"
)

func (r SplitRunParamsParsingRetrievalChunkingChunkMode) IsKnown() bool {
	switch r {
	case SplitRunParamsParsingRetrievalChunkingChunkModeVariable, SplitRunParamsParsingRetrievalChunkingChunkModeSection, SplitRunParamsParsingRetrievalChunkingChunkModePage, SplitRunParamsParsingRetrievalChunkingChunkModeDisabled, SplitRunParamsParsingRetrievalChunkingChunkModeBlock, SplitRunParamsParsingRetrievalChunkingChunkModePageSections:
		return true
	}
	return false
}

type SplitRunParamsParsingRetrievalFilterBlock string

const (
	SplitRunParamsParsingRetrievalFilterBlockHeader        SplitRunParamsParsingRetrievalFilterBlock = "Header"
	SplitRunParamsParsingRetrievalFilterBlockFooter        SplitRunParamsParsingRetrievalFilterBlock = "Footer"
	SplitRunParamsParsingRetrievalFilterBlockTitle         SplitRunParamsParsingRetrievalFilterBlock = "Title"
	SplitRunParamsParsingRetrievalFilterBlockSectionHeader SplitRunParamsParsingRetrievalFilterBlock = "Section Header"
	SplitRunParamsParsingRetrievalFilterBlockPageNumber    SplitRunParamsParsingRetrievalFilterBlock = "Page Number"
	SplitRunParamsParsingRetrievalFilterBlockListItem      SplitRunParamsParsingRetrievalFilterBlock = "List Item"
	SplitRunParamsParsingRetrievalFilterBlockFigure        SplitRunParamsParsingRetrievalFilterBlock = "Figure"
	SplitRunParamsParsingRetrievalFilterBlockTable         SplitRunParamsParsingRetrievalFilterBlock = "Table"
	SplitRunParamsParsingRetrievalFilterBlockKeyValue      SplitRunParamsParsingRetrievalFilterBlock = "Key Value"
	SplitRunParamsParsingRetrievalFilterBlockText          SplitRunParamsParsingRetrievalFilterBlock = "Text"
	SplitRunParamsParsingRetrievalFilterBlockComment       SplitRunParamsParsingRetrievalFilterBlock = "Comment"
	SplitRunParamsParsingRetrievalFilterBlockSignature     SplitRunParamsParsingRetrievalFilterBlock = "Signature"
)

func (r SplitRunParamsParsingRetrievalFilterBlock) IsKnown() bool {
	switch r {
	case SplitRunParamsParsingRetrievalFilterBlockHeader, SplitRunParamsParsingRetrievalFilterBlockFooter, SplitRunParamsParsingRetrievalFilterBlockTitle, SplitRunParamsParsingRetrievalFilterBlockSectionHeader, SplitRunParamsParsingRetrievalFilterBlockPageNumber, SplitRunParamsParsingRetrievalFilterBlockListItem, SplitRunParamsParsingRetrievalFilterBlockFigure, SplitRunParamsParsingRetrievalFilterBlockTable, SplitRunParamsParsingRetrievalFilterBlockKeyValue, SplitRunParamsParsingRetrievalFilterBlockText, SplitRunParamsParsingRetrievalFilterBlockComment, SplitRunParamsParsingRetrievalFilterBlockSignature:
		return true
	}
	return false
}

type SplitRunParamsParsingSettings struct {
	// Password to decrypt password-protected documents.
	DocumentPassword param.Field[string] `json:"document_password"`
	// If True, embed OCR metadata into the returned PDF. Defaults to False.
	EmbedPdfMetadata param.Field[bool] `json:"embed_pdf_metadata"`
	// Force the URL to be downloaded as a specific file extension (e.g. `.png`).
	ForceFileExtension param.Field[string] `json:"force_file_extension"`
	// Force the result to be returned in URL form.
	ForceURLResult param.Field[bool] `json:"force_url_result"`
	// Standard is our best multilingual OCR system. Legacy only supports germanic
	// languages and is available for backwards compatibility.
	OcrSystem param.Field[SplitRunParamsParsingSettingsOcrSystem] `json:"ocr_system"`
	// The page range to process (1-indexed). By default, the entire document is
	// processed.
	PageRange param.Field[SplitRunParamsParsingSettingsPageRangeUnion] `json:"page_range"`
	// If True, persist the results indefinitely. Defaults to False.
	PersistResults param.Field[bool] `json:"persist_results"`
	// Whether to return images for the specified block types. By default, no images
	// are returned.
	ReturnImages param.Field[[]SplitRunParamsParsingSettingsReturnImage] `json:"return_images"`
	// If True, return OCR data in the result. Defaults to False.
	ReturnOcrData param.Field[bool] `json:"return_ocr_data"`
	// The timeout for the job in seconds. Defaults to 900.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r SplitRunParamsParsingSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Standard is our best multilingual OCR system. Legacy only supports germanic
// languages and is available for backwards compatibility.
type SplitRunParamsParsingSettingsOcrSystem string

const (
	SplitRunParamsParsingSettingsOcrSystemStandard SplitRunParamsParsingSettingsOcrSystem = "standard"
	SplitRunParamsParsingSettingsOcrSystemLegacy   SplitRunParamsParsingSettingsOcrSystem = "legacy"
)

func (r SplitRunParamsParsingSettingsOcrSystem) IsKnown() bool {
	switch r {
	case SplitRunParamsParsingSettingsOcrSystemStandard, SplitRunParamsParsingSettingsOcrSystemLegacy:
		return true
	}
	return false
}

// The page range to process (1-indexed). By default, the entire document is
// processed.
//
// Satisfied by [shared.PageRangeParam],
// [SplitRunParamsParsingSettingsPageRangeArray],
// [SplitRunParamsParsingSettingsPageRangeArray].
type SplitRunParamsParsingSettingsPageRangeUnion interface {
	ImplementsSplitRunParamsParsingSettingsPageRangeUnion()
}

type SplitRunParamsParsingSettingsPageRangeArray []shared.PageRangeParam

func (r SplitRunParamsParsingSettingsPageRangeArray) ImplementsSplitRunParamsParsingSettingsPageRangeUnion() {
}

type SplitRunParamsParsingSettingsReturnImage string

const (
	SplitRunParamsParsingSettingsReturnImageFigure SplitRunParamsParsingSettingsReturnImage = "figure"
	SplitRunParamsParsingSettingsReturnImageTable  SplitRunParamsParsingSettingsReturnImage = "table"
)

func (r SplitRunParamsParsingSettingsReturnImage) IsKnown() bool {
	switch r {
	case SplitRunParamsParsingSettingsReturnImageFigure, SplitRunParamsParsingSettingsReturnImageTable:
		return true
	}
	return false
}

type SplitRunParamsParsingSpreadsheet struct {
	// In a spreadsheet with different tables inside, we enable splitting up the tables
	// by default. Accurate mode applies more powerful models for superior accuracy, at
	// 5× the default per-cell rate. Disabling will register as one large table.
	Clustering param.Field[SplitRunParamsParsingSpreadsheetClustering] `json:"clustering"`
	// Whether to exclude hidden sheets, rows, or columns in the output.
	Exclude param.Field[[]SplitRunParamsParsingSpreadsheetExclude] `json:"exclude"`
	// Whether to include cell color and formula information in the output.
	Include          param.Field[[]SplitRunParamsParsingSpreadsheetInclude]        `json:"include"`
	SplitLargeTables param.Field[SplitRunParamsParsingSpreadsheetSplitLargeTables] `json:"split_large_tables"`
}

func (r SplitRunParamsParsingSpreadsheet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// In a spreadsheet with different tables inside, we enable splitting up the tables
// by default. Accurate mode applies more powerful models for superior accuracy, at
// 5× the default per-cell rate. Disabling will register as one large table.
type SplitRunParamsParsingSpreadsheetClustering string

const (
	SplitRunParamsParsingSpreadsheetClusteringAccurate SplitRunParamsParsingSpreadsheetClustering = "accurate"
	SplitRunParamsParsingSpreadsheetClusteringFast     SplitRunParamsParsingSpreadsheetClustering = "fast"
	SplitRunParamsParsingSpreadsheetClusteringDisabled SplitRunParamsParsingSpreadsheetClustering = "disabled"
)

func (r SplitRunParamsParsingSpreadsheetClustering) IsKnown() bool {
	switch r {
	case SplitRunParamsParsingSpreadsheetClusteringAccurate, SplitRunParamsParsingSpreadsheetClusteringFast, SplitRunParamsParsingSpreadsheetClusteringDisabled:
		return true
	}
	return false
}

type SplitRunParamsParsingSpreadsheetExclude string

const (
	SplitRunParamsParsingSpreadsheetExcludeHiddenSheets SplitRunParamsParsingSpreadsheetExclude = "hidden_sheets"
	SplitRunParamsParsingSpreadsheetExcludeHiddenRows   SplitRunParamsParsingSpreadsheetExclude = "hidden_rows"
	SplitRunParamsParsingSpreadsheetExcludeHiddenCols   SplitRunParamsParsingSpreadsheetExclude = "hidden_cols"
)

func (r SplitRunParamsParsingSpreadsheetExclude) IsKnown() bool {
	switch r {
	case SplitRunParamsParsingSpreadsheetExcludeHiddenSheets, SplitRunParamsParsingSpreadsheetExcludeHiddenRows, SplitRunParamsParsingSpreadsheetExcludeHiddenCols:
		return true
	}
	return false
}

type SplitRunParamsParsingSpreadsheetInclude string

const (
	SplitRunParamsParsingSpreadsheetIncludeCellColors SplitRunParamsParsingSpreadsheetInclude = "cell_colors"
	SplitRunParamsParsingSpreadsheetIncludeFormula    SplitRunParamsParsingSpreadsheetInclude = "formula"
)

func (r SplitRunParamsParsingSpreadsheetInclude) IsKnown() bool {
	switch r {
	case SplitRunParamsParsingSpreadsheetIncludeCellColors, SplitRunParamsParsingSpreadsheetIncludeFormula:
		return true
	}
	return false
}

type SplitRunParamsParsingSpreadsheetSplitLargeTables struct {
	// If True, split large tables into smaller tables. Defaults to True.
	Enabled param.Field[bool] `json:"enabled"`
	// The size of the tables to split into. Defaults to 50.
	Size param.Field[int64] `json:"size"`
}

func (r SplitRunParamsParsingSpreadsheetSplitLargeTables) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The settings for split processing.
type SplitRunParamsSettings struct {
	// If tables should be truncated to the first few rows or if all content should be
	// preserved. truncate improves latency, preserve is recommended for cases where
	// partition_key is being used and the partition_key may be included within the
	// table. Defaults to truncate
	TableCutoff param.Field[SplitRunParamsSettingsTableCutoff] `json:"table_cutoff"`
}

func (r SplitRunParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If tables should be truncated to the first few rows or if all content should be
// preserved. truncate improves latency, preserve is recommended for cases where
// partition_key is being used and the partition_key may be included within the
// table. Defaults to truncate
type SplitRunParamsSettingsTableCutoff string

const (
	SplitRunParamsSettingsTableCutoffTruncate SplitRunParamsSettingsTableCutoff = "truncate"
	SplitRunParamsSettingsTableCutoffPreserve SplitRunParamsSettingsTableCutoff = "preserve"
)

func (r SplitRunParamsSettingsTableCutoff) IsKnown() bool {
	switch r {
	case SplitRunParamsSettingsTableCutoffTruncate, SplitRunParamsSettingsTableCutoffPreserve:
		return true
	}
	return false
}

type SplitRunJobParams struct {
	Body interface{} `json:"body,required"`
}

func (r SplitRunJobParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
