# GoLearn - Day 3 Progress: JSON Persistence + Error Handling

## ✅ What You've Built Today

Integrated **full job persistence** to disk. Your Job Runner now survives restarts with complete state preservation.

### Key Features Implemented

1. **JSON File Persistence**
   - All jobs automatically saved to `jobs.json` after mutations
   - Jobs automatically loaded on startup
   - State (including completed/failed status) is preserved

2. **File I/O + Error Handling**
   - `SaveJobsToFile()` - Uses `json.MarshalIndent()` for readable output
   - `LoadJobsFromFile()` - Uses `json.Unmarshal()` with error wrapping
   - `os.ReadFile()` / `os.WriteFile()` for disk access
   - Graceful handling of missing files (first run)

3. **New Method: AddJob()**
   - Distinguishes between new jobs (`Submit()`) and loaded jobs (`AddJob()`)
   - Critical fix: preserves job status from file instead of resetting to pending

4. **Automatic Persistence Integration**
   - Save after: `Submit()`, `ExecuteJob()`, `DeleteJob()`
   - Load once at startup
   - Works in both CLI and interactive modes

---

## 🧪 What You Learned Today

### JSON Encoding/Decoding

**Before (PHP/Symfony):**
```php
$jobs = json_encode($jobArray);
file_put_contents('jobs.json', $jobs);
```

**In Go:**
```go
data, err := json.MarshalIndent(jobs, "", "  ")
if err != nil {
    return fmt.Errorf("marshal failed: %w", err)
}
err = os.WriteFile("jobs.json", data, 0644)
```

Key differences:
- `json.MarshalIndent()` handles pretty-printing (safer than string building)
- File permissions are explicit (`0644` = rw-r--r--)
- Every step returns an error (no exceptions)

### Error Wrapping with `%w`

When you return errors, wrap them with context:

```go
// Bad: loses context
if err != nil {
    return err
}

// Good: provides context
if err != nil {
    return fmt.Errorf("failed to write file: %w", err)
}
```

The `%w` verb wraps the original error so the caller can still inspect it:

```go
err := SaveJobsToFile(runner, "jobs.json")
if errors.Is(err, os.ErrPermission) {
    // Handle permission error specifically
}
```

### The Critical Bug You Fixed (Submit vs AddJob)

**The Problem:**

```go
// OLD CODE (BUG)
func (jr *JobRunner) Submit(job *Job) error {
    // ... validate ...
    job.Status = StatusPending  // ← Always resets status!
    jr.jobs[job.ID] = job
    return nil
}

// When loading from file:
LoadJobsFromFile(...) {
    // Load job with status="completed"
    runner.Submit(job)  // ← Resets status to "pending"!
}
```

This is a **critical lesson**: Functions that mutate state have contracts. `Submit()` is for new jobs, not restored ones.

**The Solution:**

Created a separate `AddJob()` method:

```go
// For NEW jobs - sets initial status
func (jr *JobRunner) Submit(job *Job) error {
    job.Status = StatusPending
    // ... add to map ...
}

// For LOADED jobs - preserves existing status
func (jr *JobRunner) AddJob(job *Job) error {
    // Validates but does NOT change status
    // ... add to map ...
}
```

**This teaches you**: Design your APIs carefully. Don't conflate "create new" with "load existing".

---

## 📝 Error Handling Patterns You've Learned

### Pattern 1: File Not Found (Graceful Default)

```go
data, err := os.ReadFile(filename)
if err != nil {
    if os.IsNotExist(err) {
        return nil  // No file yet—fine for first run
    }
    return fmt.Errorf("failed to read: %w", err)
}
```

### Pattern 2: Error Wrapping

```go
data, err := json.MarshalIndent(jobs, "", "  ")
if err != nil {
    return fmt.Errorf("failed to marshal jobs: %w", err)
}
```

The `%w` preserves the original error. The caller can inspect it:

```go
if errors.Is(err, io.EOF) { ... }
if errors.Is(err, ErrInvalidJSON) { ... }
```

### Pattern 3: Validation Before Persistence

```go
if err := job.Validate(); err != nil {
    return fmt.Errorf("invalid job: %w", err)
}
// Only persist valid jobs
```

---

## 🔍 Files Modified Today

### `main.go`
- Added `const jobsFile = "jobs.json"`
- Load jobs from file at startup
- Save jobs after every mutation (Submit, Execute, Delete)
- Works in both CLI and interactive modes

### `runner.go`
- Split `Submit()` (for new jobs) and `AddJob()` (for loaded jobs)
- `Submit()` sets initial status to pending
- `AddJob()` preserves existing status

