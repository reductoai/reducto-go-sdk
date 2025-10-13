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
	// This field can have the runtime type of [[]interface{}].
	Result interface{} `json:"result,required"`
	// This field can have the runtime type of [shared.ExtractResponseUsage],
	// [ExtractRunResponseV3ExtractResponseUsage].
	Usage interface{} `json:"usage,required"`
	// This field can have the runtime type of [[]interface{}].
	Citations interface{} `json:"citations"`
	JobID     string      `json:"job_id,nullable"`
	// The link to the studio pipeline for the document.
	StudioLink string                 `json:"studio_link,nullable"`
	JSON       extractRunResponseJSON `json:"-"`
	union      ExtractRunResponseUnion
}

// extractRunResponseJSON contains the JSON metadata for the struct
// [ExtractRunResponse]
type extractRunResponseJSON struct {
	Result      apijson.Field
	Usage       apijson.Field
	Citations   apijson.Field
	JobID       apijson.Field
	StudioLink  apijson.Field
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
// Possible runtime types of the union are [shared.ExtractResponse],
// [ExtractRunResponseV3ExtractResponse].
func (r ExtractRunResponse) AsUnion() ExtractRunResponseUnion {
	return r.union
}

// Union satisfied by [shared.ExtractResponse] or
// [ExtractRunResponseV3ExtractResponse].
type ExtractRunResponseUnion interface {
	ImplementsExtractRunResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ExtractRunResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.ExtractResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExtractRunResponseV3ExtractResponse{}),
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

func (r ExtractRunResponseV3ExtractResponse) ImplementsExtractRunResponse() {}

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
	AdvancedOptions         param.Field[shared.AdvancedProcessingOptionsParam] `json:"advanced_options"`
	AgentExtract            param.Field[interface{}]                           `json:"agent_extract"`
	AlphaBigExtractionModel param.Field[bool]                                  `json:"alpha_big_extraction_model"`
	AlphaDeepExtract        param.Field[bool]                                  `json:"alpha_deep_extract"`
	AlphaTableCitations     param.Field[bool]                                  `json:"alpha_table_citations"`
	// The configuration options for array extract
	ArrayExtract param.Field[shared.ArrayExtractConfigParam] `json:"array_extract"`
	Async        param.Field[interface{}]                    `json:"async"`
	// The configuration options for citations.
	CitationsOptions    param.Field[shared.AdvancedCitationsConfigParam]       `json:"citations_options"`
	DocumentURL         param.Field[interface{}]                               `json:"document_url"`
	ExperimentalOptions param.Field[shared.ExperimentalProcessingOptionsParam] `json:"experimental_options"`
	// If table citations should be generated for the extracted content.
	ExperimentalTableCitations param.Field[bool] `json:"experimental_table_citations"`
	// If citations should be generated for the extracted content.
	GenerateCitations param.Field[bool] `json:"generate_citations"`
	// If images should be passed directly for extractions. Can only be enabled for
	// documents with less than 10 pages. Defaults to False.
	IncludeImages param.Field[bool]        `json:"include_images"`
	Input         param.Field[interface{}] `json:"input"`
	Instructions  param.Field[interface{}] `json:"instructions"`
	// If True, the job will be processed with lower latency and higher priority. Uses
	// 2x the cost of a regular job. Defaults to False.
	LatencySensitive param.Field[bool]        `json:"latency_sensitive"`
	NormalizedSchema param.Field[interface{}] `json:"normalized_schema"`
	Options          param.Field[interface{}] `json:"options"`
	ParseConfig      param.Field[interface{}] `json:"parse_config"`
	Parsing          param.Field[interface{}] `json:"parsing"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool]        `json:"priority"`
	Schema   param.Field[interface{}] `json:"schema"`
	Settings param.Field[interface{}] `json:"settings"`
	// If spreadsheet agent should be used for extraction.
	SpreadsheetAgent param.Field[bool] `json:"spreadsheet_agent"`
	// A system prompt to use for the extraction. This is a general prompt that is
	// applied to the entire document before any other prompts.
	SystemPrompt param.Field[string] `json:"system_prompt"`
	// If chunking should be used for the extraction. Defaults to False.
	UseChunking param.Field[bool]        `json:"use_chunking"`
	UserConfig  param.Field[interface{}] `json:"user_config"`
}

func (r ExtractRunParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBody) implementsExtractRunParamsBodyUnion() {}

// Satisfied by [ExtractConfigParam], [ExtractRunParamsBodyExtractConfig],
// [ExtractRunParamsBodySyncExtractConfig], [ExtractRunParamsBody].
type ExtractRunParamsBodyUnion interface {
	implementsExtractRunParamsBodyUnion()
}

type ExtractRunParamsBodyExtractConfig struct {
	DocumentURL param.Field[ExtractRunParamsBodyExtractConfigDocumentURLUnion] `json:"document_url,required"`
	// The JSON schema to use for extraction.
	Schema param.Field[interface{}] `json:"schema,required"`
	// If agent extraction should be used for extraction.
	AgentExtract            param.Field[bool] `json:"agent_extract"`
	AlphaBigExtractionModel param.Field[bool] `json:"alpha_big_extraction_model"`
	AlphaDeepExtract        param.Field[bool] `json:"alpha_deep_extract"`
	AlphaTableCitations     param.Field[bool] `json:"alpha_table_citations"`
	// The configuration options for asynchronous processing (default synchronous).
	Async         param.Field[ExtractRunParamsBodyExtractConfigAsync] `json:"async"`
	IncludeImages param.Field[bool]                                   `json:"include_images"`
	// If True, the job will be processed with lower latency and higher priority. Uses
	// 2x the cost of a regular job. Defaults to False.
	LatencySensitive param.Field[bool] `json:"latency_sensitive"`
	// The normalized JSON schema to use for extraction.
	NormalizedSchema param.Field[interface{}] `json:"normalized_schema"`
	// The configuration options for extraction.
	Options param.Field[ExtractRunParamsBodyExtractConfigOptions] `json:"options"`
	// The configuration options for extraction.
	ParseConfig param.Field[ExtractRunParamsBodyExtractConfigParseConfig] `json:"parse_config"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool] `json:"priority"`
	// A system prompt to use for the extraction. This is a general prompt that is
	// applied to the entire document before any other prompts.
	SystemPrompt param.Field[string] `json:"system_prompt"`
	// User-specific configuration options.
	UserConfig param.Field[map[string]interface{}] `json:"user_config"`
}

func (r ExtractRunParamsBodyExtractConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBodyExtractConfig) implementsExtractRunParamsBodyUnion() {}

// Satisfied by [shared.UnionString],
// [ExtractRunParamsBodyExtractConfigDocumentURLArray].
type ExtractRunParamsBodyExtractConfigDocumentURLUnion interface {
	ImplementsExtractRunParamsBodyExtractConfigDocumentURLUnion()
}

type ExtractRunParamsBodyExtractConfigDocumentURLArray []string

func (r ExtractRunParamsBodyExtractConfigDocumentURLArray) ImplementsExtractRunParamsBodyExtractConfigDocumentURLUnion() {
}

// The configuration options for asynchronous processing (default synchronous).
type ExtractRunParamsBodyExtractConfigAsync struct {
	Enabled  param.Field[bool]                                          `json:"enabled"`
	Priority param.Field[bool]                                          `json:"priority"`
	Webhook  param.Field[ExtractRunParamsBodyExtractConfigAsyncWebhook] `json:"webhook"`
}

func (r ExtractRunParamsBodyExtractConfigAsync) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunParamsBodyExtractConfigAsyncWebhook struct {
	// A list of Svix channels the message will be delivered down, omit to send to all
	// channels.
	Channels param.Field[[]string]                                          `json:"channels"`
	Metadata param.Field[interface{}]                                       `json:"metadata"`
	Mode     param.Field[ExtractRunParamsBodyExtractConfigAsyncWebhookMode] `json:"mode"`
	URL      param.Field[string]                                            `json:"url"`
}

func (r ExtractRunParamsBodyExtractConfigAsyncWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunParamsBodyExtractConfigAsyncWebhookMode string

const (
	ExtractRunParamsBodyExtractConfigAsyncWebhookModeDirect ExtractRunParamsBodyExtractConfigAsyncWebhookMode = "direct"
	ExtractRunParamsBodyExtractConfigAsyncWebhookModeSvix   ExtractRunParamsBodyExtractConfigAsyncWebhookMode = "svix"
)

func (r ExtractRunParamsBodyExtractConfigAsyncWebhookMode) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyExtractConfigAsyncWebhookModeDirect, ExtractRunParamsBodyExtractConfigAsyncWebhookModeSvix:
		return true
	}
	return false
}

// The configuration options for extraction.
type ExtractRunParamsBodyExtractConfigOptions struct {
	// Array extraction allows you to extract long lists of information from lengthy
	// documents. It makes parallel calls on overlapping sections of the document.
	ArrayExtract param.Field[bool] `json:"array_extract"`
	// Length of each segment, in pages, for parallel calls with array extraction.
	ArrayExtractPages param.Field[int64] `json:"array_extract_pages"`
	// If table citations should be generated for the extracted content.
	ExperimentalTableCitations param.Field[bool] `json:"experimental_table_citations"`
	// The array extraction version to use.
	ExtractAlgorithm param.Field[ExtractRunParamsBodyExtractConfigOptionsExtractAlgorithm] `json:"extract_algorithm"`
	// If citations should be generated for the extracted content.
	GenerateCitations param.Field[bool] `json:"generate_citations"`
	// If True, enable numeric citation confidence scores. Defaults to False.
	NumericalConfidence param.Field[bool] `json:"numerical_confidence"`
	// If spreadsheet agent should be used for extraction.
	SpreadsheetAgent param.Field[bool] `json:"spreadsheet_agent"`
	// Number of items to extract in each stream call.
	StreamingExtractItemDensity param.Field[int64] `json:"streaming_extract_item_density"`
}

func (r ExtractRunParamsBodyExtractConfigOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The array extraction version to use.
type ExtractRunParamsBodyExtractConfigOptionsExtractAlgorithm string

const (
	ExtractRunParamsBodyExtractConfigOptionsExtractAlgorithmAuto      ExtractRunParamsBodyExtractConfigOptionsExtractAlgorithm = "auto"
	ExtractRunParamsBodyExtractConfigOptionsExtractAlgorithmLegacy    ExtractRunParamsBodyExtractConfigOptionsExtractAlgorithm = "legacy"
	ExtractRunParamsBodyExtractConfigOptionsExtractAlgorithmStreaming ExtractRunParamsBodyExtractConfigOptionsExtractAlgorithm = "streaming"
	ExtractRunParamsBodyExtractConfigOptionsExtractAlgorithmNoOverlap ExtractRunParamsBodyExtractConfigOptionsExtractAlgorithm = "no_overlap"
)

func (r ExtractRunParamsBodyExtractConfigOptionsExtractAlgorithm) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyExtractConfigOptionsExtractAlgorithmAuto, ExtractRunParamsBodyExtractConfigOptionsExtractAlgorithmLegacy, ExtractRunParamsBodyExtractConfigOptionsExtractAlgorithmStreaming, ExtractRunParamsBodyExtractConfigOptionsExtractAlgorithmNoOverlap:
		return true
	}
	return false
}

