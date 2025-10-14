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
	return
}

// Async Parse
func (r *ParseService) RunJob(ctx context.Context, body ParseRunJobParams, opts ...option.RequestOption) (res *ParseRunJobResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "parse_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ParseRunResponse struct {
	JobID string `json:"job_id,required"`
	// The duration of the parse request in seconds.
	Duration float64 `json:"duration"`
	// The storage URL of the converted PDF file.
	PdfURL string `json:"pdf_url,nullable"`
	// This field can have the runtime type of [shared.ParseResponseResult].
	Result interface{} `json:"result"`
	// The link to the studio pipeline for the document.
	StudioLink string               `json:"studio_link,nullable"`
	Usage      shared.ParseUsage    `json:"usage"`
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
// Possible runtime types of the union are [shared.ParseResponse],
// [ParseRunResponseAsyncParseResponse].
func (r ParseRunResponse) AsUnion() ParseRunResponseUnion {
	return r.union
}

// Union satisfied by [shared.ParseResponse] or
// [ParseRunResponseAsyncParseResponse].
type ParseRunResponseUnion interface {
	ImplementsParseRunResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ParseRunResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.ParseResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParseRunResponseAsyncParseResponse{}),
		},
	)
}

type ParseRunResponseAsyncParseResponse struct {
	JobID string                                 `json:"job_id,required"`
	JSON  parseRunResponseAsyncParseResponseJSON `json:"-"`
}

// parseRunResponseAsyncParseResponseJSON contains the JSON metadata for the struct
// [ParseRunResponseAsyncParseResponse]
type parseRunResponseAsyncParseResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParseRunResponseAsyncParseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseRunResponseAsyncParseResponseJSON) RawJSON() string {
	return r.raw
}

func (r ParseRunResponseAsyncParseResponse) ImplementsParseRunResponse() {}

type ParseRunJobResponse struct {
	JobID string                  `json:"job_id,required"`
	JSON  parseRunJobResponseJSON `json:"-"`
}

// parseRunJobResponseJSON contains the JSON metadata for the struct
// [ParseRunJobResponse]
type parseRunJobResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParseRunJobResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseRunJobResponseJSON) RawJSON() string {
	return r.raw
}

type ParseRunParams struct {
	Body ParseRunParamsBodyUnion `json:"body,required"`
}

