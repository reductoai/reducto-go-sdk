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
func (r *EditService) Submit(ctx context.Context, body EditSubmitParams, opts ...option.RequestOption) (res *EditResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "edit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type BoundingBox struct {
	Height float64 `json:"height" api:"required"`
	Left   float64 `json:"left" api:"required"`
	// The page number of the bounding box (1-indexed).
	Page  int64   `json:"page" api:"required"`
	Top   float64 `json:"top" api:"required"`
	Width float64 `json:"width" api:"required"`
	// The page number in the original document of the bounding box (1-indexed).
	OriginalPage int64           `json:"original_page"`
	JSON         boundingBoxJSON `json:"-"`
}

// boundingBoxJSON contains the JSON metadata for the struct [BoundingBox]
type boundingBoxJSON struct {
	Height       apijson.Field
	Left         apijson.Field
	Page         apijson.Field
	Top          apijson.Field
	Width        apijson.Field
	OriginalPage apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BoundingBox) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r boundingBoxJSON) RawJSON() string {
	return r.raw
}

type BoundingBoxParam struct {
	Height param.Field[float64] `json:"height" api:"required"`
	Left   param.Field[float64] `json:"left" api:"required"`
	// The page number of the bounding box (1-indexed).
	Page  param.Field[int64]   `json:"page" api:"required"`
	Top   param.Field[float64] `json:"top" api:"required"`
	Width param.Field[float64] `json:"width" api:"required"`
	// The page number in the original document of the bounding box (1-indexed).
	OriginalPage param.Field[int64] `json:"original_page"`
}

