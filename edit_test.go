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

func TestEditRunWithOptionalParams(t *testing.T) {
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
	_, err := client.Edit.Run(context.TODO(), reducto.EditRunParams{
		DocumentURL:      reducto.F[reducto.EditRunParamsDocumentURLUnion](shared.UnionString("string")),
		EditInstructions: reducto.F("edit_instructions"),
		EditOptions: reducto.F(reducto.EditRunParamsEditOptions{
			Color:                 reducto.F("#e1cb97"),
			EnableOverflowPages:   reducto.F(true),
			Flatten:               reducto.F(true),
			FontSize:              reducto.F(1.000000),
			LlmProviderPreference: reducto.F(reducto.EditRunParamsEditOptionsLlmProviderPreferenceOpenAI),
		}),
		FormSchema: reducto.F([]reducto.EditRunParamsFormSchema{{
			Bbox: reducto.F(shared.BoundingBoxParam{
				Height:       reducto.F(0.000000),
				Left:         reducto.F(0.000000),
				Page:         reducto.F(int64(0)),
				Top:          reducto.F(0.000000),
				Width:        reducto.F(0.000000),
				OriginalPage: reducto.F(int64(0)),
			}),
			Description: reducto.F("description"),
			Type:        reducto.F(reducto.EditRunParamsFormSchemaTypeText),
			Fill:        reducto.F(true),
			FontSize:    reducto.F(1.000000),
			Value:       reducto.F("value"),
		}}),
		Priority: reducto.F(true),
	})
	if err != nil {
		var apierr *reducto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEditRunJobWithOptionalParams(t *testing.T) {
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
	_, err := client.Edit.RunJob(context.TODO(), reducto.EditRunJobParams{
		DocumentURL:      reducto.F[reducto.EditRunJobParamsDocumentURLUnion](shared.UnionString("string")),
		EditInstructions: reducto.F("edit_instructions"),
		EditOptions: reducto.F(reducto.EditRunJobParamsEditOptions{
			Color:                 reducto.F("#e1cb97"),
			EnableOverflowPages:   reducto.F(true),
			Flatten:               reducto.F(true),
			FontSize:              reducto.F(1.000000),
			LlmProviderPreference: reducto.F(reducto.EditRunJobParamsEditOptionsLlmProviderPreferenceOpenAI),
		}),
		FormSchema: reducto.F([]reducto.EditRunJobParamsFormSchema{{
			Bbox: reducto.F(shared.BoundingBoxParam{
				Height:       reducto.F(0.000000),
				Left:         reducto.F(0.000000),
				Page:         reducto.F(int64(0)),
				Top:          reducto.F(0.000000),
				Width:        reducto.F(0.000000),
				OriginalPage: reducto.F(int64(0)),
			}),
			Description: reducto.F("description"),
			Type:        reducto.F(reducto.EditRunJobParamsFormSchemaTypeText),
			Fill:        reducto.F(true),
			FontSize:    reducto.F(1.000000),
			Value:       reducto.F("value"),
		}}),
		Priority: reducto.F(true),
		Webhook: reducto.F(shared.WebhookConfigNewParam{
			Channels: reducto.F([]string{"string"}),
			Metadata: reducto.F[any](map[string]interface{}{}),
			Mode:     reducto.F(shared.WebhookConfigNewModeDisabled),
			URL:      reducto.F("url"),
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