func (r ParseRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ParseRunParamsBody struct {
	Input       param.Field[interface{}] `json:"input,required"`
	Async       param.Field[interface{}] `json:"async"`
	Enhance     param.Field[interface{}] `json:"enhance"`
	Formatting  param.Field[interface{}] `json:"formatting"`
	Retrieval   param.Field[interface{}] `json:"retrieval"`
	Settings    param.Field[interface{}] `json:"settings"`
	Spreadsheet param.Field[interface{}] `json:"spreadsheet"`
}

func (r ParseRunParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBody) implementsParseRunParamsBodyUnion() {}

// Satisfied by [ParseRunParamsBodySyncParseConfig],
// [ParseRunParamsBodyAsyncParseConfig], [ParseRunParamsBody].
type ParseRunParamsBodyUnion interface {
	implementsParseRunParamsBodyUnion()
}

type ParseRunParamsBodySyncParseConfig struct {
	// The URL of the document to be processed. You can provide one of the
	// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
	// prefixed URL obtained from the /upload endpoint after directly uploading a
	// document 4. A jobid:// prefixed URL obtained from a previous /parse invocation
	Input       param.Field[ParseRunParamsBodySyncParseConfigInputUnion]  `json:"input,required"`
	Enhance     param.Field[ParseRunParamsBodySyncParseConfigEnhance]     `json:"enhance"`
	Formatting  param.Field[ParseRunParamsBodySyncParseConfigFormatting]  `json:"formatting"`
	Retrieval   param.Field[ParseRunParamsBodySyncParseConfigRetrieval]   `json:"retrieval"`
	Settings    param.Field[ParseRunParamsBodySyncParseConfigSettings]    `json:"settings"`
	Spreadsheet param.Field[ParseRunParamsBodySyncParseConfigSpreadsheet] `json:"spreadsheet"`
}

func (r ParseRunParamsBodySyncParseConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBodySyncParseConfig) implementsParseRunParamsBodyUnion() {}

// The URL of the document to be processed. You can provide one of the
// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
// prefixed URL obtained from the /upload endpoint after directly uploading a
// document 4. A jobid:// prefixed URL obtained from a previous /parse invocation
//
// Satisfied by [shared.UnionString], [shared.UploadParam].
type ParseRunParamsBodySyncParseConfigInputUnion interface {
	ImplementsParseRunParamsBodySyncParseConfigInputUnion()
}

type ParseRunParamsBodySyncParseConfigEnhance struct {
	// Agentic uses vision language models to enhance the accuracy of the output of
	// different types of extraction. This will incur a cost and latency increase.
	Agentic param.Field[[]ParseRunParamsBodySyncParseConfigEnhanceAgenticUnion] `json:"agentic"`
	// If True, summarize figures using a small vision language model. Defaults to
	// True.
	SummarizeFigures param.Field[bool] `json:"summarize_figures"`
}

func (r ParseRunParamsBodySyncParseConfigEnhance) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ParseRunParamsBodySyncParseConfigEnhanceAgentic struct {
	Scope param.Field[ParseRunParamsBodySyncParseConfigEnhanceAgenticScope] `json:"scope,required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ParseRunParamsBodySyncParseConfigEnhanceAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBodySyncParseConfigEnhanceAgentic) implementsParseRunParamsBodySyncParseConfigEnhanceAgenticUnion() {
}

// Satisfied by [ParseRunParamsBodySyncParseConfigEnhanceAgenticTableAgentic],
// [ParseRunParamsBodySyncParseConfigEnhanceAgenticFigureAgentic],
// [ParseRunParamsBodySyncParseConfigEnhanceAgenticTextAgentic],
// [ParseRunParamsBodySyncParseConfigEnhanceAgentic].
type ParseRunParamsBodySyncParseConfigEnhanceAgenticUnion interface {
	implementsParseRunParamsBodySyncParseConfigEnhanceAgenticUnion()
}

type ParseRunParamsBodySyncParseConfigEnhanceAgenticTableAgentic struct {
	Scope param.Field[ParseRunParamsBodySyncParseConfigEnhanceAgenticTableAgenticScope] `json:"scope,required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ParseRunParamsBodySyncParseConfigEnhanceAgenticTableAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBodySyncParseConfigEnhanceAgenticTableAgentic) implementsParseRunParamsBodySyncParseConfigEnhanceAgenticUnion() {
}

type ParseRunParamsBodySyncParseConfigEnhanceAgenticTableAgenticScope string

const (
	ParseRunParamsBodySyncParseConfigEnhanceAgenticTableAgenticScopeTable ParseRunParamsBodySyncParseConfigEnhanceAgenticTableAgenticScope = "table"
)

func (r ParseRunParamsBodySyncParseConfigEnhanceAgenticTableAgenticScope) IsKnown() bool {
	switch r {
	case ParseRunParamsBodySyncParseConfigEnhanceAgenticTableAgenticScopeTable:
		return true
	}
	return false
}

type ParseRunParamsBodySyncParseConfigEnhanceAgenticFigureAgentic struct {
	Scope param.Field[ParseRunParamsBodySyncParseConfigEnhanceAgenticFigureAgenticScope] `json:"scope,required"`
	// Custom prompt for figure agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ParseRunParamsBodySyncParseConfigEnhanceAgenticFigureAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBodySyncParseConfigEnhanceAgenticFigureAgentic) implementsParseRunParamsBodySyncParseConfigEnhanceAgenticUnion() {
}

type ParseRunParamsBodySyncParseConfigEnhanceAgenticFigureAgenticScope string

const (
	ParseRunParamsBodySyncParseConfigEnhanceAgenticFigureAgenticScopeFigure ParseRunParamsBodySyncParseConfigEnhanceAgenticFigureAgenticScope = "figure"
)

func (r ParseRunParamsBodySyncParseConfigEnhanceAgenticFigureAgenticScope) IsKnown() bool {
	switch r {
	case ParseRunParamsBodySyncParseConfigEnhanceAgenticFigureAgenticScopeFigure:
		return true
	}
	return false
}

type ParseRunParamsBodySyncParseConfigEnhanceAgenticTextAgentic struct {
	Scope param.Field[ParseRunParamsBodySyncParseConfigEnhanceAgenticTextAgenticScope] `json:"scope,required"`
}

func (r ParseRunParamsBodySyncParseConfigEnhanceAgenticTextAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBodySyncParseConfigEnhanceAgenticTextAgentic) implementsParseRunParamsBodySyncParseConfigEnhanceAgenticUnion() {
}

type ParseRunParamsBodySyncParseConfigEnhanceAgenticTextAgenticScope string

const (
	ParseRunParamsBodySyncParseConfigEnhanceAgenticTextAgenticScopeText ParseRunParamsBodySyncParseConfigEnhanceAgenticTextAgenticScope = "text"
)

func (r ParseRunParamsBodySyncParseConfigEnhanceAgenticTextAgenticScope) IsKnown() bool {
	switch r {
	case ParseRunParamsBodySyncParseConfigEnhanceAgenticTextAgenticScopeText:
		return true
	}
	return false
}

type ParseRunParamsBodySyncParseConfigEnhanceAgenticScope string

const (
	ParseRunParamsBodySyncParseConfigEnhanceAgenticScopeTable  ParseRunParamsBodySyncParseConfigEnhanceAgenticScope = "table"
	ParseRunParamsBodySyncParseConfigEnhanceAgenticScopeFigure ParseRunParamsBodySyncParseConfigEnhanceAgenticScope = "figure"
	ParseRunParamsBodySyncParseConfigEnhanceAgenticScopeText   ParseRunParamsBodySyncParseConfigEnhanceAgenticScope = "text"
)

func (r ParseRunParamsBodySyncParseConfigEnhanceAgenticScope) IsKnown() bool {
	switch r {
	case ParseRunParamsBodySyncParseConfigEnhanceAgenticScopeTable, ParseRunParamsBodySyncParseConfigEnhanceAgenticScopeFigure, ParseRunParamsBodySyncParseConfigEnhanceAgenticScopeText:
		return true
	}
	return false
}

type ParseRunParamsBodySyncParseConfigFormatting struct {
	// If True, add page markers to the output. Defaults to False. Useful for
	// extracting data with page specific information.
	AddPageMarkers param.Field[bool] `json:"add_page_markers"`
	// A list of formatting to include in the output. [insert description of each
	// option here later]
	Include param.Field[[]ParseRunParamsBodySyncParseConfigFormattingInclude] `json:"include"`
	// A flag to indicate if consecutive tables with the same number of columns should
	// be merged. Defaults to False.
	MergeTables param.Field[bool] `json:"merge_tables"`
	// The mode to use for table output. Defaults to dynamic, which returns md for
	// simpler tables and html for more complex tables.
	TableOutputFormat param.Field[ParseRunParamsBodySyncParseConfigFormattingTableOutputFormat] `json:"table_output_format"`
}

func (r ParseRunParamsBodySyncParseConfigFormatting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ParseRunParamsBodySyncParseConfigFormattingInclude string

const (
	ParseRunParamsBodySyncParseConfigFormattingIncludeChangeTracking ParseRunParamsBodySyncParseConfigFormattingInclude = "change_tracking"
	ParseRunParamsBodySyncParseConfigFormattingIncludeHighlight      ParseRunParamsBodySyncParseConfigFormattingInclude = "highlight"
	ParseRunParamsBodySyncParseConfigFormattingIncludeComments       ParseRunParamsBodySyncParseConfigFormattingInclude = "comments"
)

func (r ParseRunParamsBodySyncParseConfigFormattingInclude) IsKnown() bool {
	switch r {
	case ParseRunParamsBodySyncParseConfigFormattingIncludeChangeTracking, ParseRunParamsBodySyncParseConfigFormattingIncludeHighlight, ParseRunParamsBodySyncParseConfigFormattingIncludeComments:
		return true
	}
	return false
}

// The mode to use for table output. Defaults to dynamic, which returns md for
// simpler tables and html for more complex tables.
type ParseRunParamsBodySyncParseConfigFormattingTableOutputFormat string

const (
	ParseRunParamsBodySyncParseConfigFormattingTableOutputFormatHTML     ParseRunParamsBodySyncParseConfigFormattingTableOutputFormat = "html"
	ParseRunParamsBodySyncParseConfigFormattingTableOutputFormatJson     ParseRunParamsBodySyncParseConfigFormattingTableOutputFormat = "json"
	ParseRunParamsBodySyncParseConfigFormattingTableOutputFormatMd       ParseRunParamsBodySyncParseConfigFormattingTableOutputFormat = "md"
	ParseRunParamsBodySyncParseConfigFormattingTableOutputFormatJsonbbox ParseRunParamsBodySyncParseConfigFormattingTableOutputFormat = "jsonbbox"
	ParseRunParamsBodySyncParseConfigFormattingTableOutputFormatDynamic  ParseRunParamsBodySyncParseConfigFormattingTableOutputFormat = "dynamic"
	ParseRunParamsBodySyncParseConfigFormattingTableOutputFormatCsv      ParseRunParamsBodySyncParseConfigFormattingTableOutputFormat = "csv"
)

func (r ParseRunParamsBodySyncParseConfigFormattingTableOutputFormat) IsKnown() bool {
	switch r {
	case ParseRunParamsBodySyncParseConfigFormattingTableOutputFormatHTML, ParseRunParamsBodySyncParseConfigFormattingTableOutputFormatJson, ParseRunParamsBodySyncParseConfigFormattingTableOutputFormatMd, ParseRunParamsBodySyncParseConfigFormattingTableOutputFormatJsonbbox, ParseRunParamsBodySyncParseConfigFormattingTableOutputFormatDynamic, ParseRunParamsBodySyncParseConfigFormattingTableOutputFormatCsv:
		return true
	}
	return false
}

type ParseRunParamsBodySyncParseConfigRetrieval struct {
	Chunking param.Field[ParseRunParamsBodySyncParseConfigRetrievalChunking] `json:"chunking"`
	// If True, use embedding optimized mode. Defaults to False.
	EmbeddingOptimized param.Field[bool] `json:"embedding_optimized"`
	// A list of block types to filter out from 'content' and 'embed' fields. By
	// default, no blocks are filtered.
	FilterBlocks param.Field[[]ParseRunParamsBodySyncParseConfigRetrievalFilterBlock] `json:"filter_blocks"`
}

func (r ParseRunParamsBodySyncParseConfigRetrieval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ParseRunParamsBodySyncParseConfigRetrievalChunking struct {
	// Choose how to partition chunks. Variable mode chunks by character length and
	// visual context. Section mode chunks by section headers. Page mode chunks
	// according to pages. Page sections mode chunks first by page, then by sections
	// within each page. Disabled returns one single chunk.
	ChunkMode param.Field[ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkMode] `json:"chunk_mode"`
	// The approximate size of chunks (in characters) that the document will be split
	// into. Defaults to null, in which case the chunk size is variable between 250 -
	// 1500 characters.
	ChunkSize param.Field[int64] `json:"chunk_size"`
}

func (r ParseRunParamsBodySyncParseConfigRetrievalChunking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Choose how to partition chunks. Variable mode chunks by character length and
// visual context. Section mode chunks by section headers. Page mode chunks
// according to pages. Page sections mode chunks first by page, then by sections
// within each page. Disabled returns one single chunk.
type ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkMode string

const (
	ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkModeVariable     ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkMode = "variable"
	ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkModeSection      ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkMode = "section"
	ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkModePage         ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkMode = "page"
	ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkModeDisabled     ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkMode = "disabled"
	ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkModeBlock        ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkMode = "block"
	ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkModePageSections ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkMode = "page_sections"
)

func (r ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkMode) IsKnown() bool {
	switch r {
	case ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkModeVariable, ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkModeSection, ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkModePage, ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkModeDisabled, ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkModeBlock, ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkModePageSections:
		return true
	}
	return false
}

type ParseRunParamsBodySyncParseConfigRetrievalFilterBlock string

const (
	ParseRunParamsBodySyncParseConfigRetrievalFilterBlockHeader        ParseRunParamsBodySyncParseConfigRetrievalFilterBlock = "Header"
	ParseRunParamsBodySyncParseConfigRetrievalFilterBlockFooter        ParseRunParamsBodySyncParseConfigRetrievalFilterBlock = "Footer"
	ParseRunParamsBodySyncParseConfigRetrievalFilterBlockTitle         ParseRunParamsBodySyncParseConfigRetrievalFilterBlock = "Title"
	ParseRunParamsBodySyncParseConfigRetrievalFilterBlockSectionHeader ParseRunParamsBodySyncParseConfigRetrievalFilterBlock = "Section Header"
	ParseRunParamsBodySyncParseConfigRetrievalFilterBlockPageNumber    ParseRunParamsBodySyncParseConfigRetrievalFilterBlock = "Page Number"
	ParseRunParamsBodySyncParseConfigRetrievalFilterBlockListItem      ParseRunParamsBodySyncParseConfigRetrievalFilterBlock = "List Item"
	ParseRunParamsBodySyncParseConfigRetrievalFilterBlockFigure        ParseRunParamsBodySyncParseConfigRetrievalFilterBlock = "Figure"
	ParseRunParamsBodySyncParseConfigRetrievalFilterBlockTable         ParseRunParamsBodySyncParseConfigRetrievalFilterBlock = "Table"
	ParseRunParamsBodySyncParseConfigRetrievalFilterBlockKeyValue      ParseRunParamsBodySyncParseConfigRetrievalFilterBlock = "Key Value"
	ParseRunParamsBodySyncParseConfigRetrievalFilterBlockText          ParseRunParamsBodySyncParseConfigRetrievalFilterBlock = "Text"
	ParseRunParamsBodySyncParseConfigRetrievalFilterBlockComment       ParseRunParamsBodySyncParseConfigRetrievalFilterBlock = "Comment"
	ParseRunParamsBodySyncParseConfigRetrievalFilterBlockSignature     ParseRunParamsBodySyncParseConfigRetrievalFilterBlock = "Signature"
)

func (r ParseRunParamsBodySyncParseConfigRetrievalFilterBlock) IsKnown() bool {
	switch r {
	case ParseRunParamsBodySyncParseConfigRetrievalFilterBlockHeader, ParseRunParamsBodySyncParseConfigRetrievalFilterBlockFooter, ParseRunParamsBodySyncParseConfigRetrievalFilterBlockTitle, ParseRunParamsBodySyncParseConfigRetrievalFilterBlockSectionHeader, ParseRunParamsBodySyncParseConfigRetrievalFilterBlockPageNumber, ParseRunParamsBodySyncParseConfigRetrievalFilterBlockListItem, ParseRunParamsBodySyncParseConfigRetrievalFilterBlockFigure, ParseRunParamsBodySyncParseConfigRetrievalFilterBlockTable, ParseRunParamsBodySyncParseConfigRetrievalFilterBlockKeyValue, ParseRunParamsBodySyncParseConfigRetrievalFilterBlockText, ParseRunParamsBodySyncParseConfigRetrievalFilterBlockComment, ParseRunParamsBodySyncParseConfigRetrievalFilterBlockSignature:
		return true
	}
	return false
}

type ParseRunParamsBodySyncParseConfigSettings struct {
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
	OcrSystem param.Field[ParseRunParamsBodySyncParseConfigSettingsOcrSystem] `json:"ocr_system"`
	// The page range to process (1-indexed). By default, the entire document is
	// processed.
	PageRange param.Field[ParseRunParamsBodySyncParseConfigSettingsPageRangeUnion] `json:"page_range"`
	// If True, persist the results indefinitely. Defaults to False.
	PersistResults param.Field[bool] `json:"persist_results"`
	// Whether to return images for the specified block types. By default, no images
	// are returned.
	ReturnImages param.Field[[]ParseRunParamsBodySyncParseConfigSettingsReturnImage] `json:"return_images"`
	// If True, return OCR data in the result. Defaults to False.
	ReturnOcrData param.Field[bool] `json:"return_ocr_data"`
	// The timeout for the job in seconds. Defaults to 900.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r ParseRunParamsBodySyncParseConfigSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Standard is our best multilingual OCR system. Legacy only supports germanic
// languages and is available for backwards compatibility.
type ParseRunParamsBodySyncParseConfigSettingsOcrSystem string

const (
	ParseRunParamsBodySyncParseConfigSettingsOcrSystemStandard ParseRunParamsBodySyncParseConfigSettingsOcrSystem = "standard"
	ParseRunParamsBodySyncParseConfigSettingsOcrSystemLegacy   ParseRunParamsBodySyncParseConfigSettingsOcrSystem = "legacy"
)

func (r ParseRunParamsBodySyncParseConfigSettingsOcrSystem) IsKnown() bool {
	switch r {
	case ParseRunParamsBodySyncParseConfigSettingsOcrSystemStandard, ParseRunParamsBodySyncParseConfigSettingsOcrSystemLegacy:
		return true
	}
	return false
}

// The page range to process (1-indexed). By default, the entire document is
// processed.
//
// Satisfied by [shared.PageRangeParam],
// [ParseRunParamsBodySyncParseConfigSettingsPageRangeArray],
// [ParseRunParamsBodySyncParseConfigSettingsPageRangeArray].
type ParseRunParamsBodySyncParseConfigSettingsPageRangeUnion interface {
	ImplementsParseRunParamsBodySyncParseConfigSettingsPageRangeUnion()
}

type ParseRunParamsBodySyncParseConfigSettingsPageRangeArray []shared.PageRangeParam

func (r ParseRunParamsBodySyncParseConfigSettingsPageRangeArray) ImplementsParseRunParamsBodySyncParseConfigSettingsPageRangeUnion() {
}

type ParseRunParamsBodySyncParseConfigSettingsReturnImage string

const (
	ParseRunParamsBodySyncParseConfigSettingsReturnImageFigure ParseRunParamsBodySyncParseConfigSettingsReturnImage = "figure"
	ParseRunParamsBodySyncParseConfigSettingsReturnImageTable  ParseRunParamsBodySyncParseConfigSettingsReturnImage = "table"
)

func (r ParseRunParamsBodySyncParseConfigSettingsReturnImage) IsKnown() bool {
	switch r {
	case ParseRunParamsBodySyncParseConfigSettingsReturnImageFigure, ParseRunParamsBodySyncParseConfigSettingsReturnImageTable:
		return true
	}
	return false
}

type ParseRunParamsBodySyncParseConfigSpreadsheet struct {
	// In a spreadsheet with different tables inside, we enable splitting up the tables
	// by default. Accurate mode applies more powerful models for superior accuracy, at
	// 5× the default per-cell rate. Disabling will register as one large table.
	Clustering param.Field[ParseRunParamsBodySyncParseConfigSpreadsheetClustering] `json:"clustering"`
	// Whether to exclude hidden sheets, rows, or columns in the output.
	Exclude param.Field[[]ParseRunParamsBodySyncParseConfigSpreadsheetExclude] `json:"exclude"`
	// Whether to include cell color and formula information in the output.
	Include          param.Field[[]ParseRunParamsBodySyncParseConfigSpreadsheetInclude]        `json:"include"`
	SplitLargeTables param.Field[ParseRunParamsBodySyncParseConfigSpreadsheetSplitLargeTables] `json:"split_large_tables"`
}

func (r ParseRunParamsBodySyncParseConfigSpreadsheet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// In a spreadsheet with different tables inside, we enable splitting up the tables
// by default. Accurate mode applies more powerful models for superior accuracy, at
// 5× the default per-cell rate. Disabling will register as one large table.
type ParseRunParamsBodySyncParseConfigSpreadsheetClustering string

const (
	ParseRunParamsBodySyncParseConfigSpreadsheetClusteringAccurate ParseRunParamsBodySyncParseConfigSpreadsheetClustering = "accurate"
	ParseRunParamsBodySyncParseConfigSpreadsheetClusteringFast     ParseRunParamsBodySyncParseConfigSpreadsheetClustering = "fast"
	ParseRunParamsBodySyncParseConfigSpreadsheetClusteringDisabled ParseRunParamsBodySyncParseConfigSpreadsheetClustering = "disabled"
)

func (r ParseRunParamsBodySyncParseConfigSpreadsheetClustering) IsKnown() bool {
	switch r {
	case ParseRunParamsBodySyncParseConfigSpreadsheetClusteringAccurate, ParseRunParamsBodySyncParseConfigSpreadsheetClusteringFast, ParseRunParamsBodySyncParseConfigSpreadsheetClusteringDisabled:
		return true
	}
	return false
}

type ParseRunParamsBodySyncParseConfigSpreadsheetExclude string

const (
	ParseRunParamsBodySyncParseConfigSpreadsheetExcludeHiddenSheets ParseRunParamsBodySyncParseConfigSpreadsheetExclude = "hidden_sheets"
	ParseRunParamsBodySyncParseConfigSpreadsheetExcludeHiddenRows   ParseRunParamsBodySyncParseConfigSpreadsheetExclude = "hidden_rows"
	ParseRunParamsBodySyncParseConfigSpreadsheetExcludeHiddenCols   ParseRunParamsBodySyncParseConfigSpreadsheetExclude = "hidden_cols"
)

func (r ParseRunParamsBodySyncParseConfigSpreadsheetExclude) IsKnown() bool {
	switch r {
	case ParseRunParamsBodySyncParseConfigSpreadsheetExcludeHiddenSheets, ParseRunParamsBodySyncParseConfigSpreadsheetExcludeHiddenRows, ParseRunParamsBodySyncParseConfigSpreadsheetExcludeHiddenCols:
		return true
	}
	return false
}

type ParseRunParamsBodySyncParseConfigSpreadsheetInclude string

const (
	ParseRunParamsBodySyncParseConfigSpreadsheetIncludeCellColors ParseRunParamsBodySyncParseConfigSpreadsheetInclude = "cell_colors"
	ParseRunParamsBodySyncParseConfigSpreadsheetIncludeFormula    ParseRunParamsBodySyncParseConfigSpreadsheetInclude = "formula"
)

func (r ParseRunParamsBodySyncParseConfigSpreadsheetInclude) IsKnown() bool {
	switch r {
	case ParseRunParamsBodySyncParseConfigSpreadsheetIncludeCellColors, ParseRunParamsBodySyncParseConfigSpreadsheetIncludeFormula:
		return true
	}
	return false
}

type ParseRunParamsBodySyncParseConfigSpreadsheetSplitLargeTables struct {
	// If True, split large tables into smaller tables. Defaults to True.
	Enabled param.Field[bool] `json:"enabled"`
	// The size of the tables to split into. Defaults to 50.
	Size param.Field[int64] `json:"size"`
}

func (r ParseRunParamsBodySyncParseConfigSpreadsheetSplitLargeTables) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ParseRunParamsBodyAsyncParseConfig struct {
	// The URL of the document to be processed. You can provide one of the
	// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
	// prefixed URL obtained from the /upload endpoint after directly uploading a
	// document 4. A jobid:// prefixed URL obtained from a previous /parse invocation
	Input param.Field[ParseRunParamsBodyAsyncParseConfigInputUnion] `json:"input,required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async       param.Field[ParseRunParamsBodyAsyncParseConfigAsync]       `json:"async"`
	Enhance     param.Field[ParseRunParamsBodyAsyncParseConfigEnhance]     `json:"enhance"`
	Formatting  param.Field[ParseRunParamsBodyAsyncParseConfigFormatting]  `json:"formatting"`
	Retrieval   param.Field[ParseRunParamsBodyAsyncParseConfigRetrieval]   `json:"retrieval"`
	Settings    param.Field[ParseRunParamsBodyAsyncParseConfigSettings]    `json:"settings"`
	Spreadsheet param.Field[ParseRunParamsBodyAsyncParseConfigSpreadsheet] `json:"spreadsheet"`
}

func (r ParseRunParamsBodyAsyncParseConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBodyAsyncParseConfig) implementsParseRunParamsBodyUnion() {}

// The URL of the document to be processed. You can provide one of the
// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
// prefixed URL obtained from the /upload endpoint after directly uploading a
// document 4. A jobid:// prefixed URL obtained from a previous /parse invocation
//
// Satisfied by [shared.UnionString], [shared.UploadParam].
type ParseRunParamsBodyAsyncParseConfigInputUnion interface {
	ImplementsParseRunParamsBodyAsyncParseConfigInputUnion()
}

// The configuration options for asynchronous processing (default synchronous).
type ParseRunParamsBodyAsyncParseConfigAsync struct {
	// JSON metadata included in webhook request body. Defaults to None.
	Metadata param.Field[interface{}] `json:"metadata"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool] `json:"priority"`
	// The webhook configuration for the asynchronous processing.
	Webhook param.Field[ParseRunParamsBodyAsyncParseConfigAsyncWebhookUnion] `json:"webhook"`
}

func (r ParseRunParamsBodyAsyncParseConfigAsync) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The webhook configuration for the asynchronous processing.
type ParseRunParamsBodyAsyncParseConfigAsyncWebhook struct {
	Channels param.Field[interface{}]                                        `json:"channels"`
	Mode     param.Field[ParseRunParamsBodyAsyncParseConfigAsyncWebhookMode] `json:"mode"`
	URL      param.Field[string]                                             `json:"url"`
}

func (r ParseRunParamsBodyAsyncParseConfigAsyncWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBodyAsyncParseConfigAsyncWebhook) implementsParseRunParamsBodyAsyncParseConfigAsyncWebhookUnion() {
}

// The webhook configuration for the asynchronous processing.
//
// Satisfied by [ParseRunParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfig],
// [ParseRunParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfig],
// [ParseRunParamsBodyAsyncParseConfigAsyncWebhook].
type ParseRunParamsBodyAsyncParseConfigAsyncWebhookUnion interface {
	implementsParseRunParamsBodyAsyncParseConfigAsyncWebhookUnion()
}

type ParseRunParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfig struct {
	// A list of Svix channels the message will be delivered down, omit to send to all
	// channels.
	Channels param.Field[[]string]                                                            `json:"channels"`
	Mode     param.Field[ParseRunParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfigMode] `json:"mode"`
}

func (r ParseRunParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfig) implementsParseRunParamsBodyAsyncParseConfigAsyncWebhookUnion() {
}

type ParseRunParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfigMode string

const (
	ParseRunParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfigModeSvix ParseRunParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfigMode = "svix"
)

func (r ParseRunParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfigMode) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfigModeSvix:
		return true
	}
	return false
}

type ParseRunParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfig struct {
	URL  param.Field[string]                                                                `json:"url,required"`
	Mode param.Field[ParseRunParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfigMode] `json:"mode"`
}

func (r ParseRunParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfig) implementsParseRunParamsBodyAsyncParseConfigAsyncWebhookUnion() {
}

type ParseRunParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfigMode string

const (
	ParseRunParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfigModeDirect ParseRunParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfigMode = "direct"
)

func (r ParseRunParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfigMode) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfigModeDirect:
		return true
	}
	return false
}

type ParseRunParamsBodyAsyncParseConfigAsyncWebhookMode string

const (
	ParseRunParamsBodyAsyncParseConfigAsyncWebhookModeSvix   ParseRunParamsBodyAsyncParseConfigAsyncWebhookMode = "svix"
	ParseRunParamsBodyAsyncParseConfigAsyncWebhookModeDirect ParseRunParamsBodyAsyncParseConfigAsyncWebhookMode = "direct"
)

func (r ParseRunParamsBodyAsyncParseConfigAsyncWebhookMode) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyAsyncParseConfigAsyncWebhookModeSvix, ParseRunParamsBodyAsyncParseConfigAsyncWebhookModeDirect:
		return true
	}
	return false
}

