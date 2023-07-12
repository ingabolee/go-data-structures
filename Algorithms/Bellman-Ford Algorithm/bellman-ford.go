package main

import "fmt"

// The Bellman-Ford algorithm is a graph algorithm that finds the shortest path
// between a source node and all other nodes in the graph, even if the edges have negative weights.
// The algorithm works by relaxing each edge of the graph for n - 1 times,
// where n is the number of nodes. Relaxing an edge means updating the distance
// of the destination node if it can be improved by going through the source node.
// After n - 1 iterations, the shortest distances are guaranteed to be found,
// unless there is a negative cycle in the graph.
// A negative cycle is a cycle of edges that has a negative total weight.
// This means that traversing the cycle repeatedly can reduce the distance indefinitely.
// The algorithm can detect a negative cycle by performing one more iteration and checking
// if any distance can be improved further. If so, there is a negative cycle and the shortest
// path problem is not well-defined.
// The algorithm defines a Node struct that represents a node in the graph. The struct
// contains the node's ID, its distance from the source node, and a list of its neighbors.

type Node struct {
	// Value is the value of the node
	Value string
	// Edges is a slice of adjacent nodes and their edge weights
	Edges [][2]interface{}
}

// We define a constant for infinity
const Infinity = int(^uint(0) >> 1)

// BellmanFord takes a source node and a graph, and returns two values:
// one for the shortest distances from the source to each node,
// and one for whether there is a negative cycle in the graph or not.
func BellmanFord(source *Node, graph []*Node) (map[*Node]int, bool) {
	// Initialize the distances map with infinity for each node
	distances := make(map[*Node]int)
	for _, node := range graph {
		distances[node] = Infinity
	}
	// Set the distance of the source to zero
	distances[source] = 0

	// Loop for n-1 times, where n is the number of nodes
	n := len(graph)
	for i := 0; i < n-1; i++ {
		// Loop through each node in the graph
		for _, u := range graph {
			// Loop through each edge from u
			for _, edge := range u.Edges {
				// Get the destination node and the edge weight
				v := edge[0].(*Node)
				w := edge[1].(int)
				// Relax the edge by updating the distance to v if it can be improved
				if distances[u]+w < distances[v] {
					distances[v] = distances[u] + w
				}
			}
		}
	}

	// Loop through each node in the graph again
	for _, u := range graph {
		// Loop through each edge from u again
		for _, edge := range u.Edges {
			// Get the destination node and the edge weight again
			v := edge[0].(*Node)
			w := edge[1].(int)
			// Check if we can still improve the distance to v
			if distances[u]+w < distances[v] {
				// If so, there is a negative cycle in the graph
				return nil, true
			}
		}
	}

	// Return the distances map and false to indicate no negative cycle
	return distances, false

}

// main function to test the Bellman-Ford algorithm
func main() {
	// Create some nodes
	a := &Node{Value: "A", Edges: make([][2]interface{}, 0)}
	b := &Node{Value: "B", Edges: make([][2]interface{}, 0)}
	c := &Node{Value: "C", Edges: make([][2]interface{}, 0)}
	d := &Node{Value: "D", Edges: make([][2]interface{}, 0)}

	// Create some edges with weights (positive or negative)
	a.Edges = append(a.Edges, [2]interface{}{b, -1})
	a.Edges = append(a.Edges, [2]interface{}{c, 4})
	b.Edges = append(b.Edges, [2]interface{}{c, 3})
	b.Edges = append(b.Edges, [2]interface{}{d, 2})
	b.Edges = append(b.Edges, [2]interface{}{a, 1})
	d.Edges = append(d.Edges, [2]interface{}{b, 3})

	// Create a graph with the nodes
	graph := []*Node{a, b, c, d}

	// Run Bellman-Ford algorithm from node A
	distances, negativeCycle := BellmanFord(a, graph)

	// Check if there is a negative cycle in the graph
	if negativeCycle {
		// If so, print a message and exit
		fmt.Println("The graph has a negative cycle.")
		return
	}

	// Print the shortest distances from A to each node
	for node, distance := range distances {
		fmt.Printf("Distance from A to %s: %d\n", node.Value, distance)
	}

}
