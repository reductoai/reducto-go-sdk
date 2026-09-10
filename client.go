// Package reducto is a Go client for the Reducto document processing API.
package reducto

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math/rand/v2"
	"net"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	BaseURLProduction     = "https://platform.reducto.ai"
	BaseURLEU             = "https://eu.platform.reducto.ai"
	BaseURLAU             = "https://au.platform.reducto.ai"
	DefaultBaseURL        = BaseURLProduction
	DefaultMaxRetries     = 2
	DefaultTimeout        = time.Hour
	DefaultConnectTimeout = 5 * time.Second
	DefaultMaxIdleConns   = 20
	Version               = "0.1.0"

	maxBackoff    = 8 * time.Second
	maxRetryAfter = 60 * time.Second
)

// Client talks to the Reducto API. Create one with New and share it; it is safe for concurrent use.
type Client struct {
	baseURL       string
	apiKey        string
	httpClient    *http.Client
	userAgent     string
	app           string
	appVersion    string
	clientInfo    bool
	maxRetries    int
	timeout       time.Duration
	headers       http.Header
	logger        *slog.Logger
	maxUploadSize int64
	retryBase     time.Duration

	// Per-call state. Options set these on the shallow copy that each call works on.
	responseInto *Response
}

// Option configures a Client. Pass options to New for the whole client, or as trailing
// arguments on any method to apply them to that call only:
//
//	c.Parse(ctx, req, reducto.WithMaxRetries(0), reducto.WithHeader("traceparent", tp))
type Option func(*Client)

// WithBaseURL points the client at a different host: a region (BaseURLEU, BaseURLAU) or a VPC
// deployment. New also reads REDUCTO_BASE_URL.
func WithBaseURL(u string) Option {
	return func(c *Client) { c.baseURL = strings.TrimRight(u, "/") }
}

// WithHTTPClient replaces the default http.Client (5 second connect timeout, no request timeout;
// WithTimeout bounds each attempt instead).
func WithHTTPClient(h *http.Client) Option {
	return func(c *Client) { c.httpClient = h }
}

// WithMaxRetries sets how many times a request is retried. See the package docs for the policy.
func WithMaxRetries(n int) Option {
	return func(c *Client) { c.maxRetries = n }
}

// WithTimeout bounds each attempt, including reading the body. Zero disables it; use ctx instead.
// Default one hour, the same as the Python SDK; sync parse of a large document can run long.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) { c.timeout = d }
}

// WithUserAgent replaces the whole User-Agent header. WithAppInfo is the usual choice;
// it keeps the SDK and platform details and puts your application in front.
func WithUserAgent(ua string) Option {
	return func(c *Client) { c.userAgent = ua }
}

// WithAppInfo names the application built on the SDK. It becomes the first product in
// User-Agent ("reducto-cli/0.1.0 Reducto/Go 0.1.0") and is sent as X-Reducto-Client and
// X-Reducto-Client-Version.
func WithAppInfo(name, version string) Option {
	return func(c *Client) { c.app, c.appVersion = name, version }
}

// WithClientInfo turns the X-Reducto-* client headers on or off. They are on by default
// and carry the same facts as the Python SDK: language, package version, OS, architecture,
// runtime, retry count and read timeout. User-Agent is sent either way.
func WithClientInfo(on bool) Option {
	return func(c *Client) { c.clientInfo = on }
}

// WithHeader adds one header to every request the client (or call) makes. It overrides
// headers the client sets itself, including Authorization.
func WithHeader(key, value string) Option {
	return func(c *Client) {
		c.headers = c.headers.Clone()
		if c.headers == nil {
			c.headers = http.Header{}
		}
		c.headers.Set(key, value)
	}
}

// WithHeaders adds several headers. See WithHeader.
func WithHeaders(h map[string]string) Option {
	return func(c *Client) {
		for k, v := range h {
			WithHeader(k, v)(c)
		}
	}
}

// WithLogger logs each request, response and retry at debug level (retries at info).
// New also honours REDUCTO_LOG=debug|info, which logs to stderr.
func WithLogger(l *slog.Logger) Option {
	return func(c *Client) { c.logger = l }
}

// WithMaxUploadSize makes Upload fail before sending when the file is larger than n bytes.
// Zero (the default) disables the check.
func WithMaxUploadSize(n int64) Option {
	return func(c *Client) { c.maxUploadSize = n }
}

// WithResponseInto stores the status, headers, request id and raw body of the final attempt
// in dst. It is filled on success and on APIError alike. Body is the place to read a field
// the typed structs do not have yet.
//
//	var resp reducto.Response
//	out, err := c.Parse(ctx, req, reducto.WithResponseInto(&resp))
func WithResponseInto(dst *Response) Option {
	return func(c *Client) { c.responseInto = dst }
}

// Response carries the transport-level details of a call. Capture it with WithResponseInto.
type Response struct {
	StatusCode int
	Header     http.Header
	// RequestID is the X-Request-ID (or Request-ID) header, or "" when the server sent none.
	RequestID string
	// Body is the raw response body.
	Body []byte
}

