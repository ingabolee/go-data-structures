package main

import (
	"fmt"
)

// Define the size of the hash table
const tableSize = 100

// Define the structure of a node in the hash table
type Node struct {
	key   int
	value int
	next  *Node
}

// Define the structure of the hash table
type HashTable struct {
	table map[int]*Node
}

// Define the hash function that maps a key to an index in the hash table
func hashFunction(key int) int {
	return key % tableSize
}

// Define the Insert method that inserts a new key-value pair into the hash table
func (h *HashTable) Insert(key int, value int) {
	index := hashFunction(key)
	if h.table[index] == nil {
		h.table[index] = &Node{key, value, nil}
	} else {
		newNode := &Node{key, value, nil}
		currentNode := h.table[index]
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
}

// Define the Search method that searches for a key in the hash table and returns true if found, false otherwise
func (h *HashTable) Search(key int) bool {
	index := hashFunction(key)
	if h.table[index] == nil {
		return false
	} else {
		currentNode := h.table[index]
		for currentNode != nil && currentNode.key != key {
			currentNode = currentNode.next
		}
		if currentNode == nil {
			return false
		} else {
			return true
		}
	}
}

// Define the Delete method that deletes a key-value pair from the hash table given its key
func (h *HashTable) Delete(key int) {
	index := hashFunction(key)
	if h.table[index] == nil {
		fmt.Println("Key not found")
	} else if h.table[index].key == key && h.table[index].next == nil {
		h.table[index] = nil
	} else if h.table[index].key == key && h.table[index].next != nil {
		h.table[index] = h.table[index].next
	} else {
		previousNode := h.table[index]
		currentNode := previousNode.next
		for currentNode != nil && currentNode.key != key {
			previousNode = currentNode
			currentNode = currentNode.next
		}
		if currentNode == nil {
			fmt.Println("Key not found")
		} else {
			previousNode.next = currentNode.next
		}
	}
}

func main() {

	// Create a new instance of the HashTable struct with an empty map as its table field
	hashTable := &HashTable{table: make(map[int]*Node)}

	// Insert some key-value pairs into the hash table
	hashTable.Insert(1, 10)
	hashTable.Insert(2, 20)
	hashTable.Insert(3, 30)

	// Search for some keys in the hash table and print out whether they were found or not
	fmt.Println(hashTable.Search(1))
	fmt.Println(hashTable.Search(4))

	// Delete a key-value pair from the hash table given its key
	hashTable.Delete(1)

}
