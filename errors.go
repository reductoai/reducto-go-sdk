package reducto

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// APIError is returned for any non-2xx response.
type APIError struct {
	StatusCode int
	Method     string
	Path       string
	Body       []byte
	// Header holds the response headers.
	Header http.Header
	// RequestID is the X-Request-ID (or Request-ID) response header, or "" when none was sent.
	RequestID string
	// Detail is set when the server returned a structured ErrorDetail.
	Detail *ErrorDetail
	// Validation is set for 422 responses.
	Validation []ValidationError
	// Message is the best human-readable summary the body offered.
	Message string
}

func (e *APIError) Error() string {
	var b strings.Builder
	fmt.Fprintf(&b, "reducto: %s %s: %d", e.Method, e.Path, e.StatusCode)
	if e.Detail != nil {
		fmt.Fprintf(&b, " %s", e.Detail.Name)
	}
	if e.Message != "" {
		fmt.Fprintf(&b, ": %s", e.Message)
	}
	if e.RequestID != "" {
		fmt.Fprintf(&b, " (request id %s)", e.RequestID)
	}
	return b.String()
}

// Code returns the machine-readable error name, or "" if the body had none.
func (e *APIError) Code() ErrorCode {
	if e.Detail == nil {
		return ""
	}
	return e.Detail.Name
}

// RetryAfter parses Retry-After-Ms, then Retry-After (seconds or an HTTP date).
// ok is false when both are absent or malformed.
func (e *APIError) RetryAfter() (d time.Duration, ok bool) {
	if ms, err := strconv.ParseFloat(strings.TrimSpace(e.Header.Get("Retry-After-Ms")), 64); err == nil {
		return time.Duration(ms * float64(time.Millisecond)), true
	}
	return parseRetryAfter(e.Header.Get("Retry-After"))
}

// APIConnectionError is returned when no HTTP response was received: DNS failure, refused
// connection, reset, or a malformed response. Err holds the underlying transport error.
type APIConnectionError struct {
	Method string
	Path   string
	Err    error
}

func (e *APIConnectionError) Error() string {
	return fmt.Sprintf("reducto: %s %s: connection error: %v", e.Method, e.Path, e.Err)
}

func (e *APIConnectionError) Unwrap() error { return e.Err }

// APITimeoutError is returned when a single attempt exceeds the client timeout (WithTimeout).
// It matches errors.Is(err, context.DeadlineExceeded). A deadline on the caller's ctx is
// returned as the plain ctx error instead.
type APITimeoutError struct {
	Method string
	Path   string
	// After is the per-attempt timeout that elapsed.
	After time.Duration
	Err   error
}

func (e *APITimeoutError) Error() string {
	return fmt.Sprintf("reducto: %s %s: timed out after %s", e.Method, e.Path, e.After)
}

func (e *APITimeoutError) Unwrap() error { return e.Err }

func (e *APITimeoutError) Is(target error) bool { return target == context.DeadlineExceeded }

// Timeout reports true so the error satisfies the net.Error convention.
func (e *APITimeoutError) Timeout() bool { return true }

func newAPIError(method, path string, status int, header http.Header, body []byte) *APIError {
	e := &APIError{
		StatusCode: status,
		Method:     method,
		Path:       path,
		Body:       body,
		Header:     header,
		RequestID:  requestID(header),
	}
	var envelope struct {
		Error  *ErrorDetail    `json:"error"`
		Detail json.RawMessage `json:"detail"`
	}
	if json.Unmarshal(body, &envelope) != nil {
		e.Message = strings.TrimSpace(string(body))
		return e
	}
	if envelope.Error != nil {
		e.Detail = envelope.Error
		e.Message = envelope.Error.Message
		return e
	}
	switch jsonKind(envelope.Detail) {
	case 's':
		_ = json.Unmarshal(envelope.Detail, &e.Message)
	case 'o':
		var d ErrorDetail
		if json.Unmarshal(envelope.Detail, &d) == nil && d.Message != "" {
			e.Detail = &d
			e.Message = d.Message
		}
	case 'a':
		if json.Unmarshal(envelope.Detail, &e.Validation) == nil && len(e.Validation) > 0 {
			e.Message = e.Validation[0].Msg
		}
	}
	if e.Message == "" && len(body) > 0 {
		e.Message = jsonHead(body)
	}
	return e
}

func requestID(h http.Header) string {
	if v := h.Get("X-Request-ID"); v != "" {
		return v
	}
	return h.Get("Request-ID")
}

func parseRetryAfter(v string) (time.Duration, bool) {
	v = strings.TrimSpace(v)
	if v == "" {
		return 0, false
	}
	if secs, err := strconv.ParseFloat(v, 64); err == nil {
		return time.Duration(secs * float64(time.Second)), true
	}
	if t, err := http.ParseTime(v); err == nil {
		return time.Until(t), true
	}
	return 0, false
}
