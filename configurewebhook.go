// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reducto

import (
	"context"
	"net/http"
	"slices"

	"github.com/reductoai/reducto-go-sdk/internal/requestconfig"
	"github.com/reductoai/reducto-go-sdk/option"
)

// ConfigureWebhookService contains methods and other services that help with
// interacting with the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigureWebhookService] method instead.
type ConfigureWebhookService struct {
	Options []option.RequestOption
}

// NewConfigureWebhookService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConfigureWebhookService(opts ...option.RequestOption) (r *ConfigureWebhookService) {
	r = &ConfigureWebhookService{}
	r.Options = opts
	return
}

// Webhook Portal
func (r *ConfigureWebhookService) New(ctx context.Context, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "configure_webhook"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}
