package main

import "fmt"

// Node represents an element in a singly linked list.
type Node struct {
	Data int
	Next *Node
}

// PrintData is a method with a pointer receiver that can be called on a nil receiver.
func (n *Node) PrintData() {
	if n == nil {
		fmt.Println("Node is nil.")
		return
	}
	fmt.Println("Node data:", n.Data)
}

// AddNode adds a new node to the linked list. Returns the new node.
func (n *Node) AddNode(data int) *Node {
	newNode := &Node{Data: data}
	if n != nil {
		n.Next = newNode
	}
	return newNode
}

func main() {
	var head *Node // nil initially

	// Call PrintData on a nil pointer (this is valid in Go).
	head.PrintData() // Expected Output: Node is nil.

	// Create the first node in the list.
	head = head.AddNode(10)
	head.PrintData() // Expected Output: Node data: 10

	// Add another node.
	head.Next = head.AddNode(20)
	head.PrintData() // Expected Output: Node data: 10

	// Print the next node (which is still not nil).
	head.Next.PrintData() // Expected Output: Node data: 20

	// Try calling PrintData on a nil pointer (next node is nil).
	head.Next.Next.PrintData() // Expected Output: Node is nil.

	//here head.Next.Next is nil but still we are able to call reciever PrintData on it
	//see here u can call a method on a nil reciever and u can handle the edge case in the reciever
	//this is not possible in java/c++ like languages
}
