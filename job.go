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
	// This field can have the runtime type of [JobGetResponseAsyncJobResponseError],
	// [JobGetResponseEnhancedAsyncJobResponseError].
	Error     interface{} `json:"error"`
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
	Error       apijson.Field
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
	Status JobGetResponseAsyncJobResponseStatus `json:"status" api:"required"`
	// Structured error body returned to customers.
	//
	// Matches the format specified in `ERROR_POLICY.md`.
	Error    JobGetResponseAsyncJobResponseError `json:"error" api:"nullable"`
	Progress float64                             `json:"progress" api:"nullable"`
	Reason   string                              `json:"reason" api:"nullable"`
	// Response from classify job - returned when polling /job/{job_id}
	Result JobGetResponseAsyncJobResponseResult `json:"result" api:"nullable"`
	JSON   jobGetResponseAsyncJobResponseJSON   `json:"-"`
}

// jobGetResponseAsyncJobResponseJSON contains the JSON metadata for the struct
// [JobGetResponseAsyncJobResponse]
type jobGetResponseAsyncJobResponseJSON struct {
	Status      apijson.Field
	Error       apijson.Field
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

// Structured error body returned to customers.
//
// Matches the format specified in `ERROR_POLICY.md`.
type JobGetResponseAsyncJobResponseError struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Machine-readable error names returned in API error responses.
	//
	// Each member maps to a category (Transient / Processing / Input) and a default
	// HTTP status code defined in `ERROR_CODE_DEFAULTS`. The enum value is the string
	// customers see in the `error.name` field.
	Name  JobGetResponseAsyncJobResponseErrorName `json:"name" api:"required"`
	JobID string                                  `json:"job_id" api:"nullable"`
	JSON  jobGetResponseAsyncJobResponseErrorJSON `json:"-"`
}

// jobGetResponseAsyncJobResponseErrorJSON contains the JSON metadata for the
// struct [JobGetResponseAsyncJobResponseError]
type jobGetResponseAsyncJobResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Name        apijson.Field
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseAsyncJobResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseAsyncJobResponseErrorJSON) RawJSON() string {
	return r.raw
}

// Machine-readable error names returned in API error responses.
//
// Each member maps to a category (Transient / Processing / Input) and a default
// HTTP status code defined in `ERROR_CODE_DEFAULTS`. The enum value is the string
// customers see in the `error.name` field.
type JobGetResponseAsyncJobResponseErrorName string

const (
	JobGetResponseAsyncJobResponseErrorNameTimeout                    JobGetResponseAsyncJobResponseErrorName = "TIMEOUT"
	JobGetResponseAsyncJobResponseErrorNameCapacityTimeout            JobGetResponseAsyncJobResponseErrorName = "CAPACITY_TIMEOUT"
	JobGetResponseAsyncJobResponseErrorNameCustomerTimeout            JobGetResponseAsyncJobResponseErrorName = "CUSTOMER_TIMEOUT"
	JobGetResponseAsyncJobResponseErrorNameInternalError              JobGetResponseAsyncJobResponseErrorName = "INTERNAL_ERROR"
	JobGetResponseAsyncJobResponseErrorNameServiceUnavailable         JobGetResponseAsyncJobResponseErrorName = "SERVICE_UNAVAILABLE"
	JobGetResponseAsyncJobResponseErrorNameGPUAllocationError         JobGetResponseAsyncJobResponseErrorName = "GPU_ALLOCATION_ERROR"
	JobGetResponseAsyncJobResponseErrorNameGPUPoolSaturated           JobGetResponseAsyncJobResponseErrorName = "GPU_POOL_SATURATED"
	JobGetResponseAsyncJobResponseErrorNameBatchQueueFull             JobGetResponseAsyncJobResponseErrorName = "BATCH_QUEUE_FULL"
	JobGetResponseAsyncJobResponseErrorNameJobStateError              JobGetResponseAsyncJobResponseErrorName = "JOB_STATE_ERROR"
	JobGetResponseAsyncJobResponseErrorNameDocumentCorrupt            JobGetResponseAsyncJobResponseErrorName = "DOCUMENT_CORRUPT"
	JobGetResponseAsyncJobResponseErrorNameDocumentEmpty              JobGetResponseAsyncJobResponseErrorName = "DOCUMENT_EMPTY"
	JobGetResponseAsyncJobResponseErrorNameDocumentUnsupported        JobGetResponseAsyncJobResponseErrorName = "DOCUMENT_UNSUPPORTED"
	JobGetResponseAsyncJobResponseErrorNameDocumentTooLarge           JobGetResponseAsyncJobResponseErrorName = "DOCUMENT_TOO_LARGE"
	JobGetResponseAsyncJobResponseErrorNameOfficeConversionTooLarge   JobGetResponseAsyncJobResponseErrorName = "OFFICE_CONVERSION_TOO_LARGE"
	JobGetResponseAsyncJobResponseErrorNameImageTooLarge              JobGetResponseAsyncJobResponseErrorName = "IMAGE_TOO_LARGE"
	JobGetResponseAsyncJobResponseErrorNameImageTooSmall              JobGetResponseAsyncJobResponseErrorName = "IMAGE_TOO_SMALL"
	JobGetResponseAsyncJobResponseErrorNameImageInvalidAspectRatio    JobGetResponseAsyncJobResponseErrorName = "IMAGE_INVALID_ASPECT_RATIO"
	JobGetResponseAsyncJobResponseErrorNameDocumentPasswordProtected  JobGetResponseAsyncJobResponseErrorName = "DOCUMENT_PASSWORD_PROTECTED"
	JobGetResponseAsyncJobResponseErrorNameFormFillFailed             JobGetResponseAsyncJobResponseErrorName = "FORM_FILL_FAILED"
	JobGetResponseAsyncJobResponseErrorNameInternalInvariantViolation JobGetResponseAsyncJobResponseErrorName = "INTERNAL_INVARIANT_VIOLATION"
	JobGetResponseAsyncJobResponseErrorNameGPUUnavailable             JobGetResponseAsyncJobResponseErrorName = "GPU_UNAVAILABLE"
	JobGetResponseAsyncJobResponseErrorNameContextWindowExceeded      JobGetResponseAsyncJobResponseErrorName = "CONTEXT_WINDOW_EXCEEDED"
	JobGetResponseAsyncJobResponseErrorNameProcessingFailed           JobGetResponseAsyncJobResponseErrorName = "PROCESSING_FAILED"
	JobGetResponseAsyncJobResponseErrorNameInferenceMethodUnsupported JobGetResponseAsyncJobResponseErrorName = "INFERENCE_METHOD_UNSUPPORTED"
	JobGetResponseAsyncJobResponseErrorNameSubprocessCrashed          JobGetResponseAsyncJobResponseErrorName = "SUBPROCESS_CRASHED"
	JobGetResponseAsyncJobResponseErrorNameBatchOrphaned              JobGetResponseAsyncJobResponseErrorName = "BATCH_ORPHANED"
	JobGetResponseAsyncJobResponseErrorNameOversizedResult            JobGetResponseAsyncJobResponseErrorName = "OVERSIZED_RESULT"
	JobGetResponseAsyncJobResponseErrorNameLlmOutputParseFailed       JobGetResponseAsyncJobResponseErrorName = "LLM_OUTPUT_PARSE_FAILED"
	JobGetResponseAsyncJobResponseErrorNameLlmProviderError           JobGetResponseAsyncJobResponseErrorName = "LLM_PROVIDER_ERROR"
	JobGetResponseAsyncJobResponseErrorNameInvalidConfig              JobGetResponseAsyncJobResponseErrorName = "INVALID_CONFIG"
	JobGetResponseAsyncJobResponseErrorNameInvalidSchema              JobGetResponseAsyncJobResponseErrorName = "INVALID_SCHEMA"
	JobGetResponseAsyncJobResponseErrorNameAuthError                  JobGetResponseAsyncJobResponseErrorName = "AUTH_ERROR"
	JobGetResponseAsyncJobResponseErrorNameNotApplicable              JobGetResponseAsyncJobResponseErrorName = "NOT_APPLICABLE"
	JobGetResponseAsyncJobResponseErrorNameRegionUnavailable          JobGetResponseAsyncJobResponseErrorName = "REGION_UNAVAILABLE"
	JobGetResponseAsyncJobResponseErrorNameNotFound                   JobGetResponseAsyncJobResponseErrorName = "NOT_FOUND"
	JobGetResponseAsyncJobResponseErrorNameJobDeletionInProgress      JobGetResponseAsyncJobResponseErrorName = "JOB_DELETION_IN_PROGRESS"
	JobGetResponseAsyncJobResponseErrorNameJobDeleted                 JobGetResponseAsyncJobResponseErrorName = "JOB_DELETED"
	JobGetResponseAsyncJobResponseErrorNameJobNotComplete             JobGetResponseAsyncJobResponseErrorName = "JOB_NOT_COMPLETE"
	JobGetResponseAsyncJobResponseErrorNameJobCancelled               JobGetResponseAsyncJobResponseErrorName = "JOB_CANCELLED"
	JobGetResponseAsyncJobResponseErrorNameRateLimit                  JobGetResponseAsyncJobResponseErrorName = "RATE_LIMIT"
	JobGetResponseAsyncJobResponseErrorNameCellCountExceeded          JobGetResponseAsyncJobResponseErrorName = "CELL_COUNT_EXCEEDED"
	JobGetResponseAsyncJobResponseErrorNameURLNotAllowed              JobGetResponseAsyncJobResponseErrorName = "URL_NOT_ALLOWED"
)

