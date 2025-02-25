// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reductoai_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/reductoai-go"
	"github.com/stainless-sdks/reductoai-go/internal/testutil"
	"github.com/stainless-sdks/reductoai-go/option"
	"github.com/stainless-sdks/reductoai-go/shared"
)

func TestSplitRunWithOptionalParams(t *testing.T) {
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
	_, err := client.Split.Run(context.TODO(), reductoai.SplitRunParams{
		DocumentURL: reductoai.F("document_url"),
		SplitDescription: reductoai.F([]shared.SplitCategoryParam{{
			Description:  reductoai.F("description"),
			Name:         reductoai.F("name"),
			PartitionKey: reductoai.F("partition_key"),
		}}),
		AdvancedOptions: reductoai.F(shared.AdvancedProcessingOptionsParam{
			AddPageMarkers:     reductoai.F(true),
			ContinueHierarchy:  reductoai.F(true),
			DocumentPassword:   reductoai.F("document_password"),
			ForceFileExtension: reductoai.F("force_file_extension"),
			KeepLineBreaks:     reductoai.F(true),
			LargeTableChunking: reductoai.F(shared.AdvancedProcessingOptionsLargeTableChunkingParam{
				Enabled: reductoai.F(true),
				Size:    reductoai.F(int64(0)),
			}),
			MergeTables: reductoai.F(true),
			OcrSystem:   reductoai.F(shared.AdvancedProcessingOptionsOcrSystemHighres),
			PageRange: reductoai.F[shared.AdvancedProcessingOptionsPageRangeUnionParam](shared.PageRangeParam{
				End:   reductoai.F(int64(0)),
				Start: reductoai.F(int64(0)),
			}),
			RemoveTextFormatting:       reductoai.F(true),
			ReturnOcrData:              reductoai.F(true),
			SpreadsheetTableClustering: reductoai.F(shared.AdvancedProcessingOptionsSpreadsheetTableClusteringDefault),
			TableOutputFormat:          reductoai.F(shared.AdvancedProcessingOptionsTableOutputFormatHTML),
		}),
		ExperimentalOptions: reductoai.F(shared.ExperimentalProcessingOptionsParam{
			DangerFilterWideBoxes: reductoai.F(true),
			EnableCheckboxes:      reductoai.F(true),
			EnableEquations:       reductoai.F(true),
			EnableScripts:         reductoai.F(true),
			EnableUnderlines:      reductoai.F(true),
			Enrich: reductoai.F(shared.ExperimentalProcessingOptionsEnrichParam{
				Enabled: reductoai.F(true),
				Prompt:  reductoai.F("prompt"),
			}),
			NativeOfficeConversion: reductoai.F(true),
			ReturnFigureImages:     reductoai.F(true),
			RotatePages:            reductoai.F(true),
		}),
		Options: reductoai.F(shared.BaseProcessingOptionsParam{
			Chunking: reductoai.F(shared.BaseProcessingOptionsChunkingParam{
				ChunkMode: reductoai.F(shared.BaseProcessingOptionsChunkingChunkModeVariable),
				ChunkSize: reductoai.F(int64(0)),
			}),
			ExtractionMode: reductoai.F(shared.BaseProcessingOptionsExtractionModeOcr),
			FigureSummary: reductoai.F(shared.BaseProcessingOptionsFigureSummaryParam{
				Enabled:  reductoai.F(true),
				Override: reductoai.F(true),
				Prompt:   reductoai.F("prompt"),
			}),
			FilterBlocks:   reductoai.F([]shared.BaseProcessingOptionsFilterBlock{shared.BaseProcessingOptionsFilterBlockHeader}),
			ForceURLResult: reductoai.F(true),
			TableSummary: reductoai.F(shared.BaseProcessingOptionsTableSummaryParam{
				Enabled: reductoai.F(true),
				Prompt:  reductoai.F("prompt"),
			}),
		}),
		SplitRules: reductoai.F("split_rules"),
	})
	if err != nil {
		var apierr *reductoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSplitRunJobWithOptionalParams(t *testing.T) {
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
	_, err := client.Split.RunJob(context.TODO(), reductoai.SplitRunJobParams{
		DocumentURL: reductoai.F("document_url"),
		SplitDescription: reductoai.F([]shared.SplitCategoryParam{{
			Description:  reductoai.F("description"),
			Name:         reductoai.F("name"),
			PartitionKey: reductoai.F("partition_key"),
		}}),
		AdvancedOptions: reductoai.F(shared.AdvancedProcessingOptionsParam{
			AddPageMarkers:     reductoai.F(true),
			ContinueHierarchy:  reductoai.F(true),
			DocumentPassword:   reductoai.F("document_password"),
			ForceFileExtension: reductoai.F("force_file_extension"),
			KeepLineBreaks:     reductoai.F(true),
			LargeTableChunking: reductoai.F(shared.AdvancedProcessingOptionsLargeTableChunkingParam{
				Enabled: reductoai.F(true),
				Size:    reductoai.F(int64(0)),
			}),
			MergeTables: reductoai.F(true),
			OcrSystem:   reductoai.F(shared.AdvancedProcessingOptionsOcrSystemHighres),
			PageRange: reductoai.F[shared.AdvancedProcessingOptionsPageRangeUnionParam](shared.PageRangeParam{
				End:   reductoai.F(int64(0)),
				Start: reductoai.F(int64(0)),
			}),
			RemoveTextFormatting:       reductoai.F(true),
			ReturnOcrData:              reductoai.F(true),
			SpreadsheetTableClustering: reductoai.F(shared.AdvancedProcessingOptionsSpreadsheetTableClusteringDefault),
			TableOutputFormat:          reductoai.F(shared.AdvancedProcessingOptionsTableOutputFormatHTML),
		}),
		ExperimentalOptions: reductoai.F(shared.ExperimentalProcessingOptionsParam{
			DangerFilterWideBoxes: reductoai.F(true),
			EnableCheckboxes:      reductoai.F(true),
			EnableEquations:       reductoai.F(true),
			EnableScripts:         reductoai.F(true),
			EnableUnderlines:      reductoai.F(true),
			Enrich: reductoai.F(shared.ExperimentalProcessingOptionsEnrichParam{
				Enabled: reductoai.F(true),
				Prompt:  reductoai.F("prompt"),
			}),
			NativeOfficeConversion: reductoai.F(true),
			ReturnFigureImages:     reductoai.F(true),
			RotatePages:            reductoai.F(true),
		}),
		Options: reductoai.F(shared.BaseProcessingOptionsParam{
			Chunking: reductoai.F(shared.BaseProcessingOptionsChunkingParam{
				ChunkMode: reductoai.F(shared.BaseProcessingOptionsChunkingChunkModeVariable),
				ChunkSize: reductoai.F(int64(0)),
			}),
			ExtractionMode: reductoai.F(shared.BaseProcessingOptionsExtractionModeOcr),
			FigureSummary: reductoai.F(shared.BaseProcessingOptionsFigureSummaryParam{
				Enabled:  reductoai.F(true),
				Override: reductoai.F(true),
				Prompt:   reductoai.F("prompt"),
			}),
			FilterBlocks:   reductoai.F([]shared.BaseProcessingOptionsFilterBlock{shared.BaseProcessingOptionsFilterBlockHeader}),
			ForceURLResult: reductoai.F(true),
			TableSummary: reductoai.F(shared.BaseProcessingOptionsTableSummaryParam{
				Enabled: reductoai.F(true),
				Prompt:  reductoai.F("prompt"),
			}),
		}),
		Priority:   reductoai.F(true),
		SplitRules: reductoai.F("split_rules"),
		Webhook: reductoai.F(shared.WebhookConfigNewParam{
			Channels: reductoai.F([]string{"string"}),
			Metadata: reductoai.F[any](map[string]interface{}{}),
			Mode:     reductoai.F(shared.WebhookConfigNewModeDisabled),
			URL:      reductoai.F("url"),
		}),
	})
	if err != nil {
		var apierr *reductoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
