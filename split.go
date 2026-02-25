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
	opts = slices.Concat(r.Options, opts)
	path := "split"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Split Async
func (r *SplitService) RunJob(ctx context.Context, body SplitRunJobParams, opts ...option.RequestOption) (res *SplitRunJobResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "split_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SplitRunJobResponse struct {
	JobID string                  `json:"job_id" api:"required"`
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
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input param.Field[SplitRunParamsInputUnion] `json:"input" api:"required"`
	// The configuration options for processing the document.
	SplitDescription param.Field[[]shared.SplitCategoryParam] `json:"split_description" api:"required"`
	// The configuration options for parsing the document. If you are passing in a
	// jobid:// URL for the file, then this configuration will be ignored.
	Parsing param.Field[shared.ParseOptionsParam] `json:"parsing"`
	// The settings for split processing.
	Settings param.Field[SplitRunParamsSettings] `json:"settings"`
	// The prompt that describes rules for splitting the document.
	SplitRules param.Field[string] `json:"split_rules"`
}

func (r SplitRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For parse/split/extract pipelines, the URL of the document to be processed. You
// can provide one of the following: 1. A publicly available URL 2. A presigned S3
// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
// directly uploading a document 4. A jobid:// prefixed URL obtained from a
// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
// API only)
//
//	For edit pipelines, this should be a string containing the edit instructions
//
// Satisfied by [shared.UnionString], [SplitRunParamsInputArray],
// [shared.UploadParam].
type SplitRunParamsInputUnion interface {
	ImplementsSplitRunParamsInputUnion()
}

type SplitRunParamsInputArray []string

func (r SplitRunParamsInputArray) ImplementsSplitRunParamsInputUnion() {}

// The settings for split processing.
type SplitRunParamsSettings struct {
	// If tables should be truncated to the first few rows or if all content should be
	// preserved. truncate improves latency, preserve is recommended for cases where
	// partition_key is being used and the partition_key may be included within the
	// table. Defaults to truncate
	TableCutoff param.Field[SplitRunParamsSettingsTableCutoff] `json:"table_cutoff"`
}

func (r SplitRunParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If tables should be truncated to the first few rows or if all content should be
// preserved. truncate improves latency, preserve is recommended for cases where
// partition_key is being used and the partition_key may be included within the
// table. Defaults to truncate
type SplitRunParamsSettingsTableCutoff string

const (
	SplitRunParamsSettingsTableCutoffTruncate SplitRunParamsSettingsTableCutoff = "truncate"
	SplitRunParamsSettingsTableCutoffPreserve SplitRunParamsSettingsTableCutoff = "preserve"
)

func (r SplitRunParamsSettingsTableCutoff) IsKnown() bool {
	switch r {
	case SplitRunParamsSettingsTableCutoffTruncate, SplitRunParamsSettingsTableCutoffPreserve:
		return true
	}
	return false
}

type SplitRunJobParams struct {
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input param.Field[SplitRunJobParamsInputUnion] `json:"input" api:"required"`
	// The configuration options for processing the document.
	SplitDescription param.Field[[]shared.SplitCategoryParam] `json:"split_description" api:"required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async param.Field[shared.ConfigV3AsyncConfigParam] `json:"async"`
	// The configuration options for parsing the document. If you are passing in a
	// jobid:// URL for the file, then this configuration will be ignored.
	Parsing param.Field[shared.ParseOptionsParam] `json:"parsing"`
	// The settings for split processing.
	Settings param.Field[SplitRunJobParamsSettings] `json:"settings"`
	// The prompt that describes rules for splitting the document.
	SplitRules param.Field[string] `json:"split_rules"`
}

func (r SplitRunJobParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For parse/split/extract pipelines, the URL of the document to be processed. You
// can provide one of the following: 1. A publicly available URL 2. A presigned S3
// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
// directly uploading a document 4. A jobid:// prefixed URL obtained from a
// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
// API only)
//
//	For edit pipelines, this should be a string containing the edit instructions
//
// Satisfied by [shared.UnionString], [SplitRunJobParamsInputArray],
// [shared.UploadParam].
type SplitRunJobParamsInputUnion interface {
	ImplementsSplitRunJobParamsInputUnion()
}

type SplitRunJobParamsInputArray []string

func (r SplitRunJobParamsInputArray) ImplementsSplitRunJobParamsInputUnion() {}

// The settings for split processing.
type SplitRunJobParamsSettings struct {
	// If tables should be truncated to the first few rows or if all content should be
	// preserved. truncate improves latency, preserve is recommended for cases where
	// partition_key is being used and the partition_key may be included within the
	// table. Defaults to truncate
	TableCutoff param.Field[SplitRunJobParamsSettingsTableCutoff] `json:"table_cutoff"`
}

func (r SplitRunJobParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If tables should be truncated to the first few rows or if all content should be
// preserved. truncate improves latency, preserve is recommended for cases where
// partition_key is being used and the partition_key may be included within the
// table. Defaults to truncate
type SplitRunJobParamsSettingsTableCutoff string

const (
	SplitRunJobParamsSettingsTableCutoffTruncate SplitRunJobParamsSettingsTableCutoff = "truncate"
	SplitRunJobParamsSettingsTableCutoffPreserve SplitRunJobParamsSettingsTableCutoff = "preserve"
)

func (r SplitRunJobParamsSettingsTableCutoff) IsKnown() bool {
	switch r {
	case SplitRunJobParamsSettingsTableCutoffTruncate, SplitRunJobParamsSettingsTableCutoffPreserve:
		return true
	}
	return false
}