func (r JobGetResponseAsyncJobResponseErrorName) IsKnown() bool {
	switch r {
	case JobGetResponseAsyncJobResponseErrorNameTimeout, JobGetResponseAsyncJobResponseErrorNameCapacityTimeout, JobGetResponseAsyncJobResponseErrorNameCustomerTimeout, JobGetResponseAsyncJobResponseErrorNameInternalError, JobGetResponseAsyncJobResponseErrorNameServiceUnavailable, JobGetResponseAsyncJobResponseErrorNameGPUAllocationError, JobGetResponseAsyncJobResponseErrorNameGPUPoolSaturated, JobGetResponseAsyncJobResponseErrorNameBatchQueueFull, JobGetResponseAsyncJobResponseErrorNameJobStateError, JobGetResponseAsyncJobResponseErrorNameDocumentCorrupt, JobGetResponseAsyncJobResponseErrorNameDocumentEmpty, JobGetResponseAsyncJobResponseErrorNameDocumentUnsupported, JobGetResponseAsyncJobResponseErrorNameDocumentTooLarge, JobGetResponseAsyncJobResponseErrorNameOfficeConversionTooLarge, JobGetResponseAsyncJobResponseErrorNameImageTooLarge, JobGetResponseAsyncJobResponseErrorNameImageTooSmall, JobGetResponseAsyncJobResponseErrorNameImageInvalidAspectRatio, JobGetResponseAsyncJobResponseErrorNameDocumentPasswordProtected, JobGetResponseAsyncJobResponseErrorNameFormFillFailed, JobGetResponseAsyncJobResponseErrorNameInternalInvariantViolation, JobGetResponseAsyncJobResponseErrorNameGPUUnavailable, JobGetResponseAsyncJobResponseErrorNameContextWindowExceeded, JobGetResponseAsyncJobResponseErrorNameProcessingFailed, JobGetResponseAsyncJobResponseErrorNameInferenceMethodUnsupported, JobGetResponseAsyncJobResponseErrorNameSubprocessCrashed, JobGetResponseAsyncJobResponseErrorNameBatchOrphaned, JobGetResponseAsyncJobResponseErrorNameOversizedResult, JobGetResponseAsyncJobResponseErrorNameLlmOutputParseFailed, JobGetResponseAsyncJobResponseErrorNameLlmProviderError, JobGetResponseAsyncJobResponseErrorNameInvalidConfig, JobGetResponseAsyncJobResponseErrorNameInvalidSchema, JobGetResponseAsyncJobResponseErrorNameAuthError, JobGetResponseAsyncJobResponseErrorNameNotApplicable, JobGetResponseAsyncJobResponseErrorNameRegionUnavailable, JobGetResponseAsyncJobResponseErrorNameNotFound, JobGetResponseAsyncJobResponseErrorNameJobDeletionInProgress, JobGetResponseAsyncJobResponseErrorNameJobDeleted, JobGetResponseAsyncJobResponseErrorNameJobNotComplete, JobGetResponseAsyncJobResponseErrorNameJobCancelled, JobGetResponseAsyncJobResponseErrorNameRateLimit, JobGetResponseAsyncJobResponseErrorNameCellCountExceeded, JobGetResponseAsyncJobResponseErrorNameURLNotAllowed:
		return true
	}
	return false
}