type ParseRunParamsBodyAsyncParseConfigEnhance struct {
	// Agentic uses vision language models to enhance the accuracy of the output of
	// different types of extraction. This will incur a cost and latency increase.
	Agentic param.Field[[]ParseRunParamsBodyAsyncParseConfigEnhanceAgenticUnion] `json:"agentic"`
	// If True, summarize figures using a small vision language model. Defaults to
	// True.
	SummarizeFigures param.Field[bool] `json:"summarize_figures"`
}

func (r ParseRunParamsBodyAsyncParseConfigEnhance) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ParseRunParamsBodyAsyncParseConfigEnhanceAgentic struct {
	Scope param.Field[ParseRunParamsBodyAsyncParseConfigEnhanceAgenticScope] `json:"scope,required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ParseRunParamsBodyAsyncParseConfigEnhanceAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBodyAsyncParseConfigEnhanceAgentic) implementsParseRunParamsBodyAsyncParseConfigEnhanceAgenticUnion() {
}

// Satisfied by [ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTableAgentic],
// [ParseRunParamsBodyAsyncParseConfigEnhanceAgenticFigureAgentic],
// [ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTextAgentic],
// [ParseRunParamsBodyAsyncParseConfigEnhanceAgentic].
type ParseRunParamsBodyAsyncParseConfigEnhanceAgenticUnion interface {
	implementsParseRunParamsBodyAsyncParseConfigEnhanceAgenticUnion()
}

type ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTableAgentic struct {
	Scope param.Field[ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTableAgenticScope] `json:"scope,required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTableAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTableAgentic) implementsParseRunParamsBodyAsyncParseConfigEnhanceAgenticUnion() {
}

type ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTableAgenticScope string

const (
	ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTableAgenticScopeTable ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTableAgenticScope = "table"
)

func (r ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTableAgenticScope) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTableAgenticScopeTable:
		return true
	}
	return false
}

type ParseRunParamsBodyAsyncParseConfigEnhanceAgenticFigureAgentic struct {
	Scope param.Field[ParseRunParamsBodyAsyncParseConfigEnhanceAgenticFigureAgenticScope] `json:"scope,required"`
	// Custom prompt for figure agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ParseRunParamsBodyAsyncParseConfigEnhanceAgenticFigureAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBodyAsyncParseConfigEnhanceAgenticFigureAgentic) implementsParseRunParamsBodyAsyncParseConfigEnhanceAgenticUnion() {
}

type ParseRunParamsBodyAsyncParseConfigEnhanceAgenticFigureAgenticScope string

const (
	ParseRunParamsBodyAsyncParseConfigEnhanceAgenticFigureAgenticScopeFigure ParseRunParamsBodyAsyncParseConfigEnhanceAgenticFigureAgenticScope = "figure"
)

func (r ParseRunParamsBodyAsyncParseConfigEnhanceAgenticFigureAgenticScope) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyAsyncParseConfigEnhanceAgenticFigureAgenticScopeFigure:
		return true
	}
	return false
}

type ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTextAgentic struct {
	Scope param.Field[ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTextAgenticScope] `json:"scope,required"`
}

func (r ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTextAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTextAgentic) implementsParseRunParamsBodyAsyncParseConfigEnhanceAgenticUnion() {
}

type ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTextAgenticScope string

const (
	ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTextAgenticScopeText ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTextAgenticScope = "text"
)

func (r ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTextAgenticScope) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyAsyncParseConfigEnhanceAgenticTextAgenticScopeText:
		return true
	}
	return false
}

type ParseRunParamsBodyAsyncParseConfigEnhanceAgenticScope string

const (
	ParseRunParamsBodyAsyncParseConfigEnhanceAgenticScopeTable  ParseRunParamsBodyAsyncParseConfigEnhanceAgenticScope = "table"
	ParseRunParamsBodyAsyncParseConfigEnhanceAgenticScopeFigure ParseRunParamsBodyAsyncParseConfigEnhanceAgenticScope = "figure"
	ParseRunParamsBodyAsyncParseConfigEnhanceAgenticScopeText   ParseRunParamsBodyAsyncParseConfigEnhanceAgenticScope = "text"
)

func (r ParseRunParamsBodyAsyncParseConfigEnhanceAgenticScope) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyAsyncParseConfigEnhanceAgenticScopeTable, ParseRunParamsBodyAsyncParseConfigEnhanceAgenticScopeFigure, ParseRunParamsBodyAsyncParseConfigEnhanceAgenticScopeText:
		return true
	}
	return false
}

type ParseRunParamsBodyAsyncParseConfigFormatting struct {
	// If True, add page markers to the output. Defaults to False. Useful for
	// extracting data with page specific information.
	AddPageMarkers param.Field[bool] `json:"add_page_markers"`
	// A list of formatting to include in the output. [insert description of each
	// option here later]
	Include param.Field[[]ParseRunParamsBodyAsyncParseConfigFormattingInclude] `json:"include"`
	// A flag to indicate if consecutive tables with the same number of columns should
	// be merged. Defaults to False.
	MergeTables param.Field[bool] `json:"merge_tables"`
	// The mode to use for table output. Defaults to dynamic, which returns md for
	// simpler tables and html for more complex tables.
	TableOutputFormat param.Field[ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormat] `json:"table_output_format"`
}

