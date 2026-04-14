# GoLearn - Day 1 Progress

## ✅ What You've Built Today

A basic **Job Runner** with the following features:

### Core Components (The Mental Model)

1. **Job Struct** (`job.go`)
   - Demonstrates how Go uses structs instead of classes
   - Fields with struct tags for JSON marshaling (you'll use this later)
   - Methods on structs:
     - `Validate()` - Returns an error (not exceptions!)
     - `Execute()` - Simulates job work with a 2-second sleep

2. **JobRunner Type** (`runner.go`)
   - Shows receiver methods (the `(jr *JobRunner)` syntax)
   - Uses of pointers: the receiver is a **pointer receiver** because we mutate the state
   - Mutex (`sync.Mutex`) for protecting shared state (you'll understand why this is needed on Day 6-7)
   - Key methods:
     - `NewJobRunner()` - Constructor pattern (explicit initialization)
     - `Submit(job *Job)` - Add jobs to storage
     - `GetStatus(jobID string)` - Retrieve job details
     - `ListJobs()` - Show all jobs

3. **CLI Interface** (`main.go`)
   - Demonstrates the `flag` package for argument parsing
   - Error handling: `if err != nil` pattern everywhere
   - Interactive mode for testing (easier than re-running the binary)

### Key Go Concepts You've Touched

| Concept | Where You See It | Why Go Works This Way |
|---------|------------------|----------------------|
| **Structs** | `Job` and `JobRunner` | Go has no classes—composition over inheritance |
| **Methods** | `(j *Job) Validate()`, `(jr *JobRunner) Submit()` | Methods are attached to types via receivers |
| **Pointers** | `*Job`, `*JobRunner` | Allows mutation and shared references |
| **Error Handling** | `if err != nil` | Errors are values, not exceptions (explicit control flow) |
| **Mutex** | `jr.mu.Mutex` in JobRunner | Protects shared state (you'll see why this matters on Day 6) |
| **Struct Tags** | `` json:"id"`` | Metadata for encoding/JSON (hook for reflection) |

---

## 🧪 How to Use Day 1 Build

### Interactive Mode (Recommended for Testing)

```bash
cd c:\@personal\golearn
.\golearn.exe interactive
```

Then type commands:
```
> submit -id job1 -name "My Task" -payload "data.csv"
> list
> status job1
> quit
```

### Non-Interactive Mode (One Command at a Time)

```bash
.\golearn.exe submit -id job1 -name "Task 1"
.\golearn.exe status job1
.\golearn.exe list
```

---

## 🔍 What You SHOULD Notice (Mental Model Resets)

### 1. **No Classes, But Structure Is Clear**
   - Go has no base classes or inheritance
   - Yet, the code is organized: `Job` and `JobRunner` are clear concepts
   - This is **composition over inheritance**—a foreign idea if you know PHP/Symfony

### 2. **Pointers Aren't Scary**
   - `JobRunner` uses a pointer receiver: `(jr *JobRunner)`
   - Why? Because `Submit()` mutates the `jobs` map
   - Go forces you to be explicit: if you want to change state, you need a pointer
   - This is a safety feature, not a burden

### 3. **Errors Are Values, Not Exceptions**
   - No try/catch blocks
   - `Validate()` returns an error: `error` (a built-in interface)
   - The caller decides what to do: log it, return it, ignore it? They control the flow
   - This is intentional—errors are as important as success paths

### 4. **Mutex Is Here but We're Not Using It Properly Yet**
   - You'll see `mu sync.Mutex` and `defer jr.mu.Unlock()`
   - In Day 1, it looks like overkill (single-threaded CLI)
   - Keep it here—Day 6+ when you add goroutines, you'll see why this matters

---

## 📝 Code Reading Exercise

Open `job.go` and answer these questions:

1. **Why is the Job struct useful?** (Not a class, just data)
   - What are the five fields? What do they represent?
   - How is `Status` different from a string? (It's a custom type!)

2. **The Validate() method:**
   - Why return error instead of throwing an exception?
   - What happens in the caller if `Validate()` returns an error?

3. **In runner.go, what does this line do?**
   ```go
   jr.mu.Lock()
   defer jr.mu.Unlock()
   ```
   - `defer` is Go's way of saying "do this when the function exits"
   - This guarantees the mutex is unlocked (no deadlocks)
   - Why might this be important? (Think: if an error happens mid-function)

---

## 🚀 What's Next (Preview of Day 2)

**Day 2: Refactor + Pointers + Struct Composition**

Tomorrow you'll:
1. Move structs to separate files (organization)
2. Experience the first pointer error (and fix it)
3. Add more methods that show why pointers are essential

We're keeping the same project—just improving the core structure.

---

## 🎯 Reflection: Was This Efficient?

- **Time spent:** ~30 min reading, 20 min coding, 10 min testing = 60 min
- **Code written:** ~200 lines
- **Concepts embedded:** Structs, methods, pointers, error handling, mutex awareness
- **No tutorials read:** You built and tested immediately

This is the Go way: small pieces, test often, understand by doing.

---

## 📂 File Structure So Far

```
c:\@personal\golearn\
├── go.mod              # Module definition
├── go.sum              # (Created by `go mod` if you added dependencies)
├── job.go              # Job struct + methods
├── runner.go           # JobRunner + core logic
├── main.go             # CLI entry point
└── golearn.exe         # Compiled binary
```

---

## ⚠️ Known Limitations (Fixed in Days 2-4)

1. **Jobs aren't persisted** — Every restart loses all jobs (fixed on Day 3-4)
2. **Jobs don't actually execute** — They're stored but not run (fixed on Day 5)
3. **No concurrency** — Everything runs sequentially (fixed on Day 6+)
4. **No logging** — Just basic fmt.Println output (fixed on Day 10)

This is intentional—we're building a foundation first.
