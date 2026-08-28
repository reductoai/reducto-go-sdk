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
	return res, err
}

// Split Async
func (r *SplitService) RunJob(ctx context.Context, body SplitRunJobParams, opts ...option.RequestOption) (res *shared.AsyncSplitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "split_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type DeepSplitPageEvidence struct {
	Evidence   string                          `json:"evidence" api:"required"`
	PageNumber int64                           `json:"page_number" api:"required"`
	Confidence DeepSplitPageEvidenceConfidence `json:"confidence" api:"nullable"`
	JSON       deepSplitPageEvidenceJSON       `json:"-"`
}

// deepSplitPageEvidenceJSON contains the JSON metadata for the struct
// [DeepSplitPageEvidence]
type deepSplitPageEvidenceJSON struct {
	Evidence    apijson.Field
	PageNumber  apijson.Field
	Confidence  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeepSplitPageEvidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deepSplitPageEvidenceJSON) RawJSON() string {
	return r.raw
}

type DeepSplitPageEvidenceConfidence string

const (
	DeepSplitPageEvidenceConfidenceHigh   DeepSplitPageEvidenceConfidence = "high"
	DeepSplitPageEvidenceConfidenceMedium DeepSplitPageEvidenceConfidence = "medium"
	DeepSplitPageEvidenceConfidenceLow    DeepSplitPageEvidenceConfidence = "low"
)

func (r DeepSplitPageEvidenceConfidence) IsKnown() bool {
	switch r {
	case DeepSplitPageEvidenceConfidenceHigh, DeepSplitPageEvidenceConfidenceMedium, DeepSplitPageEvidenceConfidenceLow:
		return true
	}
	return false
}

type ParseUsage struct {
	NumPages        int64              `json:"num_pages" api:"required"`
	CreditBreakdown map[string]float64 `json:"credit_breakdown" api:"nullable"`
	Credits         float64            `json:"credits" api:"nullable"`
	// Total non-empty cells across all sheets. Only set for spreadsheet inputs.
	NonEmptyCellCount int64 `json:"non_empty_cell_count" api:"nullable"`
	// Per-page breakdown of features used. Maps 1-indexed page numbers (as strings) to
	// the list of billing features applied on that page (e.g. 'page', 'complex',
	// 'chart_agent').
	PageBillingBreakdown map[string][]ParseUsagePageBillingBreakdown `json:"page_billing_breakdown" api:"nullable"`
	// Raw usage quantities. Only set for accounts on the new pricing model; credit
	// fields are omitted for those accounts.
	UsageBreakdown ParseUsageUsageBreakdown `json:"usage_breakdown" api:"nullable"`
	JSON           parseUsageJSON           `json:"-"`
}

// parseUsageJSON contains the JSON metadata for the struct [ParseUsage]
type parseUsageJSON struct {
	NumPages             apijson.Field
	CreditBreakdown      apijson.Field
	Credits              apijson.Field
	NonEmptyCellCount    apijson.Field
	PageBillingBreakdown apijson.Field
	UsageBreakdown       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ParseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseUsageJSON) RawJSON() string {
	return r.raw
}

type ParseUsagePageBillingBreakdown string

const (
	ParseUsagePageBillingBreakdownPage                     ParseUsagePageBillingBreakdown = "page"
	ParseUsagePageBillingBreakdownHTMLPage                 ParseUsagePageBillingBreakdown = "html_page"
	ParseUsagePageBillingBreakdownDocxNativePage           ParseUsagePageBillingBreakdown = "docx_native_page"
	ParseUsagePageBillingBreakdownAgentic                  ParseUsagePageBillingBreakdown = "agentic"
	ParseUsagePageBillingBreakdownComplex                  ParseUsagePageBillingBreakdown = "complex"
	ParseUsagePageBillingBreakdownChartAgent               ParseUsagePageBillingBreakdown = "chart_agent"
	ParseUsagePageBillingBreakdownSpreadsheetCells         ParseUsagePageBillingBreakdown = "spreadsheet_cells"
	ParseUsagePageBillingBreakdownBillableSpreadsheetPages ParseUsagePageBillingBreakdown = "billable_spreadsheet_pages"
	ParseUsagePageBillingBreakdownEnrichTable              ParseUsagePageBillingBreakdown = "enrich_table"
	ParseUsagePageBillingBreakdownFigureSummary            ParseUsagePageBillingBreakdown = "figure_summary"
	ParseUsagePageBillingBreakdownTableSummary             ParseUsagePageBillingBreakdown = "table_summary"
	ParseUsagePageBillingBreakdownKeyValue                 ParseUsagePageBillingBreakdown = "key_value"
	ParseUsagePageBillingBreakdownAgenticText              ParseUsagePageBillingBreakdown = "agentic_text"
	ParseUsagePageBillingBreakdownPromptableAgenticText    ParseUsagePageBillingBreakdown = "promptable_agentic_text"
	ParseUsagePageBillingBreakdownReductoLitePage          ParseUsagePageBillingBreakdown = "reducto_lite_page"
)

