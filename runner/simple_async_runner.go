package runner

import (
	"context"
	"errors"
	"fmt"
	"github.com/shindakioku/ciscor/actions"
	"log"
	"time"
)

// BeforeJobRunningLogger simple logger. Just use it if you need some logging info.
func BeforeJobRunningLogger(job AsyncJob) AsyncJob {
	log.Println(fmt.Sprintf(
		"Job [%s] will be execute.\nArgs: %v\nTime: %s",
		job.Action.Identification(),
		job.Args,
		time.Now().UTC().Format("2006-01-02 15-01-05"),
	))

	return job
}

// AfterJobRunningLogger simple logger. Just use it if you need some logging info.
func AfterJobRunningLogger(job AsyncJob, jobResult AsyncJobResult) {
	jobStatus := "Successfully"
	if jobResult.Error != nil {
		jobStatus = "Failure"
	}

	log.Println(fmt.Sprintf(
		"Job [%s] executed.Status: %s\n\nArgs: %v\nTime: %s\nJob output: %v",
		job.Action.Identification(),
		jobStatus,
		job.Args,
		time.Now().UTC().Format("2006-01-02 15-01-05"),
		jobResult.ReturnValue,
	))
}

type AsyncJob struct {
	Action actions.Action
	Args   any
}

type AsyncJobResult struct {
	ReturnValue any
	Error       error
}

type SimpleASyncRunner struct {
	jobs             chan AsyncJob
	beforeJobRunning func(job AsyncJob) AsyncJob
	afterJobRunning  func(job AsyncJob, jobResult AsyncJobResult)
}

func (s *SimpleASyncRunner) Add(job AsyncJob) {
	s.jobs <- job
}

func (s *SimpleASyncRunner) Run(ctx context.Context) error {
	for {
		select {
		case job, ok := <-s.jobs:
			if !ok {
				return errors.New("jobs channel closed")
			}

			if s.beforeJobRunning != nil {
				job = s.beforeJobRunning(job)
			}

			result, err := job.Action.Handle(job.Args)
			if s.afterJobRunning != nil {
				s.afterJobRunning(job, AsyncJobResult{
					ReturnValue: result,
					Error:       err,
				})
			}
		case <-ctx.Done():
			close(s.jobs)

			return nil
		}
	}
}

func (s *SimpleASyncRunner) SetBeforeJobRunning(f func(job AsyncJob) AsyncJob) {
	s.beforeJobRunning = f
}

func (s *SimpleASyncRunner) SetAfterJobRunning(f func(job AsyncJob, jobResult AsyncJobResult)) {
	s.afterJobRunning = f
}

func NewSimpleASyncRunner(chanBuffer int) ASyncRunner {
	return &SimpleASyncRunner{
		jobs: make(chan AsyncJob, chanBuffer),
	}
}
