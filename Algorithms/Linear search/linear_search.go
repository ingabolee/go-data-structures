package main

import "fmt"

func linearSearch(array []int, x int) int {
	// Initialize the index of the element to be -1.
	index := -1

	// Iterate through the array.
	for i := 0; i < len(array); i++ {
		// If the element is found, update the index and break the loop.
		if array[i] == x {
			index = i
			break
		}
	}

	// Return the index of the element.
	return index
}

func main() {
	// Create an array of integers.
	array := []int{1, 2, 3, 4, 5}

	// Find the index of the element 3.
	index := linearSearch(array, 3)

	// Print the index of the element.
	fmt.Println("The index of 3 is", index)
}
