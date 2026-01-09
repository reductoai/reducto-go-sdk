// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reducto

import (
	"github.com/reductoai/reducto-go-sdk/internal/apierror"
	"github.com/reductoai/reducto-go-sdk/shared"
)

type Error = apierror.Error

// This is an alias to an internal type.
type BoundingBox = shared.BoundingBox

// This is an alias to an internal type.
type BoundingBoxParam = shared.BoundingBoxParam

// This is an alias to an internal type.
type ChunkingParam = shared.ChunkingParam

// Choose how to partition chunks. Variable mode chunks by character length and
// visual context. Section mode chunks by section headers. Page mode chunks
// according to pages. Page sections mode chunks first by page, then by sections
// within each page. Disabled returns one single chunk.
//
// This is an alias to an internal type.
type ChunkingChunkMode = shared.ChunkingChunkMode

// This is an alias to an internal value.
const ChunkingChunkModeVariable = shared.ChunkingChunkModeVariable

// This is an alias to an internal value.
const ChunkingChunkModeSection = shared.ChunkingChunkModeSection

// This is an alias to an internal value.
const ChunkingChunkModePage = shared.ChunkingChunkModePage

// This is an alias to an internal value.
const ChunkingChunkModeDisabled = shared.ChunkingChunkModeDisabled

// This is an alias to an internal value.
const ChunkingChunkModeBlock = shared.ChunkingChunkModeBlock

// This is an alias to an internal value.
const ChunkingChunkModePageSections = shared.ChunkingChunkModePageSections

// This is an alias to an internal type.
type ConfigV3AsyncConfigParam = shared.ConfigV3AsyncConfigParam

// The webhook configuration for the asynchronous processing.
//
// This is an alias to an internal type.
type ConfigV3AsyncConfigWebhookUnionParam = shared.ConfigV3AsyncConfigWebhookUnionParam

// This is an alias to an internal type.
type ConfigV3AsyncConfigWebhookMode = shared.ConfigV3AsyncConfigWebhookMode

// This is an alias to an internal value.
const ConfigV3AsyncConfigWebhookModeSvix = shared.ConfigV3AsyncConfigWebhookModeSvix

// This is an alias to an internal value.
const ConfigV3AsyncConfigWebhookModeDirect = shared.ConfigV3AsyncConfigWebhookModeDirect

// This is an alias to an internal type.
type DirectWebhookConfigParam = shared.DirectWebhookConfigParam

// This is an alias to an internal type.
type DirectWebhookConfigMode = shared.DirectWebhookConfigMode

// This is an alias to an internal value.
const DirectWebhookConfigModeDirect = shared.DirectWebhookConfigModeDirect

// This is an alias to an internal type.
type EditResponse = shared.EditResponse

// This is an alias to an internal type.
type EditResponseFormSchema = shared.EditResponseFormSchema

// Type of the form widget
//
// This is an alias to an internal type.
type EditResponseFormSchemaType = shared.EditResponseFormSchemaType

// This is an alias to an internal value.
const EditResponseFormSchemaTypeText = shared.EditResponseFormSchemaTypeText

// This is an alias to an internal value.
const EditResponseFormSchemaTypeCheckbox = shared.EditResponseFormSchemaTypeCheckbox

// This is an alias to an internal value.
const EditResponseFormSchemaTypeRadio = shared.EditResponseFormSchemaTypeRadio

// This is an alias to an internal value.
const EditResponseFormSchemaTypeDropdown = shared.EditResponseFormSchemaTypeDropdown

// This is an alias to an internal value.
const EditResponseFormSchemaTypeBarcode = shared.EditResponseFormSchemaTypeBarcode

// This is an alias to an internal type.
type EnhanceParam = shared.EnhanceParam

// This is an alias to an internal type.
type EnhanceAgenticUnionParam = shared.EnhanceAgenticUnionParam

// This is an alias to an internal type.
type EnhanceAgenticScope = shared.EnhanceAgenticScope

// This is an alias to an internal value.
const EnhanceAgenticScopeTable = shared.EnhanceAgenticScopeTable

