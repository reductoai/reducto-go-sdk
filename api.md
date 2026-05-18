# Shared Params Types

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#ChunkingParam">ChunkingParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#DirectWebhookConfigParam">DirectWebhookConfigParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#FigureAgenticParam">FigureAgenticParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#PageRangeParam">PageRangeParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#SplitLargeTablesParam">SplitLargeTablesParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#SvixWebhookConfigParam">SvixWebhookConfigParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#TableAgenticParam">TableAgenticParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#TextAgenticParam">TextAgenticParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#UploadParam">UploadParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#WebhookConfigNewParam">WebhookConfigNewParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#AsyncEditResponse">AsyncEditResponse</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#AsyncExtractResponse">AsyncExtractResponse</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#AsyncParseResponse">AsyncParseResponse</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#AsyncPipelineResponse">AsyncPipelineResponse</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#AsyncSplitResponse">AsyncSplitResponse</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#ClassifyResponse">ClassifyResponse</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#EditResponse">EditResponse</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#ExtractResponse">ExtractResponse</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#ParseResponse">ParseResponse</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#PipelineResponse">PipelineResponse</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#SplitResponse">SplitResponse</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#Upload">Upload</a>

# reducto

Methods:

- <code title="get /version">client.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ReductoService.APIVersion">APIVersion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /upload">client.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ReductoService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#UploadParams">UploadParams</a>) (\*<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#Upload">Upload</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Parse

Params Types:

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#AsyncConfigV3Param">AsyncConfigV3Param</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#AsyncParseConfigParam">AsyncParseConfigParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#EnhanceParam">EnhanceParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#FormattingParam">FormattingParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#RetrievalParam">RetrievalParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#SettingsParam">SettingsParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#SpreadsheetParam">SpreadsheetParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ParseRunResponse">ParseRunResponse</a>

Methods:

- <code title="post /parse">client.Parse.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ParseService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ParseRunParams">ParseRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ParseRunResponse">ParseRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /parse_async">client.Parse.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ParseService.RunJob">RunJob</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ParseRunJobParams">ParseRunJobParams</a>) (\*<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#AsyncParseResponse">AsyncParseResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Extract

Params Types:

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#AsyncExtractConfigParam">AsyncExtractConfigParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ExtractSettingsParam">ExtractSettingsParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#InstructionsParam">InstructionsParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ParseOptionsParam">ParseOptionsParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#V3Extract">V3Extract</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ExtractRunResponseUnion">ExtractRunResponseUnion</a>

Methods:

- <code title="post /extract">client.Extract.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ExtractService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ExtractRunParams">ExtractRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ExtractRunResponseUnion">ExtractRunResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /extract_async">client.Extract.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ExtractService.RunJob">RunJob</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ExtractRunJobParams">ExtractRunJobParams</a>) (\*<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#AsyncExtractResponse">AsyncExtractResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Split

Params Types:

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#SplitCategoryParam">SplitCategoryParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#DeepSplitPageEvidence">DeepSplitPageEvidence</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ParseUsage">ParseUsage</a>

Methods:

- <code title="post /split">client.Split.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#SplitService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#SplitRunParams">SplitRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#SplitResponse">SplitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /split_async">client.Split.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#SplitService.RunJob">RunJob</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#SplitRunJobParams">SplitRunJobParams</a>) (\*<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#AsyncSplitResponse">AsyncSplitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Edit

Params Types:

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#BoundingBoxParam">BoundingBoxParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#EditOptionsParam">EditOptionsParam</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#EditWidgetParam">EditWidgetParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#BoundingBox">BoundingBox</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#EditWidget">EditWidget</a>

Methods:

- <code title="post /edit">client.Edit.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#EditService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#EditRunParams">EditRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#EditResponse">EditResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /edit_async">client.Edit.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#EditService.RunJob">RunJob</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#EditRunJobParams">EditRunJobParams</a>) (\*<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#AsyncEditResponse">AsyncEditResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Pipeline

Params Types:

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#PipelineSettingsParam">PipelineSettingsParam</a>

Methods:

- <code title="post /pipeline">client.Pipeline.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#PipelineService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#PipelineRunParams">PipelineRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#PipelineResponse">PipelineResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /pipeline_async">client.Pipeline.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#PipelineService.RunJob">RunJob</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#PipelineRunJobParams">PipelineRunJobParams</a>) (\*<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#AsyncPipelineResponse">AsyncPipelineResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Classify

Methods:

- <code title="post /classify">client.Classify.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ClassifyService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#ClassifyRunParams">ClassifyRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared">shared</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk/shared#ClassifyResponse">ClassifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhook

Methods:

- <code title="post /configure_webhook">client.Webhook.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#WebhookService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Job

Response Types:

- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#JobCancelResponse">JobCancelResponse</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#JobGetResponse">JobGetResponse</a>
- <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#JobGetAllResponse">JobGetAllResponse</a>

Methods:

- <code title="post /cancel/{job_id}">client.Job.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#JobService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#JobCancelResponse">JobCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /job/{job_id}">client.Job.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#JobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#JobGetResponse">JobGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs">client.Job.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#JobService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#JobGetAllParams">JobGetAllParams</a>) (\*<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk">reducto</a>.<a href="https://pkg.go.dev/github.com/reductoai/reducto-go-sdk#JobGetAllResponse">JobGetAllResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
