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
	JobID string `json:"job_id" api:"nullable"`
	// This field can have the runtime type of [[]interface{}].
	Result interface{} `json:"result"`
	// The link to the studio pipeline for the document.
	StudioLink string                 `json:"studio_link" api:"nullable"`
	Usage      shared.ExtractUsage    `json:"usage"`
	JSON       extractRunResponseJSON `json:"-"`
	union      ExtractRunResponseUnion
}

// extractRunResponseJSON contains the JSON metadata for the struct
// [ExtractRunResponse]
type extractRunResponseJSON struct {
	JobID       apijson.Field
	Result      apijson.Field
	StudioLink  apijson.Field
	Usage       apijson.Field
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
// Possible runtime types of the union are [shared.V3ExtractResponse],
// [ExtractRunResponseAsyncExtractResponse].
func (r ExtractRunResponse) AsUnion() ExtractRunResponseUnion {
	return r.union
}

// Union satisfied by [shared.V3ExtractResponse] or
// [ExtractRunResponseAsyncExtractResponse].
type ExtractRunResponseUnion interface {
	ImplementsExtractRunResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ExtractRunResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.V3ExtractResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExtractRunResponseAsyncExtractResponse{}),
		},
	)
}

type ExtractRunResponseAsyncExtractResponse struct {
	JobID string                                     `json:"job_id" api:"required"`
	JSON  extractRunResponseAsyncExtractResponseJSON `json:"-"`
}

// extractRunResponseAsyncExtractResponseJSON contains the JSON metadata for the
// struct [ExtractRunResponseAsyncExtractResponse]
type extractRunResponseAsyncExtractResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtractRunResponseAsyncExtractResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extractRunResponseAsyncExtractResponseJSON) RawJSON() string {
	return r.raw
}

func (r ExtractRunResponseAsyncExtractResponse) ImplementsExtractRunResponse() {}

type ExtractRunJobResponse struct {
	JobID string                    `json:"job_id" api:"required"`
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
	Body ExtractRunParamsBodyUnion `json:"body" api:"required"`
}

