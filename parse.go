// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package reducto

import (
	"context"
	"net/http"
	"reflect"
	"slices"

	"github.com/reductoai/reducto-go-sdk/internal/apijson"
	"github.com/reductoai/reducto-go-sdk/internal/param"
	"github.com/reductoai/reducto-go-sdk/internal/requestconfig"
	"github.com/reductoai/reducto-go-sdk/option"
	"github.com/tidwall/gjson"
)

// ParseService contains methods and other services that help with interacting with
// the reducto API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewParseService] method instead.
type ParseService struct {
	Options []option.RequestOption
}

// NewParseService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewParseService(opts ...option.RequestOption) (r *ParseService) {
	r = &ParseService{}
	r.Options = opts
	return
}

// Parse
func (r *ParseService) New(ctx context.Context, body ParseNewParams, opts ...option.RequestOption) (res *ParseNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "parse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ParseResponse struct {
	// The duration of the parse request in seconds.
	Duration float64 `json:"duration" api:"required"`
	JobID    string  `json:"job_id" api:"required"`
	// The response from the document processing service. Note that there can be two
	// types of responses, Full Result and URL Result. This is due to limitations on
	// the max return size on HTTPS. If the response is too large, it will be returned
	// as a presigned URL in the URL response. You should handle this in your
	// application.
	Result ParseResponseResult `json:"result" api:"required"`
	Usage  ParseUsage          `json:"usage" api:"required"`
	// The storage URL of the converted PDF file.
	PdfURL string `json:"pdf_url" api:"nullable"`
	// The link to the studio pipeline for the document.
	StudioLink string            `json:"studio_link" api:"nullable"`
	JSON       parseResponseJSON `json:"-"`
}

// parseResponseJSON contains the JSON metadata for the struct [ParseResponse]
type parseResponseJSON struct {
	Duration    apijson.Field
	JobID       apijson.Field
	Result      apijson.Field
	Usage       apijson.Field
	PdfURL      apijson.Field
	StudioLink  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseJSON) RawJSON() string {
	return r.raw
}

func (r ParseResponse) implementsParseNewResponse() {}

func (r ParseResponse) implementsPipelineResponseResultParseUnion() {}

func (r ParseResponse) implementsJobGetResponseAsyncJobResponseResult() {}

func (r ParseResponse) implementsJobGetResponseEnhancedAsyncJobResponseResult() {}

// The response from the document processing service. Note that there can be two
// types of responses, Full Result and URL Result. This is due to limitations on
// the max return size on HTTPS. If the response is too large, it will be returned
// as a presigned URL in the URL response. You should handle this in your
// application.
type ParseResponseResult struct {
	// type = 'full'
	Type ParseResponseResultType `json:"type" api:"required"`
	// This field can have the runtime type of [[]ParseResponseResultFullResultChunk].
	Chunks interface{} `json:"chunks"`
	// This field can have the runtime type of [interface{}].
	Custom interface{} `json:"custom"`
	// This field can have the runtime type of [ParseResponseResultFullResultOcr].
	Ocr      interface{}             `json:"ocr"`
	ResultID string                  `json:"result_id"`
	URL      string                  `json:"url"`
	JSON     parseResponseResultJSON `json:"-"`
	union    ParseResponseResultUnion
}

// parseResponseResultJSON contains the JSON metadata for the struct
// [ParseResponseResult]
type parseResponseResultJSON struct {
	Type        apijson.Field
	Chunks      apijson.Field
	Custom      apijson.Field
	Ocr         apijson.Field
	ResultID    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r parseResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *ParseResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = ParseResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ParseResponseResultUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ParseResponseResultFullResult],
// [ParseResponseResultURLResult].
func (r ParseResponseResult) AsUnion() ParseResponseResultUnion {
	return r.union
}

// The response from the document processing service. Note that there can be two
// types of responses, Full Result and URL Result. This is due to limitations on
// the max return size on HTTPS. If the response is too large, it will be returned
// as a presigned URL in the URL response. You should handle this in your
// application.
//
// Union satisfied by [ParseResponseResultFullResult] or
// [ParseResponseResultURLResult].
type ParseResponseResultUnion interface {
	implementsParseResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ParseResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParseResponseResultFullResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParseResponseResultURLResult{}),
		},
	)
}

