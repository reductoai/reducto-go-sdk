// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reducto

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/reductoai/reducto-go-sdk/internal/requestconfig"
	"github.com/reductoai/reducto-go-sdk/option"
)

// CancelService contains methods and other services that help with interacting
// with the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCancelService] method instead.
type CancelService struct {
	Options []option.RequestOption
}

// NewCancelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCancelService(opts ...option.RequestOption) (r *CancelService) {
	r = &CancelService{}
	r.Options = opts
	return
}

// Cancel Job
func (r *CancelService) CancelJob(ctx context.Context, jobID string, opts ...option.RequestOption) (res *CancelCancelJobResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("cancel/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type CancelCancelJobResponse = interface{}
