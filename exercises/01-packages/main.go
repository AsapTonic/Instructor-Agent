// ============================================================
// Exercise: Package Power
// ============================================================
//
// Go Concepts: Packages, Imports, Exported Names
// Pragmatic Principle: DRY (Don't Repeat Yourself)
//
// Difficulty: Beginner
// Estimated Time: 20-30 minutes
//
// ============================================================
// Scenario:
// ============================================================
// You're building a simple greeting application for a company.
// The company has three departments, and each needs to greet
// visitors in a slightly different way, but they all share
// some common greeting logic.
//
// Currently, the greeting code is duplicated. Your job is to
// organize it properly using packages and the DRY principle.
//
// ============================================================
// Requirements:
// ============================================================
// 1. Create a `greetings` package with:
//    - An exported function `Formal(name string) string`
//      Returns: "Good day, [name]. Welcome to our establishment."
//    - An exported function `Casual(name string) string`
//      Returns: "Hey [name]! What's up?"
//    - An unexported helper function `capitalize(s string) string`
//      (used internally to ensure names are capitalized)
//
// 2. The main package should:
//    - Import your greetings package
//    - Use both greeting functions
//    - NOT duplicate any greeting logic
//
// 3. All names should be properly capitalized using your helper
//
// ============================================================
// Hints (ask go-tutor for these progressively):
// ============================================================
// Level 1: Think about what code is shared between Formal and Casual
// Level 2: The capitalize function should be unexported - why?
// Level 3: Look at the strings package for help with capitalization
//
// ============================================================
// Project Structure:
// ============================================================
// exercises/01-packages/
// ├── main.go           (this file - complete the TODOs)
// ├── greetings/
// │   └── greetings.go  (create this file)
// └── exercise_test.go  (tests to verify your solution)
//
// ============================================================
// Getting Started:
// ============================================================
// 1. Create the greetings/ directory
// 2. Create greetings/greetings.go with the package declaration
// 3. Implement the three functions
// 4. Complete main.go below
// 5. Run: go run . (from exercises/01-packages/)
// 6. Run: go test -v (to verify)
//
// ============================================================

package main

import (
	"fmt"
	// TODO: Import your greetings package
	// Hint: The import path will be based on your module name
	// Example: "github.com/yourusername/go-learning/exercises/01-packages/greetings"
)

func main() {
	// TODO: Use the greetings package to greet these visitors
	
	// Sales department uses formal greetings
	salesVisitor := "alice"
	// Should print: "Good day, Alice. Welcome to our establishment."
	fmt.Println("Sales greeting:", /* TODO: call Formal */)
	
	// Tech department uses casual greetings  
	techVisitor := "bob"
	// Should print: "Hey Bob! What's up?"
	fmt.Println("Tech greeting:", /* TODO: call Casual */)
	
	// Marketing uses formal for external, casual for internal
	externalGuest := "charlie"
	internalColleague := "diana"
	fmt.Println("External:", /* TODO */)
	fmt.Println("Internal:", /* TODO */)
}

// ============================================================
// Questions to Consider:
// ============================================================
// 1. Why is capitalize unexported? What would happen if another
//    package tried to use it?
//
// 2. How does organizing code into packages support DRY?
//
// 3. What if we needed to add a third greeting style? How would
//    the package structure help?
//
// ============================================================
