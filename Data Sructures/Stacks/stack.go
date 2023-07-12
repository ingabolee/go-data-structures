package main

import "fmt"

// Define a struct to represent the stack
type Stack struct {
	items []int // slice to store the items in the stack
}

// Push method adds an item to the top of the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item) // add item to the end of the slice
}

// Pop method removes and returns the top item from the stack
func (s *Stack) Pop() int {
	if len(s.items) == 0 { // check if the stack is empty
		return -1 // return -1 to indicate an error
	}
	top := s.items[len(s.items)-1]     // get the top item from the end of the slice
	s.items = s.items[:len(s.items)-1] // remove the top item from the slice
	return top                         // return the top item
}

// Peek method returns the top item from the stack without removing it
func (s *Stack) Peek() int {
	if len(s.items) == 0 { // check if the stack is empty
		return -1 // return -1 to indicate an error
	}
	return s.items[len(s.items)-1] // return the top item from the end of the slice
}

// Size method returns the number of items in the stack
func (s *Stack) Size() int {
	return len(s.items) // return the length of the slice
}

func main() {
	// Create a new stack
	stack := Stack{}

	// Push some items onto the stack
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Pop the top item from the stack and print it
	fmt.Println(stack.Pop()) // Output: 3

	// Peek at the top item in the stack and print it
	fmt.Println(stack.Peek()) // Output: 2

	// Print the number of items in the stack
	fmt.Println(stack.Size()) // Output: 2
}
