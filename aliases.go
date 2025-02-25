// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reducto

import (
	"github.com/reductoai/reducto-go-sdk/internal/apierror"
	"github.com/reductoai/reducto-go-sdk/shared"
)

type Error = apierror.Error

// This is an alias to an internal type.
type AdvancedProcessingOptionsParam = shared.AdvancedProcessingOptionsParam

// The configuration options for large table chunking (currently only supported on
// spreadsheet and CSV files).
//
// This is an alias to an internal type.
type AdvancedProcessingOptionsLargeTableChunkingParam = shared.AdvancedProcessingOptionsLargeTableChunkingParam

// The OCR system to use. Highres is recommended for documents with English
// characters.
//
// This is an alias to an internal type.
type AdvancedProcessingOptionsOcrSystem = shared.AdvancedProcessingOptionsOcrSystem

// This is an alias to an internal value.
const AdvancedProcessingOptionsOcrSystemHighres = shared.AdvancedProcessingOptionsOcrSystemHighres

// This is an alias to an internal value.
const AdvancedProcessingOptionsOcrSystemMultilingual = shared.AdvancedProcessingOptionsOcrSystemMultilingual

// This is an alias to an internal value.
const AdvancedProcessingOptionsOcrSystemCombined = shared.AdvancedProcessingOptionsOcrSystemCombined

// The page range to process. By default, the entire document is processed.
//
// This is an alias to an internal type.
type AdvancedProcessingOptionsPageRangeUnionParam = shared.AdvancedProcessingOptionsPageRangeUnionParam

// This is an alias to an internal type.
type AdvancedProcessingOptionsPageRangeArrayParam = shared.AdvancedProcessingOptionsPageRangeArrayParam

// On a spreadsheet, the algorithm that is used to split up sheets into multiple
// tables.
//
// This is an alias to an internal type.
type AdvancedProcessingOptionsSpreadsheetTableClustering = shared.AdvancedProcessingOptionsSpreadsheetTableClustering

// This is an alias to an internal value.
const AdvancedProcessingOptionsSpreadsheetTableClusteringDefault = shared.AdvancedProcessingOptionsSpreadsheetTableClusteringDefault

// This is an alias to an internal value.
const AdvancedProcessingOptionsSpreadsheetTableClusteringDisabled = shared.AdvancedProcessingOptionsSpreadsheetTableClusteringDisabled

// The mode to use for table output. Dynamic returns md for simpler tables and html
// for more complex tables.
//
// This is an alias to an internal type.
type AdvancedProcessingOptionsTableOutputFormat = shared.AdvancedProcessingOptionsTableOutputFormat

// This is an alias to an internal value.
const AdvancedProcessingOptionsTableOutputFormatHTML = shared.AdvancedProcessingOptionsTableOutputFormatHTML

// This is an alias to an internal value.
const AdvancedProcessingOptionsTableOutputFormatJson = shared.AdvancedProcessingOptionsTableOutputFormatJson

// This is an alias to an internal value.
const AdvancedProcessingOptionsTableOutputFormatMd = shared.AdvancedProcessingOptionsTableOutputFormatMd

// This is an alias to an internal value.
const AdvancedProcessingOptionsTableOutputFormatJsonbbox = shared.AdvancedProcessingOptionsTableOutputFormatJsonbbox

// This is an alias to an internal value.
const AdvancedProcessingOptionsTableOutputFormatDynamic = shared.AdvancedProcessingOptionsTableOutputFormatDynamic

// This is an alias to an internal value.
const AdvancedProcessingOptionsTableOutputFormatAIJson = shared.AdvancedProcessingOptionsTableOutputFormatAIJson

// This is an alias to an internal type.
type ArrayExtractConfigParam = shared.ArrayExtractConfigParam

// The array extraction version to use.
//
// This is an alias to an internal type.
type ArrayExtractConfigMode = shared.ArrayExtractConfigMode

// This is an alias to an internal value.
const ArrayExtractConfigModeAuto = shared.ArrayExtractConfigModeAuto

// This is an alias to an internal value.
const ArrayExtractConfigModeLegacy = shared.ArrayExtractConfigModeLegacy

// This is an alias to an internal value.
const ArrayExtractConfigModeStreaming = shared.ArrayExtractConfigModeStreaming

// This is an alias to an internal value.
const ArrayExtractConfigModeNoOverlap = shared.ArrayExtractConfigModeNoOverlap

// This is an alias to an internal type.
type BaseProcessingOptionsParam = shared.BaseProcessingOptionsParam

// The configuration options for chunking.
//
// This is an alias to an internal type.
type BaseProcessingOptionsChunkingParam = shared.BaseProcessingOptionsChunkingParam

// The mode to use for chunking. Section chunks according to sections in the
// document. Page chunks according to pages. Disabled returns a single chunk.
//
// This is an alias to an internal type.
type BaseProcessingOptionsChunkingChunkMode = shared.BaseProcessingOptionsChunkingChunkMode

