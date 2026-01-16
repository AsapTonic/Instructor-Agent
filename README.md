# Go Learning Journey

> A personalized Go learning environment powered by Claude Code
> 
> Combines: **Go Tour** + **Pragmatic Programmer** + **Software Architecture**

---

## 🚀 Quick Start

### 1. Prerequisites

- [VS Code](https://code.visualstudio.com/)
- [Claude Code Extension](https://marketplace.visualstudio.com/items?itemName=anthropic.claude-code)
- [Go](https://go.dev/dl/) (1.22+)
- [Node.js](https://nodejs.org/) (for MCP servers)
- Claude Pro, Max, or API key

### 2. Setup

```bash
# Clone or copy this directory to your machine
cp -r go-learning-journey ~/go-learning-journey

# Navigate to the project
cd ~/go-learning-journey

# Install MCP server dependencies (first time only)
npm install -g @modelcontextprotocol/server-filesystem
npm install -g @modelcontextprotocol/server-memory
npm install -g @modelcontextprotocol/server-fetch

# Open in VS Code
code .
```

### 3. Start Learning

Open the Claude Code panel (click the ✱ Spark icon) and type:

```
/next-lesson
```

---

## 📚 Available Commands

| Command | Description |
|---------|-------------|
| `/next-lesson` | Auto-select and start the next appropriate lesson |
| `/start-lesson [topic]` | Start a specific topic (or get suggestions if empty) |
| `/complete-lesson [topic]` | Verify mastery and mark a lesson complete |
| `/show-progress` | View your visual progress dashboard |
| `/generate-exercise [topic]` | Create targeted practice problems |
| `/review-code [file]` | Get Socratic code review feedback |

---

## 📁 Project Structure

```
go-learning-journey/
├── .claude/
│   ├── commands/           # Slash command definitions
│   │   ├── start-lesson.md
│   │   ├── next-lesson.md
│   │   ├── complete-lesson.md
│   │   ├── show-progress.md
│   │   ├── generate-exercise.md
│   │   └── review-code.md
│   └── subagents/
│       └── go-tutor.md     # Main teaching agent
├── lessons/                # Lesson content (generated)
├── exercises/              # Practice problems (generated)
├── projects/               # Larger capstone projects
├── notes/
│   ├── curriculum.md       # Learning path & checkboxes
│   ├── progress.md         # Session-by-session log
│   └── concepts.md         # Personal knowledge base
├── .mcp.json               # MCP server configuration
├── go.mod                  # Go module file
└── README.md               # This file
```

---

## 🎯 Learning Path Overview

### Phase 1: Foundations (1-2 weeks)
- Packages, Variables, Functions, Flow Control
- **Pragmatic Focus**: DRY, basic error handling

### Phase 2: Data Structures (2-3 weeks)
- Pointers, Arrays, Slices, Maps, Structs
- **Pragmatic Focus**: Orthogonality, data-driven design
- **Architecture**: Introduction to layered design

### Phase 3: Methods & Interfaces (2-3 weeks)
- Methods, Interfaces, Error Handling
- **Pragmatic Focus**: Design by Contract, Reversibility
- **Architecture**: Dependency Injection, Hexagonal

### Phase 4: Generics (1-2 weeks)
- Type Parameters, Constraints
- **Pragmatic Focus**: DRY at the type level

### Phase 5: Concurrency (3-4 weeks)
- Goroutines, Channels, Select, Sync, Context
- **Pragmatic Focus**: Fail Fast in concurrent code
- **Architecture**: Event-Driven patterns

### Phase 6: Real-World Patterns (3-4 weeks)
- HTTP, Middleware, Testing, Project Organization
- **All principles combined**
- **Architecture**: Clean Architecture implementation

---

## 🧠 Teaching Philosophy

The go-tutor agent follows the **Socratic method**:

- ❌ Never gives direct answers
- ✅ Asks guiding questions
- ✅ Provides progressive hints
- ✅ Celebrates discoveries
- ✅ Connects concepts to real-world applications

**Example interaction:**

```
You: How do I fix this error?

go-tutor: Let's work through this together! 
Looking at the error message, what does it tell you 
about what Go is expecting vs what it received?
```

---

## 📈 Tracking Progress

Your progress is tracked in two places:

1. **`notes/curriculum.md`** - Checkbox tracker for all topics
2. **`notes/progress.md`** - Detailed session-by-session log

Use `/show-progress` for a visual dashboard anytime!

---

## 🔧 Customization

### Adjusting Difficulty

Edit `.claude/subagents/go-tutor.md` to change:
- Hint levels
- Assessment criteria
- Teaching style

### Adding Resources

Add your own references to `notes/concepts.md` as you learn.

### Extending the Curriculum

Add new topics to `notes/curriculum.md` following the existing format.

---

## 💡 Tips for Success

1. **Consistency over intensity** - 30 min daily beats 3 hours weekly
2. **Type the code** - Don't copy-paste; muscle memory matters
3. **Struggle is learning** - Ask for hints, not answers
4. **Connect concepts** - Look for patterns across topics
5. **Build projects** - Apply what you learn to real problems

---

## 🐛 Troubleshooting

### MCP servers not connecting
```bash
# Reinstall MCP servers
npm install -g @modelcontextprotocol/server-filesystem
npm install -g @modelcontextprotocol/server-memory
```

### Commands not recognized
Make sure you're in the `go-learning-journey` directory when using Claude Code.

### Agent not using Socratic method
Remind it: "Remember to use the Socratic method - ask me questions instead of giving answers."

---

## 📚 Additional Resources

- [Go Tour](https://go.dev/tour/) - Official Go tutorial
- [Effective Go](https://go.dev/doc/effective_go) - Go best practices
- [Go by Example](https://gobyexample.com/) - Annotated examples
- [The Pragmatic Programmer](https://pragprog.com/titles/tpp20/) - Software craftsmanship

---

Happy learning! 🚀

*"The best time to plant a tree was 20 years ago. The second best time is now."*
