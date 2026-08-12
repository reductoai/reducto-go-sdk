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
	return res, err
}

// Async Parse
func (r *ParseService) RunJob(ctx context.Context, body ParseRunJobParams, opts ...option.RequestOption) (res *shared.AsyncParseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "parse_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AsyncConfigV3Param struct {
	// JSON metadata included in webhook request body. Defaults to None.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Workers poll the priority queue ahead of the standard queue, so priority jobs
	// start sooner when there is queued work; sync jobs are prioritized above async
	// jobs by default.
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

func (r AsyncConfigV3WebhookParam) ImplementsAsyncConfigV3WebhookUnionParam() {}

// The webhook configuration for the asynchronous processing.
//
// Satisfied by [shared.SvixWebhookConfigParam], [shared.DirectWebhookConfigParam],
// [AsyncConfigV3WebhookParam].
type AsyncConfigV3WebhookUnionParam interface {
	ImplementsAsyncConfigV3WebhookUnionParam()
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
	// capacity is available. 'auto' (alias: 'standard') uses the default queue.
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
// capacity is available. 'auto' (alias: 'standard') uses the default queue.
type AsyncParseConfigQueuePriority string

const (
	AsyncParseConfigQueuePriorityAuto     AsyncParseConfigQueuePriority = "auto"
	AsyncParseConfigQueuePriorityStandard AsyncParseConfigQueuePriority = "standard"
	AsyncParseConfigQueuePriorityBatch    AsyncParseConfigQueuePriority = "batch"
)

func (r AsyncParseConfigQueuePriority) IsKnown() bool {
	switch r {
	case AsyncParseConfigQueuePriorityAuto, AsyncParseConfigQueuePriorityStandard, AsyncParseConfigQueuePriorityBatch:
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
	// If True, run advanced chart extraction on figures classified as charts: an
	// agentic extractor that returns full structured series data (chart_data) plus a
	// reconstruction image re-drawn from that data (extra.chart_reconstruction).
	// Higher latency. Defaults to False.
	AdvancedChartAgent param.Field[bool] `json:"advanced_chart_agent"`
	// Mode for table agentic: 'default' selectively applies enrichment only to tables
	// likely to benefit, and 'max' runs enrichment on all tables.
	Mode param.Field[EnhanceAgenticMode] `json:"mode"`
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

// Mode for table agentic: 'default' selectively applies enrichment only to tables
// likely to benefit, and 'max' runs enrichment on all tables.
type EnhanceAgenticMode string

const (
	EnhanceAgenticModeDefault EnhanceAgenticMode = "default"
	EnhanceAgenticModeAuto    EnhanceAgenticMode = "auto"
	EnhanceAgenticModeMax     EnhanceAgenticMode = "max"
)

func (r EnhanceAgenticMode) IsKnown() bool {
	switch r {
	case EnhanceAgenticModeDefault, EnhanceAgenticModeAuto, EnhanceAgenticModeMax:
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

type RetrievalParam struct {
	Chunking param.Field[shared.ChunkingParam] `json:"chunking"`
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
	// Render DPI used when rasterizing the source PDF before embedding the OCR text
	// layer (only applies when `embed_pdf_metadata` is True). Lower values produce
	// dramatically smaller output PDFs; higher values preserve more detail when zoomed
	// past 200%. Defaults to 100 (good for on-screen viewing); raise toward the source
	// scan DPI for crisper output. Min 50, max 250.
	EmbedPdfMetadataDpi param.Field[int64] `json:"embed_pdf_metadata_dpi"`
	// If True, return properties embedded in the original document. Defaults to False.
	ExtractDocumentProperties param.Field[bool] `json:"extract_document_properties"`
	// The mode to use for text extraction from PDFs. OCR mode uses optical character
	// recognition only. Hybrid mode combines OCR with embedded PDF text for best
	// accuracy (default).
	ExtractionMode param.Field[SettingsExtractionMode] `json:"extraction_mode"`
	// Force the URL to be downloaded as a specific file extension (e.g. `.png`).
	ForceFileExtension param.Field[string] `json:"force_file_extension"`
	// Force the result to be returned in URL form.
	ForceURLResult param.Field[bool] `json:"force_url_result"`
	// Hybrid VPC request-scoped settings.
	HybridVpc param.Field[SettingsHybridVpcParam] `json:"hybrid_vpc"`
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
	// Per-tenant throttling for multi-tenant applications. Tag each request with your
	// tenant's id to bound how much of your account's concurrency a single tenant can
	// consume. Account-level throttles still apply.
	TenantThrottling param.Field[SettingsTenantThrottlingParam] `json:"tenant_throttling"`
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

// Hybrid VPC request-scoped settings.
type SettingsHybridVpcParam struct {
	// Named Hybrid VPC environment to use for this request. Only applies when your
	// organization has Hybrid VPC environments configured.
	Environment param.Field[string] `json:"environment"`
}

func (r SettingsHybridVpcParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
// Satisfied by [shared.PageRangeParam], [SettingsPageRangeArrayParam],
// [SettingsPageRangeArrayParam], [SettingsPageRangeArrayParam].
type SettingsPageRangeUnionParam interface {
	ImplementsSettingsPageRangeUnionParam()
}

type SettingsPageRangeArrayParam []shared.PageRangeParam

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

// Per-tenant throttling for multi-tenant applications. Tag each request with your
// tenant's id to bound how much of your account's concurrency a single tenant can
// consume. Account-level throttles still apply.
type SettingsTenantThrottlingParam struct {
	// Your identifier for the tenant (customer, workspace, organization) this request
	// belongs to. Used only for noisy-neighbor throttling inside your account.
	TenantID param.Field[string] `json:"tenant_id" api:"required"`
	// Maximum fraction of your account's concurrency ceiling this tenant may use,
	// between 0 (exclusive) and 1. Defaults to 0.5.
	MaxShare param.Field[float64] `json:"max_share"`
}

func (r SettingsTenantThrottlingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SpreadsheetParam struct {
	// In a spreadsheet with different tables inside, we enable splitting up the tables
	// by default. Accurate mode applies more powerful models for superior accuracy, at
	// 5× the default per-cell rate. Disabling will register as one large table.
	Clustering param.Field[SpreadsheetClustering] `json:"clustering"`
	// Whether to exclude hidden sheets, rows, or columns in the output.
	Exclude param.Field[[]SpreadsheetExclude] `json:"exclude"`
	// Whether to include cell color, formula, and dropdown information in the output.
	Include param.Field[[]SpreadsheetInclude] `json:"include"`
	// Maximum total non-empty cells allowed across all sheets. If exceeded, the
	// request is rejected with a 422 error. Set to null to disable the limit. Defaults
	// to null.
	MaxCellCount     param.Field[int64]                        `json:"max_cell_count"`
	SplitLargeTables param.Field[shared.SplitLargeTablesParam] `json:"split_large_tables"`
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

type ParseRunResponse struct {
	JobID string `json:"job_id" api:"required"`
	// This field can have the runtime type of
	// [shared.ParseResponseDocumentProperties].
	DocumentProperties interface{} `json:"document_properties"`
	// The duration of the parse request in seconds.
	Duration float64 `json:"duration"`
	// The storage URL of the converted PDF file.
	PdfURL       string                       `json:"pdf_url" api:"nullable"`
	ResponseType ParseRunResponseResponseType `json:"response_type"`
	// This field can have the runtime type of [shared.ParseResponseResult].
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
	JobID              apijson.Field
	DocumentProperties apijson.Field
	Duration           apijson.Field
	PdfURL             apijson.Field
	ResponseType       apijson.Field
	Result             apijson.Field
	StudioLink         apijson.Field
	Usage              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
// [shared.AsyncParseResponse].
func (r ParseRunResponse) AsUnion() ParseRunResponseUnion {
	return r.union
}

// Union satisfied by [shared.ParseResponse] or [shared.AsyncParseResponse].
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
			Type:       reflect.TypeOf(shared.AsyncParseResponse{}),
		},
	)
}

type ParseRunResponseResponseType string

const (
	ParseRunResponseResponseTypeParse ParseRunResponseResponseType = "parse"
)

func (r ParseRunResponseResponseType) IsKnown() bool {
	switch r {
	case ParseRunResponseResponseTypeParse:
		return true
	}
	return false
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
	// capacity is available. 'auto' (alias: 'standard') uses the default queue.
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
// capacity is available. 'auto' (alias: 'standard') uses the default queue.
type ParseRunParamsBodyQueuePriority string

const (
	ParseRunParamsBodyQueuePriorityAuto     ParseRunParamsBodyQueuePriority = "auto"
	ParseRunParamsBodyQueuePriorityStandard ParseRunParamsBodyQueuePriority = "standard"
	ParseRunParamsBodyQueuePriorityBatch    ParseRunParamsBodyQueuePriority = "batch"
)

func (r ParseRunParamsBodyQueuePriority) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyQueuePriorityAuto, ParseRunParamsBodyQueuePriorityStandard, ParseRunParamsBodyQueuePriorityBatch:
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