// New returns a Client. An empty apiKey falls back to the REDUCTO_API_KEY environment variable.
func New(apiKey string, opts ...Option) *Client {
	if apiKey == "" {
		apiKey = os.Getenv("REDUCTO_API_KEY")
	}
	c := &Client{
		baseURL:    DefaultBaseURL,
		apiKey:     apiKey,
		httpClient: defaultHTTPClient(),
		clientInfo: true,
		maxRetries: DefaultMaxRetries,
		timeout:    DefaultTimeout,
		logger:     loggerFromEnv(),
		retryBase:  500 * time.Millisecond,
	}
	if u := os.Getenv("REDUCTO_BASE_URL"); u != "" {
		c.baseURL = strings.TrimRight(u, "/")
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

// WithOptions returns a shallow copy of c with opts applied. The copy shares the http.Client.
func (c *Client) WithOptions(opts ...Option) *Client {
	cc := *c
	for _, o := range opts {
		o(&cc)
	}
	return &cc
}

func (c *Client) with(opts []Option) *Client {
	if len(opts) == 0 {
		return c
	}
	return c.WithOptions(opts...)
}

func defaultHTTPClient() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.DialContext = (&net.Dialer{Timeout: DefaultConnectTimeout, KeepAlive: 30 * time.Second}).DialContext
	t.MaxIdleConnsPerHost = DefaultMaxIdleConns
	return &http.Client{Transport: t}
}

func loggerFromEnv() *slog.Logger {
	var level slog.Level
	switch strings.ToLower(os.Getenv("REDUCTO_LOG")) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	default:
		return nil
	}
	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level}))
}

// ---- transport ----

type request struct {
	method      string
	path        string
	query       url.Values
	contentType string
	body        []byte
	// stream, when set, supplies the body for each attempt. length is the exact byte count.
	stream func() (body io.Reader, length int64, err error)
}

type response struct {
	status int
	header http.Header
	body   []byte
}

func (c *Client) do(ctx context.Context, method, path string, query url.Values, body any, out any, opts []Option) error {
	req := request{method: method, path: path, query: query}
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("reducto: encode request: %w", err)
		}
		req.body, req.contentType = b, "application/json"
	}
	return c.with(opts).send(ctx, req, out)
}

func (c *Client) send(ctx context.Context, r request, out any) error {
	var lastErr error
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		if attempt > 0 {
			d := c.backoff(attempt, lastErr)
			c.logInfo(ctx, "reducto: retrying", "method", r.method, "path", r.path, "attempt", attempt, "wait", d, "cause", lastErr)
			if err := sleep(ctx, d); err != nil {
				return err
			}
		}
		res, err := c.attempt(ctx, r, attempt)
		if err != nil {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			lastErr = err
			if !c.shouldRetry(r, err) {
				return err
			}
			continue
		}
		c.capture(res)
		if res.status < 200 || res.status > 299 {
			apiErr := newAPIError(r.method, r.path, res.status, res.header, res.body)
			lastErr = apiErr
			if retryableResponse(res.status, res.header) {
				continue
			}
			return apiErr
		}
		return decode(res.body, out)
	}
	return lastErr
}

func (c *Client) attempt(ctx context.Context, r request, n int) (*response, error) {
	actx := ctx
	if c.timeout > 0 {
		var cancel context.CancelFunc
		actx, cancel = context.WithTimeout(ctx, c.timeout)
		defer cancel()
	}
	req, err := c.newRequest(actx, r, n)
	if err != nil {
		return nil, err
	}
	c.logDebug(ctx, "reducto: request", "method", r.method, "path", r.path, "attempt", n)
	start := time.Now()
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, c.transportError(ctx, actx, r, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, c.transportError(ctx, actx, r, err)
	}
	c.logDebug(ctx, "reducto: response", "method", r.method, "path", r.path, "status", resp.StatusCode,
		"duration", time.Since(start), "request_id", requestID(resp.Header))
	return &response{status: resp.StatusCode, header: resp.Header, body: body}, nil
}

func (c *Client) newRequest(ctx context.Context, r request, attempt int) (*http.Request, error) {
	u := c.baseURL + r.path
	if len(r.query) > 0 {
		u += "?" + r.query.Encode()
	}
	var rd io.Reader
	length := int64(-1)
	if r.stream != nil {
		body, n, err := r.stream()
		if err != nil {
			return nil, err
		}
		rd, length = body, n
	} else if r.body != nil {
		rd = bytes.NewReader(r.body)
	}
	req, err := http.NewRequestWithContext(ctx, r.method, u, rd)
	if err != nil {
		return nil, err
	}
	if length >= 0 {
		req.ContentLength = length
	}
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}
	req.Header.Set("Accept", "application/json")
	c.setClientHeaders(req.Header, attempt)
	if r.contentType != "" {
		req.Header.Set("Content-Type", r.contentType)
	}
	for k, vs := range c.headers {
		req.Header[http.CanonicalHeaderKey(k)] = vs
	}
	return req, nil
}