### `persistence.go`
- Implements `SaveJobsToFile()` - pretty JSON output
- Implements `LoadJobsFromFile()` - loads with error handling
- Uses `AddJob()` (not Submit) to preserve job state

### `golearn.exe`
- Rebuilt with persistence support
- `jobs.json` file created on first submit

---

## ✅ Verification Checklist (Day 3)

- ✅ Submit a job → `jobs.json` is created
- ✅ JSON has correct format (indented, readable)
- ✅ Exit the program and restart
- ✅ Jobs are loaded and listed correctly
- ✅ Execute a job, restart, status is preserved as "completed"
- ✅ Delete a job, restart, it's gone
- ✅ Job statuses survive restart (not reset to pending)
- ✅ Interactive mode also saves/loads
- ✅ First run with no `jobs.json` doesn't crash
- ✅ Error handling for file I/O issues works

---

## 🧠 Mental Model Shift: API Design Matters

Today's lesson: **Method names and contracts are part of the API**.

You designed two distinct methods:
- `Submit(job)` - "Create a new job with initial state"
- `AddJob(job)` - "Add an existing job preserving its state"

This is clearer than:
- `Submit(job)` - "Add a job (maybe new, maybe loaded, we'll decide)"

Go's explicitness shines here. Your code intentions are clear.

---

## 📊 Real-World Insight

This is how persistence works in production systems:

1. **Load** - Restore state from disk/database
2. **Process** - Handle requests that mutate state
3. **Save** - Persist changes immediately
4. **Restart** - Gracefully load and resume

You've just built this pattern. Every production database follows this cycle.

---

## 🚀 What's Next (Days 4-5 Preview)

### Day 4: Add a Benchmark

You'll time how long it takes to:
- Submit 10 jobs
- Execute all 10 synchronously

Expected: ~20 seconds (2 seconds per job)

This identifies the bottleneck that motivates Day 6-7 (concurrency).

### Days 5-6: Worker Pool Concurrency

When you hit the bottleneck, you'll add:
- Goroutines (lightweight concurrent tasks)
- Channels (communication between goroutines)
- Worker pool pattern (reusable workers)

Expected result: 10 jobs in ~5-7 seconds (3 workers in parallel)

---

## 📂 Current File Structure

```
c:\@personal\golearn\
├── go.mod
├── main.go              (loads/saves jobs, integrates persistence)
├── job.go
├── runner.go            (Submit vs AddJob distinction)
├── persistence.go       (SaveJobsToFile, LoadJobsFromFile)
├── golearn.exe          (rebuilt)
├── jobs.json            (auto-created on first submit)
├── README.md
├── DAY1_PROGRESS.md
├── DAY2_PROGRESS.md
└── DAY3_PROGRESS.md     (this file)
```

---

## 💾 JSON Format Example

```json
[
  {
    "id": "task1",
    "name": "First Task",
    "payload": "",
    "status": "completed",
    "created_at": "2026-04-16T08:10:01-03:00",
    "started_at": "2026-04-16T08:10:04-03:00",
    "completed_at": "2026-04-16T08:10:06-03:00"
  }
]
```

Notice:
- Human-readable (indented)
- Full state preserved (status, timestamps, error messages)
- Zero-value times are excluded in JSON output (or shown as "0001-01-01T00:00:00Z")

---

## 🎓 Key Takeaways

1. **JSON marshaling in Go is explicit** - You control formatting, struct tags, error handling
2. **Error wrapping preserves context** - Use `%w` to let callers inspect root causes
3. **API design matters** - `Submit()` vs `AddJob()` prevents silent state corruption
4. **Graceful degradation** - Missing files aren't errors on first run
5. **Persistence is foundational** - Everything after this assumes state survives restarts

---

## 🧪 Manual Testing (You Can Replay This Anytime)

```bash
# Clean slate
rm jobs.json

# Create jobs
.\golearn.exe submit -id task1 -name "First"
.\golearn.exe submit -id task2 -name "Second"

# Mutate
.\golearn.exe execute task1
.\golearn.exe delete task2

# Verify before restart
.\golearn.exe list
# Should show: task1 (completed)

# Simulate restart (fresh process)
.\golearn.exe list
# Should show the SAME: task1 (completed)

# Verify file
cat jobs.json
```

---

**Your Job Runner is now production-adjacent. It survives restarts, handles errors gracefully, and persists all state.**

Next: Day 4 will add benchmarking to show the bottleneck, then Days 6-7 unleash goroutines and channels.
