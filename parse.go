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

// ParseService contains methods and other services that help with interacting with
// the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewParseService] method instead.
type ParseService struct {
	Options []option.RequestOption
}

// NewParseService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewParseService(opts ...option.RequestOption) (r *ParseService) {
	r = &ParseService{}
	r.Options = opts
	return
}

// Parse
func (r *ParseService) Run(ctx context.Context, body ParseRunParams, opts ...option.RequestOption) (res *shared.ParseResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "parse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Async Parse
func (r *ParseService) RunJob(ctx context.Context, body ParseRunJobParams, opts ...option.RequestOption) (res *ParseRunJobResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "parse_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ParseRunJobResponse struct {
	JobID string                  `json:"job_id,required"`
	JSON  parseRunJobResponseJSON `json:"-"`
}

// parseRunJobResponseJSON contains the JSON metadata for the struct
// [ParseRunJobResponse]
type parseRunJobResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParseRunJobResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseRunJobResponseJSON) RawJSON() string {
	return r.raw
}

type ParseRunParams struct {
	// The URL of the document to be processed. You can provide one of the following:
	//
	//  1. A publicly available URL
	//  2. A presigned S3 URL
	//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
	//     uploading a document
	DocumentURL         param.Field[string]                                    `json:"document_url,required"`
	AdvancedOptions     param.Field[shared.AdvancedProcessingOptionsParam]     `json:"advanced_options"`
	ExperimentalOptions param.Field[shared.ExperimentalProcessingOptionsParam] `json:"experimental_options"`
	Options             param.Field[shared.BaseProcessingOptionsParam]         `json:"options"`
}

func (r ParseRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ParseRunJobParams struct {
	// The URL of the document to be processed. You can provide one of the following:
	//
	//  1. A publicly available URL
	//  2. A presigned S3 URL
	//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
	//     uploading a document
	DocumentURL         param.Field[string]                                    `json:"document_url,required"`
	AdvancedOptions     param.Field[shared.AdvancedProcessingOptionsParam]     `json:"advanced_options"`
	ExperimentalOptions param.Field[shared.ExperimentalProcessingOptionsParam] `json:"experimental_options"`
	Options             param.Field[shared.BaseProcessingOptionsParam]         `json:"options"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool]                         `json:"priority"`
	Webhook  param.Field[shared.WebhookConfigNewParam] `json:"webhook"`
}

func (r ParseRunJobParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
