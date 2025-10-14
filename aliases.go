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
const EditResponseFormSchemaTypeDropdown = shared.EditResponseFormSchemaTypeDropdown

// This is an alias to an internal value.
const EditResponseFormSchemaTypeBarcode = shared.EditResponseFormSchemaTypeBarcode

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
type PipelineResponseResultExtractArrayResultV3ExtractResponse = shared.PipelineResponseResultExtractArrayResultV3ExtractResponse

// This is an alias to an internal type.
type PipelineResponseResultExtractArrayResultV3ExtractResponseUsage = shared.PipelineResponseResultExtractArrayResultV3ExtractResponseUsage

// This is an alias to an internal type.
type PipelineResponseResultExtractV3ExtractResponse = shared.PipelineResponseResultExtractV3ExtractResponse

// This is an alias to an internal type.
type PipelineResponseResultExtractV3ExtractResponseUsage = shared.PipelineResponseResultExtractV3ExtractResponseUsage

// This is an alias to an internal type.
type SplitCategoryParam = shared.SplitCategoryParam

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