// The configuration options for extraction.
type ExtractRunParamsBodyExtractConfigParseConfig struct {
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
	ChunkMode param.Field[ExtractRunParamsBodyExtractConfigParseConfigChunkMode] `json:"chunk_mode"`
	// The approximate size of chunks (in characters) that the document will be split
	// into. Defaults to None, in which case the chunk size is variable between 250 -
	// 1500 characters.
	ChunkSize param.Field[int64] `json:"chunk_size"`
	// A flag to indicate if the hierarchy of the document should be continued from
	// chunk to chunk. E.g. ## Prev Section (cont.)
	ContinueHierarchy param.Field[bool]                                                     `json:"continue_hierarchy"`
	CustomFormat      param.Field[ExtractRunParamsBodyExtractConfigParseConfigCustomFormat] `json:"custom_format"`
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
	EnrichMode param.Field[ExtractRunParamsBodyExtractConfigParseConfigEnrichMode] `json:"enrich_mode"`
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
	IgnoreBlocks param.Field[[]ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlock] `json:"ignore_blocks"`
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
	LayoutModel param.Field[ExtractRunParamsBodyExtractConfigParseConfigLayoutModel] `json:"layout_model"`
	// The maximum number of pages to process in a single batch. Defaults to 10.
	MaxBatchSize param.Field[int64] `json:"max_batch_size"`
	// A flag to indicate if consecutive tables with the same number of columns should
	// be merged. Defaults to False.
	MergeTables param.Field[bool] `json:"merge_tables"`
	// The type of document to be processed. Defaults to document. If auto is
	// specified, the orientation of the first page will be used to determine the
	// document type.
	Mode param.Field[ExtractRunParamsBodyExtractConfigParseConfigMode] `json:"mode"`
	// The dimension of the OCR crops along each axis. num_ocr_crops^2 is the total
	// number of crops. Defaults to 2.
	NumOcrCrops param.Field[int64] `json:"num_ocr_crops"`
	// If True, enable numeric parse confidence scores in granular_confidence field.
	// Defaults to False.
	NumericalParseConfidence param.Field[bool] `json:"numerical_parse_confidence"`
	// The mode to use for OCR. If agentic is enabled, table OCR will be automatically
	// edited.
	OcrMode param.Field[ExtractRunParamsBodyExtractConfigParseConfigOcrMode] `json:"ocr_mode"`
	// The OCR system to use. Defaults to cloud (AWS Textract/Azure DocAI/etc).
	OcrSystem param.Field[ExtractRunParamsBodyExtractConfigParseConfigOcrSystem] `json:"ocr_system"`
	// The threshold for box overlap. Defaults to 0.5.
	OverlapThreshold param.Field[float64] `json:"overlap_threshold"`
	// The page number to stop processing at.
	PageEnd param.Field[int64] `json:"page_end"`
	// The page range to process.
	PageRange param.Field[ExtractRunParamsBodyExtractConfigParseConfigPageRangeUnion] `json:"page_range"`
	// The page number to start processing from.
	PageStart param.Field[int64] `json:"page_start"`
	// The method to use for OCR. hybrid uses the PDF text first, then OCR. pdf only
	// uses the PDF text. ocr only uses OCR.
	PdfOcr param.Field[ExtractRunParamsBodyExtractConfigParseConfigPdfOcr] `json:"pdf_ocr"`
	// If True, persist the results indefinitely. Defaults to False.
	PersistResults param.Field[bool] `json:"persist_results"`
	// Forces all external API calls to be routed to specified country/region.
	RegionPreference param.Field[ExtractRunParamsBodyExtractConfigParseConfigRegionPreference] `json:"region_preference"`
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
	SpreadsheetTableClustering param.Field[ExtractRunParamsBodyExtractConfigParseConfigSpreadsheetTableClustering] `json:"spreadsheet_table_clustering"`
	// If True, enable figure summaries for all figures regardless of size (onprem
	// only). Defaults to False.
	SummarizeAllFigures param.Field[bool] `json:"summarize_all_figures"`
	// The mode to use for table output. Defaults to html.
	TableOutputFormat param.Field[ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormat] `json:"table_output_format"`
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
	Version param.Field[ExtractRunParamsBodyExtractConfigParseConfigVersion] `json:"version"`
}

func (r ExtractRunParamsBodyExtractConfigParseConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The mode to use for chunking. Defaults to 'variable'. Section chunks according
// to sections in the document. Page chunks according to pages. Page sections
// chunks according to both pages and sections. Disabled returns a single chunk.
type ExtractRunParamsBodyExtractConfigParseConfigChunkMode string

const (
	ExtractRunParamsBodyExtractConfigParseConfigChunkModeVariable     ExtractRunParamsBodyExtractConfigParseConfigChunkMode = "variable"
	ExtractRunParamsBodyExtractConfigParseConfigChunkModeSection      ExtractRunParamsBodyExtractConfigParseConfigChunkMode = "section"
	ExtractRunParamsBodyExtractConfigParseConfigChunkModePage         ExtractRunParamsBodyExtractConfigParseConfigChunkMode = "page"
	ExtractRunParamsBodyExtractConfigParseConfigChunkModeDisabled     ExtractRunParamsBodyExtractConfigParseConfigChunkMode = "disabled"
	ExtractRunParamsBodyExtractConfigParseConfigChunkModeBlock        ExtractRunParamsBodyExtractConfigParseConfigChunkMode = "block"
	ExtractRunParamsBodyExtractConfigParseConfigChunkModePageSections ExtractRunParamsBodyExtractConfigParseConfigChunkMode = "page_sections"
)

func (r ExtractRunParamsBodyExtractConfigParseConfigChunkMode) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyExtractConfigParseConfigChunkModeVariable, ExtractRunParamsBodyExtractConfigParseConfigChunkModeSection, ExtractRunParamsBodyExtractConfigParseConfigChunkModePage, ExtractRunParamsBodyExtractConfigParseConfigChunkModeDisabled, ExtractRunParamsBodyExtractConfigParseConfigChunkModeBlock, ExtractRunParamsBodyExtractConfigParseConfigChunkModePageSections:
		return true
	}
	return false
}

type ExtractRunParamsBodyExtractConfigParseConfigCustomFormat string

const (
	ExtractRunParamsBodyExtractConfigParseConfigCustomFormatAml     ExtractRunParamsBodyExtractConfigParseConfigCustomFormat = "aml"
	ExtractRunParamsBodyExtractConfigParseConfigCustomFormatAIUsage ExtractRunParamsBodyExtractConfigParseConfigCustomFormat = "ai_usage"
)

func (r ExtractRunParamsBodyExtractConfigParseConfigCustomFormat) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyExtractConfigParseConfigCustomFormatAml, ExtractRunParamsBodyExtractConfigParseConfigCustomFormatAIUsage:
		return true
	}
	return false
}

// The mode to use for enrichment. Defaults to standard
type ExtractRunParamsBodyExtractConfigParseConfigEnrichMode string

const (
	ExtractRunParamsBodyExtractConfigParseConfigEnrichModeStandard ExtractRunParamsBodyExtractConfigParseConfigEnrichMode = "standard"
	ExtractRunParamsBodyExtractConfigParseConfigEnrichModePage     ExtractRunParamsBodyExtractConfigParseConfigEnrichMode = "page"
	ExtractRunParamsBodyExtractConfigParseConfigEnrichModeTable    ExtractRunParamsBodyExtractConfigParseConfigEnrichMode = "table"
)

func (r ExtractRunParamsBodyExtractConfigParseConfigEnrichMode) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyExtractConfigParseConfigEnrichModeStandard, ExtractRunParamsBodyExtractConfigParseConfigEnrichModePage, ExtractRunParamsBodyExtractConfigParseConfigEnrichModeTable:
		return true
	}
	return false
}

type ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlock string

