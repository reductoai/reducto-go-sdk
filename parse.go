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
	// Configuration for parsing a document
	Body ParseRunParamsBodyUnion `json:"body,required"`
}

func (r ParseRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Configuration for parsing a document
type ParseRunParamsBody struct {
	AdvancedOptions     param.Field[shared.AdvancedProcessingOptionsParam]     `json:"advanced_options"`
	Async               param.Field[interface{}]                               `json:"async"`
	Config              param.Field[interface{}]                               `json:"config"`
	DocumentURL         param.Field[interface{}]                               `json:"document_url"`
	Enhance             param.Field[interface{}]                               `json:"enhance"`
	ExperimentalOptions param.Field[shared.ExperimentalProcessingOptionsParam] `json:"experimental_options"`
	Formatting          param.Field[interface{}]                               `json:"formatting"`
	Input               param.Field[interface{}]                               `json:"input"`
	Options             param.Field[shared.BaseProcessingOptionsParam]         `json:"options"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority    param.Field[bool]        `json:"priority"`
	Retrieval   param.Field[interface{}] `json:"retrieval"`
	Settings    param.Field[interface{}] `json:"settings"`
	Spreadsheet param.Field[interface{}] `json:"spreadsheet"`
	UserConfig  param.Field[interface{}] `json:"user_config"`
}

func (r ParseRunParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBody) implementsParseRunParamsBodyUnion() {}

// Configuration for parsing a document
//
// Satisfied by [ParseConfigParam], [ParseRunParamsBodyParseConfig],
// [ParseRunParamsBodySyncParseConfig], [ParseRunParamsBody].
type ParseRunParamsBodyUnion interface {
	implementsParseRunParamsBodyUnion()
}

// Configuration for parsing a document
type ParseRunParamsBodyParseConfig struct {
	// The URL of the document to process. Either a public URL or a presigned URL with
	// a valid expiration time.
	DocumentURL param.Field[ParseRunParamsBodyParseConfigDocumentURLUnion] `json:"document_url,required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async param.Field[ParseRunParamsBodyParseConfigAsync] `json:"async"`
	// The configuration options for processing the document.
	Config param.Field[ParseRunParamsBodyParseConfigConfig] `json:"config"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool] `json:"priority"`
	// User-specific configuration options.
	UserConfig param.Field[map[string]interface{}] `json:"user_config"`
}

func (r ParseRunParamsBodyParseConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBodyParseConfig) implementsParseRunParamsBodyUnion() {}

// The URL of the document to process. Either a public URL or a presigned URL with
// a valid expiration time.
//
// Satisfied by [shared.UnionString],
// [ParseRunParamsBodyParseConfigDocumentURLArray].
type ParseRunParamsBodyParseConfigDocumentURLUnion interface {
	ImplementsParseRunParamsBodyParseConfigDocumentURLUnion()
}

type ParseRunParamsBodyParseConfigDocumentURLArray []string

func (r ParseRunParamsBodyParseConfigDocumentURLArray) ImplementsParseRunParamsBodyParseConfigDocumentURLUnion() {
}

// The configuration options for asynchronous processing (default synchronous).
type ParseRunParamsBodyParseConfigAsync struct {
	Enabled  param.Field[bool]                                      `json:"enabled"`
	Priority param.Field[bool]                                      `json:"priority"`
	Webhook  param.Field[ParseRunParamsBodyParseConfigAsyncWebhook] `json:"webhook"`
}

func (r ParseRunParamsBodyParseConfigAsync) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ParseRunParamsBodyParseConfigAsyncWebhook struct {
	// A list of Svix channels the message will be delivered down, omit to send to all
	// channels.
	Channels param.Field[[]string]                                      `json:"channels"`
	Metadata param.Field[interface{}]                                   `json:"metadata"`
	Mode     param.Field[ParseRunParamsBodyParseConfigAsyncWebhookMode] `json:"mode"`
	URL      param.Field[string]                                        `json:"url"`
}

func (r ParseRunParamsBodyParseConfigAsyncWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ParseRunParamsBodyParseConfigAsyncWebhookMode string

const (
	ParseRunParamsBodyParseConfigAsyncWebhookModeDirect ParseRunParamsBodyParseConfigAsyncWebhookMode = "direct"
	ParseRunParamsBodyParseConfigAsyncWebhookModeSvix   ParseRunParamsBodyParseConfigAsyncWebhookMode = "svix"
)

func (r ParseRunParamsBodyParseConfigAsyncWebhookMode) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyParseConfigAsyncWebhookModeDirect, ParseRunParamsBodyParseConfigAsyncWebhookModeSvix:
		return true
	}
	return false
}

// The configuration options for processing the document.
type ParseRunParamsBodyParseConfigConfig struct {
	// If True, add page markers to the output. Defaults to False.
	AddPageMarkers param.Field[bool] `json:"add_page_markers"`
	// If enabled, a large language/vision model will be used to postprocess the layout
	// predictions. Defaults to False.
	BetaLayoutEnrichment param.Field[bool] `json:"beta_layout_enrichment"`
	// The name of the bucket to use for the document.
	BucketName param.Field[string] `json:"bucket_name"`
	// A flag to indicate if bar chart extraction should be performed (requires
	// figure_summary=True). Defaults to False.
	ChartExtract param.Field[bool] `json:"chart_extract"`
	// The mode to use for chunking. Defaults to 'variable'. Section chunks according
	// to sections in the document. Page chunks according to pages. Page sections
	// chunks according to both pages and sections. Disabled returns a single chunk.
	ChunkMode param.Field[ParseRunParamsBodyParseConfigConfigChunkMode] `json:"chunk_mode"`
	// The approximate size of chunks (in characters) that the document will be split
	// into. Defaults to None, in which case the chunk size is variable between 250 -
	// 1500 characters.
	ChunkSize param.Field[int64] `json:"chunk_size"`
	// A flag to indicate if the hierarchy of the document should be continued from
	// chunk to chunk. E.g. ## Prev Section (cont.)
	ContinueHierarchy param.Field[bool]                                            `json:"continue_hierarchy"`
	CustomFormat      param.Field[ParseRunParamsBodyParseConfigConfigCustomFormat] `json:"custom_format"`
	// Override the customer ID for the request. Defaults to None.
	CustomerID param.Field[string] `json:"customer_id"`
	// If True, filter out boxes with width greater than 50% of the document width.
	// Defaults to False. You probably don't want to use this.
	DangerFilterWideBoxes param.Field[bool] `json:"danger_filter_wide_boxes"`
	// If True, detect signatures in the document. Defaults to False.
	DetectSignatures param.Field[bool] `json:"detect_signatures"`
	// DEPRECATED, use chunk_mode=disabled instead
	DisableChunking param.Field[bool] `json:"disable_chunking"`
	// Password to decrypt password-protected documents.
	DocumentPassword param.Field[string] `json:"document_password"`
	// The dots per inch (DPI) setting for image processing. Higher values increase
	// resolution but increase latency. The maximum recommended value is 300.
	Dpi param.Field[int64] `json:"dpi"`
	// If True, embed text metadata into the returned PDF. Defaults to False.
	EmbedTextMetadataPdf param.Field[bool] `json:"embed_text_metadata_pdf"`
	// Add <u> tag around text that's underlined and surround strikethroughs and
	// underlines with <change> tags, defaults to False
	EnableChangeTracking param.Field[bool] `json:"enable_change_tracking"`
	// Pull PDF comments from the document, defaults to False
	EnableComments param.Field[bool] `json:"enable_comments"`
	// Add <mark> tags around highlighted text detected using the segmentation model,
	// defaults to False
	EnableHighlightDetection param.Field[bool] `json:"enable_highlight_detection"`
	// Add <sub> tag around subscripts and <sup> tag around superscripts, defaults to
	// False
	EnableScripts param.Field[bool] `json:"enable_scripts"`
	// Instead of using LibreOffice, when enabled, this flag uses a Windows VM to
	// convert docx files. This is slower but more accurate.
	EnhancedDocxConversion param.Field[bool] `json:"enhanced_docx_conversion"`
	// If True, use enhanced figure summarization pipeline for complex charts. Defaults
	// to False.
	EnhancedFigureSummary param.Field[bool] `json:"enhanced_figure_summary"`
	// If enabled, a large language/vision model will be used to postprocess the
	// extracted content. Defaults to False.
	Enrich param.Field[bool] `json:"enrich"`
	// The mode to use for enrichment. Defaults to standard
	EnrichMode param.Field[ParseRunParamsBodyParseConfigConfigEnrichMode] `json:"enrich_mode"`
	// Add information to the prompt for enrichment.
	EnrichPrompt param.Field[string] `json:"enrich_prompt"`
	// Skip hidden rows and cols in Excel files. Defaults to False.
	ExcludeHiddenRowsCols param.Field[bool] `json:"exclude_hidden_rows_cols"`
	// Skip hidden sheets in Excel files. Defaults to False.
	ExcludeHiddenSheets param.Field[bool] `json:"exclude_hidden_sheets"`
	// Note, this is an alpha feature subject to change at any time. If enabled, large
	// tables will be chunked into multiple tables. Defaults to False.
	ExperimentalLargeSpreadsheetTableChunking param.Field[bool] `json:"experimental_large_spreadsheet_table_chunking"`
	// Extra metadata to be added to logs.
	ExtraMetadata param.Field[map[string]interface{}] `json:"extra_metadata"`
	// A flag to indicate if figure summarization should be performed. Defaults to
	// False.
	FigureSummary param.Field[bool] `json:"figure_summary"`
	// If the figure summary prompt should override our default prompt.
	FigureSummaryOverride param.Field[bool] `json:"figure_summary_override"`
	// Add information to the prompt for figure summarization.
	FigureSummaryPrompt param.Field[string] `json:"figure_summary_prompt"`
	// If True, filter out line numbers from the output. Defaults to False.
	FilterLineNumbers param.Field[bool] `json:"filter_line_numbers"`
	// Force the URL to be downloaded as a specific file extension (e.g. `.png`).
	ForceFileExtension param.Field[string] `json:"force_file_extension"`
	// Force the result to be returned in URL form.
	ForceURLResult param.Field[bool] `json:"force_url_result"`
	// A list of block types to ignore. Defaults to ['Page Number', 'Header', 'Footer',
	// 'Comment'].
	IgnoreBlocks param.Field[[]ParseRunParamsBodyParseConfigConfigIgnoreBlock] `json:"ignore_blocks"`
	// If True, preserve color information in spreadsheet cells by wrapping text with
	// LaTeX color commands during parsing.
	IncludeColorInformation param.Field[bool] `json:"include_color_information"`
	// If True, preserve formula information in spreadsheet cells by wrapping text with
	// LaTeX formula commands during parsing.
	IncludeFormulaInformation param.Field[bool] `json:"include_formula_information"`
	// If table cell colors should be used to determine table structure. Defaults to
	// False.
	InferTableColor param.Field[bool] `json:"infer_table_color"`
	// If True, include bounding box information in JSON table output. Defaults to
	// False.
	JsonBbox param.Field[bool] `json:"json_bbox"`
	// If line breaks should be preserved in the text. Defaults to False.
	KeepLineBreaks param.Field[bool] `json:"keep_line_breaks"`
	// The AWS KMS key to use for the document.
	KmsArn param.Field[string] `json:"kms_arn"`
	// If large tables should be chunked into smaller tables, currently only supported
	// on spreadsheet and CSV files.
	LargeTableChunking param.Field[bool] `json:"large_table_chunking"`
	// The max row/column size for a table to be chunked. Defaults to 50.
	LargeTableChunkingSize param.Field[int64] `json:"large_table_chunking_size"`
	// The layout model to use for the document. This will be deprecated in the future.
	LayoutModel param.Field[ParseRunParamsBodyParseConfigConfigLayoutModel] `json:"layout_model"`
	// The maximum number of pages to process in a single batch. Defaults to 10.
	MaxBatchSize param.Field[int64] `json:"max_batch_size"`
	// A flag to indicate if consecutive tables with the same number of columns should
	// be merged. Defaults to False.
	MergeTables param.Field[bool] `json:"merge_tables"`
	// The type of document to be processed. Defaults to document. If auto is
	// specified, the orientation of the first page will be used to determine the
	// document type.
	Mode param.Field[ParseRunParamsBodyParseConfigConfigMode] `json:"mode"`
	// The dimension of the OCR crops along each axis. num_ocr_crops^2 is the total
	// number of crops. Defaults to 2.
	NumOcrCrops param.Field[int64] `json:"num_ocr_crops"`
	// If True, enable numeric parse confidence scores in granular_confidence field.
	// Defaults to False.
	NumericalParseConfidence param.Field[bool] `json:"numerical_parse_confidence"`
	// The mode to use for OCR. If agentic is enabled, table OCR will be automatically
	// edited.
	OcrMode param.Field[ParseRunParamsBodyParseConfigConfigOcrMode] `json:"ocr_mode"`
	// The OCR system to use. Defaults to cloud (AWS Textract/Azure DocAI/etc).
	OcrSystem param.Field[ParseRunParamsBodyParseConfigConfigOcrSystem] `json:"ocr_system"`
	// The threshold for box overlap. Defaults to 0.5.
	OverlapThreshold param.Field[float64] `json:"overlap_threshold"`
	// The page number to stop processing at.
	PageEnd param.Field[int64] `json:"page_end"`
	// The page range to process.
	PageRange param.Field[ParseRunParamsBodyParseConfigConfigPageRangeUnion] `json:"page_range"`
	// The page number to start processing from.
	PageStart param.Field[int64] `json:"page_start"`
	// The method to use for OCR. hybrid uses the PDF text first, then OCR. pdf only
	// uses the PDF text. ocr only uses OCR.
	PdfOcr param.Field[ParseRunParamsBodyParseConfigConfigPdfOcr] `json:"pdf_ocr"`
	// If True, persist the results indefinitely. Defaults to False.
	PersistResults param.Field[bool] `json:"persist_results"`
	// Forces all external API calls to be routed to specified country/region.
	RegionPreference param.Field[ParseRunParamsBodyParseConfigConfigRegionPreference] `json:"region_preference"`
	// If True, remove text formatting from the output (e.g. hyphens for list items).
	// Defaults to False.
	RemoveTextFormatting param.Field[bool] `json:"remove_text_formatting"`
	// If figure images should be returned in the result. Defaults to False.
	ReturnFigureImages param.Field[bool] `json:"return_figure_images"`
	// If True, return OCR data in the result. Defaults to False.
	ReturnOcrData param.Field[bool] `json:"return_ocr_data"`
	// If table images should be returned in the result. Defaults to False.
	ReturnTableImages param.Field[bool] `json:"return_table_images"`
	// Use an orientation model to detect and rotate figures as needed, defaults to
	// False
	RotateFigures param.Field[bool] `json:"rotate_figures"`
	// Use an orientation model to detect and rotate pages as needed, defaults to False
	RotatePages param.Field[bool] `json:"rotate_pages"`
	// On a spreadsheet, the algorithm that is used to split up sheets into multiple
	// tables.
	SpreadsheetTableClustering param.Field[ParseRunParamsBodyParseConfigConfigSpreadsheetTableClustering] `json:"spreadsheet_table_clustering"`
	// If True, enable figure summaries for all figures regardless of size (onprem
	// only). Defaults to False.
	SummarizeAllFigures param.Field[bool] `json:"summarize_all_figures"`
	// The mode to use for table output. Defaults to html.
	TableOutputFormat param.Field[ParseRunParamsBodyParseConfigConfigTableOutputFormat] `json:"table_output_format"`
	// If tables should be summarized for embedding. Defaults to True.
	TableSummary param.Field[bool] `json:"table_summary"`
	// Add information to the prompt for table summarization.
	TableSummaryPrompt param.Field[string] `json:"table_summary_prompt"`
	// LEGACY: For sync/on-prem only. The timeout for the job in seconds. Defaults
	// to 1800.
	Timeout param.Field[float64] `json:"timeout"`
	// Add checkboxes to the output, defaults to False
	UseCheckboxes param.Field[bool] `json:"use_checkboxes"`
	// Add equations to the output, defaults to False
	UseEquations param.Field[bool] `json:"use_equations"`
	// Use a faster inference model for parsing. Defaults to False.
	UseFastInference param.Field[bool] `json:"use_fast_inference"`
	// Use GPU acceleration for OCR processing. Defaults to False.
	UseGPUOcr param.Field[bool] `json:"use_gpu_ocr"`
	// A user specified timeout, defaults to None
	UserSpecifiedTimeoutSeconds param.Field[float64] `json:"user_specified_timeout_seconds"`
	// The version of the processing options.
	Version param.Field[ParseRunParamsBodyParseConfigConfigVersion] `json:"version"`
}

func (r ParseRunParamsBodyParseConfigConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The mode to use for chunking. Defaults to 'variable'. Section chunks according
// to sections in the document. Page chunks according to pages. Page sections
// chunks according to both pages and sections. Disabled returns a single chunk.
type ParseRunParamsBodyParseConfigConfigChunkMode string

const (
	ParseRunParamsBodyParseConfigConfigChunkModeVariable     ParseRunParamsBodyParseConfigConfigChunkMode = "variable"
	ParseRunParamsBodyParseConfigConfigChunkModeSection      ParseRunParamsBodyParseConfigConfigChunkMode = "section"
	ParseRunParamsBodyParseConfigConfigChunkModePage         ParseRunParamsBodyParseConfigConfigChunkMode = "page"
	ParseRunParamsBodyParseConfigConfigChunkModeDisabled     ParseRunParamsBodyParseConfigConfigChunkMode = "disabled"
	ParseRunParamsBodyParseConfigConfigChunkModeBlock        ParseRunParamsBodyParseConfigConfigChunkMode = "block"
	ParseRunParamsBodyParseConfigConfigChunkModePageSections ParseRunParamsBodyParseConfigConfigChunkMode = "page_sections"
)

func (r ParseRunParamsBodyParseConfigConfigChunkMode) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyParseConfigConfigChunkModeVariable, ParseRunParamsBodyParseConfigConfigChunkModeSection, ParseRunParamsBodyParseConfigConfigChunkModePage, ParseRunParamsBodyParseConfigConfigChunkModeDisabled, ParseRunParamsBodyParseConfigConfigChunkModeBlock, ParseRunParamsBodyParseConfigConfigChunkModePageSections:
		return true
	}
	return false
}

type ParseRunParamsBodyParseConfigConfigCustomFormat string

const (
	ParseRunParamsBodyParseConfigConfigCustomFormatAml     ParseRunParamsBodyParseConfigConfigCustomFormat = "aml"
	ParseRunParamsBodyParseConfigConfigCustomFormatAIUsage ParseRunParamsBodyParseConfigConfigCustomFormat = "ai_usage"
)

func (r ParseRunParamsBodyParseConfigConfigCustomFormat) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyParseConfigConfigCustomFormatAml, ParseRunParamsBodyParseConfigConfigCustomFormatAIUsage:
		return true
	}
	return false
}

// The mode to use for enrichment. Defaults to standard
type ParseRunParamsBodyParseConfigConfigEnrichMode string

const (
	ParseRunParamsBodyParseConfigConfigEnrichModeStandard ParseRunParamsBodyParseConfigConfigEnrichMode = "standard"
	ParseRunParamsBodyParseConfigConfigEnrichModePage     ParseRunParamsBodyParseConfigConfigEnrichMode = "page"
	ParseRunParamsBodyParseConfigConfigEnrichModeTable    ParseRunParamsBodyParseConfigConfigEnrichMode = "table"
)

func (r ParseRunParamsBodyParseConfigConfigEnrichMode) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyParseConfigConfigEnrichModeStandard, ParseRunParamsBodyParseConfigConfigEnrichModePage, ParseRunParamsBodyParseConfigConfigEnrichModeTable:
		return true
	}
	return false
}

type ParseRunParamsBodyParseConfigConfigIgnoreBlock string

const (
	ParseRunParamsBodyParseConfigConfigIgnoreBlockHeader        ParseRunParamsBodyParseConfigConfigIgnoreBlock = "Header"
	ParseRunParamsBodyParseConfigConfigIgnoreBlockFooter        ParseRunParamsBodyParseConfigConfigIgnoreBlock = "Footer"
	ParseRunParamsBodyParseConfigConfigIgnoreBlockTitle         ParseRunParamsBodyParseConfigConfigIgnoreBlock = "Title"
	ParseRunParamsBodyParseConfigConfigIgnoreBlockSectionHeader ParseRunParamsBodyParseConfigConfigIgnoreBlock = "Section Header"
	ParseRunParamsBodyParseConfigConfigIgnoreBlockPageNumber    ParseRunParamsBodyParseConfigConfigIgnoreBlock = "Page Number"
	ParseRunParamsBodyParseConfigConfigIgnoreBlockListItem      ParseRunParamsBodyParseConfigConfigIgnoreBlock = "List Item"
	ParseRunParamsBodyParseConfigConfigIgnoreBlockFigure        ParseRunParamsBodyParseConfigConfigIgnoreBlock = "Figure"
	ParseRunParamsBodyParseConfigConfigIgnoreBlockTable         ParseRunParamsBodyParseConfigConfigIgnoreBlock = "Table"
	ParseRunParamsBodyParseConfigConfigIgnoreBlockKeyValue      ParseRunParamsBodyParseConfigConfigIgnoreBlock = "Key Value"
	ParseRunParamsBodyParseConfigConfigIgnoreBlockText          ParseRunParamsBodyParseConfigConfigIgnoreBlock = "Text"
	ParseRunParamsBodyParseConfigConfigIgnoreBlockComment       ParseRunParamsBodyParseConfigConfigIgnoreBlock = "Comment"
	ParseRunParamsBodyParseConfigConfigIgnoreBlockSignature     ParseRunParamsBodyParseConfigConfigIgnoreBlock = "Signature"
)

func (r ParseRunParamsBodyParseConfigConfigIgnoreBlock) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyParseConfigConfigIgnoreBlockHeader, ParseRunParamsBodyParseConfigConfigIgnoreBlockFooter, ParseRunParamsBodyParseConfigConfigIgnoreBlockTitle, ParseRunParamsBodyParseConfigConfigIgnoreBlockSectionHeader, ParseRunParamsBodyParseConfigConfigIgnoreBlockPageNumber, ParseRunParamsBodyParseConfigConfigIgnoreBlockListItem, ParseRunParamsBodyParseConfigConfigIgnoreBlockFigure, ParseRunParamsBodyParseConfigConfigIgnoreBlockTable, ParseRunParamsBodyParseConfigConfigIgnoreBlockKeyValue, ParseRunParamsBodyParseConfigConfigIgnoreBlockText, ParseRunParamsBodyParseConfigConfigIgnoreBlockComment, ParseRunParamsBodyParseConfigConfigIgnoreBlockSignature:
		return true
	}
	return false
}

// The layout model to use for the document. This will be deprecated in the future.
type ParseRunParamsBodyParseConfigConfigLayoutModel string

const (
	ParseRunParamsBodyParseConfigConfigLayoutModelDefault ParseRunParamsBodyParseConfigConfigLayoutModel = "default"
	ParseRunParamsBodyParseConfigConfigLayoutModelBeta    ParseRunParamsBodyParseConfigConfigLayoutModel = "beta"
)

func (r ParseRunParamsBodyParseConfigConfigLayoutModel) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyParseConfigConfigLayoutModelDefault, ParseRunParamsBodyParseConfigConfigLayoutModelBeta:
		return true
	}
	return false
}

// The type of document to be processed. Defaults to document. If auto is
// specified, the orientation of the first page will be used to determine the
// document type.
type ParseRunParamsBodyParseConfigConfigMode string

const (
	ParseRunParamsBodyParseConfigConfigModeDocument ParseRunParamsBodyParseConfigConfigMode = "document"
	ParseRunParamsBodyParseConfigConfigModeDeck     ParseRunParamsBodyParseConfigConfigMode = "deck"
	ParseRunParamsBodyParseConfigConfigModeAuto     ParseRunParamsBodyParseConfigConfigMode = "auto"
)

func (r ParseRunParamsBodyParseConfigConfigMode) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyParseConfigConfigModeDocument, ParseRunParamsBodyParseConfigConfigModeDeck, ParseRunParamsBodyParseConfigConfigModeAuto:
		return true
	}
	return false
}

// The mode to use for OCR. If agentic is enabled, table OCR will be automatically
// edited.
type ParseRunParamsBodyParseConfigConfigOcrMode string

const (
	ParseRunParamsBodyParseConfigConfigOcrModeStandard ParseRunParamsBodyParseConfigConfigOcrMode = "standard"
	ParseRunParamsBodyParseConfigConfigOcrModeAgentic  ParseRunParamsBodyParseConfigConfigOcrMode = "agentic"
)

func (r ParseRunParamsBodyParseConfigConfigOcrMode) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyParseConfigConfigOcrModeStandard, ParseRunParamsBodyParseConfigConfigOcrModeAgentic:
		return true
	}
	return false
}

// The OCR system to use. Defaults to cloud (AWS Textract/Azure DocAI/etc).
type ParseRunParamsBodyParseConfigConfigOcrSystem string

const (
	ParseRunParamsBodyParseConfigConfigOcrSystemGcloud    ParseRunParamsBodyParseConfigConfigOcrSystem = "gcloud"
	ParseRunParamsBodyParseConfigConfigOcrSystemTextract  ParseRunParamsBodyParseConfigConfigOcrSystem = "textract"
	ParseRunParamsBodyParseConfigConfigOcrSystemTesseract ParseRunParamsBodyParseConfigConfigOcrSystem = "tesseract"
	ParseRunParamsBodyParseConfigConfigOcrSystemCombined  ParseRunParamsBodyParseConfigConfigOcrSystem = "combined"
)

func (r ParseRunParamsBodyParseConfigConfigOcrSystem) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyParseConfigConfigOcrSystemGcloud, ParseRunParamsBodyParseConfigConfigOcrSystemTextract, ParseRunParamsBodyParseConfigConfigOcrSystemTesseract, ParseRunParamsBodyParseConfigConfigOcrSystemCombined:
		return true
	}
	return false
}

// The page range to process.
//
// Satisfied by [shared.PageRangeParam],
// [ParseRunParamsBodyParseConfigConfigPageRangeArray],
// [ParseRunParamsBodyParseConfigConfigPageRangeArray].
type ParseRunParamsBodyParseConfigConfigPageRangeUnion interface {
	ImplementsParseRunParamsBodyParseConfigConfigPageRangeUnion()
}

type ParseRunParamsBodyParseConfigConfigPageRangeArray []shared.PageRangeParam

func (r ParseRunParamsBodyParseConfigConfigPageRangeArray) ImplementsParseRunParamsBodyParseConfigConfigPageRangeUnion() {
}

// The method to use for OCR. hybrid uses the PDF text first, then OCR. pdf only
// uses the PDF text. ocr only uses OCR.
type ParseRunParamsBodyParseConfigConfigPdfOcr string

const (
	ParseRunParamsBodyParseConfigConfigPdfOcrHybrid ParseRunParamsBodyParseConfigConfigPdfOcr = "hybrid"
	ParseRunParamsBodyParseConfigConfigPdfOcrPdf    ParseRunParamsBodyParseConfigConfigPdfOcr = "pdf"
	ParseRunParamsBodyParseConfigConfigPdfOcrOcr    ParseRunParamsBodyParseConfigConfigPdfOcr = "ocr"
)

func (r ParseRunParamsBodyParseConfigConfigPdfOcr) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyParseConfigConfigPdfOcrHybrid, ParseRunParamsBodyParseConfigConfigPdfOcrPdf, ParseRunParamsBodyParseConfigConfigPdfOcrOcr:
		return true
	}
	return false
}

// Forces all external API calls to be routed to specified country/region.
type ParseRunParamsBodyParseConfigConfigRegionPreference string

const (
	ParseRunParamsBodyParseConfigConfigRegionPreferenceUs ParseRunParamsBodyParseConfigConfigRegionPreference = "us"
)

func (r ParseRunParamsBodyParseConfigConfigRegionPreference) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyParseConfigConfigRegionPreferenceUs:
		return true
	}
	return false
}

// On a spreadsheet, the algorithm that is used to split up sheets into multiple
// tables.
type ParseRunParamsBodyParseConfigConfigSpreadsheetTableClustering string

const (
	ParseRunParamsBodyParseConfigConfigSpreadsheetTableClusteringDefault     ParseRunParamsBodyParseConfigConfigSpreadsheetTableClustering = "default"
	ParseRunParamsBodyParseConfigConfigSpreadsheetTableClusteringDisabled    ParseRunParamsBodyParseConfigConfigSpreadsheetTableClustering = "disabled"
	ParseRunParamsBodyParseConfigConfigSpreadsheetTableClusteringIntelligent ParseRunParamsBodyParseConfigConfigSpreadsheetTableClustering = "intelligent"
)

func (r ParseRunParamsBodyParseConfigConfigSpreadsheetTableClustering) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyParseConfigConfigSpreadsheetTableClusteringDefault, ParseRunParamsBodyParseConfigConfigSpreadsheetTableClusteringDisabled, ParseRunParamsBodyParseConfigConfigSpreadsheetTableClusteringIntelligent:
		return true
	}
	return false
}

// The mode to use for table output. Defaults to html.
type ParseRunParamsBodyParseConfigConfigTableOutputFormat string

const (
	ParseRunParamsBodyParseConfigConfigTableOutputFormatHTML    ParseRunParamsBodyParseConfigConfigTableOutputFormat = "html"
	ParseRunParamsBodyParseConfigConfigTableOutputFormatJson    ParseRunParamsBodyParseConfigConfigTableOutputFormat = "json"
	ParseRunParamsBodyParseConfigConfigTableOutputFormatMd      ParseRunParamsBodyParseConfigConfigTableOutputFormat = "md"
	ParseRunParamsBodyParseConfigConfigTableOutputFormatDynamic ParseRunParamsBodyParseConfigConfigTableOutputFormat = "dynamic"
	ParseRunParamsBodyParseConfigConfigTableOutputFormatAIJson  ParseRunParamsBodyParseConfigConfigTableOutputFormat = "ai_json"
	ParseRunParamsBodyParseConfigConfigTableOutputFormatCsv     ParseRunParamsBodyParseConfigConfigTableOutputFormat = "csv"
)

func (r ParseRunParamsBodyParseConfigConfigTableOutputFormat) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyParseConfigConfigTableOutputFormatHTML, ParseRunParamsBodyParseConfigConfigTableOutputFormatJson, ParseRunParamsBodyParseConfigConfigTableOutputFormatMd, ParseRunParamsBodyParseConfigConfigTableOutputFormatDynamic, ParseRunParamsBodyParseConfigConfigTableOutputFormatAIJson, ParseRunParamsBodyParseConfigConfigTableOutputFormatCsv:
		return true
	}
	return false
}

// The version of the processing options.
type ParseRunParamsBodyParseConfigConfigVersion string

const (
	ParseRunParamsBodyParseConfigConfigVersionV1 ParseRunParamsBodyParseConfigConfigVersion = "v1"
	ParseRunParamsBodyParseConfigConfigVersionV2 ParseRunParamsBodyParseConfigConfigVersion = "v2"
	ParseRunParamsBodyParseConfigConfigVersionV3 ParseRunParamsBodyParseConfigConfigVersion = "v3"
)

func (r ParseRunParamsBodyParseConfigConfigVersion) IsKnown() bool {
	switch r {
	case ParseRunParamsBodyParseConfigConfigVersionV1, ParseRunParamsBodyParseConfigConfigVersionV2, ParseRunParamsBodyParseConfigConfigVersionV3:
		return true
	}
	return false
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

type ParseRunJobParams struct {
	Body ParseRunJobParamsBodyUnion `json:"body,required"`
}

func (r ParseRunJobParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ParseRunJobParamsBody struct {
	AdvancedOptions     param.Field[shared.AdvancedProcessingOptionsParam]     `json:"advanced_options"`
	Async               param.Field[interface{}]                               `json:"async"`
	DocumentURL         param.Field[interface{}]                               `json:"document_url"`
	Enhance             param.Field[interface{}]                               `json:"enhance"`
	ExperimentalOptions param.Field[shared.ExperimentalProcessingOptionsParam] `json:"experimental_options"`
	Formatting          param.Field[interface{}]                               `json:"formatting"`
	Input               param.Field[interface{}]                               `json:"input"`
	Options             param.Field[shared.BaseProcessingOptionsParam]         `json:"options"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority    param.Field[bool]                         `json:"priority"`
	Retrieval   param.Field[interface{}]                  `json:"retrieval"`
	Settings    param.Field[interface{}]                  `json:"settings"`
	Spreadsheet param.Field[interface{}]                  `json:"spreadsheet"`
	Webhook     param.Field[shared.WebhookConfigNewParam] `json:"webhook"`
}

func (r ParseRunJobParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunJobParamsBody) implementsParseRunJobParamsBodyUnion() {}

// Satisfied by [ParseRunJobParamsBodyAsyncParseConfigNew],
// [ParseRunJobParamsBodyAsyncParseConfig], [ParseRunJobParamsBody].
type ParseRunJobParamsBodyUnion interface {
	implementsParseRunJobParamsBodyUnion()
}

type ParseRunJobParamsBodyAsyncParseConfigNew struct {
	// The URL of the document to be processed. You can provide one of the following:
	//
	//  1. A publicly available URL
	//  2. A presigned S3 URL
	//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
	//     uploading a document
	DocumentURL         param.Field[ParseRunJobParamsBodyAsyncParseConfigNewDocumentURLUnion] `json:"document_url,required"`
	AdvancedOptions     param.Field[shared.AdvancedProcessingOptionsParam]                    `json:"advanced_options"`
	ExperimentalOptions param.Field[shared.ExperimentalProcessingOptionsParam]                `json:"experimental_options"`
	Options             param.Field[shared.BaseProcessingOptionsParam]                        `json:"options"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool]                         `json:"priority"`
	Webhook  param.Field[shared.WebhookConfigNewParam] `json:"webhook"`
}

func (r ParseRunJobParamsBodyAsyncParseConfigNew) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunJobParamsBodyAsyncParseConfigNew) implementsParseRunJobParamsBodyUnion() {}

// The URL of the document to be processed. You can provide one of the following:
//
//  1. A publicly available URL
//  2. A presigned S3 URL
//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
//     uploading a document
//
// Satisfied by [shared.UnionString], [shared.UploadParam].
type ParseRunJobParamsBodyAsyncParseConfigNewDocumentURLUnion interface {
	ImplementsParseRunJobParamsBodyAsyncParseConfigNewDocumentURLUnion()
}

type ParseRunJobParamsBodyAsyncParseConfig struct {
	// The URL of the document to be processed. You can provide one of the
	// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
	// prefixed URL obtained from the /upload endpoint after directly uploading a
	// document 4. A jobid:// prefixed URL obtained from a previous /parse invocation
	Input param.Field[ParseRunJobParamsBodyAsyncParseConfigInputUnion] `json:"input,required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async       param.Field[ParseRunJobParamsBodyAsyncParseConfigAsync]       `json:"async"`
	Enhance     param.Field[ParseRunJobParamsBodyAsyncParseConfigEnhance]     `json:"enhance"`
	Formatting  param.Field[ParseRunJobParamsBodyAsyncParseConfigFormatting]  `json:"formatting"`
	Retrieval   param.Field[ParseRunJobParamsBodyAsyncParseConfigRetrieval]   `json:"retrieval"`
	Settings    param.Field[ParseRunJobParamsBodyAsyncParseConfigSettings]    `json:"settings"`
	Spreadsheet param.Field[ParseRunJobParamsBodyAsyncParseConfigSpreadsheet] `json:"spreadsheet"`
}

func (r ParseRunJobParamsBodyAsyncParseConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunJobParamsBodyAsyncParseConfig) implementsParseRunJobParamsBodyUnion() {}

// The URL of the document to be processed. You can provide one of the
// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
// prefixed URL obtained from the /upload endpoint after directly uploading a
// document 4. A jobid:// prefixed URL obtained from a previous /parse invocation
//
// Satisfied by [shared.UnionString], [shared.UploadParam].
type ParseRunJobParamsBodyAsyncParseConfigInputUnion interface {
	ImplementsParseRunJobParamsBodyAsyncParseConfigInputUnion()
}

// The configuration options for asynchronous processing (default synchronous).
type ParseRunJobParamsBodyAsyncParseConfigAsync struct {
	// JSON metadata included in webhook request body. Defaults to None.
	Metadata param.Field[interface{}] `json:"metadata"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool] `json:"priority"`
	// The webhook configuration for the asynchronous processing.
	Webhook param.Field[ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookUnion] `json:"webhook"`
}

func (r ParseRunJobParamsBodyAsyncParseConfigAsync) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The webhook configuration for the asynchronous processing.
type ParseRunJobParamsBodyAsyncParseConfigAsyncWebhook struct {
	Channels param.Field[interface{}]                                           `json:"channels"`
	Mode     param.Field[ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookMode] `json:"mode"`
	URL      param.Field[string]                                                `json:"url"`
}

func (r ParseRunJobParamsBodyAsyncParseConfigAsyncWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunJobParamsBodyAsyncParseConfigAsyncWebhook) implementsParseRunJobParamsBodyAsyncParseConfigAsyncWebhookUnion() {
}

// The webhook configuration for the asynchronous processing.
//
// Satisfied by
// [ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfig],
// [ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfig],
// [ParseRunJobParamsBodyAsyncParseConfigAsyncWebhook].
type ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookUnion interface {
	implementsParseRunJobParamsBodyAsyncParseConfigAsyncWebhookUnion()
}

type ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfig struct {
	// A list of Svix channels the message will be delivered down, omit to send to all
	// channels.
	Channels param.Field[[]string]                                                               `json:"channels"`
	Mode     param.Field[ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfigMode] `json:"mode"`
}

func (r ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfig) implementsParseRunJobParamsBodyAsyncParseConfigAsyncWebhookUnion() {
}

type ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfigMode string

const (
	ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfigModeSvix ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfigMode = "svix"
)

func (r ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfigMode) IsKnown() bool {
	switch r {
	case ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookSvixWebhookConfigModeSvix:
		return true
	}
	return false
}

type ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfig struct {
	URL  param.Field[string]                                                                   `json:"url,required"`
	Mode param.Field[ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfigMode] `json:"mode"`
}

func (r ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfig) implementsParseRunJobParamsBodyAsyncParseConfigAsyncWebhookUnion() {
}

type ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfigMode string

const (
	ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfigModeDirect ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfigMode = "direct"
)

func (r ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfigMode) IsKnown() bool {
	switch r {
	case ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookDirectWebhookConfigModeDirect:
		return true
	}
	return false
}

type ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookMode string

const (
	ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookModeSvix   ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookMode = "svix"
	ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookModeDirect ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookMode = "direct"
)

func (r ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookMode) IsKnown() bool {
	switch r {
	case ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookModeSvix, ParseRunJobParamsBodyAsyncParseConfigAsyncWebhookModeDirect:
		return true
	}
	return false
}

type ParseRunJobParamsBodyAsyncParseConfigEnhance struct {
	// Agentic uses vision language models to enhance the accuracy of the output of
	// different types of extraction. This will incur a cost and latency increase.
	Agentic param.Field[[]ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticUnion] `json:"agentic"`
	// If True, summarize figures using a small vision language model. Defaults to
	// True.
	SummarizeFigures param.Field[bool] `json:"summarize_figures"`
}

func (r ParseRunJobParamsBodyAsyncParseConfigEnhance) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ParseRunJobParamsBodyAsyncParseConfigEnhanceAgentic struct {
	Scope param.Field[ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticScope] `json:"scope,required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ParseRunJobParamsBodyAsyncParseConfigEnhanceAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunJobParamsBodyAsyncParseConfigEnhanceAgentic) implementsParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticUnion() {
}

// Satisfied by [ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTableAgentic],
// [ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticFigureAgentic],
// [ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTextAgentic],
// [ParseRunJobParamsBodyAsyncParseConfigEnhanceAgentic].
type ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticUnion interface {
	implementsParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticUnion()
}

type ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTableAgentic struct {
	Scope param.Field[ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTableAgenticScope] `json:"scope,required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTableAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTableAgentic) implementsParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticUnion() {
}

type ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTableAgenticScope string

const (
	ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTableAgenticScopeTable ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTableAgenticScope = "table"
)

func (r ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTableAgenticScope) IsKnown() bool {
	switch r {
	case ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTableAgenticScopeTable:
		return true
	}
	return false
}

type ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticFigureAgentic struct {
	Scope param.Field[ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticFigureAgenticScope] `json:"scope,required"`
	// Custom prompt for figure agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticFigureAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticFigureAgentic) implementsParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticUnion() {
}

type ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticFigureAgenticScope string

const (
	ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticFigureAgenticScopeFigure ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticFigureAgenticScope = "figure"
)

func (r ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticFigureAgenticScope) IsKnown() bool {
	switch r {
	case ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticFigureAgenticScopeFigure:
		return true
	}
	return false
}

type ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTextAgentic struct {
	Scope param.Field[ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTextAgenticScope] `json:"scope,required"`
}

func (r ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTextAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTextAgentic) implementsParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticUnion() {
}

type ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTextAgenticScope string

const (
	ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTextAgenticScopeText ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTextAgenticScope = "text"
)

func (r ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTextAgenticScope) IsKnown() bool {
	switch r {
	case ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticTextAgenticScopeText:
		return true
	}
	return false
}

type ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticScope string

const (
	ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticScopeTable  ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticScope = "table"
	ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticScopeFigure ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticScope = "figure"
	ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticScopeText   ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticScope = "text"
)

func (r ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticScope) IsKnown() bool {
	switch r {
	case ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticScopeTable, ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticScopeFigure, ParseRunJobParamsBodyAsyncParseConfigEnhanceAgenticScopeText:
		return true
	}
	return false
}

type ParseRunJobParamsBodyAsyncParseConfigFormatting struct {
	// If True, add page markers to the output. Defaults to False. Useful for
	// extracting data with page specific information.
	AddPageMarkers param.Field[bool] `json:"add_page_markers"`
	// A list of formatting to include in the output. [insert description of each
	// option here later]
	Include param.Field[[]ParseRunJobParamsBodyAsyncParseConfigFormattingInclude] `json:"include"`
	// A flag to indicate if consecutive tables with the same number of columns should
	// be merged. Defaults to False.
	MergeTables param.Field[bool] `json:"merge_tables"`
	// The mode to use for table output. Defaults to dynamic, which returns md for
	// simpler tables and html for more complex tables.
	TableOutputFormat param.Field[ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormat] `json:"table_output_format"`
}

func (r ParseRunJobParamsBodyAsyncParseConfigFormatting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ParseRunJobParamsBodyAsyncParseConfigFormattingInclude string

const (
	ParseRunJobParamsBodyAsyncParseConfigFormattingIncludeChangeTracking ParseRunJobParamsBodyAsyncParseConfigFormattingInclude = "change_tracking"
	ParseRunJobParamsBodyAsyncParseConfigFormattingIncludeHighlight      ParseRunJobParamsBodyAsyncParseConfigFormattingInclude = "highlight"
	ParseRunJobParamsBodyAsyncParseConfigFormattingIncludeComments       ParseRunJobParamsBodyAsyncParseConfigFormattingInclude = "comments"
)

func (r ParseRunJobParamsBodyAsyncParseConfigFormattingInclude) IsKnown() bool {
	switch r {
	case ParseRunJobParamsBodyAsyncParseConfigFormattingIncludeChangeTracking, ParseRunJobParamsBodyAsyncParseConfigFormattingIncludeHighlight, ParseRunJobParamsBodyAsyncParseConfigFormattingIncludeComments:
		return true
	}
	return false
}

// The mode to use for table output. Defaults to dynamic, which returns md for
// simpler tables and html for more complex tables.
type ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormat string

const (
	ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormatHTML     ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormat = "html"
	ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormatJson     ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormat = "json"
	ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormatMd       ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormat = "md"
	ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormatJsonbbox ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormat = "jsonbbox"
	ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormatDynamic  ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormat = "dynamic"
	ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormatCsv      ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormat = "csv"
)

func (r ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormat) IsKnown() bool {
	switch r {
	case ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormatHTML, ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormatJson, ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormatMd, ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormatJsonbbox, ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormatDynamic, ParseRunJobParamsBodyAsyncParseConfigFormattingTableOutputFormatCsv:
		return true
	}
	return false
}

type ParseRunJobParamsBodyAsyncParseConfigRetrieval struct {
	Chunking param.Field[ParseRunJobParamsBodyAsyncParseConfigRetrievalChunking] `json:"chunking"`
	// If True, use embedding optimized mode. Defaults to False.
	EmbeddingOptimized param.Field[bool] `json:"embedding_optimized"`
	// A list of block types to filter out from 'content' and 'embed' fields. By
	// default, no blocks are filtered.
	FilterBlocks param.Field[[]ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlock] `json:"filter_blocks"`
}

func (r ParseRunJobParamsBodyAsyncParseConfigRetrieval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ParseRunJobParamsBodyAsyncParseConfigRetrievalChunking struct {
	// Choose how to partition chunks. Variable mode chunks by character length and
	// visual context. Section mode chunks by section headers. Page mode chunks
	// according to pages. Page sections mode chunks first by page, then by sections
	// within each page. Disabled returns one single chunk.
	ChunkMode param.Field[ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkMode] `json:"chunk_mode"`
	// The approximate size of chunks (in characters) that the document will be split
	// into. Defaults to null, in which case the chunk size is variable between 250 -
	// 1500 characters.
	ChunkSize param.Field[int64] `json:"chunk_size"`
}

func (r ParseRunJobParamsBodyAsyncParseConfigRetrievalChunking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Choose how to partition chunks. Variable mode chunks by character length and
// visual context. Section mode chunks by section headers. Page mode chunks
// according to pages. Page sections mode chunks first by page, then by sections
// within each page. Disabled returns one single chunk.
type ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkMode string

const (
	ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkModeVariable     ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkMode = "variable"
	ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkModeSection      ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkMode = "section"
	ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkModePage         ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkMode = "page"
	ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkModeDisabled     ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkMode = "disabled"
	ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkModeBlock        ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkMode = "block"
	ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkModePageSections ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkMode = "page_sections"
)

func (r ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkMode) IsKnown() bool {
	switch r {
	case ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkModeVariable, ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkModeSection, ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkModePage, ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkModeDisabled, ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkModeBlock, ParseRunJobParamsBodyAsyncParseConfigRetrievalChunkingChunkModePageSections:
		return true
	}
	return false
}

type ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlock string

const (
	ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockHeader        ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlock = "Header"
	ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockFooter        ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlock = "Footer"
	ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockTitle         ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlock = "Title"
	ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockSectionHeader ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlock = "Section Header"
	ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockPageNumber    ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlock = "Page Number"
	ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockListItem      ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlock = "List Item"
	ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockFigure        ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlock = "Figure"
	ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockTable         ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlock = "Table"
	ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockKeyValue      ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlock = "Key Value"
	ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockText          ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlock = "Text"
	ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockComment       ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlock = "Comment"
	ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockSignature     ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlock = "Signature"
)

func (r ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlock) IsKnown() bool {
	switch r {
	case ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockHeader, ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockFooter, ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockTitle, ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockSectionHeader, ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockPageNumber, ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockListItem, ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockFigure, ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockTable, ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockKeyValue, ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockText, ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockComment, ParseRunJobParamsBodyAsyncParseConfigRetrievalFilterBlockSignature:
		return true
	}
	return false
}

type ParseRunJobParamsBodyAsyncParseConfigSettings struct {
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
	OcrSystem param.Field[ParseRunJobParamsBodyAsyncParseConfigSettingsOcrSystem] `json:"ocr_system"`
	// The page range to process (1-indexed). By default, the entire document is
	// processed.
	PageRange param.Field[ParseRunJobParamsBodyAsyncParseConfigSettingsPageRangeUnion] `json:"page_range"`
	// If True, persist the results indefinitely. Defaults to False.
	PersistResults param.Field[bool] `json:"persist_results"`
	// Whether to return images for the specified block types. By default, no images
	// are returned.
	ReturnImages param.Field[[]ParseRunJobParamsBodyAsyncParseConfigSettingsReturnImage] `json:"return_images"`
	// If True, return OCR data in the result. Defaults to False.
	ReturnOcrData param.Field[bool] `json:"return_ocr_data"`
	// The timeout for the job in seconds. Defaults to 900.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r ParseRunJobParamsBodyAsyncParseConfigSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Standard is our best multilingual OCR system. Legacy only supports germanic
// languages and is available for backwards compatibility.
type ParseRunJobParamsBodyAsyncParseConfigSettingsOcrSystem string

const (
	ParseRunJobParamsBodyAsyncParseConfigSettingsOcrSystemStandard ParseRunJobParamsBodyAsyncParseConfigSettingsOcrSystem = "standard"
	ParseRunJobParamsBodyAsyncParseConfigSettingsOcrSystemLegacy   ParseRunJobParamsBodyAsyncParseConfigSettingsOcrSystem = "legacy"
)

func (r ParseRunJobParamsBodyAsyncParseConfigSettingsOcrSystem) IsKnown() bool {
	switch r {
	case ParseRunJobParamsBodyAsyncParseConfigSettingsOcrSystemStandard, ParseRunJobParamsBodyAsyncParseConfigSettingsOcrSystemLegacy:
		return true
	}
	return false
}

// The page range to process (1-indexed). By default, the entire document is
// processed.
//
// Satisfied by [shared.PageRangeParam],
// [ParseRunJobParamsBodyAsyncParseConfigSettingsPageRangeArray],
// [ParseRunJobParamsBodyAsyncParseConfigSettingsPageRangeArray].
type ParseRunJobParamsBodyAsyncParseConfigSettingsPageRangeUnion interface {
	ImplementsParseRunJobParamsBodyAsyncParseConfigSettingsPageRangeUnion()
}

type ParseRunJobParamsBodyAsyncParseConfigSettingsPageRangeArray []shared.PageRangeParam

func (r ParseRunJobParamsBodyAsyncParseConfigSettingsPageRangeArray) ImplementsParseRunJobParamsBodyAsyncParseConfigSettingsPageRangeUnion() {
}

type ParseRunJobParamsBodyAsyncParseConfigSettingsReturnImage string

const (
	ParseRunJobParamsBodyAsyncParseConfigSettingsReturnImageFigure ParseRunJobParamsBodyAsyncParseConfigSettingsReturnImage = "figure"
	ParseRunJobParamsBodyAsyncParseConfigSettingsReturnImageTable  ParseRunJobParamsBodyAsyncParseConfigSettingsReturnImage = "table"
)

func (r ParseRunJobParamsBodyAsyncParseConfigSettingsReturnImage) IsKnown() bool {
	switch r {
	case ParseRunJobParamsBodyAsyncParseConfigSettingsReturnImageFigure, ParseRunJobParamsBodyAsyncParseConfigSettingsReturnImageTable:
		return true
	}
	return false
}

type ParseRunJobParamsBodyAsyncParseConfigSpreadsheet struct {
	// In a spreadsheet with different tables inside, we enable splitting up the tables
	// by default. Accurate mode applies more powerful models for superior accuracy, at
	// 5× the default per-cell rate. Disabling will register as one large table.
	Clustering param.Field[ParseRunJobParamsBodyAsyncParseConfigSpreadsheetClustering] `json:"clustering"`
	// Whether to exclude hidden sheets, rows, or columns in the output.
	Exclude param.Field[[]ParseRunJobParamsBodyAsyncParseConfigSpreadsheetExclude] `json:"exclude"`
	// Whether to include cell color and formula information in the output.
	Include          param.Field[[]ParseRunJobParamsBodyAsyncParseConfigSpreadsheetInclude]        `json:"include"`
	SplitLargeTables param.Field[ParseRunJobParamsBodyAsyncParseConfigSpreadsheetSplitLargeTables] `json:"split_large_tables"`
}

func (r ParseRunJobParamsBodyAsyncParseConfigSpreadsheet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// In a spreadsheet with different tables inside, we enable splitting up the tables
// by default. Accurate mode applies more powerful models for superior accuracy, at
// 5× the default per-cell rate. Disabling will register as one large table.
type ParseRunJobParamsBodyAsyncParseConfigSpreadsheetClustering string

const (
	ParseRunJobParamsBodyAsyncParseConfigSpreadsheetClusteringAccurate ParseRunJobParamsBodyAsyncParseConfigSpreadsheetClustering = "accurate"
	ParseRunJobParamsBodyAsyncParseConfigSpreadsheetClusteringFast     ParseRunJobParamsBodyAsyncParseConfigSpreadsheetClustering = "fast"
	ParseRunJobParamsBodyAsyncParseConfigSpreadsheetClusteringDisabled ParseRunJobParamsBodyAsyncParseConfigSpreadsheetClustering = "disabled"
)

func (r ParseRunJobParamsBodyAsyncParseConfigSpreadsheetClustering) IsKnown() bool {
	switch r {
	case ParseRunJobParamsBodyAsyncParseConfigSpreadsheetClusteringAccurate, ParseRunJobParamsBodyAsyncParseConfigSpreadsheetClusteringFast, ParseRunJobParamsBodyAsyncParseConfigSpreadsheetClusteringDisabled:
		return true
	}
	return false
}

type ParseRunJobParamsBodyAsyncParseConfigSpreadsheetExclude string

const (
	ParseRunJobParamsBodyAsyncParseConfigSpreadsheetExcludeHiddenSheets ParseRunJobParamsBodyAsyncParseConfigSpreadsheetExclude = "hidden_sheets"
	ParseRunJobParamsBodyAsyncParseConfigSpreadsheetExcludeHiddenRows   ParseRunJobParamsBodyAsyncParseConfigSpreadsheetExclude = "hidden_rows"
	ParseRunJobParamsBodyAsyncParseConfigSpreadsheetExcludeHiddenCols   ParseRunJobParamsBodyAsyncParseConfigSpreadsheetExclude = "hidden_cols"
)

func (r ParseRunJobParamsBodyAsyncParseConfigSpreadsheetExclude) IsKnown() bool {
	switch r {
	case ParseRunJobParamsBodyAsyncParseConfigSpreadsheetExcludeHiddenSheets, ParseRunJobParamsBodyAsyncParseConfigSpreadsheetExcludeHiddenRows, ParseRunJobParamsBodyAsyncParseConfigSpreadsheetExcludeHiddenCols:
		return true
	}
	return false
}

type ParseRunJobParamsBodyAsyncParseConfigSpreadsheetInclude string

const (
	ParseRunJobParamsBodyAsyncParseConfigSpreadsheetIncludeCellColors ParseRunJobParamsBodyAsyncParseConfigSpreadsheetInclude = "cell_colors"
	ParseRunJobParamsBodyAsyncParseConfigSpreadsheetIncludeFormula    ParseRunJobParamsBodyAsyncParseConfigSpreadsheetInclude = "formula"
)

func (r ParseRunJobParamsBodyAsyncParseConfigSpreadsheetInclude) IsKnown() bool {
	switch r {
	case ParseRunJobParamsBodyAsyncParseConfigSpreadsheetIncludeCellColors, ParseRunJobParamsBodyAsyncParseConfigSpreadsheetIncludeFormula:
		return true
	}
	return false
}

type ParseRunJobParamsBodyAsyncParseConfigSpreadsheetSplitLargeTables struct {
	// If True, split large tables into smaller tables. Defaults to True.
	Enabled param.Field[bool] `json:"enabled"`
	// The size of the tables to split into. Defaults to 50.
	Size param.Field[int64] `json:"size"`
}

func (r ParseRunJobParamsBodyAsyncParseConfigSpreadsheetSplitLargeTables) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