func (r ParseUsagePageBillingBreakdown) IsKnown() bool {
	switch r {
	case ParseUsagePageBillingBreakdownPage, ParseUsagePageBillingBreakdownHTMLPage, ParseUsagePageBillingBreakdownDocxNativePage, ParseUsagePageBillingBreakdownAgentic, ParseUsagePageBillingBreakdownComplex, ParseUsagePageBillingBreakdownChartAgent, ParseUsagePageBillingBreakdownSpreadsheetCells, ParseUsagePageBillingBreakdownBillableSpreadsheetPages, ParseUsagePageBillingBreakdownEnrichTable, ParseUsagePageBillingBreakdownFigureSummary, ParseUsagePageBillingBreakdownTableSummary, ParseUsagePageBillingBreakdownKeyValue, ParseUsagePageBillingBreakdownAgenticText, ParseUsagePageBillingBreakdownPromptableAgenticText, ParseUsagePageBillingBreakdownReductoLitePage:
		return true
	}
	return false
}

// Raw usage quantities. Only set for accounts on the new pricing model; credit
// fields are omitted for those accounts.
type ParseUsageUsageBreakdown struct {
	Charts             int64                              `json:"charts"`
	EditModel          ParseUsageUsageBreakdownEditModel  `json:"edit_model"`
	EditPages          int64                              `json:"edit_pages"`
	LegacyParseCredits float64                            `json:"legacy_parse_credits"`
	OcrPages           int64                              `json:"ocr_pages"`
	ParseModel         ParseUsageUsageBreakdownParseModel `json:"parse_model"`
	ParseNativePages   int64                              `json:"parse_native_pages"`
	ParsePages         int64                              `json:"parse_pages"`
	PrefillPages       int64                              `json:"prefill_pages"`
	PromptedBlocks     int64                              `json:"prompted_blocks"`
	SplitModel         ParseUsageUsageBreakdownSplitModel `json:"split_model"`
	SplitPages         int64                              `json:"split_pages"`
	Tier               ParseUsageUsageBreakdownTier       `json:"tier"`
	JSON               parseUsageUsageBreakdownJSON       `json:"-"`
	union              ParseUsageUsageBreakdownUnion
}

// parseUsageUsageBreakdownJSON contains the JSON metadata for the struct
// [ParseUsageUsageBreakdown]
type parseUsageUsageBreakdownJSON struct {
	Charts             apijson.Field
	EditModel          apijson.Field
	EditPages          apijson.Field
	LegacyParseCredits apijson.Field
	OcrPages           apijson.Field
	ParseModel         apijson.Field
	ParseNativePages   apijson.Field
	ParsePages         apijson.Field
	PrefillPages       apijson.Field
	PromptedBlocks     apijson.Field
	SplitModel         apijson.Field
	SplitPages         apijson.Field
	Tier               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r parseUsageUsageBreakdownJSON) RawJSON() string {
	return r.raw
}

