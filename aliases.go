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

// This is an alias to an internal type.
type ClassifyResponseResultClassifyResponseCategory = shared.ClassifyResponseResultClassifyResponseCategory

// This is an alias to an internal type.
type ClassifyResponseResultURLResult = shared.ClassifyResponseResultURLResult

// type = 'url'
//
// This is an alias to an internal type.
type ClassifyResponseResultURLResultType = shared.ClassifyResponseResultURLResultType

// This is an alias to an internal value.
const ClassifyResponseResultURLResultTypeURL = shared.ClassifyResponseResultURLResultTypeURL

// type = 'url'
//
// This is an alias to an internal type.
type ClassifyResponseResultType = shared.ClassifyResponseResultType

// This is an alias to an internal value.
const ClassifyResponseResultTypeURL = shared.ClassifyResponseResultTypeURL

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
type ClassifyResponseResponseType = shared.ClassifyResponseResponseType

// This is an alias to an internal value.
const ClassifyResponseResponseTypeClassify = shared.ClassifyResponseResponseTypeClassify

// This is an alias to an internal type.
type ClassifyResponseUsage = shared.ClassifyResponseUsage

// This is an alias to an internal type.
type DirectWebhookConfigParam = shared.DirectWebhookConfigParam

// This is an alias to an internal type.
type DirectWebhookConfigMode = shared.DirectWebhookConfigMode

// This is an alias to an internal value.
const DirectWebhookConfigModeDirect = shared.DirectWebhookConfigModeDirect

// This is an alias to an internal type.
type EditResponse = shared.EditResponse

// This is an alias to an internal type.
type EditResponseResponseType = shared.EditResponseResponseType

// This is an alias to an internal value.
const EditResponseResponseTypeEdit = shared.EditResponseResponseTypeEdit

// This is an alias to an internal type.
type ExtractResponse = shared.ExtractResponse

// This is an alias to an internal type.
type FigureAgenticParam = shared.FigureAgenticParam

// This is an alias to an internal type.
type FigureAgenticScope = shared.FigureAgenticScope

// This is an alias to an internal value.
const FigureAgenticScopeFigure = shared.FigureAgenticScopeFigure

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

// Which pipeline produced this response. `lite` means Reducto Flash Lite served
// the request; `base` is the standard pipeline. Optional / nullable for forward
// compatibility — older API instances or persisted responses written before this
// field existed will leave it `None`; treat `None` as `base`.
//
// This is an alias to an internal type.
type ParseResponseParseMode = shared.ParseResponseParseMode

// This is an alias to an internal value.
const ParseResponseParseModeBase = shared.ParseResponseParseModeBase

// This is an alias to an internal value.
const ParseResponseParseModeLite = shared.ParseResponseParseModeLite

// This is an alias to an internal type.
type ParseResponseResponseType = shared.ParseResponseResponseType

// This is an alias to an internal value.
const ParseResponseResponseTypeParse = shared.ParseResponseResponseTypeParse

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
type PipelineResponseResultExtractArrayResultUnion = shared.PipelineResponseResultExtractArrayResultUnion

// This is an alias to an internal type.
type PipelineResponseResultParseUnion = shared.PipelineResponseResultParseUnion

// This is an alias to an internal type.
type PipelineResponseResultParseArray = shared.PipelineResponseResultParseArray

// This is an alias to an internal type.
type PipelineResponseResponseType = shared.PipelineResponseResponseType

// This is an alias to an internal value.
const PipelineResponseResponseTypePipeline = shared.PipelineResponseResponseTypePipeline

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

// The split result. If force_url_result is True, this is returned as a URL result.
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
type SplitResponseResultURLResult = shared.SplitResponseResultURLResult

// type = 'url'
//
// This is an alias to an internal type.
type SplitResponseResultURLResultType = shared.SplitResponseResultURLResultType

// This is an alias to an internal value.
const SplitResponseResultURLResultTypeURL = shared.SplitResponseResultURLResultTypeURL

// type = 'url'
//
// This is an alias to an internal type.
type SplitResponseResultType = shared.SplitResponseResultType

// This is an alias to an internal value.
const SplitResponseResultTypeURL = shared.SplitResponseResultTypeURL

// This is an alias to an internal type.
type SplitResponseResponseType = shared.SplitResponseResponseType

// This is an alias to an internal value.
const SplitResponseResponseTypeSplit = shared.SplitResponseResponseTypeSplit

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

// Mode for table agentic: 'default' selectively applies enrichment only to tables
// likely to benefit, and 'max' runs enrichment on all tables.
//
// This is an alias to an internal type.
type TableAgenticMode = shared.TableAgenticMode

// This is an alias to an internal value.
const TableAgenticModeDefault = shared.TableAgenticModeDefault

// This is an alias to an internal value.
const TableAgenticModeAuto = shared.TableAgenticModeAuto

// This is an alias to an internal value.
const TableAgenticModeMax = shared.TableAgenticModeMax

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