const (
	ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockHeader        ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlock = "Header"
	ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockFooter        ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlock = "Footer"
	ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockTitle         ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlock = "Title"
	ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockSectionHeader ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlock = "Section Header"
	ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockPageNumber    ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlock = "Page Number"
	ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockListItem      ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlock = "List Item"
	ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockFigure        ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlock = "Figure"
	ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockTable         ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlock = "Table"
	ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockKeyValue      ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlock = "Key Value"
	ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockText          ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlock = "Text"
	ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockComment       ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlock = "Comment"
	ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockSignature     ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlock = "Signature"
)

func (r ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlock) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockHeader, ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockFooter, ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockTitle, ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockSectionHeader, ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockPageNumber, ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockListItem, ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockFigure, ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockTable, ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockKeyValue, ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockText, ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockComment, ExtractRunParamsBodyExtractConfigParseConfigIgnoreBlockSignature:
		return true
	}
	return false
}

// The layout model to use for the document. This will be deprecated in the future.
type ExtractRunParamsBodyExtractConfigParseConfigLayoutModel string

const (
	ExtractRunParamsBodyExtractConfigParseConfigLayoutModelDefault ExtractRunParamsBodyExtractConfigParseConfigLayoutModel = "default"
	ExtractRunParamsBodyExtractConfigParseConfigLayoutModelBeta    ExtractRunParamsBodyExtractConfigParseConfigLayoutModel = "beta"
)

func (r ExtractRunParamsBodyExtractConfigParseConfigLayoutModel) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyExtractConfigParseConfigLayoutModelDefault, ExtractRunParamsBodyExtractConfigParseConfigLayoutModelBeta:
		return true
	}
	return false
}

// The type of document to be processed. Defaults to document. If auto is
// specified, the orientation of the first page will be used to determine the
// document type.
type ExtractRunParamsBodyExtractConfigParseConfigMode string

const (
	ExtractRunParamsBodyExtractConfigParseConfigModeDocument ExtractRunParamsBodyExtractConfigParseConfigMode = "document"
	ExtractRunParamsBodyExtractConfigParseConfigModeDeck     ExtractRunParamsBodyExtractConfigParseConfigMode = "deck"
	ExtractRunParamsBodyExtractConfigParseConfigModeAuto     ExtractRunParamsBodyExtractConfigParseConfigMode = "auto"
)

func (r ExtractRunParamsBodyExtractConfigParseConfigMode) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyExtractConfigParseConfigModeDocument, ExtractRunParamsBodyExtractConfigParseConfigModeDeck, ExtractRunParamsBodyExtractConfigParseConfigModeAuto:
		return true
	}
	return false
}

// The mode to use for OCR. If agentic is enabled, table OCR will be automatically
// edited.
type ExtractRunParamsBodyExtractConfigParseConfigOcrMode string

const (
	ExtractRunParamsBodyExtractConfigParseConfigOcrModeStandard ExtractRunParamsBodyExtractConfigParseConfigOcrMode = "standard"
	ExtractRunParamsBodyExtractConfigParseConfigOcrModeAgentic  ExtractRunParamsBodyExtractConfigParseConfigOcrMode = "agentic"
)

func (r ExtractRunParamsBodyExtractConfigParseConfigOcrMode) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyExtractConfigParseConfigOcrModeStandard, ExtractRunParamsBodyExtractConfigParseConfigOcrModeAgentic:
		return true
	}
	return false
}

// The OCR system to use. Defaults to cloud (AWS Textract/Azure DocAI/etc).
type ExtractRunParamsBodyExtractConfigParseConfigOcrSystem string

const (
	ExtractRunParamsBodyExtractConfigParseConfigOcrSystemGcloud    ExtractRunParamsBodyExtractConfigParseConfigOcrSystem = "gcloud"
	ExtractRunParamsBodyExtractConfigParseConfigOcrSystemTextract  ExtractRunParamsBodyExtractConfigParseConfigOcrSystem = "textract"
	ExtractRunParamsBodyExtractConfigParseConfigOcrSystemTesseract ExtractRunParamsBodyExtractConfigParseConfigOcrSystem = "tesseract"
	ExtractRunParamsBodyExtractConfigParseConfigOcrSystemCombined  ExtractRunParamsBodyExtractConfigParseConfigOcrSystem = "combined"
)

func (r ExtractRunParamsBodyExtractConfigParseConfigOcrSystem) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyExtractConfigParseConfigOcrSystemGcloud, ExtractRunParamsBodyExtractConfigParseConfigOcrSystemTextract, ExtractRunParamsBodyExtractConfigParseConfigOcrSystemTesseract, ExtractRunParamsBodyExtractConfigParseConfigOcrSystemCombined:
		return true
	}
	return false
}

// The page range to process.
//
// Satisfied by [shared.PageRangeParam],
// [ExtractRunParamsBodyExtractConfigParseConfigPageRangeArray],
// [ExtractRunParamsBodyExtractConfigParseConfigPageRangeArray].
type ExtractRunParamsBodyExtractConfigParseConfigPageRangeUnion interface {
	ImplementsExtractRunParamsBodyExtractConfigParseConfigPageRangeUnion()
}

type ExtractRunParamsBodyExtractConfigParseConfigPageRangeArray []shared.PageRangeParam

func (r ExtractRunParamsBodyExtractConfigParseConfigPageRangeArray) ImplementsExtractRunParamsBodyExtractConfigParseConfigPageRangeUnion() {
}

// The method to use for OCR. hybrid uses the PDF text first, then OCR. pdf only
// uses the PDF text. ocr only uses OCR.
type ExtractRunParamsBodyExtractConfigParseConfigPdfOcr string

const (
	ExtractRunParamsBodyExtractConfigParseConfigPdfOcrHybrid ExtractRunParamsBodyExtractConfigParseConfigPdfOcr = "hybrid"
	ExtractRunParamsBodyExtractConfigParseConfigPdfOcrPdf    ExtractRunParamsBodyExtractConfigParseConfigPdfOcr = "pdf"
	ExtractRunParamsBodyExtractConfigParseConfigPdfOcrOcr    ExtractRunParamsBodyExtractConfigParseConfigPdfOcr = "ocr"
)

func (r ExtractRunParamsBodyExtractConfigParseConfigPdfOcr) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyExtractConfigParseConfigPdfOcrHybrid, ExtractRunParamsBodyExtractConfigParseConfigPdfOcrPdf, ExtractRunParamsBodyExtractConfigParseConfigPdfOcrOcr:
		return true
	}
	return false
}

// Forces all external API calls to be routed to specified country/region.
type ExtractRunParamsBodyExtractConfigParseConfigRegionPreference string