func (r ParseRunParamsBodyAsyncParseConfigFormatting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ParseRunParamsBodyAsyncParseConfigFormattingInclude string

const (
	ParseRunParamsBodyAsyncParseConfigFormattingIncludeChangeTracking ParseRunParamsBodyAsyncParseConfigFormattingInclude = "change_tracking"
	ParseRunParamsBodyAsyncParseConfigFormattingIncludeHighlight      ParseRunParamsBodyAsyncParseConfigFormattingInclude = "highlight"
	ParseRunParamsBodyAsyncParseConfigFormattingIncludeComments       ParseRunParamsBodyAsyncParseConfigFormattingInclude = "comments"
)

func (r ParseRunParamsBodyAsyncParseConfigFormattingInclude) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyAsyncParseConfigFormattingIncludeChangeTracking, ParseRunParamsBodyAsyncParseConfigFormattingIncludeHighlight, ParseRunParamsBodyAsyncParseConfigFormattingIncludeComments:
		return true
	}
	return false
}

// The mode to use for table output. Defaults to dynamic, which returns md for
// simpler tables and html for more complex tables.
type ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormat string

const (
	ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormatHTML     ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormat = "html"
	ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormatJson     ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormat = "json"
	ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormatMd       ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormat = "md"
	ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormatJsonbbox ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormat = "jsonbbox"
	ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormatDynamic  ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormat = "dynamic"
	ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormatCsv      ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormat = "csv"
)

func (r ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormat) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormatHTML, ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormatJson, ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormatMd, ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormatJsonbbox, ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormatDynamic, ParseRunParamsBodyAsyncParseConfigFormattingTableOutputFormatCsv:
		return true
	}
	return false
}

type ParseRunParamsBodyAsyncParseConfigRetrieval struct {
	Chunking param.Field[ParseRunParamsBodyAsyncParseConfigRetrievalChunking] `json:"chunking"`
	// If True, use embedding optimized mode. Defaults to False.
	EmbeddingOptimized param.Field[bool] `json:"embedding_optimized"`
	// A list of block types to filter out from 'content' and 'embed' fields. By
	// default, no blocks are filtered.
	FilterBlocks param.Field[[]ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlock] `json:"filter_blocks"`
}

func (r ParseRunParamsBodyAsyncParseConfigRetrieval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ParseRunParamsBodyAsyncParseConfigRetrievalChunking struct {
	// Choose how to partition chunks. Variable mode chunks by character length and
	// visual context. Section mode chunks by section headers. Page mode chunks
	// according to pages. Page sections mode chunks first by page, then by sections
	// within each page. Disabled returns one single chunk.
	ChunkMode param.Field[ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkMode] `json:"chunk_mode"`
	// The approximate size of chunks (in characters) that the document will be split
	// into. Defaults to null, in which case the chunk size is variable between 250 -
	// 1500 characters.
	ChunkSize param.Field[int64] `json:"chunk_size"`
}

func (r ParseRunParamsBodyAsyncParseConfigRetrievalChunking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Choose how to partition chunks. Variable mode chunks by character length and
// visual context. Section mode chunks by section headers. Page mode chunks
// according to pages. Page sections mode chunks first by page, then by sections
// within each page. Disabled returns one single chunk.
type ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkMode string

const (
	ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkModeVariable     ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkMode = "variable"
	ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkModeSection      ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkMode = "section"
	ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkModePage         ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkMode = "page"
	ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkModeDisabled     ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkMode = "disabled"
	ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkModeBlock        ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkMode = "block"
	ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkModePageSections ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkMode = "page_sections"
)

func (r ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkMode) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkModeVariable, ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkModeSection, ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkModePage, ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkModeDisabled, ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkModeBlock, ParseRunParamsBodyAsyncParseConfigRetrievalChunkingChunkModePageSections:
		return true
	}
	return false
}

type ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlock string

const (
	ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockHeader        ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlock = "Header"
	ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockFooter        ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlock = "Footer"
	ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockTitle         ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlock = "Title"
	ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockSectionHeader ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlock = "Section Header"
	ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockPageNumber    ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlock = "Page Number"
	ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockListItem      ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlock = "List Item"
	ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockFigure        ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlock = "Figure"
	ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockTable         ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlock = "Table"
	ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockKeyValue      ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlock = "Key Value"
	ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockText          ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlock = "Text"
	ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockComment       ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlock = "Comment"
	ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockSignature     ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlock = "Signature"
)

func (r ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlock) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockHeader, ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockFooter, ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockTitle, ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockSectionHeader, ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockPageNumber, ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockListItem, ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockFigure, ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockTable, ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockKeyValue, ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockText, ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockComment, ParseRunParamsBodyAsyncParseConfigRetrievalFilterBlockSignature:
		return true
	}
	return false
}

type ParseRunParamsBodyAsyncParseConfigSettings struct {
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
	OcrSystem param.Field[ParseRunParamsBodyAsyncParseConfigSettingsOcrSystem] `json:"ocr_system"`
	// The page range to process (1-indexed). By default, the entire document is
	// processed.
	PageRange param.Field[ParseRunParamsBodyAsyncParseConfigSettingsPageRangeUnion] `json:"page_range"`
	// If True, persist the results indefinitely. Defaults to False.
	PersistResults param.Field[bool] `json:"persist_results"`
	// Whether to return images for the specified block types. By default, no images
	// are returned.
	ReturnImages param.Field[[]ParseRunParamsBodyAsyncParseConfigSettingsReturnImage] `json:"return_images"`
	// If True, return OCR data in the result. Defaults to False.
	ReturnOcrData param.Field[bool] `json:"return_ocr_data"`
	// The timeout for the job in seconds. Defaults to 900.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r ParseRunParamsBodyAsyncParseConfigSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Standard is our best multilingual OCR system. Legacy only supports germanic
// languages and is available for backwards compatibility.
type ParseRunParamsBodyAsyncParseConfigSettingsOcrSystem string

const (
	ParseRunParamsBodyAsyncParseConfigSettingsOcrSystemStandard ParseRunParamsBodyAsyncParseConfigSettingsOcrSystem = "standard"
	ParseRunParamsBodyAsyncParseConfigSettingsOcrSystemLegacy   ParseRunParamsBodyAsyncParseConfigSettingsOcrSystem = "legacy"
)

func (r ParseRunParamsBodyAsyncParseConfigSettingsOcrSystem) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyAsyncParseConfigSettingsOcrSystemStandard, ParseRunParamsBodyAsyncParseConfigSettingsOcrSystemLegacy:
		return true
	}
	return false
}

// The page range to process (1-indexed). By default, the entire document is
// processed.
//
// Satisfied by [shared.PageRangeParam],
// [ParseRunParamsBodyAsyncParseConfigSettingsPageRangeArray],
// [ParseRunParamsBodyAsyncParseConfigSettingsPageRangeArray].
type ParseRunParamsBodyAsyncParseConfigSettingsPageRangeUnion interface {
	ImplementsParseRunParamsBodyAsyncParseConfigSettingsPageRangeUnion()
}

type ParseRunParamsBodyAsyncParseConfigSettingsPageRangeArray []shared.PageRangeParam

func (r ParseRunParamsBodyAsyncParseConfigSettingsPageRangeArray) ImplementsParseRunParamsBodyAsyncParseConfigSettingsPageRangeUnion() {
}

type ParseRunParamsBodyAsyncParseConfigSettingsReturnImage string

const (
	ParseRunParamsBodyAsyncParseConfigSettingsReturnImageFigure ParseRunParamsBodyAsyncParseConfigSettingsReturnImage = "figure"
	ParseRunParamsBodyAsyncParseConfigSettingsReturnImageTable  ParseRunParamsBodyAsyncParseConfigSettingsReturnImage = "table"
)

func (r ParseRunParamsBodyAsyncParseConfigSettingsReturnImage) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyAsyncParseConfigSettingsReturnImageFigure, ParseRunParamsBodyAsyncParseConfigSettingsReturnImageTable:
		return true
	}
	return false
}

type ParseRunParamsBodyAsyncParseConfigSpreadsheet struct {
	// In a spreadsheet with different tables inside, we enable splitting up the tables
	// by default. Accurate mode applies more powerful models for superior accuracy, at
	// 5× the default per-cell rate. Disabling will register as one large table.
	Clustering param.Field[ParseRunParamsBodyAsyncParseConfigSpreadsheetClustering] `json:"clustering"`
	// Whether to exclude hidden sheets, rows, or columns in the output.
	Exclude param.Field[[]ParseRunParamsBodyAsyncParseConfigSpreadsheetExclude] `json:"exclude"`
	// Whether to include cell color and formula information in the output.
	Include          param.Field[[]ParseRunParamsBodyAsyncParseConfigSpreadsheetInclude]        `json:"include"`
	SplitLargeTables param.Field[ParseRunParamsBodyAsyncParseConfigSpreadsheetSplitLargeTables] `json:"split_large_tables"`
}