// Response from classify job - returned when polling /job/{job_id}
type JobGetResponseAsyncJobResponseResult struct {
	// This field can have the runtime type of [shared.ExtractResponseCitationsUnion].
	Citations interface{} `json:"citations"`
	// Optional document-level deep extract confidence label.
	Confidence JobGetResponseAsyncJobResponseResultConfidence `json:"confidence" api:"nullable"`
	// Optional explanation for the document-level confidence label.
	ConfidenceReason string `json:"confidence_reason" api:"nullable"`
	// This field can have the runtime type of
	// [shared.ParseResponseDocumentProperties].
	DocumentProperties interface{} `json:"document_properties"`
	// Presigned URL to download the edited document.
	DocumentURL string `json:"document_url"`
	// The duration of the parse request in seconds.
	Duration float64 `json:"duration" api:"nullable"`
	// This field can have the runtime type of [map[string]string].
	ExtraMetadata interface{} `json:"extra_metadata"`
	// This field can have the runtime type of [[]EditWidget].
	FormSchema interface{} `json:"form_schema"`
	JobID      string      `json:"job_id" api:"nullable"`
	// The storage URL of the converted PDF file.
	PdfURL string `json:"pdf_url" api:"nullable"`
	// This field can have the runtime type of [map[string]interface{}],
	// [shared.ClassifyResponseResponseConfidence].
	ResponseConfidence interface{}                                      `json:"response_confidence"`
	ResponseType       JobGetResponseAsyncJobResponseResultResponseType `json:"response_type"`
	// This field can have the runtime type of [shared.ParseResponseResult],
	// [shared.ExtractResponseResultUnion], [shared.SplitResponseResult],
	// [shared.PipelineResponseResult], [[]interface{}],
	// [shared.ClassifyResponseResult],
	// [JobGetResponseAsyncJobResponseResultChartResponseResult].
	Result interface{} `json:"result"`
	// The link to the studio pipeline for the document.
	StudioLink string `json:"studio_link" api:"nullable"`
	// This field can have the runtime type of [ParseUsage], [ExtractUsage],
	// [shared.ClassifyResponseUsage],
	// [JobGetResponseAsyncJobResponseResultChartResponseUsage].
	Usage interface{}                              `json:"usage"`
	JSON  jobGetResponseAsyncJobResponseResultJSON `json:"-"`
	union JobGetResponseAsyncJobResponseResultUnion
}

// jobGetResponseAsyncJobResponseResultJSON contains the JSON metadata for the
// struct [JobGetResponseAsyncJobResponseResult]
type jobGetResponseAsyncJobResponseResultJSON struct {
	Citations          apijson.Field
	Confidence         apijson.Field
	ConfidenceReason   apijson.Field
	DocumentProperties apijson.Field
	DocumentURL        apijson.Field
	Duration           apijson.Field
	ExtraMetadata      apijson.Field
	FormSchema         apijson.Field
	JobID              apijson.Field
	PdfURL             apijson.Field
	ResponseConfidence apijson.Field
	ResponseType       apijson.Field
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
// Possible runtime types of the union are [shared.ParseResponse],
// [shared.ExtractResponse], [shared.SplitResponse], [shared.EditResponse],
// [shared.PipelineResponse], [V3Extract], [shared.ClassifyResponse],
// [JobGetResponseAsyncJobResponseResultChartResponse].
func (r JobGetResponseAsyncJobResponseResult) AsUnion() JobGetResponseAsyncJobResponseResultUnion {
	return r.union
}

// Response from classify job - returned when polling /job/{job_id}
//
// Union satisfied by [shared.ParseResponse], [shared.ExtractResponse],
// [shared.SplitResponse], [shared.EditResponse], [shared.PipelineResponse],
// [V3Extract], [shared.ClassifyResponse] or
// [JobGetResponseAsyncJobResponseResultChartResponse].
type JobGetResponseAsyncJobResponseResultUnion interface {
	ImplementsJobGetResponseAsyncJobResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*JobGetResponseAsyncJobResponseResultUnion)(nil)).Elem(),
		"response_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.ParseResponse{}),
			DiscriminatorValue: "parse",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.ExtractResponse{}),
			DiscriminatorValue: "extract",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.SplitResponse{}),
			DiscriminatorValue: "split",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.EditResponse{}),
			DiscriminatorValue: "edit",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.PipelineResponse{}),
			DiscriminatorValue: "pipeline",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(V3Extract{}),
			DiscriminatorValue: "v3_extract",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.ClassifyResponse{}),
			DiscriminatorValue: "classify",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(JobGetResponseAsyncJobResponseResultChartResponse{}),
			DiscriminatorValue: "chart",
		},
	)
}

// Response from synchronous and persisted chart extraction jobs.
type JobGetResponseAsyncJobResponseResultChartResponse struct {
	JobID        string                                                        `json:"job_id" api:"required"`
	Result       JobGetResponseAsyncJobResponseResultChartResponseResult       `json:"result" api:"required"`
	Duration     float64                                                       `json:"duration" api:"nullable"`
	ResponseType JobGetResponseAsyncJobResponseResultChartResponseResponseType `json:"response_type"`
	Usage        JobGetResponseAsyncJobResponseResultChartResponseUsage        `json:"usage"`
	JSON         jobGetResponseAsyncJobResponseResultChartResponseJSON         `json:"-"`
}

// jobGetResponseAsyncJobResponseResultChartResponseJSON contains the JSON metadata
// for the struct [JobGetResponseAsyncJobResponseResultChartResponse]
type jobGetResponseAsyncJobResponseResultChartResponseJSON struct {
	JobID        apijson.Field
	Result       apijson.Field
	Duration     apijson.Field
	ResponseType apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *JobGetResponseAsyncJobResponseResultChartResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseAsyncJobResponseResultChartResponseJSON) RawJSON() string {
	return r.raw
}

func (r JobGetResponseAsyncJobResponseResultChartResponse) ImplementsJobGetResponseAsyncJobResponseResult() {
}

type JobGetResponseAsyncJobResponseResultChartResponseResult struct {
	ChartData         map[string]interface{}                                      `json:"chart_data" api:"required"`
	ReconstructionURL string                                                      `json:"reconstruction_url" api:"required"`
	Summary           string                                                      `json:"summary" api:"required"`
	Verified          bool                                                        `json:"verified" api:"required"`
	JSON              jobGetResponseAsyncJobResponseResultChartResponseResultJSON `json:"-"`
}

// jobGetResponseAsyncJobResponseResultChartResponseResultJSON contains the JSON
// metadata for the struct
// [JobGetResponseAsyncJobResponseResultChartResponseResult]
type jobGetResponseAsyncJobResponseResultChartResponseResultJSON struct {
	ChartData         apijson.Field
	ReconstructionURL apijson.Field
	Summary           apijson.Field
	Verified          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *JobGetResponseAsyncJobResponseResultChartResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseAsyncJobResponseResultChartResponseResultJSON) RawJSON() string {
	return r.raw
}

type JobGetResponseAsyncJobResponseResultChartResponseResponseType string

const (
	JobGetResponseAsyncJobResponseResultChartResponseResponseTypeChart JobGetResponseAsyncJobResponseResultChartResponseResponseType = "chart"
)

