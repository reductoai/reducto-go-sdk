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
		ExtractConfig: reducto.ExtractConfigParam{
			DocumentURL: reducto.F[reducto.ExtractConfigDocumentURLUnionParam](shared.UnionString("string")),
			Schema:      reducto.F[any](map[string]interface{}{}),
			AdvancedOptions: reducto.F(shared.AdvancedProcessingOptionsParam{
				AddPageMarkers:            reducto.F(true),
				ContinueHierarchy:         reducto.F(true),
				DocumentPassword:          reducto.F("document_password"),
				EnableChangeTracking:      reducto.F(true),
				EnableHighlightDetection:  reducto.F(true),
				ExcludeHiddenRowsCols:     reducto.F(true),
				ExcludeHiddenSheets:       reducto.F(true),
				FilterLineNumbers:         reducto.F(true),
				ForceFileExtension:        reducto.F("force_file_extension"),
				IncludeColorInformation:   reducto.F(true),
				IncludeFormulaInformation: reducto.F(true),
				KeepLineBreaks:            reducto.F(true),
				LargeTableChunking: reducto.F(shared.LargeTableChunkingConfigParam{
					Enabled: reducto.F(true),
					Size:    reducto.F(int64(0)),
				}),
				MergeTables: reducto.F(true),
				OcrSystem:   reducto.F(shared.AdvancedProcessingOptionsOcrSystemHighres),
				PageRange: reducto.F[shared.AdvancedProcessingOptionsPageRangeUnionParam](shared.PageRangeParam{
					End:   reducto.F(int64(0)),
					Start: reducto.F(int64(0)),
				}),
				PersistResults:             reducto.F(true),
				ReadComments:               reducto.F(true),
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
			CitationsOptions: reducto.F(shared.AdvancedCitationsConfigParam{
				NumericalConfidence: reducto.F(true),
			}),
			ExperimentalOptions: reducto.F(shared.ExperimentalProcessingOptionsParam{
				DangerFilterWideBoxes: reducto.F(true),
				DetectSignatures:      reducto.F(true),
				EmbedTextMetadataPdf:  reducto.F(true),
				EnableCheckboxes:      reducto.F(true),
				EnableEquations:       reducto.F(true),
				EnableScripts:         reducto.F(true),
				Enrich: reducto.F(shared.EnrichConfigParam{
					Enabled: reducto.F(true),
					Mode:    reducto.F(shared.EnrichConfigModeStandard),
					Prompt:  reducto.F("prompt"),
				}),
				LayoutModel:                 reducto.F(shared.ExperimentalProcessingOptionsLayoutModelDefault),
				NativeOfficeConversion:      reducto.F(true),
				ReturnFigureImages:          reducto.F(true),
				ReturnTableImages:           reducto.F(true),
				RotateFigures:               reducto.F(true),
				RotatePages:                 reducto.F(true),
				UserSpecifiedTimeoutSeconds: reducto.F(0.000000),
			}),
			ExperimentalTableCitations: reducto.F(true),
			GenerateCitations:          reducto.F(true),
			IncludeImages:              reducto.F(true),
			Options: reducto.F(shared.BaseProcessingOptionsParam{
				Chunking: reducto.F(shared.ChunkingConfigParam{
					ChunkMode: reducto.F(shared.ChunkingConfigChunkModeVariable),
					ChunkSize: reducto.F(int64(0)),
				}),
				ExtractionMode: reducto.F(shared.BaseProcessingOptionsExtractionModeOcr),
				FigureSummary: reducto.F(shared.FigureSummaryConfigParam{
					Enabled:  reducto.F(true),
					Override: reducto.F(true),
					Prompt:   reducto.F("prompt"),
				}),
				FilterBlocks:   reducto.F([]shared.BaseProcessingOptionsFilterBlock{shared.BaseProcessingOptionsFilterBlockHeader}),
				ForceURLResult: reducto.F(true),
				OcrMode:        reducto.F(shared.BaseProcessingOptionsOcrModeStandard),
				TableSummary: reducto.F(shared.TableSummaryConfigParam{
					Enabled: reducto.F(true),
					Prompt:  reducto.F("prompt"),
				}),
			}),
			Priority:         reducto.F(true),
			SpreadsheetAgent: reducto.F(true),
			SystemPrompt:     reducto.F("system_prompt"),
			UseChunking:      reducto.F(true),
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
		DocumentURL: reducto.F[reducto.ExtractRunJobParamsDocumentURLUnion](shared.UnionString("string")),
		Schema:      reducto.F[any](map[string]interface{}{}),
		AdvancedOptions: reducto.F(shared.AdvancedProcessingOptionsParam{
			AddPageMarkers:            reducto.F(true),
			ContinueHierarchy:         reducto.F(true),
			DocumentPassword:          reducto.F("document_password"),
			EnableChangeTracking:      reducto.F(true),
			EnableHighlightDetection:  reducto.F(true),
			ExcludeHiddenRowsCols:     reducto.F(true),
			ExcludeHiddenSheets:       reducto.F(true),
			FilterLineNumbers:         reducto.F(true),
			ForceFileExtension:        reducto.F("force_file_extension"),
			IncludeColorInformation:   reducto.F(true),
			IncludeFormulaInformation: reducto.F(true),
			KeepLineBreaks:            reducto.F(true),
			LargeTableChunking: reducto.F(shared.LargeTableChunkingConfigParam{
				Enabled: reducto.F(true),
				Size:    reducto.F(int64(0)),
			}),
			MergeTables: reducto.F(true),
			OcrSystem:   reducto.F(shared.AdvancedProcessingOptionsOcrSystemHighres),
			PageRange: reducto.F[shared.AdvancedProcessingOptionsPageRangeUnionParam](shared.PageRangeParam{
				End:   reducto.F(int64(0)),
				Start: reducto.F(int64(0)),
			}),
			PersistResults:             reducto.F(true),
			ReadComments:               reducto.F(true),
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
		CitationsOptions: reducto.F(shared.AdvancedCitationsConfigParam{
			NumericalConfidence: reducto.F(true),
		}),
		ExperimentalOptions: reducto.F(shared.ExperimentalProcessingOptionsParam{
			DangerFilterWideBoxes: reducto.F(true),
			DetectSignatures:      reducto.F(true),
			EmbedTextMetadataPdf:  reducto.F(true),
			EnableCheckboxes:      reducto.F(true),
			EnableEquations:       reducto.F(true),
			EnableScripts:         reducto.F(true),
			Enrich: reducto.F(shared.EnrichConfigParam{
				Enabled: reducto.F(true),
				Mode:    reducto.F(shared.EnrichConfigModeStandard),
				Prompt:  reducto.F("prompt"),
			}),
			LayoutModel:                 reducto.F(shared.ExperimentalProcessingOptionsLayoutModelDefault),
			NativeOfficeConversion:      reducto.F(true),
			ReturnFigureImages:          reducto.F(true),
			ReturnTableImages:           reducto.F(true),
			RotateFigures:               reducto.F(true),
			RotatePages:                 reducto.F(true),
			UserSpecifiedTimeoutSeconds: reducto.F(0.000000),
		}),
		ExperimentalTableCitations: reducto.F(true),
		GenerateCitations:          reducto.F(true),
		IncludeImages:              reducto.F(true),
		Options: reducto.F(shared.BaseProcessingOptionsParam{
			Chunking: reducto.F(shared.ChunkingConfigParam{
				ChunkMode: reducto.F(shared.ChunkingConfigChunkModeVariable),
				ChunkSize: reducto.F(int64(0)),
			}),
			ExtractionMode: reducto.F(shared.BaseProcessingOptionsExtractionModeOcr),
			FigureSummary: reducto.F(shared.FigureSummaryConfigParam{
				Enabled:  reducto.F(true),
				Override: reducto.F(true),
				Prompt:   reducto.F("prompt"),
			}),
			FilterBlocks:   reducto.F([]shared.BaseProcessingOptionsFilterBlock{shared.BaseProcessingOptionsFilterBlockHeader}),
			ForceURLResult: reducto.F(true),
			OcrMode:        reducto.F(shared.BaseProcessingOptionsOcrModeStandard),
			TableSummary: reducto.F(shared.TableSummaryConfigParam{
				Enabled: reducto.F(true),
				Prompt:  reducto.F("prompt"),
			}),
		}),
		Priority:         reducto.F(true),
		SpreadsheetAgent: reducto.F(true),
		SystemPrompt:     reducto.F("system_prompt"),
		UseChunking:      reducto.F(true),
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