// This is an alias to an internal value.
const EnhanceAgenticScopeFigure = shared.EnhanceAgenticScopeFigure

// This is an alias to an internal value.
const EnhanceAgenticScopeText = shared.EnhanceAgenticScopeText

// This is an alias to an internal type.
type ExtractResponse = shared.ExtractResponse

// This is an alias to an internal type.
type ExtractUsage = shared.ExtractUsage

// This is an alias to an internal type.
type FigureAgenticParam = shared.FigureAgenticParam

// This is an alias to an internal type.
type FigureAgenticScope = shared.FigureAgenticScope

// This is an alias to an internal value.
const FigureAgenticScopeFigure = shared.FigureAgenticScopeFigure

// This is an alias to an internal type.
type FormattingParam = shared.FormattingParam

// This is an alias to an internal type.
type FormattingInclude = shared.FormattingInclude

// This is an alias to an internal value.
const FormattingIncludeChangeTracking = shared.FormattingIncludeChangeTracking

// This is an alias to an internal value.
const FormattingIncludeHighlight = shared.FormattingIncludeHighlight

// This is an alias to an internal value.
const FormattingIncludeComments = shared.FormattingIncludeComments

// This is an alias to an internal value.
const FormattingIncludeHyperlinks = shared.FormattingIncludeHyperlinks

// This is an alias to an internal value.
const FormattingIncludeSignatures = shared.FormattingIncludeSignatures

// The mode to use for table output. Defaults to dynamic, which returns md for
// simpler tables and html for more complex tables.
//
// This is an alias to an internal type.
type FormattingTableOutputFormat = shared.FormattingTableOutputFormat

// This is an alias to an internal value.
const FormattingTableOutputFormatHTML = shared.FormattingTableOutputFormatHTML

// This is an alias to an internal value.
const FormattingTableOutputFormatJson = shared.FormattingTableOutputFormatJson

// This is an alias to an internal value.
const FormattingTableOutputFormatMd = shared.FormattingTableOutputFormatMd

// This is an alias to an internal value.
const FormattingTableOutputFormatJsonbbox = shared.FormattingTableOutputFormatJsonbbox

// This is an alias to an internal value.
const FormattingTableOutputFormatDynamic = shared.FormattingTableOutputFormatDynamic

// This is an alias to an internal value.
const FormattingTableOutputFormatCsv = shared.FormattingTableOutputFormatCsv

// This is an alias to an internal type.
type PageRangeParam = shared.PageRangeParam

// This is an alias to an internal type.
type ParseOptionsParam = shared.ParseOptionsParam

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
const ParseResponseResultFullResultChunksBlocksTypeSignature = shared.ParseResponseResultFullResultChunksBlocksTypeSignature

// Granular confidence scores for the block. It is a dictionary of confidence
// scores for the block. The confidence scores will not be None if the user has
// enabled numeric confidence scores.
//
// This is an alias to an internal type.
type ParseResponseResultFullResultChunksBlocksGranularConfidence = shared.ParseResponseResultFullResultChunksBlocksGranularConfidence

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
type PipelineResponse = shared.PipelineResponse

// This is an alias to an internal type.
type PipelineResponseResult = shared.PipelineResponseResult

// This is an alias to an internal type.
type PipelineResponseResultExtractUnion = shared.PipelineResponseResultExtractUnion

// This is an alias to an internal type.
type PipelineResponseResultExtractArray = shared.PipelineResponseResultExtractArray

// This is the response format for Extract -> Split Pipelines
//
// This is an alias to an internal type.
type PipelineResponseResultExtractArrayItem = shared.PipelineResponseResultExtractArrayItem

// This is an alias to an internal type.
type PipelineResponseResultExtractArrayResult = shared.PipelineResponseResultExtractArrayResult

// This is an alias to an internal type.
type PipelineResponseResultParseUnion = shared.PipelineResponseResultParseUnion

// This is an alias to an internal type.
type PipelineResponseResultParseArray = shared.PipelineResponseResultParseArray

// This is an alias to an internal type.
type RetrievalParam = shared.RetrievalParam

// This is an alias to an internal type.
type RetrievalFilterBlock = shared.RetrievalFilterBlock

