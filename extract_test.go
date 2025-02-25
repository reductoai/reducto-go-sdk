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
	t.Skip("skipped: tests are disabled for the time being")
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
		DocumentURL: reducto.F("document_url"),
		Schema:      reducto.F[any](map[string]interface{}{}),
		AdvancedOptions: reducto.F(shared.AdvancedProcessingOptionsParam{
			AddPageMarkers:     reducto.F(true),
			ContinueHierarchy:  reducto.F(true),
			DocumentPassword:   reducto.F("document_password"),
			ForceFileExtension: reducto.F("force_file_extension"),
			KeepLineBreaks:     reducto.F(true),
			LargeTableChunking: reducto.F(shared.AdvancedProcessingOptionsLargeTableChunkingParam{
				Enabled: reducto.F(true),
				Size:    reducto.F(int64(0)),
			}),
			MergeTables: reducto.F(true),
			OcrSystem:   reducto.F(shared.AdvancedProcessingOptionsOcrSystemHighres),
			PageRange: reducto.F[shared.AdvancedProcessingOptionsPageRangeUnionParam](shared.PageRangeParam{
				End:   reducto.F(int64(0)),
				Start: reducto.F(int64(0)),
			}),
			RemoveTextFormatting:       reducto.F(true),
			ReturnOcrData:              reducto.F(true),
			SpreadsheetTableClustering: reducto.F(shared.AdvancedProcessingOptionsSpreadsheetTableClusteringDefault),
			TableOutputFormat:          reducto.F(shared.AdvancedProcessingOptionsTableOutputFormatHTML),
		}),
		ArrayExtract: reducto.F(shared.ArrayExtractConfigParam{
			Enabled:                     reducto.F(true),
			Mode:                        reducto.F(shared.ArrayExtractConfigModeAuto),
			PagesPerSegment:             reducto.F(int64(0)),
			StreamingExtractItemDensity: reducto.F(int64(0)),
		}),
		ExperimentalOptions: reducto.F(shared.ExperimentalProcessingOptionsParam{
			DangerFilterWideBoxes: reducto.F(true),
			EnableCheckboxes:      reducto.F(true),
			EnableEquations:       reducto.F(true),
			EnableScripts:         reducto.F(true),
			EnableUnderlines:      reducto.F(true),
			Enrich: reducto.F(shared.ExperimentalProcessingOptionsEnrichParam{
				Enabled: reducto.F(true),
				Prompt:  reducto.F("prompt"),
			}),
			NativeOfficeConversion: reducto.F(true),
			ReturnFigureImages:     reducto.F(true),
			RotatePages:            reducto.F(true),
		}),
		GenerateCitations: reducto.F(true),
		Options: reducto.F(shared.BaseProcessingOptionsParam{
			Chunking: reducto.F(shared.BaseProcessingOptionsChunkingParam{
				ChunkMode: reducto.F(shared.BaseProcessingOptionsChunkingChunkModeVariable),
				ChunkSize: reducto.F(int64(0)),
			}),
			ExtractionMode: reducto.F(shared.BaseProcessingOptionsExtractionModeOcr),
			FigureSummary: reducto.F(shared.BaseProcessingOptionsFigureSummaryParam{
				Enabled:  reducto.F(true),
				Override: reducto.F(true),
				Prompt:   reducto.F("prompt"),
			}),
			FilterBlocks:   reducto.F([]shared.BaseProcessingOptionsFilterBlock{shared.BaseProcessingOptionsFilterBlockHeader}),
			ForceURLResult: reducto.F(true),
			TableSummary: reducto.F(shared.BaseProcessingOptionsTableSummaryParam{
				Enabled: reducto.F(true),
				Prompt:  reducto.F("prompt"),
			}),
		}),
		SystemPrompt: reducto.F("system_prompt"),
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
	t.Skip("skipped: tests are disabled for the time being")
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
		DocumentURL: reducto.F("document_url"),
		Schema:      reducto.F[any](map[string]interface{}{}),
		AdvancedOptions: reducto.F(shared.AdvancedProcessingOptionsParam{
			AddPageMarkers:     reducto.F(true),
			ContinueHierarchy:  reducto.F(true),
			DocumentPassword:   reducto.F("document_password"),
			ForceFileExtension: reducto.F("force_file_extension"),
			KeepLineBreaks:     reducto.F(true),
			LargeTableChunking: reducto.F(shared.AdvancedProcessingOptionsLargeTableChunkingParam{
				Enabled: reducto.F(true),
				Size:    reducto.F(int64(0)),
			}),
			MergeTables: reducto.F(true),
			OcrSystem:   reducto.F(shared.AdvancedProcessingOptionsOcrSystemHighres),
			PageRange: reducto.F[shared.AdvancedProcessingOptionsPageRangeUnionParam](shared.PageRangeParam{
				End:   reducto.F(int64(0)),
				Start: reducto.F(int64(0)),
			}),
			RemoveTextFormatting:       reducto.F(true),
			ReturnOcrData:              reducto.F(true),
			SpreadsheetTableClustering: reducto.F(shared.AdvancedProcessingOptionsSpreadsheetTableClusteringDefault),
			TableOutputFormat:          reducto.F(shared.AdvancedProcessingOptionsTableOutputFormatHTML),
		}),
		ArrayExtract: reducto.F(shared.ArrayExtractConfigParam{
			Enabled:                     reducto.F(true),
			Mode:                        reducto.F(shared.ArrayExtractConfigModeAuto),
			PagesPerSegment:             reducto.F(int64(0)),
			StreamingExtractItemDensity: reducto.F(int64(0)),
		}),
		ExperimentalOptions: reducto.F(shared.ExperimentalProcessingOptionsParam{
			DangerFilterWideBoxes: reducto.F(true),
			EnableCheckboxes:      reducto.F(true),
			EnableEquations:       reducto.F(true),
			EnableScripts:         reducto.F(true),
			EnableUnderlines:      reducto.F(true),
			Enrich: reducto.F(shared.ExperimentalProcessingOptionsEnrichParam{
				Enabled: reducto.F(true),
				Prompt:  reducto.F("prompt"),
			}),
			NativeOfficeConversion: reducto.F(true),
			ReturnFigureImages:     reducto.F(true),
			RotatePages:            reducto.F(true),
		}),
		GenerateCitations: reducto.F(true),
		Options: reducto.F(shared.BaseProcessingOptionsParam{
			Chunking: reducto.F(shared.BaseProcessingOptionsChunkingParam{
				ChunkMode: reducto.F(shared.BaseProcessingOptionsChunkingChunkModeVariable),
				ChunkSize: reducto.F(int64(0)),
			}),
			ExtractionMode: reducto.F(shared.BaseProcessingOptionsExtractionModeOcr),
			FigureSummary: reducto.F(shared.BaseProcessingOptionsFigureSummaryParam{
				Enabled:  reducto.F(true),
				Override: reducto.F(true),
				Prompt:   reducto.F("prompt"),
			}),
			FilterBlocks:   reducto.F([]shared.BaseProcessingOptionsFilterBlock{shared.BaseProcessingOptionsFilterBlockHeader}),
			ForceURLResult: reducto.F(true),
			TableSummary: reducto.F(shared.BaseProcessingOptionsTableSummaryParam{
				Enabled: reducto.F(true),
				Prompt:  reducto.F("prompt"),
			}),
		}),
		Priority:     reducto.F(true),
		SystemPrompt: reducto.F("system_prompt"),
		Webhook: reducto.F(shared.WebhookConfigNewParam{
			Channels: reducto.F([]string{"string"}),
			Metadata: reducto.F[any](map[string]interface{}{}),
			Mode:     reducto.F(shared.WebhookConfigNewModeDisabled),
			URL:      reducto.F("url"),
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
