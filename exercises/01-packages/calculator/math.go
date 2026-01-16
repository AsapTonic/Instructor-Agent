// Package calculator provides basic arithmetic operations
package calculator

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Multiply returns the product of two integers
func Multiply(a, b int) int {
	return a * b
}

// helper is unexported - only this package can use it
func helper(x int) int {
	return x * 2
}

// Double uses the helper function internally
func Double(x int) int {
	return helper(x)
}
