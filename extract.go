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
func (r *ExtractService) Run(ctx context.Context, body ExtractRunParams, opts ...option.RequestOption) (res *shared.ExtractResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "extract"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Extract Async
func (r *ExtractService) RunJob(ctx context.Context, body ExtractRunJobParams, opts ...option.RequestOption) (res *ExtractRunJobResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "extract_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ExtractRunJobResponse struct {
	JobID string                    `json:"job_id,required"`
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
	ExtractConfig ExtractConfigParam `json:"extract_config,required"`
}

func (r ExtractRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ExtractConfig)
}

type ExtractRunJobParams struct {
	// The URL of the document to be processed. You can provide one of the following:
	//
	//  1. A publicly available URL
	//  2. A presigned S3 URL
	//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
	//     uploading a document
	//  4. A job_id (jobid://) or a list of job_ids (jobid://)
	DocumentURL param.Field[ExtractRunJobParamsDocumentURLUnion] `json:"document_url,required"`
	// The JSON schema to use for extraction.
	Schema          param.Field[interface{}]                           `json:"schema,required"`
	AdvancedOptions param.Field[shared.AdvancedProcessingOptionsParam] `json:"advanced_options"`
	// The configuration options for array extract
	ArrayExtract        param.Field[shared.ArrayExtractConfigParam]            `json:"array_extract"`
	ExperimentalOptions param.Field[shared.ExperimentalProcessingOptionsParam] `json:"experimental_options"`
	// If citations should be generated for the extracted content.
	GenerateCitations param.Field[bool]                              `json:"generate_citations"`
	Options           param.Field[shared.BaseProcessingOptionsParam] `json:"options"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool] `json:"priority"`
	// A system prompt to use for the extraction. This is a general prompt that is
	// applied to the entire document before any other prompts.
	SystemPrompt param.Field[string]                       `json:"system_prompt"`
	Webhook      param.Field[shared.WebhookConfigNewParam] `json:"webhook"`
}

func (r ExtractRunJobParams) MarshalJSON() (data []byte, err error) {
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
// Satisfied by [shared.UnionString], [ExtractRunJobParamsDocumentURLArray],
// [shared.UploadParam].
type ExtractRunJobParamsDocumentURLUnion interface {
	ImplementsExtractRunJobParamsDocumentURLUnion()
}

type ExtractRunJobParamsDocumentURLArray []string

func (r ExtractRunJobParamsDocumentURLArray) ImplementsExtractRunJobParamsDocumentURLUnion() {}
