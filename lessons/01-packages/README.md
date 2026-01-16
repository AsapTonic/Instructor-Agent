# Lesson 1: Packages & Imports

> **Go Concept**: Packages, Imports, Exported Names
> **Pragmatic Principle**: DRY (Don't Repeat Yourself)
> **Estimated Time**: 30-45 minutes

---

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
1. Create and organize Go packages
2. Import and use standard library packages
3. Understand exported vs unexported names
4. Apply DRY by organizing reusable code into packages

---

## 📖 Concept Introduction

### What is a Package?

Think of a package like a folder in your desk drawer. Just as you might have separate folders for "Bills", "Recipes", and "Taxes", Go organizes code into packages based on what it does.

Every Go file starts with a package declaration:

```go
package main  // or package mypackage
```

### The `main` Package

The `main` package is special - it's the entry point for executable programs:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

**Key points:**
- Must have `package main`
- Must have `func main()`
- This is where execution begins

### Imports

To use code from other packages, you import them:

```go
import "fmt"                    // Single import
import "math/rand"              // Subpackage

import (                        // Multiple imports (grouped)
    "fmt"
    "math"
)
```

### Exported Names

In Go, a name is **exported** (public) if it starts with a capital letter:

```go
package mymath

func Add(a, b int) int {        // Exported - other packages CAN use this
    return a + b
}

func subtract(a, b int) int {   // Unexported - other packages CANNOT use this
    return a - b
}
```

This is Go's approach to encapsulation - no `public`/`private` keywords needed!

---

## 🔗 Pragmatic Connection: DRY

**The Pragmatic Programmer** teaches us:
> "Every piece of knowledge must have a single, unambiguous, authoritative representation within a system."

### How Packages Enable DRY

Without packages, you might repeat code across files:

```go
// file1.go
func calculateTax(amount float64) float64 {
    return amount * 0.08
}

// file2.go (DUPLICATED!)
func calculateTax(amount float64) float64 {
    return amount * 0.08
}
```

With packages, you write it once:

```go
// tax/tax.go
package tax

func Calculate(amount float64) float64 {
    return amount * 0.08
}

// file1.go and file2.go both import "tax"
import "myproject/tax"
tax.Calculate(100.00)
```

**Benefits:**
- One place to update if tax rate changes
- Clear ownership of the code
- Easier to test

---

## 🧪 Practice Exercise

Now it's your turn! Complete the exercise in `exercises/01-packages/`.

The exercise will have you:
1. Create a multi-file Go program
2. Organize code into packages
3. Use exported/unexported names correctly
4. Apply DRY by avoiding code duplication

---

## 📝 Quick Reference

| Concept | Example | Notes |
|---------|---------|-------|
| Package declaration | `package main` | First line of every Go file |
| Single import | `import "fmt"` | Standard library |
| Grouped imports | `import ("fmt"  "math")` | Multiple packages |
| Exported name | `Calculate()` | Capital letter = public |
| Unexported name | `calculate()` | Lowercase = private |
| Main function | `func main()` | Entry point |

---

## ✅ Mastery Checklist

Before moving on, make sure you can:
- [ ] Explain what a package is in your own words
- [ ] Create a file with proper package declaration
- [ ] Import and use standard library packages
- [ ] Predict which names are exported vs unexported
- [ ] Explain how packages help with DRY

---

## 🔜 Next Up

**Lesson 2: Variables & Types** - We'll explore Go's type system and see how type safety helps prevent bugs (connecting to the Pragmatic principle of "Fail Fast").
