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
			Parsing: reducto.F(shared.ParseOptionsParam{
				Enhance: reducto.F(shared.EnhanceParam{
					Agentic: reducto.F([]shared.EnhanceAgenticUnionParam{shared.TableAgenticParam{
						Scope:  reducto.F(shared.TableAgenticScopeTable),
						Prompt: reducto.F("prompt"),
					}}),
					IntelligentOrdering: reducto.F(true),
					SummarizeFigures:    reducto.F(true),
				}),
				Formatting: reducto.F(shared.FormattingParam{
					AddPageMarkers:    reducto.F(true),
					Include:           reducto.F([]shared.FormattingInclude{shared.FormattingIncludeChangeTracking}),
					MergeTables:       reducto.F(true),
					TableOutputFormat: reducto.F(shared.FormattingTableOutputFormatHTML),
				}),
				Retrieval: reducto.F(shared.RetrievalParam{
					Chunking: reducto.F(shared.ChunkingParam{
						ChunkMode: reducto.F(shared.ChunkingChunkModeVariable),
						ChunkSize: reducto.F(int64(0)),
					}),
					EmbeddingOptimized: reducto.F(true),
					FilterBlocks:       reducto.F([]shared.RetrievalFilterBlock{shared.RetrievalFilterBlockHeader}),
				}),
				Settings: reducto.F(shared.SettingsParam{
					DocumentPassword:   reducto.F("document_password"),
					EmbedPdfMetadata:   reducto.F(true),
					ExtractionMode:     reducto.F(shared.SettingsExtractionModeOcr),
					ForceFileExtension: reducto.F("force_file_extension"),
					ForceURLResult:     reducto.F(true),
					OcrSystem:          reducto.F(shared.SettingsOcrSystemStandard),
					PageRange: reducto.F[shared.SettingsPageRangeUnionParam](shared.PageRangeParam{
						End:   reducto.F(int64(0)),
						Start: reducto.F(int64(0)),
					}),
					PersistResults: reducto.F(true),
					ReturnImages:   reducto.F([]shared.SettingsReturnImage{shared.SettingsReturnImageFigure}),
					ReturnOcrData:  reducto.F(true),
					Timeout:        reducto.F(0.000000),
				}),
				Spreadsheet: reducto.F(shared.SpreadsheetParam{
					Clustering: reducto.F(shared.SpreadsheetClusteringAccurate),
					Exclude:    reducto.F([]shared.SpreadsheetExclude{shared.SpreadsheetExcludeHiddenSheets}),
					Include:    reducto.F([]shared.SpreadsheetInclude{shared.SpreadsheetIncludeCellColors}),
					SplitLargeTables: reducto.F(shared.SplitLargeTablesParam{
						Enabled: reducto.F(true),
						Size:    reducto.F[shared.SplitLargeTablesSizeUnionParam](shared.UnionInt(int64(0))),
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
		Async: reducto.F(shared.ConfigV3AsyncConfigParam{
			Metadata: reducto.F[any](map[string]interface{}{}),
			Priority: reducto.F(true),
			Webhook: reducto.F[shared.ConfigV3AsyncConfigWebhookUnionParam](shared.SvixWebhookConfigParam{
				Channels: reducto.F([]string{"string"}),
				Mode:     reducto.F(shared.SvixWebhookConfigModeSvix),
			}),
		}),
		Instructions: reducto.F(reducto.ExtractRunJobParamsInstructions{
			Schema:       reducto.F[any](map[string]interface{}{}),
			SystemPrompt: reducto.F("system_prompt"),
		}),
		Parsing: reducto.F(shared.ParseOptionsParam{
			Enhance: reducto.F(shared.EnhanceParam{
				Agentic: reducto.F([]shared.EnhanceAgenticUnionParam{shared.TableAgenticParam{
					Scope:  reducto.F(shared.TableAgenticScopeTable),
					Prompt: reducto.F("prompt"),
				}}),
				IntelligentOrdering: reducto.F(true),
				SummarizeFigures:    reducto.F(true),
			}),
			Formatting: reducto.F(shared.FormattingParam{
				AddPageMarkers:    reducto.F(true),
				Include:           reducto.F([]shared.FormattingInclude{shared.FormattingIncludeChangeTracking}),
				MergeTables:       reducto.F(true),
				TableOutputFormat: reducto.F(shared.FormattingTableOutputFormatHTML),
			}),
			Retrieval: reducto.F(shared.RetrievalParam{
				Chunking: reducto.F(shared.ChunkingParam{
					ChunkMode: reducto.F(shared.ChunkingChunkModeVariable),
					ChunkSize: reducto.F(int64(0)),
				}),
				EmbeddingOptimized: reducto.F(true),
				FilterBlocks:       reducto.F([]shared.RetrievalFilterBlock{shared.RetrievalFilterBlockHeader}),
			}),
			Settings: reducto.F(shared.SettingsParam{
				DocumentPassword:   reducto.F("document_password"),
				EmbedPdfMetadata:   reducto.F(true),
				ExtractionMode:     reducto.F(shared.SettingsExtractionModeOcr),
				ForceFileExtension: reducto.F("force_file_extension"),
				ForceURLResult:     reducto.F(true),
				OcrSystem:          reducto.F(shared.SettingsOcrSystemStandard),
				PageRange: reducto.F[shared.SettingsPageRangeUnionParam](shared.PageRangeParam{
					End:   reducto.F(int64(0)),
					Start: reducto.F(int64(0)),
				}),
				PersistResults: reducto.F(true),
				ReturnImages:   reducto.F([]shared.SettingsReturnImage{shared.SettingsReturnImageFigure}),
				ReturnOcrData:  reducto.F(true),
				Timeout:        reducto.F(0.000000),
			}),
			Spreadsheet: reducto.F(shared.SpreadsheetParam{
				Clustering: reducto.F(shared.SpreadsheetClusteringAccurate),
				Exclude:    reducto.F([]shared.SpreadsheetExclude{shared.SpreadsheetExcludeHiddenSheets}),
				Include:    reducto.F([]shared.SpreadsheetInclude{shared.SpreadsheetIncludeCellColors}),
				SplitLargeTables: reducto.F(shared.SplitLargeTablesParam{
					Enabled: reducto.F(true),
					Size:    reducto.F[shared.SplitLargeTablesSizeUnionParam](shared.UnionInt(int64(0))),
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