func (r BoundingBoxParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EditOptionsParam struct {
	// The color to use for edits, in hex format.
	Color param.Field[string] `json:"color"`
	// If True, creates overflow pages for text that doesn't fit in form fields.
	// Defaults to False.
	EnableOverflowPages param.Field[bool] `json:"enable_overflow_pages"`
	// If True, flattens form fields after filling, converting them to static content.
	// Defaults to False.
	Flatten param.Field[bool] `json:"flatten"`
	// The font size (in points) to use for filled text fields. If not specified, font
	// size is automatically calculated based on field dimensions.
	FontSize param.Field[float64] `json:"font_size"`
	// The LLM provider to use for edit processing. If not specified, defaults to
	// 'google'
	LlmProviderPreference param.Field[EditOptionsLlmProviderPreference] `json:"llm_provider_preference"`
}

func (r EditOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The LLM provider to use for edit processing. If not specified, defaults to
// 'google'
type EditOptionsLlmProviderPreference string

const (
	EditOptionsLlmProviderPreferenceOpenAI    EditOptionsLlmProviderPreference = "openai"
	EditOptionsLlmProviderPreferenceAnthropic EditOptionsLlmProviderPreference = "anthropic"
	EditOptionsLlmProviderPreferenceGoogle    EditOptionsLlmProviderPreference = "google"
)

func (r EditOptionsLlmProviderPreference) IsKnown() bool {
	switch r {
	case EditOptionsLlmProviderPreferenceOpenAI, EditOptionsLlmProviderPreferenceAnthropic, EditOptionsLlmProviderPreferenceGoogle:
		return true
	}
	return false
}

type EditResponse struct {
	// Presigned URL to download the edited document.
	DocumentURL string `json:"document_url" api:"required"`
	// Form schema for PDF forms. List of widgets with their types, descriptions, and
	// bounding boxes.
	FormSchema []EditWidget `json:"form_schema" api:"nullable"`
	// Usage information for the edit operation, including number of pages and credits
	// charged.
	Usage ParseUsage       `json:"usage" api:"nullable"`
	JSON  editResponseJSON `json:"-"`
}

// editResponseJSON contains the JSON metadata for the struct [EditResponse]
type editResponseJSON struct {
	DocumentURL apijson.Field
	FormSchema  apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editResponseJSON) RawJSON() string {
	return r.raw
}

func (r EditResponse) implementsJobGetResponseAsyncJobResponseResult() {}

func (r EditResponse) implementsJobGetResponseEnhancedAsyncJobResponseResult() {}

type EditWidget struct {
	// Bounding box coordinates of the widget
	Bbox BoundingBox `json:"bbox" api:"required"`
	// Description of the widget extracted from the document
	Description string `json:"description" api:"required"`
	// Type of the form widget
	Type EditWidgetType `json:"type" api:"required"`
	// If True (default), the system will attempt to fill this widget. If False, the
	// widget will be created but intentionally left unfilled.
	Fill bool `json:"fill"`
	// Font size in points for this specific field. Takes priority over the global
	// font_size in EditOptions. If not set, falls back to the global font_size, then
	// to auto-calculated sizing.
	FontSize float64 `json:"font_size" api:"nullable"`
	// If provided, this value will be used directly instead of attempting to
	// intelligently determine the field value.
	Value string         `json:"value" api:"nullable"`
	JSON  editWidgetJSON `json:"-"`
}

// editWidgetJSON contains the JSON metadata for the struct [EditWidget]
type editWidgetJSON struct {
	Bbox        apijson.Field
	Description apijson.Field
	Type        apijson.Field
	Fill        apijson.Field
	FontSize    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EditWidget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editWidgetJSON) RawJSON() string {
	return r.raw
}

// Type of the form widget
type EditWidgetType string

const (
	EditWidgetTypeText     EditWidgetType = "text"
	EditWidgetTypeCheckbox EditWidgetType = "checkbox"
	EditWidgetTypeRadio    EditWidgetType = "radio"
	EditWidgetTypeDropdown EditWidgetType = "dropdown"
	EditWidgetTypeBarcode  EditWidgetType = "barcode"
)

func (r EditWidgetType) IsKnown() bool {
	switch r {
	case EditWidgetTypeText, EditWidgetTypeCheckbox, EditWidgetTypeRadio, EditWidgetTypeDropdown, EditWidgetTypeBarcode:
		return true
	}
	return false
}

type EditWidgetParam struct {
	// Bounding box coordinates of the widget
	Bbox param.Field[BoundingBoxParam] `json:"bbox" api:"required"`
	// Description of the widget extracted from the document
	Description param.Field[string] `json:"description" api:"required"`
	// Type of the form widget
	Type param.Field[EditWidgetType] `json:"type" api:"required"`
	// If True (default), the system will attempt to fill this widget. If False, the
	// widget will be created but intentionally left unfilled.
	Fill param.Field[bool] `json:"fill"`
	// Font size in points for this specific field. Takes priority over the global
	// font_size in EditOptions. If not set, falls back to the global font_size, then
	// to auto-calculated sizing.
	FontSize param.Field[float64] `json:"font_size"`
	// If provided, this value will be used directly instead of attempting to
	// intelligently determine the field value.
	Value param.Field[string] `json:"value"`
}

func (r EditWidgetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EditSubmitParams struct {
	// The URL of the document to be processed. You can provide one of the following:
	//
	//  1. A publicly available URL
	//  2. A presigned S3 URL
	//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
	//     uploading a document
	DocumentURL param.Field[EditSubmitParamsDocumentURLUnion] `json:"document_url" api:"required"`
	// The instructions for the edit.
	EditInstructions param.Field[string]           `json:"edit_instructions" api:"required"`
	EditOptions      param.Field[EditOptionsParam] `json:"edit_options"`
	// Form schema for PDF forms. List of widgets with their types, descriptions, and
	// bounding boxes. Only works for PDFs.
	FormSchema param.Field[[]EditWidgetParam] `json:"form_schema"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool] `json:"priority"`
}

func (r EditSubmitParams) MarshalJSON() (data []byte, err error) {
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
type EditSubmitParamsDocumentURLUnion interface {
	ImplementsEditSubmitParamsDocumentURLUnion()
}