type ParseResponseResultFullResult struct {
	Chunks []ParseResponseResultFullResultChunk `json:"chunks" api:"required"`
	// type = 'full'
	Type   ParseResponseResultFullResultType `json:"type" api:"required"`
	Custom interface{}                       `json:"custom"`
	Ocr    ParseResponseResultFullResultOcr  `json:"ocr" api:"nullable"`
	JSON   parseResponseResultFullResultJSON `json:"-"`
}

// parseResponseResultFullResultJSON contains the JSON metadata for the struct
// [ParseResponseResultFullResult]
type parseResponseResultFullResultJSON struct {
	Chunks      apijson.Field
	Type        apijson.Field
	Custom      apijson.Field
	Ocr         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParseResponseResultFullResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseResultFullResultJSON) RawJSON() string {
	return r.raw
}

func (r ParseResponseResultFullResult) implementsParseResponseResult() {}

type ParseResponseResultFullResultChunk struct {
	Blocks []ParseResponseResultFullResultChunksBlock `json:"blocks" api:"required"`
	// The content of the chunk extracted from the document.
	Content string `json:"content" api:"required"`
	// Chunk content optimized for embedding and retrieval.
	Embed string `json:"embed" api:"required"`
	// The enriched content of the chunk extracted from the document.
	Enriched string `json:"enriched" api:"required,nullable"`
	// Whether the enrichment was successful.
	EnrichmentSuccess bool                                   `json:"enrichment_success"`
	JSON              parseResponseResultFullResultChunkJSON `json:"-"`
}

// parseResponseResultFullResultChunkJSON contains the JSON metadata for the struct
// [ParseResponseResultFullResultChunk]
type parseResponseResultFullResultChunkJSON struct {
	Blocks            apijson.Field
	Content           apijson.Field
	Embed             apijson.Field
	Enriched          apijson.Field
	EnrichmentSuccess apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ParseResponseResultFullResultChunk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseResultFullResultChunkJSON) RawJSON() string {
	return r.raw
}

type ParseResponseResultFullResultChunksBlock struct {
	// The bounding box of the block extracted from the document.
	Bbox BoundingBox `json:"bbox" api:"required"`
	// The content of the block extracted from the document.
	Content string `json:"content" api:"required"`
	// The type of block extracted from the document.
	Type ParseResponseResultFullResultChunksBlocksType `json:"type" api:"required"`
	// (Experimental) The URL/link to chart data JSON for figure blocks processed by
	// chart agent.
	ChartData []string `json:"chart_data" api:"nullable"`
	// The confidence for the block. It is either low or high and takes into account
	// factors like OCR and table structure
	Confidence string `json:"confidence" api:"nullable"`
	// Extra metadata fields for the block. Fields like 'is_chart' will only appear
	// when set to True.
	Extra map[string]interface{} `json:"extra" api:"nullable"`
	// Granular confidence scores for the block. It is a dictionary of confidence
	// scores for the block. The confidence scores will not be None if the user has
	// enabled numeric confidence scores.
	GranularConfidence ParseResponseResultFullResultChunksBlocksGranularConfidence `json:"granular_confidence" api:"nullable"`
	// (Experimental) The URL of the image associated with the block.
	ImageURL string                                       `json:"image_url" api:"nullable"`
	JSON     parseResponseResultFullResultChunksBlockJSON `json:"-"`
}

// parseResponseResultFullResultChunksBlockJSON contains the JSON metadata for the
// struct [ParseResponseResultFullResultChunksBlock]
type parseResponseResultFullResultChunksBlockJSON struct {
	Bbox               apijson.Field
	Content            apijson.Field
	Type               apijson.Field
	ChartData          apijson.Field
	Confidence         apijson.Field
	Extra              apijson.Field
	GranularConfidence apijson.Field
	ImageURL           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ParseResponseResultFullResultChunksBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseResultFullResultChunksBlockJSON) RawJSON() string {
	return r.raw
}

// The type of block extracted from the document.
type ParseResponseResultFullResultChunksBlocksType string

