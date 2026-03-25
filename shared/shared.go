// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/reductoai/reducto-go-sdk/internal/apijson"
	"github.com/reductoai/reducto-go-sdk/internal/param"
)

type Upload struct {
	FileID       string     `json:"file_id" api:"required"`
	PresignedURL string     `json:"presigned_url" api:"nullable"`
	JSON         uploadJSON `json:"-"`
}

// uploadJSON contains the JSON metadata for the struct [Upload]
type uploadJSON struct {
	FileID       apijson.Field
	PresignedURL apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Upload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uploadJSON) RawJSON() string {
	return r.raw
}

type UploadParam struct {
	FileID       param.Field[string] `json:"file_id" api:"required"`
	PresignedURL param.Field[string] `json:"presigned_url"`
}

func (r UploadParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UploadParam) ImplementsAsyncParseConfigInputUnionParam() {}

func (r UploadParam) ImplementsParseRunParamsBodySyncParseConfigInputUnion() {}

func (r UploadParam) ImplementsAsyncExtractConfigInputUnionParam() {}

func (r UploadParam) ImplementsExtractRunParamsBodySyncExtractConfigInputUnion() {}

func (r UploadParam) ImplementsSplitRunParamsInputUnion() {}

func (r UploadParam) ImplementsSplitRunJobParamsInputUnion() {}

func (r UploadParam) ImplementsEditRunParamsDocumentURLUnion() {}

func (r UploadParam) ImplementsEditRunJobParamsDocumentURLUnion() {}

func (r UploadParam) ImplementsPipelineRunParamsInputUnion() {}

func (r UploadParam) ImplementsPipelineRunJobParamsInputUnion() {}

func (r UploadParam) ImplementsClassifyRunParamsInputUnion() {}
