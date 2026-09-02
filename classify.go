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
func (r *ClassifyService) Run(ctx context.Context, body ClassifyRunParams, opts ...option.RequestOption) (res *shared.ClassifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "classify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ClassifyRunParams struct {
	// The URL of the document to be classified. You can provide one of the following:
	//
	//  1. A publicly available URL
	//  2. A presigned S3 URL
	//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
	//     uploading a document
	Input param.Field[ClassifyRunParamsInputUnion] `json:"input" api:"required"`
	// A mapping of higher-level classify groups to the category labels that belong to
	// each group. When provided, the response includes `extra_metadata.grouping` with
	// the matched group name, or `ungrouped` if the selected category is not in any
	// group.
	CategoryGroups param.Field[map[string][]string] `json:"category_groups"`
	// A list of classification categories and their matching criteria.
	ClassificationSchema param.Field[[]ClassifyRunParamsClassificationSchema] `json:"classification_schema"`
	// Optional document-level metadata to include in classification prompts.
	DocumentMetadata param.Field[string] `json:"document_metadata"`
	// Force the endpoint result to be returned in URL form.
	ForceURLResult param.Field[bool] `json:"force_url_result"`
	// The classification model to use. Set to "accurate" to run Deep Classify for
	// higher accuracy on hard documents. Defaults to "default".
	Model param.Field[ClassifyRunParamsModel] `json:"model"`
	// The page range to process (1-indexed). By default, the first 5 pages are used.
	// At most 10 pages can be selected. Only applies to PDFs; ignored for other
	// document types.
	PageRange param.Field[ClassifyRunParamsPageRangeUnion] `json:"page_range"`
	// Workers poll the priority queue ahead of the standard queue, so priority jobs
	// start sooner when there is queued work; sync jobs are prioritized above async
	// jobs by default.
	Priority param.Field[bool] `json:"priority"`
}

func (r ClassifyRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The URL of the document to be classified. You can provide one of the following:
//
//  1. A publicly available URL
//  2. A presigned S3 URL
//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
//     uploading a document
//
// Satisfied by [shared.UnionString], [ClassifyRunParamsInputArray],
// [shared.UploadParam].
type ClassifyRunParamsInputUnion interface {
	ImplementsClassifyRunParamsInputUnion()
}

type ClassifyRunParamsInputArray []string

func (r ClassifyRunParamsInputArray) ImplementsClassifyRunParamsInputUnion() {}

// A single classification category with its matching criteria.
type ClassifyRunParamsClassificationSchema struct {
	// The category name/label that documents will be classified into (e.g., 'invoice',
	// 'contract', 'receipt').
	Category param.Field[string] `json:"category" api:"required"`
	// A list of criteria, keywords, or descriptions that define what characteristics a
	// document must have to be classified into this category (e.g., ['contains billing
	// information', 'has itemized charges']).
	Criteria param.Field[[]string] `json:"criteria" api:"required"`
}

func (r ClassifyRunParamsClassificationSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The classification model to use. Set to "accurate" to run Deep Classify for
// higher accuracy on hard documents. Defaults to "default".
type ClassifyRunParamsModel string

const (
	ClassifyRunParamsModelDefault  ClassifyRunParamsModel = "default"
	ClassifyRunParamsModelAccurate ClassifyRunParamsModel = "accurate"
)

func (r ClassifyRunParamsModel) IsKnown() bool {
	switch r {
	case ClassifyRunParamsModelDefault, ClassifyRunParamsModelAccurate:
		return true
	}
	return false
}

// The page range to process (1-indexed). By default, the first 5 pages are used.
// At most 10 pages can be selected. Only applies to PDFs; ignored for other
// document types.
//
// Satisfied by [shared.PageRangeParam], [ClassifyRunParamsPageRangeArray],
// [ClassifyRunParamsPageRangeArray].
type ClassifyRunParamsPageRangeUnion interface {
	ImplementsClassifyRunParamsPageRangeUnion()
}

type ClassifyRunParamsPageRangeArray []shared.PageRangeParam

func (r ClassifyRunParamsPageRangeArray) ImplementsClassifyRunParamsPageRangeUnion() {}