const (
	ExtractRunParamsBodyExtractConfigParseConfigRegionPreferenceUs ExtractRunParamsBodyExtractConfigParseConfigRegionPreference = "us"
)

func (r ExtractRunParamsBodyExtractConfigParseConfigRegionPreference) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyExtractConfigParseConfigRegionPreferenceUs:
		return true
	}
	return false
}

// On a spreadsheet, the algorithm that is used to split up sheets into multiple
// tables.
type ExtractRunParamsBodyExtractConfigParseConfigSpreadsheetTableClustering string

const (
	ExtractRunParamsBodyExtractConfigParseConfigSpreadsheetTableClusteringDefault     ExtractRunParamsBodyExtractConfigParseConfigSpreadsheetTableClustering = "default"
	ExtractRunParamsBodyExtractConfigParseConfigSpreadsheetTableClusteringDisabled    ExtractRunParamsBodyExtractConfigParseConfigSpreadsheetTableClustering = "disabled"
	ExtractRunParamsBodyExtractConfigParseConfigSpreadsheetTableClusteringIntelligent ExtractRunParamsBodyExtractConfigParseConfigSpreadsheetTableClustering = "intelligent"
)

func (r ExtractRunParamsBodyExtractConfigParseConfigSpreadsheetTableClustering) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyExtractConfigParseConfigSpreadsheetTableClusteringDefault, ExtractRunParamsBodyExtractConfigParseConfigSpreadsheetTableClusteringDisabled, ExtractRunParamsBodyExtractConfigParseConfigSpreadsheetTableClusteringIntelligent:
		return true
	}
	return false
}

// The mode to use for table output. Defaults to html.
type ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormat string

const (
	ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormatHTML    ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormat = "html"
	ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormatJson    ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormat = "json"
	ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormatMd      ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormat = "md"
	ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormatDynamic ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormat = "dynamic"
	ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormatAIJson  ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormat = "ai_json"
	ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormatCsv     ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormat = "csv"
)

func (r ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormat) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormatHTML, ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormatJson, ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormatMd, ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormatDynamic, ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormatAIJson, ExtractRunParamsBodyExtractConfigParseConfigTableOutputFormatCsv:
		return true
	}
	return false
}

// The version of the processing options.
type ExtractRunParamsBodyExtractConfigParseConfigVersion string

const (
	ExtractRunParamsBodyExtractConfigParseConfigVersionV1 ExtractRunParamsBodyExtractConfigParseConfigVersion = "v1"
	ExtractRunParamsBodyExtractConfigParseConfigVersionV2 ExtractRunParamsBodyExtractConfigParseConfigVersion = "v2"
	ExtractRunParamsBodyExtractConfigParseConfigVersionV3 ExtractRunParamsBodyExtractConfigParseConfigVersion = "v3"
)

