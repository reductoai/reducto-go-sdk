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
	_, err := client.Extract.Run(context.TODO(), reducto.ExtractRunParams{
		Body: reducto.ExtractRunParamsBodySyncExtractConfig{
			Input: reducto.F[reducto.ExtractRunParamsBodySyncExtractConfigInputUnion](shared.UnionString("string")),
			Instructions: reducto.F(reducto.InstructionsParam{
				Schema:       reducto.F[any](map[string]interface{}{}),
				SystemPrompt: reducto.F("system_prompt"),
			}),
			Parsing: reducto.F(reducto.ParseOptionsParam{
				Enhance: reducto.F(reducto.EnhanceParam{
					Agentic: reducto.F([]reducto.EnhanceAgenticUnionParam{shared.TableAgenticParam{
						Scope:  reducto.F(shared.TableAgenticScopeTable),
						Mode:   reducto.F(shared.TableAgenticModeDefault),
						Prompt: reducto.F("prompt"),
					}}),
					IntelligentOrdering: reducto.F(true),
					SummarizeFigures:    reducto.F(true),
				}),
				Formatting: reducto.F(reducto.FormattingParam{
					AddPageMarkers:    reducto.F(true),
					Include:           reducto.F([]reducto.FormattingInclude{reducto.FormattingIncludeChangeTracking}),
					MergeTables:       reducto.F(true),
					TableOutputFormat: reducto.F(reducto.FormattingTableOutputFormatHTML),
				}),
				Retrieval: reducto.F(reducto.RetrievalParam{
					Chunking: reducto.F(shared.ChunkingParam{
						ChunkMode:    reducto.F(shared.ChunkingChunkModeVariable),
						ChunkOverlap: reducto.F(int64(0)),
						ChunkSize:    reducto.F(int64(0)),
					}),
					EmbeddingOptimized: reducto.F(true),
					FilterBlocks:       reducto.F([]reducto.RetrievalFilterBlock{reducto.RetrievalFilterBlockHeader}),
				}),
				Settings: reducto.F(reducto.SettingsParam{
					DocumentPassword:    reducto.F("document_password"),
					EmbedPdfMetadata:    reducto.F(true),
					EmbedPdfMetadataDpi: reducto.F(int64(50)),
					ExtractionMode:      reducto.F(reducto.SettingsExtractionModeOcr),
					ForceFileExtension:  reducto.F("force_file_extension"),
					ForceURLResult:      reducto.F(true),
					HybridVpc: reducto.F(reducto.SettingsHybridVpcParam{
						Environment: reducto.F("environment"),
					}),
					OcrSystem: reducto.F(reducto.SettingsOcrSystemStandard),
					PageRange: reducto.F[reducto.SettingsPageRangeUnionParam](shared.PageRangeParam{
						End:   reducto.F(int64(0)),
						Start: reducto.F(int64(0)),
					}),
					PersistResults: reducto.F(true),
					ReturnImages:   reducto.F([]reducto.SettingsReturnImage{reducto.SettingsReturnImageFigure}),
					ReturnOcrData:  reducto.F(true),
					Timeout:        reducto.F(0.000000),
				}),
				Spreadsheet: reducto.F(reducto.SpreadsheetParam{
					Clustering:   reducto.F(reducto.SpreadsheetClusteringAccurate),
					Exclude:      reducto.F([]reducto.SpreadsheetExclude{reducto.SpreadsheetExcludeHiddenSheets}),
					Include:      reducto.F([]reducto.SpreadsheetInclude{reducto.SpreadsheetIncludeCellColors}),
					MaxCellCount: reducto.F(int64(1)),
					SplitLargeTables: reducto.F(shared.SplitLargeTablesParam{
						Enabled: reducto.F(true),
						Size:    reducto.F[shared.SplitLargeTablesSizeUnionParam](shared.UnionInt(int64(0))),
					}),
				}),
			}),
			Settings: reducto.F(reducto.ExtractSettingsParam{
				ArrayExtract: reducto.F(true),
				Citations: reducto.F(reducto.ExtractSettingsCitationsParam{
					Enabled:             reducto.F(true),
					NumericalConfidence: reducto.F(true),
				}),
				DeepExtract:        reducto.F(true),
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
	_, err := client.Extract.RunJob(context.TODO(), reducto.ExtractRunJobParams{
		AsyncExtractConfig: reducto.AsyncExtractConfigParam{
			Input: reducto.F[reducto.AsyncExtractConfigInputUnionParam](shared.UnionString("string")),
			Async: reducto.F(reducto.AsyncConfigV3Param{
				Metadata: reducto.F[any](map[string]interface{}{}),
				Priority: reducto.F(true),
				Webhook: reducto.F[reducto.AsyncConfigV3WebhookUnionParam](shared.SvixWebhookConfigParam{
					Channels: reducto.F([]string{"string"}),
					Mode:     reducto.F(shared.SvixWebhookConfigModeSvix),
				}),
			}),
			Instructions: reducto.F(reducto.InstructionsParam{
				Schema:       reducto.F[any](map[string]interface{}{}),
				SystemPrompt: reducto.F("system_prompt"),
			}),
			Parsing: reducto.F(reducto.ParseOptionsParam{
				Enhance: reducto.F(reducto.EnhanceParam{
					Agentic: reducto.F([]reducto.EnhanceAgenticUnionParam{shared.TableAgenticParam{
						Scope:  reducto.F(shared.TableAgenticScopeTable),
						Mode:   reducto.F(shared.TableAgenticModeDefault),
						Prompt: reducto.F("prompt"),
					}}),
					IntelligentOrdering: reducto.F(true),
					SummarizeFigures:    reducto.F(true),
				}),
				Formatting: reducto.F(reducto.FormattingParam{
					AddPageMarkers:    reducto.F(true),
					Include:           reducto.F([]reducto.FormattingInclude{reducto.FormattingIncludeChangeTracking}),
					MergeTables:       reducto.F(true),
					TableOutputFormat: reducto.F(reducto.FormattingTableOutputFormatHTML),
				}),
				Retrieval: reducto.F(reducto.RetrievalParam{
					Chunking: reducto.F(shared.ChunkingParam{
						ChunkMode:    reducto.F(shared.ChunkingChunkModeVariable),
						ChunkOverlap: reducto.F(int64(0)),
						ChunkSize:    reducto.F(int64(0)),
					}),
					EmbeddingOptimized: reducto.F(true),
					FilterBlocks:       reducto.F([]reducto.RetrievalFilterBlock{reducto.RetrievalFilterBlockHeader}),
				}),
				Settings: reducto.F(reducto.SettingsParam{
					DocumentPassword:    reducto.F("document_password"),
					EmbedPdfMetadata:    reducto.F(true),
					EmbedPdfMetadataDpi: reducto.F(int64(50)),
					ExtractionMode:      reducto.F(reducto.SettingsExtractionModeOcr),
					ForceFileExtension:  reducto.F("force_file_extension"),
					ForceURLResult:      reducto.F(true),
					HybridVpc: reducto.F(reducto.SettingsHybridVpcParam{
						Environment: reducto.F("environment"),
					}),
					OcrSystem: reducto.F(reducto.SettingsOcrSystemStandard),
					PageRange: reducto.F[reducto.SettingsPageRangeUnionParam](shared.PageRangeParam{
						End:   reducto.F(int64(0)),
						Start: reducto.F(int64(0)),
					}),
					PersistResults: reducto.F(true),
					ReturnImages:   reducto.F([]reducto.SettingsReturnImage{reducto.SettingsReturnImageFigure}),
					ReturnOcrData:  reducto.F(true),
					Timeout:        reducto.F(0.000000),
				}),
				Spreadsheet: reducto.F(reducto.SpreadsheetParam{
					Clustering:   reducto.F(reducto.SpreadsheetClusteringAccurate),
					Exclude:      reducto.F([]reducto.SpreadsheetExclude{reducto.SpreadsheetExcludeHiddenSheets}),
					Include:      reducto.F([]reducto.SpreadsheetInclude{reducto.SpreadsheetIncludeCellColors}),
					MaxCellCount: reducto.F(int64(1)),
					SplitLargeTables: reducto.F(shared.SplitLargeTablesParam{
						Enabled: reducto.F(true),
						Size:    reducto.F[shared.SplitLargeTablesSizeUnionParam](shared.UnionInt(int64(0))),
					}),
				}),
			}),
			Settings: reducto.F(reducto.ExtractSettingsParam{
				ArrayExtract: reducto.F(true),
				Citations: reducto.F(reducto.ExtractSettingsCitationsParam{
					Enabled:             reducto.F(true),
					NumericalConfidence: reducto.F(true),
				}),
				DeepExtract:        reducto.F(true),
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
