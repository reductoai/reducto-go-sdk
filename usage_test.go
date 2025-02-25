// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reductoai_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/reductoai-go"
	"github.com/stainless-sdks/reductoai-go/internal/testutil"
	"github.com/stainless-sdks/reductoai-go/option"
	"github.com/stainless-sdks/reductoai-go/shared"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := reductoai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	splitResponse, err := client.Split.Run(context.TODO(), reductoai.SplitRunParams{
		DocumentURL: reductoai.F("document_url"),
		SplitDescription: reductoai.F([]shared.SplitCategoryParam{{
			Description: reductoai.F("description"),
			Name:        reductoai.F("name"),
		}}),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", splitResponse.Result)
}
