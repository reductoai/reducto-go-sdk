// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reducto

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/reductoai/reducto-go-sdk/internal/apijson"
	"github.com/reductoai/reducto-go-sdk/internal/apiquery"
	"github.com/reductoai/reducto-go-sdk/internal/param"
	"github.com/reductoai/reducto-go-sdk/internal/requestconfig"
	"github.com/reductoai/reducto-go-sdk/option"
	"github.com/tidwall/gjson"
)

// JobService contains methods and other services that help with interacting with
// the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobService] method instead.
type JobService struct {
	Options []option.RequestOption
}

// NewJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJobService(opts ...option.RequestOption) (r *JobService) {
	r = &JobService{}
	r.Options = opts
	return
}

// Cancel Job
func (r *JobService) Cancel(ctx context.Context, jobID string, opts ...option.RequestOption) (res *JobCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cancel/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Retrieve Parse
func (r *JobService) Get(ctx context.Context, jobID string, opts ...option.RequestOption) (res *JobGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("job/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get Jobs
func (r *JobService) GetAll(ctx context.Context, query JobGetAllParams, opts ...option.RequestOption) (res *JobGetAllResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ExtractResponse struct {
	// The citations corresponding to the extracted response.
	Citations []interface{} `json:"citations" api:"required,nullable"`
	// The extracted response in your provided schema. This is a list of dictionaries.
	// If disable_chunking is True (default), then it will be a list of length one.
	Result []interface{} `json:"result" api:"required"`
	Usage  ExtractUsage  `json:"usage" api:"required"`
	JobID  string        `json:"job_id" api:"nullable"`
	// The link to the studio pipeline for the document.
	StudioLink string              `json:"studio_link" api:"nullable"`
	JSON       extractResponseJSON `json:"-"`
}

// extractResponseJSON contains the JSON metadata for the struct [ExtractResponse]
type extractResponseJSON struct {
	Citations   apijson.Field
	Result      apijson.Field
	Usage       apijson.Field
	JobID       apijson.Field
	StudioLink  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtractResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extractResponseJSON) RawJSON() string {
	return r.raw
}

func (r ExtractResponse) implementsPipelineResponseResultExtractUnion() {}

func (r ExtractResponse) implementsPipelineResponseResultExtractExtractVariant0Result() {}

func (r ExtractResponse) implementsJobGetResponseAsyncJobResponseResult() {}

func (r ExtractResponse) implementsJobGetResponseEnhancedAsyncJobResponseResult() {}

type JobCancelResponse = interface{}

type JobGetResponse struct {
	Status JobGetResponseStatus `json:"status" api:"required"`
	// This field can have the runtime type of [interface{}].
	Bucket    interface{} `json:"bucket"`
	CreatedAt time.Time   `json:"created_at" api:"nullable" format:"date-time"`
	Duration  float64     `json:"duration" api:"nullable"`
	NumPages  int64       `json:"num_pages" api:"nullable"`
	Progress  float64     `json:"progress" api:"nullable"`
	RawConfig string      `json:"raw_config" api:"nullable"`
	Reason    string      `json:"reason" api:"nullable"`
	// This field can have the runtime type of [JobGetResponseAsyncJobResponseResult],
	// [JobGetResponseEnhancedAsyncJobResponseResult].
	Result interface{} `json:"result"`
	// This field can have the runtime type of [interface{}].
	Source     interface{}        `json:"source"`
	TotalPages int64              `json:"total_pages" api:"nullable"`
	Type       JobGetResponseType `json:"type" api:"nullable"`
	JSON       jobGetResponseJSON `json:"-"`
	union      JobGetResponseUnion
}

// jobGetResponseJSON contains the JSON metadata for the struct [JobGetResponse]
type jobGetResponseJSON struct {
	Status      apijson.Field
	Bucket      apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	NumPages    apijson.Field
	Progress    apijson.Field
	RawConfig   apijson.Field
	Reason      apijson.Field
	Result      apijson.Field
	Source      apijson.Field
	TotalPages  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r jobGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *JobGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = JobGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [JobGetResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [JobGetResponseAsyncJobResponse],
// [JobGetResponseEnhancedAsyncJobResponse].
func (r JobGetResponse) AsUnion() JobGetResponseUnion {
	return r.union
}

// Union satisfied by [JobGetResponseAsyncJobResponse] or
// [JobGetResponseEnhancedAsyncJobResponse].
type JobGetResponseUnion interface {
	implementsJobGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*JobGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(JobGetResponseAsyncJobResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(JobGetResponseEnhancedAsyncJobResponse{}),
		},
	)
}

type JobGetResponseAsyncJobResponse struct {
	Status   JobGetResponseAsyncJobResponseStatus `json:"status" api:"required"`
	Progress float64                              `json:"progress" api:"nullable"`
	Reason   string                               `json:"reason" api:"nullable"`
	// Response from classify job - returned when polling /job/{job_id}
	Result JobGetResponseAsyncJobResponseResult `json:"result" api:"nullable"`
	JSON   jobGetResponseAsyncJobResponseJSON   `json:"-"`
}

// jobGetResponseAsyncJobResponseJSON contains the JSON metadata for the struct
// [JobGetResponseAsyncJobResponse]
type jobGetResponseAsyncJobResponseJSON struct {
	Status      apijson.Field
	Progress    apijson.Field
	Reason      apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseAsyncJobResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseAsyncJobResponseJSON) RawJSON() string {
	return r.raw
}

func (r JobGetResponseAsyncJobResponse) implementsJobGetResponse() {}

type JobGetResponseAsyncJobResponseStatus string

const (
	JobGetResponseAsyncJobResponseStatusPending   JobGetResponseAsyncJobResponseStatus = "Pending"
	JobGetResponseAsyncJobResponseStatusCompleted JobGetResponseAsyncJobResponseStatus = "Completed"
	JobGetResponseAsyncJobResponseStatusFailed    JobGetResponseAsyncJobResponseStatus = "Failed"
	JobGetResponseAsyncJobResponseStatusIdle      JobGetResponseAsyncJobResponseStatus = "Idle"
)

func (r JobGetResponseAsyncJobResponseStatus) IsKnown() bool {
	switch r {
	case JobGetResponseAsyncJobResponseStatusPending, JobGetResponseAsyncJobResponseStatusCompleted, JobGetResponseAsyncJobResponseStatusFailed, JobGetResponseAsyncJobResponseStatusIdle:
		return true
	}
	return false
}

// Response from classify job - returned when polling /job/{job_id}
type JobGetResponseAsyncJobResponseResult struct {
	// This field can have the runtime type of [[]interface{}].
	Citations interface{} `json:"citations"`
	// Presigned URL to download the edited document.
	DocumentURL string `json:"document_url"`
	// The duration of the parse request in seconds.
	Duration float64 `json:"duration" api:"nullable"`
	// This field can have the runtime type of [[]EditWidget].
	FormSchema interface{} `json:"form_schema"`
	JobID      string      `json:"job_id" api:"nullable"`
	// The storage URL of the converted PDF file.
	PdfURL string `json:"pdf_url" api:"nullable"`
	// This field can have the runtime type of [ClassifyResponseResponseConfidence].
	ResponseConfidence interface{} `json:"response_confidence"`
	// This field can have the runtime type of [ParseResponseResult], [[]interface{}],
	// [SplitResponseResult], [PipelineResponseResult], [ClassifyResponseResult].
	Result interface{} `json:"result"`
	// The link to the studio pipeline for the document.
	StudioLink string `json:"studio_link" api:"nullable"`
	// This field can have the runtime type of [ParseUsage], [ExtractUsage].
	Usage interface{}                              `json:"usage"`
	JSON  jobGetResponseAsyncJobResponseResultJSON `json:"-"`
	union JobGetResponseAsyncJobResponseResultUnion
}

// jobGetResponseAsyncJobResponseResultJSON contains the JSON metadata for the
// struct [JobGetResponseAsyncJobResponseResult]
type jobGetResponseAsyncJobResponseResultJSON struct {
	Citations          apijson.Field
	DocumentURL        apijson.Field
	Duration           apijson.Field
	FormSchema         apijson.Field
	JobID              apijson.Field
	PdfURL             apijson.Field
	ResponseConfidence apijson.Field
	Result             apijson.Field
	StudioLink         apijson.Field
	Usage              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r jobGetResponseAsyncJobResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *JobGetResponseAsyncJobResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = JobGetResponseAsyncJobResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [JobGetResponseAsyncJobResponseResultUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [ParseResponse], [ExtractResponse],
// [SplitResponse], [EditResponse], [PipelineResponse], [V3Extract],
// [ClassifyResponse].
func (r JobGetResponseAsyncJobResponseResult) AsUnion() JobGetResponseAsyncJobResponseResultUnion {
	return r.union
}

// Response from classify job - returned when polling /job/{job_id}
//
// Union satisfied by [ParseResponse], [ExtractResponse], [SplitResponse],
// [EditResponse], [PipelineResponse], [V3Extract] or [ClassifyResponse].
type JobGetResponseAsyncJobResponseResultUnion interface {
	implementsJobGetResponseAsyncJobResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*JobGetResponseAsyncJobResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParseResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExtractResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplitResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EditResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V3Extract{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ClassifyResponse{}),
		},
	)
}

type JobGetResponseEnhancedAsyncJobResponse struct {
	Status    JobGetResponseEnhancedAsyncJobResponseStatus `json:"status" api:"required"`
	Bucket    interface{}                                  `json:"bucket"`
	CreatedAt time.Time                                    `json:"created_at" api:"nullable" format:"date-time"`
	Duration  float64                                      `json:"duration" api:"nullable"`
	NumPages  int64                                        `json:"num_pages" api:"nullable"`
	Progress  float64                                      `json:"progress" api:"nullable"`
	RawConfig string                                       `json:"raw_config" api:"nullable"`
	Reason    string                                       `json:"reason" api:"nullable"`
	// Response from classify job - returned when polling /job/{job_id}
	Result     JobGetResponseEnhancedAsyncJobResponseResult `json:"result" api:"nullable"`
	Source     interface{}                                  `json:"source"`
	TotalPages int64                                        `json:"total_pages" api:"nullable"`
	Type       JobGetResponseEnhancedAsyncJobResponseType   `json:"type" api:"nullable"`
	JSON       jobGetResponseEnhancedAsyncJobResponseJSON   `json:"-"`
}

// jobGetResponseEnhancedAsyncJobResponseJSON contains the JSON metadata for the
// struct [JobGetResponseEnhancedAsyncJobResponse]
type jobGetResponseEnhancedAsyncJobResponseJSON struct {
	Status      apijson.Field
	Bucket      apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	NumPages    apijson.Field
	Progress    apijson.Field
	RawConfig   apijson.Field
	Reason      apijson.Field
	Result      apijson.Field
	Source      apijson.Field
	TotalPages  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseEnhancedAsyncJobResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseEnhancedAsyncJobResponseJSON) RawJSON() string {
	return r.raw
}

func (r JobGetResponseEnhancedAsyncJobResponse) implementsJobGetResponse() {}

type JobGetResponseEnhancedAsyncJobResponseStatus string

const (
	JobGetResponseEnhancedAsyncJobResponseStatusPending   JobGetResponseEnhancedAsyncJobResponseStatus = "Pending"
	JobGetResponseEnhancedAsyncJobResponseStatusCompleted JobGetResponseEnhancedAsyncJobResponseStatus = "Completed"
	JobGetResponseEnhancedAsyncJobResponseStatusFailed    JobGetResponseEnhancedAsyncJobResponseStatus = "Failed"
	JobGetResponseEnhancedAsyncJobResponseStatusIdle      JobGetResponseEnhancedAsyncJobResponseStatus = "Idle"
)

func (r JobGetResponseEnhancedAsyncJobResponseStatus) IsKnown() bool {
	switch r {
	case JobGetResponseEnhancedAsyncJobResponseStatusPending, JobGetResponseEnhancedAsyncJobResponseStatusCompleted, JobGetResponseEnhancedAsyncJobResponseStatusFailed, JobGetResponseEnhancedAsyncJobResponseStatusIdle:
		return true
	}
	return false
}

// Response from classify job - returned when polling /job/{job_id}
type JobGetResponseEnhancedAsyncJobResponseResult struct {
	// This field can have the runtime type of [[]interface{}].
	Citations interface{} `json:"citations"`
	// Presigned URL to download the edited document.
	DocumentURL string `json:"document_url"`
	// The duration of the parse request in seconds.
	Duration float64 `json:"duration" api:"nullable"`
	// This field can have the runtime type of [[]EditWidget].
	FormSchema interface{} `json:"form_schema"`
	JobID      string      `json:"job_id" api:"nullable"`
	// The storage URL of the converted PDF file.
	PdfURL string `json:"pdf_url" api:"nullable"`
	// This field can have the runtime type of [ClassifyResponseResponseConfidence].
	ResponseConfidence interface{} `json:"response_confidence"`
	// This field can have the runtime type of [ParseResponseResult], [[]interface{}],
	// [SplitResponseResult], [PipelineResponseResult], [ClassifyResponseResult].
	Result interface{} `json:"result"`
	// The link to the studio pipeline for the document.
	StudioLink string `json:"studio_link" api:"nullable"`
	// This field can have the runtime type of [ParseUsage], [ExtractUsage].
	Usage interface{}                                      `json:"usage"`
	JSON  jobGetResponseEnhancedAsyncJobResponseResultJSON `json:"-"`
	union JobGetResponseEnhancedAsyncJobResponseResultUnion
}

// jobGetResponseEnhancedAsyncJobResponseResultJSON contains the JSON metadata for
// the struct [JobGetResponseEnhancedAsyncJobResponseResult]
type jobGetResponseEnhancedAsyncJobResponseResultJSON struct {
	Citations          apijson.Field
	DocumentURL        apijson.Field
	Duration           apijson.Field
	FormSchema         apijson.Field
	JobID              apijson.Field
	PdfURL             apijson.Field
	ResponseConfidence apijson.Field
	Result             apijson.Field
	StudioLink         apijson.Field
	Usage              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r jobGetResponseEnhancedAsyncJobResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *JobGetResponseEnhancedAsyncJobResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = JobGetResponseEnhancedAsyncJobResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [JobGetResponseEnhancedAsyncJobResponseResultUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [ParseResponse], [ExtractResponse],
// [SplitResponse], [EditResponse], [PipelineResponse], [V3Extract],
// [ClassifyResponse].
func (r JobGetResponseEnhancedAsyncJobResponseResult) AsUnion() JobGetResponseEnhancedAsyncJobResponseResultUnion {
	return r.union
}

// Response from classify job - returned when polling /job/{job_id}
//
// Union satisfied by [ParseResponse], [ExtractResponse], [SplitResponse],
// [EditResponse], [PipelineResponse], [V3Extract] or [ClassifyResponse].
type JobGetResponseEnhancedAsyncJobResponseResultUnion interface {
	implementsJobGetResponseEnhancedAsyncJobResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*JobGetResponseEnhancedAsyncJobResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParseResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExtractResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SplitResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EditResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V3Extract{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ClassifyResponse{}),
		},
	)
}

type JobGetResponseEnhancedAsyncJobResponseType string

const (
	JobGetResponseEnhancedAsyncJobResponseTypeParse    JobGetResponseEnhancedAsyncJobResponseType = "Parse"
	JobGetResponseEnhancedAsyncJobResponseTypeExtract  JobGetResponseEnhancedAsyncJobResponseType = "Extract"
	JobGetResponseEnhancedAsyncJobResponseTypeSplit    JobGetResponseEnhancedAsyncJobResponseType = "Split"
	JobGetResponseEnhancedAsyncJobResponseTypeEdit     JobGetResponseEnhancedAsyncJobResponseType = "Edit"
	JobGetResponseEnhancedAsyncJobResponseTypePipeline JobGetResponseEnhancedAsyncJobResponseType = "Pipeline"
	JobGetResponseEnhancedAsyncJobResponseTypeClassify JobGetResponseEnhancedAsyncJobResponseType = "Classify"
)

func (r JobGetResponseEnhancedAsyncJobResponseType) IsKnown() bool {
	switch r {
	case JobGetResponseEnhancedAsyncJobResponseTypeParse, JobGetResponseEnhancedAsyncJobResponseTypeExtract, JobGetResponseEnhancedAsyncJobResponseTypeSplit, JobGetResponseEnhancedAsyncJobResponseTypeEdit, JobGetResponseEnhancedAsyncJobResponseTypePipeline, JobGetResponseEnhancedAsyncJobResponseTypeClassify:
		return true
	}
	return false
}

type JobGetResponseStatus string

const (
	JobGetResponseStatusPending   JobGetResponseStatus = "Pending"
	JobGetResponseStatusCompleted JobGetResponseStatus = "Completed"
	JobGetResponseStatusFailed    JobGetResponseStatus = "Failed"
	JobGetResponseStatusIdle      JobGetResponseStatus = "Idle"
)

func (r JobGetResponseStatus) IsKnown() bool {
	switch r {
	case JobGetResponseStatusPending, JobGetResponseStatusCompleted, JobGetResponseStatusFailed, JobGetResponseStatusIdle:
		return true
	}
	return false
}

type JobGetResponseType string

const (
	JobGetResponseTypeParse    JobGetResponseType = "Parse"
	JobGetResponseTypeExtract  JobGetResponseType = "Extract"
	JobGetResponseTypeSplit    JobGetResponseType = "Split"
	JobGetResponseTypeEdit     JobGetResponseType = "Edit"
	JobGetResponseTypePipeline JobGetResponseType = "Pipeline"
	JobGetResponseTypeClassify JobGetResponseType = "Classify"
)

func (r JobGetResponseType) IsKnown() bool {
	switch r {
	case JobGetResponseTypeParse, JobGetResponseTypeExtract, JobGetResponseTypeSplit, JobGetResponseTypeEdit, JobGetResponseTypePipeline, JobGetResponseTypeClassify:
		return true
	}
	return false
}

type JobGetAllResponse struct {
	// List of jobs with their job_id, status, type, raw_config, created_at, num_pages
	// and duration
	Jobs []JobGetAllResponseJob `json:"jobs" api:"required"`
	// Cursor to fetch the next page of results. If null, there are no more results.
	NextCursor string                `json:"next_cursor" api:"nullable"`
	JSON       jobGetAllResponseJSON `json:"-"`
}

// jobGetAllResponseJSON contains the JSON metadata for the struct
// [JobGetAllResponse]
type jobGetAllResponseJSON struct {
	Jobs        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetAllResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetAllResponseJSON) RawJSON() string {
	return r.raw
}

type JobGetAllResponseJob struct {
	CreatedAt  time.Time                   `json:"created_at" api:"required" format:"date-time"`
	Duration   float64                     `json:"duration" api:"required,nullable"`
	JobID      string                      `json:"job_id" api:"required"`
	NumPages   int64                       `json:"num_pages" api:"required,nullable"`
	RawConfig  string                      `json:"raw_config" api:"required"`
	Status     JobGetAllResponseJobsStatus `json:"status" api:"required"`
	TotalPages int64                       `json:"total_pages" api:"required,nullable"`
	Type       JobGetAllResponseJobsType   `json:"type" api:"required"`
	Bucket     interface{}                 `json:"bucket"`
	Source     interface{}                 `json:"source"`
	JSON       jobGetAllResponseJobJSON    `json:"-"`
}

// jobGetAllResponseJobJSON contains the JSON metadata for the struct
// [JobGetAllResponseJob]
type jobGetAllResponseJobJSON struct {
	CreatedAt   apijson.Field
	Duration    apijson.Field
	JobID       apijson.Field
	NumPages    apijson.Field
	RawConfig   apijson.Field
	Status      apijson.Field
	TotalPages  apijson.Field
	Type        apijson.Field
	Bucket      apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetAllResponseJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetAllResponseJobJSON) RawJSON() string {
	return r.raw
}

type JobGetAllResponseJobsStatus string

const (
	JobGetAllResponseJobsStatusPending    JobGetAllResponseJobsStatus = "Pending"
	JobGetAllResponseJobsStatusCompleted  JobGetAllResponseJobsStatus = "Completed"
	JobGetAllResponseJobsStatusFailed     JobGetAllResponseJobsStatus = "Failed"
	JobGetAllResponseJobsStatusIdle       JobGetAllResponseJobsStatus = "Idle"
	JobGetAllResponseJobsStatusInProgress JobGetAllResponseJobsStatus = "InProgress"
	JobGetAllResponseJobsStatusCompleting JobGetAllResponseJobsStatus = "Completing"
	JobGetAllResponseJobsStatusCancelled  JobGetAllResponseJobsStatus = "Cancelled"
)

func (r JobGetAllResponseJobsStatus) IsKnown() bool {
	switch r {
	case JobGetAllResponseJobsStatusPending, JobGetAllResponseJobsStatusCompleted, JobGetAllResponseJobsStatusFailed, JobGetAllResponseJobsStatusIdle, JobGetAllResponseJobsStatusInProgress, JobGetAllResponseJobsStatusCompleting, JobGetAllResponseJobsStatusCancelled:
		return true
	}
	return false
}

type JobGetAllResponseJobsType string

const (
	JobGetAllResponseJobsTypeParse    JobGetAllResponseJobsType = "Parse"
	JobGetAllResponseJobsTypeExtract  JobGetAllResponseJobsType = "Extract"
	JobGetAllResponseJobsTypeSplit    JobGetAllResponseJobsType = "Split"
	JobGetAllResponseJobsTypeEdit     JobGetAllResponseJobsType = "Edit"
	JobGetAllResponseJobsTypePipeline JobGetAllResponseJobsType = "Pipeline"
	JobGetAllResponseJobsTypeClassify JobGetAllResponseJobsType = "Classify"
)

func (r JobGetAllResponseJobsType) IsKnown() bool {
	switch r {
	case JobGetAllResponseJobsTypeParse, JobGetAllResponseJobsTypeExtract, JobGetAllResponseJobsTypeSplit, JobGetAllResponseJobsTypeEdit, JobGetAllResponseJobsTypePipeline, JobGetAllResponseJobsTypeClassify:
		return true
	}
	return false
}

type JobGetAllParams struct {
	// Cursor for pagination. Use the next_cursor from the previous response to fetch
	// the next page.
	Cursor param.Field[string] `query:"cursor"`
	// Exclude raw_config from response to reduce size
	ExcludeConfigs param.Field[bool] `query:"exclude_configs"`
	// Maximum number of jobs to return per page. Defaults to 100, max 500.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [JobGetAllParams]'s query parameters as `url.Values`.
func (r JobGetAllParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
