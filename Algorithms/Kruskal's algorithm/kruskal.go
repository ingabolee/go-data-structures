// Kruskal's algorithm is a graph algorithm that finds the minimum spanning tree
// of a connected, undirected, weighted graph. A minimum spanning tree is a subset
// of the edges that connects all the nodes with the minimum total weight.

// It works by sorting all the edges of the graph by their weights and adding them
// to the tree one by one, as long as they do not create a cycle. It uses a disjoint-set
// data structure to keep track of the connected components of the tree and to check
// for cycles efficiently.

// We define a Node struct that represents a node in the graph
package main

import (
	"fmt"
	"sort"
)

type Node struct {
	// Value is the value of the node
	Value string
	// Parent is the parent node in the disjoint-set
	Parent *Node
	// Rank is the rank of the node in the disjoint-set
	Rank int
}

// We define an Edge struct that represents an edge in the graph
type Edge struct {
	// Source is the source node of the edge
	Source *Node
	// Destination is the destination node of the edge
	Destination *Node
	// Weight is the weight of the edge
	Weight int
}

// FindSet takes a node and returns its representative node in the disjoint-set
func FindSet(node *Node) *Node {
	// If the node is its own parent, it is the representative
	if node.Parent == node {
		return node
	}
	// Otherwise, recursively find the parent of the node and compress the path
	node.Parent = FindSet(node.Parent)
	return node.Parent
}

// Union takes two nodes and merges their sets in the disjoint-set
func Union(node1, node2 *Node) {
	// Find the representatives of each node
	root1 := FindSet(node1)
	root2 := FindSet(node2)

	// If they are already in the same set, do nothing
	if root1 == root2 {
		return
	}

	// Otherwise, compare their ranks and attach the lower rank to the higher rank
	if root1.Rank < root2.Rank {
		root1.Parent = root2
	} else if root1.Rank > root2.Rank {
		root2.Parent = root1
	} else {
		// If they have equal ranks, arbitrarily make one the parent of the other and increment its rank
		root2.Parent = root1
		root1.Rank++

	}
}

// Kruskal takes a graph as a slice of edges, and returns two values:
// one for the total weight of the minimum spanning tree,
// and one for the slice of edges in the minimum spanning tree.
func Kruskal(graph []*Edge) (int, []*Edge) {
	// Sort the edges by their weights in non-decreasing order
	sort.Slice(graph, func(i, j int) bool {
		return graph[i].Weight < graph[j].Weight
	})

	// Initialize an empty slice to store the edges in the minimum spanning tree
	mst := make([]*Edge, 0)

	// Initialize a variable to store the total weight of the minimum spanning tree
	totalWeight := 0

	// Loop through each edge in the sorted graph
	for _, edge := range graph {
		// Get the source and destination nodes of the edge
		u := edge.Source
		v := edge.Destination
		// Find their representatives in the disjoint-set
		uSet := FindSet(u)
		vSet := FindSet(v)
		// If they are not in the same set, it means they do not create a cycle
		if uSet != vSet {
			// Add them to different sets using union operation
			Union(uSet, vSet)
			// Add this edge to mst slice
			mst = append(mst, edge)
			// Add its weight to total weight
			totalWeight += edge.Weight
		}

	}

	// Return total weight and mst slice
	return totalWeight, mst

}

// main function to test Kruskal's algorithm
func main() {
	// Create some nodes
	a := &Node{Value: "A"}
	b := &Node{Value: "B"}
	c := &Node{Value: "C"}
	d := &Node{Value: "D"}
	e := &Node{Value: "E"}
	f := &Node{Value: "F"}

	// Initialize each node as its own parent and rank zero
	a.Parent = a
	b.Parent = b
	c.Parent = c
	d.Parent = d
	e.Parent = e
	f.Parent = f

	// Create some edges with weights
	ab := &Edge{Source: a, Destination: b, Weight: 5}
	ac := &Edge{Source: a, Destination: c, Weight: 1}
	bd := &Edge{Source: b, Destination: d, Weight: 1}
	be := &Edge{Source: b, Destination: e, Weight: 3}
	ce := &Edge{Source: c, Destination: e, Weight: 2}
	df := &Edge{Source: d, Destination: f, Weight: 2}
	ef := &Edge{Source: e, Destination: f, Weight: 3}

	// Create a graph with the edges
	graph := []*Edge{ab, ac, bd, be, ce, df, ef}

	// Run Kruskal's algorithm on the graph
	totalWeight, mst := Kruskal(graph)

	// Print the total weight of the minimum spanning tree
	fmt.Println("Total weight:", totalWeight)

	// Print the edges in the minimum spanning tree
	fmt.Println("Edges:")
	for _, edge := range mst {
		fmt.Printf("%s - %s\n", edge.Source.Value, edge.Destination.Value)
	}

}
