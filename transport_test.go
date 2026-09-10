package reducto

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func hijackAndClose(w http.ResponseWriter) {
	conn, _, err := w.(http.Hijacker).Hijack()
	if err == nil {
		conn.Close()
	}
}

func TestHeadersClientAndPerCall(t *testing.T) {
	var seen []http.Header
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		seen = append(seen, r.Header.Clone())
		io.WriteString(w, `"ok"`)
	})
	c = c.WithOptions(WithHeaders(map[string]string{"X-Org": "acme", "Traceparent": "00-abc"}))
	if _, err := c.Version(context.Background(), WithHeader("X-Call", "1"), WithHeader("X-Org", "override")); err != nil {
		t.Fatal(err)
	}
	if _, err := c.Version(context.Background()); err != nil {
		t.Fatal(err)
	}
	first, second := seen[0], seen[1]
	if first.Get("X-Call") != "1" || first.Get("X-Org") != "override" || first.Get("Traceparent") != "00-abc" {
		t.Errorf("per-call headers = %v", first)
	}
	if second.Get("X-Call") != "" || second.Get("X-Org") != "acme" {
		t.Errorf("per-call header leaked into client: %v", second)
	}
	if first.Get("Authorization") != "Bearer test-key" {
		t.Errorf("auth lost: %v", first)
	}
}

func TestResponseCaptureAndRequestID(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-ID", "req_123")
		if r.URL.Path == "/fail" {
			w.WriteHeader(500)
			return
		}
		io.WriteString(w, `"ok"`)
	})
	var resp Response
	if _, err := c.Version(context.Background(), WithResponseInto(&resp)); err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 || resp.RequestID != "req_123" || resp.Header.Get("X-Request-Id") != "req_123" {
		t.Fatalf("resp = %+v", resp)
	}
	_, err := c.Do(context.Background(), "GET", "/fail", nil, nil, WithResponseInto(&resp))
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.RequestID != "req_123" || resp.StatusCode != 500 {
		t.Fatalf("err = %v, resp = %+v", err, resp)
	}
	if !strings.Contains(err.Error(), "req_123") {
		t.Fatal(err)
	}
}

func TestPostIsRetriedAfterConnectionError(t *testing.T) {
	var calls atomic.Int32
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		hijackAndClose(w)
	})
	c.maxRetries = 2
	c.httpClient.Transport.(*http.Transport).DisableKeepAlives = true

	_, err := c.ParseAsync(context.Background(), &AsyncParseConfig{Input: DocumentInputFromString("x")})
	var connErr *APIConnectionError
	if !errors.As(err, &connErr) {
		t.Fatalf("err = %T %v", err, err)
	}
	if calls.Load() != 3 {
		t.Fatalf("POST should retry: %d calls", calls.Load())
	}

	calls.Store(0)
	_, err = c.Version(context.Background())
	if !errors.As(err, &connErr) || calls.Load() != 3 {
		t.Fatalf("GET should retry: err = %v, calls = %d", err, calls.Load())
	}
}

func TestPostIsRetriedOnDialError(t *testing.T) {
	srv := httptest.NewServer(http.NotFoundHandler())
	addr := srv.URL
	srv.Close()
	c := New("k", WithBaseURL(addr), WithMaxRetries(1))
	c.retryBase = time.Millisecond
	start := time.Now()
	_, err := c.ParseAsync(context.Background(), &AsyncParseConfig{Input: DocumentInputFromString("x")})
	var connErr *APIConnectionError
	if !errors.As(err, &connErr) {
		t.Fatalf("err = %T %v", err, err)
	}
	if time.Since(start) < time.Millisecond {
		t.Fatal("expected at least one backoff sleep, so at least one retry")
	}
}

func TestPostIsRetriedOn5xx(t *testing.T) {
	var calls atomic.Int32
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if calls.Add(1) == 1 {
			w.WriteHeader(503)
			return
		}
		io.WriteString(w, `{"job_id": "j"}`)
	})
	c.maxRetries = 1
	out, err := c.ParseAsync(context.Background(), &AsyncParseConfig{Input: DocumentInputFromString("x")})
	if err != nil || out.JobID != "j" || calls.Load() != 2 {
		t.Fatalf("out = %+v, err = %v, calls = %d", out, err, calls.Load())
	}
}