func (r JobGetResponseAsyncJobResponseResultChartResponseResponseType) IsKnown() bool {
	switch r {
	case JobGetResponseAsyncJobResponseResultChartResponseResponseTypeChart:
		return true
	}
	return false
}

type JobGetResponseAsyncJobResponseResultChartResponseUsage struct {
	Credits   float64                                                         `json:"credits"`
	NumCharts JobGetResponseAsyncJobResponseResultChartResponseUsageNumCharts `json:"num_charts"`
	JSON      jobGetResponseAsyncJobResponseResultChartResponseUsageJSON      `json:"-"`
}

// jobGetResponseAsyncJobResponseResultChartResponseUsageJSON contains the JSON
// metadata for the struct [JobGetResponseAsyncJobResponseResultChartResponseUsage]
type jobGetResponseAsyncJobResponseResultChartResponseUsageJSON struct {
	Credits     apijson.Field
	NumCharts   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseAsyncJobResponseResultChartResponseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseAsyncJobResponseResultChartResponseUsageJSON) RawJSON() string {
	return r.raw
}

type JobGetResponseAsyncJobResponseResultChartResponseUsageNumCharts int64

const (
	JobGetResponseAsyncJobResponseResultChartResponseUsageNumCharts1 JobGetResponseAsyncJobResponseResultChartResponseUsageNumCharts = 1
)

func (r JobGetResponseAsyncJobResponseResultChartResponseUsageNumCharts) IsKnown() bool {
	switch r {
	case JobGetResponseAsyncJobResponseResultChartResponseUsageNumCharts1:
		return true
	}
	return false
}

// Optional document-level deep extract confidence label.
type JobGetResponseAsyncJobResponseResultConfidence string

const (
	JobGetResponseAsyncJobResponseResultConfidenceHigh JobGetResponseAsyncJobResponseResultConfidence = "high"
	JobGetResponseAsyncJobResponseResultConfidenceLow  JobGetResponseAsyncJobResponseResultConfidence = "low"
)

func (r JobGetResponseAsyncJobResponseResultConfidence) IsKnown() bool {
	switch r {
	case JobGetResponseAsyncJobResponseResultConfidenceHigh, JobGetResponseAsyncJobResponseResultConfidenceLow:
		return true
	}
	return false
}

type JobGetResponseAsyncJobResponseResultResponseType string

const (
	JobGetResponseAsyncJobResponseResultResponseTypeParse     JobGetResponseAsyncJobResponseResultResponseType = "parse"
	JobGetResponseAsyncJobResponseResultResponseTypeExtract   JobGetResponseAsyncJobResponseResultResponseType = "extract"
	JobGetResponseAsyncJobResponseResultResponseTypeSplit     JobGetResponseAsyncJobResponseResultResponseType = "split"
	JobGetResponseAsyncJobResponseResultResponseTypeEdit      JobGetResponseAsyncJobResponseResultResponseType = "edit"
	JobGetResponseAsyncJobResponseResultResponseTypePipeline  JobGetResponseAsyncJobResponseResultResponseType = "pipeline"
	JobGetResponseAsyncJobResponseResultResponseTypeV3Extract JobGetResponseAsyncJobResponseResultResponseType = "v3_extract"
	JobGetResponseAsyncJobResponseResultResponseTypeClassify  JobGetResponseAsyncJobResponseResultResponseType = "classify"
	JobGetResponseAsyncJobResponseResultResponseTypeChart     JobGetResponseAsyncJobResponseResultResponseType = "chart"
)

func (r JobGetResponseAsyncJobResponseResultResponseType) IsKnown() bool {
	switch r {
	case JobGetResponseAsyncJobResponseResultResponseTypeParse, JobGetResponseAsyncJobResponseResultResponseTypeExtract, JobGetResponseAsyncJobResponseResultResponseTypeSplit, JobGetResponseAsyncJobResponseResultResponseTypeEdit, JobGetResponseAsyncJobResponseResultResponseTypePipeline, JobGetResponseAsyncJobResponseResultResponseTypeV3Extract, JobGetResponseAsyncJobResponseResultResponseTypeClassify, JobGetResponseAsyncJobResponseResultResponseTypeChart:
		return true
	}
	return false
}

type JobGetResponseEnhancedAsyncJobResponse struct {
	Status    JobGetResponseEnhancedAsyncJobResponseStatus `json:"status" api:"required"`
	Bucket    interface{}                                  `json:"bucket"`
	CreatedAt time.Time                                    `json:"created_at" api:"nullable" format:"date-time"`
	Duration  float64                                      `json:"duration" api:"nullable"`
	// Structured error body returned to customers.
	//
	// Matches the format specified in `ERROR_POLICY.md`.
	Error     JobGetResponseEnhancedAsyncJobResponseError `json:"error" api:"nullable"`
	NumPages  int64                                       `json:"num_pages" api:"nullable"`
	Progress  float64                                     `json:"progress" api:"nullable"`
	RawConfig string                                      `json:"raw_config" api:"nullable"`
	Reason    string                                      `json:"reason" api:"nullable"`
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
	Error       apijson.Field
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

// Structured error body returned to customers.
//
// Matches the format specified in `ERROR_POLICY.md`.
type JobGetResponseEnhancedAsyncJobResponseError struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Machine-readable error names returned in API error responses.
	//
	// Each member maps to a category (Transient / Processing / Input) and a default
	// HTTP status code defined in `ERROR_CODE_DEFAULTS`. The enum value is the string
	// customers see in the `error.name` field.
	Name  JobGetResponseEnhancedAsyncJobResponseErrorName `json:"name" api:"required"`
	JobID string                                          `json:"job_id" api:"nullable"`
	JSON  jobGetResponseEnhancedAsyncJobResponseErrorJSON `json:"-"`
}

// jobGetResponseEnhancedAsyncJobResponseErrorJSON contains the JSON metadata for
// the struct [JobGetResponseEnhancedAsyncJobResponseError]
type jobGetResponseEnhancedAsyncJobResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Name        apijson.Field
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseEnhancedAsyncJobResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseEnhancedAsyncJobResponseErrorJSON) RawJSON() string {
	return r.raw
}

// Machine-readable error names returned in API error responses.
//
// Each member maps to a category (Transient / Processing / Input) and a default
// HTTP status code defined in `ERROR_CODE_DEFAULTS`. The enum value is the string
// customers see in the `error.name` field.
type JobGetResponseEnhancedAsyncJobResponseErrorName string

