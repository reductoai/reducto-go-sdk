package reducto

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// TypedExtract is an extract response whose items were decoded into T.
type TypedExtract[T any] struct {
	Result []T
	// V3ExtractResponse is set when the API answered in the v3 shape.
	V3ExtractResponse *V3ExtractResponse
	// ExtractResponse is set when the API answered in the legacy shape. Which shape you get
	// depends on the account and on settings such as citations.
	ExtractResponse *ExtractResponse
}

// Usage returns the usage block of the underlying response.
func (t *TypedExtract[T]) Usage() ExtractUsage {
	if t.ExtractResponse != nil {
		return t.ExtractResponse.Usage
	}
	return t.V3ExtractResponse.Usage
}

// TypedExtractError is returned by ExtractAs and ValidateExtract when the result cannot be
// decoded into T. Response is the API response as received. Err holds the JSON decode error
// when an item did not match.
type TypedExtractError struct {
	Reason   string
	Response any
	Err      error
}

func (e *TypedExtractError) Error() string { return "reducto: extract: " + e.Reason }

func (e *TypedExtractError) Unwrap() error { return e.Err }

// ExtractAs calls Extract with schema as instructions.schema and decodes the items into T.
// schema is a JSON Schema for one item, as a map or any value that marshals to one. Any
// Instructions.Schema already set on req is replaced; other fields are sent as given.
//
// It returns a *TypedExtractError when the document was queued as a job or the result came back
// as a URL; see ValidateExtract for the follow-up in those cases.
//
//	type Invoice struct {
//		Total float64 `json:"total"`
//	}
//	out, err := reducto.ExtractAs[Invoice](ctx, client, map[string]any{
//		"type": "object", "properties": map[string]any{"total": map[string]any{"type": "number"}},
//	}, &reducto.SyncExtractConfig{Input: reducto.DocumentInputFromString(url)})
func ExtractAs[T any](ctx context.Context, c *Client, schema any, req *SyncExtractConfig, opts ...Option) (*TypedExtract[T], error) {
	body := *req
	var ins Instructions
	if req.Instructions != nil {
		ins = *req.Instructions
	}
	ins.Schema = schema
	body.Instructions = &ins
	out, err := c.Extract(ctx, &body, opts...)
	if err != nil {
		return nil, err
	}
	return ValidateExtract[T](out)
}

// ValidateExtract decodes the items of an extract response into T.
//
// response may be an *ExtractOutput (from Extract), a *JobResult (from a finished job), or an
// *ExtractResponse / *V3ExtractResponse directly; values are accepted too. A single object is
// treated as a list of one. It returns a *TypedExtractError when the response is a queued job,
// a URL result, not an extract response at all, or when an item does not decode into T.
func ValidateExtract[T any](response any) (*TypedExtract[T], error) {
	var (
		legacy *ExtractResponse
		v3     *V3ExtractResponse
		queued *AsyncExtractResponse
	)
	switch r := response.(type) {
	case *ExtractOutput:
		if r == nil {
			return nil, &TypedExtractError{Reason: "expected an extract response, got nil", Response: response}
		}
		legacy, v3, queued = r.ExtractResponse, r.V3ExtractResponse, r.AsyncExtractResponse
	case ExtractOutput:
		legacy, v3, queued = r.ExtractResponse, r.V3ExtractResponse, r.AsyncExtractResponse
	case *JobResult:
		if r == nil {
			return nil, &TypedExtractError{Reason: "expected an extract response, got nil", Response: response}
		}
		legacy, v3 = r.ExtractResponse, r.V3ExtractResponse
	case JobResult:
		legacy, v3 = r.ExtractResponse, r.V3ExtractResponse
	case *ExtractResponse:
		legacy = r
	case ExtractResponse:
		legacy = &r
	case *V3ExtractResponse:
		v3 = r
	case V3ExtractResponse:
		v3 = &r
	default:
		return nil, &TypedExtractError{Reason: fmt.Sprintf("expected an extract response, got %T", response), Response: response}
	}

	if queued != nil {
		return nil, &TypedExtractError{
			Reason:   fmt.Sprintf("the document was queued as job %s; call WaitForJob, then ValidateExtract(job.Result)", queued.JobID),
			Response: response,
		}
	}
	var raw any
	switch {
	case legacy != nil:
		if u := legacy.Result.UrlResult; u != nil {
			return nil, urlResultError(u, response)
		}
		raw = legacy.Result.AnyArray
	case v3 != nil:
		raw = v3.Result
		if m, ok := raw.(map[string]any); ok && m["type"] == "url" {
			var u UrlResult
			if b, err := json.Marshal(m); err == nil && json.Unmarshal(b, &u) == nil {
				return nil, urlResultError(&u, response)
			}
		}
	default:
		return nil, &TypedExtractError{Reason: "expected an extract response, got an empty result", Response: response}
	}

	items, ok := raw.([]any)
	if !ok {
		items = []any{raw}
	}
	result := make([]T, 0, len(items))
	for i, item := range items {
		b, err := json.Marshal(item)
		if err != nil {
			return nil, &TypedExtractError{Reason: fmt.Sprintf("item %d: %v", i, err), Response: response, Err: err}
		}
		var v T
		if err := json.Unmarshal(b, &v); err != nil {
			reason := fmt.Sprintf("item %d does not match %T: %s", i, v, strings.TrimPrefix(err.Error(), "json: "))
			return nil, &TypedExtractError{Reason: reason, Response: response, Err: err}
		}
		result = append(result, v)
	}
	return &TypedExtract[T]{Result: result, V3ExtractResponse: v3, ExtractResponse: legacy}, nil
}

func urlResultError(u *UrlResult, response any) *TypedExtractError {
	return &TypedExtractError{
		Reason:   fmt.Sprintf("the result was returned as a URL (%s); fetch it and decode the items yourself", u.URL),
		Response: response,
	}
}
