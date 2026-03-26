// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsAsyncParseConfigInputUnionParam()                 {}
func (UnionString) ImplementsParseRunParamsBodySyncParseConfigInputUnion()     {}
func (UnionString) ImplementsAsyncExtractConfigInputUnionParam()               {}
func (UnionString) ImplementsExtractRunParamsBodySyncExtractConfigInputUnion() {}
func (UnionString) ImplementsSplitRunParamsInputUnion()                        {}
func (UnionString) ImplementsSplitRunJobParamsInputUnion()                     {}
func (UnionString) ImplementsEditRunParamsDocumentURLUnion()                   {}
func (UnionString) ImplementsEditRunJobParamsDocumentURLUnion()                {}
func (UnionString) ImplementsPipelineRunParamsInputUnion()                     {}
func (UnionString) ImplementsPipelineRunJobParamsInputUnion()                  {}
func (UnionString) ImplementsClassifyRunParamsInputUnion()                     {}

type UnionInt int64

func (UnionInt) ImplementsSplitLargeTablesSizeUnionParam() {}
