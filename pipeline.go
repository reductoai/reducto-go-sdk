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

// PipelineService contains methods and other services that help with interacting
// with the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPipelineService] method instead.
type PipelineService struct {
	Options []option.RequestOption
}

// NewPipelineService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPipelineService(opts ...option.RequestOption) (r *PipelineService) {
	r = &PipelineService{}
	r.Options = opts
	return
}

// Pipeline
func (r *PipelineService) New(ctx context.Context, body PipelineNewParams, opts ...option.RequestOption) (res *PipelineResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pipeline"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PipelineResponse struct {
	JobID  string                 `json:"job_id" api:"required"`
	Result PipelineResponseResult `json:"result" api:"required"`
	Usage  ParseUsage             `json:"usage" api:"required"`
	JSON   pipelineResponseJSON   `json:"-"`
}

// pipelineResponseJSON contains the JSON metadata for the struct
// [PipelineResponse]
type pipelineResponseJSON struct {
	JobID       apijson.Field
	Result      apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineResponseJSON) RawJSON() string {
	return r.raw
}

func (r PipelineResponse) implementsJobGetResponseAsyncJobResponseResult() {}

func (r PipelineResponse) implementsJobGetResponseEnhancedAsyncJobResponseResult() {}

type PipelineResponseResult struct {
	Extract PipelineResponseResultExtractUnion `json:"extract" api:"required,nullable"`
	Parse   PipelineResponseResultParseUnion   `json:"parse" api:"required,nullable"`
	Split   SplitResponse                      `json:"split" api:"required,nullable"`
	Edit    EditResponse                       `json:"edit" api:"nullable"`
	JSON    pipelineResponseResultJSON         `json:"-"`
}

// pipelineResponseResultJSON contains the JSON metadata for the struct
// [PipelineResponseResult]
type pipelineResponseResultJSON struct {
	Extract     apijson.Field
	Parse       apijson.Field
	Split       apijson.Field
	Edit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineResponseResultJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [PipelineResponseResultExtractArray], [ExtractResponse] or
// [V3Extract].
type PipelineResponseResultExtractUnion interface {
	implementsPipelineResponseResultExtractUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PipelineResponseResultExtractUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineResponseResultExtractArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExtractResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V3Extract{}),
		},
	)
}

type PipelineResponseResultExtractArray []PipelineResponseResultExtractArrayItem

func (r PipelineResponseResultExtractArray) implementsPipelineResponseResultExtractUnion() {}

// This is the response format for Extract -> Split Pipelines
type PipelineResponseResultExtractArrayItem struct {
	PageRange []int64                                    `json:"page_range" api:"required"`
	Result    PipelineResponseResultExtractArrayResult   `json:"result" api:"required"`
	SplitName string                                     `json:"split_name" api:"required"`
	Partition string                                     `json:"partition" api:"nullable"`
	JSON      pipelineResponseResultExtractArrayItemJSON `json:"-"`
}

// pipelineResponseResultExtractArrayItemJSON contains the JSON metadata for the
// struct [PipelineResponseResultExtractArrayItem]
type pipelineResponseResultExtractArrayItemJSON struct {
	PageRange   apijson.Field
	Result      apijson.Field
	SplitName   apijson.Field
	Partition   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineResponseResultExtractArrayItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineResponseResultExtractArrayItemJSON) RawJSON() string {
	return r.raw
}

type PipelineResponseResultExtractArrayResult struct {
	// This field can have the runtime type of [[]interface{}].
	Result interface{}  `json:"result" api:"required"`
	Usage  ExtractUsage `json:"usage" api:"required"`
	// This field can have the runtime type of [[]interface{}].
	Citations interface{} `json:"citations"`
	JobID     string      `json:"job_id" api:"nullable"`
	// The link to the studio pipeline for the document.
	StudioLink string                                       `json:"studio_link" api:"nullable"`
	JSON       pipelineResponseResultExtractArrayResultJSON `json:"-"`
	union      PipelineResponseResultExtractArrayResultUnion
}

// pipelineResponseResultExtractArrayResultJSON contains the JSON metadata for the
// struct [PipelineResponseResultExtractArrayResult]
type pipelineResponseResultExtractArrayResultJSON struct {
	Result      apijson.Field
	Usage       apijson.Field
	Citations   apijson.Field
	JobID       apijson.Field
	StudioLink  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r pipelineResponseResultExtractArrayResultJSON) RawJSON() string {
	return r.raw
}

func (r *PipelineResponseResultExtractArrayResult) UnmarshalJSON(data []byte) (err error) {
	*r = PipelineResponseResultExtractArrayResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PipelineResponseResultExtractArrayResultUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [ExtractResponse], [V3Extract].
func (r PipelineResponseResultExtractArrayResult) AsUnion() PipelineResponseResultExtractArrayResultUnion {
	return r.union
}

// Union satisfied by [ExtractResponse] or [V3Extract].
type PipelineResponseResultExtractArrayResultUnion interface {
	implementsPipelineResponseResultExtractArrayResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PipelineResponseResultExtractArrayResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExtractResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V3Extract{}),
		},
	)
}

// Union satisfied by [ParseResponse] or [PipelineResponseResultParseArray].
type PipelineResponseResultParseUnion interface {
	implementsPipelineResponseResultParseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PipelineResponseResultParseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParseResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineResponseResultParseArray{}),
		},
	)
}

type PipelineResponseResultParseArray []ParseResponse

func (r PipelineResponseResultParseArray) implementsPipelineResponseResultParseUnion() {}

// Settings for pipeline execution that override pipeline defaults.
type PipelineSettingsParam struct {
	// Password to decrypt password-protected documents.
	DocumentPassword param.Field[string] `json:"document_password"`
}

func (r PipelineSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineNewParams struct {
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input param.Field[PipelineNewParamsInputUnion] `json:"input" api:"required"`
	// The ID of the pipeline to use for the document.
	PipelineID param.Field[string] `json:"pipeline_id" api:"required"`
	// Settings for pipeline execution that override pipeline defaults.
	Settings param.Field[PipelineSettingsParam] `json:"settings"`
}

func (r PipelineNewParams) MarshalJSON() (data []byte, err error) {
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
// Satisfied by [shared.UnionString], [PipelineNewParamsInputArray],
// [UploadResponseParam].
type PipelineNewParamsInputUnion interface {
	ImplementsPipelineNewParamsInputUnion()
}

type PipelineNewParamsInputArray []string

func (r PipelineNewParamsInputArray) ImplementsPipelineNewParamsInputUnion() {}