const (
	JobGetResponseEnhancedAsyncJobResponseErrorNameTimeout                    JobGetResponseEnhancedAsyncJobResponseErrorName = "TIMEOUT"
	JobGetResponseEnhancedAsyncJobResponseErrorNameCapacityTimeout            JobGetResponseEnhancedAsyncJobResponseErrorName = "CAPACITY_TIMEOUT"
	JobGetResponseEnhancedAsyncJobResponseErrorNameCustomerTimeout            JobGetResponseEnhancedAsyncJobResponseErrorName = "CUSTOMER_TIMEOUT"
	JobGetResponseEnhancedAsyncJobResponseErrorNameInternalError              JobGetResponseEnhancedAsyncJobResponseErrorName = "INTERNAL_ERROR"
	JobGetResponseEnhancedAsyncJobResponseErrorNameServiceUnavailable         JobGetResponseEnhancedAsyncJobResponseErrorName = "SERVICE_UNAVAILABLE"
	JobGetResponseEnhancedAsyncJobResponseErrorNameGPUAllocationError         JobGetResponseEnhancedAsyncJobResponseErrorName = "GPU_ALLOCATION_ERROR"
	JobGetResponseEnhancedAsyncJobResponseErrorNameGPUPoolSaturated           JobGetResponseEnhancedAsyncJobResponseErrorName = "GPU_POOL_SATURATED"
	JobGetResponseEnhancedAsyncJobResponseErrorNameBatchQueueFull             JobGetResponseEnhancedAsyncJobResponseErrorName = "BATCH_QUEUE_FULL"
	JobGetResponseEnhancedAsyncJobResponseErrorNameJobStateError              JobGetResponseEnhancedAsyncJobResponseErrorName = "JOB_STATE_ERROR"
	JobGetResponseEnhancedAsyncJobResponseErrorNameDocumentCorrupt            JobGetResponseEnhancedAsyncJobResponseErrorName = "DOCUMENT_CORRUPT"
	JobGetResponseEnhancedAsyncJobResponseErrorNameDocumentEmpty              JobGetResponseEnhancedAsyncJobResponseErrorName = "DOCUMENT_EMPTY"
	JobGetResponseEnhancedAsyncJobResponseErrorNameDocumentUnsupported        JobGetResponseEnhancedAsyncJobResponseErrorName = "DOCUMENT_UNSUPPORTED"
	JobGetResponseEnhancedAsyncJobResponseErrorNameDocumentTooLarge           JobGetResponseEnhancedAsyncJobResponseErrorName = "DOCUMENT_TOO_LARGE"
	JobGetResponseEnhancedAsyncJobResponseErrorNameOfficeConversionTooLarge   JobGetResponseEnhancedAsyncJobResponseErrorName = "OFFICE_CONVERSION_TOO_LARGE"
	JobGetResponseEnhancedAsyncJobResponseErrorNameImageTooLarge              JobGetResponseEnhancedAsyncJobResponseErrorName = "IMAGE_TOO_LARGE"
	JobGetResponseEnhancedAsyncJobResponseErrorNameImageTooSmall              JobGetResponseEnhancedAsyncJobResponseErrorName = "IMAGE_TOO_SMALL"
	JobGetResponseEnhancedAsyncJobResponseErrorNameImageInvalidAspectRatio    JobGetResponseEnhancedAsyncJobResponseErrorName = "IMAGE_INVALID_ASPECT_RATIO"
	JobGetResponseEnhancedAsyncJobResponseErrorNameDocumentPasswordProtected  JobGetResponseEnhancedAsyncJobResponseErrorName = "DOCUMENT_PASSWORD_PROTECTED"
	JobGetResponseEnhancedAsyncJobResponseErrorNameFormFillFailed             JobGetResponseEnhancedAsyncJobResponseErrorName = "FORM_FILL_FAILED"
	JobGetResponseEnhancedAsyncJobResponseErrorNameInternalInvariantViolation JobGetResponseEnhancedAsyncJobResponseErrorName = "INTERNAL_INVARIANT_VIOLATION"
	JobGetResponseEnhancedAsyncJobResponseErrorNameGPUUnavailable             JobGetResponseEnhancedAsyncJobResponseErrorName = "GPU_UNAVAILABLE"
	JobGetResponseEnhancedAsyncJobResponseErrorNameContextWindowExceeded      JobGetResponseEnhancedAsyncJobResponseErrorName = "CONTEXT_WINDOW_EXCEEDED"
	JobGetResponseEnhancedAsyncJobResponseErrorNameProcessingFailed           JobGetResponseEnhancedAsyncJobResponseErrorName = "PROCESSING_FAILED"
	JobGetResponseEnhancedAsyncJobResponseErrorNameInferenceMethodUnsupported JobGetResponseEnhancedAsyncJobResponseErrorName = "INFERENCE_METHOD_UNSUPPORTED"
	JobGetResponseEnhancedAsyncJobResponseErrorNameSubprocessCrashed          JobGetResponseEnhancedAsyncJobResponseErrorName = "SUBPROCESS_CRASHED"
	JobGetResponseEnhancedAsyncJobResponseErrorNameBatchOrphaned              JobGetResponseEnhancedAsyncJobResponseErrorName = "BATCH_ORPHANED"
	JobGetResponseEnhancedAsyncJobResponseErrorNameOversizedResult            JobGetResponseEnhancedAsyncJobResponseErrorName = "OVERSIZED_RESULT"
	JobGetResponseEnhancedAsyncJobResponseErrorNameLlmOutputParseFailed       JobGetResponseEnhancedAsyncJobResponseErrorName = "LLM_OUTPUT_PARSE_FAILED"
	JobGetResponseEnhancedAsyncJobResponseErrorNameLlmProviderError           JobGetResponseEnhancedAsyncJobResponseErrorName = "LLM_PROVIDER_ERROR"
	JobGetResponseEnhancedAsyncJobResponseErrorNameInvalidConfig              JobGetResponseEnhancedAsyncJobResponseErrorName = "INVALID_CONFIG"
	JobGetResponseEnhancedAsyncJobResponseErrorNameInvalidSchema              JobGetResponseEnhancedAsyncJobResponseErrorName = "INVALID_SCHEMA"
	JobGetResponseEnhancedAsyncJobResponseErrorNameAuthError                  JobGetResponseEnhancedAsyncJobResponseErrorName = "AUTH_ERROR"
	JobGetResponseEnhancedAsyncJobResponseErrorNameNotApplicable              JobGetResponseEnhancedAsyncJobResponseErrorName = "NOT_APPLICABLE"
	JobGetResponseEnhancedAsyncJobResponseErrorNameRegionUnavailable          JobGetResponseEnhancedAsyncJobResponseErrorName = "REGION_UNAVAILABLE"
	JobGetResponseEnhancedAsyncJobResponseErrorNameNotFound                   JobGetResponseEnhancedAsyncJobResponseErrorName = "NOT_FOUND"
	JobGetResponseEnhancedAsyncJobResponseErrorNameJobDeletionInProgress      JobGetResponseEnhancedAsyncJobResponseErrorName = "JOB_DELETION_IN_PROGRESS"
	JobGetResponseEnhancedAsyncJobResponseErrorNameJobDeleted                 JobGetResponseEnhancedAsyncJobResponseErrorName = "JOB_DELETED"
	JobGetResponseEnhancedAsyncJobResponseErrorNameJobNotComplete             JobGetResponseEnhancedAsyncJobResponseErrorName = "JOB_NOT_COMPLETE"
	JobGetResponseEnhancedAsyncJobResponseErrorNameJobCancelled               JobGetResponseEnhancedAsyncJobResponseErrorName = "JOB_CANCELLED"
	JobGetResponseEnhancedAsyncJobResponseErrorNameRateLimit                  JobGetResponseEnhancedAsyncJobResponseErrorName = "RATE_LIMIT"
	JobGetResponseEnhancedAsyncJobResponseErrorNameCellCountExceeded          JobGetResponseEnhancedAsyncJobResponseErrorName = "CELL_COUNT_EXCEEDED"
	JobGetResponseEnhancedAsyncJobResponseErrorNameURLNotAllowed              JobGetResponseEnhancedAsyncJobResponseErrorName = "URL_NOT_ALLOWED"
)

