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

// ExecuteJob runs a job synchronously and updates its status
// This is a pointer receiver because we mutate the job's state
func (jr *JobRunner) ExecuteJob(jobID string) error {
	jr.mu.Lock()
	job, exists := jr.jobs[jobID]
	jr.mu.Unlock()

	if !exists {
		return fmt.Errorf("job %s not found", jobID)
	}

	// Execute the job
	if err := job.Execute(); err != nil {
		jr.mu.Lock()
		job.Status = StatusFailed
		job.Error = err.Error()
		jr.mu.Unlock()
		return fmt.Errorf("job execution failed: %w", err)
	}

	// Job executed successfully (Execute() already updated status)
	return nil
}

// UpdateJobStatus updates a job's status
// Note: This is a mutation operation—why do we need a pointer receiver?
// Because we're modifying the JobRunner's state (the jobs map and its contents)
func (jr *JobRunner) UpdateJobStatus(jobID string, status JobStatus) error {
	jr.mu.Lock()
	defer jr.mu.Unlock()

	job, exists := jr.jobs[jobID]
	if !exists {
		return fmt.Errorf("job %s not found", jobID)
	}

	job.Status = status
	return nil
}

// DeleteJob removes a job from the runner
// This clearly shows why pointer receivers are necessary:
// We're modifying the JobRunner's internal state (the jobs map)
func (jr *JobRunner) DeleteJob(jobID string) error {
	jr.mu.Lock()
	defer jr.mu.Unlock()

	if _, exists := jr.jobs[jobID]; !exists {
		return fmt.Errorf("job %s not found", jobID)
	}

	delete(jr.jobs, jobID)
	return nil
}

// GetStats returns statistics about the jobs
// This is a read-only method (no mutations), so it COULD use a value receiver,
// but we use pointer receiver for consistency with other methods
func (jr *JobRunner) GetStats() JobStats {
	jr.mu.Lock()
	defer jr.mu.Unlock()

	stats := JobStats{
		TotalJobs:     len(jr.jobs),
		PendingJobs:   0,
		RunningJobs:   0,
		CompletedJobs: 0,
		FailedJobs:    0,
	}

	for _, job := range jr.jobs {
		switch job.Status {
		case StatusPending:
			stats.PendingJobs++
		case StatusRunning:
			stats.RunningJobs++
		case StatusCompleted:
			stats.CompletedJobs++
		case StatusFailed:
			stats.FailedJobs++
		}
	}

	return stats
}
