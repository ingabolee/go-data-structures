package main

import (
	"fmt"
)

// Node represents a node in the graph with an ID.
type Node struct {
	ID int
}

// Edge represents an edge in the graph with a source, target, and weight.
type Edge struct {
	Source, Target Node
	Weight         float64
}

// Graph represents a graph with nodes and edges.
type Graph struct {
	Nodes    map[int]Node
	Edges    map[int][]Edge
	Directed bool
	Weighted bool
}

// NewGraph creates a new graph with the given directed and weighted properties.
func NewGraph(directed, weighted bool) Graph {
	return Graph{
		Nodes:    make(map[int]Node),
		Edges:    make(map[int][]Edge),
		Directed: directed,
		Weighted: weighted,
	}
}

// AddNode adds a node with the given ID to the graph.
func (g Graph) AddNode(id int) {
	if _, exists := g.Nodes[id]; !exists {
		g.Nodes[id] = Node{ID: id}
		g.Edges[id] = []Edge{}
	}
}

// AddEdge adds an edge between the source and target nodes with the given weight.
func (g Graph) AddEdge(source, target int, weight float64) {
	if g.Weighted == false {
		weight = 1
	}

	// Check if both nodes exist, if not, add them to the graph.
	g.AddNode(source)
	g.AddNode(target)

	// Add the edge between the source and target nodes.
	g.Edges[source] = append(g.Edges[source], Edge{
		Source: g.Nodes[source],
		Target: g.Nodes[target],
		Weight: weight,
	})

	// If the graph is undirected, add the reverse edge as well.
	if !g.Directed {
		g.Edges[target] = append(g.Edges[target], Edge{
			Source: g.Nodes[target],
			Target: g.Nodes[source],
			Weight: weight,
		})
	}
}

func main() {
	// Create a directed, unweighted graph.
	directedUnweighted := NewGraph(true, false)
	directedUnweighted.AddEdge(1, 2, 0)
	directedUnweighted.AddEdge(2, 3, 0)
	fmt.Println("Directed, Unweighted Graph:", directedUnweighted)

	// Create an undirected, unweighted graph.
	undirectedUnweighted := NewGraph(false, false)
	undirectedUnweighted.AddEdge(1, 2, 0)
	undirectedUnweighted.AddEdge(2, 3, 0)
	fmt.Println("Undirected, Unweighted Graph:", undirectedUnweighted)

	// Create a directed, weighted graph.
	directedWeighted := NewGraph(true, true)
	directedWeighted.AddEdge(1, 2, 2.5)
	directedWeighted.AddEdge(2, 3, 1.5)
	fmt.Println("Directed, Weighted Graph:", directedWeighted)

	// Create an undirected, weighted graph.
	undirectedWeighted := NewGraph(false, true)
	undirectedWeighted.AddEdge(1, 2, 2.5)
	undirectedWeighted.AddEdge(2, 3, 1.5)
	fmt.Println("Undirected, Weighted Graph:", undirectedWeighted)
}
