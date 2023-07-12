package main

import "fmt"

func depthFirstSearch(graph map[string][]string, startVertex string) {
	// Initialize a visited set to track which nodes have been visited.
	visited := make(map[string]bool)

	// Create a stack to store the nodes that need to be visited.
	stack := []string{startVertex}

	// Loop until the stack is empty.
	for len(stack) > 0 {
		// Pop the top node from the stack.
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// If the node has not been visited, mark it as visited and print it.
		if !visited[node] {
			visited[node] = true
			fmt.Println(node)

			// Add all of the node's neighbors to the stack.
			for _, neighbor := range graph[node] {
				stack = append(stack, neighbor)
			}
		}
	}
}

func main() {
	// Create a graph.
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"C", "D"},
		"C": {"D"},
		"D": {},
	}

	// Perform a depth first search on the graph.
	depthFirstSearch(graph, "A")
}
