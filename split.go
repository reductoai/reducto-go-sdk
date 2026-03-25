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
func (r *SplitService) Run(ctx context.Context, body SplitRunParams, opts ...option.RequestOption) (res *SplitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "split"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Split Async
func (r *SplitService) RunJob(ctx context.Context, body SplitRunJobParams, opts ...option.RequestOption) (res *SplitRunJobResponse, err error) {
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
	JSON            parseUsageJSON     `json:"-"`
}

// parseUsageJSON contains the JSON metadata for the struct [ParseUsage]
type parseUsageJSON struct {
	NumPages        apijson.Field
	CreditBreakdown apijson.Field
	Credits         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ParseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseUsageJSON) RawJSON() string {
	return r.raw
}

type SplitCategoryParam struct {
	Description  param.Field[string] `json:"description" api:"required"`
	Name         param.Field[string] `json:"name" api:"required"`
	PartitionKey param.Field[string] `json:"partition_key"`
}

func (r SplitCategoryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SplitResponse struct {
	// The split result.
	Result SplitResponseResult `json:"result" api:"required"`
	Usage  ParseUsage          `json:"usage" api:"required"`
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

func (r SplitResponse) implementsJobGetResponseAsyncJobResponseResult() {}

func (r SplitResponse) implementsJobGetResponseEnhancedAsyncJobResponseResult() {}

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
	Pages      []DeepSplitPageEvidence                             `json:"pages" api:"required"`
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
	Pages []DeepSplitPageEvidence                               `json:"pages" api:"required"`
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

type SplitTableOptionsParam struct {
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

type SplitRunJobResponse struct {
	JobID string                  `json:"job_id" api:"required"`
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
