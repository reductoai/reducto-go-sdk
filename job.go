// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reducto

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/reductoai/reducto-go-sdk/internal/apijson"
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
	opts = append(r.Options[:], opts...)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("cancel/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Retrieve Parse
func (r *JobService) Get(ctx context.Context, jobID string, opts ...option.RequestOption) (res *JobGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("job/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
	// [[]interface{}], [shared.SplitResponseResult],
	// [JobGetResponseAsyncJobResponseResultPipelineResponseResult].
	Result interface{} `json:"result"`
	// The link to the studio pipeline for the document.
	StudioLink string `json:"studio_link,nullable"`
	// This field can have the runtime type of [shared.ParseUsage],
	// [shared.ExtractResponseUsage].
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
// [JobGetResponseAsyncJobResponseResultPipelineResponse].
func (r JobGetResponseAsyncJobResponseResult) AsUnion() JobGetResponseAsyncJobResponseResultUnion {
	return r.union
}

// Union satisfied by [shared.ParseResponse], [shared.ExtractResponse],
// [shared.SplitResponse], [shared.EditResponse] or
// [JobGetResponseAsyncJobResponseResultPipelineResponse].
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
			Type:       reflect.TypeOf(JobGetResponseAsyncJobResponseResultPipelineResponse{}),
		},
	)
}

type JobGetResponseAsyncJobResponseResultPipelineResponse struct {
	JobID  string                                                     `json:"job_id,required"`
	Result JobGetResponseAsyncJobResponseResultPipelineResponseResult `json:"result,required"`
	Usage  shared.ParseUsage                                          `json:"usage,required"`
	JSON   jobGetResponseAsyncJobResponseResultPipelineResponseJSON   `json:"-"`
}

// jobGetResponseAsyncJobResponseResultPipelineResponseJSON contains the JSON
// metadata for the struct [JobGetResponseAsyncJobResponseResultPipelineResponse]
type jobGetResponseAsyncJobResponseResultPipelineResponseJSON struct {
	JobID       apijson.Field
	Result      apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseAsyncJobResponseResultPipelineResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseAsyncJobResponseResultPipelineResponseJSON) RawJSON() string {
	return r.raw
}

func (r JobGetResponseAsyncJobResponseResultPipelineResponse) ImplementsJobGetResponseAsyncJobResponseResult() {
}

type JobGetResponseAsyncJobResponseResultPipelineResponseResult struct {
	Extract JobGetResponseAsyncJobResponseResultPipelineResponseResultExtractUnion `json:"extract,required,nullable"`
	Parse   shared.ParseResponse                                                   `json:"parse,required,nullable"`
	Split   shared.SplitResponse                                                   `json:"split,required,nullable"`
	JSON    jobGetResponseAsyncJobResponseResultPipelineResponseResultJSON         `json:"-"`
}

// jobGetResponseAsyncJobResponseResultPipelineResponseResultJSON contains the JSON
// metadata for the struct
// [JobGetResponseAsyncJobResponseResultPipelineResponseResult]
type jobGetResponseAsyncJobResponseResultPipelineResponseResultJSON struct {
	Extract     apijson.Field
	Parse       apijson.Field
	Split       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseAsyncJobResponseResultPipelineResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseAsyncJobResponseResultPipelineResponseResultJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [JobGetResponseAsyncJobResponseResultPipelineResponseResultExtractArray] or
// [shared.ExtractResponse].
type JobGetResponseAsyncJobResponseResultPipelineResponseResultExtractUnion interface {
	ImplementsJobGetResponseAsyncJobResponseResultPipelineResponseResultExtractUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*JobGetResponseAsyncJobResponseResultPipelineResponseResultExtractUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(JobGetResponseAsyncJobResponseResultPipelineResponseResultExtractArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.ExtractResponse{}),
		},
	)
}

type JobGetResponseAsyncJobResponseResultPipelineResponseResultExtractArray []JobGetResponseAsyncJobResponseResultPipelineResponseResultExtractArrayItem

func (r JobGetResponseAsyncJobResponseResultPipelineResponseResultExtractArray) ImplementsJobGetResponseAsyncJobResponseResultPipelineResponseResultExtractUnion() {
}

// This is the response format for Extract -> Split Pipelines
type JobGetResponseAsyncJobResponseResultPipelineResponseResultExtractArrayItem struct {
	PageRange []int64                                                                        `json:"page_range,required"`
	Result    shared.ExtractResponse                                                         `json:"result,required"`
	SplitName string                                                                         `json:"split_name,required"`
	Partition string                                                                         `json:"partition,nullable"`
	JSON      jobGetResponseAsyncJobResponseResultPipelineResponseResultExtractArrayItemJSON `json:"-"`
}