func TestRetryAfterIsBoundedAndParsesDates(t *testing.T) {
	c := New("k")
	c.retryBase = time.Millisecond
	mk := func(v string) *APIError {
		return &APIError{StatusCode: 429, Header: http.Header{"Retry-After": {v}}}
	}
	if d := c.backoff(1, mk("3600")); d > maxBackoff*2 {
		t.Fatalf("huge Retry-After was honoured: %s", d)
	}
	if d := c.backoff(1, mk("2")); d != 2*time.Second {
		t.Fatalf("Retry-After: 2 -> %s", d)
	}
	date := time.Now().Add(3 * time.Second).UTC().Format(http.TimeFormat)
	if d := c.backoff(1, mk(date)); d < 2*time.Second || d > 3*time.Second {
		t.Fatalf("HTTP-date Retry-After -> %s", d)
	}
	if d := c.backoff(1, mk("garbage")); d > maxBackoff*2 {
		t.Fatalf("garbage -> %s", d)
	}
	if d := c.backoff(40, nil); d > maxBackoff*2 || d <= 0 {
		t.Fatalf("large attempt -> %s", d)
	}
}

func TestTimeoutErrorType(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
		case <-time.After(2 * time.Second):
		}
	})
	_, err := c.Version(context.Background(), WithTimeout(20*time.Millisecond))
	var te *APITimeoutError
	if !errors.As(err, &te) || !errors.Is(err, context.DeadlineExceeded) || te.After != 20*time.Millisecond {
		t.Fatalf("err = %T %v", err, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	_, err = c.Version(ctx)
	if errors.As(err, &te) || !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("caller deadline should surface raw: %T %v", err, err)
	}
}