const (
	ParseResponseResultFullResultChunksBlocksTypeHeader        ParseResponseResultFullResultChunksBlocksType = "Header"
	ParseResponseResultFullResultChunksBlocksTypeFooter        ParseResponseResultFullResultChunksBlocksType = "Footer"
	ParseResponseResultFullResultChunksBlocksTypeTitle         ParseResponseResultFullResultChunksBlocksType = "Title"
	ParseResponseResultFullResultChunksBlocksTypeSectionHeader ParseResponseResultFullResultChunksBlocksType = "Section Header"
	ParseResponseResultFullResultChunksBlocksTypePageNumber    ParseResponseResultFullResultChunksBlocksType = "Page Number"
	ParseResponseResultFullResultChunksBlocksTypeListItem      ParseResponseResultFullResultChunksBlocksType = "List Item"
	ParseResponseResultFullResultChunksBlocksTypeFigure        ParseResponseResultFullResultChunksBlocksType = "Figure"
	ParseResponseResultFullResultChunksBlocksTypeTable         ParseResponseResultFullResultChunksBlocksType = "Table"
	ParseResponseResultFullResultChunksBlocksTypeKeyValue      ParseResponseResultFullResultChunksBlocksType = "Key Value"
	ParseResponseResultFullResultChunksBlocksTypeText          ParseResponseResultFullResultChunksBlocksType = "Text"
	ParseResponseResultFullResultChunksBlocksTypeComment       ParseResponseResultFullResultChunksBlocksType = "Comment"
	ParseResponseResultFullResultChunksBlocksTypeSignature     ParseResponseResultFullResultChunksBlocksType = "Signature"
)

func (r ParseResponseResultFullResultChunksBlocksType) IsKnown() bool {
	switch r {
	case ParseResponseResultFullResultChunksBlocksTypeHeader, ParseResponseResultFullResultChunksBlocksTypeFooter, ParseResponseResultFullResultChunksBlocksTypeTitle, ParseResponseResultFullResultChunksBlocksTypeSectionHeader, ParseResponseResultFullResultChunksBlocksTypePageNumber, ParseResponseResultFullResultChunksBlocksTypeListItem, ParseResponseResultFullResultChunksBlocksTypeFigure, ParseResponseResultFullResultChunksBlocksTypeTable, ParseResponseResultFullResultChunksBlocksTypeKeyValue, ParseResponseResultFullResultChunksBlocksTypeText, ParseResponseResultFullResultChunksBlocksTypeComment, ParseResponseResultFullResultChunksBlocksTypeSignature:
		return true
	}
	return false
}

// Granular confidence scores for the block. It is a dictionary of confidence
// scores for the block. The confidence scores will not be None if the user has
// enabled numeric confidence scores.
type ParseResponseResultFullResultChunksBlocksGranularConfidence struct {
	ExtractConfidence float64                                                         `json:"extract_confidence" api:"nullable"`
	ParseConfidence   float64                                                         `json:"parse_confidence" api:"nullable"`
	JSON              parseResponseResultFullResultChunksBlocksGranularConfidenceJSON `json:"-"`
}

// parseResponseResultFullResultChunksBlocksGranularConfidenceJSON contains the
// JSON metadata for the struct
// [ParseResponseResultFullResultChunksBlocksGranularConfidence]
type parseResponseResultFullResultChunksBlocksGranularConfidenceJSON struct {
	ExtractConfidence apijson.Field
	ParseConfidence   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ParseResponseResultFullResultChunksBlocksGranularConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseResultFullResultChunksBlocksGranularConfidenceJSON) RawJSON() string {
	return r.raw
}

// type = 'full'
type ParseResponseResultFullResultType string

const (
	ParseResponseResultFullResultTypeFull ParseResponseResultFullResultType = "full"
)

func (r ParseResponseResultFullResultType) IsKnown() bool {
	switch r {
	case ParseResponseResultFullResultTypeFull:
		return true
	}
	return false
}

type ParseResponseResultFullResultOcr struct {
	Lines []ParseResponseResultFullResultOcrLine `json:"lines" api:"required"`
	Words []ParseResponseResultFullResultOcrWord `json:"words" api:"required"`
	JSON  parseResponseResultFullResultOcrJSON   `json:"-"`
}

