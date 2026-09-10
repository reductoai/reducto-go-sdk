package reducto

import (
	"context"
	"fmt"
	"iter"
	"time"
)

// WaitOptions tune Client.WaitForJob.
type WaitOptions struct {
	// Interval between polls. Default 2s.
	Interval time.Duration
	// Timeout bounds the total wait. Zero means wait until ctx is done.
	Timeout time.Duration
}

// JobFailedError is returned by WaitForJob when the job ends in a Failed state.
type JobFailedError struct {
	JobID string
	Job   *EnhancedAsyncJobResponse
}

func (e *JobFailedError) Error() string {
	msg := "job failed"
	if e.Job.Error != nil {
		msg = fmt.Sprintf("%s: %s", e.Job.Error.Name, e.Job.Error.Message)
	} else if e.Job.Reason != nil {
		msg = *e.Job.Reason
	}
	return fmt.Sprintf("reducto: job %s: %s", e.JobID, msg)
}

// JobTimeoutError is returned by WaitForJob when WaitOptions.Timeout elapses first.
// Job holds the last state seen, if any. It matches errors.Is(err, context.DeadlineExceeded).
type JobTimeoutError struct {
	JobID   string
	Timeout time.Duration
	Job     *EnhancedAsyncJobResponse
}

func (e *JobTimeoutError) Error() string {
	return fmt.Sprintf("reducto: job %s did not finish within %s", e.JobID, e.Timeout)
}

func (e *JobTimeoutError) Is(target error) bool { return target == context.DeadlineExceeded }

// WaitForJob polls GetJob until the job is Completed or Failed and returns the final state.
// opts apply to each GetJob call.
func (c *Client) WaitForJob(ctx context.Context, jobID string, wait *WaitOptions, opts ...Option) (*EnhancedAsyncJobResponse, error) {
	interval := 2 * time.Second
	var timeout time.Duration
	if wait != nil {
		if wait.Interval > 0 {
			interval = wait.Interval
		}
		timeout = wait.Timeout
	}
	wctx := ctx
	if timeout > 0 {
		var cancel context.CancelFunc
		wctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	var last *EnhancedAsyncJobResponse
	for {
		job, err := c.GetJob(wctx, jobID, opts...)
		if err != nil {
			if wctx.Err() != nil && ctx.Err() == nil {
				return last, &JobTimeoutError{JobID: jobID, Timeout: timeout, Job: last}
			}
			return last, err
		}
		last = job
		switch job.Status {
		case JobStatusCompleted:
			return job, nil
		case JobStatusFailed:
			return job, &JobFailedError{JobID: jobID, Job: job}
		}
		if err := sleep(wctx, interval); err != nil {
			if ctx.Err() == nil {
				return job, &JobTimeoutError{JobID: jobID, Timeout: timeout, Job: job}
			}
			return job, err
		}
	}
}

// IterJobs pages through ListJobs and yields every job. params.Limit sets the page size;
// the cursor is managed for you. Iteration stops at the first error, which is yielded last.
//
//	for job, err := range c.IterJobs(ctx, nil) {
//		if err != nil { return err }
//		fmt.Println(job.JobID)
//	}
func (c *Client) IterJobs(ctx context.Context, params *ListJobsParams, opts ...Option) iter.Seq2[*SingleJob, error] {
	return func(yield func(*SingleJob, error) bool) {
		var p ListJobsParams
		if params != nil {
			p = *params
		}
		for {
			page, err := c.ListJobs(ctx, &p, opts...)
			if err != nil {
				yield(nil, err)
				return
			}
			for i := range page.Jobs {
				if !yield(&page.Jobs[i], nil) {
					return
				}
			}
			next := page.NextCursor
			if next == nil || *next == "" || (p.Cursor != nil && *p.Cursor == *next) {
				return
			}
			p.Cursor = next
		}
	}
}
