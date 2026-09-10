# Contributing

## Setup

Install Go 1.23 or later. Then:

```sh
go test ./...
```

There are no other dependencies. The module uses only the standard library.

## Layout

| File | Contents | Edit by hand? |
| --- | --- | --- |
| `types.go` | Request and response types, enums, unions | No |
| `api.go` | One method per API endpoint | No |
| `client.go` | Client, options, transport, retries | Yes |
| `errors.go` | `APIError` and friends | Yes |
| `upload.go` | `Upload`, `UploadFile` | Yes |
| `jobs.go` | `WaitForJob`, `IterJobs` | Yes |
| `extract.go` | `ExtractAs`, `ValidateExtract` | Yes |
| `webhook.go` | `VerifyWebhook` | Yes |
| `raw.go` | `Do`, the raw escape hatch | Yes |
| `union.go`, `time.go` | JSON helpers | Yes |

`types.go` and `api.go` are generated from Reducto's OpenAPI document. A hand edit is lost on
the next regeneration. If a type or an endpoint is wrong or missing, open an issue.

## Tests

`go test ./...` runs offline against an in-process HTTP server.

`smoke_test.go` calls the real API. It is skipped unless you set `REDUCTO_SMOKE=1`:

```sh
REDUCTO_SMOKE=1 REDUCTO_API_KEY=... go test -run TestSmoke -v -timeout 20m
```

## Style

Run `gofmt` before you commit. CI fails on unformatted files. Prefer no comments; when a
comment is needed, say why, not what.

## Releases

A release is a git tag on `main`, for example `v0.2.0`. Update `Version` in `client.go` and
add a section to `CHANGELOG.md` in the same commit. Never move or delete a tag: Go modules
and the module proxy pin to it.