// This is an alias to an internal value.
const BaseProcessingOptionsChunkingChunkModeVariable = shared.BaseProcessingOptionsChunkingChunkModeVariable

// This is an alias to an internal value.
const BaseProcessingOptionsChunkingChunkModeSection = shared.BaseProcessingOptionsChunkingChunkModeSection

// This is an alias to an internal value.
const BaseProcessingOptionsChunkingChunkModePage = shared.BaseProcessingOptionsChunkingChunkModePage

// This is an alias to an internal value.
const BaseProcessingOptionsChunkingChunkModeBlock = shared.BaseProcessingOptionsChunkingChunkModeBlock

// This is an alias to an internal value.
const BaseProcessingOptionsChunkingChunkModeDisabled = shared.BaseProcessingOptionsChunkingChunkModeDisabled

// The mode to use for extraction.
//
// This is an alias to an internal type.
type BaseProcessingOptionsExtractionMode = shared.BaseProcessingOptionsExtractionMode

// This is an alias to an internal value.
const BaseProcessingOptionsExtractionModeOcr = shared.BaseProcessingOptionsExtractionModeOcr

// This is an alias to an internal value.
const BaseProcessingOptionsExtractionModeMetadata = shared.BaseProcessingOptionsExtractionModeMetadata

// This is an alias to an internal value.
const BaseProcessingOptionsExtractionModeHybrid = shared.BaseProcessingOptionsExtractionModeHybrid

// The configuration options for figure summarization.
//
// This is an alias to an internal type.
type BaseProcessingOptionsFigureSummaryParam = shared.BaseProcessingOptionsFigureSummaryParam

// This is an alias to an internal type.
type BaseProcessingOptionsFilterBlock = shared.BaseProcessingOptionsFilterBlock

// This is an alias to an internal value.
const BaseProcessingOptionsFilterBlockHeader = shared.BaseProcessingOptionsFilterBlockHeader

// This is an alias to an internal value.
const BaseProcessingOptionsFilterBlockFooter = shared.BaseProcessingOptionsFilterBlockFooter

// This is an alias to an internal value.
const BaseProcessingOptionsFilterBlockTitle = shared.BaseProcessingOptionsFilterBlockTitle

// This is an alias to an internal value.
const BaseProcessingOptionsFilterBlockSectionHeader = shared.BaseProcessingOptionsFilterBlockSectionHeader

// This is an alias to an internal value.
const BaseProcessingOptionsFilterBlockPageNumber = shared.BaseProcessingOptionsFilterBlockPageNumber

// This is an alias to an internal value.
const BaseProcessingOptionsFilterBlockListItem = shared.BaseProcessingOptionsFilterBlockListItem

// This is an alias to an internal value.
const BaseProcessingOptionsFilterBlockFigure = shared.BaseProcessingOptionsFilterBlockFigure

// This is an alias to an internal value.
const BaseProcessingOptionsFilterBlockTable = shared.BaseProcessingOptionsFilterBlockTable

// This is an alias to an internal value.
const BaseProcessingOptionsFilterBlockKeyValue = shared.BaseProcessingOptionsFilterBlockKeyValue

// This is an alias to an internal value.
const BaseProcessingOptionsFilterBlockText = shared.BaseProcessingOptionsFilterBlockText

// This is an alias to an internal value.
const BaseProcessingOptionsFilterBlockComment = shared.BaseProcessingOptionsFilterBlockComment

// This is an alias to an internal value.
const BaseProcessingOptionsFilterBlockDiscard = shared.BaseProcessingOptionsFilterBlockDiscard

// The configuration options for table summarization.
//
// This is an alias to an internal type.
type BaseProcessingOptionsTableSummaryParam = shared.BaseProcessingOptionsTableSummaryParam

// This is an alias to an internal type.
type BoundingBox = shared.BoundingBox

// This is an alias to an internal type.
type ExperimentalProcessingOptionsParam = shared.ExperimentalProcessingOptionsParam

// The configuration options for enrichment.
//
// This is an alias to an internal type.
type ExperimentalProcessingOptionsEnrichParam = shared.ExperimentalProcessingOptionsEnrichParam

// This is an alias to an internal type.
type ExtractResponse = shared.ExtractResponse

// This is an alias to an internal type.
type ExtractResponseUsage = shared.ExtractResponseUsage

// This is an alias to an internal type.
type PageRangeParam = shared.PageRangeParam

// This is an alias to an internal type.
type ParseResponse = shared.ParseResponse

// The response from the document processing service. Note that there can be two
// types of responses, Full Result and URL Result. This is due to limitations on
// the max return size on HTTPS. If the response is too large, it will be returned
// as a presigned URL in the URL response. You should handle this in your
// application.
//
// This is an alias to an internal type.
type ParseResponseResult = shared.ParseResponseResult

// This is an alias to an internal type.
type ParseResponseResultFullResult = shared.ParseResponseResultFullResult

// This is an alias to an internal type.
type ParseResponseResultFullResultChunk = shared.ParseResponseResultFullResultChunk

