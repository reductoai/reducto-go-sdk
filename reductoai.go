// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reductoai

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/url"

	"github.com/stainless-sdks/reductoai-go/internal/apiform"
	"github.com/stainless-sdks/reductoai-go/internal/apijson"
	"github.com/stainless-sdks/reductoai-go/internal/apiquery"
	"github.com/stainless-sdks/reductoai-go/internal/param"
)

type APIVersionResponse = interface{}

type UploadResponse struct {
	FileID       string             `json:"file_id,required"`
	PresignedURL string             `json:"presigned_url,nullable"`
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

type UploadParams struct {
	Extension param.Field[string]    `query:"extension"`
	File      param.Field[io.Reader] `json:"file" format:"binary"`
}

func (r UploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes [UploadParams]'s query parameters as `url.Values`.
func (r UploadParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
