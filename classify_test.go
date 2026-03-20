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

func TestClassifyClassifyWithOptionalParams(t *testing.T) {
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
	_, err := client.Classify.Classify(context.TODO(), reducto.ClassifyClassifyParams{
		Input: reducto.F[reducto.ClassifyClassifyParamsInputUnion](shared.UnionString("string")),
		ClassificationSchema: reducto.F([]reducto.ClassifyClassifyParamsClassificationSchema{{
			Category: reducto.F("category"),
			Criteria: reducto.F([]string{"string"}),
		}}),
		DocumentMetadata: reducto.F("document_metadata"),
		PageRange: reducto.F[reducto.ClassifyClassifyParamsPageRangeUnion](reducto.PageRangeParam{
			End:   reducto.F(int64(0)),
			Start: reducto.F(int64(0)),
		}),
		PersistResults: reducto.F(true),
	})
	if err != nil {
		var apierr *reducto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
