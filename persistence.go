package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Persistence functions will be implemented in Days 3-4
// These stubs show the interface we'll implement

// SaveJobsToFile saves all jobs to a JSON file
// TODO (Day 3): Implement JSON marshaling and file writing
func SaveJobsToFile(runner *JobRunner, filename string) error {
	jobs := runner.ListJobs()
	data, err := json.MarshalIndent(jobs, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal jobs: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// LoadJobsFromFile loads jobs from a JSON file
// TODO (Day 3): Implement file reading and JSON unmarshaling
func LoadJobsFromFile(runner *JobRunner, filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist yet—this is OK on first run
			return nil
		}
		return fmt.Errorf("failed to read file: %w", err)
	}

	var jobs []*Job
	err = json.Unmarshal(data, &jobs)
	if err != nil {
		return fmt.Errorf("failed to unmarshal jobs: %w", err)
	}

	// Add all loaded jobs to the runner
	for _, job := range jobs {
		if err := runner.Submit(job); err != nil {
			// Note: We might want to handle conflicts (job already exists)
			// This will be discussed in Day 3-4
			return fmt.Errorf("failed to load job %s: %w", job.ID, err)
		}
	}

	return nil
}

// Note on error handling and JSON:
// - json.Marshal/Unmarshal return error values (not exceptions)
// - os.ReadFile and os.WriteFile also return error values
// - Notice how we "wrap" errors with %w for context
// - This allows the caller to see the full error chain
//
// This is the Go way: explicit error propagation, not try/catch
