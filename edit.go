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
func (r *EditService) Run(ctx context.Context, body EditRunParams, opts ...option.RequestOption) (res *shared.EditResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "edit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Edit Async
func (r *EditService) RunJob(ctx context.Context, body EditRunJobParams, opts ...option.RequestOption) (res *EditRunJobResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "edit_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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
	// Form schema for PDF forms. List of widgets with their types, descriptions, and
	// bounding boxes. Only works for PDFs.
	FormSchema param.Field[[]EditRunParamsFormSchema] `json:"form_schema"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool] `json:"priority"`
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
	// The LLM provider to use for edit processing. If not specified, defaults to
	// 'google'
	LlmProviderPreference param.Field[EditRunParamsEditOptionsLlmProviderPreference] `json:"llm_provider_preference"`
}

func (r EditRunParamsEditOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The LLM provider to use for edit processing. If not specified, defaults to
// 'google'
type EditRunParamsEditOptionsLlmProviderPreference string

const (
	EditRunParamsEditOptionsLlmProviderPreferenceOpenAI    EditRunParamsEditOptionsLlmProviderPreference = "openai"
	EditRunParamsEditOptionsLlmProviderPreferenceAnthropic EditRunParamsEditOptionsLlmProviderPreference = "anthropic"
	EditRunParamsEditOptionsLlmProviderPreferenceGoogle    EditRunParamsEditOptionsLlmProviderPreference = "google"
)

func (r EditRunParamsEditOptionsLlmProviderPreference) IsKnown() bool {
	switch r {
	case EditRunParamsEditOptionsLlmProviderPreferenceOpenAI, EditRunParamsEditOptionsLlmProviderPreferenceAnthropic, EditRunParamsEditOptionsLlmProviderPreferenceGoogle:
		return true
	}
	return false
}

type EditRunParamsFormSchema struct {
	// Bounding box coordinates of the widget
	Bbox param.Field[shared.BoundingBoxParam] `json:"bbox,required"`
	// Description of the widget extracted from the document
	Description param.Field[string] `json:"description,required"`
	// Type of the form widget
	Type param.Field[EditRunParamsFormSchemaType] `json:"type,required"`
	// If True (default), the system will attempt to fill this widget. If False, the
	// widget will be created but intentionally left unfilled.
	Fill param.Field[bool] `json:"fill"`
	// If provided, this value will be used directly instead of attempting to
	// intelligently determine the field value.
	Value param.Field[string] `json:"value"`
}

func (r EditRunParamsFormSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the form widget
type EditRunParamsFormSchemaType string

const (
	EditRunParamsFormSchemaTypeText     EditRunParamsFormSchemaType = "text"
	EditRunParamsFormSchemaTypeCheckbox EditRunParamsFormSchemaType = "checkbox"
	EditRunParamsFormSchemaTypeDropdown EditRunParamsFormSchemaType = "dropdown"
	EditRunParamsFormSchemaTypeBarcode  EditRunParamsFormSchemaType = "barcode"
)

func (r EditRunParamsFormSchemaType) IsKnown() bool {
	switch r {
	case EditRunParamsFormSchemaTypeText, EditRunParamsFormSchemaTypeCheckbox, EditRunParamsFormSchemaTypeDropdown, EditRunParamsFormSchemaTypeBarcode:
		return true
	}
	return false
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
	// Form schema for PDF forms. List of widgets with their types, descriptions, and
	// bounding boxes. Only works for PDFs.
	FormSchema param.Field[[]EditRunJobParamsFormSchema] `json:"form_schema"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool]                         `json:"priority"`
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
	// The LLM provider to use for edit processing. If not specified, defaults to
	// 'google'
	LlmProviderPreference param.Field[EditRunJobParamsEditOptionsLlmProviderPreference] `json:"llm_provider_preference"`
}

func (r EditRunJobParamsEditOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The LLM provider to use for edit processing. If not specified, defaults to
// 'google'
type EditRunJobParamsEditOptionsLlmProviderPreference string

const (
	EditRunJobParamsEditOptionsLlmProviderPreferenceOpenAI    EditRunJobParamsEditOptionsLlmProviderPreference = "openai"
	EditRunJobParamsEditOptionsLlmProviderPreferenceAnthropic EditRunJobParamsEditOptionsLlmProviderPreference = "anthropic"
	EditRunJobParamsEditOptionsLlmProviderPreferenceGoogle    EditRunJobParamsEditOptionsLlmProviderPreference = "google"
)

func (r EditRunJobParamsEditOptionsLlmProviderPreference) IsKnown() bool {
	switch r {
	case EditRunJobParamsEditOptionsLlmProviderPreferenceOpenAI, EditRunJobParamsEditOptionsLlmProviderPreferenceAnthropic, EditRunJobParamsEditOptionsLlmProviderPreferenceGoogle:
		return true
	}
	return false
}

type EditRunJobParamsFormSchema struct {
	// Bounding box coordinates of the widget
	Bbox param.Field[shared.BoundingBoxParam] `json:"bbox,required"`
	// Description of the widget extracted from the document
	Description param.Field[string] `json:"description,required"`
	// Type of the form widget
	Type param.Field[EditRunJobParamsFormSchemaType] `json:"type,required"`
	// If True (default), the system will attempt to fill this widget. If False, the
	// widget will be created but intentionally left unfilled.
	Fill param.Field[bool] `json:"fill"`
	// If provided, this value will be used directly instead of attempting to
	// intelligently determine the field value.
	Value param.Field[string] `json:"value"`
}

func (r EditRunJobParamsFormSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the form widget
type EditRunJobParamsFormSchemaType string

const (
	EditRunJobParamsFormSchemaTypeText     EditRunJobParamsFormSchemaType = "text"
	EditRunJobParamsFormSchemaTypeCheckbox EditRunJobParamsFormSchemaType = "checkbox"
	EditRunJobParamsFormSchemaTypeDropdown EditRunJobParamsFormSchemaType = "dropdown"
	EditRunJobParamsFormSchemaTypeBarcode  EditRunJobParamsFormSchemaType = "barcode"
)

func (r EditRunJobParamsFormSchemaType) IsKnown() bool {
	switch r {
	case EditRunJobParamsFormSchemaTypeText, EditRunJobParamsFormSchemaTypeCheckbox, EditRunJobParamsFormSchemaTypeDropdown, EditRunJobParamsFormSchemaTypeBarcode:
		return true
	}
	return false
}
