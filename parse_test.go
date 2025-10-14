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

func TestParseRunWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Parse.Run(context.TODO(), reducto.ParseRunParams{
		Body: reducto.ParseRunParamsBodySyncParseConfig{
			Input: reducto.F[reducto.ParseRunParamsBodySyncParseConfigInputUnion](shared.UnionString("string")),
			Enhance: reducto.F(reducto.ParseRunParamsBodySyncParseConfigEnhance{
				Agentic: reducto.F([]reducto.ParseRunParamsBodySyncParseConfigEnhanceAgenticUnion{reducto.ParseRunParamsBodySyncParseConfigEnhanceAgenticTableAgentic{
					Scope:  reducto.F(reducto.ParseRunParamsBodySyncParseConfigEnhanceAgenticTableAgenticScopeTable),
					Prompt: reducto.F("prompt"),
				}}),
				SummarizeFigures: reducto.F(true),
			}),
			Formatting: reducto.F(reducto.ParseRunParamsBodySyncParseConfigFormatting{
				AddPageMarkers:    reducto.F(true),
				Include:           reducto.F([]reducto.ParseRunParamsBodySyncParseConfigFormattingInclude{reducto.ParseRunParamsBodySyncParseConfigFormattingIncludeChangeTracking}),
				MergeTables:       reducto.F(true),
				TableOutputFormat: reducto.F(reducto.ParseRunParamsBodySyncParseConfigFormattingTableOutputFormatHTML),
			}),
			Retrieval: reducto.F(reducto.ParseRunParamsBodySyncParseConfigRetrieval{
				Chunking: reducto.F(reducto.ParseRunParamsBodySyncParseConfigRetrievalChunking{
					ChunkMode: reducto.F(reducto.ParseRunParamsBodySyncParseConfigRetrievalChunkingChunkModeVariable),
					ChunkSize: reducto.F(int64(0)),
				}),
				EmbeddingOptimized: reducto.F(true),
				FilterBlocks:       reducto.F([]reducto.ParseRunParamsBodySyncParseConfigRetrievalFilterBlock{reducto.ParseRunParamsBodySyncParseConfigRetrievalFilterBlockHeader}),
			}),
			Settings: reducto.F(reducto.ParseRunParamsBodySyncParseConfigSettings{
				DocumentPassword:   reducto.F("document_password"),
				EmbedPdfMetadata:   reducto.F(true),
				ForceFileExtension: reducto.F("force_file_extension"),
				ForceURLResult:     reducto.F(true),
				OcrSystem:          reducto.F(reducto.ParseRunParamsBodySyncParseConfigSettingsOcrSystemStandard),
				PageRange: reducto.F[reducto.ParseRunParamsBodySyncParseConfigSettingsPageRangeUnion](shared.PageRangeParam{
					End:   reducto.F(int64(0)),
					Start: reducto.F(int64(0)),
				}),
				PersistResults: reducto.F(true),
				ReturnImages:   reducto.F([]reducto.ParseRunParamsBodySyncParseConfigSettingsReturnImage{reducto.ParseRunParamsBodySyncParseConfigSettingsReturnImageFigure}),
				ReturnOcrData:  reducto.F(true),
				Timeout:        reducto.F(0.000000),
			}),
			Spreadsheet: reducto.F(reducto.ParseRunParamsBodySyncParseConfigSpreadsheet{
				Clustering: reducto.F(reducto.ParseRunParamsBodySyncParseConfigSpreadsheetClusteringAccurate),
				Exclude:    reducto.F([]reducto.ParseRunParamsBodySyncParseConfigSpreadsheetExclude{reducto.ParseRunParamsBodySyncParseConfigSpreadsheetExcludeHiddenSheets}),
				Include:    reducto.F([]reducto.ParseRunParamsBodySyncParseConfigSpreadsheetInclude{reducto.ParseRunParamsBodySyncParseConfigSpreadsheetIncludeCellColors}),
				SplitLargeTables: reducto.F(reducto.ParseRunParamsBodySyncParseConfigSpreadsheetSplitLargeTables{
					Enabled: reducto.F(true),
					Size:    reducto.F(int64(0)),
				}),
			}),
		},
	})
	if err != nil {
		var apierr *reducto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestParseRunJobWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Parse.RunJob(context.TODO(), reducto.ParseRunJobParams{
		Input: reducto.F[reducto.ParseRunJobParamsInputUnion](shared.UnionString("string")),
		Async: reducto.F(reducto.ParseRunJobParamsAsync{
			Metadata: reducto.F[any](map[string]interface{}{}),
			Priority: reducto.F(true),
			Webhook: reducto.F[reducto.ParseRunJobParamsAsyncWebhookUnion](reducto.ParseRunJobParamsAsyncWebhookSvixWebhookConfig{
				Channels: reducto.F([]string{"string"}),
				Mode:     reducto.F(reducto.ParseRunJobParamsAsyncWebhookSvixWebhookConfigModeSvix),
			}),
		}),
		Enhance: reducto.F(reducto.ParseRunJobParamsEnhance{
			Agentic: reducto.F([]reducto.ParseRunJobParamsEnhanceAgenticUnion{reducto.ParseRunJobParamsEnhanceAgenticTableAgentic{
				Scope:  reducto.F(reducto.ParseRunJobParamsEnhanceAgenticTableAgenticScopeTable),
				Prompt: reducto.F("prompt"),
			}}),
			SummarizeFigures: reducto.F(true),
		}),
		Formatting: reducto.F(reducto.ParseRunJobParamsFormatting{
			AddPageMarkers:    reducto.F(true),
			Include:           reducto.F([]reducto.ParseRunJobParamsFormattingInclude{reducto.ParseRunJobParamsFormattingIncludeChangeTracking}),
			MergeTables:       reducto.F(true),
			TableOutputFormat: reducto.F(reducto.ParseRunJobParamsFormattingTableOutputFormatHTML),
		}),
		Retrieval: reducto.F(reducto.ParseRunJobParamsRetrieval{
			Chunking: reducto.F(reducto.ParseRunJobParamsRetrievalChunking{
				ChunkMode: reducto.F(reducto.ParseRunJobParamsRetrievalChunkingChunkModeVariable),
				ChunkSize: reducto.F(int64(0)),
			}),
			EmbeddingOptimized: reducto.F(true),
			FilterBlocks:       reducto.F([]reducto.ParseRunJobParamsRetrievalFilterBlock{reducto.ParseRunJobParamsRetrievalFilterBlockHeader}),
		}),
		Settings: reducto.F(reducto.ParseRunJobParamsSettings{
			DocumentPassword:   reducto.F("document_password"),
			EmbedPdfMetadata:   reducto.F(true),
			ForceFileExtension: reducto.F("force_file_extension"),
			ForceURLResult:     reducto.F(true),
			OcrSystem:          reducto.F(reducto.ParseRunJobParamsSettingsOcrSystemStandard),
			PageRange: reducto.F[reducto.ParseRunJobParamsSettingsPageRangeUnion](shared.PageRangeParam{
				End:   reducto.F(int64(0)),
				Start: reducto.F(int64(0)),
			}),
			PersistResults: reducto.F(true),
			ReturnImages:   reducto.F([]reducto.ParseRunJobParamsSettingsReturnImage{reducto.ParseRunJobParamsSettingsReturnImageFigure}),
			ReturnOcrData:  reducto.F(true),
			Timeout:        reducto.F(0.000000),
		}),
		Spreadsheet: reducto.F(reducto.ParseRunJobParamsSpreadsheet{
			Clustering: reducto.F(reducto.ParseRunJobParamsSpreadsheetClusteringAccurate),
			Exclude:    reducto.F([]reducto.ParseRunJobParamsSpreadsheetExclude{reducto.ParseRunJobParamsSpreadsheetExcludeHiddenSheets}),
			Include:    reducto.F([]reducto.ParseRunJobParamsSpreadsheetInclude{reducto.ParseRunJobParamsSpreadsheetIncludeCellColors}),
			SplitLargeTables: reducto.F(reducto.ParseRunJobParamsSpreadsheetSplitLargeTables{
				Enabled: reducto.F(true),
				Size:    reducto.F(int64(0)),
			}),
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
