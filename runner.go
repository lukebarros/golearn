package main

import (
	"fmt"
	"sync"
)

// JobRunner manages job storage and retrieval
type JobRunner struct {
	jobs map[string]*Job
	mu   sync.Mutex // Protect access to the jobs map
}

// NewJobRunner creates a new JobRunner instance
func NewJobRunner() *JobRunner {
	return &JobRunner{
		jobs: make(map[string]*Job),
	}
}

// Submit adds a job to the runner
// For now (Day 1), jobs are stored but not executed
func (jr *JobRunner) Submit(job *Job) error {
	// Validate the job before accepting it
	if err := job.Validate(); err != nil {
		return fmt.Errorf("invalid job: %w", err)
	}

	jr.mu.Lock()
	defer jr.mu.Unlock()

	// Check if job ID already exists
	if _, exists := jr.jobs[job.ID]; exists {
		return fmt.Errorf("job with ID %s already exists", job.ID)
	}

	// Set initial status
	job.Status = StatusPending
	jr.jobs[job.ID] = job

	return nil
}

// GetStatus retrieves the current status of a job by ID
func (jr *JobRunner) GetStatus(jobID string) (*Job, error) {
	jr.mu.Lock()
	defer jr.mu.Unlock()

	job, exists := jr.jobs[jobID]
	if !exists {
		return nil, fmt.Errorf("job %s not found", jobID)
	}

	return job, nil
}

// ListJobs returns all jobs currently in the runner
func (jr *JobRunner) ListJobs() []*Job {
	jr.mu.Lock()
	defer jr.mu.Unlock()

	jobs := make([]*Job, 0, len(jr.jobs))
	for _, job := range jr.jobs {
		jobs = append(jobs, job)
	}
	return jobs
}
