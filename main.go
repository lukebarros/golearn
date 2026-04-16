package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const jobsFile = "jobs.json"

func main() {
	// Create a new job runner
	runner := NewJobRunner()

	// Load existing jobs from file
	if err := LoadJobsFromFile(runner, jobsFile); err != nil {
		log.Fatalf("Failed to load jobs from file: %v", err)
	}

	// Check for interactive mode
	if len(os.Args) > 1 && os.Args[1] == "interactive" {
		runInteractiveMode(runner)
		return
	}

	// Define subcommands
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "submit":
		handleSubmit(runner, os.Args[2:])
	case "status":
		handleStatus(runner, os.Args[2:])
	case "list":
		handleList(runner)
	case "execute":
		handleExecute(runner, os.Args[2:])
	case "delete":
		handleDelete(runner, os.Args[2:])
	case "stats":
		handleStats(runner)
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

// handleSubmit processes the "submit" command
func handleSubmit(runner *JobRunner, args []string) {
	fs := flag.NewFlagSet("submit", flag.ExitOnError)
	id := fs.String("id", "", "Job ID (required)")
	name := fs.String("name", "", "Job name (required)")
	payload := fs.String("payload", "", "Job payload (optional)")

	fs.Parse(args)

	if *id == "" || *name == "" {
		fmt.Println("Error: id and name are required")
		fmt.Println("Usage: submit -id <id> -name <name> [-payload <payload>]")
		os.Exit(1)
	}

	job := &Job{
		ID:        *id,
		Name:      *name,
		Payload:   *payload,
		CreatedAt: time.Now(),
	}

	if err := runner.Submit(job); err != nil {
		log.Fatalf("Failed to submit job: %v", err)
	}

	// Save jobs to file after successful submission
	if err := SaveJobsToFile(runner, jobsFile); err != nil {
		log.Fatalf("Failed to save jobs to file: %v", err)
	}

	fmt.Printf("✓ Job submitted: ID=%s, Name=%s\n", job.ID, job.Name)
}

// handleStatus processes the "status" command
func handleStatus(runner *JobRunner, args []string) {
	if len(args) == 0 {
		fmt.Println("Error: job ID required")
		fmt.Println("Usage: status <job-id>")
		os.Exit(1)
	}

	jobID := args[0]
	job, err := runner.GetStatus(jobID)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	printJobDetails(job)
}

// handleList processes the "list" command
func handleList(runner *JobRunner) {
	jobs := runner.ListJobs()
	if len(jobs) == 0 {
		fmt.Println("No jobs found.")
		return
	}

	fmt.Println("\n=== Job List ===")
	fmt.Printf("%-20s | %-20s | %-15s | %-10s\n", "ID", "Name", "Status", "Created")
	fmt.Println(strings.Repeat("-", 75))

	for _, job := range jobs {
		createdStr := job.CreatedAt.Format("2006-01-02 15:04")
		fmt.Printf("%-20s | %-20s | %-15s | %-10s\n", job.ID, job.Name, job.Status, createdStr)
	}
}

// printJobDetails prints detailed information about a job
func printJobDetails(job *Job) {
	fmt.Println("\n=== Job Details ===")
	fmt.Printf("ID:            %s\n", job.ID)
	fmt.Printf("Name:          %s\n", job.Name)
	fmt.Printf("Status:        %s\n", job.Status)
	fmt.Printf("Created:       %s\n", job.CreatedAt.Format("2006-01-02 15:04:05"))
	if !job.StartedAt.IsZero() {
		fmt.Printf("Started:       %s\n", job.StartedAt.Format("2006-01-02 15:04:05"))
	}
	if !job.CompletedAt.IsZero() {
		fmt.Printf("Completed:     %s\n", job.CompletedAt.Format("2006-01-02 15:04:05"))
	}
	if job.Error != "" {
		fmt.Printf("Error:         %s\n", job.Error)
	}
	if job.Payload != "" {
		fmt.Printf("Payload:       %s\n", job.Payload)
	}
	fmt.Println()
}

// handleExecute processes the "execute" command
func handleExecute(runner *JobRunner, args []string) {
	if len(args) == 0 {
		fmt.Println("Error: job ID required")
		fmt.Println("Usage: execute <job-id>")
		os.Exit(1)
	}

	jobID := args[0]
	if err := runner.ExecuteJob(jobID); err != nil {
		log.Fatalf("Failed to execute job: %v", err)
	}

	// Save jobs to file after execution
	if err := SaveJobsToFile(runner, jobsFile); err != nil {
		log.Fatalf("Failed to save jobs to file: %v", err)
	}

	job, _ := runner.GetStatus(jobID)
	fmt.Printf("✓ Job executed: %s (Status: %s)\n", jobID, job.Status)
}

// handleDelete processes the "delete" command
func handleDelete(runner *JobRunner, args []string) {
	if len(args) == 0 {
		fmt.Println("Error: job ID required")
		fmt.Println("Usage: delete <job-id>")
		os.Exit(1)
	}

	jobID := args[0]
	if err := runner.DeleteJob(jobID); err != nil {
		log.Fatalf("Failed to delete job: %v", err)
	}

	// Save jobs to file after deletion
	if err := SaveJobsToFile(runner, jobsFile); err != nil {
		log.Fatalf("Failed to save jobs to file: %v", err)
	}

	fmt.Printf("✓ Job deleted: %s\n", jobID)
}

// handleStats processes the "stats" command
func handleStats(runner *JobRunner) {
	stats := runner.GetStats()
	fmt.Println("\n=== Job Statistics ===")
	fmt.Printf("Total Jobs:      %d\n", stats.TotalJobs)
	fmt.Printf("Pending:         %d\n", stats.PendingJobs)
	fmt.Printf("Running:         %d\n", stats.RunningJobs)
	fmt.Printf("Completed:       %d\n", stats.CompletedJobs)
	fmt.Printf("Failed:          %d\n", stats.FailedJobs)
	fmt.Println()
}

// printUsage displays the help text
func printUsage() {
	fmt.Println(`
GoLearn - Job Runner Learning Project (Day 2)
Usage: golearn <command> [options]

Commands:
  submit -id <id> -name <name> [-payload <payload>]
    Submit a new job
    
  status <job-id>
    Check the status of a job
    
  list
    List all jobs
    
  execute <job-id>
    Execute a job (synchronously)
    
  delete <job-id>
    Delete a job
    
  stats
    Show job statistics

  interactive
    Run in interactive mode (for quick testing)

  help
    Show this help message

Examples:
  golearn submit -id job1 -name "Process" -payload "data.csv"
  golearn list
  golearn execute job1
  golearn stats
  golearn delete job1
  golearn interactive
`)
}

// parseArgs splits a command line respecting quoted strings
// Example: `submit -id job1 -name "My Task"`
// Returns: []string{"submit", "-id", "job1", "-name", "My Task"}
func parseArgs(input string) []string {
	var args []string
	var current strings.Builder
	inQuotes := false

	for i := 0; i < len(input); i++ {
		ch := input[i]

		if ch == '"' {
			inQuotes = !inQuotes
			continue // Don't include the quote character itself
		}

		if ch == ' ' && !inQuotes {
			if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
			continue
		}

		current.WriteByte(ch)
	}

	if current.Len() > 0 {
		args = append(args, current.String())
	}

	return args
}

// runInteractiveMode allows multiple commands in a single session
func runInteractiveMode(runner *JobRunner) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("=== GoLearn Interactive Mode ===")
	fmt.Println("Type 'help' for commands, 'quit' to exit")
	fmt.Println()

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		parts := parseArgs(input)
		if len(parts) == 0 {
			continue
		}
		command := parts[0]

		switch command {
		case "quit", "exit":
			fmt.Println("Goodbye!")
			return
		case "submit":
			handleSubmit(runner, parts[1:])
		case "status":
			handleStatus(runner, parts[1:])
		case "list":
			handleList(runner)
		case "execute":
			handleExecute(runner, parts[1:])
		case "delete":
			handleDelete(runner, parts[1:])
		case "stats":
			handleStats(runner)
		case "help":
			printInteractiveHelp()
		default:
			fmt.Printf("Unknown command: %s\n", command)
		}
		fmt.Println()
	}
}

// printInteractiveHelp shows help for interactive mode
func printInteractiveHelp() {
	fmt.Println(`
Interactive Commands:
  submit -id <id> -name <name> [-payload <payload>]  - Submit a job
  status <job-id>                                      - Check job status
  execute <job-id>                                     - Run a job
  delete <job-id>                                      - Delete a job
  list                                                 - List all jobs
  stats                                                - Show statistics
  help                                                 - Show this help
  quit                                                 - Exit
  
Learning Goal: Observe how pointer receivers are necessary for mutations
Example session:
  > submit -id job1 -name "My Task" -payload "data.csv"
  > list
  > execute job1
  > stats
  > status job1
`)
}