func (r *ParseUsageUsageBreakdown) UnmarshalJSON(data []byte) (err error) {
	*r = ParseUsageUsageBreakdown{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ParseUsageUsageBreakdownUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ParseUsageUsageBreakdownParseUsageBreakdown],
// [ParseUsageUsageBreakdownSplitUsageBreakdown],
// [ParseUsageUsageBreakdownEditUsageBreakdown].
func (r ParseUsageUsageBreakdown) AsUnion() ParseUsageUsageBreakdownUnion {
	return r.union
}

// Raw usage quantities. Only set for accounts on the new pricing model; credit
// fields are omitted for those accounts.
//
// Union satisfied by [ParseUsageUsageBreakdownParseUsageBreakdown],
// [ParseUsageUsageBreakdownSplitUsageBreakdown] or
// [ParseUsageUsageBreakdownEditUsageBreakdown].
type ParseUsageUsageBreakdownUnion interface {
	implementsParseUsageUsageBreakdown()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ParseUsageUsageBreakdownUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParseUsageUsageBreakdownParseUsageBreakdown{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParseUsageUsageBreakdownSplitUsageBreakdown{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParseUsageUsageBreakdownEditUsageBreakdown{}),
		},
	)
}

// Raw parse quantities for accounts on the new (Q3 2026) pricing model.
//
// `parse_model` is "R-1" for the new parse model and "Legacy" for the legacy parse
// pipeline. A legacy-pipeline parse carries its cost in `legacy_parse_credits`;
// add-on quantities (`ocr_pages`, `charts`, `prompted_blocks`) apply to the new
// parse model only.
type ParseUsageUsageBreakdownParseUsageBreakdown struct {
	ParseModel         ParseUsageUsageBreakdownParseUsageBreakdownParseModel `json:"parse_model" api:"required"`
	Tier               ParseUsageUsageBreakdownParseUsageBreakdownTier       `json:"tier" api:"required"`
	Charts             int64                                                 `json:"charts"`
	LegacyParseCredits float64                                               `json:"legacy_parse_credits"`
	OcrPages           int64                                                 `json:"ocr_pages"`
	ParseNativePages   int64                                                 `json:"parse_native_pages"`
	ParsePages         int64                                                 `json:"parse_pages"`
	PromptedBlocks     int64                                                 `json:"prompted_blocks"`
	JSON               parseUsageUsageBreakdownParseUsageBreakdownJSON       `json:"-"`
}

// parseUsageUsageBreakdownParseUsageBreakdownJSON contains the JSON metadata for
// the struct [ParseUsageUsageBreakdownParseUsageBreakdown]
type parseUsageUsageBreakdownParseUsageBreakdownJSON struct {
	ParseModel         apijson.Field
	Tier               apijson.Field
	Charts             apijson.Field
	LegacyParseCredits apijson.Field
	OcrPages           apijson.Field
	ParseNativePages   apijson.Field
	ParsePages         apijson.Field
	PromptedBlocks     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ParseUsageUsageBreakdownParseUsageBreakdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseUsageUsageBreakdownParseUsageBreakdownJSON) RawJSON() string {
	return r.raw
}

func (r ParseUsageUsageBreakdownParseUsageBreakdown) implementsParseUsageUsageBreakdown() {}

type ParseUsageUsageBreakdownParseUsageBreakdownParseModel string

const (
	ParseUsageUsageBreakdownParseUsageBreakdownParseModelR1     ParseUsageUsageBreakdownParseUsageBreakdownParseModel = "R-1"
	ParseUsageUsageBreakdownParseUsageBreakdownParseModelLegacy ParseUsageUsageBreakdownParseUsageBreakdownParseModel = "Legacy"
)

func (r ParseUsageUsageBreakdownParseUsageBreakdownParseModel) IsKnown() bool {
	switch r {
	case ParseUsageUsageBreakdownParseUsageBreakdownParseModelR1, ParseUsageUsageBreakdownParseUsageBreakdownParseModelLegacy:
		return true
	}
	return false
}

type ParseUsageUsageBreakdownParseUsageBreakdownTier string

const (
	ParseUsageUsageBreakdownParseUsageBreakdownTierDefault ParseUsageUsageBreakdownParseUsageBreakdownTier = "Default"
	ParseUsageUsageBreakdownParseUsageBreakdownTierBatch   ParseUsageUsageBreakdownParseUsageBreakdownTier = "Batch"
)