func (r JobGetResponseEnhancedAsyncJobResponseErrorName) IsKnown() bool {
	switch r {
	case JobGetResponseEnhancedAsyncJobResponseErrorNameTimeout, JobGetResponseEnhancedAsyncJobResponseErrorNameCapacityTimeout, JobGetResponseEnhancedAsyncJobResponseErrorNameCustomerTimeout, JobGetResponseEnhancedAsyncJobResponseErrorNameInternalError, JobGetResponseEnhancedAsyncJobResponseErrorNameServiceUnavailable, JobGetResponseEnhancedAsyncJobResponseErrorNameGPUAllocationError, JobGetResponseEnhancedAsyncJobResponseErrorNameGPUPoolSaturated, JobGetResponseEnhancedAsyncJobResponseErrorNameBatchQueueFull, JobGetResponseEnhancedAsyncJobResponseErrorNameJobStateError, JobGetResponseEnhancedAsyncJobResponseErrorNameDocumentCorrupt, JobGetResponseEnhancedAsyncJobResponseErrorNameDocumentEmpty, JobGetResponseEnhancedAsyncJobResponseErrorNameDocumentUnsupported, JobGetResponseEnhancedAsyncJobResponseErrorNameDocumentTooLarge, JobGetResponseEnhancedAsyncJobResponseErrorNameOfficeConversionTooLarge, JobGetResponseEnhancedAsyncJobResponseErrorNameImageTooLarge, JobGetResponseEnhancedAsyncJobResponseErrorNameImageTooSmall, JobGetResponseEnhancedAsyncJobResponseErrorNameImageInvalidAspectRatio, JobGetResponseEnhancedAsyncJobResponseErrorNameDocumentPasswordProtected, JobGetResponseEnhancedAsyncJobResponseErrorNameFormFillFailed, JobGetResponseEnhancedAsyncJobResponseErrorNameInternalInvariantViolation, JobGetResponseEnhancedAsyncJobResponseErrorNameGPUUnavailable, JobGetResponseEnhancedAsyncJobResponseErrorNameContextWindowExceeded, JobGetResponseEnhancedAsyncJobResponseErrorNameProcessingFailed, JobGetResponseEnhancedAsyncJobResponseErrorNameInferenceMethodUnsupported, JobGetResponseEnhancedAsyncJobResponseErrorNameSubprocessCrashed, JobGetResponseEnhancedAsyncJobResponseErrorNameBatchOrphaned, JobGetResponseEnhancedAsyncJobResponseErrorNameOversizedResult, JobGetResponseEnhancedAsyncJobResponseErrorNameLlmOutputParseFailed, JobGetResponseEnhancedAsyncJobResponseErrorNameLlmProviderError, JobGetResponseEnhancedAsyncJobResponseErrorNameInvalidConfig, JobGetResponseEnhancedAsyncJobResponseErrorNameInvalidSchema, JobGetResponseEnhancedAsyncJobResponseErrorNameAuthError, JobGetResponseEnhancedAsyncJobResponseErrorNameNotApplicable, JobGetResponseEnhancedAsyncJobResponseErrorNameRegionUnavailable, JobGetResponseEnhancedAsyncJobResponseErrorNameNotFound, JobGetResponseEnhancedAsyncJobResponseErrorNameJobDeletionInProgress, JobGetResponseEnhancedAsyncJobResponseErrorNameJobDeleted, JobGetResponseEnhancedAsyncJobResponseErrorNameJobNotComplete, JobGetResponseEnhancedAsyncJobResponseErrorNameJobCancelled, JobGetResponseEnhancedAsyncJobResponseErrorNameRateLimit, JobGetResponseEnhancedAsyncJobResponseErrorNameCellCountExceeded, JobGetResponseEnhancedAsyncJobResponseErrorNameURLNotAllowed:
		return true
	}
	return false
}

// Response from classify job - returned when polling /job/{job_id}
type JobGetResponseEnhancedAsyncJobResponseResult struct {
	// This field can have the runtime type of [shared.ExtractResponseCitationsUnion].
	Citations interface{} `json:"citations"`
	// Optional document-level deep extract confidence label.
	Confidence JobGetResponseEnhancedAsyncJobResponseResultConfidence `json:"confidence" api:"nullable"`
	// Optional explanation for the document-level confidence label.
	ConfidenceReason string `json:"confidence_reason" api:"nullable"`
	// This field can have the runtime type of
	// [shared.ParseResponseDocumentProperties].
	DocumentProperties interface{} `json:"document_properties"`
	// Presigned URL to download the edited document.
	DocumentURL string `json:"document_url"`
	// The duration of the parse request in seconds.
	Duration float64 `json:"duration" api:"nullable"`
	// This field can have the runtime type of [map[string]string].
	ExtraMetadata interface{} `json:"extra_metadata"`
	// This field can have the runtime type of [[]EditWidget].
	FormSchema interface{} `json:"form_schema"`
	JobID      string      `json:"job_id" api:"nullable"`
	// The storage URL of the converted PDF file.
	PdfURL string `json:"pdf_url" api:"nullable"`
	// This field can have the runtime type of [map[string]interface{}],
	// [shared.ClassifyResponseResponseConfidence].
	ResponseConfidence interface{}                                              `json:"response_confidence"`
	ResponseType       JobGetResponseEnhancedAsyncJobResponseResultResponseType `json:"response_type"`
	// This field can have the runtime type of [shared.ParseResponseResult],
	// [shared.ExtractResponseResultUnion], [shared.SplitResponseResult],
	// [shared.PipelineResponseResult], [[]interface{}],
	// [shared.ClassifyResponseResult],
	// [JobGetResponseEnhancedAsyncJobResponseResultChartResponseResult].
	Result interface{} `json:"result"`
	// The link to the studio pipeline for the document.
	StudioLink string `json:"studio_link" api:"nullable"`
	// This field can have the runtime type of [ParseUsage], [ExtractUsage],
	// [shared.ClassifyResponseUsage],
	// [JobGetResponseEnhancedAsyncJobResponseResultChartResponseUsage].
	Usage interface{}                                      `json:"usage"`
	JSON  jobGetResponseEnhancedAsyncJobResponseResultJSON `json:"-"`
	union JobGetResponseEnhancedAsyncJobResponseResultUnion
}

