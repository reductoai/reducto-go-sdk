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
	return
}

// Extract Async
func (r *ExtractService) RunJob(ctx context.Context, body ExtractRunJobParams, opts ...option.RequestOption) (res *ExtractRunJobResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "extract_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ExtractRunResponse struct {
	JobID string `json:"job_id,nullable"`
	// This field can have the runtime type of [[]interface{}].
	Result interface{} `json:"result"`
	// The link to the studio pipeline for the document.
	StudioLink string `json:"studio_link,nullable"`
	// This field can have the runtime type of
	// [ExtractRunResponseV3ExtractResponseUsage].
	Usage interface{}            `json:"usage"`
	JSON  extractRunResponseJSON `json:"-"`
	union ExtractRunResponseUnion
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
// Possible runtime types of the union are [ExtractRunResponseV3ExtractResponse],
// [ExtractRunResponseAsyncExtractResponse].
func (r ExtractRunResponse) AsUnion() ExtractRunResponseUnion {
	return r.union
}

// Union satisfied by [ExtractRunResponseV3ExtractResponse] or
// [ExtractRunResponseAsyncExtractResponse].
type ExtractRunResponseUnion interface {
	implementsExtractRunResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ExtractRunResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExtractRunResponseV3ExtractResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExtractRunResponseAsyncExtractResponse{}),
		},
	)
}

type ExtractRunResponseV3ExtractResponse struct {
	// The extracted response in your provided schema. This is a list of dictionaries.
	// If disable_chunking is True (default), then it will be a list of length one.
	Result []interface{}                            `json:"result,required"`
	Usage  ExtractRunResponseV3ExtractResponseUsage `json:"usage,required"`
	JobID  string                                   `json:"job_id,nullable"`
	// The link to the studio pipeline for the document.
	StudioLink string                                  `json:"studio_link,nullable"`
	JSON       extractRunResponseV3ExtractResponseJSON `json:"-"`
}

// extractRunResponseV3ExtractResponseJSON contains the JSON metadata for the
// struct [ExtractRunResponseV3ExtractResponse]
type extractRunResponseV3ExtractResponseJSON struct {
	Result      apijson.Field
	Usage       apijson.Field
	JobID       apijson.Field
	StudioLink  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtractRunResponseV3ExtractResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extractRunResponseV3ExtractResponseJSON) RawJSON() string {
	return r.raw
}

func (r ExtractRunResponseV3ExtractResponse) implementsExtractRunResponse() {}

type ExtractRunResponseV3ExtractResponseUsage struct {
	NumFields int64                                        `json:"num_fields,required"`
	NumPages  int64                                        `json:"num_pages,required"`
	Credits   float64                                      `json:"credits,nullable"`
	JSON      extractRunResponseV3ExtractResponseUsageJSON `json:"-"`
}

// extractRunResponseV3ExtractResponseUsageJSON contains the JSON metadata for the
// struct [ExtractRunResponseV3ExtractResponseUsage]
type extractRunResponseV3ExtractResponseUsageJSON struct {
	NumFields   apijson.Field
	NumPages    apijson.Field
	Credits     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtractRunResponseV3ExtractResponseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extractRunResponseV3ExtractResponseUsageJSON) RawJSON() string {
	return r.raw
}

type ExtractRunResponseAsyncExtractResponse struct {
	JobID string                                     `json:"job_id,required"`
	JSON  extractRunResponseAsyncExtractResponseJSON `json:"-"`
}

// extractRunResponseAsyncExtractResponseJSON contains the JSON metadata for the
// struct [ExtractRunResponseAsyncExtractResponse]
type extractRunResponseAsyncExtractResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtractRunResponseAsyncExtractResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extractRunResponseAsyncExtractResponseJSON) RawJSON() string {
	return r.raw
}

func (r ExtractRunResponseAsyncExtractResponse) implementsExtractRunResponse() {}

type ExtractRunJobResponse struct {
	JobID string                    `json:"job_id,required"`
	JSON  extractRunJobResponseJSON `json:"-"`
}

// extractRunJobResponseJSON contains the JSON metadata for the struct
// [ExtractRunJobResponse]
type extractRunJobResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtractRunJobResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extractRunJobResponseJSON) RawJSON() string {
	return r.raw
}

type ExtractRunParams struct {
	Body ExtractRunParamsBodyUnion `json:"body,required"`
}

func (r ExtractRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ExtractRunParamsBody struct {
	Input        param.Field[interface{}] `json:"input,required"`
	Async        param.Field[interface{}] `json:"async"`
	Instructions param.Field[interface{}] `json:"instructions"`
	Parsing      param.Field[interface{}] `json:"parsing"`
	Settings     param.Field[interface{}] `json:"settings"`
}

func (r ExtractRunParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBody) implementsExtractRunParamsBodyUnion() {}

// Satisfied by [ExtractRunParamsBodySyncExtractConfig],
// [ExtractRunParamsBodyAsyncExtractConfig], [ExtractRunParamsBody].
type ExtractRunParamsBodyUnion interface {
	implementsExtractRunParamsBodyUnion()
}

type ExtractRunParamsBodySyncExtractConfig struct {
	// The URL of the document to be processed. You can provide one of the
	// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
	// prefixed URL obtained from the /upload endpoint after directly uploading a
	// document 4. A jobid:// prefixed URL obtained from a previous /parse invocation
	Input param.Field[ExtractRunParamsBodySyncExtractConfigInputUnion] `json:"input,required"`
	// The instructions to use for the extraction.
	Instructions param.Field[ExtractRunParamsBodySyncExtractConfigInstructions] `json:"instructions"`
	// The configuration options for parsing the document. If you are passing in a
	// jobid:// URL for the file, then this configuration will be ignored.
	Parsing param.Field[ExtractRunParamsBodySyncExtractConfigParsing] `json:"parsing"`
	// The settings to use for the extraction.
	Settings param.Field[ExtractRunParamsBodySyncExtractConfigSettings] `json:"settings"`
}

func (r ExtractRunParamsBodySyncExtractConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBodySyncExtractConfig) implementsExtractRunParamsBodyUnion() {}

// The URL of the document to be processed. You can provide one of the
// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
// prefixed URL obtained from the /upload endpoint after directly uploading a
// document 4. A jobid:// prefixed URL obtained from a previous /parse invocation
//
// Satisfied by [shared.UnionString], [shared.UploadParam].
type ExtractRunParamsBodySyncExtractConfigInputUnion interface {
	ImplementsExtractRunParamsBodySyncExtractConfigInputUnion()
}

// The instructions to use for the extraction.
type ExtractRunParamsBodySyncExtractConfigInstructions struct {
	// The JSON schema to use for the extraction.
	Schema param.Field[interface{}] `json:"schema"`
	// The system prompt to use for the extraction.
	SystemPrompt param.Field[string] `json:"system_prompt"`
}

func (r ExtractRunParamsBodySyncExtractConfigInstructions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration options for parsing the document. If you are passing in a
// jobid:// URL for the file, then this configuration will be ignored.
type ExtractRunParamsBodySyncExtractConfigParsing struct {
	Enhance     param.Field[ExtractRunParamsBodySyncExtractConfigParsingEnhance]     `json:"enhance"`
	Formatting  param.Field[ExtractRunParamsBodySyncExtractConfigParsingFormatting]  `json:"formatting"`
	Retrieval   param.Field[ExtractRunParamsBodySyncExtractConfigParsingRetrieval]   `json:"retrieval"`
	Settings    param.Field[ExtractRunParamsBodySyncExtractConfigParsingSettings]    `json:"settings"`
	Spreadsheet param.Field[ExtractRunParamsBodySyncExtractConfigParsingSpreadsheet] `json:"spreadsheet"`
}

func (r ExtractRunParamsBodySyncExtractConfigParsing) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunParamsBodySyncExtractConfigParsingEnhance struct {
	// Agentic uses vision language models to enhance the accuracy of the output of
	// different types of extraction. This will incur a cost and latency increase.
	Agentic param.Field[[]ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticUnion] `json:"agentic"`
	// If True, summarize figures using a small vision language model. Defaults to
	// True.
	SummarizeFigures param.Field[bool] `json:"summarize_figures"`
}

func (r ExtractRunParamsBodySyncExtractConfigParsingEnhance) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgentic struct {
	Scope param.Field[ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticScope] `json:"scope,required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgentic) implementsExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticUnion() {
}

// Satisfied by
// [ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTableAgentic],
// [ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticFigureAgentic],
// [ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTextAgentic],
// [ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgentic].
type ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticUnion interface {
	implementsExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticUnion()
}

type ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTableAgentic struct {
	Scope param.Field[ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTableAgenticScope] `json:"scope,required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTableAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTableAgentic) implementsExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticUnion() {
}

type ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTableAgenticScope string

const (
	ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTableAgenticScopeTable ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTableAgenticScope = "table"
)

func (r ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTableAgenticScope) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTableAgenticScopeTable:
		return true
	}
	return false
}

type ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticFigureAgentic struct {
	Scope param.Field[ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticFigureAgenticScope] `json:"scope,required"`
	// Custom prompt for figure agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticFigureAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticFigureAgentic) implementsExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticUnion() {
}

type ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticFigureAgenticScope string

const (
	ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticFigureAgenticScopeFigure ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticFigureAgenticScope = "figure"
)

func (r ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticFigureAgenticScope) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticFigureAgenticScopeFigure:
		return true
	}
	return false
}

type ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTextAgentic struct {
	Scope param.Field[ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTextAgenticScope] `json:"scope,required"`
}

func (r ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTextAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTextAgentic) implementsExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticUnion() {
}

type ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTextAgenticScope string

const (
	ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTextAgenticScopeText ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTextAgenticScope = "text"
)

func (r ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTextAgenticScope) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTextAgenticScopeText:
		return true
	}
	return false
}

type ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticScope string

const (
	ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticScopeTable  ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticScope = "table"
	ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticScopeFigure ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticScope = "figure"
	ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticScopeText   ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticScope = "text"
)

func (r ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticScope) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticScopeTable, ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticScopeFigure, ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticScopeText:
		return true
	}
	return false
}

type ExtractRunParamsBodySyncExtractConfigParsingFormatting struct {
	// If True, add page markers to the output. Defaults to False. Useful for
	// extracting data with page specific information.
	AddPageMarkers param.Field[bool] `json:"add_page_markers"`
	// A list of formatting to include in the output. [insert description of each
	// option here later]
	Include param.Field[[]ExtractRunParamsBodySyncExtractConfigParsingFormattingInclude] `json:"include"`
	// A flag to indicate if consecutive tables with the same number of columns should
	// be merged. Defaults to False.
	MergeTables param.Field[bool] `json:"merge_tables"`
	// The mode to use for table output. Defaults to dynamic, which returns md for
	// simpler tables and html for more complex tables.
	TableOutputFormat param.Field[ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormat] `json:"table_output_format"`
}

func (r ExtractRunParamsBodySyncExtractConfigParsingFormatting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunParamsBodySyncExtractConfigParsingFormattingInclude string

const (
	ExtractRunParamsBodySyncExtractConfigParsingFormattingIncludeChangeTracking ExtractRunParamsBodySyncExtractConfigParsingFormattingInclude = "change_tracking"
	ExtractRunParamsBodySyncExtractConfigParsingFormattingIncludeHighlight      ExtractRunParamsBodySyncExtractConfigParsingFormattingInclude = "highlight"
	ExtractRunParamsBodySyncExtractConfigParsingFormattingIncludeComments       ExtractRunParamsBodySyncExtractConfigParsingFormattingInclude = "comments"
)

func (r ExtractRunParamsBodySyncExtractConfigParsingFormattingInclude) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodySyncExtractConfigParsingFormattingIncludeChangeTracking, ExtractRunParamsBodySyncExtractConfigParsingFormattingIncludeHighlight, ExtractRunParamsBodySyncExtractConfigParsingFormattingIncludeComments:
		return true
	}
	return false
}

// The mode to use for table output. Defaults to dynamic, which returns md for
// simpler tables and html for more complex tables.
type ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormat string

const (
	ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormatHTML     ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormat = "html"
	ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormatJson     ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormat = "json"
	ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormatMd       ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormat = "md"
	ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormatJsonbbox ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormat = "jsonbbox"
	ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormatDynamic  ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormat = "dynamic"
	ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormatCsv      ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormat = "csv"
)

func (r ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormat) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormatHTML, ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormatJson, ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormatMd, ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormatJsonbbox, ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormatDynamic, ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormatCsv:
		return true
	}
	return false
}

type ExtractRunParamsBodySyncExtractConfigParsingRetrieval struct {
	Chunking param.Field[ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunking] `json:"chunking"`
	// If True, use embedding optimized mode. Defaults to False.
	EmbeddingOptimized param.Field[bool] `json:"embedding_optimized"`
	// A list of block types to filter out from 'content' and 'embed' fields. By
	// default, no blocks are filtered.
	FilterBlocks param.Field[[]ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlock] `json:"filter_blocks"`
}

func (r ExtractRunParamsBodySyncExtractConfigParsingRetrieval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunking struct {
	// Choose how to partition chunks. Variable mode chunks by character length and
	// visual context. Section mode chunks by section headers. Page mode chunks
	// according to pages. Page sections mode chunks first by page, then by sections
	// within each page. Disabled returns one single chunk.
	ChunkMode param.Field[ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkMode] `json:"chunk_mode"`
	// The approximate size of chunks (in characters) that the document will be split
	// into. Defaults to null, in which case the chunk size is variable between 250 -
	// 1500 characters.
	ChunkSize param.Field[int64] `json:"chunk_size"`
}

func (r ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Choose how to partition chunks. Variable mode chunks by character length and
// visual context. Section mode chunks by section headers. Page mode chunks
// according to pages. Page sections mode chunks first by page, then by sections
// within each page. Disabled returns one single chunk.
type ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkMode string

const (
	ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkModeVariable     ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkMode = "variable"
	ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkModeSection      ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkMode = "section"
	ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkModePage         ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkMode = "page"
	ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkModeDisabled     ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkMode = "disabled"
	ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkModeBlock        ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkMode = "block"
	ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkModePageSections ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkMode = "page_sections"
)

func (r ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkMode) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkModeVariable, ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkModeSection, ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkModePage, ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkModeDisabled, ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkModeBlock, ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkModePageSections:
		return true
	}
	return false
}

type ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlock string

const (
	ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockHeader        ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlock = "Header"
	ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockFooter        ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlock = "Footer"
	ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockTitle         ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlock = "Title"
	ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockSectionHeader ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlock = "Section Header"
	ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockPageNumber    ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlock = "Page Number"
	ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockListItem      ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlock = "List Item"
	ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockFigure        ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlock = "Figure"
	ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockTable         ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlock = "Table"
	ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockKeyValue      ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlock = "Key Value"
	ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockText          ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlock = "Text"
	ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockComment       ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlock = "Comment"
	ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockSignature     ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlock = "Signature"
)

func (r ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlock) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockHeader, ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockFooter, ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockTitle, ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockSectionHeader, ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockPageNumber, ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockListItem, ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockFigure, ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockTable, ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockKeyValue, ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockText, ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockComment, ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockSignature:
		return true
	}
	return false
}

type ExtractRunParamsBodySyncExtractConfigParsingSettings struct {
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
	OcrSystem param.Field[ExtractRunParamsBodySyncExtractConfigParsingSettingsOcrSystem] `json:"ocr_system"`
	// The page range to process (1-indexed). By default, the entire document is
	// processed.
	PageRange param.Field[ExtractRunParamsBodySyncExtractConfigParsingSettingsPageRangeUnion] `json:"page_range"`
	// If True, persist the results indefinitely. Defaults to False.
	PersistResults param.Field[bool] `json:"persist_results"`
	// Whether to return images for the specified block types. By default, no images
	// are returned.
	ReturnImages param.Field[[]ExtractRunParamsBodySyncExtractConfigParsingSettingsReturnImage] `json:"return_images"`
	// If True, return OCR data in the result. Defaults to False.
	ReturnOcrData param.Field[bool] `json:"return_ocr_data"`
	// The timeout for the job in seconds. Defaults to 900.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r ExtractRunParamsBodySyncExtractConfigParsingSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Standard is our best multilingual OCR system. Legacy only supports germanic
// languages and is available for backwards compatibility.
type ExtractRunParamsBodySyncExtractConfigParsingSettingsOcrSystem string

const (
	ExtractRunParamsBodySyncExtractConfigParsingSettingsOcrSystemStandard ExtractRunParamsBodySyncExtractConfigParsingSettingsOcrSystem = "standard"
	ExtractRunParamsBodySyncExtractConfigParsingSettingsOcrSystemLegacy   ExtractRunParamsBodySyncExtractConfigParsingSettingsOcrSystem = "legacy"
)

func (r ExtractRunParamsBodySyncExtractConfigParsingSettingsOcrSystem) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodySyncExtractConfigParsingSettingsOcrSystemStandard, ExtractRunParamsBodySyncExtractConfigParsingSettingsOcrSystemLegacy:
		return true
	}
	return false
}

