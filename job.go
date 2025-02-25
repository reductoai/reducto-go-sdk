// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reductoai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/reductoai-go/internal/apijson"
	"github.com/stainless-sdks/reductoai-go/internal/requestconfig"
	"github.com/stainless-sdks/reductoai-go/option"
	"github.com/stainless-sdks/reductoai-go/shared"
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
	Status   JobGetResponseStatus `json:"status,required"`
	Progress float64              `json:"progress,nullable"`
	Reason   string               `json:"reason,nullable"`
	Result   JobGetResponseResult `json:"result,nullable"`
	JSON     jobGetResponseJSON   `json:"-"`
}

// jobGetResponseJSON contains the JSON metadata for the struct [JobGetResponse]
type jobGetResponseJSON struct {
	Status      apijson.Field
	Progress    apijson.Field
	Reason      apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseJSON) RawJSON() string {
	return r.raw
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

type JobGetResponseResult struct {
	// This field can have the runtime type of [shared.ParseResponseResult],
	// [[]interface{}], [shared.SplitResponseResult].
	Result interface{} `json:"result,required"`
	// This field can have the runtime type of [shared.ParseUsage],
	// [shared.ExtractResponseUsage].
	Usage interface{} `json:"usage,required"`
	// This field can have the runtime type of [[]interface{}].
	Citations interface{} `json:"citations"`
	// The duration of the parse request in seconds.
	Duration float64 `json:"duration"`
	JobID    string  `json:"job_id"`
	// The storage URL of the converted PDF file.
	PdfURL string                   `json:"pdf_url,nullable"`
	JSON   jobGetResponseResultJSON `json:"-"`
	union  JobGetResponseResultUnion
}

// jobGetResponseResultJSON contains the JSON metadata for the struct
// [JobGetResponseResult]
type jobGetResponseResultJSON struct {
	Result      apijson.Field
	Usage       apijson.Field
	Citations   apijson.Field
	Duration    apijson.Field
	JobID       apijson.Field
	PdfURL      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r jobGetResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *JobGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = JobGetResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [JobGetResponseResultUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [shared.ParseResponse],
// [shared.ExtractResponse], [shared.SplitResponse].
func (r JobGetResponseResult) AsUnion() JobGetResponseResultUnion {
	return r.union
}

// Union satisfied by [shared.ParseResponse], [shared.ExtractResponse] or
// [shared.SplitResponse].
type JobGetResponseResultUnion interface {
	ImplementsJobGetResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*JobGetResponseResultUnion)(nil)).Elem(),
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
	)
}
