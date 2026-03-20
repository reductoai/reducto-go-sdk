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

func TestExtractNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Extract.New(context.TODO(), reducto.ExtractNewParams{
		Body: reducto.ExtractNewParamsBodySyncExtractConfig{
			Input: reducto.F[reducto.ExtractNewParamsBodySyncExtractConfigInputUnion](shared.UnionString("string")),
			Instructions: reducto.F(reducto.InstructionsParam{
				Schema:       reducto.F[any](map[string]interface{}{}),
				SystemPrompt: reducto.F("system_prompt"),
			}),
			Parsing: reducto.F(reducto.ParseOptionsParam{
				Enhance: reducto.F(reducto.EnhanceParam{
					Agentic: reducto.F([]reducto.EnhanceAgenticUnionParam{reducto.EnhanceAgenticTableAgenticParam{
						Scope:  reducto.F(reducto.EnhanceAgenticTableAgenticScopeTable),
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
					Chunking: reducto.F(reducto.RetrievalChunkingParam{
						ChunkMode:    reducto.F(reducto.RetrievalChunkingChunkModeVariable),
						ChunkOverlap: reducto.F(int64(0)),
						ChunkSize:    reducto.F(int64(0)),
					}),
					EmbeddingOptimized: reducto.F(true),
					FilterBlocks:       reducto.F([]reducto.RetrievalFilterBlock{reducto.RetrievalFilterBlockHeader}),
				}),
				Settings: reducto.F(reducto.SettingsParam{
					DocumentPassword:   reducto.F("document_password"),
					EmbedPdfMetadata:   reducto.F(true),
					ExtractionMode:     reducto.F(reducto.SettingsExtractionModeOcr),
					ForceFileExtension: reducto.F("force_file_extension"),
					ForceURLResult:     reducto.F(true),
					OcrSystem:          reducto.F(reducto.SettingsOcrSystemStandard),
					PageRange: reducto.F[reducto.SettingsPageRangeUnionParam](reducto.PageRangeParam{
						End:   reducto.F(int64(0)),
						Start: reducto.F(int64(0)),
					}),
					PersistResults: reducto.F(true),
					ReturnImages:   reducto.F([]reducto.SettingsReturnImage{reducto.SettingsReturnImageFigure}),
					ReturnOcrData:  reducto.F(true),
					Timeout:        reducto.F(0.000000),
				}),
				Spreadsheet: reducto.F(reducto.SpreadsheetParam{
					Clustering: reducto.F(reducto.SpreadsheetClusteringAccurate),
					Exclude:    reducto.F([]reducto.SpreadsheetExclude{reducto.SpreadsheetExcludeHiddenSheets}),
					Include:    reducto.F([]reducto.SpreadsheetInclude{reducto.SpreadsheetIncludeCellColors}),
					SplitLargeTables: reducto.F(reducto.SpreadsheetSplitLargeTablesParam{
						Enabled: reducto.F(true),
						Size:    reducto.F[reducto.SpreadsheetSplitLargeTablesSizeUnionParam](shared.UnionInt(int64(0))),
					}),
				}),
			}),
			Settings: reducto.F(reducto.ExtractSettingsParam{
				ArrayExtract: reducto.F(true),
				Citations: reducto.F(reducto.ExtractSettingsCitationsParam{
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
