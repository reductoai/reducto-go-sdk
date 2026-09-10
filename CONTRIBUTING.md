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
| `upload.go` | `Upload`, `UploadFile`, `PresignUpload` | Yes |
| `jobs.go` | `WaitForJob`, `IterJobs` | Yes |
| `extract.go` | `ExtractAs`, `ValidateExtract` | Yes |
| `webhook.go` | `VerifyWebhook` | Yes |
| `raw.go` | `Do`, the raw escape hatch | Yes |
| `union.go`, `time.go` | JSON helpers | Yes |

`types.go` holds the request and response types; `api.go` holds one method per endpoint. Both
follow Reducto's OpenAPI document by hand. Keep the conventions you see: `json` tags with
`omitempty` on optional fields, pointers for optional scalars, typed strings with constants
for enums, and union structs with one pointer per variant plus `Unknown`.

## Spec drift

`spec/openapi.json` is a snapshot of the public OpenAPI document. `go run ./internal/specdrift`
anchors each `Client` method to an endpoint through its `c.do` call, then compares the Go types
against the spec: field names, types, enum values and required-ness. CI runs it on every pull
request and fails on drift.

```sh
go run ./internal/specdrift                    # check against the snapshot
go run ./internal/specdrift -live              # check against https://reducto.ai/openapi.json
go run ./internal/specdrift -update-snapshot   # refresh the snapshot, then check
```

Refreshing the snapshot is a manual step. Commit it together with the SDK change it calls for.

When the SDK must differ from the spec on purpose, add an entry to `spec/drift-allowlist.json`
with the `endpoint`, `location` and `kind` the tool printed, an optional `detail` to pin one
item, and a `reason` that says when to remove it. The tool reports allowed items and flags
entries that no longer match anything.

## Tests

`go test ./...` runs offline against an in-process HTTP server.

End-to-end tests in `e2e_test.go` call the live Reducto API. They are behind the `e2e` build
tag, so `go test ./...` does not compile them. They need `REDUCTO_API_KEY`. The top-level
tests run in parallel, one per endpoint group:

```sh
REDUCTO_API_KEY=... go test -tags e2e -run TestE2E -v -timeout 20m ./...
```

## Style

Run `gofmt` before you commit. CI fails on unformatted files. Prefer no comments; when a
comment is needed, say why, not what.

## Releases

A release is a git tag on `main`, for example `v0.2.0`. Update `Version` in `client.go` and
add a section to `CHANGELOG.md` in the same commit. Never move or delete a tag: Go modules
and the module proxy pin to it.
