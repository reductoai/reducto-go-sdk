// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reductoai_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/reductoai/reducto-go-sdk"
	"github.com/reductoai/reducto-go-sdk/internal/testutil"
	"github.com/reductoai/reducto-go-sdk/option"
)

func TestReductoaiAPIVersion(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.APIVersion(context.TODO())
	if err != nil {
		var apierr *reductoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReductoaiUploadWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Upload(context.TODO(), reductoai.UploadParams{
		Extension: reductoai.F("extension"),
		File:      reductoai.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
	})
	if err != nil {
		var apierr *reductoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
