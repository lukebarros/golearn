# GoLearn: Learn Go the Right Way

> A structured learning roadmap for backend developers transitioning from Python/PHP to Go. Based on real experience with 3 YOE in Python and PHP (Symfony).

**Status**: Day 1 Complete ✅ | Interactive Job Runner Ready

---

## 🚀 The Big Picture

This repo documents a **10-12 day learning journey** to build real Go competence through building a single project: a **Job Runner system** that progressively introduces concurrency, error handling, backend patterns, and production concerns.

### Why This Approach?

Most Go tutorials teach syntax in isolation. This roadmap teaches concepts through **immediate application**:

1. Learn a pattern (1-2 hours)
2. Implement it in the Job Runner
3. Hit a real problem that motivates the *next* concept
4. Repeat

### Who This Is For

- Backend developers with **3+ years Python/PHP experience**
- Tired of frameworks and magic (Symfony, Django)
- Want to truly understand concurrency, not just use it
- Have 1-2 hours daily to dedicate
- Ready to challenge your mental model

---

## 📋 The Brutal Truth (Why Most People Fail)

If you approach Go like "PHP but faster," you will plateau quickly. Go is opinionated:

| Aspect | Go's Way | Your Background | Adjustment Needed |
|--------|----------|-----------------|-------------------|
| **Abstraction** | Simplicity over abstraction | Heavy frameworks (Symfony) | Break the interface addiction |
| **Code organization** | Small packages | Namespaced classes | Think composition, not inheritance |
| **Composition** | Interfaces + structs | Base classes + inheritance | Invert your thinking |
| **Concurrency** | First-class primitive | Thread pools / async callbacks | Learn goroutines deeply |
| **Error handling** | Values, not exceptions | Try/catch mindset | Explicit control flow |
| **Dependencies** | Manual wiring | Autowiring containers | Be explicit from Day 1 |

**You must embrace Go's way, not bend Go to your way.**

---

## 📚 Full Roadmap: 10-12 Days to Competence

### Phase 1: Reset Your Mental Model (Before You Code)

**Duration**: Think about this while starting Day 1

Go is opinionated about:
- ✅ **Simplicity over abstraction** — No clever frameworks, no magic
- ✅ **Composition over inheritance** — Structs + interfaces, not base classes
- ✅ **Explicit over magical** — If it's important, make it visible
- ✅ **Concurrency as a first-class primitive** — Not add-ons, built into language

**Things you'll stop leaning on:**
- ❌ Heavy frameworks (Symfony, Laravel)
- ❌ ORMs doing magic (Doctrine, OpenQuery)
- ❌ Deep class hierarchies
- ❌ Dependency injection containers
- ❌ Exception handling

**Things you must embrace:**
- ✅ Small packages (organize by domain, not pattern)
- ✅ Interfaces used sparingly (only when needed for abstraction)
- ✅ Explicit error handling (every error is important)
- ✅ Standard library dominance (net/http, encoding/json, database/sql are powerful)

---

### Phase 2: Syntax + Core Concepts (Days 1-2)

**Duration**: 2-3 days max

