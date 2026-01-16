# Go Learning Curriculum

> **Your complete learning path: Go Tour + Pragmatic Programmer + Architecture**
> 
> Progress: Track with checkboxes. Each topic requires understanding the concept,
> completing exercises, and passing mastery verification.

---

## 📋 Legend

- `[ ]` = Not started
- `[x]` = Completed & Mastered
- `🔄` = In progress (replace `[ ]` with this when working on it)
- `→` = Prerequisite arrow (must complete before)

---

## Phase 1: Foundations (Go Tour Basics)

> **Goal**: Master Go syntax and basic constructs
> **Pragmatic Focus**: DRY, basic error handling
> **Estimated Time**: 1-2 weeks

### 1.1 Packages & Imports
- [ ] Packages and main
- [ ] Imports and exported names
- [ ] **Exercise**: Create a multi-file program

**Go Tour Sections**: Packages, Imports, Exported names

### 1.2 Variables & Types
- [ ] Variables and var statements
- [ ] Short variable declarations `:=`
- [ ] Basic types (int, float64, string, bool)
- [ ] Zero values
- [ ] Type conversions
- [ ] Constants
- [ ] **Exercise**: Config struct with type-safe settings

**Go Tour Sections**: Variables, Short declarations, Basic types, Zero values, Type conversions, Constants

**Pragmatic Integration**: 
- DRY: Avoid repeated type declarations
- Tip: "Don't repeat yourself" - extract common types

### 1.3 Functions
- [ ] Function syntax and multiple arguments
- [ ] Multiple return values
- [ ] Named return values
- [ ] Variadic functions
- [ ] **Exercise**: Calculator with clean function design

**Go Tour Sections**: Functions, Multiple results, Named results

**Pragmatic Integration**:
- DRY: Extract common logic into functions
- Tip: Functions should do one thing well

