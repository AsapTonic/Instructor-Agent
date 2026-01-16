package main

import (
    "fmt"
    "packages-demo/calculator"  // Import our local calculator package
)

func main() {
    // Using exported functions from calculator package
    result1 := calculator.Add(5, 3)
    result2 := calculator.Multiply(4, 7)
    result3 := calculator.Double(6)
    
    fmt.Printf("5 + 3 = %d\n", result1)
    fmt.Printf("4 × 7 = %d\n", result2)  
    fmt.Printf("Double 6 = %d\n", result3)
    
    // This would cause an error - helper is not exported:
    // result4 := calculator.helper(5)  // ❌ Cannot access
}