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
)

// SplitAsyncService contains methods and other services that help with interacting
// with the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSplitAsyncService] method instead.
type SplitAsyncService struct {
	Options []option.RequestOption
}

// NewSplitAsyncService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSplitAsyncService(opts ...option.RequestOption) (r *SplitAsyncService) {
	r = &SplitAsyncService{}
	r.Options = opts
	return
}

// Split Async
func (r *SplitAsyncService) New(ctx context.Context, body SplitAsyncNewParams, opts ...option.RequestOption) (res *SplitAsyncNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "split_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SplitAsyncNewResponse struct {
	JobID string                    `json:"job_id" api:"required"`
	JSON  splitAsyncNewResponseJSON `json:"-"`
}

// splitAsyncNewResponseJSON contains the JSON metadata for the struct
// [SplitAsyncNewResponse]
type splitAsyncNewResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitAsyncNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitAsyncNewResponseJSON) RawJSON() string {
	return r.raw
}

type SplitAsyncNewParams struct {
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input param.Field[SplitAsyncNewParamsInputUnion] `json:"input" api:"required"`
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

func (r SplitAsyncNewParams) MarshalJSON() (data []byte, err error) {
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
// Satisfied by [shared.UnionString], [SplitAsyncNewParamsInputArray],
// [UploadResponseParam].
type SplitAsyncNewParamsInputUnion interface {
	ImplementsSplitAsyncNewParamsInputUnion()
}

type SplitAsyncNewParamsInputArray []string

func (r SplitAsyncNewParamsInputArray) ImplementsSplitAsyncNewParamsInputUnion() {}