func (r ExtractRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ExtractRunParamsBody struct {
	Input param.Field[interface{}] `json:"input" api:"required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async        param.Field[shared.ConfigV3AsyncConfigParam] `json:"async"`
	Instructions param.Field[interface{}]                     `json:"instructions"`
	// The configuration options for parsing the document. If you are passing in a
	// jobid:// URL for the file, then this configuration will be ignored.
	Parsing  param.Field[shared.ParseOptionsParam] `json:"parsing"`
	Settings param.Field[interface{}]              `json:"settings"`
}

func (r ExtractRunParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBody) implementsExtractRunParamsBodyUnion() {}

// Satisfied by [ExtractRunParamsBodySyncExtractConfig],
// [ExtractRunParamsBodyAsyncExtractConfig], [ExtractRunParamsBody].
type ExtractRunParamsBodyUnion interface {
	implementsExtractRunParamsBodyUnion()
}

type ExtractRunParamsBodySyncExtractConfig struct {
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input param.Field[ExtractRunParamsBodySyncExtractConfigInputUnion] `json:"input" api:"required"`
	// The instructions to use for the extraction.
	Instructions param.Field[ExtractRunParamsBodySyncExtractConfigInstructions] `json:"instructions"`
	// The configuration options for parsing the document. If you are passing in a
	// jobid:// URL for the file, then this configuration will be ignored.
	Parsing param.Field[shared.ParseOptionsParam] `json:"parsing"`
	// The settings to use for the extraction.
	Settings param.Field[ExtractRunParamsBodySyncExtractConfigSettings] `json:"settings"`
}

func (r ExtractRunParamsBodySyncExtractConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBodySyncExtractConfig) implementsExtractRunParamsBodyUnion() {}

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
// [ExtractRunParamsBodySyncExtractConfigInputArray], [shared.UploadParam].
type ExtractRunParamsBodySyncExtractConfigInputUnion interface {
	ImplementsExtractRunParamsBodySyncExtractConfigInputUnion()
}

type ExtractRunParamsBodySyncExtractConfigInputArray []string

func (r ExtractRunParamsBodySyncExtractConfigInputArray) ImplementsExtractRunParamsBodySyncExtractConfigInputUnion() {
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

type ExtractRunParamsBodyAsyncExtractConfig struct {
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input param.Field[ExtractRunParamsBodyAsyncExtractConfigInputUnion] `json:"input" api:"required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async param.Field[shared.ConfigV3AsyncConfigParam] `json:"async"`
	// The instructions to use for the extraction.
	Instructions param.Field[ExtractRunParamsBodyAsyncExtractConfigInstructions] `json:"instructions"`
	// The configuration options for parsing the document. If you are passing in a
	// jobid:// URL for the file, then this configuration will be ignored.
	Parsing param.Field[shared.ParseOptionsParam] `json:"parsing"`
	// The settings to use for the extraction.
	Settings param.Field[ExtractRunParamsBodyAsyncExtractConfigSettings] `json:"settings"`
}

func (r ExtractRunParamsBodyAsyncExtractConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExtractRunParamsBodyAsyncExtractConfig) implementsExtractRunParamsBodyUnion() {}

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
// [ExtractRunParamsBodyAsyncExtractConfigInputArray], [shared.UploadParam].
type ExtractRunParamsBodyAsyncExtractConfigInputUnion interface {
	ImplementsExtractRunParamsBodyAsyncExtractConfigInputUnion()
}

type ExtractRunParamsBodyAsyncExtractConfigInputArray []string

func (r ExtractRunParamsBodyAsyncExtractConfigInputArray) ImplementsExtractRunParamsBodyAsyncExtractConfigInputUnion() {
}

// The instructions to use for the extraction.
type ExtractRunParamsBodyAsyncExtractConfigInstructions struct {
	// The JSON schema to use for the extraction.
	Schema param.Field[interface{}] `json:"schema"`
	// The system prompt to use for the extraction.
	SystemPrompt param.Field[string] `json:"system_prompt"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigInstructions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The settings to use for the extraction.
type ExtractRunParamsBodyAsyncExtractConfigSettings struct {
	// If True, use array extraction.
	ArrayExtract param.Field[bool] `json:"array_extract"`
	// The citations to use for the extraction.
	Citations param.Field[ExtractRunParamsBodyAsyncExtractConfigSettingsCitations] `json:"citations"`
	// If True, include images in the extraction.
	IncludeImages param.Field[bool] `json:"include_images"`
	// If True, jobs will be processed with a higher throughput and priority at a
	// higher cost. Defaults to False.
	OptimizeForLatency param.Field[bool] `json:"optimize_for_latency"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The citations to use for the extraction.
type ExtractRunParamsBodyAsyncExtractConfigSettingsCitations struct {
	// If True, include citations in the extraction.
	Enabled param.Field[bool] `json:"enabled"`
	// If True, enable numeric citation confidence scores. Defaults to True.
	NumericalConfidence param.Field[bool] `json:"numerical_confidence"`
}

func (r ExtractRunParamsBodyAsyncExtractConfigSettingsCitations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExtractRunJobParams struct {
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input param.Field[ExtractRunJobParamsInputUnion] `json:"input" api:"required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async param.Field[shared.ConfigV3AsyncConfigParam] `json:"async"`
	// The instructions to use for the extraction.
	Instructions param.Field[ExtractRunJobParamsInstructions] `json:"instructions"`
	// The configuration options for parsing the document. If you are passing in a
	// jobid:// URL for the file, then this configuration will be ignored.
	Parsing param.Field[shared.ParseOptionsParam] `json:"parsing"`
	// The settings to use for the extraction.
	Settings param.Field[ExtractRunJobParamsSettings] `json:"settings"`
}

func (r ExtractRunJobParams) MarshalJSON() (data []byte, err error) {
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
// Satisfied by [shared.UnionString], [ExtractRunJobParamsInputArray],
// [shared.UploadParam].
type ExtractRunJobParamsInputUnion interface {
	ImplementsExtractRunJobParamsInputUnion()
}

type ExtractRunJobParamsInputArray []string

func (r ExtractRunJobParamsInputArray) ImplementsExtractRunJobParamsInputUnion() {}

// The instructions to use for the extraction.
type ExtractRunJobParamsInstructions struct {
	// The JSON schema to use for the extraction.
	Schema param.Field[interface{}] `json:"schema"`
	// The system prompt to use for the extraction.
	SystemPrompt param.Field[string] `json:"system_prompt"`
}

func (r ExtractRunJobParamsInstructions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The settings to use for the extraction.
type ExtractRunJobParamsSettings struct {
	// If True, use array extraction.
	ArrayExtract param.Field[bool] `json:"array_extract"`
	// The citations to use for the extraction.
	Citations param.Field[ExtractRunJobParamsSettingsCitations] `json:"citations"`
	// If True, include images in the extraction.
	IncludeImages param.Field[bool] `json:"include_images"`
	// If True, jobs will be processed with a higher throughput and priority at a
	// higher cost. Defaults to False.
	OptimizeForLatency param.Field[bool] `json:"optimize_for_latency"`
}

func (r ExtractRunJobParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The citations to use for the extraction.
type ExtractRunJobParamsSettingsCitations struct {
	// If True, include citations in the extraction.
	Enabled param.Field[bool] `json:"enabled"`
	// If True, enable numeric citation confidence scores. Defaults to True.
	NumericalConfidence param.Field[bool] `json:"numerical_confidence"`
}

func (r ExtractRunJobParamsSettingsCitations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
