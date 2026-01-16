# Go Concepts Reference

> Your personal knowledge base. Add notes as you learn!
> This becomes your quick reference for Go patterns and principles.

---

## 📦 Go Language

### Variables

```go
// Long form
var name string = "Go"

// Type inference
var count = 42

// Short declaration (preferred inside functions)
message := "Hello"

// Zero values
var i int      // 0
var s string   // ""
var b bool     // false
var p *int     // nil
```

**Key Points**:
- [Add your notes here]

---

### Slices

```go
// Creation
slice := []int{1, 2, 3}
slice := make([]int, length, capacity)

// Operations
slice = append(slice, 4)
subSlice := slice[1:3]  // [2, 3]
```

**Key Points**:
- [Add your notes here]

**Gotchas**:
- [Add gotchas you discover]

---

### Maps

```go
// Creation
m := make(map[string]int)
m := map[string]int{"key": 1}

// Operations
m["key"] = value
delete(m, "key")

// Existence check
value, ok := m["key"]
if ok {
    // key exists
}
```

**Key Points**:
- [Add your notes here]

---

### Interfaces

```go
// Definition
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Implicit implementation
type MyReader struct{}

func (r MyReader) Read(p []byte) (n int, err error) {
    // implementation
}

// Type assertion
value, ok := interface.(ConcreteType)
```

**Key Points**:
- [Add your notes here]

---

### Concurrency

```go
// Goroutine
go func() {
    // runs concurrently
}()

// Channel
ch := make(chan int)
ch <- value  // send
v := <-ch    // receive

// Select
select {
case v := <-ch1:
    // handle ch1
case ch2 <- x:
    // send to ch2
default:
    // non-blocking
}
```

**Key Points**:
- [Add your notes here]

---

## 📚 Pragmatic Programmer Principles

### DRY (Don't Repeat Yourself)

**Definition**: Every piece of knowledge must have a single, unambiguous representation.

**In Go**:
- Extract repeated code into functions
- Use interfaces for common behavior
- Generics for type-varied duplication
- Table-driven tests

**Example**:
```go
// Before (WET)
if err != nil { log.Println(err); return err }
if err != nil { log.Println(err); return err }

// After (DRY)
func handleErr(err error) error {
    if err != nil { log.Println(err) }
    return err
}
```

**My Notes**:
- [Add observations from exercises]

---

### Orthogonality

**Definition**: Components should be independent; changes to one shouldn't affect others.

**In Go**:
- Package boundaries
- Interface segregation
- Avoid global state

**Signs of Poor Orthogonality**:
- Changing one thing breaks another
- Need to modify multiple files for one feature
- Circular dependencies

**My Notes**:
- [Add observations from exercises]

---

### Design by Contract

**Definition**: Functions have clear contracts: preconditions, postconditions, invariants.

**In Go**:
- Document expected inputs
- Validate inputs early
- Return errors for contract violations

**Example**:
```go
// Divide requires b != 0 (precondition)
// Returns a/b (postcondition)
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

**My Notes**:
- [Add observations from exercises]

---

### Fail Fast

**Definition**: Detect and report problems as soon as possible.

**In Go**:
- Check errors immediately
- Use early returns
- Panic for truly unrecoverable situations

**Pattern**:
```go
func Process(data []byte) error {
    if len(data) == 0 {
        return errors.New("empty data")  // Fail fast!
    }
    // ... rest of processing
}
```

**My Notes**:
- [Add observations from exercises]

---

## 🏗️ Architecture Patterns

### Layered Architecture

```
┌─────────────────────────┐
│    Presentation Layer   │  ← HTTP handlers, CLI
├─────────────────────────┤
│     Business Layer      │  ← Services, use cases
├─────────────────────────┤
│      Data Layer         │  ← Repositories, storage
└─────────────────────────┘
```

**Key Principle**: Dependencies point downward only.

**In Go**:
```
/handlers   → /services   → /repository
(HTTP)        (Business)     (Data)
```

**My Notes**:
- [Add observations from exercises]

---

### Hexagonal Architecture

```
                  ┌─────────────┐
    ┌─────────────┤    Core     ├─────────────┐
    │             │  (Domain)   │             │
    │             └─────────────┘             │
    │                                         │
┌───┴───┐                               ┌─────┴────┐
│ Ports │                               │ Adapters │
│(Interfaces)                           │(Implementations)
└───────┘                               └──────────┘
```

**Key Principle**: Core depends on nothing; everything depends on core.

**In Go**:
- Ports = Interfaces defined in core
- Adapters = Implementations in separate packages

**My Notes**:
- [Add observations from exercises]

---

### Repository Pattern

**Purpose**: Abstract data access from business logic.

```go
// Port (interface in domain)
type UserRepository interface {
    FindByID(id string) (*User, error)
    Save(user *User) error
}

// Adapter (implementation)
type PostgresUserRepo struct {
    db *sql.DB
}

func (r *PostgresUserRepo) FindByID(id string) (*User, error) {
    // SQL implementation
}
```

**My Notes**:
- [Add observations from exercises]

---

## 🔧 Go Idioms & Patterns

### Error Handling

```go
// Always check errors
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}

// Sentinel errors
var ErrNotFound = errors.New("not found")

// Checking specific errors
if errors.Is(err, ErrNotFound) {
    // handle not found
}
```

---

### Table-Driven Tests

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -1, -2},
        {"zero", 0, 0, 0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", 
                    tt.a, tt.b, got, tt.expected)
            }
        })
    }
}
```

---

### Functional Options

```go
type Server struct {
    port    int
    timeout time.Duration
}

type Option func(*Server)

func WithPort(p int) Option {
    return func(s *Server) { s.port = p }
}

func NewServer(opts ...Option) *Server {
    s := &Server{port: 8080}
    for _, opt := range opts {
        opt(s)
    }
    return s
}
```

---

## 📝 My Questions

<!-- Add questions as they come up -->

1. 
2. 
3. 

---

## 🔗 Useful Resources

- [Go Tour](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [The Pragmatic Programmer](https://pragprog.com/titles/tpp20/)

---

*Last updated: Never*