// This is an alias to an internal value.
const RetrievalFilterBlockHeader = shared.RetrievalFilterBlockHeader

// This is an alias to an internal value.
const RetrievalFilterBlockFooter = shared.RetrievalFilterBlockFooter

// This is an alias to an internal value.
const RetrievalFilterBlockTitle = shared.RetrievalFilterBlockTitle

// This is an alias to an internal value.
const RetrievalFilterBlockSectionHeader = shared.RetrievalFilterBlockSectionHeader

// This is an alias to an internal value.
const RetrievalFilterBlockPageNumber = shared.RetrievalFilterBlockPageNumber

// This is an alias to an internal value.
const RetrievalFilterBlockListItem = shared.RetrievalFilterBlockListItem

// This is an alias to an internal value.
const RetrievalFilterBlockFigure = shared.RetrievalFilterBlockFigure

// This is an alias to an internal value.
const RetrievalFilterBlockTable = shared.RetrievalFilterBlockTable

// This is an alias to an internal value.
const RetrievalFilterBlockKeyValue = shared.RetrievalFilterBlockKeyValue

// This is an alias to an internal value.
const RetrievalFilterBlockText = shared.RetrievalFilterBlockText

// This is an alias to an internal value.
const RetrievalFilterBlockComment = shared.RetrievalFilterBlockComment

// This is an alias to an internal value.
const RetrievalFilterBlockSignature = shared.RetrievalFilterBlockSignature

// This is an alias to an internal type.
type SettingsParam = shared.SettingsParam

// The mode to use for text extraction from PDFs. OCR mode uses optical character
// recognition only. Hybrid mode combines OCR with embedded PDF text for best
// accuracy (default).
//
// This is an alias to an internal type.
type SettingsExtractionMode = shared.SettingsExtractionMode

// This is an alias to an internal value.
const SettingsExtractionModeOcr = shared.SettingsExtractionModeOcr

// This is an alias to an internal value.
const SettingsExtractionModeHybrid = shared.SettingsExtractionModeHybrid

// Standard is our best multilingual OCR system. Legacy only supports germanic
// languages and is available for backwards compatibility.
//
// This is an alias to an internal type.
type SettingsOcrSystem = shared.SettingsOcrSystem

// This is an alias to an internal value.
const SettingsOcrSystemStandard = shared.SettingsOcrSystemStandard

// This is an alias to an internal value.
const SettingsOcrSystemLegacy = shared.SettingsOcrSystemLegacy

// The page range to process (1-indexed). By default, the entire document is
// processed.
//
// This is an alias to an internal type.
type SettingsPageRangeUnionParam = shared.SettingsPageRangeUnionParam

// This is an alias to an internal type.
type SettingsPageRangeArrayParam = shared.SettingsPageRangeArrayParam

// This is an alias to an internal type.
type SettingsReturnImage = shared.SettingsReturnImage

// This is an alias to an internal value.
const SettingsReturnImageFigure = shared.SettingsReturnImageFigure

// This is an alias to an internal value.
const SettingsReturnImageTable = shared.SettingsReturnImageTable

// This is an alias to an internal type.
type SplitCategoryParam = shared.SplitCategoryParam

// This is an alias to an internal type.
type SplitLargeTablesParam = shared.SplitLargeTablesParam

// The size of the tables to split into. Defaults to 50. Use 'row' and 'column' to
// independently specify the number of rows and columns to include when splitting.
// If you only want to split by rows or columns, set the other value to None.
//
// This is an alias to an internal type.
type SplitLargeTablesSizeUnionParam = shared.SplitLargeTablesSizeUnionParam

// This is an alias to an internal type.
type SplitLargeTablesSizeSplitLargeTableSizesParam = shared.SplitLargeTablesSizeSplitLargeTableSizesParam

// This is an alias to an internal type.
type SplitResponse = shared.SplitResponse

// The split result.
//
// This is an alias to an internal type.
type SplitResponseResult = shared.SplitResponseResult

// This is an alias to an internal type.
type SplitResponseResultSplit = shared.SplitResponseResultSplit