// This is an alias to an internal type.
type ParseResponseResultFullResultChunksBlock = shared.ParseResponseResultFullResultChunksBlock

// The type of block extracted from the document.
//
// This is an alias to an internal type.
type ParseResponseResultFullResultChunksBlocksType = shared.ParseResponseResultFullResultChunksBlocksType

// This is an alias to an internal value.
const ParseResponseResultFullResultChunksBlocksTypeHeader = shared.ParseResponseResultFullResultChunksBlocksTypeHeader

// This is an alias to an internal value.
const ParseResponseResultFullResultChunksBlocksTypeFooter = shared.ParseResponseResultFullResultChunksBlocksTypeFooter

// This is an alias to an internal value.
const ParseResponseResultFullResultChunksBlocksTypeTitle = shared.ParseResponseResultFullResultChunksBlocksTypeTitle

// This is an alias to an internal value.
const ParseResponseResultFullResultChunksBlocksTypeSectionHeader = shared.ParseResponseResultFullResultChunksBlocksTypeSectionHeader

// This is an alias to an internal value.
const ParseResponseResultFullResultChunksBlocksTypePageNumber = shared.ParseResponseResultFullResultChunksBlocksTypePageNumber

// This is an alias to an internal value.
const ParseResponseResultFullResultChunksBlocksTypeListItem = shared.ParseResponseResultFullResultChunksBlocksTypeListItem

// This is an alias to an internal value.
const ParseResponseResultFullResultChunksBlocksTypeFigure = shared.ParseResponseResultFullResultChunksBlocksTypeFigure

// This is an alias to an internal value.
const ParseResponseResultFullResultChunksBlocksTypeTable = shared.ParseResponseResultFullResultChunksBlocksTypeTable

// This is an alias to an internal value.
const ParseResponseResultFullResultChunksBlocksTypeKeyValue = shared.ParseResponseResultFullResultChunksBlocksTypeKeyValue

// This is an alias to an internal value.
const ParseResponseResultFullResultChunksBlocksTypeText = shared.ParseResponseResultFullResultChunksBlocksTypeText

// This is an alias to an internal value.
const ParseResponseResultFullResultChunksBlocksTypeComment = shared.ParseResponseResultFullResultChunksBlocksTypeComment

// This is an alias to an internal value.
const ParseResponseResultFullResultChunksBlocksTypeDiscard = shared.ParseResponseResultFullResultChunksBlocksTypeDiscard

// type = 'full'
//
// This is an alias to an internal type.
type ParseResponseResultFullResultType = shared.ParseResponseResultFullResultType

// This is an alias to an internal value.
const ParseResponseResultFullResultTypeFull = shared.ParseResponseResultFullResultTypeFull

// This is an alias to an internal type.
type ParseResponseResultFullResultOcr = shared.ParseResponseResultFullResultOcr

// This is an alias to an internal type.
type ParseResponseResultFullResultOcrLine = shared.ParseResponseResultFullResultOcrLine

// This is an alias to an internal type.
type ParseResponseResultFullResultOcrWord = shared.ParseResponseResultFullResultOcrWord

// This is an alias to an internal type.
type ParseResponseResultURLResult = shared.ParseResponseResultURLResult

// type = 'url'
//
// This is an alias to an internal type.
type ParseResponseResultURLResultType = shared.ParseResponseResultURLResultType

// This is an alias to an internal value.
const ParseResponseResultURLResultTypeURL = shared.ParseResponseResultURLResultTypeURL

// type = 'full'
//
// This is an alias to an internal type.
type ParseResponseResultType = shared.ParseResponseResultType

// This is an alias to an internal value.
const ParseResponseResultTypeFull = shared.ParseResponseResultTypeFull

// This is an alias to an internal value.
const ParseResponseResultTypeURL = shared.ParseResponseResultTypeURL

// This is an alias to an internal type.
type ParseUsage = shared.ParseUsage

// This is an alias to an internal type.
type SplitCategoryParam = shared.SplitCategoryParam

// This is an alias to an internal type.
type SplitResponse = shared.SplitResponse

// The extracted response in your provided schema. This is a list of dictionaries.
// If disbale_chunking is True (default), then it will be a list of length one.
//
// This is an alias to an internal type.
type SplitResponseResult = shared.SplitResponseResult

// This is an alias to an internal type.
type Upload = shared.Upload

// This is an alias to an internal type.
type WebhookConfigNewParam = shared.WebhookConfigNewParam

// The mode to use for webhook delivery. Defaults to 'disabled'. We recommend using
// 'svix' for production environments.
//
// This is an alias to an internal type.
type WebhookConfigNewMode = shared.WebhookConfigNewMode

// This is an alias to an internal value.
const WebhookConfigNewModeDisabled = shared.WebhookConfigNewModeDisabled

// This is an alias to an internal value.
const WebhookConfigNewModeSvix = shared.WebhookConfigNewModeSvix

// This is an alias to an internal value.
const WebhookConfigNewModeDirect = shared.WebhookConfigNewModeDirect