// jobGetResponseEnhancedAsyncJobResponseResultJSON contains the JSON metadata for
// the struct [JobGetResponseEnhancedAsyncJobResponseResult]
type jobGetResponseEnhancedAsyncJobResponseResultJSON struct {
	Citations          apijson.Field
	Confidence         apijson.Field
	ConfidenceReason   apijson.Field
	DocumentProperties apijson.Field
	DocumentURL        apijson.Field
	Duration           apijson.Field
	ExtraMetadata      apijson.Field
	FormSchema         apijson.Field
	JobID              apijson.Field
	PdfURL             apijson.Field
	ResponseConfidence apijson.Field
	ResponseType       apijson.Field
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
// Possible runtime types of the union are [shared.ParseResponse],
// [shared.ExtractResponse], [shared.SplitResponse], [shared.EditResponse],
// [shared.PipelineResponse], [V3Extract], [shared.ClassifyResponse],
// [JobGetResponseEnhancedAsyncJobResponseResultChartResponse].
func (r JobGetResponseEnhancedAsyncJobResponseResult) AsUnion() JobGetResponseEnhancedAsyncJobResponseResultUnion {
	return r.union
}

// Response from classify job - returned when polling /job/{job_id}
//
// Union satisfied by [shared.ParseResponse], [shared.ExtractResponse],
// [shared.SplitResponse], [shared.EditResponse], [shared.PipelineResponse],
// [V3Extract], [shared.ClassifyResponse] or
// [JobGetResponseEnhancedAsyncJobResponseResultChartResponse].
type JobGetResponseEnhancedAsyncJobResponseResultUnion interface {
	ImplementsJobGetResponseEnhancedAsyncJobResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*JobGetResponseEnhancedAsyncJobResponseResultUnion)(nil)).Elem(),
		"response_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.ParseResponse{}),
			DiscriminatorValue: "parse",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.ExtractResponse{}),
			DiscriminatorValue: "extract",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.SplitResponse{}),
			DiscriminatorValue: "split",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.EditResponse{}),
			DiscriminatorValue: "edit",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.PipelineResponse{}),
			DiscriminatorValue: "pipeline",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(V3Extract{}),
			DiscriminatorValue: "v3_extract",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.ClassifyResponse{}),
			DiscriminatorValue: "classify",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(JobGetResponseEnhancedAsyncJobResponseResultChartResponse{}),
			DiscriminatorValue: "chart",
		},
	)
}

// Response from synchronous and persisted chart extraction jobs.
type JobGetResponseEnhancedAsyncJobResponseResultChartResponse struct {
	JobID        string                                                                `json:"job_id" api:"required"`
	Result       JobGetResponseEnhancedAsyncJobResponseResultChartResponseResult       `json:"result" api:"required"`
	Duration     float64                                                               `json:"duration" api:"nullable"`
	ResponseType JobGetResponseEnhancedAsyncJobResponseResultChartResponseResponseType `json:"response_type"`
	Usage        JobGetResponseEnhancedAsyncJobResponseResultChartResponseUsage        `json:"usage"`
	JSON         jobGetResponseEnhancedAsyncJobResponseResultChartResponseJSON         `json:"-"`
}

// jobGetResponseEnhancedAsyncJobResponseResultChartResponseJSON contains the JSON
// metadata for the struct
// [JobGetResponseEnhancedAsyncJobResponseResultChartResponse]
type jobGetResponseEnhancedAsyncJobResponseResultChartResponseJSON struct {
	JobID        apijson.Field
	Result       apijson.Field
	Duration     apijson.Field
	ResponseType apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *JobGetResponseEnhancedAsyncJobResponseResultChartResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseEnhancedAsyncJobResponseResultChartResponseJSON) RawJSON() string {
	return r.raw
}

func (r JobGetResponseEnhancedAsyncJobResponseResultChartResponse) ImplementsJobGetResponseEnhancedAsyncJobResponseResult() {
}

type JobGetResponseEnhancedAsyncJobResponseResultChartResponseResult struct {
	ChartData         map[string]interface{}                                              `json:"chart_data" api:"required"`
	ReconstructionURL string                                                              `json:"reconstruction_url" api:"required"`
	Summary           string                                                              `json:"summary" api:"required"`
	Verified          bool                                                                `json:"verified" api:"required"`
	JSON              jobGetResponseEnhancedAsyncJobResponseResultChartResponseResultJSON `json:"-"`
}

// jobGetResponseEnhancedAsyncJobResponseResultChartResponseResultJSON contains the
// JSON metadata for the struct
// [JobGetResponseEnhancedAsyncJobResponseResultChartResponseResult]
type jobGetResponseEnhancedAsyncJobResponseResultChartResponseResultJSON struct {
	ChartData         apijson.Field
	ReconstructionURL apijson.Field
	Summary           apijson.Field
	Verified          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *JobGetResponseEnhancedAsyncJobResponseResultChartResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseEnhancedAsyncJobResponseResultChartResponseResultJSON) RawJSON() string {
	return r.raw
}

type JobGetResponseEnhancedAsyncJobResponseResultChartResponseResponseType string

const (
	JobGetResponseEnhancedAsyncJobResponseResultChartResponseResponseTypeChart JobGetResponseEnhancedAsyncJobResponseResultChartResponseResponseType = "chart"
)

func (r JobGetResponseEnhancedAsyncJobResponseResultChartResponseResponseType) IsKnown() bool {
	switch r {
	case JobGetResponseEnhancedAsyncJobResponseResultChartResponseResponseTypeChart:
		return true
	}
	return false
}

type JobGetResponseEnhancedAsyncJobResponseResultChartResponseUsage struct {
	Credits   float64                                                                 `json:"credits"`
	NumCharts JobGetResponseEnhancedAsyncJobResponseResultChartResponseUsageNumCharts `json:"num_charts"`
	JSON      jobGetResponseEnhancedAsyncJobResponseResultChartResponseUsageJSON      `json:"-"`
}

