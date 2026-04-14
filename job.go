package main

import (
	"fmt"
	"time"
)

// JobStatus represents the current state of a job
type JobStatus string

const (
	StatusPending   JobStatus = "pending"
	StatusRunning   JobStatus = "running"
	StatusCompleted JobStatus = "completed"
	StatusFailed    JobStatus = "failed"
)

// Job represents a unit of work
type Job struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Payload   string    `json:"payload"`
	Status    JobStatus `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	StartedAt time.Time `json:"started_at"`
	CompletedAt time.Time `json:"completed_at"`
	Error     string    `json:"error,omitempty"`
}

// Validate checks if the job has valid required fields
func (j *Job) Validate() error {
	if j.ID == "" {
		return fmt.Errorf("job ID cannot be empty")
	}
	if j.Name == "" {
		return fmt.Errorf("job name cannot be empty")
	}
	return nil
}

// Execute simulates job execution (for now, just a sleep)
// In a real system, this would do actual work
func (j *Job) Execute() error {
	j.Status = StatusRunning
	j.StartedAt = time.Now()

	// Simulate work with a 2-second delay
	time.Sleep(2 * time.Second)

	j.Status = StatusCompleted
	j.CompletedAt = time.Now()
	return nil
}