// jobGetResponseAsyncJobResponseResultPipelineResponseResultExtractArrayItemJSON
// contains the JSON metadata for the struct
// [JobGetResponseAsyncJobResponseResultPipelineResponseResultExtractArrayItem]
type jobGetResponseAsyncJobResponseResultPipelineResponseResultExtractArrayItemJSON struct {
	PageRange   apijson.Field
	Result      apijson.Field
	SplitName   apijson.Field
	Partition   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseAsyncJobResponseResultPipelineResponseResultExtractArrayItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseAsyncJobResponseResultPipelineResponseResultExtractArrayItemJSON) RawJSON() string {
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
	// [[]interface{}], [shared.SplitResponseResult],
	// [JobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResult].
	Result interface{} `json:"result"`
	// The link to the studio pipeline for the document.
	StudioLink string `json:"studio_link,nullable"`
	// This field can have the runtime type of [shared.ParseUsage],
	// [shared.ExtractResponseUsage].
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
// [JobGetResponseEnhancedAsyncJobResponseResultPipelineResponse].
func (r JobGetResponseEnhancedAsyncJobResponseResult) AsUnion() JobGetResponseEnhancedAsyncJobResponseResultUnion {
	return r.union
}

// Union satisfied by [shared.ParseResponse], [shared.ExtractResponse],
// [shared.SplitResponse], [shared.EditResponse] or
// [JobGetResponseEnhancedAsyncJobResponseResultPipelineResponse].
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
			Type:       reflect.TypeOf(JobGetResponseEnhancedAsyncJobResponseResultPipelineResponse{}),
		},
	)
}

type JobGetResponseEnhancedAsyncJobResponseResultPipelineResponse struct {
	JobID  string                                                             `json:"job_id,required"`
	Result JobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResult `json:"result,required"`
	Usage  shared.ParseUsage                                                  `json:"usage,required"`
	JSON   jobGetResponseEnhancedAsyncJobResponseResultPipelineResponseJSON   `json:"-"`
}

// jobGetResponseEnhancedAsyncJobResponseResultPipelineResponseJSON contains the
// JSON metadata for the struct
// [JobGetResponseEnhancedAsyncJobResponseResultPipelineResponse]
type jobGetResponseEnhancedAsyncJobResponseResultPipelineResponseJSON struct {
	JobID       apijson.Field
	Result      apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseEnhancedAsyncJobResponseResultPipelineResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseEnhancedAsyncJobResponseResultPipelineResponseJSON) RawJSON() string {
	return r.raw
}

func (r JobGetResponseEnhancedAsyncJobResponseResultPipelineResponse) ImplementsJobGetResponseEnhancedAsyncJobResponseResult() {
}

type JobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResult struct {
	Extract JobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultExtractUnion `json:"extract,required,nullable"`
	Parse   shared.ParseResponse                                                           `json:"parse,required,nullable"`
	Split   shared.SplitResponse                                                           `json:"split,required,nullable"`
	JSON    jobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultJSON         `json:"-"`
}

// jobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultJSON contains
// the JSON metadata for the struct
// [JobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResult]
type jobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultJSON struct {
	Extract     apijson.Field
	Parse       apijson.Field
	Split       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [JobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultExtractArray]
// or [shared.ExtractResponse].
type JobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultExtractUnion interface {
	ImplementsJobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultExtractUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*JobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultExtractUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(JobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultExtractArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.ExtractResponse{}),
		},
	)
}

type JobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultExtractArray []JobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultExtractArrayItem

func (r JobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultExtractArray) ImplementsJobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultExtractUnion() {
}

// This is the response format for Extract -> Split Pipelines
type JobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultExtractArrayItem struct {
	PageRange []int64                                                                                `json:"page_range,required"`
	Result    shared.ExtractResponse                                                                 `json:"result,required"`
	SplitName string                                                                                 `json:"split_name,required"`
	Partition string                                                                                 `json:"partition,nullable"`
	JSON      jobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultExtractArrayItemJSON `json:"-"`
}

// jobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultExtractArrayItemJSON
// contains the JSON metadata for the struct
// [JobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultExtractArrayItem]
type jobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultExtractArrayItemJSON struct {
	PageRange   apijson.Field
	Result      apijson.Field
	SplitName   apijson.Field
	Partition   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultExtractArrayItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseEnhancedAsyncJobResponseResultPipelineResponseResultExtractArrayItemJSON) RawJSON() string {
	return r.raw
}

type JobGetResponseEnhancedAsyncJobResponseType string

const (
	JobGetResponseEnhancedAsyncJobResponseTypeParse   JobGetResponseEnhancedAsyncJobResponseType = "Parse"
	JobGetResponseEnhancedAsyncJobResponseTypeExtract JobGetResponseEnhancedAsyncJobResponseType = "Extract"
	JobGetResponseEnhancedAsyncJobResponseTypeSplit   JobGetResponseEnhancedAsyncJobResponseType = "Split"
	JobGetResponseEnhancedAsyncJobResponseTypeEdit    JobGetResponseEnhancedAsyncJobResponseType = "Edit"
)

func (r JobGetResponseEnhancedAsyncJobResponseType) IsKnown() bool {
	switch r {
	case JobGetResponseEnhancedAsyncJobResponseTypeParse, JobGetResponseEnhancedAsyncJobResponseTypeExtract, JobGetResponseEnhancedAsyncJobResponseTypeSplit, JobGetResponseEnhancedAsyncJobResponseTypeEdit:
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
	JobGetResponseTypeParse   JobGetResponseType = "Parse"
	JobGetResponseTypeExtract JobGetResponseType = "Extract"
	JobGetResponseTypeSplit   JobGetResponseType = "Split"
	JobGetResponseTypeEdit    JobGetResponseType = "Edit"
)

func (r JobGetResponseType) IsKnown() bool {
	switch r {
	case JobGetResponseTypeParse, JobGetResponseTypeExtract, JobGetResponseTypeSplit, JobGetResponseTypeEdit:
		return true
	}
	return false
}
