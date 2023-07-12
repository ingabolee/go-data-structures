package main

import "fmt"

func insertionSort(array []int) {
	// Iterate through the array, starting from the second element.
	for i := 1; i < len(array); i++ {
		// The current element to be sorted.
		currentValue := array[i]

		// Iterate through the sorted portion of the array, starting from the end.
		j := i - 1
		for j >= 0 && array[j] > currentValue {
			// Shift the element at index `j` to the right.
			array[j+1] = array[j]
			j--
		}

		// Insert the current element at index `j+1`.
		array[j+1] = currentValue
	}
}

func main() {
	// Create an array of unsorted integers.
	array := []int{5, 2, 4, 7, 1, 3, 6, 8}

	// Sort the array using insertion sort.
	insertionSort(array)

	// Print the sorted array.
	fmt.Println(array)
}
