package main

import "fmt"

func selectionSort(array []int) {
	// Loop through the array, starting at the first element.
	for i := 0; i < len(array)-1; i++ {
		// Set the minimum element to the current element.
		minIndex := i

		// Loop through the remaining elements in the array.
		for j := i + 1; j < len(array); j++ {
			// If the current element is less than the minimum element,
			// update the minimum element.
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}

		// Swap the current element with the minimum element.
		array[i], array[minIndex] = array[minIndex], array[i]
	}
}

func main() {
	// Create an array of integers.
	array := []int{10, 5, 2, 7, 3, 1, 8, 9, 6}

	// Print the array before sorting.
	fmt.Println("Unsorted array:", array)

	// Sort the array.
	selectionSort(array)

	// Print the array after sorting.
	fmt.Println("Sorted array:", array)
}
