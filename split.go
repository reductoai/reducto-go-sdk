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
	// Per-page breakdown of features used. Maps 1-indexed page numbers (as strings) to
	// the list of billing features applied on that page (e.g. 'page', 'complex',
	// 'chart_agent').
	PageBillingBreakdown map[string][]ParseUsagePageBillingBreakdown `json:"page_billing_breakdown" api:"nullable"`
	JSON                 parseUsageJSON                              `json:"-"`
}

// parseUsageJSON contains the JSON metadata for the struct [ParseUsage]
type parseUsageJSON struct {
	NumPages             apijson.Field
	CreditBreakdown      apijson.Field
	Credits              apijson.Field
	PageBillingBreakdown apijson.Field
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
)

func (r ParseUsagePageBillingBreakdown) IsKnown() bool {
	switch r {
	case ParseUsagePageBillingBreakdownPage, ParseUsagePageBillingBreakdownHTMLPage, ParseUsagePageBillingBreakdownDocxNativePage, ParseUsagePageBillingBreakdownAgentic, ParseUsagePageBillingBreakdownComplex, ParseUsagePageBillingBreakdownChartAgent, ParseUsagePageBillingBreakdownSpreadsheetCells, ParseUsagePageBillingBreakdownBillableSpreadsheetPages, ParseUsagePageBillingBreakdownEnrichTable, ParseUsagePageBillingBreakdownFigureSummary, ParseUsagePageBillingBreakdownTableSummary, ParseUsagePageBillingBreakdownKeyValue, ParseUsagePageBillingBreakdownAgenticText, ParseUsagePageBillingBreakdownPromptableAgenticText:
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

type SplitTableOptionsParam struct {
	// If True, a page can belong to multiple categories/partitions. If False, each
	// page must belong to exactly one category. Defaults to True.
	AllowPageOverlap param.Field[bool] `json:"allow_page_overlap"`
	// If tables should be truncated to the first few rows or if all content should be
	// preserved. truncate improves latency, preserve is recommended for cases where
	// partition_key is being used and the partition_key may be included within the
	// table. Defaults to truncate
	TableCutoff param.Field[SplitTableOptionsTableCutoff] `json:"table_cutoff"`
}

func (r SplitTableOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If tables should be truncated to the first few rows or if all content should be
// preserved. truncate improves latency, preserve is recommended for cases where
// partition_key is being used and the partition_key may be included within the
// table. Defaults to truncate
type SplitTableOptionsTableCutoff string

const (
	SplitTableOptionsTableCutoffTruncate SplitTableOptionsTableCutoff = "truncate"
	SplitTableOptionsTableCutoffPreserve SplitTableOptionsTableCutoff = "preserve"
)

func (r SplitTableOptionsTableCutoff) IsKnown() bool {
	switch r {
	case SplitTableOptionsTableCutoffTruncate, SplitTableOptionsTableCutoffPreserve:
		return true
	}
	return false
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
	Settings param.Field[SplitTableOptionsParam] `json:"settings"`
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
	Settings param.Field[SplitTableOptionsParam] `json:"settings"`
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
