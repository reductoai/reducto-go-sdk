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
func (r *ParseService) RunJob(ctx context.Context, body ParseRunJobParams, opts ...option.RequestOption) (res *ParseRunJobResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "parse_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ParseRunResponse struct {
	JobID string `json:"job_id" api:"required"`
	// The duration of the parse request in seconds.
	Duration float64 `json:"duration"`
	// The storage URL of the converted PDF file.
	PdfURL string `json:"pdf_url" api:"nullable"`
	// This field can have the runtime type of [shared.ParseResponseResult].
	Result interface{} `json:"result"`
	// The link to the studio pipeline for the document.
	StudioLink string               `json:"studio_link" api:"nullable"`
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
	JobID string                                 `json:"job_id" api:"required"`
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
	JobID string                  `json:"job_id" api:"required"`
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
	Body ParseRunParamsBodyUnion `json:"body" api:"required"`
}

func (r ParseRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ParseRunParamsBody struct {
	Input param.Field[interface{}] `json:"input" api:"required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async       param.Field[shared.ConfigV3AsyncConfigParam] `json:"async"`
	Enhance     param.Field[shared.EnhanceParam]             `json:"enhance"`
	Formatting  param.Field[shared.FormattingParam]          `json:"formatting"`
	Retrieval   param.Field[shared.RetrievalParam]           `json:"retrieval"`
	Settings    param.Field[shared.SettingsParam]            `json:"settings"`
	Spreadsheet param.Field[shared.SpreadsheetParam]         `json:"spreadsheet"`
}

func (r ParseRunParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBody) implementsParseRunParamsBodyUnion() {}

// Satisfied by [ParseRunParamsBodySyncParseConfig],
// [ParseRunParamsBodyAsyncParseConfig], [ParseRunParamsBody].
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
	Enhance     param.Field[shared.EnhanceParam]                         `json:"enhance"`
	Formatting  param.Field[shared.FormattingParam]                      `json:"formatting"`
	Retrieval   param.Field[shared.RetrievalParam]                       `json:"retrieval"`
	Settings    param.Field[shared.SettingsParam]                        `json:"settings"`
	Spreadsheet param.Field[shared.SpreadsheetParam]                     `json:"spreadsheet"`
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

type ParseRunParamsBodyAsyncParseConfig struct {
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input param.Field[ParseRunParamsBodyAsyncParseConfigInputUnion] `json:"input" api:"required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async       param.Field[shared.ConfigV3AsyncConfigParam] `json:"async"`
	Enhance     param.Field[shared.EnhanceParam]             `json:"enhance"`
	Formatting  param.Field[shared.FormattingParam]          `json:"formatting"`
	Retrieval   param.Field[shared.RetrievalParam]           `json:"retrieval"`
	Settings    param.Field[shared.SettingsParam]            `json:"settings"`
	Spreadsheet param.Field[shared.SpreadsheetParam]         `json:"spreadsheet"`
}

func (r ParseRunParamsBodyAsyncParseConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseRunParamsBodyAsyncParseConfig) implementsParseRunParamsBodyUnion() {}

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
// [ParseRunParamsBodyAsyncParseConfigInputArray], [shared.UploadParam].
type ParseRunParamsBodyAsyncParseConfigInputUnion interface {
	ImplementsParseRunParamsBodyAsyncParseConfigInputUnion()
}

type ParseRunParamsBodyAsyncParseConfigInputArray []string

func (r ParseRunParamsBodyAsyncParseConfigInputArray) ImplementsParseRunParamsBodyAsyncParseConfigInputUnion() {
}

type ParseRunJobParams struct {
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input param.Field[ParseRunJobParamsInputUnion] `json:"input" api:"required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async       param.Field[shared.ConfigV3AsyncConfigParam] `json:"async"`
	Enhance     param.Field[shared.EnhanceParam]             `json:"enhance"`
	Formatting  param.Field[shared.FormattingParam]          `json:"formatting"`
	Retrieval   param.Field[shared.RetrievalParam]           `json:"retrieval"`
	Settings    param.Field[shared.SettingsParam]            `json:"settings"`
	Spreadsheet param.Field[shared.SpreadsheetParam]         `json:"spreadsheet"`
}

func (r ParseRunJobParams) MarshalJSON() (data []byte, err error) {
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
// Satisfied by [shared.UnionString], [ParseRunJobParamsInputArray],
// [shared.UploadParam].
type ParseRunJobParamsInputUnion interface {
	ImplementsParseRunJobParamsInputUnion()
}

type ParseRunJobParamsInputArray []string

func (r ParseRunJobParamsInputArray) ImplementsParseRunJobParamsInputUnion() {}