**You'll learn:**
- Types, structs, methods
- Receivers (value vs. pointer)
- Interfaces (but don't overuse them)
- Error handling (`if err != nil`)
- Slices vs arrays (critical distinction)
- Maps
- Pointers (simpler than C, but essential)

**Building**: Core Job Runner
- `Job` struct with validation
- `JobRunner` type with methods
- Basic in-memory storage
- CLI interface

**Key difference vs. PHP/Python:**
- No exceptions → errors are values you return
- No classes → structs + methods (simpler!)
- Interfaces are **implicit** (huge concept change)
- Pointers force you to be explicit about mutation

**Deliverable**: A CLI that creates and retrieves jobs (synchronously)

---

### Phase 3: The Real Go—Concurrency (Days 3-7)

**Duration**: 5 days (this is critical)

This is where Go becomes valuable. Without this, you're writing inferior Node.js/PHP in Go.

**Learn deeply:**
- Goroutines (lightweight threads, created by Go runtime)
- Channels (how goroutines communicate)
- The `select` statement (multiplex channel operations)
- Worker pools (idiomatic concurrency pattern)
- `context.Context` (VERY important for backend; every handler needs it)
- Graceful shutdown (prevents goroutine leaks)

**Building**: Job Runner with Concurrency
- Day 5: Benchmark synchronous execution, identify bottleneck
- Day 6-7: Implement worker pool pattern (3-5 workers)
- Day 8-9: Add context + graceful shutdown
- Observe: 10 jobs on 3 workers = ~7 seconds (vs. 20 seconds sequential)

**Key insight**: Goroutines are not threads. The Go runtime multiplexes them onto OS threads. Creating 10,000 goroutines is fine; creating 10,000 threads will crash.

**Deliverable**: A robust concurrent job runner that handles cancellation cleanly

---

### Phase 4: Build Backend Systems (Days 8-9)

**Duration**: 2 days

Now connect to what you already know from Symfony.

**Replace your Symfony mental model:**

| Symfony | Go Equivalent |
|---------|---------------|
| Controllers | HTTP handlers (functions or methods) |
| Services | Plain structs/functions |
| Dependency Injection | Manual wiring (no container) |
| Middleware | Middleware (simpler pattern) |
| Routing | `net/http` + chi/gorilla (start simple) |
| Validation | Explicit struct methods (no annotations) |

**Learn:**
- `net/http` package (don't start with frameworks)
- JSON encoding/decoding (`json.Marshal/Unmarshal`)
- Routing (start with `net/http` mux, then maybe chi or gin)
- Middleware patterns (function composition, not decorators)
- Request context handling (you learned `context.Context`—use it here)

**Building**: HTTP API for your Job Runner
- POST `/jobs` → submit a new job
- GET `/jobs/{id}` → get job details
- GET `/jobs` → list all jobs
- Middleware for logging and error handling

**Deliverable**: A working REST API without a framework

---

### Phase 5: Database Layer (Days 10-11)

**Duration**: 2 days

Given your PostgreSQL experience from PHP/Python, you'll excel here (most Go devs are weak in SQL depth).

**Learn:**
- `database/sql` (the standard library interface)
- Connection pooling (`sql.DB` handles this)
- Transactions (ACID semantics)
- Prepared statements (prevent SQL injection)
- Query result scanning (`rows.Scan()`)

**Then:**
- Try `sqlc` (generates type-safe SQL from your queries—very aligned with Go's explicit style)
- Compare with ORMs like `gorm` (you'll likely dislike them—good instinct)
- Understand the trade-off: raw SQL vs. generated code vs. ORM abstraction

**Building**: Persist Job Runner to PostgreSQL
- Schema: `jobs` table (id, name, payload, status, created_at, etc.)
- Migration approach (hand-written SQL initially, then maybe `migrate` tool)
- Update `JobRunner` to read/write from Postgres
- Connection pooling configuration

**Leverage point**: Your SQL depth is your differentiator. Most Go devs skip this and use ORMs.

**Deliverable**: Job Runner backed by real database, with proper connection pooling

---

### Phase 6: Production-Grade Concerns (Day 12+)

**Duration**: 1+ day

This is where Go shines compared to PHP/Symfony.

**Learn:**
- Profiling with `pprof` (CPU, memory, goroutine profiles)
- Memory management basics (escape analysis, stack vs. heap)
- Goroutine leak detection (`runtime.NumGoroutine()`)
- Graceful shutdowns (signal handling, cleanup)
- Structured logging (zap, slog, or simple json output)
- Observability: logs, metrics (prometheus-compatible)
- Health checks and readiness probes (for production deployment)

**Building**: Production-ready Job Runner
- Add structured logging (who, what, when?)
- Metrics: job processing rate, latency, error rate
- Graceful shutdown: drain jobs on SIGTERM
- Profiling: `pprof` endpoints for debugging
- Health check endpoint

**Why this matters**: Go applications run forever. Goroutine leaks, memory leaks, and poor shutdown handling will crash your service in production.

**Deliverable**: A production-grade Job Runner ready for deployment

---

## 🎯 What You're NOT Doing (Don't Waste Time)

- ❌ Over-learning syntax (2 days max for basics)
- ❌ Building toy CRUD apps ("build a blog in 30 minutes")
- ❌ Framework hopping (chi vs. gin vs. echo vs. fiber)
- ❌ Overusing interfaces everywhere ("make everything an interface")
- ❌ Deep diving into reflection (learn when you need it)
- ❌ Memorizing the standard library (use docs)

**Focus on building one real system end-to-end.**

---

## 📅 Suggested Timeline

### Week 1: Foundation
- **Day 1-2** (Phase 1-2): Syntax + Core (structs, methods, error handling)
  - Deliverable: Synchronous Job Runner CLI
  
- **Day 3-4** (Phase 2 continued): Methods, pointers, deeper understanding
  - Deliverable: Add JSON persistence, validation
  
- **Day 5** (Phase 3 start): Hit the bottleneck
  - Deliverable: Benchmark shows sequential bottleneck

### Week 2: Concurrency + Backend
- **Day 6-7** (Phase 3): Goroutines + channels (worker pool)
  - Deliverable: Concurrent Job Runner (3x speedup)
  
- **Day 8-9** (Phase 3 + 4): Context + graceful shutdown + HTTP API
  - Deliverable: REST API with concurrent job processing
  
- **Day 10** (Phase 4 + 5): Database layer (PostgreSQL)
  - Deliverable: Job Runner backed by real database

### Week 3+: Polish + Production
- **Day 11-12** (Phase 6): Production concerns (logging, metrics, profiling)
  - Deliverable: Production-ready system

---

## 💪 Efficient Learning Strategy (Based on Your Profile)

**Do this instead of passive learning:**

1. **Learn concept** (1-2 hours reading/docs)
2. **Immediately implement** it in the Job Runner
3. **Break it under load** or with edge cases
4. **Fix it** (now you understand why the pattern exists)

This feedback loop matters more than 10x any tutorial.

### For Weak Areas (Math Intuition + Concurrency Thinking)

Go will force you to:
- Think about **resource usage** (how many goroutines? channels?)
- Reason about **parallel execution** (race conditions, deadlocks)
- Understand **trade-offs** explicitly (throughput vs. latency, memory vs. speed)

**You can't avoid this.** If you resist the discomfort, you'll stagnate.

---

## 🏗️ Project Structure

```
c:\@personal\golearn\
├── README.md                # This file
├── DAY1_PROGRESS.md         # Day 1 reflection
├── go.mod                   # Module definition
├── main.go                  # CLI entry point
├── job.go                   # Job struct + methods
├── runner.go                # JobRunner type
├── persistence.go           # (Days 3-4) JSON file I/O
├── database.go              # (Day 10) PostgreSQL integration
├── server.go                # (Days 8-9) HTTP server
├── logger.go                # (Day 12) Structured logging
└── golearn.exe              # Compiled binary
```

---

## 🚀 Getting Started

### Prerequisites
- Go 1.23+ (installed via `winget install GoLang.Go` on Windows)
- PostgreSQL 13+ (for Phase 5, optional until Day 10)

### Running the Project

**Interactive mode (for testing):**
```bash
cd c:\@personal\golearn
.\golearn.exe interactive
```

Commands:
```
> submit -id job1 -name "Process Data" -payload "data.csv"
> list
> status job1
> quit
```

**Build from source:**
```bash
go build -o golearn.exe
```

---

## 🧠 The Mental Model Shift You Need

### From PHP/Symfony → Go Thinking

#### 1. Classes → Structs + Methods

```go
// Old (PHP/Symfony)
class JobRunner {
    private $jobs = [];
    public function submit(Job $job) { ... }
}

// New (Go)
type JobRunner struct {
    jobs map[string]*Job
}
func (jr *JobRunner) Submit(job *Job) error { ... }
```

**Key difference**: Receivers are explicit. If you need to mutate state, you use a **pointer receiver** `*JobRunner`. If you just read, you use a **value receiver** `JobRunner`.

#### 2. Exceptions → Error Values

```go
// Old (PHP/Symfony)
try {
    $job = $this->jobService->submit($jobData);
} catch (ValidationException $e) {
    return error($e->getMessage());
}

// New (Go)
job := &Job{...}
if err := job.Validate(); err != nil {
    return fmt.Errorf("validation failed: %w", err)
}
```

**Key difference**: Error handling is inline. Every possible error path is visible. Less magic, more control.

#### 3. Dependency Injection Container → Manual Wiring

```go
// Old (Symfony)
class MyController {
    public function __construct(JobRepository $repo, Logger $logger) { ... }
}
// Framework injects dependencies

// New (Go)
func NewJobRunner(logger Logger, db *sql.DB) *JobRunner {
    return &JobRunner{logger: logger, db: db}
}
// You control the flow
```

**Key difference**: You're responsible for wiring dependencies. This is tedious at first, but you always know where dependencies come from.

#### 4. Goroutines ≠ Threads

```go
// Old (parallel thinking)
// Each request = 1 thread (expensive, limited pool)

// New (Go thinking)
// Each request spawns goroutines (lightweight, millions possible)
// The Go runtime multiplexes them efficiently
for {
    conn, _ := listener.Accept()
    go handleConnection(conn)  // Cheap, reusable
}
```

**Key difference**: Goroutines are not OS threads. They're managed by the Go runtime. You can safely create thousands of them.

#### 5. Channels > Locks

```go
// Old (shared memory thinking)
$results = [];
$mutex->lock();
$results[] = $computation();
$mutex->unlock();

// New (Go thinking)
results := make(chan int)
go func() {
    results <- computation()
}()
value := <-results  // Safe, built-in synchronization
```

**Key difference**: Goroutines communicate via channels, not shared memory. This prevents race conditions elegantly.

---

## ✅ Verification Checkpoints

### After Phase 1-2 (Days 1-2): Syntax Mastery
- ✅ Can create structs with methods
- ✅ Understand pointer receivers and value receivers
- ✅ Error handling is comfortable (if err != nil)
- ✅ Can use maps and slices confidently
- ✅ No panic about pointers

### After Phase 3 (Days 6-7): Concurrency Understanding
- ✅ Understand goroutines (not threads)
- ✅ Can create channels and communicate between goroutines
- ✅ Worker pool pattern is clear
- ✅ `context.Context` usage is natural
- ✅ Can identify and prevent goroutine leaks

### After Phase 4-5 (Days 8-11): Backend Competence
- ✅ HTTP handlers without a framework
- ✅ JSON marshaling/unmarshaling
- ✅ Database transactions and prepared statements
- ✅ Connection pooling intuition
- ✅ SQL depth is a strength

### After Phase 6 (Day 12+): Production Ready
- ✅ Can profile a system (pprof)
- ✅ Understand memory management basics
- ✅ Graceful shutdown is automated
- ✅ Observability (logging + metrics) is built in
- ✅ No goroutine leaks, no memory leaks

---

## 🎓 Key Concepts to Get Right

### 1. Interfaces Are Implicit

Go interfaces don't require explicit declaration. If a type has the methods an interface needs, it satisfies that interface.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

// This satisfies Reader WITHOUT declaring it
type MyFile struct { ... }
func (f *MyFile) Read(p []byte) (int, error) { ... }
```

**Why?** Decoupling. You can write code that works with any reader without the reader knowing about your interface.

### 2. Composition Over Inheritance

Go has no inheritance. Instead, you compose structs:

```go
type Engine struct { ... }
type Car struct {
    Engine  // Embedded, not inherited
    Wheels int
}

car.Start()  // Delegates to Engine if Engine has Start()
```

### 3. Package-Level Organization

Go organizes code by packages, not by pattern. A package is a folder:

```
myapp/
├── job/        // All job-related code
│   ├── job.go
│   ├── runner.go
│   └── store.go
├── server/     // HTTP server code
│   ├── handler.go
│   └── middleware.go
└── main.go
```

Not:

```
myapp/
├── models/
├── services/
├── controllers/
└── repositories/
```

### 4. Goroutines + Channels = Concurrency Primitive

This is the biggest mind-shift. You don't manage threads; you spawn lightweight goroutines and let the runtime handle them.

```go
// Spawn 1000 goroutines? Fine.
for i := 0; i < 1000; i++ {
    go worker()
}

// They'll run efficiently. The runtime multiplexes them.
```

---

## 🔥 Common Pitfalls (You'll Hit These, That's OK)

1. **Forgetting to close channels** → Deadlocks in concurrent code
2. **Forgetting `defer mu.Unlock()`** → Race conditions
3. **Forgetting `context.Context`** → Goroutine leaks
4. **Using value receivers when you need pointers** → Receiver not updated
5. **Ignoring errors silently** → Hidden bugs
6. **Over-using interfaces** → Unnecessary abstraction

**Each pitfall is a learning opportunity.**

---

## 📖 Further Reading (Not Required)

When you hit a concept you want deeper understanding:

- **Official Go Blog**: https://go.dev/blog (excellent for concurrency patterns)
- **Effective Go**: https://go.dev/doc/effective_go (style guide + patterns)
- **Go Memory Model**: https://go.dev/ref/mem (concurrency semantics)
- **Standard Library Docs**: https://pkg.go.dev/std (gold mine)

Don't read these upfront. Read them when you need to understand something specific.

---

## 🏁 By the End (Week 3+)

You'll have built a **production-grade Job Runner** that is:

- ✅ Concurrent (worker pool pattern)
- ✅ Gracefully shutdown-capable
- ✅ Backed by PostgreSQL
- ✅ With a REST API
- ✅ Structured logging
- ✅ Metrics + profiling
- ✅ No goroutine leaks
- ✅ No memory leaks

More importantly, you'll **understand Go's philosophy** and why it's different.

---

## 💡 One More Thing

> "Your main bottleneck is not Go itself. It's how you think about resource usage, parallel execution, and trade-offs. Go will force you to think intentionally. If you resist that discomfort, you'll stagnate."

The roadmap is challenging because Go is cleaner than PHP/Symfony. There's less to hide behind. Embrace that.

---

**Start with Day 1. Build something. Test it. Understand it. Move to Day 2.**

The feedback loop is your teacher.
