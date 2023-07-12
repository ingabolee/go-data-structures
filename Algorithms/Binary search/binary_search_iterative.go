package main

import "fmt"

func binarySearch(list []int, target int) int {
	// Initialize the low and high indices of the list.
	low := 0
	high := len(list) - 1

	// Loop until the low index is greater than or equal to the high index.
	for low <= high {
		// Find the middle element of the list.
		mid := (low + high) / 2

		// If the target is equal to the middle element, return the index of the middle element.
		if list[mid] == target {
			return mid
		}

		// If the target is less than the middle element, update the high index to the middle element - 1.
		if list[mid] > target {
			high = mid - 1
		} else { // If the target is greater than the middle element, update the low index to the middle element + 1.
			low = mid + 1
		}
	}

	// If the target is not found, return -1.
	return -1
}

func main() {
	// Create a list of integers.
	list := []int{1, 3, 5, 7, 9}

	// Search for the target value in the list.
	target := 9
	index := binarySearch(list, target)

	// If the target is found, print the index of the target value.
	if index != -1 {
		fmt.Println("Target found at index:", index)
	} else {
		fmt.Println("Target not found.")
	}
}
