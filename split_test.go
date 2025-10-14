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
		Parsing: reducto.F(reducto.SplitRunParamsParsing{
			Enhance: reducto.F(reducto.SplitRunParamsParsingEnhance{
				Agentic: reducto.F([]reducto.SplitRunParamsParsingEnhanceAgenticUnion{reducto.SplitRunParamsParsingEnhanceAgenticTableAgentic{
					Scope:  reducto.F(reducto.SplitRunParamsParsingEnhanceAgenticTableAgenticScopeTable),
					Prompt: reducto.F("prompt"),
				}}),
				SummarizeFigures: reducto.F(true),
			}),
			Formatting: reducto.F(reducto.SplitRunParamsParsingFormatting{
				AddPageMarkers:    reducto.F(true),
				Include:           reducto.F([]reducto.SplitRunParamsParsingFormattingInclude{reducto.SplitRunParamsParsingFormattingIncludeChangeTracking}),
				MergeTables:       reducto.F(true),
				TableOutputFormat: reducto.F(reducto.SplitRunParamsParsingFormattingTableOutputFormatHTML),
			}),
			Retrieval: reducto.F(reducto.SplitRunParamsParsingRetrieval{
				Chunking: reducto.F(reducto.SplitRunParamsParsingRetrievalChunking{
					ChunkMode: reducto.F(reducto.SplitRunParamsParsingRetrievalChunkingChunkModeVariable),
					ChunkSize: reducto.F(int64(0)),
				}),
				EmbeddingOptimized: reducto.F(true),
				FilterBlocks:       reducto.F([]reducto.SplitRunParamsParsingRetrievalFilterBlock{reducto.SplitRunParamsParsingRetrievalFilterBlockHeader}),
			}),
			Settings: reducto.F(reducto.SplitRunParamsParsingSettings{
				DocumentPassword:   reducto.F("document_password"),
				EmbedPdfMetadata:   reducto.F(true),
				ForceFileExtension: reducto.F("force_file_extension"),
				ForceURLResult:     reducto.F(true),
				OcrSystem:          reducto.F(reducto.SplitRunParamsParsingSettingsOcrSystemStandard),
				PageRange: reducto.F[reducto.SplitRunParamsParsingSettingsPageRangeUnion](shared.PageRangeParam{
					End:   reducto.F(int64(0)),
					Start: reducto.F(int64(0)),
				}),
				PersistResults: reducto.F(true),
				ReturnImages:   reducto.F([]reducto.SplitRunParamsParsingSettingsReturnImage{reducto.SplitRunParamsParsingSettingsReturnImageFigure}),
				ReturnOcrData:  reducto.F(true),
				Timeout:        reducto.F(0.000000),
			}),
			Spreadsheet: reducto.F(reducto.SplitRunParamsParsingSpreadsheet{
				Clustering: reducto.F(reducto.SplitRunParamsParsingSpreadsheetClusteringAccurate),
				Exclude:    reducto.F([]reducto.SplitRunParamsParsingSpreadsheetExclude{reducto.SplitRunParamsParsingSpreadsheetExcludeHiddenSheets}),
				Include:    reducto.F([]reducto.SplitRunParamsParsingSpreadsheetInclude{reducto.SplitRunParamsParsingSpreadsheetIncludeCellColors}),
				SplitLargeTables: reducto.F(reducto.SplitRunParamsParsingSpreadsheetSplitLargeTables{
					Enabled: reducto.F(true),
					Size:    reducto.F(int64(0)),
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

func TestSplitRunJob(t *testing.T) {
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
		Body: map[string]interface{}{},
	})
	if err != nil {
		var apierr *reducto.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
