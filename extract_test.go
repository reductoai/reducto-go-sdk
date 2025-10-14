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

func TestExtractRunWithOptionalParams(t *testing.T) {
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
	_, err := client.Extract.Run(context.TODO(), reducto.ExtractRunParams{
		Body: reducto.ExtractRunParamsBodySyncExtractConfig{
			Input: reducto.F[reducto.ExtractRunParamsBodySyncExtractConfigInputUnion](shared.UnionString("string")),
			Instructions: reducto.F(reducto.ExtractRunParamsBodySyncExtractConfigInstructions{
				Schema:       reducto.F[any](map[string]interface{}{}),
				SystemPrompt: reducto.F("system_prompt"),
			}),
			Parsing: reducto.F(reducto.ExtractRunParamsBodySyncExtractConfigParsing{
				Enhance: reducto.F(reducto.ExtractRunParamsBodySyncExtractConfigParsingEnhance{
					Agentic: reducto.F([]reducto.ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticUnion{reducto.ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTableAgentic{
						Scope:  reducto.F(reducto.ExtractRunParamsBodySyncExtractConfigParsingEnhanceAgenticTableAgenticScopeTable),
						Prompt: reducto.F("prompt"),
					}}),
					SummarizeFigures: reducto.F(true),
				}),
				Formatting: reducto.F(reducto.ExtractRunParamsBodySyncExtractConfigParsingFormatting{
					AddPageMarkers:    reducto.F(true),
					Include:           reducto.F([]reducto.ExtractRunParamsBodySyncExtractConfigParsingFormattingInclude{reducto.ExtractRunParamsBodySyncExtractConfigParsingFormattingIncludeChangeTracking}),
					MergeTables:       reducto.F(true),
					TableOutputFormat: reducto.F(reducto.ExtractRunParamsBodySyncExtractConfigParsingFormattingTableOutputFormatHTML),
				}),
				Retrieval: reducto.F(reducto.ExtractRunParamsBodySyncExtractConfigParsingRetrieval{
					Chunking: reducto.F(reducto.ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunking{
						ChunkMode: reducto.F(reducto.ExtractRunParamsBodySyncExtractConfigParsingRetrievalChunkingChunkModeVariable),
						ChunkSize: reducto.F(int64(0)),
					}),
					EmbeddingOptimized: reducto.F(true),
					FilterBlocks:       reducto.F([]reducto.ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlock{reducto.ExtractRunParamsBodySyncExtractConfigParsingRetrievalFilterBlockHeader}),
				}),
				Settings: reducto.F(reducto.ExtractRunParamsBodySyncExtractConfigParsingSettings{
					DocumentPassword:   reducto.F("document_password"),
					EmbedPdfMetadata:   reducto.F(true),
					ForceFileExtension: reducto.F("force_file_extension"),
					ForceURLResult:     reducto.F(true),
					OcrSystem:          reducto.F(reducto.ExtractRunParamsBodySyncExtractConfigParsingSettingsOcrSystemStandard),
					PageRange: reducto.F[reducto.ExtractRunParamsBodySyncExtractConfigParsingSettingsPageRangeUnion](shared.PageRangeParam{
						End:   reducto.F(int64(0)),
						Start: reducto.F(int64(0)),
					}),
					PersistResults: reducto.F(true),
					ReturnImages:   reducto.F([]reducto.ExtractRunParamsBodySyncExtractConfigParsingSettingsReturnImage{reducto.ExtractRunParamsBodySyncExtractConfigParsingSettingsReturnImageFigure}),
					ReturnOcrData:  reducto.F(true),
					Timeout:        reducto.F(0.000000),
				}),
				Spreadsheet: reducto.F(reducto.ExtractRunParamsBodySyncExtractConfigParsingSpreadsheet{
					Clustering: reducto.F(reducto.ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetClusteringAccurate),
					Exclude:    reducto.F([]reducto.ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetExclude{reducto.ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetExcludeHiddenSheets}),
					Include:    reducto.F([]reducto.ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetInclude{reducto.ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetIncludeCellColors}),
					SplitLargeTables: reducto.F(reducto.ExtractRunParamsBodySyncExtractConfigParsingSpreadsheetSplitLargeTables{
						Enabled: reducto.F(true),
						Size:    reducto.F(int64(0)),
					}),
				}),
			}),
			Settings: reducto.F(reducto.ExtractRunParamsBodySyncExtractConfigSettings{
				ArrayExtract: reducto.F(true),
				Citations: reducto.F(reducto.ExtractRunParamsBodySyncExtractConfigSettingsCitations{
					Enabled:             reducto.F(true),
					NumericalConfidence: reducto.F(true),
				}),
				IncludeImages:      reducto.F(true),
				OptimizeForLatency: reducto.F(true),
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

func TestExtractRunJobWithOptionalParams(t *testing.T) {
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
	_, err := client.Extract.RunJob(context.TODO(), reducto.ExtractRunJobParams{
		Input: reducto.F[reducto.ExtractRunJobParamsInputUnion](shared.UnionString("string")),
		Async: reducto.F(reducto.ExtractRunJobParamsAsync{
			Metadata: reducto.F[any](map[string]interface{}{}),
			Priority: reducto.F(true),
			Webhook: reducto.F[reducto.ExtractRunJobParamsAsyncWebhookUnion](reducto.ExtractRunJobParamsAsyncWebhookSvixWebhookConfig{
				Channels: reducto.F([]string{"string"}),
				Mode:     reducto.F(reducto.ExtractRunJobParamsAsyncWebhookSvixWebhookConfigModeSvix),
			}),
		}),
		Instructions: reducto.F(reducto.ExtractRunJobParamsInstructions{
			Schema:       reducto.F[any](map[string]interface{}{}),
			SystemPrompt: reducto.F("system_prompt"),
		}),
		Parsing: reducto.F(reducto.ExtractRunJobParamsParsing{
			Enhance: reducto.F(reducto.ExtractRunJobParamsParsingEnhance{
				Agentic: reducto.F([]reducto.ExtractRunJobParamsParsingEnhanceAgenticUnion{reducto.ExtractRunJobParamsParsingEnhanceAgenticTableAgentic{
					Scope:  reducto.F(reducto.ExtractRunJobParamsParsingEnhanceAgenticTableAgenticScopeTable),
					Prompt: reducto.F("prompt"),
				}}),
				SummarizeFigures: reducto.F(true),
			}),
			Formatting: reducto.F(reducto.ExtractRunJobParamsParsingFormatting{
				AddPageMarkers:    reducto.F(true),
				Include:           reducto.F([]reducto.ExtractRunJobParamsParsingFormattingInclude{reducto.ExtractRunJobParamsParsingFormattingIncludeChangeTracking}),
				MergeTables:       reducto.F(true),
				TableOutputFormat: reducto.F(reducto.ExtractRunJobParamsParsingFormattingTableOutputFormatHTML),
			}),
			Retrieval: reducto.F(reducto.ExtractRunJobParamsParsingRetrieval{
				Chunking: reducto.F(reducto.ExtractRunJobParamsParsingRetrievalChunking{
					ChunkMode: reducto.F(reducto.ExtractRunJobParamsParsingRetrievalChunkingChunkModeVariable),
					ChunkSize: reducto.F(int64(0)),
				}),
				EmbeddingOptimized: reducto.F(true),
				FilterBlocks:       reducto.F([]reducto.ExtractRunJobParamsParsingRetrievalFilterBlock{reducto.ExtractRunJobParamsParsingRetrievalFilterBlockHeader}),
			}),
			Settings: reducto.F(reducto.ExtractRunJobParamsParsingSettings{
				DocumentPassword:   reducto.F("document_password"),
				EmbedPdfMetadata:   reducto.F(true),
				ForceFileExtension: reducto.F("force_file_extension"),
				ForceURLResult:     reducto.F(true),
				OcrSystem:          reducto.F(reducto.ExtractRunJobParamsParsingSettingsOcrSystemStandard),
				PageRange: reducto.F[reducto.ExtractRunJobParamsParsingSettingsPageRangeUnion](shared.PageRangeParam{
					End:   reducto.F(int64(0)),
					Start: reducto.F(int64(0)),
				}),
				PersistResults: reducto.F(true),
				ReturnImages:   reducto.F([]reducto.ExtractRunJobParamsParsingSettingsReturnImage{reducto.ExtractRunJobParamsParsingSettingsReturnImageFigure}),
				ReturnOcrData:  reducto.F(true),
				Timeout:        reducto.F(0.000000),
			}),
			Spreadsheet: reducto.F(reducto.ExtractRunJobParamsParsingSpreadsheet{
				Clustering: reducto.F(reducto.ExtractRunJobParamsParsingSpreadsheetClusteringAccurate),
				Exclude:    reducto.F([]reducto.ExtractRunJobParamsParsingSpreadsheetExclude{reducto.ExtractRunJobParamsParsingSpreadsheetExcludeHiddenSheets}),
				Include:    reducto.F([]reducto.ExtractRunJobParamsParsingSpreadsheetInclude{reducto.ExtractRunJobParamsParsingSpreadsheetIncludeCellColors}),
				SplitLargeTables: reducto.F(reducto.ExtractRunJobParamsParsingSpreadsheetSplitLargeTables{
					Enabled: reducto.F(true),
					Size:    reducto.F(int64(0)),
				}),
			}),
		}),
		Settings: reducto.F(reducto.ExtractRunJobParamsSettings{
			ArrayExtract: reducto.F(true),
			Citations: reducto.F(reducto.ExtractRunJobParamsSettingsCitations{
				Enabled:             reducto.F(true),
				NumericalConfidence: reducto.F(true),
			}),
			IncludeImages:      reducto.F(true),
			OptimizeForLatency: reducto.F(true),
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