// The page range to process (1-indexed). By default, the entire document is
// processed.
//
// Satisfied by [shared.PageRangeParam],
// [ExtractRunParamsBodySyncExtractConfigParsingSettingsPageRangeArray],
// [ExtractRunParamsBodySyncExtractConfigParsingSettingsPageRangeArray].
type ExtractRunParamsBodySyncExtractConfigParsingSettingsPageRangeUnion interface {
	ImplementsExtractRunParamsBodySyncExtractConfigParsingSettingsPageRangeUnion()
}

type ExtractRunParamsBodySyncExtractConfigParsingSettingsPageRangeArray []shared.PageRangeParam

func (r ExtractRunParamsBodySyncExtractConfigParsingSettingsPageRangeArray) ImplementsExtractRunParamsBodySyncExtractConfigParsingSettingsPageRangeUnion() {
}

type ExtractRunParamsBodySyncExtractConfigParsingSettingsReturnImage string

const (
	ExtractRunParamsBodySyncExtractConfigParsingSettingsReturnImageFigure ExtractRunParamsBodySyncExtractConfigParsingSettingsReturnImage = "figure"
	ExtractRunParamsBodySyncExtractConfigParsingSettingsReturnImageTable  ExtractRunParamsBodySyncExtractConfigParsingSettingsReturnImage = "table"
)

func (r ExtractRunParamsBodySyncExtractConfigParsingSettingsReturnImage) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodySyncExtractConfigParsingSettingsReturnImageFigure, ExtractRunParamsBodySyncExtractConfigParsingSettingsReturnImageTable:
		return true
	}
	return false
}

type ExtractRunParamsBodySyncExtractConfigParsingSpreadsheet struct {
	// In a spreadsheet with different tables inside, we enable splitting up the tables
	// by default. Accurate mode applies more powerful models for superior accuracy, at
	// 5× the default per-cell rate. Disabling will register as one large table.
	Clustering param.Field[ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetClustering] `json:"clustering"`
	// Whether to exclude hidden sheets, rows, or columns in the output.
	Exclude param.Field[[]ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetExclude] `json:"exclude"`
	// Whether to include cell color and formula information in the output.
	Include          param.Field[[]ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetInclude]        `json:"include"`
	SplitLargeTables param.Field[ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetSplitLargeTables] `json:"split_large_tables"`
}

func (r ExtractRunParamsBodySyncExtractConfigParsingSpreadsheet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// In a spreadsheet with different tables inside, we enable splitting up the tables
// by default. Accurate mode applies more powerful models for superior accuracy, at
// 5× the default per-cell rate. Disabling will register as one large table.
type ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetClustering string

const (
	ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetClusteringAccurate ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetClustering = "accurate"
	ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetClusteringFast     ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetClustering = "fast"
	ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetClusteringDisabled ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetClustering = "disabled"
)

func (r ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetClustering) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetClusteringAccurate, ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetClusteringFast, ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetClusteringDisabled:
		return true
	}
	return false
}

type ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetExclude string

const (
	ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetExcludeHiddenSheets ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetExclude = "hidden_sheets"
	ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetExcludeHiddenRows   ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetExclude = "hidden_rows"
	ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetExcludeHiddenCols   ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetExclude = "hidden_cols"
)

func (r ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetExclude) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetExcludeHiddenSheets, ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetExcludeHiddenRows, ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetExcludeHiddenCols:
		return true
	}
	return false
}

type ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetInclude string

const (
	ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetIncludeCellColors ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetInclude = "cell_colors"
	ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetIncludeFormula    ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetInclude = "formula"
)

func (r ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetInclude) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetIncludeCellColors, ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetIncludeFormula:
		return true
	}
	return false
}

type ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetSplitLargeTables struct {
	// If True, split large tables into smaller tables. Defaults to True.
	Enabled param.Field[bool] `json:"enabled"`
	// The size of the tables to split into. Defaults to 50.
	Size param.Field[int64] `json:"size"`
}

func (r ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetSplitLargeTables) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The settings to use for the extraction.
type ExtractRunParamsBodySyncExtractConfigSettings struct {
	// If True, use array extraction.
	ArrayExtract param.Field[bool] `json:"array_extract"`
	// The citations to use for the extraction.
	Citations param.Field[ExtractRunParamsBodySyncExtractConfigSettingsCitations] `json:"citations"`
	// If True, include images in the extraction.
	IncludeImages param.Field[bool] `json:"include_images"`
	// If True, jobs will be processed with a higher throughput and priority at a
	// higher cost. Defaults to False.
	OptimizeForLatency param.Field[bool] `json:"optimize_for_latency"`
}

func (r ExtractRunParamsBodySyncExtractConfigSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The citations to use for the extraction.
type ExtractRunParamsBodySyncExtractConfigSettingsCitations struct {
	// If True, include citations in the extraction.
	Enabled param.Field[bool] `json:"enabled"`
	// If True, enable numeric citation confidence scores. Defaults to True.
	NumericalConfidence param.Field[bool] `json:"numerical_confidence"`
}

func (r ExtractRunParamsBodySyncExtractConfigSettingsCitations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunParamsBodyAsyncExtractConfig struct {
	// The URL of the document to be processed. You can provide one of the
	// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
	// prefixed URL obtained from the /upload endpoint after directly uploading a
	// document 4. A jobid:// prefixed URL obtained from a previous /parse invocation
	Input param.Field[ExtractRunParamsBodyAsyncExtractConfigInputUnion] `json:"input,required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async param.Field[ExtractRunParamsBodyAsyncExtractConfigAsync] `json:"async"`
	// The instructions to use for the extraction.
	Instructions param.Field[ExtractRunParamsBodyAsyncExtractConfigInstructions] `json:"instructions"`
	// The configuration options for parsing the document. If you are passing in a
	// jobid:// URL for the file, then this configuration will be ignored.
	Parsing param.Field[ExtractRunParamsBodyAsyncExtractConfigParsing] `json:"parsing"`
	// The settings to use for the extraction.
	Settings param.Field[ExtractRunParamsBodyAsyncExtractConfigSettings] `json:"settings"`
}

func (r ExtractRunParamsBodyAsyncExtractConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBodyAsyncExtractConfig) implementsExtractRunParamsBodyUnion() {}

// The URL of the document to be processed. You can provide one of the
// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
// prefixed URL obtained from the /upload endpoint after directly uploading a
// document 4. A jobid:// prefixed URL obtained from a previous /parse invocation
//
// Satisfied by [shared.UnionString], [shared.UploadParam].
type ExtractRunParamsBodyAsyncExtractConfigInputUnion interface {
	ImplementsExtractRunParamsBodyAsyncExtractConfigInputUnion()
}

// The configuration options for asynchronous processing (default synchronous).
type ExtractRunParamsBodyAsyncExtractConfigAsync struct {
	// JSON metadata included in webhook request body. Defaults to None.
	Metadata param.Field[interface{}] `json:"metadata"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool] `json:"priority"`
	// The webhook configuration for the asynchronous processing.
	Webhook param.Field[ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookUnion] `json:"webhook"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigAsync) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The webhook configuration for the asynchronous processing.
type ExtractRunParamsBodyAsyncExtractConfigAsyncWebhook struct {
	Channels param.Field[interface{}]                                            `json:"channels"`
	Mode     param.Field[ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookMode] `json:"mode"`
	URL      param.Field[string]                                                 `json:"url"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigAsyncWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBodyAsyncExtractConfigAsyncWebhook) implementsExtractRunParamsBodyAsyncExtractConfigAsyncWebhookUnion() {
}

// The webhook configuration for the asynchronous processing.
//
// Satisfied by
// [ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfig],
// [ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfig],
// [ExtractRunParamsBodyAsyncExtractConfigAsyncWebhook].
type ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookUnion interface {
	implementsExtractRunParamsBodyAsyncExtractConfigAsyncWebhookUnion()
}

type ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfig struct {
	// A list of Svix channels the message will be delivered down, omit to send to all
	// channels.
	Channels param.Field[[]string]                                                                `json:"channels"`
	Mode     param.Field[ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfigMode] `json:"mode"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfig) implementsExtractRunParamsBodyAsyncExtractConfigAsyncWebhookUnion() {
}

type ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfigMode string

const (
	ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfigModeSvix ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfigMode = "svix"
)

func (r ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfigMode) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfigModeSvix:
		return true
	}
	return false
}

type ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfig struct {
	URL  param.Field[string]                                                                    `json:"url,required"`
	Mode param.Field[ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfigMode] `json:"mode"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfig) implementsExtractRunParamsBodyAsyncExtractConfigAsyncWebhookUnion() {
}

type ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfigMode string

const (
	ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfigModeDirect ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfigMode = "direct"
)

func (r ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfigMode) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfigModeDirect:
		return true
	}
	return false
}

type ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookMode string

const (
	ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookModeSvix   ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookMode = "svix"
	ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookModeDirect ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookMode = "direct"
)

func (r ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookMode) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookModeSvix, ExtractRunParamsBodyAsyncExtractConfigAsyncWebhookModeDirect:
		return true
	}
	return false
}

// The instructions to use for the extraction.
type ExtractRunParamsBodyAsyncExtractConfigInstructions struct {
	// The JSON schema to use for the extraction.
	Schema param.Field[interface{}] `json:"schema"`
	// The system prompt to use for the extraction.
	SystemPrompt param.Field[string] `json:"system_prompt"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigInstructions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration options for parsing the document. If you are passing in a
// jobid:// URL for the file, then this configuration will be ignored.
type ExtractRunParamsBodyAsyncExtractConfigParsing struct {
	Enhance     param.Field[ExtractRunParamsBodyAsyncExtractConfigParsingEnhance]     `json:"enhance"`
	Formatting  param.Field[ExtractRunParamsBodyAsyncExtractConfigParsingFormatting]  `json:"formatting"`
	Retrieval   param.Field[ExtractRunParamsBodyAsyncExtractConfigParsingRetrieval]   `json:"retrieval"`
	Settings    param.Field[ExtractRunParamsBodyAsyncExtractConfigParsingSettings]    `json:"settings"`
	Spreadsheet param.Field[ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheet] `json:"spreadsheet"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigParsing) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunParamsBodyAsyncExtractConfigParsingEnhance struct {
	// Agentic uses vision language models to enhance the accuracy of the output of
	// different types of extraction. This will incur a cost and latency increase.
	Agentic param.Field[[]ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticUnion] `json:"agentic"`
	// If True, summarize figures using a small vision language model. Defaults to
	// True.
	SummarizeFigures param.Field[bool] `json:"summarize_figures"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigParsingEnhance) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgentic struct {
	Scope param.Field[ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticScope] `json:"scope,required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgentic) implementsExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticUnion() {
}

// Satisfied by
// [ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgentic],
// [ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgentic],
// [ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgentic],
// [ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgentic].
type ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticUnion interface {
	implementsExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticUnion()
}

type ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgentic struct {
	Scope param.Field[ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgenticScope] `json:"scope,required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgentic) implementsExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticUnion() {
}

type ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgenticScope string

const (
	ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgenticScopeTable ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgenticScope = "table"
)

func (r ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgenticScope) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgenticScopeTable:
		return true
	}
	return false
}

type ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgentic struct {
	Scope param.Field[ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgenticScope] `json:"scope,required"`
	// Custom prompt for figure agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgentic) implementsExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticUnion() {
}

type ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgenticScope string

const (
	ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgenticScopeFigure ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgenticScope = "figure"
)

func (r ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgenticScope) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgenticScopeFigure:
		return true
	}
	return false
}

type ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgentic struct {
	Scope param.Field[ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgenticScope] `json:"scope,required"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgentic) implementsExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticUnion() {
}

type ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgenticScope string

const (
	ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgenticScopeText ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgenticScope = "text"
)

func (r ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgenticScope) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgenticScopeText:
		return true
	}
	return false
}

type ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticScope string

const (
	ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticScopeTable  ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticScope = "table"
	ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticScopeFigure ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticScope = "figure"
	ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticScopeText   ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticScope = "text"
)

func (r ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticScope) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticScopeTable, ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticScopeFigure, ExtractRunParamsBodyAsyncExtractConfigParsingEnhanceAgenticScopeText:
		return true
	}
	return false
}

type ExtractRunParamsBodyAsyncExtractConfigParsingFormatting struct {
	// If True, add page markers to the output. Defaults to False. Useful for
	// extracting data with page specific information.
	AddPageMarkers param.Field[bool] `json:"add_page_markers"`
	// A list of formatting to include in the output. [insert description of each
	// option here later]
	Include param.Field[[]ExtractRunParamsBodyAsyncExtractConfigParsingFormattingInclude] `json:"include"`
	// A flag to indicate if consecutive tables with the same number of columns should
	// be merged. Defaults to False.
	MergeTables param.Field[bool] `json:"merge_tables"`
	// The mode to use for table output. Defaults to dynamic, which returns md for
	// simpler tables and html for more complex tables.
	TableOutputFormat param.Field[ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormat] `json:"table_output_format"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigParsingFormatting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunParamsBodyAsyncExtractConfigParsingFormattingInclude string

const (
	ExtractRunParamsBodyAsyncExtractConfigParsingFormattingIncludeChangeTracking ExtractRunParamsBodyAsyncExtractConfigParsingFormattingInclude = "change_tracking"
	ExtractRunParamsBodyAsyncExtractConfigParsingFormattingIncludeHighlight      ExtractRunParamsBodyAsyncExtractConfigParsingFormattingInclude = "highlight"
	ExtractRunParamsBodyAsyncExtractConfigParsingFormattingIncludeComments       ExtractRunParamsBodyAsyncExtractConfigParsingFormattingInclude = "comments"
)

func (r ExtractRunParamsBodyAsyncExtractConfigParsingFormattingInclude) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyAsyncExtractConfigParsingFormattingIncludeChangeTracking, ExtractRunParamsBodyAsyncExtractConfigParsingFormattingIncludeHighlight, ExtractRunParamsBodyAsyncExtractConfigParsingFormattingIncludeComments:
		return true
	}
	return false
}

// The mode to use for table output. Defaults to dynamic, which returns md for
// simpler tables and html for more complex tables.
type ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormat string

const (
	ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatHTML     ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormat = "html"
	ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatJson     ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormat = "json"
	ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatMd       ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormat = "md"
	ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatJsonbbox ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormat = "jsonbbox"
	ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatDynamic  ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormat = "dynamic"
	ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatCsv      ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormat = "csv"
)

func (r ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormat) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatHTML, ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatJson, ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatMd, ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatJsonbbox, ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatDynamic, ExtractRunParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatCsv:
		return true
	}
	return false
}

type ExtractRunParamsBodyAsyncExtractConfigParsingRetrieval struct {
	Chunking param.Field[ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunking] `json:"chunking"`
	// If True, use embedding optimized mode. Defaults to False.
	EmbeddingOptimized param.Field[bool] `json:"embedding_optimized"`
	// A list of block types to filter out from 'content' and 'embed' fields. By
	// default, no blocks are filtered.
	FilterBlocks param.Field[[]ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock] `json:"filter_blocks"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigParsingRetrieval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunking struct {
	// Choose how to partition chunks. Variable mode chunks by character length and
	// visual context. Section mode chunks by section headers. Page mode chunks
	// according to pages. Page sections mode chunks first by page, then by sections
	// within each page. Disabled returns one single chunk.
	ChunkMode param.Field[ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkMode] `json:"chunk_mode"`
	// The approximate size of chunks (in characters) that the document will be split
	// into. Defaults to null, in which case the chunk size is variable between 250 -
	// 1500 characters.
	ChunkSize param.Field[int64] `json:"chunk_size"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Choose how to partition chunks. Variable mode chunks by character length and
// visual context. Section mode chunks by section headers. Page mode chunks
// according to pages. Page sections mode chunks first by page, then by sections
// within each page. Disabled returns one single chunk.
type ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkMode string

const (
	ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModeVariable     ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkMode = "variable"
	ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModeSection      ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkMode = "section"
	ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModePage         ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkMode = "page"
	ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModeDisabled     ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkMode = "disabled"
	ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModeBlock        ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkMode = "block"
	ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModePageSections ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkMode = "page_sections"
)

func (r ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkMode) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModeVariable, ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModeSection, ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModePage, ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModeDisabled, ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModeBlock, ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModePageSections:
		return true
	}
	return false
}

type ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock string

const (
	ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockHeader        ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Header"
	ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockFooter        ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Footer"
	ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockTitle         ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Title"
	ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockSectionHeader ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Section Header"
	ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockPageNumber    ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Page Number"
	ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockListItem      ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "List Item"
	ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockFigure        ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Figure"
	ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockTable         ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Table"
	ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockKeyValue      ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Key Value"
	ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockText          ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Text"
	ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockComment       ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Comment"
	ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockSignature     ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Signature"
)

func (r ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockHeader, ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockFooter, ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockTitle, ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockSectionHeader, ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockPageNumber, ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockListItem, ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockFigure, ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockTable, ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockKeyValue, ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockText, ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockComment, ExtractRunParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockSignature:
		return true
	}
	return false
}

type ExtractRunParamsBodyAsyncExtractConfigParsingSettings struct {
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
	OcrSystem param.Field[ExtractRunParamsBodyAsyncExtractConfigParsingSettingsOcrSystem] `json:"ocr_system"`
	// The page range to process (1-indexed). By default, the entire document is
	// processed.
	PageRange param.Field[ExtractRunParamsBodyAsyncExtractConfigParsingSettingsPageRangeUnion] `json:"page_range"`
	// If True, persist the results indefinitely. Defaults to False.
	PersistResults param.Field[bool] `json:"persist_results"`
	// Whether to return images for the specified block types. By default, no images
	// are returned.
	ReturnImages param.Field[[]ExtractRunParamsBodyAsyncExtractConfigParsingSettingsReturnImage] `json:"return_images"`
	// If True, return OCR data in the result. Defaults to False.
	ReturnOcrData param.Field[bool] `json:"return_ocr_data"`
	// The timeout for the job in seconds. Defaults to 900.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigParsingSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Standard is our best multilingual OCR system. Legacy only supports germanic
// languages and is available for backwards compatibility.
type ExtractRunParamsBodyAsyncExtractConfigParsingSettingsOcrSystem string

const (
	ExtractRunParamsBodyAsyncExtractConfigParsingSettingsOcrSystemStandard ExtractRunParamsBodyAsyncExtractConfigParsingSettingsOcrSystem = "standard"
	ExtractRunParamsBodyAsyncExtractConfigParsingSettingsOcrSystemLegacy   ExtractRunParamsBodyAsyncExtractConfigParsingSettingsOcrSystem = "legacy"
)

func (r ExtractRunParamsBodyAsyncExtractConfigParsingSettingsOcrSystem) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyAsyncExtractConfigParsingSettingsOcrSystemStandard, ExtractRunParamsBodyAsyncExtractConfigParsingSettingsOcrSystemLegacy:
		return true
	}
	return false
}

// The page range to process (1-indexed). By default, the entire document is
// processed.
//
// Satisfied by [shared.PageRangeParam],
// [ExtractRunParamsBodyAsyncExtractConfigParsingSettingsPageRangeArray],
// [ExtractRunParamsBodyAsyncExtractConfigParsingSettingsPageRangeArray].
type ExtractRunParamsBodyAsyncExtractConfigParsingSettingsPageRangeUnion interface {
	ImplementsExtractRunParamsBodyAsyncExtractConfigParsingSettingsPageRangeUnion()
}

type ExtractRunParamsBodyAsyncExtractConfigParsingSettingsPageRangeArray []shared.PageRangeParam

func (r ExtractRunParamsBodyAsyncExtractConfigParsingSettingsPageRangeArray) ImplementsExtractRunParamsBodyAsyncExtractConfigParsingSettingsPageRangeUnion() {
}

type ExtractRunParamsBodyAsyncExtractConfigParsingSettingsReturnImage string

const (
	ExtractRunParamsBodyAsyncExtractConfigParsingSettingsReturnImageFigure ExtractRunParamsBodyAsyncExtractConfigParsingSettingsReturnImage = "figure"
	ExtractRunParamsBodyAsyncExtractConfigParsingSettingsReturnImageTable  ExtractRunParamsBodyAsyncExtractConfigParsingSettingsReturnImage = "table"
)

func (r ExtractRunParamsBodyAsyncExtractConfigParsingSettingsReturnImage) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyAsyncExtractConfigParsingSettingsReturnImageFigure, ExtractRunParamsBodyAsyncExtractConfigParsingSettingsReturnImageTable:
		return true
	}
	return false
}

type ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheet struct {
	// In a spreadsheet with different tables inside, we enable splitting up the tables
	// by default. Accurate mode applies more powerful models for superior accuracy, at
	// 5× the default per-cell rate. Disabling will register as one large table.
	Clustering param.Field[ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetClustering] `json:"clustering"`
	// Whether to exclude hidden sheets, rows, or columns in the output.
	Exclude param.Field[[]ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetExclude] `json:"exclude"`
	// Whether to include cell color and formula information in the output.
	Include          param.Field[[]ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetInclude]        `json:"include"`
	SplitLargeTables param.Field[ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetSplitLargeTables] `json:"split_large_tables"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// In a spreadsheet with different tables inside, we enable splitting up the tables
// by default. Accurate mode applies more powerful models for superior accuracy, at
// 5× the default per-cell rate. Disabling will register as one large table.
type ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetClustering string

const (
	ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetClusteringAccurate ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetClustering = "accurate"
	ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetClusteringFast     ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetClustering = "fast"
	ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetClusteringDisabled ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetClustering = "disabled"
)

func (r ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetClustering) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetClusteringAccurate, ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetClusteringFast, ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetClusteringDisabled:
		return true
	}
	return false
}

type ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetExclude string

const (
	ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetExcludeHiddenSheets ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetExclude = "hidden_sheets"
	ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetExcludeHiddenRows   ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetExclude = "hidden_rows"
	ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetExcludeHiddenCols   ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetExclude = "hidden_cols"
)

func (r ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetExclude) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetExcludeHiddenSheets, ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetExcludeHiddenRows, ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetExcludeHiddenCols:
		return true
	}
	return false
}

type ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetInclude string

const (
	ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetIncludeCellColors ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetInclude = "cell_colors"
	ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetIncludeFormula    ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetInclude = "formula"
)

func (r ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetInclude) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetIncludeCellColors, ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetIncludeFormula:
		return true
	}
	return false
}

type ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetSplitLargeTables struct {
	// If True, split large tables into smaller tables. Defaults to True.
	Enabled param.Field[bool] `json:"enabled"`
	// The size of the tables to split into. Defaults to 50.
	Size param.Field[int64] `json:"size"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigParsingSpreadsheetSplitLargeTables) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The settings to use for the extraction.
type ExtractRunParamsBodyAsyncExtractConfigSettings struct {
	// If True, use array extraction.
	ArrayExtract param.Field[bool] `json:"array_extract"`
	// The citations to use for the extraction.
	Citations param.Field[ExtractRunParamsBodyAsyncExtractConfigSettingsCitations] `json:"citations"`
	// If True, include images in the extraction.
	IncludeImages param.Field[bool] `json:"include_images"`
	// If True, jobs will be processed with a higher throughput and priority at a
	// higher cost. Defaults to False.
	OptimizeForLatency param.Field[bool] `json:"optimize_for_latency"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The citations to use for the extraction.
type ExtractRunParamsBodyAsyncExtractConfigSettingsCitations struct {
	// If True, include citations in the extraction.
	Enabled param.Field[bool] `json:"enabled"`
	// If True, enable numeric citation confidence scores. Defaults to True.
	NumericalConfidence param.Field[bool] `json:"numerical_confidence"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigSettingsCitations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunJobParams struct {
	// The URL of the document to be processed. You can provide one of the
	// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
	// prefixed URL obtained from the /upload endpoint after directly uploading a
	// document 4. A jobid:// prefixed URL obtained from a previous /parse invocation
	Input param.Field[ExtractRunJobParamsInputUnion] `json:"input,required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async param.Field[ExtractRunJobParamsAsync] `json:"async"`
	// The instructions to use for the extraction.
	Instructions param.Field[ExtractRunJobParamsInstructions] `json:"instructions"`
	// The configuration options for parsing the document. If you are passing in a
	// jobid:// URL for the file, then this configuration will be ignored.
	Parsing param.Field[ExtractRunJobParamsParsing] `json:"parsing"`
	// The settings to use for the extraction.
	Settings param.Field[ExtractRunJobParamsSettings] `json:"settings"`
}

func (r ExtractRunJobParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The URL of the document to be processed. You can provide one of the
// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
// prefixed URL obtained from the /upload endpoint after directly uploading a
// document 4. A jobid:// prefixed URL obtained from a previous /parse invocation
//
// Satisfied by [shared.UnionString], [shared.UploadParam].
type ExtractRunJobParamsInputUnion interface {
	ImplementsExtractRunJobParamsInputUnion()
}

// The configuration options for asynchronous processing (default synchronous).
type ExtractRunJobParamsAsync struct {
	// JSON metadata included in webhook request body. Defaults to None.
	Metadata param.Field[interface{}] `json:"metadata"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool] `json:"priority"`
	// The webhook configuration for the asynchronous processing.
	Webhook param.Field[ExtractRunJobParamsAsyncWebhookUnion] `json:"webhook"`
}

func (r ExtractRunJobParamsAsync) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The webhook configuration for the asynchronous processing.
type ExtractRunJobParamsAsyncWebhook struct {
	Channels param.Field[interface{}]                         `json:"channels"`
	Mode     param.Field[ExtractRunJobParamsAsyncWebhookMode] `json:"mode"`
	URL      param.Field[string]                              `json:"url"`
}

func (r ExtractRunJobParamsAsyncWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunJobParamsAsyncWebhook) implementsExtractRunJobParamsAsyncWebhookUnion() {}

// The webhook configuration for the asynchronous processing.
//
// Satisfied by [ExtractRunJobParamsAsyncWebhookSvixWebhookConfig],
// [ExtractRunJobParamsAsyncWebhookDirectWebhookConfig],
// [ExtractRunJobParamsAsyncWebhook].
type ExtractRunJobParamsAsyncWebhookUnion interface {
	implementsExtractRunJobParamsAsyncWebhookUnion()
}

type ExtractRunJobParamsAsyncWebhookSvixWebhookConfig struct {
	// A list of Svix channels the message will be delivered down, omit to send to all
	// channels.
	Channels param.Field[[]string]                                             `json:"channels"`
	Mode     param.Field[ExtractRunJobParamsAsyncWebhookSvixWebhookConfigMode] `json:"mode"`
}

func (r ExtractRunJobParamsAsyncWebhookSvixWebhookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunJobParamsAsyncWebhookSvixWebhookConfig) implementsExtractRunJobParamsAsyncWebhookUnion() {
}

type ExtractRunJobParamsAsyncWebhookSvixWebhookConfigMode string

const (
	ExtractRunJobParamsAsyncWebhookSvixWebhookConfigModeSvix ExtractRunJobParamsAsyncWebhookSvixWebhookConfigMode = "svix"
)

func (r ExtractRunJobParamsAsyncWebhookSvixWebhookConfigMode) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsAsyncWebhookSvixWebhookConfigModeSvix:
		return true
	}
	return false
}

type ExtractRunJobParamsAsyncWebhookDirectWebhookConfig struct {
	URL  param.Field[string]                                                 `json:"url,required"`
	Mode param.Field[ExtractRunJobParamsAsyncWebhookDirectWebhookConfigMode] `json:"mode"`
}

func (r ExtractRunJobParamsAsyncWebhookDirectWebhookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunJobParamsAsyncWebhookDirectWebhookConfig) implementsExtractRunJobParamsAsyncWebhookUnion() {
}

type ExtractRunJobParamsAsyncWebhookDirectWebhookConfigMode string

const (
	ExtractRunJobParamsAsyncWebhookDirectWebhookConfigModeDirect ExtractRunJobParamsAsyncWebhookDirectWebhookConfigMode = "direct"
)

func (r ExtractRunJobParamsAsyncWebhookDirectWebhookConfigMode) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsAsyncWebhookDirectWebhookConfigModeDirect:
		return true
	}
	return false
}

type ExtractRunJobParamsAsyncWebhookMode string

const (
	ExtractRunJobParamsAsyncWebhookModeSvix   ExtractRunJobParamsAsyncWebhookMode = "svix"
	ExtractRunJobParamsAsyncWebhookModeDirect ExtractRunJobParamsAsyncWebhookMode = "direct"
)

func (r ExtractRunJobParamsAsyncWebhookMode) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsAsyncWebhookModeSvix, ExtractRunJobParamsAsyncWebhookModeDirect:
		return true
	}
	return false
}

// The instructions to use for the extraction.
type ExtractRunJobParamsInstructions struct {
	// The JSON schema to use for the extraction.
	Schema param.Field[interface{}] `json:"schema"`
	// The system prompt to use for the extraction.
	SystemPrompt param.Field[string] `json:"system_prompt"`
}

func (r ExtractRunJobParamsInstructions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration options for parsing the document. If you are passing in a
// jobid:// URL for the file, then this configuration will be ignored.
type ExtractRunJobParamsParsing struct {
	Enhance     param.Field[ExtractRunJobParamsParsingEnhance]     `json:"enhance"`
	Formatting  param.Field[ExtractRunJobParamsParsingFormatting]  `json:"formatting"`
	Retrieval   param.Field[ExtractRunJobParamsParsingRetrieval]   `json:"retrieval"`
	Settings    param.Field[ExtractRunJobParamsParsingSettings]    `json:"settings"`
	Spreadsheet param.Field[ExtractRunJobParamsParsingSpreadsheet] `json:"spreadsheet"`
}

func (r ExtractRunJobParamsParsing) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunJobParamsParsingEnhance struct {
	// Agentic uses vision language models to enhance the accuracy of the output of
	// different types of extraction. This will incur a cost and latency increase.
	Agentic param.Field[[]ExtractRunJobParamsParsingEnhanceAgenticUnion] `json:"agentic"`
	// If True, summarize figures using a small vision language model. Defaults to
	// True.
	SummarizeFigures param.Field[bool] `json:"summarize_figures"`
}

func (r ExtractRunJobParamsParsingEnhance) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunJobParamsParsingEnhanceAgentic struct {
	Scope param.Field[ExtractRunJobParamsParsingEnhanceAgenticScope] `json:"scope,required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ExtractRunJobParamsParsingEnhanceAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunJobParamsParsingEnhanceAgentic) implementsExtractRunJobParamsParsingEnhanceAgenticUnion() {
}

// Satisfied by [ExtractRunJobParamsParsingEnhanceAgenticTableAgentic],
// [ExtractRunJobParamsParsingEnhanceAgenticFigureAgentic],
// [ExtractRunJobParamsParsingEnhanceAgenticTextAgentic],
// [ExtractRunJobParamsParsingEnhanceAgentic].
type ExtractRunJobParamsParsingEnhanceAgenticUnion interface {
	implementsExtractRunJobParamsParsingEnhanceAgenticUnion()
}

type ExtractRunJobParamsParsingEnhanceAgenticTableAgentic struct {
	Scope param.Field[ExtractRunJobParamsParsingEnhanceAgenticTableAgenticScope] `json:"scope,required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ExtractRunJobParamsParsingEnhanceAgenticTableAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunJobParamsParsingEnhanceAgenticTableAgentic) implementsExtractRunJobParamsParsingEnhanceAgenticUnion() {
}

type ExtractRunJobParamsParsingEnhanceAgenticTableAgenticScope string

const (
	ExtractRunJobParamsParsingEnhanceAgenticTableAgenticScopeTable ExtractRunJobParamsParsingEnhanceAgenticTableAgenticScope = "table"
)

func (r ExtractRunJobParamsParsingEnhanceAgenticTableAgenticScope) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsParsingEnhanceAgenticTableAgenticScopeTable:
		return true
	}
	return false
}

type ExtractRunJobParamsParsingEnhanceAgenticFigureAgentic struct {
	Scope param.Field[ExtractRunJobParamsParsingEnhanceAgenticFigureAgenticScope] `json:"scope,required"`
	// Custom prompt for figure agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ExtractRunJobParamsParsingEnhanceAgenticFigureAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunJobParamsParsingEnhanceAgenticFigureAgentic) implementsExtractRunJobParamsParsingEnhanceAgenticUnion() {
}

type ExtractRunJobParamsParsingEnhanceAgenticFigureAgenticScope string

const (
	ExtractRunJobParamsParsingEnhanceAgenticFigureAgenticScopeFigure ExtractRunJobParamsParsingEnhanceAgenticFigureAgenticScope = "figure"
)

func (r ExtractRunJobParamsParsingEnhanceAgenticFigureAgenticScope) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsParsingEnhanceAgenticFigureAgenticScopeFigure:
		return true
	}
	return false
}

type ExtractRunJobParamsParsingEnhanceAgenticTextAgentic struct {
	Scope param.Field[ExtractRunJobParamsParsingEnhanceAgenticTextAgenticScope] `json:"scope,required"`
}

func (r ExtractRunJobParamsParsingEnhanceAgenticTextAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunJobParamsParsingEnhanceAgenticTextAgentic) implementsExtractRunJobParamsParsingEnhanceAgenticUnion() {
}

type ExtractRunJobParamsParsingEnhanceAgenticTextAgenticScope string

const (
	ExtractRunJobParamsParsingEnhanceAgenticTextAgenticScopeText ExtractRunJobParamsParsingEnhanceAgenticTextAgenticScope = "text"
)

func (r ExtractRunJobParamsParsingEnhanceAgenticTextAgenticScope) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsParsingEnhanceAgenticTextAgenticScopeText:
		return true
	}
	return false
}

type ExtractRunJobParamsParsingEnhanceAgenticScope string

const (
	ExtractRunJobParamsParsingEnhanceAgenticScopeTable  ExtractRunJobParamsParsingEnhanceAgenticScope = "table"
	ExtractRunJobParamsParsingEnhanceAgenticScopeFigure ExtractRunJobParamsParsingEnhanceAgenticScope = "figure"
	ExtractRunJobParamsParsingEnhanceAgenticScopeText   ExtractRunJobParamsParsingEnhanceAgenticScope = "text"
)

func (r ExtractRunJobParamsParsingEnhanceAgenticScope) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsParsingEnhanceAgenticScopeTable, ExtractRunJobParamsParsingEnhanceAgenticScopeFigure, ExtractRunJobParamsParsingEnhanceAgenticScopeText:
		return true
	}
	return false
}

type ExtractRunJobParamsParsingFormatting struct {
	// If True, add page markers to the output. Defaults to False. Useful for
	// extracting data with page specific information.
	AddPageMarkers param.Field[bool] `json:"add_page_markers"`
	// A list of formatting to include in the output. [insert description of each
	// option here later]
	Include param.Field[[]ExtractRunJobParamsParsingFormattingInclude] `json:"include"`
	// A flag to indicate if consecutive tables with the same number of columns should
	// be merged. Defaults to False.
	MergeTables param.Field[bool] `json:"merge_tables"`
	// The mode to use for table output. Defaults to dynamic, which returns md for
	// simpler tables and html for more complex tables.
	TableOutputFormat param.Field[ExtractRunJobParamsParsingFormattingTableOutputFormat] `json:"table_output_format"`
}

func (r ExtractRunJobParamsParsingFormatting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunJobParamsParsingFormattingInclude string

const (
	ExtractRunJobParamsParsingFormattingIncludeChangeTracking ExtractRunJobParamsParsingFormattingInclude = "change_tracking"
	ExtractRunJobParamsParsingFormattingIncludeHighlight      ExtractRunJobParamsParsingFormattingInclude = "highlight"
	ExtractRunJobParamsParsingFormattingIncludeComments       ExtractRunJobParamsParsingFormattingInclude = "comments"
)

func (r ExtractRunJobParamsParsingFormattingInclude) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsParsingFormattingIncludeChangeTracking, ExtractRunJobParamsParsingFormattingIncludeHighlight, ExtractRunJobParamsParsingFormattingIncludeComments:
		return true
	}
	return false
}

// The mode to use for table output. Defaults to dynamic, which returns md for
// simpler tables and html for more complex tables.
type ExtractRunJobParamsParsingFormattingTableOutputFormat string

const (
	ExtractRunJobParamsParsingFormattingTableOutputFormatHTML     ExtractRunJobParamsParsingFormattingTableOutputFormat = "html"
	ExtractRunJobParamsParsingFormattingTableOutputFormatJson     ExtractRunJobParamsParsingFormattingTableOutputFormat = "json"
	ExtractRunJobParamsParsingFormattingTableOutputFormatMd       ExtractRunJobParamsParsingFormattingTableOutputFormat = "md"
	ExtractRunJobParamsParsingFormattingTableOutputFormatJsonbbox ExtractRunJobParamsParsingFormattingTableOutputFormat = "jsonbbox"
	ExtractRunJobParamsParsingFormattingTableOutputFormatDynamic  ExtractRunJobParamsParsingFormattingTableOutputFormat = "dynamic"
	ExtractRunJobParamsParsingFormattingTableOutputFormatCsv      ExtractRunJobParamsParsingFormattingTableOutputFormat = "csv"
)

func (r ExtractRunJobParamsParsingFormattingTableOutputFormat) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsParsingFormattingTableOutputFormatHTML, ExtractRunJobParamsParsingFormattingTableOutputFormatJson, ExtractRunJobParamsParsingFormattingTableOutputFormatMd, ExtractRunJobParamsParsingFormattingTableOutputFormatJsonbbox, ExtractRunJobParamsParsingFormattingTableOutputFormatDynamic, ExtractRunJobParamsParsingFormattingTableOutputFormatCsv:
		return true
	}
	return false
}

type ExtractRunJobParamsParsingRetrieval struct {
	Chunking param.Field[ExtractRunJobParamsParsingRetrievalChunking] `json:"chunking"`
	// If True, use embedding optimized mode. Defaults to False.
	EmbeddingOptimized param.Field[bool] `json:"embedding_optimized"`
	// A list of block types to filter out from 'content' and 'embed' fields. By
	// default, no blocks are filtered.
	FilterBlocks param.Field[[]ExtractRunJobParamsParsingRetrievalFilterBlock] `json:"filter_blocks"`
}

func (r ExtractRunJobParamsParsingRetrieval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunJobParamsParsingRetrievalChunking struct {
	// Choose how to partition chunks. Variable mode chunks by character length and
	// visual context. Section mode chunks by section headers. Page mode chunks
	// according to pages. Page sections mode chunks first by page, then by sections
	// within each page. Disabled returns one single chunk.
	ChunkMode param.Field[ExtractRunJobParamsParsingRetrievalChunkingChunkMode] `json:"chunk_mode"`
	// The approximate size of chunks (in characters) that the document will be split
	// into. Defaults to null, in which case the chunk size is variable between 250 -
	// 1500 characters.
	ChunkSize param.Field[int64] `json:"chunk_size"`
}

func (r ExtractRunJobParamsParsingRetrievalChunking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Choose how to partition chunks. Variable mode chunks by character length and
// visual context. Section mode chunks by section headers. Page mode chunks
// according to pages. Page sections mode chunks first by page, then by sections
// within each page. Disabled returns one single chunk.
type ExtractRunJobParamsParsingRetrievalChunkingChunkMode string

const (
	ExtractRunJobParamsParsingRetrievalChunkingChunkModeVariable     ExtractRunJobParamsParsingRetrievalChunkingChunkMode = "variable"
	ExtractRunJobParamsParsingRetrievalChunkingChunkModeSection      ExtractRunJobParamsParsingRetrievalChunkingChunkMode = "section"
	ExtractRunJobParamsParsingRetrievalChunkingChunkModePage         ExtractRunJobParamsParsingRetrievalChunkingChunkMode = "page"
	ExtractRunJobParamsParsingRetrievalChunkingChunkModeDisabled     ExtractRunJobParamsParsingRetrievalChunkingChunkMode = "disabled"
	ExtractRunJobParamsParsingRetrievalChunkingChunkModeBlock        ExtractRunJobParamsParsingRetrievalChunkingChunkMode = "block"
	ExtractRunJobParamsParsingRetrievalChunkingChunkModePageSections ExtractRunJobParamsParsingRetrievalChunkingChunkMode = "page_sections"
)

func (r ExtractRunJobParamsParsingRetrievalChunkingChunkMode) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsParsingRetrievalChunkingChunkModeVariable, ExtractRunJobParamsParsingRetrievalChunkingChunkModeSection, ExtractRunJobParamsParsingRetrievalChunkingChunkModePage, ExtractRunJobParamsParsingRetrievalChunkingChunkModeDisabled, ExtractRunJobParamsParsingRetrievalChunkingChunkModeBlock, ExtractRunJobParamsParsingRetrievalChunkingChunkModePageSections:
		return true
	}
	return false
}

type ExtractRunJobParamsParsingRetrievalFilterBlock string

const (
	ExtractRunJobParamsParsingRetrievalFilterBlockHeader        ExtractRunJobParamsParsingRetrievalFilterBlock = "Header"
	ExtractRunJobParamsParsingRetrievalFilterBlockFooter        ExtractRunJobParamsParsingRetrievalFilterBlock = "Footer"
	ExtractRunJobParamsParsingRetrievalFilterBlockTitle         ExtractRunJobParamsParsingRetrievalFilterBlock = "Title"
	ExtractRunJobParamsParsingRetrievalFilterBlockSectionHeader ExtractRunJobParamsParsingRetrievalFilterBlock = "Section Header"
	ExtractRunJobParamsParsingRetrievalFilterBlockPageNumber    ExtractRunJobParamsParsingRetrievalFilterBlock = "Page Number"
	ExtractRunJobParamsParsingRetrievalFilterBlockListItem      ExtractRunJobParamsParsingRetrievalFilterBlock = "List Item"
	ExtractRunJobParamsParsingRetrievalFilterBlockFigure        ExtractRunJobParamsParsingRetrievalFilterBlock = "Figure"
	ExtractRunJobParamsParsingRetrievalFilterBlockTable         ExtractRunJobParamsParsingRetrievalFilterBlock = "Table"
	ExtractRunJobParamsParsingRetrievalFilterBlockKeyValue      ExtractRunJobParamsParsingRetrievalFilterBlock = "Key Value"
	ExtractRunJobParamsParsingRetrievalFilterBlockText          ExtractRunJobParamsParsingRetrievalFilterBlock = "Text"
	ExtractRunJobParamsParsingRetrievalFilterBlockComment       ExtractRunJobParamsParsingRetrievalFilterBlock = "Comment"
	ExtractRunJobParamsParsingRetrievalFilterBlockSignature     ExtractRunJobParamsParsingRetrievalFilterBlock = "Signature"
)

func (r ExtractRunJobParamsParsingRetrievalFilterBlock) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsParsingRetrievalFilterBlockHeader, ExtractRunJobParamsParsingRetrievalFilterBlockFooter, ExtractRunJobParamsParsingRetrievalFilterBlockTitle, ExtractRunJobParamsParsingRetrievalFilterBlockSectionHeader, ExtractRunJobParamsParsingRetrievalFilterBlockPageNumber, ExtractRunJobParamsParsingRetrievalFilterBlockListItem, ExtractRunJobParamsParsingRetrievalFilterBlockFigure, ExtractRunJobParamsParsingRetrievalFilterBlockTable, ExtractRunJobParamsParsingRetrievalFilterBlockKeyValue, ExtractRunJobParamsParsingRetrievalFilterBlockText, ExtractRunJobParamsParsingRetrievalFilterBlockComment, ExtractRunJobParamsParsingRetrievalFilterBlockSignature:
		return true
	}
	return false
}

type ExtractRunJobParamsParsingSettings struct {
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
	OcrSystem param.Field[ExtractRunJobParamsParsingSettingsOcrSystem] `json:"ocr_system"`
	// The page range to process (1-indexed). By default, the entire document is
	// processed.
	PageRange param.Field[ExtractRunJobParamsParsingSettingsPageRangeUnion] `json:"page_range"`
	// If True, persist the results indefinitely. Defaults to False.
	PersistResults param.Field[bool] `json:"persist_results"`
	// Whether to return images for the specified block types. By default, no images
	// are returned.
	ReturnImages param.Field[[]ExtractRunJobParamsParsingSettingsReturnImage] `json:"return_images"`
	// If True, return OCR data in the result. Defaults to False.
	ReturnOcrData param.Field[bool] `json:"return_ocr_data"`
	// The timeout for the job in seconds. Defaults to 900.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r ExtractRunJobParamsParsingSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Standard is our best multilingual OCR system. Legacy only supports germanic
// languages and is available for backwards compatibility.
type ExtractRunJobParamsParsingSettingsOcrSystem string

const (
	ExtractRunJobParamsParsingSettingsOcrSystemStandard ExtractRunJobParamsParsingSettingsOcrSystem = "standard"
	ExtractRunJobParamsParsingSettingsOcrSystemLegacy   ExtractRunJobParamsParsingSettingsOcrSystem = "legacy"
)

func (r ExtractRunJobParamsParsingSettingsOcrSystem) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsParsingSettingsOcrSystemStandard, ExtractRunJobParamsParsingSettingsOcrSystemLegacy:
		return true
	}
	return false
}

// The page range to process (1-indexed). By default, the entire document is
// processed.
//
// Satisfied by [shared.PageRangeParam],
// [ExtractRunJobParamsParsingSettingsPageRangeArray],
// [ExtractRunJobParamsParsingSettingsPageRangeArray].
type ExtractRunJobParamsParsingSettingsPageRangeUnion interface {
	ImplementsExtractRunJobParamsParsingSettingsPageRangeUnion()
}

type ExtractRunJobParamsParsingSettingsPageRangeArray []shared.PageRangeParam

func (r ExtractRunJobParamsParsingSettingsPageRangeArray) ImplementsExtractRunJobParamsParsingSettingsPageRangeUnion() {
}

type ExtractRunJobParamsParsingSettingsReturnImage string

const (
	ExtractRunJobParamsParsingSettingsReturnImageFigure ExtractRunJobParamsParsingSettingsReturnImage = "figure"
	ExtractRunJobParamsParsingSettingsReturnImageTable  ExtractRunJobParamsParsingSettingsReturnImage = "table"
)

func (r ExtractRunJobParamsParsingSettingsReturnImage) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsParsingSettingsReturnImageFigure, ExtractRunJobParamsParsingSettingsReturnImageTable:
		return true
	}
	return false
}

type ExtractRunJobParamsParsingSpreadsheet struct {
	// In a spreadsheet with different tables inside, we enable splitting up the tables
	// by default. Accurate mode applies more powerful models for superior accuracy, at
	// 5× the default per-cell rate. Disabling will register as one large table.
	Clustering param.Field[ExtractRunJobParamsParsingSpreadsheetClustering] `json:"clustering"`
	// Whether to exclude hidden sheets, rows, or columns in the output.
	Exclude param.Field[[]ExtractRunJobParamsParsingSpreadsheetExclude] `json:"exclude"`
	// Whether to include cell color and formula information in the output.
	Include          param.Field[[]ExtractRunJobParamsParsingSpreadsheetInclude]        `json:"include"`
	SplitLargeTables param.Field[ExtractRunJobParamsParsingSpreadsheetSplitLargeTables] `json:"split_large_tables"`
}

func (r ExtractRunJobParamsParsingSpreadsheet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// In a spreadsheet with different tables inside, we enable splitting up the tables
// by default. Accurate mode applies more powerful models for superior accuracy, at
// 5× the default per-cell rate. Disabling will register as one large table.
type ExtractRunJobParamsParsingSpreadsheetClustering string

const (
	ExtractRunJobParamsParsingSpreadsheetClusteringAccurate ExtractRunJobParamsParsingSpreadsheetClustering = "accurate"
	ExtractRunJobParamsParsingSpreadsheetClusteringFast     ExtractRunJobParamsParsingSpreadsheetClustering = "fast"
	ExtractRunJobParamsParsingSpreadsheetClusteringDisabled ExtractRunJobParamsParsingSpreadsheetClustering = "disabled"
)

func (r ExtractRunJobParamsParsingSpreadsheetClustering) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsParsingSpreadsheetClusteringAccurate, ExtractRunJobParamsParsingSpreadsheetClusteringFast, ExtractRunJobParamsParsingSpreadsheetClusteringDisabled:
		return true
	}
	return false
}

type ExtractRunJobParamsParsingSpreadsheetExclude string

const (
	ExtractRunJobParamsParsingSpreadsheetExcludeHiddenSheets ExtractRunJobParamsParsingSpreadsheetExclude = "hidden_sheets"
	ExtractRunJobParamsParsingSpreadsheetExcludeHiddenRows   ExtractRunJobParamsParsingSpreadsheetExclude = "hidden_rows"
	ExtractRunJobParamsParsingSpreadsheetExcludeHiddenCols   ExtractRunJobParamsParsingSpreadsheetExclude = "hidden_cols"
)

func (r ExtractRunJobParamsParsingSpreadsheetExclude) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsParsingSpreadsheetExcludeHiddenSheets, ExtractRunJobParamsParsingSpreadsheetExcludeHiddenRows, ExtractRunJobParamsParsingSpreadsheetExcludeHiddenCols:
		return true
	}
	return false
}

type ExtractRunJobParamsParsingSpreadsheetInclude string

const (
	ExtractRunJobParamsParsingSpreadsheetIncludeCellColors ExtractRunJobParamsParsingSpreadsheetInclude = "cell_colors"
	ExtractRunJobParamsParsingSpreadsheetIncludeFormula    ExtractRunJobParamsParsingSpreadsheetInclude = "formula"
)

func (r ExtractRunJobParamsParsingSpreadsheetInclude) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsParsingSpreadsheetIncludeCellColors, ExtractRunJobParamsParsingSpreadsheetIncludeFormula:
		return true
	}
	return false
}

type ExtractRunJobParamsParsingSpreadsheetSplitLargeTables struct {
	// If True, split large tables into smaller tables. Defaults to True.
	Enabled param.Field[bool] `json:"enabled"`
	// The size of the tables to split into. Defaults to 50.
	Size param.Field[int64] `json:"size"`
}

func (r ExtractRunJobParamsParsingSpreadsheetSplitLargeTables) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The settings to use for the extraction.
type ExtractRunJobParamsSettings struct {
	// If True, use array extraction.
	ArrayExtract param.Field[bool] `json:"array_extract"`
	// The citations to use for the extraction.
	Citations param.Field[ExtractRunJobParamsSettingsCitations] `json:"citations"`
	// If True, include images in the extraction.
	IncludeImages param.Field[bool] `json:"include_images"`
	// If True, jobs will be processed with a higher throughput and priority at a
	// higher cost. Defaults to False.
	OptimizeForLatency param.Field[bool] `json:"optimize_for_latency"`
}

func (r ExtractRunJobParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The citations to use for the extraction.
type ExtractRunJobParamsSettingsCitations struct {
	// If True, include citations in the extraction.
	Enabled param.Field[bool] `json:"enabled"`
	// If True, enable numeric citation confidence scores. Defaults to True.
	NumericalConfidence param.Field[bool] `json:"numerical_confidence"`
}

func (r ExtractRunJobParamsSettingsCitations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