// This is an alias to an internal type.
type SplitResponseResultSplitsConf = shared.SplitResponseResultSplitsConf

// This is an alias to an internal value.
const SplitResponseResultSplitsConfHigh = shared.SplitResponseResultSplitsConfHigh

// This is an alias to an internal value.
const SplitResponseResultSplitsConfLow = shared.SplitResponseResultSplitsConfLow

// This is an alias to an internal type.
type SplitResponseResultSplitsPartition = shared.SplitResponseResultSplitsPartition

// This is an alias to an internal type.
type SplitResponseResultSplitsPartitionsConf = shared.SplitResponseResultSplitsPartitionsConf

// This is an alias to an internal value.
const SplitResponseResultSplitsPartitionsConfHigh = shared.SplitResponseResultSplitsPartitionsConfHigh

// This is an alias to an internal value.
const SplitResponseResultSplitsPartitionsConfLow = shared.SplitResponseResultSplitsPartitionsConfLow

// This is an alias to an internal type.
type SpreadsheetParam = shared.SpreadsheetParam

// In a spreadsheet with different tables inside, we enable splitting up the tables
// by default. Accurate mode applies more powerful models for superior accuracy, at
// 5× the default per-cell rate. Disabling will register as one large table.
//
// This is an alias to an internal type.
type SpreadsheetClustering = shared.SpreadsheetClustering

// This is an alias to an internal value.
const SpreadsheetClusteringAccurate = shared.SpreadsheetClusteringAccurate

// This is an alias to an internal value.
const SpreadsheetClusteringFast = shared.SpreadsheetClusteringFast

// This is an alias to an internal value.
const SpreadsheetClusteringDisabled = shared.SpreadsheetClusteringDisabled

// This is an alias to an internal type.
type SpreadsheetExclude = shared.SpreadsheetExclude

// This is an alias to an internal value.
const SpreadsheetExcludeHiddenSheets = shared.SpreadsheetExcludeHiddenSheets

// This is an alias to an internal value.
const SpreadsheetExcludeHiddenRows = shared.SpreadsheetExcludeHiddenRows

// This is an alias to an internal value.
const SpreadsheetExcludeHiddenCols = shared.SpreadsheetExcludeHiddenCols

// This is an alias to an internal value.
const SpreadsheetExcludeStyling = shared.SpreadsheetExcludeStyling

// This is an alias to an internal value.
const SpreadsheetExcludeSpreadsheetImages = shared.SpreadsheetExcludeSpreadsheetImages

// This is an alias to an internal type.
type SpreadsheetInclude = shared.SpreadsheetInclude

// This is an alias to an internal value.
const SpreadsheetIncludeCellColors = shared.SpreadsheetIncludeCellColors

// This is an alias to an internal value.
const SpreadsheetIncludeFormula = shared.SpreadsheetIncludeFormula

// This is an alias to an internal value.
const SpreadsheetIncludeDropdowns = shared.SpreadsheetIncludeDropdowns

// This is an alias to an internal type.
type SvixWebhookConfigParam = shared.SvixWebhookConfigParam

// This is an alias to an internal type.
type SvixWebhookConfigMode = shared.SvixWebhookConfigMode

// This is an alias to an internal value.
const SvixWebhookConfigModeSvix = shared.SvixWebhookConfigModeSvix

// This is an alias to an internal type.
type TableAgenticParam = shared.TableAgenticParam

// This is an alias to an internal type.
type TableAgenticScope = shared.TableAgenticScope

// This is an alias to an internal value.
const TableAgenticScopeTable = shared.TableAgenticScopeTable

// This is an alias to an internal type.
type TextAgenticParam = shared.TextAgenticParam

// This is an alias to an internal type.
type TextAgenticScope = shared.TextAgenticScope

// This is an alias to an internal value.
const TextAgenticScopeText = shared.TextAgenticScopeText

// This is an alias to an internal type.
type Upload = shared.Upload

// This is an alias to an internal type.
type UploadParam = shared.UploadParam

// This is an alias to an internal type.
type V3ExtractResponse = shared.V3ExtractResponse

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