func (r ParseRunParamsBodyAsyncParseConfigSpreadsheet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// In a spreadsheet with different tables inside, we enable splitting up the tables
// by default. Accurate mode applies more powerful models for superior accuracy, at
// 5× the default per-cell rate. Disabling will register as one large table.
type ParseRunParamsBodyAsyncParseConfigSpreadsheetClustering string

const (
	ParseRunParamsBodyAsyncParseConfigSpreadsheetClusteringAccurate ParseRunParamsBodyAsyncParseConfigSpreadsheetClustering = "accurate"
	ParseRunParamsBodyAsyncParseConfigSpreadsheetClusteringFast     ParseRunParamsBodyAsyncParseConfigSpreadsheetClustering = "fast"
	ParseRunParamsBodyAsyncParseConfigSpreadsheetClusteringDisabled ParseRunParamsBodyAsyncParseConfigSpreadsheetClustering = "disabled"
)

func (r ParseRunParamsBodyAsyncParseConfigSpreadsheetClustering) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyAsyncParseConfigSpreadsheetClusteringAccurate, ParseRunParamsBodyAsyncParseConfigSpreadsheetClusteringFast, ParseRunParamsBodyAsyncParseConfigSpreadsheetClusteringDisabled:
		return true
	}
	return false
}

type ParseRunParamsBodyAsyncParseConfigSpreadsheetExclude string

const (
	ParseRunParamsBodyAsyncParseConfigSpreadsheetExcludeHiddenSheets ParseRunParamsBodyAsyncParseConfigSpreadsheetExclude = "hidden_sheets"
	ParseRunParamsBodyAsyncParseConfigSpreadsheetExcludeHiddenRows   ParseRunParamsBodyAsyncParseConfigSpreadsheetExclude = "hidden_rows"
	ParseRunParamsBodyAsyncParseConfigSpreadsheetExcludeHiddenCols   ParseRunParamsBodyAsyncParseConfigSpreadsheetExclude = "hidden_cols"
)

func (r ParseRunParamsBodyAsyncParseConfigSpreadsheetExclude) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyAsyncParseConfigSpreadsheetExcludeHiddenSheets, ParseRunParamsBodyAsyncParseConfigSpreadsheetExcludeHiddenRows, ParseRunParamsBodyAsyncParseConfigSpreadsheetExcludeHiddenCols:
		return true
	}
	return false
}

type ParseRunParamsBodyAsyncParseConfigSpreadsheetInclude string

const (
	ParseRunParamsBodyAsyncParseConfigSpreadsheetIncludeCellColors ParseRunParamsBodyAsyncParseConfigSpreadsheetInclude = "cell_colors"
	ParseRunParamsBodyAsyncParseConfigSpreadsheetIncludeFormula    ParseRunParamsBodyAsyncParseConfigSpreadsheetInclude = "formula"
)

func (r ParseRunParamsBodyAsyncParseConfigSpreadsheetInclude) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyAsyncParseConfigSpreadsheetIncludeCellColors, ParseRunParamsBodyAsyncParseConfigSpreadsheetIncludeFormula:
		return true
	}
	return false
}

type ParseRunParamsBodyAsyncParseConfigSpreadsheetSplitLargeTables struct {
	// If True, split large tables into smaller tables. Defaults to True.
	Enabled param.Field[bool] `json:"enabled"`
	// The size of the tables to split into. Defaults to 50.
	Size param.Field[int64] `json:"size"`
}

func (r ParseRunParamsBodyAsyncParseConfigSpreadsheetSplitLargeTables) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ParseRunJobParams struct {
	// The URL of the document to be processed. You can provide one of the
	// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
	// prefixed URL obtained from the /upload endpoint after directly uploading a
	// document 4. A jobid:// prefixed URL obtained from a previous /parse invocation
	Input param.Field[ParseRunJobParamsInputUnion] `json:"input,required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async       param.Field[ParseRunJobParamsAsync]       `json:"async"`
	Enhance     param.Field[ParseRunJobParamsEnhance]     `json:"enhance"`
	Formatting  param.Field[ParseRunJobParamsFormatting]  `json:"formatting"`
	Retrieval   param.Field[ParseRunJobParamsRetrieval]   `json:"retrieval"`
	Settings    param.Field[ParseRunJobParamsSettings]    `json:"settings"`
	Spreadsheet param.Field[ParseRunJobParamsSpreadsheet] `json:"spreadsheet"`
}

func (r ParseRunJobParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The URL of the document to be processed. You can provide one of the
// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
// prefixed URL obtained from the /upload endpoint after directly uploading a
// document 4. A jobid:// prefixed URL obtained from a previous /parse invocation
//
// Satisfied by [shared.UnionString], [shared.UploadParam].
type ParseRunJobParamsInputUnion interface {
	ImplementsParseRunJobParamsInputUnion()
}

// The configuration options for asynchronous processing (default synchronous).
type ParseRunJobParamsAsync struct {
	// JSON metadata included in webhook request body. Defaults to None.
	Metadata param.Field[interface{}] `json:"metadata"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool] `json:"priority"`
	// The webhook configuration for the asynchronous processing.
	Webhook param.Field[ParseRunJobParamsAsyncWebhookUnion] `json:"webhook"`
}

func (r ParseRunJobParamsAsync) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The webhook configuration for the asynchronous processing.
type ParseRunJobParamsAsyncWebhook struct {
	Channels param.Field[interface{}]                       `json:"channels"`
	Mode     param.Field[ParseRunJobParamsAsyncWebhookMode] `json:"mode"`
	URL      param.Field[string]                            `json:"url"`
}

func (r ParseRunJobParamsAsyncWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunJobParamsAsyncWebhook) implementsParseRunJobParamsAsyncWebhookUnion() {}

// The webhook configuration for the asynchronous processing.
//
// Satisfied by [ParseRunJobParamsAsyncWebhookSvixWebhookConfig],
// [ParseRunJobParamsAsyncWebhookDirectWebhookConfig],
// [ParseRunJobParamsAsyncWebhook].
type ParseRunJobParamsAsyncWebhookUnion interface {
	implementsParseRunJobParamsAsyncWebhookUnion()
}

type ParseRunJobParamsAsyncWebhookSvixWebhookConfig struct {
	// A list of Svix channels the message will be delivered down, omit to send to all
	// channels.
	Channels param.Field[[]string]                                           `json:"channels"`
	Mode     param.Field[ParseRunJobParamsAsyncWebhookSvixWebhookConfigMode] `json:"mode"`
}

func (r ParseRunJobParamsAsyncWebhookSvixWebhookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunJobParamsAsyncWebhookSvixWebhookConfig) implementsParseRunJobParamsAsyncWebhookUnion() {
}

type ParseRunJobParamsAsyncWebhookSvixWebhookConfigMode string

const (
	ParseRunJobParamsAsyncWebhookSvixWebhookConfigModeSvix ParseRunJobParamsAsyncWebhookSvixWebhookConfigMode = "svix"
)

func (r ParseRunJobParamsAsyncWebhookSvixWebhookConfigMode) IsKnown() bool {
	switch r {
	case ParseRunJobParamsAsyncWebhookSvixWebhookConfigModeSvix:
		return true
	}
	return false
}

type ParseRunJobParamsAsyncWebhookDirectWebhookConfig struct {
	URL  param.Field[string]                                               `json:"url,required"`
	Mode param.Field[ParseRunJobParamsAsyncWebhookDirectWebhookConfigMode] `json:"mode"`
}

func (r ParseRunJobParamsAsyncWebhookDirectWebhookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunJobParamsAsyncWebhookDirectWebhookConfig) implementsParseRunJobParamsAsyncWebhookUnion() {
}

type ParseRunJobParamsAsyncWebhookDirectWebhookConfigMode string

const (
	ParseRunJobParamsAsyncWebhookDirectWebhookConfigModeDirect ParseRunJobParamsAsyncWebhookDirectWebhookConfigMode = "direct"
)

func (r ParseRunJobParamsAsyncWebhookDirectWebhookConfigMode) IsKnown() bool {
	switch r {
	case ParseRunJobParamsAsyncWebhookDirectWebhookConfigModeDirect:
		return true
	}
	return false
}

type ParseRunJobParamsAsyncWebhookMode string

const (
	ParseRunJobParamsAsyncWebhookModeSvix   ParseRunJobParamsAsyncWebhookMode = "svix"
	ParseRunJobParamsAsyncWebhookModeDirect ParseRunJobParamsAsyncWebhookMode = "direct"
)

func (r ParseRunJobParamsAsyncWebhookMode) IsKnown() bool {
	switch r {
	case ParseRunJobParamsAsyncWebhookModeSvix, ParseRunJobParamsAsyncWebhookModeDirect:
		return true
	}
	return false
}

type ParseRunJobParamsEnhance struct {
	// Agentic uses vision language models to enhance the accuracy of the output of
	// different types of extraction. This will incur a cost and latency increase.
	Agentic param.Field[[]ParseRunJobParamsEnhanceAgenticUnion] `json:"agentic"`
	// If True, summarize figures using a small vision language model. Defaults to
	// True.
	SummarizeFigures param.Field[bool] `json:"summarize_figures"`
}

