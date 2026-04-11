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
	"github.com/reductoai/reducto-go-sdk/shared"
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
	// This field can have the runtime type of
	// [JobGetResponseAsyncJobResponseResultUnion],
	// [JobGetResponseEnhancedAsyncJobResponseResultUnion].
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
	Result JobGetResponseAsyncJobResponseResultUnion `json:"result" api:"nullable"`
	JSON   jobGetResponseAsyncJobResponseJSON        `json:"-"`
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
//
// Union satisfied by [shared.ParseResponse], [shared.ExtractResponse],
// [shared.SplitResponse], [shared.EditResponse], [shared.PipelineResponse],
// [V3Extract] or [shared.ClassifyResponse].
type JobGetResponseAsyncJobResponseResultUnion interface {
	ImplementsJobGetResponseAsyncJobResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*JobGetResponseAsyncJobResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.ParseResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.ExtractResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.SplitResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.EditResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.PipelineResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V3Extract{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.ClassifyResponse{}),
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
	Result     JobGetResponseEnhancedAsyncJobResponseResultUnion `json:"result" api:"nullable"`
	Source     interface{}                                       `json:"source"`
	TotalPages int64                                             `json:"total_pages" api:"nullable"`
	Type       JobGetResponseEnhancedAsyncJobResponseType        `json:"type" api:"nullable"`
	JSON       jobGetResponseEnhancedAsyncJobResponseJSON        `json:"-"`
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
//
// Union satisfied by [shared.ParseResponse], [shared.ExtractResponse],
// [shared.SplitResponse], [shared.EditResponse], [shared.PipelineResponse],
// [V3Extract] or [shared.ClassifyResponse].
type JobGetResponseEnhancedAsyncJobResponseResultUnion interface {
	ImplementsJobGetResponseEnhancedAsyncJobResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*JobGetResponseEnhancedAsyncJobResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.ParseResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.ExtractResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.SplitResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.EditResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.PipelineResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V3Extract{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.ClassifyResponse{}),
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
