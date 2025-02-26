# Shared Params Types

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#AdvancedProcessingOptionsParam">AdvancedProcessingOptionsParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#ArrayExtractConfigParam">ArrayExtractConfigParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#BaseProcessingOptionsParam">BaseProcessingOptionsParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#ExperimentalProcessingOptionsParam">ExperimentalProcessingOptionsParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#PageRangeParam">PageRangeParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#SplitCategoryParam">SplitCategoryParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#UploadParam">UploadParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#WebhookConfigNewParam">WebhookConfigNewParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#BoundingBox">BoundingBox</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#ExtractResponse">ExtractResponse</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#ParseResponse">ParseResponse</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#ParseUsage">ParseUsage</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#SplitResponse">SplitResponse</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#Upload">Upload</a>

# reducto

Response Types:

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#APIVersionResponse">APIVersionResponse</a>

Methods:

- <code title="get /version">client.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ReductoService.APIVersion">APIVersion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#APIVersionResponse">APIVersionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /upload">client.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ReductoService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#UploadParams">UploadParams</a>) (<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#Upload">Upload</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Job

Response Types:

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#JobCancelResponse">JobCancelResponse</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#JobGetResponse">JobGetResponse</a>

Methods:

- <code title="post /cancel/{job_id}">client.Job.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#JobService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#JobCancelResponse">JobCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /job/{job_id}">client.Job.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#JobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#JobGetResponse">JobGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Split

Response Types:

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#SplitRunJobResponse">SplitRunJobResponse</a>

Methods:

- <code title="post /split">client.Split.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#SplitService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#SplitRunParams">SplitRunParams</a>) (<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#SplitResponse">SplitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /split_async">client.Split.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#SplitService.RunJob">RunJob</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#SplitRunJobParams">SplitRunJobParams</a>) (<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#SplitRunJobResponse">SplitRunJobResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Parse

Response Types:

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ParseRunJobResponse">ParseRunJobResponse</a>

Methods:

- <code title="post /parse">client.Parse.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ParseService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ParseRunParams">ParseRunParams</a>) (<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#ParseResponse">ParseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /parse_async">client.Parse.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ParseService.RunJob">RunJob</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ParseRunJobParams">ParseRunJobParams</a>) (<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ParseRunJobResponse">ParseRunJobResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Extract

Response Types:

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ExtractRunJobResponse">ExtractRunJobResponse</a>

Methods:

- <code title="post /extract">client.Extract.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ExtractService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ExtractRunParams">ExtractRunParams</a>) (<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#ExtractResponse">ExtractResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /extract_async">client.Extract.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ExtractService.RunJob">RunJob</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ExtractRunJobParams">ExtractRunJobParams</a>) (<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ExtractRunJobResponse">ExtractRunJobResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhook

Methods:

- <code title="post /configure_webhook">client.Webhook.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#WebhookService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Config

Params Types:

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ExtractConfigParam">ExtractConfigParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ParseConfigParam">ParseConfigParam</a>