### 1.4 Flow Control
- [ ] For loops (Go's only loop construct)
- [ ] If statements
- [ ] If with short statement
- [ ] Switch statements
- [ ] Defer
- [ ] **Exercise**: FizzBuzz with proper control flow

**Go Tour Sections**: For, If, Switch, Defer

**Pragmatic Integration**:
- Fail Fast: Use early returns
- Tip: "Crash early" - validate inputs at function start

---

## Phase 2: Data Structures

> **Goal**: Master Go's built-in data structures
> **Pragmatic Focus**: Orthogonality, data-driven design
> **Architecture**: Introduction to layered design
> **Prerequisites**: Phase 1 complete
> **Estimated Time**: 2-3 weeks

### 2.1 Pointers
- [ ] Pointers and memory
- [ ] Pointer dereferencing
- [ ] No pointer arithmetic in Go
- [ ] **Exercise**: Swap function, value vs pointer semantics

**Go Tour Sections**: Pointers

### 2.2 Arrays
- [ ] Array declaration and initialization
- [ ] Array length is part of type
- [ ] Arrays are values (copied)
- [ ] **Exercise**: Matrix operations

**Go Tour Sections**: Arrays

### 2.3 Slices ⭐ (Critical Topic)
- [ ] Slice basics and literals
- [ ] Slice length and capacity
- [ ] Nil slices
- [ ] Creating with make
- [ ] Appending to slices
- [ ] Slice internals (header, backing array)
- [ ] Range iteration
- [ ] **Exercise**: Dynamic inventory system

**Go Tour Sections**: Slices, Slice length and capacity, Nil slices, Making slices, Appending, Range

**Pragmatic Integration**:
- DRY: Table-driven operations on slices
- Tip: "Eliminate effects between unrelated things"

**⚠️ Common Gotchas**: 
- Slice aliasing (shared backing array)
- Append may create new array

### 2.4 Maps
- [ ] Map declaration and literals
- [ ] Mutating maps (insert, delete, lookup)
- [ ] Zero value of map is nil
- [ ] Comma-ok idiom for existence check
- [ ] **Exercise**: Word frequency counter

**Go Tour Sections**: Maps, Mutating Maps

**Pragmatic Integration**:
- Orthogonality: Maps for decoupled key-value relationships
- Design by Contract: Always check map existence

### 2.5 Structs
- [ ] Struct declaration
- [ ] Struct fields and access
- [ ] Struct literals
- [ ] Pointers to structs
- [ ] Embedded structs (composition)
- [ ] **Exercise**: User management with nested structs

**Go Tour Sections**: Structs, Struct Fields, Pointers to structs, Struct Literals

**Architecture Introduction**:
- Entity design
- Data Transfer Objects (DTOs)
- Separation of data and behavior

---

## Phase 3: Methods & Interfaces

> **Goal**: Master Go's approach to OOP through composition
> **Pragmatic Focus**: Design by Contract, Orthogonality, Reversibility
> **Architecture**: Dependency Injection, Hexagonal Architecture
> **Prerequisites**: Phase 2 complete (especially Structs)
> **Estimated Time**: 2-3 weeks

### 3.1 Methods
- [ ] Method syntax
- [ ] Value receivers
- [ ] Pointer receivers
- [ ] When to use pointer vs value receivers
- [ ] Methods on non-struct types
- [ ] **Exercise**: Bank account with methods

**Go Tour Sections**: Methods, Pointer receivers, Methods continued

**Pragmatic Integration**:
- Design by Contract: Methods define the contract
- Tip: "Design with contracts" - clear pre/post conditions

### 3.2 Interfaces ⭐ (Critical Topic)
- [ ] Interface definition
- [ ] Implicit implementation
- [ ] Interface values
- [ ] Empty interface `any`
- [ ] Type assertions
- [ ] Type switches
- [ ] **Exercise**: Payment processor with multiple payment types

**Go Tour Sections**: Interfaces, Interfaces are satisfied implicitly, Interface values, Empty interface, Type assertions, Type switches

**Pragmatic Integration**:
- Orthogonality: Interfaces decouple components
- Reversibility: Swap implementations easily
- Tip: "Program to an interface, not an implementation"

**Architecture**:
- Dependency Injection pattern
- Repository pattern introduction
- Hexagonal Architecture (Ports = Interfaces, Adapters = Implementations)

### 3.3 Common Interfaces
- [ ] Stringer interface
- [ ] Error interface
- [ ] Reader/Writer interfaces
- [ ] **Exercise**: Custom logger implementing standard interfaces

**Go Tour Sections**: Stringers, Errors, Readers

### 3.4 Error Handling
- [ ] Error type and creation
- [ ] Error wrapping (Go 1.13+)
- [ ] errors.Is and errors.As
- [ ] Panic and recover
- [ ] When to use panic vs error
- [ ] **Exercise**: File processor with comprehensive error handling

**Go Tour Sections**: Errors (extended)

**Pragmatic Integration**:
- Fail Fast: Return errors immediately
- Crash Early: Panic only for unrecoverable situations
- Tip: "Use exceptions for exceptional problems"

---

## Phase 4: Generics

> **Goal**: Master Go's type parameters for reusable code
> **Pragmatic Focus**: DRY at the type level
> **Prerequisites**: Interfaces (Phase 3.2)
> **Estimated Time**: 1-2 weeks

### 4.1 Type Parameters
- [ ] Generic function syntax
- [ ] Type parameter constraints
- [ ] The `any` and `comparable` constraints
- [ ] **Exercise**: Generic stack implementation

**Go Tour Sections**: Type parameters

### 4.2 Generic Types
- [ ] Generic type definitions
- [ ] Generic methods
- [ ] Type inference
- [ ] **Exercise**: Generic cache with TTL

**Go Tour Sections**: Generic types

### 4.3 Type Constraints
- [ ] Interface as constraint
- [ ] Union constraints
- [ ] The `constraints` package
- [ ] **Exercise**: Generic math operations

**Pragmatic Integration**:
- DRY: Eliminate type-specific code duplication
- Tip: Write generic code only when you have 3+ duplications

---

## Phase 5: Concurrency ⭐

> **Goal**: Master Go's concurrency primitives
> **Pragmatic Focus**: Fail Fast in concurrent code, Orthogonality
> **Architecture**: Event-Driven, Producer-Consumer patterns
> **Prerequisites**: Phase 3 complete (especially Interfaces)
> **Estimated Time**: 3-4 weeks

### 5.1 Goroutines
- [ ] Goroutine basics
- [ ] Starting goroutines
- [ ] Anonymous goroutines
- [ ] Goroutine lifecycle
- [ ] **Exercise**: Concurrent prime number finder

**Go Tour Sections**: Goroutines

### 5.2 Channels
- [ ] Channel basics and creation
- [ ] Sending and receiving
- [ ] Buffered channels
- [ ] Channel directions (send-only, receive-only)
- [ ] Closing channels
- [ ] Range over channels
- [ ] **Exercise**: Pipeline processor

**Go Tour Sections**: Channels, Buffered Channels, Range and Close

**Architecture**:
- Producer-Consumer pattern
- Pipeline pattern

### 5.3 Select
- [ ] Select basics
- [ ] Select with default
- [ ] Select for timeouts
- [ ] Non-blocking operations
- [ ] **Exercise**: Multiplexer with timeout

**Go Tour Sections**: Select, Default Selection

### 5.4 Sync Package
- [ ] sync.Mutex
- [ ] sync.RWMutex
- [ ] sync.WaitGroup
- [ ] sync.Once
- [ ] sync.Map (when to use)
- [ ] **Exercise**: Thread-safe counter with benchmarks

**Go Tour Sections**: sync.Mutex

**Pragmatic Integration**:
- Fail Fast: Handle goroutine errors properly
- Orthogonality: Keep concurrent components independent
- Tip: "Don't share memory; share by communicating"

### 5.5 Context Package
- [ ] Context basics
- [ ] Context with cancel
- [ ] Context with timeout
- [ ] Context with values
- [ ] Passing context through call chains
- [ ] **Exercise**: Cancellable web crawler

**Architecture**:
- Event-Driven Architecture
- Graceful shutdown patterns

---

## Phase 6: Real-World Patterns

> **Goal**: Build production-ready Go applications
> **Pragmatic Focus**: All principles combined
> **Architecture**: Microservices, API design, Middleware
> **Prerequisites**: All previous phases
> **Estimated Time**: 3-4 weeks

### 6.1 HTTP Servers
- [ ] net/http basics
- [ ] Handlers and HandlerFunc
- [ ] Routing patterns
- [ ] Request parsing
- [ ] Response writing
- [ ] **Exercise**: RESTful API for todo list

### 6.2 Middleware Patterns
- [ ] Middleware concept
- [ ] Chaining middleware
- [ ] Common middleware (logging, auth, recovery)
- [ ] **Exercise**: API with authentication middleware

**Architecture**:
- Middleware chain pattern
- Decorator pattern in Go

### 6.3 Testing in Go
- [ ] Table-driven tests
- [ ] Test helpers
- [ ] Mocking with interfaces
- [ ] Benchmarks
- [ ] Test coverage
- [ ] **Exercise**: Full test suite for previous projects

**Pragmatic Integration**:
- DRY: Table-driven tests eliminate repetition
- Tip: "Test your software, or your users will"

### 6.4 JSON & Encoding
- [ ] JSON marshaling/unmarshaling
- [ ] Struct tags
- [ ] Custom marshalers
- [ ] **Exercise**: API client with JSON handling

### 6.5 Project Organization
- [ ] Standard Go project layout
- [ ] Internal packages
- [ ] Dependency management (go.mod)
- [ ] **Exercise**: Refactor projects into proper structure

**Architecture**:
- Clean Architecture in Go
- Hexagonal Architecture implementation

---

## 🏗️ Capstone Projects

Complete these to demonstrate mastery:

### Project 1: CLI Task Manager
- [ ] Uses: Slices, Maps, Structs, Files, Flags
- [ ] Principles: DRY, Orthogonality
- [ ] Pattern: Layered Architecture

### Project 2: REST API Service
- [ ] Uses: HTTP, JSON, Interfaces, Testing
- [ ] Principles: Design by Contract, Fail Fast
- [ ] Pattern: Hexagonal Architecture

### Project 3: Concurrent Data Processor
- [ ] Uses: Goroutines, Channels, Context
- [ ] Principles: All
- [ ] Pattern: Event-Driven Architecture

---

## 📊 Pragmatic Principles Tracker

Track which principles you've practiced:

- [ ] **DRY** (Don't Repeat Yourself) - First applied in: ___
- [ ] **Orthogonality** - First applied in: ___
- [ ] **Reversibility** - First applied in: ___
- [ ] **Tracer Bullets** - First applied in: ___
- [ ] **Prototyping** - First applied in: ___
- [ ] **Design by Contract** - First applied in: ___
- [ ] **Fail Fast** - First applied in: ___
- [ ] **Decoupling** - First applied in: ___

---

## 🏛️ Architecture Patterns Tracker

Track which patterns you've implemented:

- [ ] **Layered Architecture** - First applied in: ___
- [ ] **Repository Pattern** - First applied in: ___
- [ ] **Hexagonal Architecture** - First applied in: ___
- [ ] **Dependency Injection** - First applied in: ___
- [ ] **Middleware Chain** - First applied in: ___
- [ ] **Event-Driven** - First applied in: ___
- [ ] **Producer-Consumer** - First applied in: ___

---

## Notes

- ⭐ = Critical topic, take extra time
- Topics with **Exercise** require hands-on completion
- Check `notes/progress.md` for session-by-session details
- Use `/show-progress` for visual dashboard
