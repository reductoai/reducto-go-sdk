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

// EditAsyncService contains methods and other services that help with interacting
// with the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEditAsyncService] method instead.
type EditAsyncService struct {
	Options []option.RequestOption
}

// NewEditAsyncService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEditAsyncService(opts ...option.RequestOption) (r *EditAsyncService) {
	r = &EditAsyncService{}
	r.Options = opts
	return
}

// Edit Async
func (r *EditAsyncService) New(ctx context.Context, body EditAsyncNewParams, opts ...option.RequestOption) (res *EditAsyncNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "edit_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type EditAsyncNewResponse struct {
	JobID string                   `json:"job_id" api:"required"`
	JSON  editAsyncNewResponseJSON `json:"-"`
}

// editAsyncNewResponseJSON contains the JSON metadata for the struct
// [EditAsyncNewResponse]
type editAsyncNewResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EditAsyncNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editAsyncNewResponseJSON) RawJSON() string {
	return r.raw
}

type EditAsyncNewParams struct {
	// The URL of the document to be processed. You can provide one of the following:
	//
	//  1. A publicly available URL
	//  2. A presigned S3 URL
	//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
	//     uploading a document
	DocumentURL param.Field[EditAsyncNewParamsDocumentURLUnion] `json:"document_url" api:"required"`
	// The instructions for the edit.
	EditInstructions param.Field[string]           `json:"edit_instructions" api:"required"`
	EditOptions      param.Field[EditOptionsParam] `json:"edit_options"`
	// Form schema for PDF forms. List of widgets with their types, descriptions, and
	// bounding boxes. Only works for PDFs.
	FormSchema param.Field[[]EditWidgetParam] `json:"form_schema"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool]                      `json:"priority"`
	Webhook  param.Field[EditAsyncNewParamsWebhook] `json:"webhook"`
}

func (r EditAsyncNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The URL of the document to be processed. You can provide one of the following:
//
//  1. A publicly available URL
//  2. A presigned S3 URL
//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
//     uploading a document
//
// Satisfied by [shared.UnionString], [UploadResponseParam].
type EditAsyncNewParamsDocumentURLUnion interface {
	ImplementsEditAsyncNewParamsDocumentURLUnion()
}

type EditAsyncNewParamsWebhook struct {
	// A list of Svix channels the message will be delivered down, omit to send to all
	// channels.
	Channels param.Field[[]string] `json:"channels"`
	// JSON metadata included in webhook request body
	Metadata param.Field[interface{}] `json:"metadata"`
	// The mode to use for webhook delivery. Defaults to 'disabled'. We recommend using
	// 'svix' for production environments.
	Mode param.Field[EditAsyncNewParamsWebhookMode] `json:"mode"`
	// The URL to send the webhook to (if using direct webhoook).
	URL param.Field[string] `json:"url"`
}

func (r EditAsyncNewParamsWebhook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The mode to use for webhook delivery. Defaults to 'disabled'. We recommend using
// 'svix' for production environments.
type EditAsyncNewParamsWebhookMode string

const (
	EditAsyncNewParamsWebhookModeDisabled EditAsyncNewParamsWebhookMode = "disabled"
	EditAsyncNewParamsWebhookModeSvix     EditAsyncNewParamsWebhookMode = "svix"
	EditAsyncNewParamsWebhookModeDirect   EditAsyncNewParamsWebhookMode = "direct"
)

func (r EditAsyncNewParamsWebhookMode) IsKnown() bool {
	switch r {
	case EditAsyncNewParamsWebhookModeDisabled, EditAsyncNewParamsWebhookModeSvix, EditAsyncNewParamsWebhookModeDirect:
		return true
	}
	return false
}
