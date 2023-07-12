package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	// Value is the value of the node
	Value string
	// Edges is a map of adjacent nodes and their edge weights
	Edges map[*Node]int
}

// We define a PriorityQueue struct that implements heap.Interface
type PriorityQueue struct {
	// Nodes is a slice of nodes in the queue
	Nodes []*Node
	// Distances is a map of nodes and their distances from the source
	Distances map[*Node]int
}

// Len returns the length of the queue
func (pq *PriorityQueue) Len() int {
	return len(pq.Nodes)
}

// Less returns true if the node at i has less distance than the node at j
func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.Distances[pq.Nodes[i]] < pq.Distances[pq.Nodes[j]]
}

// Swap swaps the nodes at i and j
func (pq *PriorityQueue) Swap(i, j int) {
	pq.Nodes[i], pq.Nodes[j] = pq.Nodes[j], pq.Nodes[i]
}

// Push adds a node to the queue
func (pq *PriorityQueue) Push(x interface{}) {
	pq.Nodes = append(pq.Nodes, x.(*Node))
}

// Pop removes and returns the node with the minimum distance
func (pq *PriorityQueue) Pop() interface{} {
	n := len(pq.Nodes)
	node := pq.Nodes[n-1]
	pq.Nodes = pq.Nodes[:n-1]
	return node
}

// We define a constant for infinity
const Infinity = int(^uint(0) >> 1)

// Dijkstra takes a source node and a graph, and returns two maps:
// one for the shortest distances from the source to each node,
// and one for the previous node on the shortest path to each node.
func Dijkstra(source *Node, graph []*Node) (map[*Node]int, map[*Node]*Node) {
	// Initialize an empty priority queue
	pq := &PriorityQueue{
		Nodes:     make([]*Node, 0),
		Distances: make(map[*Node]int),
	}
	// Initialize the distances map with infinity for each node
	distances := make(map[*Node]int)
	for _, node := range graph {
		distances[node] = Infinity
	}
	// Initialize the previous map with nil for each node
	previous := make(map[*Node]*Node)
	for _, node := range graph {
		previous[node] = nil
	}
	// Set the distance of the source to zero and push it to the queue
	distances[source] = 0
	heap.Push(pq, source)

	// Loop until the queue is empty
	for pq.Len() > 0 {
		// Pop the node with the minimum distance
		u := heap.Pop(pq).(*Node)
		// Loop through its adjacent nodes
		for v, w := range u.Edges {
			// Calculate the distance to v through u
			alt := distances[u] + w
			// If this distance is smaller than the current distance to v,
			if alt < distances[v] {
				// Update the distance to v and its previous node
				distances[v] = alt
				previous[v] = u
				// Push v to the queue with its new distance
				heap.Push(pq, v)
			}
		}
	}

	// Return the distances and previous maps
	return distances, previous

}

// PrintPath takes a target node and a previous map, and prints the shortest path to it from the source.
func PrintPath(target *Node, previous map[*Node]*Node) {
	// Initialize an empty slice to store the path nodes
	path := make([]*Node, 0)
	// Start from the target node and loop backwards until reaching nil (the source)
	for node := target; node != nil; node = previous[node] {
		// Prepend the node to the path slice
		path = append([]*Node{node}, path...)
	}
	// Print each node in the path slice separated by arrows
	for i, node := range path {
		if i > 0 {
			fmt.Print(" -> ")
		}
		fmt.Print(node.Value)
	}
	fmt.Println()
}

// main function to test the Dijkstra's algorithm
func main() {
	// Create some nodes
	a := &Node{Value: "A", Edges: make(map[*Node]int)}
	b := &Node{Value: "B", Edges: make(map[*Node]int)}
	c := &Node{Value: "C", Edges: make(map[*Node]int)}
	d := &Node{Value: "D", Edges: make(map[*Node]int)}
	e := &Node{Value: "E", Edges: make(map[*Node]int)}
	f := &Node{Value: "F", Edges: make(map[*Node]int)}

	// Create some edges with weights
	a.Edges[b] = 5
	a.Edges[c] = 1
	b.Edges[a] = 5
	b.Edges[d] = 1
	b.Edges[e] = 3
	c.Edges[a] = 1
	c.Edges[e] = 2
	d.Edges[b] = 1
	d.Edges[f] = 2
	e.Edges[b] = 3
	e.Edges[c] = 2
	e.Edges[f] = 3
	f.Edges[d] = 2
	f.Edges[e] = 3

	// Create a graph with the nodes
	graph := []*Node{a, b, c, d, e, f}

	// Run Dijkstra's algorithm from node A
	distances, previous := Dijkstra(a, graph)

	// Print the shortest distances from A to each node
	for node, distance := range distances {
		fmt.Printf("Distance from A to %s: %d\n", node.Value, distance)
		// Print the shortest path from A to each node
		fmt.Printf("Path from A to %s: ", node.Value)
		PrintPath(node, previous)
		fmt.Println()

	}

}
