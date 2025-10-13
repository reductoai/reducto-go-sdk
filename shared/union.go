// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsSplitRunParamsDocumentURLUnion()                               {}
func (UnionString) ImplementsSplitRunJobParamsDocumentURLUnion()                            {}
func (UnionString) ImplementsParseConfigDocumentURLUnionParam()                             {}
func (UnionString) ImplementsParseRunParamsBodyParseConfigDocumentURLUnion()                {}
func (UnionString) ImplementsParseRunParamsBodySyncParseConfigInputUnion()                  {}
func (UnionString) ImplementsParseRunJobParamsBodyAsyncParseConfigNewDocumentURLUnion()     {}
func (UnionString) ImplementsParseRunJobParamsBodyAsyncParseConfigInputUnion()              {}
func (UnionString) ImplementsExtractConfigDocumentURLUnionParam()                           {}
func (UnionString) ImplementsExtractRunParamsBodyExtractConfigDocumentURLUnion()            {}
func (UnionString) ImplementsExtractRunParamsBodySyncExtractConfigInputUnion()              {}
func (UnionString) ImplementsExtractRunJobParamsBodyAsyncExtractConfigNewDocumentURLUnion() {}
func (UnionString) ImplementsExtractRunJobParamsBodyAsyncExtractConfigInputUnion()          {}
func (UnionString) ImplementsEditRunParamsDocumentURLUnion()                                {}
func (UnionString) ImplementsEditRunJobParamsDocumentURLUnion()                             {}
func (UnionString) ImplementsPipelineRunParamsDocumentURLUnion()                            {}
func (UnionString) ImplementsPipelineRunJobParamsDocumentURLUnion()                         {}