func (r ExtractRunParamsBodyExtractConfigParseConfigVersion) IsKnown() bool {
	switch r {
	case ExtractRunParamsBodyExtractConfigParseConfigVersionV1, ExtractRunParamsBodyExtractConfigParseConfigVersionV2, ExtractRunParamsBodyExtractConfigParseConfigVersionV3:
		return true
	}
	return false
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

type ExtractRunJobParams struct {
	Body ExtractRunJobParamsBodyUnion `json:"body,required"`
}

func (r ExtractRunJobParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ExtractRunJobParamsBody struct {
	AdvancedOptions param.Field[shared.AdvancedProcessingOptionsParam] `json:"advanced_options"`
	AgentExtract    param.Field[interface{}]                           `json:"agent_extract"`
	// The configuration options for array extract
	ArrayExtract param.Field[shared.ArrayExtractConfigParam] `json:"array_extract"`
	Async        param.Field[interface{}]                    `json:"async"`
	// The configuration options for citations.
	CitationsOptions    param.Field[shared.AdvancedCitationsConfigParam]       `json:"citations_options"`
	DocumentURL         param.Field[interface{}]                               `json:"document_url"`
	ExperimentalOptions param.Field[shared.ExperimentalProcessingOptionsParam] `json:"experimental_options"`
	// If table citations should be generated for the extracted content.
	ExperimentalTableCitations param.Field[bool] `json:"experimental_table_citations"`
	// If citations should be generated for the extracted content.
	GenerateCitations param.Field[bool] `json:"generate_citations"`
	// If images should be passed directly for extractions. Can only be enabled for
	// documents with less than 10 pages. Defaults to False.
	IncludeImages param.Field[bool]        `json:"include_images"`
	Input         param.Field[interface{}] `json:"input"`
	Instructions  param.Field[interface{}] `json:"instructions"`
	// If True, the job will be processed with lower latency and higher priority. Uses
	// 2x the cost of a regular job. Defaults to False.
	LatencySensitive param.Field[bool]                              `json:"latency_sensitive"`
	Options          param.Field[shared.BaseProcessingOptionsParam] `json:"options"`
	Parsing          param.Field[interface{}]                       `json:"parsing"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool]        `json:"priority"`
	Schema   param.Field[interface{}] `json:"schema"`
	Settings param.Field[interface{}] `json:"settings"`
	// If spreadsheet agent should be used for extraction.
	SpreadsheetAgent param.Field[bool] `json:"spreadsheet_agent"`
	// A system prompt to use for the extraction. This is a general prompt that is
	// applied to the entire document before any other prompts.
	SystemPrompt param.Field[string] `json:"system_prompt"`
	// If chunking should be used for the extraction. Defaults to False.
	UseChunking param.Field[bool]                         `json:"use_chunking"`
	Webhook     param.Field[shared.WebhookConfigNewParam] `json:"webhook"`
}

func (r ExtractRunJobParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunJobParamsBody) implementsExtractRunJobParamsBodyUnion() {}

// Satisfied by [ExtractRunJobParamsBodyAsyncExtractConfigNew],
// [ExtractRunJobParamsBodyAsyncExtractConfig], [ExtractRunJobParamsBody].
type ExtractRunJobParamsBodyUnion interface {
	implementsExtractRunJobParamsBodyUnion()
}

type ExtractRunJobParamsBodyAsyncExtractConfigNew struct {
	// The URL of the document to be processed. You can provide one of the following:
	//
	//  1. A publicly available URL
	//  2. A presigned S3 URL
	//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
	//     uploading a document
	//  4. A job_id (jobid://) or a list of job_ids (jobid://) obtained from a previous
	//     /parse endpoint
	DocumentURL param.Field[ExtractRunJobParamsBodyAsyncExtractConfigNewDocumentURLUnion] `json:"document_url,required"`
	// The JSON schema to use for extraction.
	Schema          param.Field[interface{}]                           `json:"schema,required"`
	AdvancedOptions param.Field[shared.AdvancedProcessingOptionsParam] `json:"advanced_options"`
	// The configuration options for agent extract
	AgentExtract param.Field[ExtractRunJobParamsBodyAsyncExtractConfigNewAgentExtract] `json:"agent_extract"`
	// The configuration options for array extract
	ArrayExtract param.Field[shared.ArrayExtractConfigParam] `json:"array_extract"`
	// The configuration options for citations.
	CitationsOptions    param.Field[shared.AdvancedCitationsConfigParam]       `json:"citations_options"`
	ExperimentalOptions param.Field[shared.ExperimentalProcessingOptionsParam] `json:"experimental_options"`
	// If table citations should be generated for the extracted content.
	ExperimentalTableCitations param.Field[bool] `json:"experimental_table_citations"`
	// If citations should be generated for the extracted content.
	GenerateCitations param.Field[bool] `json:"generate_citations"`
	// If images should be passed directly for extractions. Can only be enabled for
	// documents with less than 10 pages. Defaults to False.
	IncludeImages param.Field[bool] `json:"include_images"`
	// If True, the job will be processed with lower latency and higher priority. Uses
	// 2x the cost of a regular job. Defaults to False.
	LatencySensitive param.Field[bool]                              `json:"latency_sensitive"`
	Options          param.Field[shared.BaseProcessingOptionsParam] `json:"options"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool] `json:"priority"`
	// If spreadsheet agent should be used for extraction.
	SpreadsheetAgent param.Field[bool] `json:"spreadsheet_agent"`
	// A system prompt to use for the extraction. This is a general prompt that is
	// applied to the entire document before any other prompts.
	SystemPrompt param.Field[string] `json:"system_prompt"`
	// If chunking should be used for the extraction. Defaults to False.
	UseChunking param.Field[bool]                         `json:"use_chunking"`
	Webhook     param.Field[shared.WebhookConfigNewParam] `json:"webhook"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigNew) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigNew) implementsExtractRunJobParamsBodyUnion() {}

// The URL of the document to be processed. You can provide one of the following:
//
//  1. A publicly available URL
//  2. A presigned S3 URL
//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
//     uploading a document
//  4. A job_id (jobid://) or a list of job_ids (jobid://) obtained from a previous
//     /parse endpoint
//
// Satisfied by [shared.UnionString],
// [ExtractRunJobParamsBodyAsyncExtractConfigNewDocumentURLArray],
// [shared.UploadParam].
type ExtractRunJobParamsBodyAsyncExtractConfigNewDocumentURLUnion interface {
	ImplementsExtractRunJobParamsBodyAsyncExtractConfigNewDocumentURLUnion()
}

type ExtractRunJobParamsBodyAsyncExtractConfigNewDocumentURLArray []string

func (r ExtractRunJobParamsBodyAsyncExtractConfigNewDocumentURLArray) ImplementsExtractRunJobParamsBodyAsyncExtractConfigNewDocumentURLUnion() {
}

// The configuration options for agent extract
type ExtractRunJobParamsBodyAsyncExtractConfigNewAgentExtract struct {
	// If agent extraction should be used for extraction.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigNewAgentExtract) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunJobParamsBodyAsyncExtractConfig struct {
	// The URL of the document to be processed. You can provide one of the
	// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
	// prefixed URL obtained from the /upload endpoint after directly uploading a
	// document 4. A jobid:// prefixed URL obtained from a previous /parse invocation
	Input param.Field[ExtractRunJobParamsBodyAsyncExtractConfigInputUnion] `json:"input,required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async param.Field[ExtractRunJobParamsBodyAsyncExtractConfigAsync] `json:"async"`
	// The instructions to use for the extraction.
	Instructions param.Field[ExtractRunJobParamsBodyAsyncExtractConfigInstructions] `json:"instructions"`
	// The configuration options for parsing the document. If you are passing in a
	// jobid:// URL for the file, then this configuration will be ignored.
	Parsing param.Field[ExtractRunJobParamsBodyAsyncExtractConfigParsing] `json:"parsing"`
	// The settings to use for the extraction.
	Settings param.Field[ExtractRunJobParamsBodyAsyncExtractConfigSettings] `json:"settings"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunJobParamsBodyAsyncExtractConfig) implementsExtractRunJobParamsBodyUnion() {}

// The URL of the document to be processed. You can provide one of the
// following: 1. A publicly available URL 2. A presigned S3 URL 3. A reducto://
// prefixed URL obtained from the /upload endpoint after directly uploading a
// document 4. A jobid:// prefixed URL obtained from a previous /parse invocation
//
// Satisfied by [shared.UnionString], [shared.UploadParam].
type ExtractRunJobParamsBodyAsyncExtractConfigInputUnion interface {
	ImplementsExtractRunJobParamsBodyAsyncExtractConfigInputUnion()
}

// The configuration options for asynchronous processing (default synchronous).
type ExtractRunJobParamsBodyAsyncExtractConfigAsync struct {
	// JSON metadata included in webhook request body. Defaults to None.
	Metadata param.Field[interface{}] `json:"metadata"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool] `json:"priority"`
	// The webhook configuration for the asynchronous processing.
	Webhook param.Field[ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookUnion] `json:"webhook"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigAsync) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The webhook configuration for the asynchronous processing.
type ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhook struct {
	Channels param.Field[interface{}]                                               `json:"channels"`
	Mode     param.Field[ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookMode] `json:"mode"`
	URL      param.Field[string]                                                    `json:"url"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhook) implementsExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookUnion() {
}

// The webhook configuration for the asynchronous processing.
//
// Satisfied by
// [ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfig],
// [ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfig],
// [ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhook].
type ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookUnion interface {
	implementsExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookUnion()
}

type ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfig struct {
	// A list of Svix channels the message will be delivered down, omit to send to all
	// channels.
	Channels param.Field[[]string]                                                                   `json:"channels"`
	Mode     param.Field[ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfigMode] `json:"mode"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfig) implementsExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookUnion() {
}

type ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfigMode string

const (
	ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfigModeSvix ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfigMode = "svix"
)

func (r ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfigMode) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookSvixWebhookConfigModeSvix:
		return true
	}
	return false
}

type ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfig struct {
	URL  param.Field[string]                                                                       `json:"url,required"`
	Mode param.Field[ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfigMode] `json:"mode"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfig) implementsExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookUnion() {
}

type ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfigMode string

const (
	ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfigModeDirect ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfigMode = "direct"
)

func (r ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfigMode) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookDirectWebhookConfigModeDirect:
		return true
	}
	return false
}

type ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookMode string

const (
	ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookModeSvix   ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookMode = "svix"
	ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookModeDirect ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookMode = "direct"
)

func (r ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookMode) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookModeSvix, ExtractRunJobParamsBodyAsyncExtractConfigAsyncWebhookModeDirect:
		return true
	}
	return false
}

// The instructions to use for the extraction.
type ExtractRunJobParamsBodyAsyncExtractConfigInstructions struct {
	// The JSON schema to use for the extraction.
	Schema param.Field[interface{}] `json:"schema"`
	// The system prompt to use for the extraction.
	SystemPrompt param.Field[string] `json:"system_prompt"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigInstructions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration options for parsing the document. If you are passing in a
// jobid:// URL for the file, then this configuration will be ignored.
type ExtractRunJobParamsBodyAsyncExtractConfigParsing struct {
	Enhance     param.Field[ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhance]     `json:"enhance"`
	Formatting  param.Field[ExtractRunJobParamsBodyAsyncExtractConfigParsingFormatting]  `json:"formatting"`
	Retrieval   param.Field[ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrieval]   `json:"retrieval"`
	Settings    param.Field[ExtractRunJobParamsBodyAsyncExtractConfigParsingSettings]    `json:"settings"`
	Spreadsheet param.Field[ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheet] `json:"spreadsheet"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsing) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhance struct {
	// Agentic uses vision language models to enhance the accuracy of the output of
	// different types of extraction. This will incur a cost and latency increase.
	Agentic param.Field[[]ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticUnion] `json:"agentic"`
	// If True, summarize figures using a small vision language model. Defaults to
	// True.
	SummarizeFigures param.Field[bool] `json:"summarize_figures"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhance) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgentic struct {
	Scope param.Field[ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticScope] `json:"scope,required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgentic) implementsExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticUnion() {
}

// Satisfied by
// [ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgentic],
// [ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgentic],
// [ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgentic],
// [ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgentic].
type ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticUnion interface {
	implementsExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticUnion()
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgentic struct {
	Scope param.Field[ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgenticScope] `json:"scope,required"`
	// Custom prompt for table agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgentic) implementsExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticUnion() {
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgenticScope string

const (
	ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgenticScopeTable ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgenticScope = "table"
)

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgenticScope) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTableAgenticScopeTable:
		return true
	}
	return false
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgentic struct {
	Scope param.Field[ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgenticScope] `json:"scope,required"`
	// Custom prompt for figure agentic.
	Prompt param.Field[string] `json:"prompt"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgentic) implementsExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticUnion() {
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgenticScope string

const (
	ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgenticScopeFigure ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgenticScope = "figure"
)

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgenticScope) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticFigureAgenticScopeFigure:
		return true
	}
	return false
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgentic struct {
	Scope param.Field[ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgenticScope] `json:"scope,required"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgentic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgentic) implementsExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticUnion() {
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgenticScope string

const (
	ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgenticScopeText ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgenticScope = "text"
)

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgenticScope) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticTextAgenticScopeText:
		return true
	}
	return false
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticScope string

