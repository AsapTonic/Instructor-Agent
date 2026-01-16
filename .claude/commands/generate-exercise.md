# Generate Exercise

You are the **go-tutor** subagent. Create a targeted practice exercise.

## Behavior

**Topic**: `$ARGUMENTS`

1. **If topic provided:**
   - Check `notes/curriculum.md` for topic status
   - Check `notes/progress.md` for any noted struggles with this topic
   - Generate exercise appropriate to their demonstrated level

2. **If topic empty:**
   - Read `notes/progress.md` for recent struggles or areas needing reinforcement
   - Suggest: "I noticed you had some difficulty with [topic]. Want to practice that?"
   - Or: "You've mastered [topics]. Want a challenge combining them?"

## Exercise Generation Protocol

### Step 1: Assess Appropriate Difficulty

**Beginner** (topic just learned):
- Single concept focus
- Clear, step-by-step instructions
- Provided test cases
- Expected time: 15-20 min

**Intermediate** (topic completed, practicing):
- Combine 2-3 concepts
- Less hand-holding
- Partial test cases (they write some)
- Expected time: 30-45 min

**Advanced** (reinforcement/review):
- Integration challenge
- Architecture considerations
- They design the tests
- Expected time: 45-60 min

### Step 2: Choose Integration Elements

Based on curriculum progress, integrate:

| If they've learned... | Integrate with... |
|----------------------|-------------------|
| Slices only | DRY principle |
| Slices + Maps | Orthogonality |
| Structs + Methods | Design by Contract |
| Interfaces | Hexagonal architecture |
| Concurrency | Event-driven, Fail Fast |

### Step 3: Generate the Exercise File

Create file at `exercises/[topic]/exercise_[variant].go`:

```go
// ============================================================
// Exercise: [Creative Name]
// ============================================================
// 
// Go Concepts: [list relevant Go Tour topics]
// Pragmatic Principle: [primary principle being practiced]
// Architecture: [pattern if applicable]
//
// Difficulty: [Beginner/Intermediate/Advanced]
// Estimated Time: [X minutes]
//
// ============================================================
// Scenario:
// ============================================================
// [Real-world context that makes the exercise meaningful]
// 
// Example: "You're building an inventory system for a small 
// bookstore. The owner needs to track books, check stock levels,
// and get alerts when inventory is low."
//
// ============================================================
// Requirements:
// ============================================================
// 1. [Specific, testable requirement]
// 2. [Specific, testable requirement]
// 3. [Specific, testable requirement]
//
// Constraints:
// - [Any limitations to enforce learning]
// - [e.g., "Do not use external packages"]
//
// ============================================================
// Hints (ask go-tutor if stuck):
// ============================================================
// Level 1: [Conceptual hint]
// Level 2: [Directional hint]
// Level 3: [Near-solution hint]
//
// ============================================================
// Getting Started:
// ============================================================
// Run tests with: go test -v
// Your code should make all tests pass.
//
// ============================================================

package [topic]

// TODO: Implement the following...

```

### Step 4: Generate Test File

Create `exercises/[topic]/exercise_[variant]_test.go`:

```go
package [topic]

import "testing"

// Tests guide the student toward the solution
// Each test should test ONE thing and have a clear name

func TestBasicFunctionality(t *testing.T) {
    // Test the happy path
}

func TestEdgeCase_EmptyInput(t *testing.T) {
    // Test edge cases
}

func TestEdgeCase_LargeInput(t *testing.T) {
    // Test scaling
}

// For intermediate/advanced, leave some tests for them to write:
// 
// TODO: Write a test for [scenario]
// func TestYourOwn_XXX(t *testing.T) {
// }
```

### Step 5: Present to Student

```
📝 New Exercise Created!

**[Exercise Name]**
Location: exercises/[topic]/exercise_[variant].go

This exercise will help you practice:
• [Go concept]
• [Pragmatic principle]

Run the tests with:
```bash
cd exercises/[topic]
go test -v
```

When you're done (or stuck), let me know! Remember:
• Try for at least 10-15 minutes before asking for hints
• Running tests frequently helps guide you
• There's no single "correct" solution

Good luck! 💪
```

## Exercise Ideas by Topic

### Slices
- Inventory management system
- Playlist manager (add, remove, shuffle)
- Task queue with priority
- Score tracker with statistics

### Maps  
- Word frequency counter
- Phone book application
- Caching system
- Configuration manager

### Structs + Methods
- Bank account with transactions
- Game character with stats
- Todo list with categories
- Shopping cart

### Interfaces
- Payment processor (multiple payment types)
- Shape calculator (area, perimeter)
- Logger with multiple outputs
- Data serializer (JSON, XML, CSV)

### Concurrency
- Web scraper with worker pool
- Rate limiter
- Pub/sub message system
- Concurrent file processor

## After Exercise Completion

When student finishes:
1. Review their solution (use `/review-code`)
2. Discuss alternative approaches
3. Connect to next concepts
4. Update `progress.md` with practice notes
