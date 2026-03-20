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

// ClassifyService contains methods and other services that help with interacting
// with the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClassifyService] method instead.
type ClassifyService struct {
	Options []option.RequestOption
}

// NewClassifyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClassifyService(opts ...option.RequestOption) (r *ClassifyService) {
	r = &ClassifyService{}
	r.Options = opts
	return
}

// Classify
func (r *ClassifyService) New(ctx context.Context, body ClassifyNewParams, opts ...option.RequestOption) (res *ClassifyNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "classify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Response from classify job - returned when polling /job/{job_id}
type ClassifyNewResponse struct {
	JobID  string                    `json:"job_id" api:"required"`
	Result ClassifyNewResponseResult `json:"result" api:"required"`
	// The duration of the classify request in seconds.
	Duration float64 `json:"duration" api:"nullable"`
	// Overall confidence breakdown for classification response.
	ResponseConfidence ClassifyNewResponseResponseConfidence `json:"response_confidence" api:"nullable"`
	JSON               classifyNewResponseJSON               `json:"-"`
}

// classifyNewResponseJSON contains the JSON metadata for the struct
// [ClassifyNewResponse]
type classifyNewResponseJSON struct {
	JobID              apijson.Field
	Result             apijson.Field
	Duration           apijson.Field
	ResponseConfidence apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ClassifyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r classifyNewResponseJSON) RawJSON() string {
	return r.raw
}

type ClassifyNewResponseResult struct {
	Category string                        `json:"category" api:"required"`
	JSON     classifyNewResponseResultJSON `json:"-"`
}

// classifyNewResponseResultJSON contains the JSON metadata for the struct
// [ClassifyNewResponseResult]
type classifyNewResponseResultJSON struct {
	Category    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClassifyNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r classifyNewResponseResultJSON) RawJSON() string {
	return r.raw
}

// Overall confidence breakdown for classification response.
type ClassifyNewResponseResponseConfidence struct {
	Categories []ClassifyNewResponseResponseConfidenceCategory `json:"categories" api:"required"`
	JSON       classifyNewResponseResponseConfidenceJSON       `json:"-"`
}

// classifyNewResponseResponseConfidenceJSON contains the JSON metadata for the
// struct [ClassifyNewResponseResponseConfidence]
type classifyNewResponseResponseConfidenceJSON struct {
	Categories  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClassifyNewResponseResponseConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r classifyNewResponseResponseConfidenceJSON) RawJSON() string {
	return r.raw
}

// Confidence result for a category.
type ClassifyNewResponseResponseConfidenceCategory struct {
	Category           string                                                              `json:"category" api:"required"`
	Confidence         float64                                                             `json:"confidence" api:"required"`
	CriteriaConfidence []ClassifyNewResponseResponseConfidenceCategoriesCriteriaConfidence `json:"criteria_confidence" api:"required"`
	JSON               classifyNewResponseResponseConfidenceCategoryJSON                   `json:"-"`
}

// classifyNewResponseResponseConfidenceCategoryJSON contains the JSON metadata for
// the struct [ClassifyNewResponseResponseConfidenceCategory]
type classifyNewResponseResponseConfidenceCategoryJSON struct {
	Category           apijson.Field
	Confidence         apijson.Field
	CriteriaConfidence apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ClassifyNewResponseResponseConfidenceCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r classifyNewResponseResponseConfidenceCategoryJSON) RawJSON() string {
	return r.raw
}

// Confidence result for a single criterion.
type ClassifyNewResponseResponseConfidenceCategoriesCriteriaConfidence struct {
	Confidence ClassifyNewResponseResponseConfidenceCategoriesCriteriaConfidenceConfidence `json:"confidence" api:"required"`
	Criterion  string                                                                      `json:"criterion" api:"required"`
	JSON       classifyNewResponseResponseConfidenceCategoriesCriteriaConfidenceJSON       `json:"-"`
}

// classifyNewResponseResponseConfidenceCategoriesCriteriaConfidenceJSON contains
// the JSON metadata for the struct
// [ClassifyNewResponseResponseConfidenceCategoriesCriteriaConfidence]
type classifyNewResponseResponseConfidenceCategoriesCriteriaConfidenceJSON struct {
	Confidence  apijson.Field
	Criterion   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClassifyNewResponseResponseConfidenceCategoriesCriteriaConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r classifyNewResponseResponseConfidenceCategoriesCriteriaConfidenceJSON) RawJSON() string {
	return r.raw
}

type ClassifyNewResponseResponseConfidenceCategoriesCriteriaConfidenceConfidence string

const (
	ClassifyNewResponseResponseConfidenceCategoriesCriteriaConfidenceConfidenceHigh ClassifyNewResponseResponseConfidenceCategoriesCriteriaConfidenceConfidence = "high"
	ClassifyNewResponseResponseConfidenceCategoriesCriteriaConfidenceConfidenceLow  ClassifyNewResponseResponseConfidenceCategoriesCriteriaConfidenceConfidence = "low"
)

func (r ClassifyNewResponseResponseConfidenceCategoriesCriteriaConfidenceConfidence) IsKnown() bool {
	switch r {
	case ClassifyNewResponseResponseConfidenceCategoriesCriteriaConfidenceConfidenceHigh, ClassifyNewResponseResponseConfidenceCategoriesCriteriaConfidenceConfidenceLow:
		return true
	}
	return false
}

type ClassifyNewParams struct {
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input param.Field[ClassifyNewParamsInputUnion] `json:"input" api:"required"`
	// A list of classification categories and their matching criteria.
	ClassificationSchema param.Field[[]ClassifyNewParamsClassificationSchema] `json:"classification_schema"`
	// Optional document-level metadata to include in classification prompts.
	DocumentMetadata param.Field[string] `json:"document_metadata"`
	// The page range to process (1-indexed). By default, the first 5 pages are used.
	// If more than 25 pages are selected, only the first 25 (after sorting) are used.
	// Only applies to PDFs; ignored for other document types.
	PageRange param.Field[ClassifyNewParamsPageRangeUnion] `json:"page_range"`
	// If True, persist the results indefinitely. Defaults to False.
	PersistResults param.Field[bool] `json:"persist_results"`
}

func (r ClassifyNewParams) MarshalJSON() (data []byte, err error) {
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
// Satisfied by [shared.UnionString], [ClassifyNewParamsInputArray],
// [shared.UploadParam].
type ClassifyNewParamsInputUnion interface {
	ImplementsClassifyNewParamsInputUnion()
}

type ClassifyNewParamsInputArray []string

func (r ClassifyNewParamsInputArray) ImplementsClassifyNewParamsInputUnion() {}

// A single classification category with its matching criteria.
type ClassifyNewParamsClassificationSchema struct {
	// The category name/label that documents will be classified into (e.g., 'invoice',
	// 'contract', 'receipt').
	Category param.Field[string] `json:"category" api:"required"`
	// A list of criteria, keywords, or descriptions that define what characteristics a
	// document must have to be classified into this category (e.g., ['contains billing
	// information', 'has itemized charges']).
	Criteria param.Field[[]string] `json:"criteria" api:"required"`
}

func (r ClassifyNewParamsClassificationSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The page range to process (1-indexed). By default, the first 5 pages are used.
// If more than 25 pages are selected, only the first 25 (after sorting) are used.
// Only applies to PDFs; ignored for other document types.
//
// Satisfied by [shared.PageRangeParam], [ClassifyNewParamsPageRangeArray],
// [ClassifyNewParamsPageRangeArray].
type ClassifyNewParamsPageRangeUnion interface {
	ImplementsClassifyNewParamsPageRangeUnion()
}

type ClassifyNewParamsPageRangeArray []shared.PageRangeParam

func (r ClassifyNewParamsPageRangeArray) ImplementsClassifyNewParamsPageRangeUnion() {}
