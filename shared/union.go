// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsAsyncParseConfigInputUnionParam()                 {}
func (UnionString) ImplementsParseNewParamsBodySyncParseConfigInputUnion()     {}
func (UnionString) ImplementsAsyncExtractConfigInputUnionParam()               {}
func (UnionString) ImplementsExtractNewParamsBodySyncExtractConfigInputUnion() {}
func (UnionString) ImplementsSplitNewParamsInputUnion()                        {}
func (UnionString) ImplementsSplitAsyncNewParamsInputUnion()                   {}
func (UnionString) ImplementsEditSubmitParamsDocumentURLUnion()                {}
func (UnionString) ImplementsEditAsyncNewParamsDocumentURLUnion()              {}
func (UnionString) ImplementsPipelineNewParamsInputUnion()                     {}
func (UnionString) ImplementsPipelineAsyncNewParamsInputUnion()                {}
func (UnionString) ImplementsClassifyNewParamsInputUnion()                     {}

type UnionInt int64

func (UnionInt) ImplementsSpreadsheetSplitLargeTablesSizeUnionParam() {}
