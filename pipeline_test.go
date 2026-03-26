// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reducto_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/reductoai/reducto-go-sdk"
	"github.com/reductoai/reducto-go-sdk/internal/testutil"
	"github.com/reductoai/reducto-go-sdk/option"
	"github.com/reductoai/reducto-go-sdk/shared"
)

func TestPipelineRunWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := reducto.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pipeline.Run(context.TODO(), reducto.PipelineRunParams{
		Input:      reducto.F[reducto.PipelineRunParamsInputUnion](shared.UnionString("string")),
		PipelineID: reducto.F("pipeline_id"),
		Settings: reducto.F(reducto.PipelineSettingsParam{
			DocumentPassword: reducto.F("document_password"),
		}),
	})
	if err != nil {
		var apierr *reducto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPipelineRunJobWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := reducto.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pipeline.RunJob(context.TODO(), reducto.PipelineRunJobParams{
		Input:      reducto.F[reducto.PipelineRunJobParamsInputUnion](shared.UnionString("string")),
		PipelineID: reducto.F("pipeline_id"),
		Async: reducto.F(reducto.AsyncConfigV3Param{
			Metadata: reducto.F[any](map[string]interface{}{}),
			Priority: reducto.F(true),
			Webhook: reducto.F[reducto.AsyncConfigV3WebhookUnionParam](shared.SvixWebhookConfigParam{
				Channels: reducto.F([]string{"string"}),
				Mode:     reducto.F(shared.SvixWebhookConfigModeSvix),
			}),
		}),
		Settings: reducto.F(reducto.PipelineSettingsParam{
			DocumentPassword: reducto.F("document_password"),
		}),
	})
	if err != nil {
		var apierr *reducto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