func TestWaitForJobTimeoutErrorType(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"status": "Pending", "progress": 0.1}`)
	})
	job, err := c.WaitForJob(context.Background(), "j", &WaitOptions{Interval: time.Millisecond, Timeout: 30 * time.Millisecond})
	var jt *JobTimeoutError
	if !errors.As(err, &jt) || jt.JobID != "j" || !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("err = %T %v", err, err)
	}
	if job == nil || job.Status != JobStatusPending {
		t.Fatalf("last job state lost: %+v", job)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	_, err = c.WaitForJob(ctx, "j", &WaitOptions{Interval: time.Millisecond})
	if errors.As(err, &jt) || !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("caller deadline should surface raw: %T %v", err, err)
	}
}

func TestIterJobsFollowsCursor(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Query().Get("cursor") {
		case "":
			io.WriteString(w, `{"jobs": [{"job_id": "a", "status": "Completed", "type": "Parse", "created_at": "2024-01-01T00:00:00"}], "next_cursor": "c2"}`)
		case "c2":
			io.WriteString(w, `{"jobs": [{"job_id": "b", "status": "Completed", "type": "Parse", "created_at": "2024-01-01T00:00:00"}], "next_cursor": null}`)
		default:
			io.WriteString(w, `not json`)
		}
		if r.URL.Query().Get("limit") != "1" {
			t.Errorf("limit not forwarded: %s", r.URL.RawQuery)
		}
	})
	var ids []string
	for job, err := range c.IterJobs(context.Background(), &ListJobsParams{Limit: Ptr[int64](1)}) {
		if err != nil {
			t.Fatal(err)
		}
		ids = append(ids, job.JobID)
	}
	if strings.Join(ids, ",") != "a,b" {
		t.Fatalf("ids = %v", ids)
	}

	for _, err := range c.IterJobs(context.Background(), &ListJobsParams{Cursor: Ptr("bad"), Limit: Ptr[int64](1)}) {
		if err == nil {
			t.Fatal("expected decode error for empty body")
		}
	}
}

func TestBaseURLFromEnv(t *testing.T) {
	t.Setenv("REDUCTO_BASE_URL", "https://vpc.example.com/")
	if c := New("k"); c.baseURL != "https://vpc.example.com" {
		t.Fatalf("baseURL = %q", c.baseURL)
	}
	if c := New("k", WithBaseURL("https://other")); c.baseURL != "https://other" {
		t.Fatalf("option should win: %q", c.baseURL)
	}
}

func TestUploadStreamsSeekableFileAndRetries(t *testing.T) {
	path := filepath.Join(t.TempDir(), "doc.pdf")
	content := strings.Repeat("%PDF", 1000)
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	var calls atomic.Int32
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength <= 0 {
			t.Errorf("expected known Content-Length, got %d", r.ContentLength)
		}
		f, hdr, err := r.FormFile("file")
		if err != nil {
			t.Fatal(err)
		}
		data, _ := io.ReadAll(f)
		if hdr.Filename != "doc.pdf" || string(data) != content {
			t.Errorf("attempt %d: file = %s (%d bytes)", calls.Load(), hdr.Filename, len(data))
		}
		if calls.Add(1) == 1 {
			w.WriteHeader(503)
			return
		}
		io.WriteString(w, `{"file_id": "reducto://abc"}`)
	})
	up, err := c.UploadFile(context.Background(), path, WithMaxRetries(1))
	if err != nil || up.FileID != "reducto://abc" || calls.Load() != 2 {
		t.Fatalf("up = %+v, err = %v, calls = %d", up, err, calls.Load())
	}
}

func TestUploadGuards(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		t.Error("server should not be reached")
	})
	if _, err := c.Upload(context.Background(), strings.NewReader(""), "e.pdf", nil); !errors.Is(err, ErrEmptyUpload) {
		t.Fatalf("empty -> %v", err)
	}
	_, err := c.Upload(context.Background(), strings.NewReader("12345"), "e.pdf", nil, WithMaxUploadSize(4))
	if !errors.Is(err, ErrUploadTooLarge) {
		t.Fatalf("too large -> %v", err)
	}
}

func svixHeaders(secret, id string, ts int64, payload []byte) http.Header {
	key, _ := base64.StdEncoding.DecodeString(strings.TrimPrefix(secret, "whsec_"))
	mac := hmac.New(sha256.New, key)
	fmt.Fprintf(mac, "%s.%d.", id, ts)
	mac.Write(payload)
	sig := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return http.Header{
		"Svix-Id":        {id},
		"Svix-Timestamp": {strconv.FormatInt(ts, 10)},
		"Svix-Signature": {"v1,AAAA v1," + sig},
	}
}

func TestVerifyWebhook(t *testing.T) {
	secret := "whsec_" + base64.StdEncoding.EncodeToString([]byte("0123456789abcdef0123456789abcdef"))
	payload := []byte(`{"job_id":"j1","status":"Completed"}`)
	now := time.Now()
	h := svixHeaders(secret, "msg_1", now.Unix(), payload)

	if err := verifyWebhook(secret, payload, h, now); err != nil {
		t.Fatal(err)
	}
	if err := verifyWebhook(strings.TrimPrefix(secret, "whsec_"), payload, h, now); err != nil {
		t.Fatalf("bare secret: %v", err)
	}
	var verr *WebhookVerificationError
	if err := verifyWebhook(secret, []byte("tampered"), h, now); !errors.As(err, &verr) {
		t.Fatalf("tampered -> %v", err)
	}
	if err := verifyWebhook(secret, payload, h, now.Add(10*time.Minute)); !errors.As(err, &verr) || !strings.Contains(err.Error(), "tolerance") {
		t.Fatalf("old -> %v", err)
	}
	if err := verifyWebhook(secret, payload, http.Header{}, now); !errors.As(err, &verr) {
		t.Fatalf("missing headers -> %v", err)
	}
	std := http.Header{"Webhook-Id": h["Svix-Id"], "Webhook-Timestamp": h["Svix-Timestamp"], "Webhook-Signature": h["Svix-Signature"]}
	if err := verifyWebhook(secret, payload, std, now); err != nil {
		t.Fatalf("standard-webhooks header names: %v", err)
	}
}
