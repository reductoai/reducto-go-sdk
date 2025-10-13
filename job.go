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
		return
	}
	path := fmt.Sprintf("cancel/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Retrieve Parse
func (r *JobService) Get(ctx context.Context, jobID string, query JobGetParams, opts ...option.RequestOption) (res *JobGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("job/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Jobs
func (r *JobService) GetAll(ctx context.Context, query JobGetAllParams, opts ...option.RequestOption) (res *JobGetAllResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type JobCancelResponse = interface{}

type JobGetResponse struct {
	Status JobGetResponseStatus `json:"status,required"`
	// This field can have the runtime type of [interface{}].
	Bucket    interface{} `json:"bucket"`
	CreatedAt time.Time   `json:"created_at,nullable" format:"date-time"`
	Duration  float64     `json:"duration,nullable"`
	NumPages  int64       `json:"num_pages,nullable"`
	Progress  float64     `json:"progress,nullable"`
	RawConfig string      `json:"raw_config,nullable"`
	Reason    string      `json:"reason,nullable"`
	// This field can have the runtime type of [JobGetResponseAsyncJobResponseResult],
	// [JobGetResponseEnhancedAsyncJobResponseResult].
	Result interface{} `json:"result"`
	// This field can have the runtime type of [interface{}].
	Source     interface{}        `json:"source"`
	TotalPages int64              `json:"total_pages,nullable"`
	Type       JobGetResponseType `json:"type,nullable"`
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
	Status   JobGetResponseAsyncJobResponseStatus `json:"status,required"`
	Progress float64                              `json:"progress,nullable"`
	Reason   string                               `json:"reason,nullable"`
	Result   JobGetResponseAsyncJobResponseResult `json:"result,nullable"`
	JSON     jobGetResponseAsyncJobResponseJSON   `json:"-"`
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

type JobGetResponseAsyncJobResponseResult struct {
	// This field can have the runtime type of [[]interface{}].
	Citations interface{} `json:"citations"`
	// Presigned URL to download the edited document.
	DocumentURL string `json:"document_url"`
	// The duration of the parse request in seconds.
	Duration float64 `json:"duration"`
	// This field can have the runtime type of [[]shared.EditResponseFormSchema].
	FormSchema interface{} `json:"form_schema"`
	JobID      string      `json:"job_id,nullable"`
	// The storage URL of the converted PDF file.
	PdfURL string `json:"pdf_url,nullable"`
	// This field can have the runtime type of [shared.ParseResponseResult],
	// [[]interface{}], [shared.SplitResponseResult], [shared.PipelineResponseResult].
	Result interface{} `json:"result"`
	// The link to the studio pipeline for the document.
	StudioLink string `json:"studio_link,nullable"`
	// This field can have the runtime type of [shared.ParseUsage],
	// [shared.ExtractResponseUsage],
	// [JobGetResponseAsyncJobResponseResultV3ExtractResponseUsage].
	Usage interface{}                              `json:"usage"`
	JSON  jobGetResponseAsyncJobResponseResultJSON `json:"-"`
	union JobGetResponseAsyncJobResponseResultUnion
}

// jobGetResponseAsyncJobResponseResultJSON contains the JSON metadata for the
// struct [JobGetResponseAsyncJobResponseResult]
type jobGetResponseAsyncJobResponseResultJSON struct {
	Citations   apijson.Field
	DocumentURL apijson.Field
	Duration    apijson.Field
	FormSchema  apijson.Field
	JobID       apijson.Field
	PdfURL      apijson.Field
	Result      apijson.Field
	StudioLink  apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
// Possible runtime types of the union are [shared.ParseResponse],
// [shared.ExtractResponse], [shared.SplitResponse], [shared.EditResponse],
// [shared.PipelineResponse],
// [JobGetResponseAsyncJobResponseResultV3ExtractResponse].
func (r JobGetResponseAsyncJobResponseResult) AsUnion() JobGetResponseAsyncJobResponseResultUnion {
	return r.union
}

// Union satisfied by [shared.ParseResponse], [shared.ExtractResponse],
// [shared.SplitResponse], [shared.EditResponse], [shared.PipelineResponse] or
// [JobGetResponseAsyncJobResponseResultV3ExtractResponse].
type JobGetResponseAsyncJobResponseResultUnion interface {
	ImplementsJobGetResponseAsyncJobResponseResult()
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
			Type:       reflect.TypeOf(JobGetResponseAsyncJobResponseResultV3ExtractResponse{}),
		},
	)
}

type JobGetResponseAsyncJobResponseResultV3ExtractResponse struct {
	// The extracted response in your provided schema. This is a list of dictionaries.
	// If disable_chunking is True (default), then it will be a list of length one.
	Result []interface{}                                              `json:"result,required"`
	Usage  JobGetResponseAsyncJobResponseResultV3ExtractResponseUsage `json:"usage,required"`
	JobID  string                                                     `json:"job_id,nullable"`
	// The link to the studio pipeline for the document.
	StudioLink string                                                    `json:"studio_link,nullable"`
	JSON       jobGetResponseAsyncJobResponseResultV3ExtractResponseJSON `json:"-"`
}

// jobGetResponseAsyncJobResponseResultV3ExtractResponseJSON contains the JSON
// metadata for the struct [JobGetResponseAsyncJobResponseResultV3ExtractResponse]
type jobGetResponseAsyncJobResponseResultV3ExtractResponseJSON struct {
	Result      apijson.Field
	Usage       apijson.Field
	JobID       apijson.Field
	StudioLink  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseAsyncJobResponseResultV3ExtractResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseAsyncJobResponseResultV3ExtractResponseJSON) RawJSON() string {
	return r.raw
}

func (r JobGetResponseAsyncJobResponseResultV3ExtractResponse) ImplementsJobGetResponseAsyncJobResponseResult() {
}

type JobGetResponseAsyncJobResponseResultV3ExtractResponseUsage struct {
	NumFields int64                                                          `json:"num_fields,required"`
	NumPages  int64                                                          `json:"num_pages,required"`
	Credits   float64                                                        `json:"credits,nullable"`
	JSON      jobGetResponseAsyncJobResponseResultV3ExtractResponseUsageJSON `json:"-"`
}

// jobGetResponseAsyncJobResponseResultV3ExtractResponseUsageJSON contains the JSON
// metadata for the struct
// [JobGetResponseAsyncJobResponseResultV3ExtractResponseUsage]
type jobGetResponseAsyncJobResponseResultV3ExtractResponseUsageJSON struct {
	NumFields   apijson.Field
	NumPages    apijson.Field
	Credits     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseAsyncJobResponseResultV3ExtractResponseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseAsyncJobResponseResultV3ExtractResponseUsageJSON) RawJSON() string {
	return r.raw
}

type JobGetResponseEnhancedAsyncJobResponse struct {
	Status     JobGetResponseEnhancedAsyncJobResponseStatus `json:"status,required"`
	Bucket     interface{}                                  `json:"bucket"`
	CreatedAt  time.Time                                    `json:"created_at,nullable" format:"date-time"`
	Duration   float64                                      `json:"duration,nullable"`
	NumPages   int64                                        `json:"num_pages,nullable"`
	Progress   float64                                      `json:"progress,nullable"`
	RawConfig  string                                       `json:"raw_config,nullable"`
	Reason     string                                       `json:"reason,nullable"`
	Result     JobGetResponseEnhancedAsyncJobResponseResult `json:"result,nullable"`
	Source     interface{}                                  `json:"source"`
	TotalPages int64                                        `json:"total_pages,nullable"`
	Type       JobGetResponseEnhancedAsyncJobResponseType   `json:"type,nullable"`
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

type JobGetResponseEnhancedAsyncJobResponseResult struct {
	// This field can have the runtime type of [[]interface{}].
	Citations interface{} `json:"citations"`
	// Presigned URL to download the edited document.
	DocumentURL string `json:"document_url"`
	// The duration of the parse request in seconds.
	Duration float64 `json:"duration"`
	// This field can have the runtime type of [[]shared.EditResponseFormSchema].
	FormSchema interface{} `json:"form_schema"`
	JobID      string      `json:"job_id,nullable"`
	// The storage URL of the converted PDF file.
	PdfURL string `json:"pdf_url,nullable"`
	// This field can have the runtime type of [shared.ParseResponseResult],
	// [[]interface{}], [shared.SplitResponseResult], [shared.PipelineResponseResult].
	Result interface{} `json:"result"`
	// The link to the studio pipeline for the document.
	StudioLink string `json:"studio_link,nullable"`
	// This field can have the runtime type of [shared.ParseUsage],
	// [shared.ExtractResponseUsage],
	// [JobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponseUsage].
	Usage interface{}                                      `json:"usage"`
	JSON  jobGetResponseEnhancedAsyncJobResponseResultJSON `json:"-"`
	union JobGetResponseEnhancedAsyncJobResponseResultUnion
}

// jobGetResponseEnhancedAsyncJobResponseResultJSON contains the JSON metadata for
// the struct [JobGetResponseEnhancedAsyncJobResponseResult]
type jobGetResponseEnhancedAsyncJobResponseResultJSON struct {
	Citations   apijson.Field
	DocumentURL apijson.Field
	Duration    apijson.Field
	FormSchema  apijson.Field
	JobID       apijson.Field
	PdfURL      apijson.Field
	Result      apijson.Field
	StudioLink  apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
// Possible runtime types of the union are [shared.ParseResponse],
// [shared.ExtractResponse], [shared.SplitResponse], [shared.EditResponse],
// [shared.PipelineResponse],
// [JobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponse].
func (r JobGetResponseEnhancedAsyncJobResponseResult) AsUnion() JobGetResponseEnhancedAsyncJobResponseResultUnion {
	return r.union
}

// Union satisfied by [shared.ParseResponse], [shared.ExtractResponse],
// [shared.SplitResponse], [shared.EditResponse], [shared.PipelineResponse] or
// [JobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponse].
type JobGetResponseEnhancedAsyncJobResponseResultUnion interface {
	ImplementsJobGetResponseEnhancedAsyncJobResponseResult()
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
			Type:       reflect.TypeOf(JobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponse{}),
		},
	)
}

type JobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponse struct {
	// The extracted response in your provided schema. This is a list of dictionaries.
	// If disable_chunking is True (default), then it will be a list of length one.
	Result []interface{}                                                      `json:"result,required"`
	Usage  JobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponseUsage `json:"usage,required"`
	JobID  string                                                             `json:"job_id,nullable"`
	// The link to the studio pipeline for the document.
	StudioLink string                                                            `json:"studio_link,nullable"`
	JSON       jobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponseJSON `json:"-"`
}

// jobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponseJSON contains the
// JSON metadata for the struct
// [JobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponse]
type jobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponseJSON struct {
	Result      apijson.Field
	Usage       apijson.Field
	JobID       apijson.Field
	StudioLink  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponseJSON) RawJSON() string {
	return r.raw
}

func (r JobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponse) ImplementsJobGetResponseEnhancedAsyncJobResponseResult() {
}

type JobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponseUsage struct {
	NumFields int64                                                                  `json:"num_fields,required"`
	NumPages  int64                                                                  `json:"num_pages,required"`
	Credits   float64                                                                `json:"credits,nullable"`
	JSON      jobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponseUsageJSON `json:"-"`
}

// jobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponseUsageJSON contains
// the JSON metadata for the struct
// [JobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponseUsage]
type jobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponseUsageJSON struct {
	NumFields   apijson.Field
	NumPages    apijson.Field
	Credits     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseEnhancedAsyncJobResponseResultV3ExtractResponseUsageJSON) RawJSON() string {
	return r.raw
}

type JobGetResponseEnhancedAsyncJobResponseType string

const (
	JobGetResponseEnhancedAsyncJobResponseTypeParse    JobGetResponseEnhancedAsyncJobResponseType = "Parse"
	JobGetResponseEnhancedAsyncJobResponseTypeExtract  JobGetResponseEnhancedAsyncJobResponseType = "Extract"
	JobGetResponseEnhancedAsyncJobResponseTypeSplit    JobGetResponseEnhancedAsyncJobResponseType = "Split"
	JobGetResponseEnhancedAsyncJobResponseTypeEdit     JobGetResponseEnhancedAsyncJobResponseType = "Edit"
	JobGetResponseEnhancedAsyncJobResponseTypePipeline JobGetResponseEnhancedAsyncJobResponseType = "Pipeline"
)

func (r JobGetResponseEnhancedAsyncJobResponseType) IsKnown() bool {
	switch r {
	case JobGetResponseEnhancedAsyncJobResponseTypeParse, JobGetResponseEnhancedAsyncJobResponseTypeExtract, JobGetResponseEnhancedAsyncJobResponseTypeSplit, JobGetResponseEnhancedAsyncJobResponseTypeEdit, JobGetResponseEnhancedAsyncJobResponseTypePipeline:
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
)

func (r JobGetResponseType) IsKnown() bool {
	switch r {
	case JobGetResponseTypeParse, JobGetResponseTypeExtract, JobGetResponseTypeSplit, JobGetResponseTypeEdit, JobGetResponseTypePipeline:
		return true
	}
	return false
}

type JobGetAllResponse struct {
	// List of jobs with their job_id, status, type, raw_config, created_at, num_pages
	// and duration
	Jobs []JobGetAllResponseJob `json:"jobs,required"`
	// Cursor to fetch the next page of results. If null, there are no more results.
	NextCursor string                `json:"next_cursor,nullable"`
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
	CreatedAt  time.Time                   `json:"created_at,required" format:"date-time"`
	Duration   float64                     `json:"duration,required,nullable"`
	JobID      string                      `json:"job_id,required"`
	NumPages   int64                       `json:"num_pages,required,nullable"`
	RawConfig  string                      `json:"raw_config,required"`
	Status     JobGetAllResponseJobsStatus `json:"status,required"`
	TotalPages int64                       `json:"total_pages,required,nullable"`
	Type       JobGetAllResponseJobsType   `json:"type,required"`
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
)

func (r JobGetAllResponseJobsType) IsKnown() bool {
	switch r {
	case JobGetAllResponseJobsTypeParse, JobGetAllResponseJobsTypeExtract, JobGetAllResponseJobsTypeSplit, JobGetAllResponseJobsTypeEdit, JobGetAllResponseJobsTypePipeline:
		return true
	}
	return false
}

type JobGetParams struct {
	BucketName param.Field[string] `query:"bucket_name"`
}

// URLQuery serializes [JobGetParams]'s query parameters as `url.Values`.
func (r JobGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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
