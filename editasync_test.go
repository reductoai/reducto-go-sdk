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

func TestEditAsyncNewWithOptionalParams(t *testing.T) {
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
	_, err := client.EditAsync.New(context.TODO(), reducto.EditAsyncNewParams{
		DocumentURL:      reducto.F[reducto.EditAsyncNewParamsDocumentURLUnion](shared.UnionString("string")),
		EditInstructions: reducto.F("edit_instructions"),
		EditOptions: reducto.F(reducto.EditOptionsParam{
			Color:                 reducto.F("#e1cb97"),
			EnableOverflowPages:   reducto.F(true),
			Flatten:               reducto.F(true),
			FontSize:              reducto.F(1.000000),
			LlmProviderPreference: reducto.F(reducto.EditOptionsLlmProviderPreferenceOpenAI),
		}),
		FormSchema: reducto.F([]reducto.EditWidgetParam{{
			Bbox: reducto.F(reducto.BoundingBoxParam{
				Height:       reducto.F(0.000000),
				Left:         reducto.F(0.000000),
				Page:         reducto.F(int64(0)),
				Top:          reducto.F(0.000000),
				Width:        reducto.F(0.000000),
				OriginalPage: reducto.F(int64(0)),
			}),
			Description: reducto.F("description"),
			Type:        reducto.F(reducto.EditWidgetTypeText),
			Fill:        reducto.F(true),
			FontSize:    reducto.F(1.000000),
			Value:       reducto.F("value"),
		}}),
		Priority: reducto.F(true),
		Webhook: reducto.F(reducto.EditAsyncNewParamsWebhook{
			Channels: reducto.F([]string{"string"}),
			Metadata: reducto.F[any](map[string]interface{}{}),
			Mode:     reducto.F(reducto.EditAsyncNewParamsWebhookModeDisabled),
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