func (r ParseRunJobParamsEnhance) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ParseRunJobParamsEnhanceAgentic struct {
	Scope param.Field[ParseRunJobParamsEnhanceAgenticScope] `json:"scope,required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ParseRunJobParamsEnhanceAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunJobParamsEnhanceAgentic) implementsParseRunJobParamsEnhanceAgenticUnion() {}

// Satisfied by [ParseRunJobParamsEnhanceAgenticTableAgentic],
// [ParseRunJobParamsEnhanceAgenticFigureAgentic],
// [ParseRunJobParamsEnhanceAgenticTextAgentic], [ParseRunJobParamsEnhanceAgentic].
type ParseRunJobParamsEnhanceAgenticUnion interface {
	implementsParseRunJobParamsEnhanceAgenticUnion()
}

type ParseRunJobParamsEnhanceAgenticTableAgentic struct {
	Scope param.Field[ParseRunJobParamsEnhanceAgenticTableAgenticScope] `json:"scope,required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ParseRunJobParamsEnhanceAgenticTableAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunJobParamsEnhanceAgenticTableAgentic) implementsParseRunJobParamsEnhanceAgenticUnion() {
}

type ParseRunJobParamsEnhanceAgenticTableAgenticScope string

const (
	ParseRunJobParamsEnhanceAgenticTableAgenticScopeTable ParseRunJobParamsEnhanceAgenticTableAgenticScope = "table"
)

func (r ParseRunJobParamsEnhanceAgenticTableAgenticScope) IsKnown() bool {
	switch r {
	case ParseRunJobParamsEnhanceAgenticTableAgenticScopeTable:
		return true
	}
	return false
}

type ParseRunJobParamsEnhanceAgenticFigureAgentic struct {
	Scope param.Field[ParseRunJobParamsEnhanceAgenticFigureAgenticScope] `json:"scope,required"`
	// Custom prompt for figure agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ParseRunJobParamsEnhanceAgenticFigureAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunJobParamsEnhanceAgenticFigureAgentic) implementsParseRunJobParamsEnhanceAgenticUnion() {
}

type ParseRunJobParamsEnhanceAgenticFigureAgenticScope string

const (
	ParseRunJobParamsEnhanceAgenticFigureAgenticScopeFigure ParseRunJobParamsEnhanceAgenticFigureAgenticScope = "figure"
)

func (r ParseRunJobParamsEnhanceAgenticFigureAgenticScope) IsKnown() bool {
	switch r {
	case ParseRunJobParamsEnhanceAgenticFigureAgenticScopeFigure:
		return true
	}
	return false
}

type ParseRunJobParamsEnhanceAgenticTextAgentic struct {
	Scope param.Field[ParseRunJobParamsEnhanceAgenticTextAgenticScope] `json:"scope,required"`
}

func (r ParseRunJobParamsEnhanceAgenticTextAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunJobParamsEnhanceAgenticTextAgentic) implementsParseRunJobParamsEnhanceAgenticUnion() {
}

type ParseRunJobParamsEnhanceAgenticTextAgenticScope string

const (
	ParseRunJobParamsEnhanceAgenticTextAgenticScopeText ParseRunJobParamsEnhanceAgenticTextAgenticScope = "text"
)

func (r ParseRunJobParamsEnhanceAgenticTextAgenticScope) IsKnown() bool {
	switch r {
	case ParseRunJobParamsEnhanceAgenticTextAgenticScopeText:
		return true
	}
	return false
}

type ParseRunJobParamsEnhanceAgenticScope string

const (
	ParseRunJobParamsEnhanceAgenticScopeTable  ParseRunJobParamsEnhanceAgenticScope = "table"
	ParseRunJobParamsEnhanceAgenticScopeFigure ParseRunJobParamsEnhanceAgenticScope = "figure"
	ParseRunJobParamsEnhanceAgenticScopeText   ParseRunJobParamsEnhanceAgenticScope = "text"
)

func (r ParseRunJobParamsEnhanceAgenticScope) IsKnown() bool {
	switch r {
	case ParseRunJobParamsEnhanceAgenticScopeTable, ParseRunJobParamsEnhanceAgenticScopeFigure, ParseRunJobParamsEnhanceAgenticScopeText:
		return true
	}
	return false
}

type ParseRunJobParamsFormatting struct {
	// If True, add page markers to the output. Defaults to False. Useful for
	// extracting data with page specific information.
	AddPageMarkers param.Field[bool] `json:"add_page_markers"`
	// A list of formatting to include in the output. [insert description of each
	// option here later]
	Include param.Field[[]ParseRunJobParamsFormattingInclude] `json:"include"`
	// A flag to indicate if consecutive tables with the same number of columns should
	// be merged. Defaults to False.
	MergeTables param.Field[bool] `json:"merge_tables"`
	// The mode to use for table output. Defaults to dynamic, which returns md for
	// simpler tables and html for more complex tables.
	TableOutputFormat param.Field[ParseRunJobParamsFormattingTableOutputFormat] `json:"table_output_format"`
}

