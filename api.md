# Shared Params Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared#AdvancedProcessingOptionsParam">AdvancedProcessingOptionsParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared#ArrayExtractConfigParam">ArrayExtractConfigParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared#BaseProcessingOptionsParam">BaseProcessingOptionsParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared#ExperimentalProcessingOptionsParam">ExperimentalProcessingOptionsParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared#PageRangeParam">PageRangeParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared#SplitCategoryParam">SplitCategoryParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared#WebhookConfigNewParam">WebhookConfigNewParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared#BoundingBox">BoundingBox</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared#ExtractResponse">ExtractResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared#ParseResponse">ParseResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared#ParseUsage">ParseUsage</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared#SplitResponse">SplitResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared#Upload">Upload</a>

# reductoai

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#APIVersionResponse">APIVersionResponse</a>

Methods:

- <code title="get /version">client.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#ReductoaiService.APIVersion">APIVersion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#APIVersionResponse">APIVersionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /upload">client.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#ReductoaiService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#UploadParams">UploadParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared#Upload">Upload</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Job

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#JobCancelResponse">JobCancelResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#JobGetResponse">JobGetResponse</a>

Methods:

- <code title="post /cancel/{job_id}">client.Job.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#JobService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#JobCancelResponse">JobCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /job/{job_id}">client.Job.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#JobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#JobGetResponse">JobGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Split

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#SplitRunJobResponse">SplitRunJobResponse</a>

Methods:

- <code title="post /split">client.Split.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#SplitService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#SplitRunParams">SplitRunParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared#SplitResponse">SplitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /split_async">client.Split.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#SplitService.RunJob">RunJob</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#SplitRunJobParams">SplitRunJobParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#SplitRunJobResponse">SplitRunJobResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Parse

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#ParseRunJobResponse">ParseRunJobResponse</a>

Methods:

- <code title="post /parse">client.Parse.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#ParseService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#ParseRunParams">ParseRunParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared#ParseResponse">ParseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /parse_async">client.Parse.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#ParseService.RunJob">RunJob</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#ParseRunJobParams">ParseRunJobParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#ParseRunJobResponse">ParseRunJobResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Extract

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#ExtractRunJobResponse">ExtractRunJobResponse</a>

Methods:

- <code title="post /extract">client.Extract.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#ExtractService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#ExtractRunParams">ExtractRunParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go/shared#ExtractResponse">ExtractResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /extract_async">client.Extract.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#ExtractService.RunJob">RunJob</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#ExtractRunJobParams">ExtractRunJobParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go">reductoai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#ExtractRunJobResponse">ExtractRunJobResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhook

Methods:

- <code title="post /configure_webhook">client.Webhook.<a href="https://pkg.go.dev/github.com/stainless-sdks/reductoai-go#WebhookService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