const (
	ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticScopeTable  ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticScope = "table"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticScopeFigure ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticScope = "figure"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticScopeText   ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticScope = "text"
)

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticScope) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticScopeTable, ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticScopeFigure, ExtractRunJobParamsBodyAsyncExtractConfigParsingEnhanceAgenticScopeText:
		return true
	}
	return false
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingFormatting struct {
	// If True, add page markers to the output. Defaults to False. Useful for
	// extracting data with page specific information.
	AddPageMarkers param.Field[bool] `json:"add_page_markers"`
	// A list of formatting to include in the output. [insert description of each
	// option here later]
	Include param.Field[[]ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingInclude] `json:"include"`
	// A flag to indicate if consecutive tables with the same number of columns should
	// be merged. Defaults to False.
	MergeTables param.Field[bool] `json:"merge_tables"`
	// The mode to use for table output. Defaults to dynamic, which returns md for
	// simpler tables and html for more complex tables.
	TableOutputFormat param.Field[ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormat] `json:"table_output_format"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingFormatting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingInclude string

const (
	ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingIncludeChangeTracking ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingInclude = "change_tracking"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingIncludeHighlight      ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingInclude = "highlight"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingIncludeComments       ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingInclude = "comments"
)

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingInclude) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingIncludeChangeTracking, ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingIncludeHighlight, ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingIncludeComments:
		return true
	}
	return false
}

// The mode to use for table output. Defaults to dynamic, which returns md for
// simpler tables and html for more complex tables.
type ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormat string

const (
	ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatHTML     ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormat = "html"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatJson     ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormat = "json"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatMd       ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormat = "md"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatJsonbbox ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormat = "jsonbbox"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatDynamic  ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormat = "dynamic"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatCsv      ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormat = "csv"
)

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormat) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatHTML, ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatJson, ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatMd, ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatJsonbbox, ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatDynamic, ExtractRunJobParamsBodyAsyncExtractConfigParsingFormattingTableOutputFormatCsv:
		return true
	}
	return false
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrieval struct {
	Chunking param.Field[ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunking] `json:"chunking"`
	// If True, use embedding optimized mode. Defaults to False.
	EmbeddingOptimized param.Field[bool] `json:"embedding_optimized"`
	// A list of block types to filter out from 'content' and 'embed' fields. By
	// default, no blocks are filtered.
	FilterBlocks param.Field[[]ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock] `json:"filter_blocks"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrieval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunking struct {
	// Choose how to partition chunks. Variable mode chunks by character length and
	// visual context. Section mode chunks by section headers. Page mode chunks
	// according to pages. Page sections mode chunks first by page, then by sections
	// within each page. Disabled returns one single chunk.
	ChunkMode param.Field[ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkMode] `json:"chunk_mode"`
	// The approximate size of chunks (in characters) that the document will be split
	// into. Defaults to null, in which case the chunk size is variable between 250 -
	// 1500 characters.
	ChunkSize param.Field[int64] `json:"chunk_size"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Choose how to partition chunks. Variable mode chunks by character length and
// visual context. Section mode chunks by section headers. Page mode chunks
// according to pages. Page sections mode chunks first by page, then by sections
// within each page. Disabled returns one single chunk.
type ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkMode string

const (
	ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModeVariable     ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkMode = "variable"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModeSection      ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkMode = "section"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModePage         ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkMode = "page"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModeDisabled     ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkMode = "disabled"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModeBlock        ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkMode = "block"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModePageSections ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkMode = "page_sections"
)

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkMode) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModeVariable, ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModeSection, ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModePage, ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModeDisabled, ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModeBlock, ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalChunkingChunkModePageSections:
		return true
	}
	return false
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock string

const (
	ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockHeader        ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Header"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockFooter        ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Footer"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockTitle         ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Title"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockSectionHeader ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Section Header"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockPageNumber    ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Page Number"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockListItem      ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "List Item"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockFigure        ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Figure"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockTable         ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Table"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockKeyValue      ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Key Value"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockText          ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Text"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockComment       ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Comment"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockSignature     ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock = "Signature"
)

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlock) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockHeader, ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockFooter, ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockTitle, ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockSectionHeader, ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockPageNumber, ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockListItem, ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockFigure, ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockTable, ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockKeyValue, ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockText, ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockComment, ExtractRunJobParamsBodyAsyncExtractConfigParsingRetrievalFilterBlockSignature:
		return true
	}
	return false
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingSettings struct {
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
	OcrSystem param.Field[ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsOcrSystem] `json:"ocr_system"`
	// The page range to process (1-indexed). By default, the entire document is
	// processed.
	PageRange param.Field[ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsPageRangeUnion] `json:"page_range"`
	// If True, persist the results indefinitely. Defaults to False.
	PersistResults param.Field[bool] `json:"persist_results"`
	// Whether to return images for the specified block types. By default, no images
	// are returned.
	ReturnImages param.Field[[]ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsReturnImage] `json:"return_images"`
	// If True, return OCR data in the result. Defaults to False.
	ReturnOcrData param.Field[bool] `json:"return_ocr_data"`
	// The timeout for the job in seconds. Defaults to 900.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Standard is our best multilingual OCR system. Legacy only supports germanic
// languages and is available for backwards compatibility.
type ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsOcrSystem string

const (
	ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsOcrSystemStandard ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsOcrSystem = "standard"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsOcrSystemLegacy   ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsOcrSystem = "legacy"
)

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsOcrSystem) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsOcrSystemStandard, ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsOcrSystemLegacy:
		return true
	}
	return false
}

// The page range to process (1-indexed). By default, the entire document is
// processed.
//
// Satisfied by [shared.PageRangeParam],
// [ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsPageRangeArray],
// [ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsPageRangeArray].
type ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsPageRangeUnion interface {
	ImplementsExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsPageRangeUnion()
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsPageRangeArray []shared.PageRangeParam

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsPageRangeArray) ImplementsExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsPageRangeUnion() {
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsReturnImage string

const (
	ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsReturnImageFigure ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsReturnImage = "figure"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsReturnImageTable  ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsReturnImage = "table"
)

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsReturnImage) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsReturnImageFigure, ExtractRunJobParamsBodyAsyncExtractConfigParsingSettingsReturnImageTable:
		return true
	}
	return false
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheet struct {
	// In a spreadsheet with different tables inside, we enable splitting up the tables
	// by default. Accurate mode applies more powerful models for superior accuracy, at
	// 5× the default per-cell rate. Disabling will register as one large table.
	Clustering param.Field[ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetClustering] `json:"clustering"`
	// Whether to exclude hidden sheets, rows, or columns in the output.
	Exclude param.Field[[]ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetExclude] `json:"exclude"`
	// Whether to include cell color and formula information in the output.
	Include          param.Field[[]ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetInclude]        `json:"include"`
	SplitLargeTables param.Field[ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetSplitLargeTables] `json:"split_large_tables"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// In a spreadsheet with different tables inside, we enable splitting up the tables
// by default. Accurate mode applies more powerful models for superior accuracy, at
// 5× the default per-cell rate. Disabling will register as one large table.
type ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetClustering string

const (
	ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetClusteringAccurate ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetClustering = "accurate"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetClusteringFast     ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetClustering = "fast"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetClusteringDisabled ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetClustering = "disabled"
)

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetClustering) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetClusteringAccurate, ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetClusteringFast, ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetClusteringDisabled:
		return true
	}
	return false
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetExclude string

const (
	ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetExcludeHiddenSheets ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetExclude = "hidden_sheets"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetExcludeHiddenRows   ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetExclude = "hidden_rows"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetExcludeHiddenCols   ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetExclude = "hidden_cols"
)

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetExclude) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetExcludeHiddenSheets, ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetExcludeHiddenRows, ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetExcludeHiddenCols:
		return true
	}
	return false
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetInclude string

const (
	ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetIncludeCellColors ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetInclude = "cell_colors"
	ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetIncludeFormula    ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetInclude = "formula"
)

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetInclude) IsKnown() bool {
	switch r {
	case ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetIncludeCellColors, ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetIncludeFormula:
		return true
	}
	return false
}

type ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetSplitLargeTables struct {
	// If True, split large tables into smaller tables. Defaults to True.
	Enabled param.Field[bool] `json:"enabled"`
	// The size of the tables to split into. Defaults to 50.
	Size param.Field[int64] `json:"size"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigParsingSpreadsheetSplitLargeTables) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The settings to use for the extraction.
type ExtractRunJobParamsBodyAsyncExtractConfigSettings struct {
	// If True, use array extraction.
	ArrayExtract param.Field[bool] `json:"array_extract"`
	// The citations to use for the extraction.
	Citations param.Field[ExtractRunJobParamsBodyAsyncExtractConfigSettingsCitations] `json:"citations"`
	// If True, include images in the extraction.
	IncludeImages param.Field[bool] `json:"include_images"`
	// If True, jobs will be processed with a higher throughput and priority at a
	// higher cost. Defaults to False.
	OptimizeForLatency param.Field[bool] `json:"optimize_for_latency"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The citations to use for the extraction.
type ExtractRunJobParamsBodyAsyncExtractConfigSettingsCitations struct {
	// If True, include citations in the extraction.
	Enabled param.Field[bool] `json:"enabled"`
	// If True, enable numeric citation confidence scores. Defaults to True.
	NumericalConfidence param.Field[bool] `json:"numerical_confidence"`
}

func (r ExtractRunJobParamsBodyAsyncExtractConfigSettingsCitations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
