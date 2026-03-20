// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reducto

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/reductoai/reducto-go-sdk/internal/apijson"
	"github.com/reductoai/reducto-go-sdk/internal/apiquery"
	"github.com/reductoai/reducto-go-sdk/internal/param"
	"github.com/reductoai/reducto-go-sdk/internal/requestconfig"
	"github.com/reductoai/reducto-go-sdk/option"
)

// UploadService contains methods and other services that help with interacting
// with the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUploadService] method instead.
type UploadService struct {
	Options []option.RequestOption
}

// NewUploadService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUploadService(opts ...option.RequestOption) (r *UploadService) {
	r = &UploadService{}
	r.Options = opts
	return
}

// Upload
func (r *UploadService) New(ctx context.Context, params UploadNewParams, opts ...option.RequestOption) (res *UploadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "upload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type UploadResponse struct {
	FileID       string             `json:"file_id" api:"required"`
	PresignedURL string             `json:"presigned_url" api:"nullable"`
	JSON         uploadResponseJSON `json:"-"`
}

// uploadResponseJSON contains the JSON metadata for the struct [UploadResponse]
type uploadResponseJSON struct {
	FileID       apijson.Field
	PresignedURL apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uploadResponseJSON) RawJSON() string {
	return r.raw
}

type UploadResponseParam struct {
	FileID       param.Field[string] `json:"file_id" api:"required"`
	PresignedURL param.Field[string] `json:"presigned_url"`
}

func (r UploadResponseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UploadResponseParam) ImplementsAsyncParseConfigInputUnionParam() {}

func (r UploadResponseParam) ImplementsParseNewParamsBodySyncParseConfigInputUnion() {}

func (r UploadResponseParam) ImplementsAsyncExtractConfigInputUnionParam() {}

func (r UploadResponseParam) ImplementsExtractNewParamsBodySyncExtractConfigInputUnion() {}

func (r UploadResponseParam) ImplementsSplitNewParamsInputUnion() {}

func (r UploadResponseParam) ImplementsSplitAsyncNewParamsInputUnion() {}

func (r UploadResponseParam) ImplementsEditSubmitParamsDocumentURLUnion() {}

func (r UploadResponseParam) ImplementsEditAsyncNewParamsDocumentURLUnion() {}

func (r UploadResponseParam) ImplementsPipelineNewParamsInputUnion() {}

func (r UploadResponseParam) ImplementsPipelineAsyncNewParamsInputUnion() {}

func (r UploadResponseParam) ImplementsClassifyClassifyParamsInputUnion() {}

type UploadNewParams struct {
	Extension param.Field[string] `query:"extension"`
	File      param.Field[string] `json:"file"`
}

func (r UploadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [UploadNewParams]'s query parameters as `url.Values`.
func (r UploadNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