func (r ParseRunJobParamsFormatting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ParseRunJobParamsFormattingInclude string

const (
	ParseRunJobParamsFormattingIncludeChangeTracking ParseRunJobParamsFormattingInclude = "change_tracking"
	ParseRunJobParamsFormattingIncludeHighlight      ParseRunJobParamsFormattingInclude = "highlight"
	ParseRunJobParamsFormattingIncludeComments       ParseRunJobParamsFormattingInclude = "comments"
)

func (r ParseRunJobParamsFormattingInclude) IsKnown() bool {
	switch r {
	case ParseRunJobParamsFormattingIncludeChangeTracking, ParseRunJobParamsFormattingIncludeHighlight, ParseRunJobParamsFormattingIncludeComments:
		return true
	}
	return false
}

// The mode to use for table output. Defaults to dynamic, which returns md for
// simpler tables and html for more complex tables.
type ParseRunJobParamsFormattingTableOutputFormat string

const (
	ParseRunJobParamsFormattingTableOutputFormatHTML     ParseRunJobParamsFormattingTableOutputFormat = "html"
	ParseRunJobParamsFormattingTableOutputFormatJson     ParseRunJobParamsFormattingTableOutputFormat = "json"
	ParseRunJobParamsFormattingTableOutputFormatMd       ParseRunJobParamsFormattingTableOutputFormat = "md"
	ParseRunJobParamsFormattingTableOutputFormatJsonbbox ParseRunJobParamsFormattingTableOutputFormat = "jsonbbox"
	ParseRunJobParamsFormattingTableOutputFormatDynamic  ParseRunJobParamsFormattingTableOutputFormat = "dynamic"
	ParseRunJobParamsFormattingTableOutputFormatCsv      ParseRunJobParamsFormattingTableOutputFormat = "csv"
)

func (r ParseRunJobParamsFormattingTableOutputFormat) IsKnown() bool {
	switch r {
	case ParseRunJobParamsFormattingTableOutputFormatHTML, ParseRunJobParamsFormattingTableOutputFormatJson, ParseRunJobParamsFormattingTableOutputFormatMd, ParseRunJobParamsFormattingTableOutputFormatJsonbbox, ParseRunJobParamsFormattingTableOutputFormatDynamic, ParseRunJobParamsFormattingTableOutputFormatCsv:
		return true
	}
	return false
}

type ParseRunJobParamsRetrieval struct {
	Chunking param.Field[ParseRunJobParamsRetrievalChunking] `json:"chunking"`
	// If True, use embedding optimized mode. Defaults to False.
	EmbeddingOptimized param.Field[bool] `json:"embedding_optimized"`
	// A list of block types to filter out from 'content' and 'embed' fields. By
	// default, no blocks are filtered.
	FilterBlocks param.Field[[]ParseRunJobParamsRetrievalFilterBlock] `json:"filter_blocks"`
}

func (r ParseRunJobParamsRetrieval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ParseRunJobParamsRetrievalChunking struct {
	// Choose how to partition chunks. Variable mode chunks by character length and
	// visual context. Section mode chunks by section headers. Page mode chunks
	// according to pages. Page sections mode chunks first by page, then by sections
	// within each page. Disabled returns one single chunk.
	ChunkMode param.Field[ParseRunJobParamsRetrievalChunkingChunkMode] `json:"chunk_mode"`
	// The approximate size of chunks (in characters) that the document will be split
	// into. Defaults to null, in which case the chunk size is variable between 250 -
	// 1500 characters.
	ChunkSize param.Field[int64] `json:"chunk_size"`
}

func (r ParseRunJobParamsRetrievalChunking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Choose how to partition chunks. Variable mode chunks by character length and
// visual context. Section mode chunks by section headers. Page mode chunks
// according to pages. Page sections mode chunks first by page, then by sections
// within each page. Disabled returns one single chunk.
type ParseRunJobParamsRetrievalChunkingChunkMode string

const (
	ParseRunJobParamsRetrievalChunkingChunkModeVariable     ParseRunJobParamsRetrievalChunkingChunkMode = "variable"
	ParseRunJobParamsRetrievalChunkingChunkModeSection      ParseRunJobParamsRetrievalChunkingChunkMode = "section"
	ParseRunJobParamsRetrievalChunkingChunkModePage         ParseRunJobParamsRetrievalChunkingChunkMode = "page"
	ParseRunJobParamsRetrievalChunkingChunkModeDisabled     ParseRunJobParamsRetrievalChunkingChunkMode = "disabled"
	ParseRunJobParamsRetrievalChunkingChunkModeBlock        ParseRunJobParamsRetrievalChunkingChunkMode = "block"
	ParseRunJobParamsRetrievalChunkingChunkModePageSections ParseRunJobParamsRetrievalChunkingChunkMode = "page_sections"
)

func (r ParseRunJobParamsRetrievalChunkingChunkMode) IsKnown() bool {
	switch r {
	case ParseRunJobParamsRetrievalChunkingChunkModeVariable, ParseRunJobParamsRetrievalChunkingChunkModeSection, ParseRunJobParamsRetrievalChunkingChunkModePage, ParseRunJobParamsRetrievalChunkingChunkModeDisabled, ParseRunJobParamsRetrievalChunkingChunkModeBlock, ParseRunJobParamsRetrievalChunkingChunkModePageSections:
		return true
	}
	return false
}

type ParseRunJobParamsRetrievalFilterBlock string

const (
	ParseRunJobParamsRetrievalFilterBlockHeader        ParseRunJobParamsRetrievalFilterBlock = "Header"
	ParseRunJobParamsRetrievalFilterBlockFooter        ParseRunJobParamsRetrievalFilterBlock = "Footer"
	ParseRunJobParamsRetrievalFilterBlockTitle         ParseRunJobParamsRetrievalFilterBlock = "Title"
	ParseRunJobParamsRetrievalFilterBlockSectionHeader ParseRunJobParamsRetrievalFilterBlock = "Section Header"
	ParseRunJobParamsRetrievalFilterBlockPageNumber    ParseRunJobParamsRetrievalFilterBlock = "Page Number"
	ParseRunJobParamsRetrievalFilterBlockListItem      ParseRunJobParamsRetrievalFilterBlock = "List Item"
	ParseRunJobParamsRetrievalFilterBlockFigure        ParseRunJobParamsRetrievalFilterBlock = "Figure"
	ParseRunJobParamsRetrievalFilterBlockTable         ParseRunJobParamsRetrievalFilterBlock = "Table"
	ParseRunJobParamsRetrievalFilterBlockKeyValue      ParseRunJobParamsRetrievalFilterBlock = "Key Value"
	ParseRunJobParamsRetrievalFilterBlockText          ParseRunJobParamsRetrievalFilterBlock = "Text"
	ParseRunJobParamsRetrievalFilterBlockComment       ParseRunJobParamsRetrievalFilterBlock = "Comment"
	ParseRunJobParamsRetrievalFilterBlockSignature     ParseRunJobParamsRetrievalFilterBlock = "Signature"
)

func (r ParseRunJobParamsRetrievalFilterBlock) IsKnown() bool {
	switch r {
	case ParseRunJobParamsRetrievalFilterBlockHeader, ParseRunJobParamsRetrievalFilterBlockFooter, ParseRunJobParamsRetrievalFilterBlockTitle, ParseRunJobParamsRetrievalFilterBlockSectionHeader, ParseRunJobParamsRetrievalFilterBlockPageNumber, ParseRunJobParamsRetrievalFilterBlockListItem, ParseRunJobParamsRetrievalFilterBlockFigure, ParseRunJobParamsRetrievalFilterBlockTable, ParseRunJobParamsRetrievalFilterBlockKeyValue, ParseRunJobParamsRetrievalFilterBlockText, ParseRunJobParamsRetrievalFilterBlockComment, ParseRunJobParamsRetrievalFilterBlockSignature:
		return true
	}
	return false
}

type ParseRunJobParamsSettings struct {
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
	OcrSystem param.Field[ParseRunJobParamsSettingsOcrSystem] `json:"ocr_system"`
	// The page range to process (1-indexed). By default, the entire document is
	// processed.
	PageRange param.Field[ParseRunJobParamsSettingsPageRangeUnion] `json:"page_range"`
	// If True, persist the results indefinitely. Defaults to False.
	PersistResults param.Field[bool] `json:"persist_results"`
	// Whether to return images for the specified block types. By default, no images
	// are returned.
	ReturnImages param.Field[[]ParseRunJobParamsSettingsReturnImage] `json:"return_images"`
	// If True, return OCR data in the result. Defaults to False.
	ReturnOcrData param.Field[bool] `json:"return_ocr_data"`
	// The timeout for the job in seconds. Defaults to 900.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r ParseRunJobParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Standard is our best multilingual OCR system. Legacy only supports germanic
// languages and is available for backwards compatibility.
type ParseRunJobParamsSettingsOcrSystem string

const (
	ParseRunJobParamsSettingsOcrSystemStandard ParseRunJobParamsSettingsOcrSystem = "standard"
	ParseRunJobParamsSettingsOcrSystemLegacy   ParseRunJobParamsSettingsOcrSystem = "legacy"
)

func (r ParseRunJobParamsSettingsOcrSystem) IsKnown() bool {
	switch r {
	case ParseRunJobParamsSettingsOcrSystemStandard, ParseRunJobParamsSettingsOcrSystemLegacy:
		return true
	}
	return false
}

// The page range to process (1-indexed). By default, the entire document is
// processed.
//
// Satisfied by [shared.PageRangeParam], [ParseRunJobParamsSettingsPageRangeArray],
// [ParseRunJobParamsSettingsPageRangeArray].
type ParseRunJobParamsSettingsPageRangeUnion interface {
	ImplementsParseRunJobParamsSettingsPageRangeUnion()
}

type ParseRunJobParamsSettingsPageRangeArray []shared.PageRangeParam

func (r ParseRunJobParamsSettingsPageRangeArray) ImplementsParseRunJobParamsSettingsPageRangeUnion() {
}

type ParseRunJobParamsSettingsReturnImage string

const (
	ParseRunJobParamsSettingsReturnImageFigure ParseRunJobParamsSettingsReturnImage = "figure"
	ParseRunJobParamsSettingsReturnImageTable  ParseRunJobParamsSettingsReturnImage = "table"
)

func (r ParseRunJobParamsSettingsReturnImage) IsKnown() bool {
	switch r {
	case ParseRunJobParamsSettingsReturnImageFigure, ParseRunJobParamsSettingsReturnImageTable:
		return true
	}
	return false
}

type ParseRunJobParamsSpreadsheet struct {
	// In a spreadsheet with different tables inside, we enable splitting up the tables
	// by default. Accurate mode applies more powerful models for superior accuracy, at
	// 5× the default per-cell rate. Disabling will register as one large table.
	Clustering param.Field[ParseRunJobParamsSpreadsheetClustering] `json:"clustering"`
	// Whether to exclude hidden sheets, rows, or columns in the output.
	Exclude param.Field[[]ParseRunJobParamsSpreadsheetExclude] `json:"exclude"`
	// Whether to include cell color and formula information in the output.
	Include          param.Field[[]ParseRunJobParamsSpreadsheetInclude]        `json:"include"`
	SplitLargeTables param.Field[ParseRunJobParamsSpreadsheetSplitLargeTables] `json:"split_large_tables"`
}

func (r ParseRunJobParamsSpreadsheet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// In a spreadsheet with different tables inside, we enable splitting up the tables
// by default. Accurate mode applies more powerful models for superior accuracy, at
// 5× the default per-cell rate. Disabling will register as one large table.
type ParseRunJobParamsSpreadsheetClustering string

const (
	ParseRunJobParamsSpreadsheetClusteringAccurate ParseRunJobParamsSpreadsheetClustering = "accurate"
	ParseRunJobParamsSpreadsheetClusteringFast     ParseRunJobParamsSpreadsheetClustering = "fast"
	ParseRunJobParamsSpreadsheetClusteringDisabled ParseRunJobParamsSpreadsheetClustering = "disabled"
)

func (r ParseRunJobParamsSpreadsheetClustering) IsKnown() bool {
	switch r {
	case ParseRunJobParamsSpreadsheetClusteringAccurate, ParseRunJobParamsSpreadsheetClusteringFast, ParseRunJobParamsSpreadsheetClusteringDisabled:
		return true
	}
	return false
}

type ParseRunJobParamsSpreadsheetExclude string

const (
	ParseRunJobParamsSpreadsheetExcludeHiddenSheets ParseRunJobParamsSpreadsheetExclude = "hidden_sheets"
	ParseRunJobParamsSpreadsheetExcludeHiddenRows   ParseRunJobParamsSpreadsheetExclude = "hidden_rows"
	ParseRunJobParamsSpreadsheetExcludeHiddenCols   ParseRunJobParamsSpreadsheetExclude = "hidden_cols"
)

func (r ParseRunJobParamsSpreadsheetExclude) IsKnown() bool {
	switch r {
	case ParseRunJobParamsSpreadsheetExcludeHiddenSheets, ParseRunJobParamsSpreadsheetExcludeHiddenRows, ParseRunJobParamsSpreadsheetExcludeHiddenCols:
		return true
	}
	return false
}

type ParseRunJobParamsSpreadsheetInclude string

const (
	ParseRunJobParamsSpreadsheetIncludeCellColors ParseRunJobParamsSpreadsheetInclude = "cell_colors"
	ParseRunJobParamsSpreadsheetIncludeFormula    ParseRunJobParamsSpreadsheetInclude = "formula"
)

func (r ParseRunJobParamsSpreadsheetInclude) IsKnown() bool {
	switch r {
	case ParseRunJobParamsSpreadsheetIncludeCellColors, ParseRunJobParamsSpreadsheetIncludeFormula:
		return true
	}
	return false
}

type ParseRunJobParamsSpreadsheetSplitLargeTables struct {
	// If True, split large tables into smaller tables. Defaults to True.
	Enabled param.Field[bool] `json:"enabled"`
	// The size of the tables to split into. Defaults to 50.
	Size param.Field[int64] `json:"size"`
}

func (r ParseRunJobParamsSpreadsheetSplitLargeTables) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