func (r ParseUsageUsageBreakdownParseUsageBreakdownTier) IsKnown() bool {
	switch r {
	case ParseUsageUsageBreakdownParseUsageBreakdownTierDefault, ParseUsageUsageBreakdownParseUsageBreakdownTierBatch:
		return true
	}
	return false
}

// Raw split quantities for accounts on the new pricing model.
//
// The add-on quantities (`ocr_pages`, `charts`, `prompted_blocks`) come from the
// parse bundled into the split job; its page cost is covered by `split_pages` but
// its add-ons are billed separately.
type ParseUsageUsageBreakdownSplitUsageBreakdown struct {
	SplitModel     ParseUsageUsageBreakdownSplitUsageBreakdownSplitModel `json:"split_model" api:"required"`
	Charts         int64                                                 `json:"charts"`
	OcrPages       int64                                                 `json:"ocr_pages"`
	PromptedBlocks int64                                                 `json:"prompted_blocks"`
	SplitPages     int64                                                 `json:"split_pages"`
	JSON           parseUsageUsageBreakdownSplitUsageBreakdownJSON       `json:"-"`
}

// parseUsageUsageBreakdownSplitUsageBreakdownJSON contains the JSON metadata for
// the struct [ParseUsageUsageBreakdownSplitUsageBreakdown]
type parseUsageUsageBreakdownSplitUsageBreakdownJSON struct {
	SplitModel     apijson.Field
	Charts         apijson.Field
	OcrPages       apijson.Field
	PromptedBlocks apijson.Field
	SplitPages     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ParseUsageUsageBreakdownSplitUsageBreakdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseUsageUsageBreakdownSplitUsageBreakdownJSON) RawJSON() string {
	return r.raw
}

func (r ParseUsageUsageBreakdownSplitUsageBreakdown) implementsParseUsageUsageBreakdown() {}

type ParseUsageUsageBreakdownSplitUsageBreakdownSplitModel string

const (
	ParseUsageUsageBreakdownSplitUsageBreakdownSplitModelSplit     ParseUsageUsageBreakdownSplitUsageBreakdownSplitModel = "Split"
	ParseUsageUsageBreakdownSplitUsageBreakdownSplitModelDeepSplit ParseUsageUsageBreakdownSplitUsageBreakdownSplitModel = "Deep Split"
)

func (r ParseUsageUsageBreakdownSplitUsageBreakdownSplitModel) IsKnown() bool {
	switch r {
	case ParseUsageUsageBreakdownSplitUsageBreakdownSplitModelSplit, ParseUsageUsageBreakdownSplitUsageBreakdownSplitModelDeepSplit:
		return true
	}
	return false
}

// Raw edit quantities for accounts on the new pricing model.
//
// `edit_pages` is the page count billed at the `edit_model` rate. A job with both
// normal and prefilled pages reports `edit_model="Normal"` with the prefilled
// pages in `prefill_pages`, billed at the "Prefill" rate.
type ParseUsageUsageBreakdownEditUsageBreakdown struct {
	EditModel    ParseUsageUsageBreakdownEditUsageBreakdownEditModel `json:"edit_model" api:"required"`
	EditPages    int64                                               `json:"edit_pages"`
	PrefillPages int64                                               `json:"prefill_pages"`
	JSON         parseUsageUsageBreakdownEditUsageBreakdownJSON      `json:"-"`
}

// parseUsageUsageBreakdownEditUsageBreakdownJSON contains the JSON metadata for
// the struct [ParseUsageUsageBreakdownEditUsageBreakdown]
type parseUsageUsageBreakdownEditUsageBreakdownJSON struct {
	EditModel    apijson.Field
	EditPages    apijson.Field
	PrefillPages apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ParseUsageUsageBreakdownEditUsageBreakdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseUsageUsageBreakdownEditUsageBreakdownJSON) RawJSON() string {
	return r.raw
}

func (r ParseUsageUsageBreakdownEditUsageBreakdown) implementsParseUsageUsageBreakdown() {}

type ParseUsageUsageBreakdownEditUsageBreakdownEditModel string

