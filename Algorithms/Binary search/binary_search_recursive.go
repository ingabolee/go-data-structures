package main

import "fmt"

func binarySearch(list []int, target int, low int, high int) int {
	// Base case: If the list is empty or the target is not in the list, return -1.
	if len(list) == 0 || target < list[0] || target > list[high] {
		return -1
	}

	// Find the middle element of the list.
	mid := (low + high) / 2

	// If the target is equal to the middle element, return the index of the middle element.
	if list[mid] == target {
		return mid
	}

	// If the target is less than the middle element, recursively search the left half of the list.
	if list[mid] > target {
		return binarySearch(list, target, low, mid-1)
	}

	// If the target is greater than the middle element, recursively search the right half of the list.
	return binarySearch(list, target, mid+1, high)
}

func main() {
	// Create a list of integers.
	list := []int{1, 3, 5, 7, 9}

	// Search for the target value in the list.
	target := 9
	index := binarySearch(list, target, 0, len(list)-1)

	// If the target is found, print the index of the target value.
	if index != -1 {
		fmt.Println("Target found at index:", index)
	} else {
		fmt.Println("Target not found.")
	}
}
