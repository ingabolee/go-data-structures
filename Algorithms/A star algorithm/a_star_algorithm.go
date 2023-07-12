package main

import (
	"container/heap"
	"fmt"
	"math"
)

type node struct {
	x int
	y int
}

type edge struct {
	from *node
	to   *node
	cost float64
}

type priorityQueue []*node

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].fScore() < pq[j].fScore()
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*node))
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type graph struct {
	nodes []*node
	edges []*edge
}

func (g *graph) addNode(n *node) {
	g.nodes = append(g.nodes, n)
}

func (g *graph) addEdge(from, to *node, cost float64) {
	g.edges = append(g.edges, &edge{from, to, cost})
}

func (g *graph) neighbors(n *node) []*node {
	var neighbors []*node
	for _, e := range g.edges {
		if e.from == n {
			neighbors = append(neighbors, e.to)
		}
	}
	return neighbors
}

func (g *graph) cost(from, to *node) float64 {
	for _, e := range g.edges {
		if e.from == from && e.to == to {
			return e.cost
		}
	}
	return math.Inf(1)
}

func (n *node) gScore() float64 {
	return 0 // start node has zero cost
}

func (n *node) hScore(goal *node) float64 {
	dx := math.Abs(float64(n.x - goal.x))
	dy := math.Abs(float64(n.y - goal.y))
	return math.Sqrt(dx*dx + dy*dy) // Euclidean distance heuristic
}

func (n *node) fScore() float64 {
	return n.gScore() + n.hScore(goal)
}

var start = &node{0, 0}
var goal = &node{4, 4}

func main() {
	g := &graph{}
	g.addNode(&node{0, 0})
	g.addNode(&node{0, 1})
	g.addNode(&node{0, 2})
	g.addNode(&node{1, 0})
	g.addNode(&node{1, 1})
	g.addNode(&node{1, 2})
	g.addNode(&node{2, 0})
	g.addNode(&node{2, 1})
	g.addNode(&node{2, 2})
	g.addEdge(&node{0, 0}, &node{0, 1}, 1)
	g.addEdge(&node{0, 1}, &node{0, 2}, 1)
	g.addEdge(&node{0, 2}, &node{1, 2}, 1)
	g.addEdge(&node{1, 2}, &node{2, 2}, 1)
	g.addEdge(&node{2, 2}, &node{2, 1}, 1)
	g.addEdge(&node{2, 1}, &node{2, 0}, 1)
	g.addEdge(&node{2, 0}, &node{1, 0}, 1)
	g.addEdge(&node{1, 0}, &node{0, 0}, 1)

	openSet := &priorityQueue{start}
	cameFrom := make(map[*node]*node)
	gScore := make(map[*node]float64)
	gScore[start] = 0
	fScore := make(map[*node]float64)
	fScore[start] = start.hScore(goal)

	for len(*openSet) > 0 {
		current := heap.Pop(openSet).(*node)
		if current == goal {
			path := []node{*current}
			for {
				current = cameFrom[current]
				if current == start {
					path = append(path, *current)
					break
				}
				path = append(path, *current)
			}
			for i := len(path) - 1; i >= 0; i-- {
				fmt.Printf("(%d, %d) ", path[i].x, path[i].y)
			}
			return
		}
		for _, neighbor := range g.neighbors(current) {
			tentativeGScore := gScore[current] + g.cost(current, neighbor)
			if tentativeGScore < gScore[neighbor] || gScore[neighbor] == 0 {
				cameFrom[neighbor] = current
				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = neighbor.hScore(goal) + tentativeGScore
				if !contains(openSet, neighbor) {
					heap.Push(openSet, neighbor)
				}
			}
		}
	}
	fmt.Println("No path found")
}

func contains(pq *priorityQueue, n *node) bool {
	for _, item := range *pq {
		if item == n {
			return true
		}
	}
	return false
}