// parseResponseResultFullResultOcrJSON contains the JSON metadata for the struct
// [ParseResponseResultFullResultOcr]
type parseResponseResultFullResultOcrJSON struct {
	Lines       apijson.Field
	Words       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParseResponseResultFullResultOcr) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseResultFullResultOcrJSON) RawJSON() string {
	return r.raw
}

type ParseResponseResultFullResultOcrLine struct {
	Bbox BoundingBox `json:"bbox" api:"required"`
	Text string      `json:"text" api:"required"`
	// The index of the chunk that the line belongs to.
	ChunkIndex int64 `json:"chunk_index" api:"nullable"`
	// OCR confidence score between 0 and 1, where 1 indicates highest confidence
	Confidence float64 `json:"confidence" api:"nullable"`
	// The rotation angle in degrees, from 0 to 360, counterclockwise.
	Rotation int64                                    `json:"rotation" api:"nullable"`
	JSON     parseResponseResultFullResultOcrLineJSON `json:"-"`
}

// parseResponseResultFullResultOcrLineJSON contains the JSON metadata for the
// struct [ParseResponseResultFullResultOcrLine]
type parseResponseResultFullResultOcrLineJSON struct {
	Bbox        apijson.Field
	Text        apijson.Field
	ChunkIndex  apijson.Field
	Confidence  apijson.Field
	Rotation    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParseResponseResultFullResultOcrLine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseResultFullResultOcrLineJSON) RawJSON() string {
	return r.raw
}

type ParseResponseResultFullResultOcrWord struct {
	Bbox BoundingBox `json:"bbox" api:"required"`
	Text string      `json:"text" api:"required"`
	// The index of the chunk that the word belongs to.
	ChunkIndex int64 `json:"chunk_index" api:"nullable"`
	// OCR confidence score between 0 and 1, where 1 indicates highest confidence
	Confidence float64 `json:"confidence" api:"nullable"`
	// The rotation angle in degrees, from 0 to 360, counterclockwise.
	Rotation int64                                    `json:"rotation" api:"nullable"`
	JSON     parseResponseResultFullResultOcrWordJSON `json:"-"`
}

// parseResponseResultFullResultOcrWordJSON contains the JSON metadata for the
// struct [ParseResponseResultFullResultOcrWord]
type parseResponseResultFullResultOcrWordJSON struct {
	Bbox        apijson.Field
	Text        apijson.Field
	ChunkIndex  apijson.Field
	Confidence  apijson.Field
	Rotation    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParseResponseResultFullResultOcrWord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseResultFullResultOcrWordJSON) RawJSON() string {
	return r.raw
}

type ParseResponseResultURLResult struct {
	ResultID string `json:"result_id" api:"required"`
	// type = 'url'
	Type ParseResponseResultURLResultType `json:"type" api:"required"`
	URL  string                           `json:"url" api:"required"`
	JSON parseResponseResultURLResultJSON `json:"-"`
}

// parseResponseResultURLResultJSON contains the JSON metadata for the struct
// [ParseResponseResultURLResult]
type parseResponseResultURLResultJSON struct {
	ResultID    apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ParseResponseResultURLResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r parseResponseResultURLResultJSON) RawJSON() string {
	return r.raw
}

func (r ParseResponseResultURLResult) implementsParseResponseResult() {}

// type = 'url'
type ParseResponseResultURLResultType string

const (
	ParseResponseResultURLResultTypeURL ParseResponseResultURLResultType = "url"
)

func (r ParseResponseResultURLResultType) IsKnown() bool {
	switch r {
	case ParseResponseResultURLResultTypeURL:
		return true
	}
	return false
}

// type = 'full'
type ParseResponseResultType string

const (
	ParseResponseResultTypeFull ParseResponseResultType = "full"
	ParseResponseResultTypeURL  ParseResponseResultType = "url"
)

func (r ParseResponseResultType) IsKnown() bool {
	switch r {
	case ParseResponseResultTypeFull, ParseResponseResultTypeURL:
		return true
	}
	return false
}

