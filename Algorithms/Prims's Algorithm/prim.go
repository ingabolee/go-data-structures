// Prim's algorithm is a graph algorithm that finds the minimum spanning tree
// of a connected, undirected, weighted graph. A minimum spanning tree is a subset
// of the edges that connects all the nodes with the minimum total weight.

// It works by starting from an arbitrary node and growing the tree by adding
// the edge with the minimum weight that connects a node in the tree to a node outside the tree.
// It repeats this process until all the nodes are in the tree.

// We define a Node struct that represents a node in the graph

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
	// Weights is a map of nodes and their weights from the source
	Weights map[*Node]int
}

// Len returns the length of the queue
func (pq *PriorityQueue) Len() int {
	return len(pq.Nodes)
}

// Less returns true if the node at i has less weight than the node at j
func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.Weights[pq.Nodes[i]] < pq.Weights[pq.Nodes[j]]
}

// Swap swaps the nodes at i and j
func (pq *PriorityQueue) Swap(i, j int) {
	pq.Nodes[i], pq.Nodes[j] = pq.Nodes[j], pq.Nodes[i]
}

// Push adds a node to the queue
func (pq *PriorityQueue) Push(x interface{}) {
	pq.Nodes = append(pq.Nodes, x.(*Node))
}

// Pop removes and returns the node with the minimum weight
func (pq *PriorityQueue) Pop() interface{} {
	n := len(pq.Nodes)
	node := pq.Nodes[n-1]
	pq.Nodes = pq.Nodes[:n-1]
	return node
}

// We define a constant for infinity
const Infinity = int(^uint(0) >> 1)

// Prim takes a source node and a graph, and returns two values:
// one for the total weight of the minimum spanning tree,
// and one for the slice of edges in the minimum spanning tree.
func Prim(source *Node, graph []*Node) (int, [][2]*Node) {
	// Initialize an empty priority queue
	pq := &PriorityQueue{
		Nodes:   make([]*Node, 0),
		Weights: make(map[*Node]int),
	}
	// Initialize the weights map with infinity for each node
	weights := make(map[*Node]int)
	for _, node := range graph {
		weights[node] = Infinity
	}
	// Initialize an empty visited map to keep track of visited nodes
	visited := make(map[*Node]bool)

	// Set the weight of the source to zero and push it to the queue
	weights[source] = 0
	heap.Push(pq, source)

	// Initialize an empty slice to store the edges in the minimum spanning tree
	mst := make([][2]*Node, 0)

	// Initialize a variable to store the total weight of the minimum spanning tree
	totalWeight := 0

	// Loop until the queue is empty or all nodes are visited
	for pq.Len() > 0 && len(visited) < len(graph) {
		// Pop the node with the minimum weight from the queue
		u := heap.Pop(pq).(*Node)
		// Mark it as visited and add its weight to the total weight
		visited[u] = true
		totalWeight += weights[u]
		// Loop through its adjacent nodes
		for v, w := range u.Edges {
			// If v is not visited and its weight is larger than w,
			if !visited[v] && weights[v] > w {
				// Update its weight and push it to the queue
				weights[v] = w
				heap.Push(pq, v)
				// Add the edge (u,v) to the mst slice
				mst = append(mst, [2]*Node{u, v})
			}
		}

	}

	// Return the total weight and mst slice
	return totalWeight, mst

}

// main function to test Prim's algorithm
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

	// Run Prim's algorithm from node A
	totalWeight, mst := Prim(a, graph)

	// Print the total weight of the minimum spanning tree
	fmt.Println("Total weight:", totalWeight)

	// Print the edges in the minimum spanning tree
	fmt.Println("Edges:")
	for _, edge := range mst {
		fmt.Printf("%s - %s\n", edge[0].Value, edge[1].Value)
	}

}
