package main

import "fmt"

type Node struct {
	data interface{} // Store data of any type
	next *Node       // Pointer to the next node in the list
}

type LinkedList struct {
	head *Node // Pointer to the first node in the list
}

func (list *LinkedList) AddNode(data interface{}) {
	newNode := &Node{
		data: data,
		next: nil,
	}
	fmt.Println(newNode)
	if list.head == nil {
		// If the list is empty, make the new node the head
		list.head = newNode
	} else {
		// Traverse to the end of the list and add the new node
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}
func main() {
	list := LinkedList{}
	fmt.Println(list)
	// Add nodes to the linked list
	list.AddNode(10)
	list.AddNode("Hello")
	list.AddNode(3.14)
}
