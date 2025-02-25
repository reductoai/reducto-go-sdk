// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reductoai_test

import (
	"context"
	"os"
	"testing"

	"github.com/reductoai/reducto-go-sdk"
	"github.com/reductoai/reducto-go-sdk/internal/testutil"
	"github.com/reductoai/reducto-go-sdk/option"
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
	parseResponse, err := client.Parse.Run(context.TODO(), reductoai.ParseRunParams{
		DocumentURL: reductoai.F("https://pdfobject.com/pdf/sample.pdf"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", parseResponse.JobID)
}
