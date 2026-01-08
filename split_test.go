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

func TestSplitRunWithOptionalParams(t *testing.T) {
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
	_, err := client.Split.Run(context.TODO(), reducto.SplitRunParams{
		Input: reducto.F[reducto.SplitRunParamsInputUnion](shared.UnionString("string")),
		SplitDescription: reducto.F([]shared.SplitCategoryParam{{
			Description:  reducto.F("description"),
			Name:         reducto.F("name"),
			PartitionKey: reducto.F("partition_key"),
		}}),
		Parsing: reducto.F(shared.ParseOptionsParam{
			Enhance: reducto.F(shared.EnhanceParam{
				Agentic: reducto.F([]shared.EnhanceAgenticUnionParam{shared.TableAgenticParam{
					Scope:  reducto.F(shared.TableAgenticScopeTable),
					Prompt: reducto.F("prompt"),
				}}),
				SummarizeFigures: reducto.F(true),
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
		Settings: reducto.F(reducto.SplitRunParamsSettings{
			TableCutoff: reducto.F(reducto.SplitRunParamsSettingsTableCutoffTruncate),
		}),
		SplitRules: reducto.F("split_rules"),
	})
	if err != nil {
		var apierr *reducto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSplitRunJobWithOptionalParams(t *testing.T) {
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
	_, err := client.Split.RunJob(context.TODO(), reducto.SplitRunJobParams{
		Input: reducto.F[reducto.SplitRunJobParamsInputUnion](shared.UnionString("string")),
		SplitDescription: reducto.F([]shared.SplitCategoryParam{{
			Description:  reducto.F("description"),
			Name:         reducto.F("name"),
			PartitionKey: reducto.F("partition_key"),
		}}),
		Async: reducto.F(shared.ConfigV3AsyncConfigParam{
			Metadata: reducto.F[any](map[string]interface{}{}),
			Priority: reducto.F(true),
			Webhook: reducto.F[shared.ConfigV3AsyncConfigWebhookUnionParam](shared.SvixWebhookConfigParam{
				Channels: reducto.F([]string{"string"}),
				Mode:     reducto.F(shared.SvixWebhookConfigModeSvix),
			}),
		}),
		Parsing: reducto.F(shared.ParseOptionsParam{
			Enhance: reducto.F(shared.EnhanceParam{
				Agentic: reducto.F([]shared.EnhanceAgenticUnionParam{shared.TableAgenticParam{
					Scope:  reducto.F(shared.TableAgenticScopeTable),
					Prompt: reducto.F("prompt"),
				}}),
				SummarizeFigures: reducto.F(true),
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
		Settings: reducto.F(reducto.SplitRunJobParamsSettings{
			TableCutoff: reducto.F(reducto.SplitRunJobParamsSettingsTableCutoffTruncate),
		}),
		SplitRules: reducto.F("split_rules"),
	})
	if err != nil {
		var apierr *reducto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
