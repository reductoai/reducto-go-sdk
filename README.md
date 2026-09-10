# Reducto Go API library

[![Go Reference](https://pkg.go.dev/badge/github.com/reductoai/reducto-go-sdk.svg)](https://pkg.go.dev/github.com/reductoai/reducto-go-sdk)

The Reducto Go library provides convenient access to the Reducto REST API from any Go 1.23+
application. It covers every endpoint, includes types for all request and response fields,
and depends only on the standard library.

## Documentation

The REST API documentation can be found on [docs.reductoai.com](https://docs.reductoai.com).
The full API of this library can be found on
[pkg.go.dev](https://pkg.go.dev/github.com/reductoai/reducto-go-sdk).

## Installation

```sh
go get github.com/reductoai/reducto-go-sdk
```

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"log"

	reducto "github.com/reductoai/reducto-go-sdk"
)

func main() {
	ctx := context.Background()
	client := reducto.New("") // falls back to $REDUCTO_API_KEY

	out, err := client.Parse(ctx, &reducto.SyncParseConfig{
		Input: reducto.DocumentInputFromString("https://example.com/report.pdf"),
		Settings: &reducto.Settings{
			OCRSystem: reducto.SettingsOCRSystemStandard,
			PageRange: &reducto.PageSelection{PageRange: &reducto.PageRange{Start: reducto.Ptr[int64](1), End: reducto.Ptr[int64](5)}},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Sync endpoints can still answer with a job id when the document is large.
	if out.AsyncParseResponse != nil {
		log.Fatalf("queued as job %s; use WaitForJob", out.AsyncParseResponse.JobID)
	}
	res := out.ParseResponse.Result
	if res.UrlResult != nil {
		fmt.Println("result too large, download from", res.UrlResult.URL)
		return
	}
	for _, chunk := range res.FullResult.Chunks {
		fmt.Println(chunk.Content)
	}
}
```

## Upload a local file

```go
up, err := client.UploadFile(ctx, "invoice.pdf")
if err != nil {
	log.Fatal(err)
}
out, err := client.Extract(ctx, &reducto.SyncExtractConfig{
	Input: up.Input(),
	Instructions: &reducto.Instructions{
		Schema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"total": map[string]any{"type": "number"},
			},
		},
	},
})
```

## Async jobs

```go
job, err := client.ParseAsync(ctx, &reducto.AsyncParseConfig{
	Input: reducto.DocumentInputFromString("https://example.com/big.pdf"),
	Webhook: &reducto.WebhookConfig{
		DirectWebhookConfig: &reducto.DirectWebhookConfig{URL: "https://my.app/hook"},
	},
})
if err != nil {
	log.Fatal(err)
}

final, err := client.WaitForJob(ctx, job.JobID, &reducto.WaitOptions{Interval: 3 * time.Second, Timeout: 10 * time.Minute})
var failed *reducto.JobFailedError
if errors.As(err, &failed) {
	log.Fatalf("job failed: %s", failed.Job.Error.Name)
}
var timedOut *reducto.JobTimeoutError
if errors.As(err, &timedOut) {
	log.Fatalf("still %s after %s", timedOut.Job.Status, timedOut.Timeout)
}
if final.Result.ParseResponse != nil {
	// ...
}
```

List every job without paging by hand:

```go
for job, err := range client.IterJobs(ctx, &reducto.ListJobsParams{Limit: reducto.Ptr[int64](200)}) {
	if err != nil {
		return err
	}
	fmt.Println(job.JobID, job.Status)
}
```

## Configuration

Every option works on the client and on a single call. Pass it to `New` for the
whole client, or as a trailing argument on any method:

```go
client := reducto.New("",
	reducto.WithBaseURL("https://vpc.example.com"),        // or $REDUCTO_BASE_URL
	reducto.WithMaxRetries(3),
	reducto.WithTimeout(2*time.Minute),                    // per attempt
	reducto.WithHeaders(map[string]string{"X-Org": "acme"}),
	reducto.WithLogger(slog.Default()),                    // or $REDUCTO_LOG=debug
)

out, err := client.Parse(ctx, req,
	reducto.WithMaxRetries(0),
	reducto.WithHeader("traceparent", tp),
)
```

| Option | Purpose |
| --- | --- |
| `WithBaseURL` | Host. `reducto.BaseURLEU` and `reducto.BaseURLAU` select a region. Also read from `REDUCTO_BASE_URL`. |
| `WithMaxRetries` | Retry count. Default 2. |
| `WithTimeout` | Per-attempt timeout, including the body read. Default 1h, as in the Python SDK. Zero disables. |
| `WithHeaders`, `WithHeader` | Extra headers. They override built-in headers, including `Authorization`. |
| `WithAppInfo` | Name and version of your application. Goes first in `User-Agent` and into `X-Reducto-Client`. |
| `WithClientInfo` | Send the `X-Reducto-*` client headers (default on). |
| `WithUserAgent` | Replace the whole `User-Agent` value. |
| `WithHTTPClient` | Custom transport (proxies, connection pools). |
| `WithLogger` | `*slog.Logger`. Requests and responses at debug, retries at info. |
| `WithMaxUploadSize` | Reject uploads over N bytes before sending. |
| `WithResponseInto` | Capture status, headers, request id and the raw body. |

`WithOptions` still returns a configured copy when you want to reuse one set of overrides.

## Client headers

Every request carries `User-Agent: Reducto/Go <version>` and the same platform headers as
the Python SDK:

```
X-Reducto-Lang: go                 X-Reducto-Runtime: go
X-Reducto-Package-Version: 0.1.0   X-Reducto-Runtime-Version: 1.23.1
X-Reducto-OS: MacOS                X-Reducto-Async: false
X-Reducto-Arch: arm64              X-Reducto-Retry-Count: 0
                                   X-Reducto-Read-Timeout: 300
```

They let Reducto tell client versions apart in its request logs. Nothing else is sent, and
no request is made that you did not ask for. `WithClientInfo(false)` drops the
`X-Reducto-*` headers. `WithAppInfo("my-app", "1.2")` puts your application in front of the
SDK in `User-Agent` and adds `X-Reducto-Client` and `X-Reducto-Client-Version`.

## Raw response

```go
var resp reducto.Response
out, err := client.Parse(ctx, req, reducto.WithResponseInto(&resp))
fmt.Println(resp.StatusCode, resp.RequestID, resp.Header.Get("Date"))
json.Unmarshal(resp.Body, &extra) // a field the types do not have yet
```

`resp` is filled on success and on `*APIError` alike. `APIError` also carries `Header`
and `RequestID`. The API does not send a request id header today; the field is populated
when a proxy in front of it does.

## Errors

| Type | When |
| --- | --- |
| `*APIError` | Any non-2xx response. `StatusCode`, `Code()`, `Message`, `Validation`, `Header`, `RequestID`, `RetryAfter()`. |
| `*APIConnectionError` | No response: DNS, refused, reset, malformed reply. `Unwrap` gives the transport error. |
| `*APITimeoutError` | One attempt exceeded `WithTimeout`. Matches `errors.Is(err, context.DeadlineExceeded)`. |
| `*JobFailedError` | `WaitForJob` saw `Failed`. |
| `*JobTimeoutError` | `WaitForJob` hit `WaitOptions.Timeout`. Also matches `context.DeadlineExceeded`. |
| `*WebhookVerificationError` | `VerifyWebhook` rejected a delivery. |
| `*TypedExtractError` | `ExtractAs` or `ValidateExtract` could not decode the result. `Response` holds the raw response. |

A deadline or cancellation on your own `ctx` is returned as the plain `ctx.Err()`.

```go
var apiErr *reducto.APIError
if errors.As(err, &apiErr) {
	fmt.Println(apiErr.StatusCode, apiErr.Code(), apiErr.Message, apiErr.RequestID)
	for _, v := range apiErr.Validation { // 422 details
		fmt.Println(v.Msg)
	}
}
```

## Typed extraction

`ExtractAs[T]` sends the JSON schema you pass and decodes the items into `T`. Go has no
runtime schema derivation, so the schema is written by hand or produced by a library of your
choice (for example `invopop/jsonschema`).

```go
type Invoice struct {
	Total  float64 `json:"total"`
	Vendor string  `json:"vendor"`
}

schema := map[string]any{
	"type": "object",
	"properties": map[string]any{
		"total":  map[string]any{"type": "number"},
		"vendor": map[string]any{"type": "string"},
	},
	"required": []string{"total", "vendor"},
}

out, err := reducto.ExtractAs[Invoice](ctx, client, schema, &reducto.SyncExtractConfig{
	Input: reducto.DocumentInputFromString("https://example.com/invoice.pdf"),
})
out.Result            // []Invoice
out.Usage()           // ExtractUsage
out.V3ExtractResponse // the response as received
```

It returns a `*TypedExtractError` when the document was queued as a job, when the result came
back as a URL (`ForceURLResult`), or when an item does not decode. For a queued job, wait for
it and decode the result yourself:

```go
job, err := client.WaitForJob(ctx, jobID, nil)
out, err := reducto.ValidateExtract[Invoice](job.Result)
```

## Retries

Default 2 retries with exponential backoff from 500ms, capped at 8s plus jitter.
`Retry-After-Ms`, then `Retry-After` (seconds or HTTP-date), is honoured when it is 60s or
less; larger values fall back to the normal backoff. An `x-should-retry: true|false` response
header overrides the status-code rule below.

| Failure | GET / DELETE | POST |
| --- | --- | --- |
| 408, 429, 5xx | retry | retry |
| Connection refused, DNS failure | retry | retry |
| Timeout, reset, EOF after sending | retry | retry |

POST connection errors are retried to match the published SDK. A retry can create a
second job if the server accepted the first request. Set `WithMaxRetries(0)` to
disable retries. Pass an `Idempotency-Key` with `WithHeader` if a proxy needs it.

## Webhooks

Configure Svix mode (`WebhookConfigNew{Mode: reducto.WebhookConfigNewModeSvix}`), then verify
each delivery with the endpoint secret from the portal:

```go
func handler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	if err := reducto.VerifyWebhook(os.Getenv("SVIX_SECRET"), body, r.Header); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	// body is trusted
}
```

Both `svix-*` and `webhook-*` header names are accepted. Deliveries in `direct` mode are
unsigned and cannot be verified.

## Uploads

`Upload` streams an `io.ReadSeeker` (such as `*os.File`) with a known `Content-Length` and
rewinds it on retry. Any other reader is buffered in memory. Empty files fail with
`ErrEmptyUpload`; `WithMaxUploadSize` adds an upper bound (`ErrUploadTooLarge`).

`PresignUpload` returns a handle and a presigned URL without sending the bytes. PUT the file to
the URL yourself, then use `up.Input()` as usual.

## Conventions

- Optional scalar fields are pointers (`*bool`, `*int64`, `*string`). Use
  `reducto.Ptr(v)`. This lets you send an explicit `false` or `0` and
  overrides the server default. A nil field is omitted and the server default applies.
- `anyOf`/`oneOf` schemas become union structs with one pointer per variant
  (`DocumentInput`, `ParseResult`, `JobResult`, ...). Exactly one field is
  set. Build them with the `<Union>From<Variant>` helpers. When the server sends a
  shape no variant matches (a new `type`, say), decoding does not fail: the raw JSON
  lands in `Unknown` and round-trips on marshal.
- Fields with a fixed value in the spec (`response_type`, `scope`, `mode`)
  are set for you on marshal.
- Enums are typed strings with constants: `reducto.ChunkingChunkModePage`.
- `client.Do(ctx, method, path, query, body)` is the raw escape hatch. It returns `json.RawMessage`.

The types follow Reducto's OpenAPI document. `go run ./internal/specdrift` checks them against
the snapshot in `spec/openapi.json` on every pull request; intentional differences are listed
with a reason in `spec/drift-allowlist.json`. See [CONTRIBUTING.md](CONTRIBUTING.md).

## Development

```sh
go test ./...
```

`e2e_test.go` runs every endpoint against the real API. It is behind the `e2e` build tag and
needs `REDUCTO_API_KEY`:

```sh
REDUCTO_API_KEY=... go test -tags e2e -run TestE2E -v -timeout 20m ./...
```
