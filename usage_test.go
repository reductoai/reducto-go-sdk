// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reducto_test

import (
	"context"
	"os"
	"testing"

	"github.com/reductoai/reducto-go-sdk"
	"github.com/reductoai/reducto-go-sdk/internal/testutil"
	"github.com/reductoai/reducto-go-sdk/option"
	"github.com/reductoai/reducto-go-sdk/shared"
)

func TestUsage(t *testing.T) {
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
	parseResponse, err := client.Parse.Run(context.TODO(), reducto.ParseRunParams{
		DocumentURL: reducto.F[reducto.ParseRunParamsDocumentURLUnion](shared.UnionString("https://pdfobject.com/pdf/sample.pdf")),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", parseResponse.JobID)
}
