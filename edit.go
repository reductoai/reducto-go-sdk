// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reducto

import (
	"context"
	"net/http"

	"github.com/reductoai/reducto-go-sdk/internal/apijson"
	"github.com/reductoai/reducto-go-sdk/internal/param"
	"github.com/reductoai/reducto-go-sdk/internal/requestconfig"
	"github.com/reductoai/reducto-go-sdk/option"
	"github.com/reductoai/reducto-go-sdk/shared"
)

// EditService contains methods and other services that help with interacting with
// the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEditService] method instead.
type EditService struct {
	Options []option.RequestOption
}

// NewEditService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEditService(opts ...option.RequestOption) (r *EditService) {
	r = &EditService{}
	r.Options = opts
	return
}

// Edit
func (r *EditService) Run(ctx context.Context, body EditRunParams, opts ...option.RequestOption) (res *EditRunResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "edit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Edit Async
func (r *EditService) RunJob(ctx context.Context, body EditRunJobParams, opts ...option.RequestOption) (res *EditRunJobResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "edit_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EditRunResponse struct {
	DocumentURL string              `json:"document_url,required"`
	JSON        editRunResponseJSON `json:"-"`
}

// editRunResponseJSON contains the JSON metadata for the struct [EditRunResponse]
type editRunResponseJSON struct {
	DocumentURL apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EditRunResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editRunResponseJSON) RawJSON() string {
	return r.raw
}

type EditRunJobResponse struct {
	JobID string                 `json:"job_id,required"`
	JSON  editRunJobResponseJSON `json:"-"`
}

// editRunJobResponseJSON contains the JSON metadata for the struct
// [EditRunJobResponse]
type editRunJobResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EditRunJobResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editRunJobResponseJSON) RawJSON() string {
	return r.raw
}

type EditRunParams struct {
	// The URL of the document to be processed. You can provide one of the following:
	//
	//  1. A publicly available URL
	//  2. A presigned S3 URL
	//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
	//     uploading a document
	DocumentURL param.Field[EditRunParamsDocumentURLUnion] `json:"document_url,required"`
	// The instructions for the edit.
	EditInstructions param.Field[string]                   `json:"edit_instructions,required"`
	EditOptions      param.Field[EditRunParamsEditOptions] `json:"edit_options"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool] `json:"priority"`
	// List of text snippets that can be reused throughout the document.
	Snippets param.Field[[]string] `json:"snippets"`
}

func (r EditRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The URL of the document to be processed. You can provide one of the following:
//
//  1. A publicly available URL
//  2. A presigned S3 URL
//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
//     uploading a document
//
// Satisfied by [shared.UnionString], [shared.UploadParam].
type EditRunParamsDocumentURLUnion interface {
	ImplementsEditRunParamsDocumentURLUnion()
}

type EditRunParamsEditOptions struct {
	// The color to use for edits, in hex format.
	Color param.Field[string] `json:"color"`
}

func (r EditRunParamsEditOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EditRunJobParams struct {
	// The URL of the document to be processed. You can provide one of the following:
	//
	//  1. A publicly available URL
	//  2. A presigned S3 URL
	//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
	//     uploading a document
	DocumentURL param.Field[EditRunJobParamsDocumentURLUnion] `json:"document_url,required"`
	// The instructions for the edit.
	EditInstructions param.Field[string]                      `json:"edit_instructions,required"`
	EditOptions      param.Field[EditRunJobParamsEditOptions] `json:"edit_options"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool] `json:"priority"`
	// List of text snippets that can be reused throughout the document.
	Snippets param.Field[[]string]                     `json:"snippets"`
	Webhook  param.Field[shared.WebhookConfigNewParam] `json:"webhook"`
}

func (r EditRunJobParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The URL of the document to be processed. You can provide one of the following:
//
//  1. A publicly available URL
//  2. A presigned S3 URL
//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
//     uploading a document
//
// Satisfied by [shared.UnionString], [shared.UploadParam].
type EditRunJobParamsDocumentURLUnion interface {
	ImplementsEditRunJobParamsDocumentURLUnion()
}

type EditRunJobParamsEditOptions struct {
	// The color to use for edits, in hex format.
	Color param.Field[string] `json:"color"`
}

func (r EditRunJobParamsEditOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
