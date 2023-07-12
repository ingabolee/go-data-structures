package main

import "fmt"

func Fibonacci(n int) int {
	// Base case: if n is 0 or 1, return n
	if n == 0 || n == 1 {
		return n
	}
	// Recursive case: return the sum of the previous two Fibonacci numbers
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// main function to test the Fibonacci algorithm
func main() {
	// Print the first 10 Fibonacci numbers
	for i := 0; i < 15; i++ {
		fmt.Println(Fibonacci(i))
	}
}