const (
	ParseUsageUsageBreakdownEditUsageBreakdownEditModelNormal  ParseUsageUsageBreakdownEditUsageBreakdownEditModel = "Normal"
	ParseUsageUsageBreakdownEditUsageBreakdownEditModelPrefill ParseUsageUsageBreakdownEditUsageBreakdownEditModel = "Prefill"
)

func (r ParseUsageUsageBreakdownEditUsageBreakdownEditModel) IsKnown() bool {
	switch r {
	case ParseUsageUsageBreakdownEditUsageBreakdownEditModelNormal, ParseUsageUsageBreakdownEditUsageBreakdownEditModelPrefill:
		return true
	}
	return false
}

type ParseUsageUsageBreakdownEditModel string

const (
	ParseUsageUsageBreakdownEditModelNormal  ParseUsageUsageBreakdownEditModel = "Normal"
	ParseUsageUsageBreakdownEditModelPrefill ParseUsageUsageBreakdownEditModel = "Prefill"
)

func (r ParseUsageUsageBreakdownEditModel) IsKnown() bool {
	switch r {
	case ParseUsageUsageBreakdownEditModelNormal, ParseUsageUsageBreakdownEditModelPrefill:
		return true
	}
	return false
}

type ParseUsageUsageBreakdownParseModel string

const (
	ParseUsageUsageBreakdownParseModelR1     ParseUsageUsageBreakdownParseModel = "R-1"
	ParseUsageUsageBreakdownParseModelLegacy ParseUsageUsageBreakdownParseModel = "Legacy"
)

func (r ParseUsageUsageBreakdownParseModel) IsKnown() bool {
	switch r {
	case ParseUsageUsageBreakdownParseModelR1, ParseUsageUsageBreakdownParseModelLegacy:
		return true
	}
	return false
}

type ParseUsageUsageBreakdownSplitModel string

const (
	ParseUsageUsageBreakdownSplitModelSplit     ParseUsageUsageBreakdownSplitModel = "Split"
	ParseUsageUsageBreakdownSplitModelDeepSplit ParseUsageUsageBreakdownSplitModel = "Deep Split"
)

func (r ParseUsageUsageBreakdownSplitModel) IsKnown() bool {
	switch r {
	case ParseUsageUsageBreakdownSplitModelSplit, ParseUsageUsageBreakdownSplitModelDeepSplit:
		return true
	}
	return false
}

type ParseUsageUsageBreakdownTier string

const (
	ParseUsageUsageBreakdownTierDefault ParseUsageUsageBreakdownTier = "Default"
	ParseUsageUsageBreakdownTierBatch   ParseUsageUsageBreakdownTier = "Batch"
)