// jobGetResponseEnhancedAsyncJobResponseResultChartResponseUsageJSON contains the
// JSON metadata for the struct
// [JobGetResponseEnhancedAsyncJobResponseResultChartResponseUsage]
type jobGetResponseEnhancedAsyncJobResponseResultChartResponseUsageJSON struct {
	Credits     apijson.Field
	NumCharts   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseEnhancedAsyncJobResponseResultChartResponseUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseEnhancedAsyncJobResponseResultChartResponseUsageJSON) RawJSON() string {
	return r.raw
}

type JobGetResponseEnhancedAsyncJobResponseResultChartResponseUsageNumCharts int64

const (
	JobGetResponseEnhancedAsyncJobResponseResultChartResponseUsageNumCharts1 JobGetResponseEnhancedAsyncJobResponseResultChartResponseUsageNumCharts = 1
)

func (r JobGetResponseEnhancedAsyncJobResponseResultChartResponseUsageNumCharts) IsKnown() bool {
	switch r {
	case JobGetResponseEnhancedAsyncJobResponseResultChartResponseUsageNumCharts1:
		return true
	}
	return false
}

// Optional document-level deep extract confidence label.
type JobGetResponseEnhancedAsyncJobResponseResultConfidence string

const (
	JobGetResponseEnhancedAsyncJobResponseResultConfidenceHigh JobGetResponseEnhancedAsyncJobResponseResultConfidence = "high"
	JobGetResponseEnhancedAsyncJobResponseResultConfidenceLow  JobGetResponseEnhancedAsyncJobResponseResultConfidence = "low"
)

func (r JobGetResponseEnhancedAsyncJobResponseResultConfidence) IsKnown() bool {
	switch r {
	case JobGetResponseEnhancedAsyncJobResponseResultConfidenceHigh, JobGetResponseEnhancedAsyncJobResponseResultConfidenceLow:
		return true
	}
	return false
}

type JobGetResponseEnhancedAsyncJobResponseResultResponseType string

const (
	JobGetResponseEnhancedAsyncJobResponseResultResponseTypeParse     JobGetResponseEnhancedAsyncJobResponseResultResponseType = "parse"
	JobGetResponseEnhancedAsyncJobResponseResultResponseTypeExtract   JobGetResponseEnhancedAsyncJobResponseResultResponseType = "extract"
	JobGetResponseEnhancedAsyncJobResponseResultResponseTypeSplit     JobGetResponseEnhancedAsyncJobResponseResultResponseType = "split"
	JobGetResponseEnhancedAsyncJobResponseResultResponseTypeEdit      JobGetResponseEnhancedAsyncJobResponseResultResponseType = "edit"
	JobGetResponseEnhancedAsyncJobResponseResultResponseTypePipeline  JobGetResponseEnhancedAsyncJobResponseResultResponseType = "pipeline"
	JobGetResponseEnhancedAsyncJobResponseResultResponseTypeV3Extract JobGetResponseEnhancedAsyncJobResponseResultResponseType = "v3_extract"
	JobGetResponseEnhancedAsyncJobResponseResultResponseTypeClassify  JobGetResponseEnhancedAsyncJobResponseResultResponseType = "classify"
	JobGetResponseEnhancedAsyncJobResponseResultResponseTypeChart     JobGetResponseEnhancedAsyncJobResponseResultResponseType = "chart"
)

func (r JobGetResponseEnhancedAsyncJobResponseResultResponseType) IsKnown() bool {
	switch r {
	case JobGetResponseEnhancedAsyncJobResponseResultResponseTypeParse, JobGetResponseEnhancedAsyncJobResponseResultResponseTypeExtract, JobGetResponseEnhancedAsyncJobResponseResultResponseTypeSplit, JobGetResponseEnhancedAsyncJobResponseResultResponseTypeEdit, JobGetResponseEnhancedAsyncJobResponseResultResponseTypePipeline, JobGetResponseEnhancedAsyncJobResponseResultResponseTypeV3Extract, JobGetResponseEnhancedAsyncJobResponseResultResponseTypeClassify, JobGetResponseEnhancedAsyncJobResponseResultResponseTypeChart:
		return true
	}
	return false
}

type JobGetResponseEnhancedAsyncJobResponseType string

const (
	JobGetResponseEnhancedAsyncJobResponseTypeParse    JobGetResponseEnhancedAsyncJobResponseType = "Parse"
	JobGetResponseEnhancedAsyncJobResponseTypeExtract  JobGetResponseEnhancedAsyncJobResponseType = "Extract"
	JobGetResponseEnhancedAsyncJobResponseTypeSplit    JobGetResponseEnhancedAsyncJobResponseType = "Split"
	JobGetResponseEnhancedAsyncJobResponseTypeEdit     JobGetResponseEnhancedAsyncJobResponseType = "Edit"
	JobGetResponseEnhancedAsyncJobResponseTypePipeline JobGetResponseEnhancedAsyncJobResponseType = "Pipeline"
	JobGetResponseEnhancedAsyncJobResponseTypeClassify JobGetResponseEnhancedAsyncJobResponseType = "Classify"
	JobGetResponseEnhancedAsyncJobResponseTypeChart    JobGetResponseEnhancedAsyncJobResponseType = "Chart"
)

func (r JobGetResponseEnhancedAsyncJobResponseType) IsKnown() bool {
	switch r {
	case JobGetResponseEnhancedAsyncJobResponseTypeParse, JobGetResponseEnhancedAsyncJobResponseTypeExtract, JobGetResponseEnhancedAsyncJobResponseTypeSplit, JobGetResponseEnhancedAsyncJobResponseTypeEdit, JobGetResponseEnhancedAsyncJobResponseTypePipeline, JobGetResponseEnhancedAsyncJobResponseTypeClassify, JobGetResponseEnhancedAsyncJobResponseTypeChart:
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
	JobGetResponseTypeChart    JobGetResponseType = "Chart"
)

func (r JobGetResponseType) IsKnown() bool {
	switch r {
	case JobGetResponseTypeParse, JobGetResponseTypeExtract, JobGetResponseTypeSplit, JobGetResponseTypeEdit, JobGetResponseTypePipeline, JobGetResponseTypeClassify, JobGetResponseTypeChart:
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
	JobGetAllResponseJobsTypeChart    JobGetAllResponseJobsType = "Chart"
)

func (r JobGetAllResponseJobsType) IsKnown() bool {
	switch r {
	case JobGetAllResponseJobsTypeParse, JobGetAllResponseJobsTypeExtract, JobGetAllResponseJobsTypeSplit, JobGetAllResponseJobsTypeEdit, JobGetAllResponseJobsTypePipeline, JobGetAllResponseJobsTypeClassify, JobGetAllResponseJobsTypeChart:
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
