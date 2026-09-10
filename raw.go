package reducto

import (
	"context"
	"encoding/json"
	"net/url"
)

// Do sends an arbitrary JSON request and returns the raw response body.
// It is the escape hatch for fields or endpoints the typed API does not cover.
func (c *Client) Do(ctx context.Context, method, path string, query url.Values, body any, opts ...Option) (json.RawMessage, error) {
	var out json.RawMessage
	if err := c.do(ctx, method, path, query, body, &out, opts); err != nil {
		return nil, err
	}
	return out, nil
}