func (r ParseUsageUsageBreakdownTier) IsKnown() bool {
	switch r {
	case ParseUsageUsageBreakdownTierDefault, ParseUsageUsageBreakdownTierBatch:
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

type SplitRunParams struct {
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input param.Field[SplitRunParamsInputUnion] `json:"input" api:"required"`
	// The configuration options for processing the document.
	SplitDescription param.Field[[]SplitCategoryParam] `json:"split_description" api:"required"`
	// The configuration options for parsing the document. If you are passing in a
	// jobid:// URL for the file, then this configuration will be ignored.
	Parsing param.Field[ParseOptionsParam] `json:"parsing"`
	// The settings for split processing.
	Settings param.Field[SplitRunParamsSettings] `json:"settings"`
	// The prompt that describes rules for splitting the document.
	SplitRules param.Field[string] `json:"split_rules"`
}

func (r SplitRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For parse/split/extract pipelines, the URL of the document to be processed. You
// can provide one of the following: 1. A publicly available URL 2. A presigned S3
// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
// directly uploading a document 4. A jobid:// prefixed URL obtained from a
// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
// API only)
//
//	For edit pipelines, this should be a string containing the edit instructions
//
// Satisfied by [shared.UnionString], [SplitRunParamsInputArray],
// [shared.UploadParam].
type SplitRunParamsInputUnion interface {
	ImplementsSplitRunParamsInputUnion()
}

type SplitRunParamsInputArray []string

func (r SplitRunParamsInputArray) ImplementsSplitRunParamsInputUnion() {}

// The settings for split processing.
type SplitRunParamsSettings struct {
	// If True, a page can belong to multiple categories/partitions. If False, each
	// page must belong to exactly one category. Defaults to True.
	AllowPageOverlap param.Field[bool] `json:"allow_page_overlap"`
	// If True (default), deep split may split a category into partitions even when
	// that category has no configured partition_key. If False, categories without a
	// partition_key are never partitioned, so partitioning happens only where you
	// explicitly configured a partition_key.
	AutoPartition param.Field[bool] `json:"auto_partition"`
	// If True, uses the deep split agent for higher-quality document splitting. Off by
	// default.
	DeepSplit param.Field[bool] `json:"deep_split"`
	// Force the endpoint result to be returned in URL form.
	ForceURLResult param.Field[bool] `json:"force_url_result"`
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
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input param.Field[SplitRunJobParamsInputUnion] `json:"input" api:"required"`
	// The configuration options for processing the document.
	SplitDescription param.Field[[]SplitCategoryParam] `json:"split_description" api:"required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async param.Field[AsyncConfigV3Param] `json:"async"`
	// The configuration options for parsing the document. If you are passing in a
	// jobid:// URL for the file, then this configuration will be ignored.
	Parsing param.Field[ParseOptionsParam] `json:"parsing"`
	// The settings for split processing.
	Settings param.Field[SplitRunJobParamsSettings] `json:"settings"`
	// The prompt that describes rules for splitting the document.
	SplitRules param.Field[string] `json:"split_rules"`
}

func (r SplitRunJobParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For parse/split/extract pipelines, the URL of the document to be processed. You
// can provide one of the following: 1. A publicly available URL 2. A presigned S3
// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
// directly uploading a document 4. A jobid:// prefixed URL obtained from a
// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
// API only)
//
//	For edit pipelines, this should be a string containing the edit instructions
//
// Satisfied by [shared.UnionString], [SplitRunJobParamsInputArray],
// [shared.UploadParam].
type SplitRunJobParamsInputUnion interface {
	ImplementsSplitRunJobParamsInputUnion()
}

type SplitRunJobParamsInputArray []string

func (r SplitRunJobParamsInputArray) ImplementsSplitRunJobParamsInputUnion() {}

// The settings for split processing.
type SplitRunJobParamsSettings struct {
	// If True, a page can belong to multiple categories/partitions. If False, each
	// page must belong to exactly one category. Defaults to True.
	AllowPageOverlap param.Field[bool] `json:"allow_page_overlap"`
	// If True (default), deep split may split a category into partitions even when
	// that category has no configured partition_key. If False, categories without a
	// partition_key are never partitioned, so partitioning happens only where you
	// explicitly configured a partition_key.
	AutoPartition param.Field[bool] `json:"auto_partition"`
	// If True, uses the deep split agent for higher-quality document splitting. Off by
	// default.
	DeepSplit param.Field[bool] `json:"deep_split"`
	// Force the endpoint result to be returned in URL form.
	ForceURLResult param.Field[bool] `json:"force_url_result"`
	// If tables should be truncated to the first few rows or if all content should be
	// preserved. truncate improves latency, preserve is recommended for cases where
	// partition_key is being used and the partition_key may be included within the
	// table. Defaults to truncate
	TableCutoff param.Field[SplitRunJobParamsSettingsTableCutoff] `json:"table_cutoff"`
}

func (r SplitRunJobParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If tables should be truncated to the first few rows or if all content should be
// preserved. truncate improves latency, preserve is recommended for cases where
// partition_key is being used and the partition_key may be included within the
// table. Defaults to truncate
type SplitRunJobParamsSettingsTableCutoff string

const (
	SplitRunJobParamsSettingsTableCutoffTruncate SplitRunJobParamsSettingsTableCutoff = "truncate"
	SplitRunJobParamsSettingsTableCutoffPreserve SplitRunJobParamsSettingsTableCutoff = "preserve"
)

func (r SplitRunJobParamsSettingsTableCutoff) IsKnown() bool {
	switch r {
	case SplitRunJobParamsSettingsTableCutoffTruncate, SplitRunJobParamsSettingsTableCutoffPreserve:
		return true
	}
	return false
}
