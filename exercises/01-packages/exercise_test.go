package main

import (
	"strings"
	"testing"
)

// Note: These tests are designed to guide your implementation.
// Run them with: go test -v

func TestFormalGreeting(t *testing.T) {
	// Skip if greetings package not yet created
	// Once you create the package, remove this skip
	t.Skip("Remove this skip once you've created the greetings package")
	
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "lowercase name",
			input:    "alice",
			expected: "Good day, Alice. Welcome to our establishment.",
		},
		{
			name:     "uppercase name",
			input:    "BOB",
			expected: "Good day, Bob. Welcome to our establishment.",
		},
		{
			name:     "mixed case name",
			input:    "cHaRLie",
			expected: "Good day, Charlie. Welcome to our establishment.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: Uncomment and update once greetings package exists
			// got := greetings.Formal(tt.input)
			// if got != tt.expected {
			//     t.Errorf("Formal(%q) = %q; want %q", tt.input, got, tt.expected)
			// }
		})
	}
}

func TestCasualGreeting(t *testing.T) {
	t.Skip("Remove this skip once you've created the greetings package")
	
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "lowercase name",
			input:    "diana",
			expected: "Hey Diana! What's up?",
		},
		{
			name:     "uppercase name",
			input:    "EVE",
			expected: "Hey Eve! What's up?",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: Uncomment once greetings package exists
			// got := greetings.Casual(tt.input)
			// if got != tt.expected {
			//     t.Errorf("Casual(%q) = %q; want %q", tt.input, got, tt.expected)
			// }
		})
	}
}

func TestCapitalizeIsUnexported(t *testing.T) {
	// This test verifies that 'capitalize' is NOT exported
	// If your greetings package accidentally exports it, this conceptual
	// test reminds you to fix it.
	
	// The way to verify: try to call greetings.capitalize() 
	// It should NOT compile if properly unexported.
	
	// This is a conceptual test - the real test is:
	// Does `greetings.capitalize` cause a compile error? (It should!)
	
	t.Log("Remember: capitalize should start with lowercase 'c' to be unexported")
	t.Log("Verify by trying: greetings.capitalize('test') - should not compile")
}

func TestGreetingsUseDRY(t *testing.T) {
	// This is a conceptual test for the DRY principle
	// We're checking that both functions properly capitalize names
	// without duplicating the capitalization logic
	
	t.Skip("Remove this skip once you've created the greetings package")
	
	// Both should handle the same edge case identically
	// because they share the capitalize helper
	
	weirdName := "jOhN"
	
	// TODO: Uncomment once greetings package exists
	// formal := greetings.Formal(weirdName)
	// casual := greetings.Casual(weirdName)
	
	// Both should have "John" properly capitalized
	// if !strings.Contains(formal, "John") {
	//     t.Error("Formal greeting didn't properly capitalize name")
	// }
	// if !strings.Contains(casual, "John") {
	//     t.Error("Casual greeting didn't properly capitalize name")
	// }
	
	_ = strings.Contains // suppress unused import warning during skip
}

// ============================================================
// Bonus: Write Your Own Test
// ============================================================
// Once you've completed the exercise, try writing a test for:
// - Empty string input
// - Very long names
// - Names with numbers
//
// func TestYourOwn(t *testing.T) {
//     // Your test here
// }
