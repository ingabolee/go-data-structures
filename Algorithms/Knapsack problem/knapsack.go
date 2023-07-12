// Knapsack problem is an optimization problem where we have to choose a subset of items
// that maximizes the total value while keeping the total weight within a given capacity.
// There are different variants of the problem, such as 0-1 knapsack, fractional knapsack, etc.

// We define an Item struct that represents an item in the knapsack
package main

import "fmt"

type Item struct {
	// Value is the value of the item
	Value int
	// Weight is the weight of the item
	Weight int
}

// Knapsack01 takes a slice of items, a capacity, and returns the maximum value and the subset of items
// that can be packed in the knapsack using 0-1 knapsack algorithm.
// 0-1 knapsack means that each item can be either taken or left, but not split.
func Knapsack01(items []Item, capacity int) (int, []Item) {
	// Initialize a 2D slice to store the optimal values for each subproblem
	// dp[i][w] represents the maximum value that can be obtained by using the first i items
	// and a knapsack of capacity w
	n := len(items)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// Loop through each item
	for i := 1; i <= n; i++ {
		// Loop through each possible weight from 0 to capacity
		for w := 0; w <= capacity; w++ {
			// Get the current item's value and weight
			v := items[i-1].Value
			wt := items[i-1].Weight
			// If the current item's weight is larger than the current capacity,
			if wt > w {
				// We cannot take this item, so we just inherit the previous optimal value
				dp[i][w] = dp[i-1][w]
			} else {
				// Otherwise, we have two options: take this item or leave it
				// If we take this item, we add its value to the optimal value of the remaining capacity
				take := v + dp[i-1][w-wt]
				// If we leave this item, we just inherit the previous optimal value
				leave := dp[i-1][w]
				// We choose the option that gives us the maximum value
				dp[i][w] = max(take, leave)
			}
		}

	}

	// The maximum value is stored at the bottom-right corner of the dp slice
	maxValue := dp[n][capacity]

	// To find the subset of items that gives us the maximum value,
	// we backtrack from the bottom-right corner and check which items were taken or left
	subset := make([]Item, 0)
	i := n
	w := capacity
	for i > 0 && w > 0 {
		// If the optimal value at (i,w) is different from the optimal value at (i-1,w),
		// it means that we took the ith item
		if dp[i][w] != dp[i-1][w] {
			// Add this item to the subset slice
			subset = append(subset, items[i-1])
			// Reduce the remaining capacity by this item's weight
			w -= items[i-1].Weight

		}
		// Move to the previous row
		i--

	}

	// Return the maximum value and the subset slice
	return maxValue, subset

}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// main function to test the knapsack problem algorithm
func main() {
	// Create some items with values and weights
	items := []Item{
		{Value: 60, Weight: 10},
		{Value: 100, Weight: 20},
		{Value: 120, Weight: 30},
	}

	// Set a capacity for the knapsack
	capacity := 50

	// Run Knapsack01 on the items and capacity
	maxValue, subset := Knapsack01(items, capacity)

	// Print the maximum value and the subset of items
	fmt.Println("Maximum value:", maxValue)
	fmt.Println("Subset of items:")
	for _, item := range subset {
		fmt.Printf("Value: %d, Weight: %d\n", item.Value, item.Weight)
	}

}
