// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reducto

import (
	"github.com/reductoai/reducto-go-sdk/internal/apijson"
	"github.com/reductoai/reducto-go-sdk/internal/param"
	"github.com/reductoai/reducto-go-sdk/option"
	"github.com/reductoai/reducto-go-sdk/shared"
)

// ConfigService contains methods and other services that help with interacting
// with the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigService] method instead.
type ConfigService struct {
	Options []option.RequestOption
}

// NewConfigService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConfigService(opts ...option.RequestOption) (r *ConfigService) {
	r = &ConfigService{}
	r.Options = opts
	return
}

type ExtractConfigParam struct {
	// The URL of the document to be processed. You can provide one of the following:
	//
	//  1. A publicly available URL
	//  2. A presigned S3 URL
	//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
	//     uploading a document
	//  4. A job_id (jobid://) or a list of job_ids (jobid://)
	DocumentURL param.Field[ExtractConfigDocumentURLUnionParam] `json:"document_url,required"`
	// The JSON schema to use for extraction.
	Schema          param.Field[interface{}]                           `json:"schema,required"`
	AdvancedOptions param.Field[shared.AdvancedProcessingOptionsParam] `json:"advanced_options"`
	// The configuration options for array extract
	ArrayExtract        param.Field[shared.ArrayExtractConfigParam]            `json:"array_extract"`
	ExperimentalOptions param.Field[shared.ExperimentalProcessingOptionsParam] `json:"experimental_options"`
	// If citations should be generated for the extracted content.
	GenerateCitations param.Field[bool]                              `json:"generate_citations"`
	Options           param.Field[shared.BaseProcessingOptionsParam] `json:"options"`
	// A system prompt to use for the extraction. This is a general prompt that is
	// applied to the entire document before any other prompts.
	SystemPrompt param.Field[string] `json:"system_prompt"`
}

func (r ExtractConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The URL of the document to be processed. You can provide one of the following:
//
//  1. A publicly available URL
//  2. A presigned S3 URL
//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
//     uploading a document
//  4. A job_id (jobid://) or a list of job_ids (jobid://)
//
// Satisfied by [shared.UnionString], [ExtractConfigDocumentURLArrayParam],
// [shared.UploadParam].
type ExtractConfigDocumentURLUnionParam interface {
	ImplementsExtractConfigDocumentURLUnionParam()
}

type ExtractConfigDocumentURLArrayParam []string

func (r ExtractConfigDocumentURLArrayParam) ImplementsExtractConfigDocumentURLUnionParam() {}

type ParseConfigParam struct {
	// The URL of the document to be processed. You can provide one of the following:
	//
	//  1. A publicly available URL
	//  2. A presigned S3 URL
	//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
	//     uploading a document
	DocumentURL         param.Field[ParseConfigDocumentURLUnionParam]          `json:"document_url,required"`
	AdvancedOptions     param.Field[shared.AdvancedProcessingOptionsParam]     `json:"advanced_options"`
	ExperimentalOptions param.Field[shared.ExperimentalProcessingOptionsParam] `json:"experimental_options"`
	Options             param.Field[shared.BaseProcessingOptionsParam]         `json:"options"`
}

func (r ParseConfigParam) MarshalJSON() (data []byte, err error) {
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
type ParseConfigDocumentURLUnionParam interface {
	ImplementsParseConfigDocumentURLUnionParam()
}
