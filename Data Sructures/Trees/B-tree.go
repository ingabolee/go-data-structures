package main

import (
	"fmt"
)

const t = 2 // degree of the B-tree

type node struct {
	isLeaf bool
	keys   []int
	child  []*node
	n      int
}

func newnode(leaf bool) *node {
	return &node{
		isLeaf: leaf,
		keys:   make([]int, 2*t-1),
		child:  make([]*node, 2*t),
		n:      0,
	}
}

func (n *node) search(k int) (bool, int) {
	i := 0
	for i < n.n && k > n.keys[i] {
		i++
	}
	if i < n.n && k == n.keys[i] {
		return true, i
	}
	if n.isLeaf {
		return false, -1
	}
	return n.child[i].search(k)
}

func (n *node) insertNonFull(k int) {
	i := n.n - 1
	if n.isLeaf {
		for i >= 0 && n.keys[i] > k {
			n.keys[i+1] = n.keys[i]
			i--
		}
		n.keys[i+1] = k
		n.n++
	} else {
		for i >= 0 && n.keys[i] > k {
			i--
		}
		if n.child[i+1].n == 2*t-1 {
			n.splitChild(i+1, n.child[i+1])
			if n.keys[i+1] < k {
				i++
			}
		}
		n.child[i+1].insertNonFull(k)
	}
}

func (n *node) splitChild(i int, y *node) {
	z := newnode(y.isLeaf)
	z.n = t - 1
	for j := 0; j < t-1; j++ {
		z.keys[j] = y.keys[j+t]
	}
	if !y.isLeaf {
		for j := 0; j < t; j++ {
			z.child[j] = y.child[j+t]
		}
	}
	y.n = t - 1
	for j := n.n; j >= i+1; j-- {
		n.child[j+1] = n.child[j]
	}
	n.child[i+1] = z
	for j := n.n - 1; j >= i; j-- {
		n.keys[j+1] = n.keys[j]
	}
	n.keys[i] = y.keys[t-1]
	n.n++
}

type Btree struct {
	root *node
}

func (b *Btree) search(k int) bool {
	if b.root == nil {
		return false
	}
	found, _ := b.root.search(k)
	return found
}

func (b *Btree) insert(k int) {
	if b.root == nil {
		b.root = newnode(true)
		b.root.keys[0] = k
		b.root.n = 1
	} else {
		if b.root.n == 2*t-1 {
			s := newnode(false)
			s.child[0] = b.root
			s.splitChild(0, b.root)
			i := 0
			if s.keys[0] < k {
				i++
			}
			s.child[i].insertNonFull(k)
			b.root = s
		} else {
			b.root.insertNonFull(k)
		}
	}
}

func main() {
	b := &Btree{}
	b.insert(10)
	b.insert(20)
	b.insert(5)
	b.insert(6)
	b.insert(12)
	b.insert(30)
	b.insert(7)
	fmt.Println(b.search(6)) // true
	fmt.Println(b.search(9)) // false
}
