package main

import "fmt"

func bubbleSort(array []int) {
	// Loop through the array, comparing adjacent elements.
	for i := 0; i < len(array)-1; i++ {
		// For each iteration, start at the beginning of the array and compare adjacent elements.
		for j := 0; j < len(array)-i-1; j++ {
			// If the current element is greater than the next element, swap them.
			if array[j] > array[j+1] {
				// Swap the elements.
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
}

func main() {
	// Create an array of unsorted numbers.
	array := []int{5, 2, 4, 7, 1, 3, 6, 9, 8}

	// Print the unsorted array.
	fmt.Println("Unsorted array:", array)

	// Sort the array using bubble sort.
	bubbleSort(array)

	// Print the sorted array.
	fmt.Println("Sorted array:", array)
}