// ---- client identification ----
//
// Header names and values follow the Python SDK (platform_headers in _base_client.py), so
// one log query covers every SDK.

const sdkProduct = "Reducto/Go " + Version

func (c *Client) setClientHeaders(h http.Header, attempt int) {
	h.Set("User-Agent", c.userAgentValue())
	if !c.clientInfo {
		return
	}
	h.Set("X-Reducto-Lang", "go")
	h.Set("X-Reducto-Package-Version", Version)
	h.Set("X-Reducto-OS", platformOS())
	h.Set("X-Reducto-Arch", platformArch())
	h.Set("X-Reducto-Runtime", "go")
	h.Set("X-Reducto-Runtime-Version", strings.TrimPrefix(runtime.Version(), "go"))
	h.Set("X-Reducto-Async", "false")
	h.Set("X-Reducto-Retry-Count", strconv.Itoa(attempt))
	if c.timeout > 0 {
		h.Set("X-Reducto-Read-Timeout", strconv.FormatFloat(c.timeout.Seconds(), 'f', -1, 64))
	}
	if c.app != "" {
		h.Set("X-Reducto-Client", c.app)
		h.Set("X-Reducto-Client-Version", c.appVersion)
	}
}

func (c *Client) userAgentValue() string {
	if c.userAgent != "" {
		return c.userAgent
	}
	if c.app != "" {
		return c.app + "/" + c.appVersion + " " + sdkProduct
	}
	return sdkProduct
}

func platformOS() string {
	switch runtime.GOOS {
	case "darwin":
		return "MacOS"
	case "ios":
		return "iOS"
	case "windows":
		return "Windows"
	case "linux":
		return "Linux"
	case "android":
		return "Android"
	case "freebsd":
		return "FreeBSD"
	case "openbsd":
		return "OpenBSD"
	case "":
		return "Unknown"
	}
	return runtime.GOOS
}

func platformArch() string {
	switch runtime.GOARCH {
	case "arm64":
		return "arm64"
	case "arm":
		return "arm"
	case "amd64":
		return "x64"
	case "386":
		return "x32"
	case "":
		return "unknown"
	}
	return runtime.GOARCH
}

func (c *Client) transportError(ctx, actx context.Context, r request, err error) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	var ne net.Error
	if errors.Is(actx.Err(), context.DeadlineExceeded) || (errors.As(err, &ne) && ne.Timeout()) {
		return &APITimeoutError{Method: r.method, Path: r.path, After: c.timeout, Err: err}
	}
	return &APIConnectionError{Method: r.method, Path: r.path, Err: err}
}

func (c *Client) capture(res *response) {
	if c.responseInto == nil {
		return
	}
	*c.responseInto = Response{StatusCode: res.status, Header: res.header, RequestID: requestID(res.header), Body: res.body}
}

// shouldRetry decides whether a failed attempt is worth repeating.
//
// Any method is retried on 408, 429 and 5xx, and every method is retried on timeouts and
// connection errors. This matches the published SDKs. A POST retry can create a second job
// when the server accepted the first request; set WithMaxRetries(0) where that matters.
func (c *Client) shouldRetry(r request, err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return retryableResponse(apiErr.StatusCode, apiErr.Header)
	}
	return true
}

// retryableResponse lets an explicit x-should-retry header win over the status code,
// as the Python SDK does.
func retryableResponse(status int, h http.Header) bool {
	switch strings.ToLower(h.Get("x-should-retry")) {
	case "true":
		return true
	case "false":
		return false
	}
	return status == http.StatusRequestTimeout || status == http.StatusTooManyRequests || status >= 500
}

func decode(body []byte, out any) error {
	if out == nil || len(bytes.TrimSpace(body)) == 0 {
		return nil
	}
	if raw, ok := out.(*json.RawMessage); ok {
		*raw = append((*raw)[:0], body...)
		return nil
	}
	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("reducto: decode response: %w (body: %s)", err, jsonHead(body))
	}
	return nil
}

// backoff honours Retry-After when it is between 0 and maxRetryAfter, else grows
// exponentially from retryBase and is capped at maxBackoff (plus up to 50% jitter).
func (c *Client) backoff(attempt int, lastErr error) time.Duration {
	var apiErr *APIError
	if errors.As(lastErr, &apiErr) {
		if d, ok := apiErr.RetryAfter(); ok && d > 0 && d <= maxRetryAfter {
			return d
		}
	}
	d := c.retryBase << min(attempt-1, 16)
	if d > maxBackoff || d <= 0 {
		d = maxBackoff
	}
	return d + rand.N(d/2+1)
}

func sleep(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

func (c *Client) logDebug(ctx context.Context, msg string, args ...any) {
	if c.logger != nil {
		c.logger.DebugContext(ctx, msg, args...)
	}
}

func (c *Client) logInfo(ctx context.Context, msg string, args ...any) {
	if c.logger != nil {
		c.logger.InfoContext(ctx, msg, args...)
	}
}

// Ptr returns a pointer to v. Optional request fields are pointers so that false and 0 can be sent explicitly.
func Ptr[T any](v T) *T { return &v }
