// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reducto

import (
	"github.com/reductoai/reducto-go-sdk/internal/apierror"
	"github.com/reductoai/reducto-go-sdk/shared"
)

type Error = apierror.Error

// This is an alias to an internal type.
type AsyncEditResponse = shared.AsyncEditResponse

// This is an alias to an internal type.
type AsyncExtractResponse = shared.AsyncExtractResponse

// This is an alias to an internal type.
type AsyncParseResponse = shared.AsyncParseResponse

// This is an alias to an internal type.
type AsyncPipelineResponse = shared.AsyncPipelineResponse

// This is an alias to an internal type.
type AsyncSplitResponse = shared.AsyncSplitResponse

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

// Response from classify job - returned when polling /job/{job_id}
//
// This is an alias to an internal type.
type ClassifyResponse = shared.ClassifyResponse

// This is an alias to an internal type.
type ClassifyResponseResult = shared.ClassifyResponseResult

// Overall confidence breakdown for classification response.
//
// This is an alias to an internal type.
type ClassifyResponseResponseConfidence = shared.ClassifyResponseResponseConfidence

// Confidence result for a category.
//
// This is an alias to an internal type.
type ClassifyResponseResponseConfidenceCategory = shared.ClassifyResponseResponseConfidenceCategory

// Confidence result for a single criterion.
//
// This is an alias to an internal type.
type ClassifyResponseResponseConfidenceCategoriesCriteriaConfidence = shared.ClassifyResponseResponseConfidenceCategoriesCriteriaConfidence

// This is an alias to an internal type.
type ClassifyResponseResponseConfidenceCategoriesCriteriaConfidenceConfidence = shared.ClassifyResponseResponseConfidenceCategoriesCriteriaConfidenceConfidence

// This is an alias to an internal value.
const ClassifyResponseResponseConfidenceCategoriesCriteriaConfidenceConfidenceHigh = shared.ClassifyResponseResponseConfidenceCategoriesCriteriaConfidenceConfidenceHigh

// This is an alias to an internal value.
const ClassifyResponseResponseConfidenceCategoriesCriteriaConfidenceConfidenceLow = shared.ClassifyResponseResponseConfidenceCategoriesCriteriaConfidenceConfidenceLow

// This is an alias to an internal type.
type DirectWebhookConfigParam = shared.DirectWebhookConfigParam

// This is an alias to an internal type.
type DirectWebhookConfigMode = shared.DirectWebhookConfigMode

// This is an alias to an internal value.
const DirectWebhookConfigModeDirect = shared.DirectWebhookConfigModeDirect

// This is an alias to an internal type.
type EditResponse = shared.EditResponse

// This is an alias to an internal type.
type ExtractResponse = shared.ExtractResponse

// This is an alias to an internal type.
type FigureAgenticParam = shared.FigureAgenticParam

// This is an alias to an internal type.
type FigureAgenticScope = shared.FigureAgenticScope

// This is an alias to an internal value.
const FigureAgenticScopeFigure = shared.FigureAgenticScopeFigure

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
type SplitResponseResultSplitResult = shared.SplitResponseResultSplitResult

// This is an alias to an internal type.
type SplitResponseResultSplitResultSplit = shared.SplitResponseResultSplitResultSplit

// This is an alias to an internal type.
type SplitResponseResultSplitResultSplitsConf = shared.SplitResponseResultSplitResultSplitsConf

// This is an alias to an internal value.
const SplitResponseResultSplitResultSplitsConfHigh = shared.SplitResponseResultSplitResultSplitsConfHigh

// This is an alias to an internal value.
const SplitResponseResultSplitResultSplitsConfLow = shared.SplitResponseResultSplitResultSplitsConfLow

// This is an alias to an internal type.
type SplitResponseResultSplitResultSplitsPartition = shared.SplitResponseResultSplitResultSplitsPartition

// This is an alias to an internal type.
type SplitResponseResultSplitResultSplitsPartitionsConf = shared.SplitResponseResultSplitResultSplitsPartitionsConf

// This is an alias to an internal value.
const SplitResponseResultSplitResultSplitsPartitionsConfHigh = shared.SplitResponseResultSplitResultSplitsPartitionsConfHigh

// This is an alias to an internal value.
const SplitResponseResultSplitResultSplitsPartitionsConfLow = shared.SplitResponseResultSplitResultSplitsPartitionsConfLow

// This is an alias to an internal type.
type SplitResponseResultDeepSplitResult = shared.SplitResponseResultDeepSplitResult

// This is an alias to an internal type.
type SplitResponseResultDeepSplitResultSplit = shared.SplitResponseResultDeepSplitResultSplit

// This is an alias to an internal type.
type SplitResponseResultDeepSplitResultSplitsPartition = shared.SplitResponseResultDeepSplitResultSplitsPartition

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
