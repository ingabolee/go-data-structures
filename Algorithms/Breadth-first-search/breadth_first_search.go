package main

import "fmt"

func breadthFirstSearch(graph map[string][]string, startVertex string) {
	// Initialize a visited set to track which nodes have been visited.
	visited := make(map[string]bool)

	// Create a queue to store the nodes that need to be visited.
	queue := []string{startVertex}

	// Loop until the queue is empty.
	for len(queue) > 0 {
		// Pop the first node from the queue.
		node := queue[0]
		queue = queue[1:]

		// If the node has not been visited, mark it as visited and print it.
		if !visited[node] {
			visited[node] = true
			fmt.Println(node)

			// Add all of the node's neighbors to the queue.
			for _, neighbor := range graph[node] {
				queue = append(queue, neighbor)
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

	// Perform a breadth first search on the graph.
	breadthFirstSearch(graph, "A")
}
