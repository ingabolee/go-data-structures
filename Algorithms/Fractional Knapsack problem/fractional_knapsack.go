package main

import (
	"fmt"
	"sort"
)

// Item represents an item with a weight and a value
type Item struct {
	weight, value float64
}

// FractionalKnapsack returns the maximum value that can be obtained
// by filling a knapsack of capacity W with fractional items
func FractionalKnapsack(W float64, items []Item) float64 {
	// Sort the items by their value per unit weight in decreasing order
	sort.Slice(items, func(i, j int) bool {
		return items[i].value/items[i].weight > items[j].value/items[j].weight
	})

	// Initialize the result and the remaining capacity
	res := 0.0
	cap := W

	// Iterate over the sorted items
	for _, item := range items {
		// If the item can fit entirely, take it and update the result and capacity
		if item.weight <= cap {
			res += item.value
			cap -= item.weight
		} else {
			// If the item cannot fit entirely, take a fraction of it and fill the knapsack
			res += item.value * cap / item.weight
			break // No more space left
		}
	}

	return res
}

func main() {
	items := []Item{{10, 60}, {20, 100}, {30, 120}}
	W := 50.0
	fmt.Println(FractionalKnapsack(W, items))
}
