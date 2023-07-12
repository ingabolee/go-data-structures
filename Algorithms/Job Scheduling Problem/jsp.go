package main

import (
	"fmt"
	"sort"
)

// Job represents a job with a deadline and a profit
type Job struct {
	id     string
	dead   int
	profit int
}

// compare is a helper function to sort the jobs by their profit in decreasing order
func compare(a, b Job) bool {
	return a.profit > b.profit
}

// printJobScheduling is the main function that prints the maximum profit sequence of jobs
func printJobScheduling(jobs []Job, n int) {
	// Sort the jobs by their profit in decreasing order
	sort.Slice(jobs, func(i, j int) bool {
		return compare(jobs[i], jobs[j])
	})

	// Initialize an array of slots to store the scheduled jobs
	slots := make([]string, n)

	// Iterate over the sorted jobs
	for _, job := range jobs {
		// Find the latest possible slot that does not exceed the deadline
		for j := min(n, job.dead) - 1; j >= 0; j-- {
			// If the slot is empty, assign the job to it and break
			if slots[j] == "" {
				slots[j] = job.id
				break
			}
		}
	}

	// Print the scheduled jobs
	fmt.Println("Following is maximum profit sequence of jobs:")
	for _, slot := range slots {
		if slot != "" {
			fmt.Print(slot, " ")
		}
	}
	fmt.Println()
}

// min is a helper function to return the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	jobs := []Job{
		{"a", 2, 100},
		{"b", 1, 19},
		{"c", 2, 27},
		{"d", 1, 25},
		{"e", 3, 15},
	}
	n := len(jobs)
	printJobScheduling(jobs, n)
}
