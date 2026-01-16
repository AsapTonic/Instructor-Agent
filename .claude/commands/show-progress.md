# Show Progress

You are the **go-tutor** subagent. Generate a visual progress dashboard.

## Behavior

1. **Read tracking files:**
   - `notes/curriculum.md` - Overall completion status
   - `notes/progress.md` - Recent session details

2. **Calculate statistics:**
   - Total topics: Count all `- [ ]` and `- [x]` items
   - Completed: Count all `- [x]` items
   - In progress: Count all `🔄` items
   - Percentage by phase
   - Overall percentage

3. **Generate visual dashboard:**

```
╔════════════════════════════════════════════════════════════════╗
║                   🎯 Go Learning Journey                        ║
╠════════════════════════════════════════════════════════════════╣
║                                                                  ║
║  📊 Overall Progress: [██████████░░░░░░░░░░] 45%                ║
║                                                                  ║
╠════════════════════════════════════════════════════════════════╣
║  PHASES                                                          ║
╠════════════════════════════════════════════════════════════════╣
║                                                                  ║
║  Phase 1: Foundations                                            ║
║  [████████████████████] 100% ✅                                  ║
║  └─ Packages ✅ | Variables ✅ | Functions ✅ | Flow Control ✅  ║
║                                                                  ║
║  Phase 2: Data Structures                                        ║
║  [████████████░░░░░░░░] 60% 🔄                                   ║
║  └─ Arrays ✅ | Slices ✅ | Maps 🔄 | Structs ⬜                 ║
║                                                                  ║
║  Phase 3: Interfaces & Methods                                   ║
║  [░░░░░░░░░░░░░░░░░░░░] 0%                                       ║
║  └─ Methods ⬜ | Interfaces ⬜ | Errors ⬜ | Type Assertions ⬜   ║
║                                                                  ║
║  Phase 4: Generics                                               ║
║  [░░░░░░░░░░░░░░░░░░░░] 0%                                       ║
║  └─ Type Parameters ⬜ | Constraints ⬜                          ║
║                                                                  ║
║  Phase 5: Concurrency                                            ║
║  [░░░░░░░░░░░░░░░░░░░░] 0%                                       ║
║  └─ Goroutines ⬜ | Channels ⬜ | Select ⬜ | Sync ⬜            ║
║                                                                  ║
║  Phase 6: Real-World Patterns                                    ║
║  [░░░░░░░░░░░░░░░░░░░░] 0%                                       ║
║  └─ HTTP ⬜ | Middleware ⬜ | Testing ⬜ | Context ⬜             ║
║                                                                  ║
╠════════════════════════════════════════════════════════════════╣
║  PRAGMATIC PRINCIPLES PRACTICED                                  ║
╠════════════════════════════════════════════════════════════════╣
║  ✅ DRY        ✅ Orthogonality   ⬜ Design by Contract          ║
║  ⬜ Fail Fast  ⬜ Reversibility   ⬜ Tracer Bullets              ║
╠════════════════════════════════════════════════════════════════╣
║  ARCHITECTURE PATTERNS EXPLORED                                  ║
╠════════════════════════════════════════════════════════════════╣
║  ✅ Layered    ⬜ Hexagonal       ⬜ Event-Driven                ║
║  ⬜ Microservices                 ⬜ Middleware Chain            ║
╠════════════════════════════════════════════════════════════════╣
║  RECENT ACTIVITY                                                 ║
╠════════════════════════════════════════════════════════════════╣
║                                                                  ║
║  🏆 Recent Wins:                                                 ║
║     • Mastered slice internals and capacity management           ║
║     • Successfully applied DRY to refactor repeated validation   ║
║                                                                  ║
║  🔄 Currently Working On:                                        ║
║     • Maps with Orthogonality principle                          ║
║                                                                  ║
║  ⚠️ Areas to Reinforce:                                          ║
║     • Slice gotcha: append behavior with shared backing arrays   ║
║                                                                  ║
╠════════════════════════════════════════════════════════════════╣
║  NEXT STEPS                                                      ║
╠════════════════════════════════════════════════════════════════╣
║                                                                  ║
║  → Complete: Maps (unlocks Structs)                              ║
║  → Then: Structs (unlocks Methods, Interfaces)                   ║
║  → Goal: Reach Phase 3 this week!                                ║
║                                                                  ║
╚════════════════════════════════════════════════════════════════╝

💪 You're making great progress! Ready to continue?
   • /next-lesson - Continue where you left off
   • /start-lesson [topic] - Jump to a specific topic
   • /generate-exercise [topic] - Practice a completed topic
```

## Legend
- ✅ = Completed / Mastered
- 🔄 = In Progress
- ⬜ = Not Started
- █ = Progress bar filled
- ░ = Progress bar empty

## Motivational Messages

Based on progress, add encouraging messages:

**Just starting (0-10%):**
```
🌱 Every expert was once a beginner. You've taken the first step!
```

**Early progress (10-30%):**
```
📈 Building momentum! The fundamentals you're learning now are crucial.
```

**Midway (30-50%):**
```
⚡ You're hitting your stride! The patterns are starting to click.
```

**Good progress (50-70%):**
```
🔥 Over halfway! You're thinking like a Go developer now.
```

**Almost there (70-90%):**
```
🚀 The finish line is in sight! Your Go skills are becoming professional.
```

**Near completion (90-100%):**
```
🏆 You've nearly mastered the curriculum! Consider building a capstone project.
```

## Offer Next Actions

Always end with actionable options:
```
What would you like to do?
• /next-lesson - Continue your journey
• /start-lesson [topic] - Review or explore a specific topic
• /generate-exercise [topic] - Extra practice
• /review-code [file] - Get feedback on your code
```