type ParseNewResponse struct {
	JobID string `json:"job_id" api:"required"`
	// The duration of the parse request in seconds.
	Duration float64 `json:"duration"`
	// The storage URL of the converted PDF file.
	PdfURL string `json:"pdf_url" api:"nullable"`
	// This field can have the runtime type of [ParseResponseResult].
	Result interface{} `json:"result"`
	// The link to the studio pipeline for the document.
	StudioLink string               `json:"studio_link" api:"nullable"`
	Usage      ParseUsage           `json:"usage"`
	JSON       parseNewResponseJSON `json:"-"`
	union      ParseNewResponseUnion
}

// parseNewResponseJSON contains the JSON metadata for the struct
// [ParseNewResponse]
type parseNewResponseJSON struct {
	JobID       apijson.Field
	Duration    apijson.Field
	PdfURL      apijson.Field
	Result      apijson.Field
	StudioLink  apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r parseNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ParseNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ParseNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ParseNewResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ParseResponse], [AsyncParseResponse].
func (r ParseNewResponse) AsUnion() ParseNewResponseUnion {
	return r.union
}

// Union satisfied by [ParseResponse] or [AsyncParseResponse].
type ParseNewResponseUnion interface {
	implementsParseNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ParseNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ParseResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AsyncParseResponse{}),
		},
	)
}

type ParseNewParams struct {
	Body ParseNewParamsBodyUnion `json:"body" api:"required"`
}

func (r ParseNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ParseNewParamsBody struct {
	Input param.Field[interface{}] `json:"input" api:"required"`
	// The configuration options for asynchronous processing (default synchronous).
	Async       param.Field[AsyncConfigV3Param] `json:"async"`
	Enhance     param.Field[EnhanceParam]       `json:"enhance"`
	Formatting  param.Field[FormattingParam]    `json:"formatting"`
	Retrieval   param.Field[RetrievalParam]     `json:"retrieval"`
	Settings    param.Field[SettingsParam]      `json:"settings"`
	Spreadsheet param.Field[SpreadsheetParam]   `json:"spreadsheet"`
}

func (r ParseNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseNewParamsBody) implementsParseNewParamsBodyUnion() {}

// Satisfied by [ParseNewParamsBodySyncParseConfig], [AsyncParseConfigParam],
// [ParseNewParamsBody].
type ParseNewParamsBodyUnion interface {
	implementsParseNewParamsBodyUnion()
}

type ParseNewParamsBodySyncParseConfig struct {
	// For parse/split/extract pipelines, the URL of the document to be processed. You
	// can provide one of the following: 1. A publicly available URL 2. A presigned S3
	// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
	// directly uploading a document 4. A jobid:// prefixed URL obtained from a
	// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
	// API only)
	//
	//	For edit pipelines, this should be a string containing the edit instructions
	Input       param.Field[ParseNewParamsBodySyncParseConfigInputUnion] `json:"input" api:"required"`
	Enhance     param.Field[EnhanceParam]                                `json:"enhance"`
	Formatting  param.Field[FormattingParam]                             `json:"formatting"`
	Retrieval   param.Field[RetrievalParam]                              `json:"retrieval"`
	Settings    param.Field[SettingsParam]                               `json:"settings"`
	Spreadsheet param.Field[SpreadsheetParam]                            `json:"spreadsheet"`
}

func (r ParseNewParamsBodySyncParseConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ParseNewParamsBodySyncParseConfig) implementsParseNewParamsBodyUnion() {}

// For parse/split/extract pipelines, the URL of the document to be processed. You
// can provide one of the following: 1. A publicly available URL 2. A presigned S3
// URL 3. A reducto:// prefixed URL obtained from the /upload endpoint after
// directly uploading a document 4. A jobid:// prefixed URL obtained from a
// previous /parse invocation 5. A list of URLs (for multi-document pipelines, V3
// API only)
//
//	For edit pipelines, this should be a string containing the edit instructions
//
// Satisfied by [shared.UnionString],
// [ParseNewParamsBodySyncParseConfigInputArray], [UploadResponseParam].
type ParseNewParamsBodySyncParseConfigInputUnion interface {
	ImplementsParseNewParamsBodySyncParseConfigInputUnion()
}

type ParseNewParamsBodySyncParseConfigInputArray []string

func (r ParseNewParamsBodySyncParseConfigInputArray) ImplementsParseNewParamsBodySyncParseConfigInputUnion() {
}
