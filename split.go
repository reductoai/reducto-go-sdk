// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reductoai

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/reductoai-go/internal/apijson"
	"github.com/stainless-sdks/reductoai-go/internal/param"
	"github.com/stainless-sdks/reductoai-go/internal/requestconfig"
	"github.com/stainless-sdks/reductoai-go/option"
	"github.com/stainless-sdks/reductoai-go/shared"
)

// SplitService contains methods and other services that help with interacting with
// the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSplitService] method instead.
type SplitService struct {
	Options []option.RequestOption
}

// NewSplitService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSplitService(opts ...option.RequestOption) (r *SplitService) {
	r = &SplitService{}
	r.Options = opts
	return
}

// Split
func (r *SplitService) Run(ctx context.Context, body SplitRunParams, opts ...option.RequestOption) (res *shared.SplitResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "split"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Split Async
func (r *SplitService) RunJob(ctx context.Context, body SplitRunJobParams, opts ...option.RequestOption) (res *SplitRunJobResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "split_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SplitRunJobResponse struct {
	JobID string                  `json:"job_id,required"`
	JSON  splitRunJobResponseJSON `json:"-"`
}

// splitRunJobResponseJSON contains the JSON metadata for the struct
// [SplitRunJobResponse]
type splitRunJobResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitRunJobResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitRunJobResponseJSON) RawJSON() string {
	return r.raw
}

type SplitRunParams struct {
	// The URL of the document to be processed. You can provide one of the following:
	//
	//  1. A publicly available URL
	//  2. A presigned S3 URL
	//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
	//     uploading a document
	DocumentURL param.Field[string] `json:"document_url,required"`
	// The configuration options for processing the document.
	SplitDescription    param.Field[[]shared.SplitCategoryParam]               `json:"split_description,required"`
	AdvancedOptions     param.Field[shared.AdvancedProcessingOptionsParam]     `json:"advanced_options"`
	ExperimentalOptions param.Field[shared.ExperimentalProcessingOptionsParam] `json:"experimental_options"`
	Options             param.Field[shared.BaseProcessingOptionsParam]         `json:"options"`
	// The rules for splitting the document.
	SplitRules param.Field[string] `json:"split_rules"`
}

func (r SplitRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SplitRunJobParams struct {
	// The URL of the document to be processed. You can provide one of the following:
	//
	//  1. A publicly available URL
	//  2. A presigned S3 URL
	//  3. A reducto:// prefixed URL obtained from the /upload endpoint after directly
	//     uploading a document
	DocumentURL param.Field[string] `json:"document_url,required"`
	// The configuration options for processing the document.
	SplitDescription    param.Field[[]shared.SplitCategoryParam]               `json:"split_description,required"`
	AdvancedOptions     param.Field[shared.AdvancedProcessingOptionsParam]     `json:"advanced_options"`
	ExperimentalOptions param.Field[shared.ExperimentalProcessingOptionsParam] `json:"experimental_options"`
	Options             param.Field[shared.BaseProcessingOptionsParam]         `json:"options"`
	// If True, attempts to process the job with priority if the user has priority
	// processing budget available; by default, sync jobs are prioritized above async
	// jobs.
	Priority param.Field[bool] `json:"priority"`
	// The rules for splitting the document.
	SplitRules param.Field[string]                       `json:"split_rules"`
	Webhook    param.Field[shared.WebhookConfigNewParam] `json:"webhook"`
}

func (r SplitRunJobParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
