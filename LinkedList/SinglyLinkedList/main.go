package main

import (
	"fmt"
)

// create node struct
type Node struct {
	data int
	next *Node
}

// create head struct for a starting point
type LinkedList struct {
	head *Node
}

// func for insert value
func (list *LinkedList) Insert(value int) {
	newNode := &Node{data: value} // create a newnode with value

	//if nill make the newnode at the head
	if list.head == nil {
		list.head = newNode
		return
	}

	// oterwise loop till end
	current := list.head
	for current.next != nil {
		current = current.next
	}
	// assign newnode at the end
	current.next = newNode
}

// func for print
func (list *LinkedList) Print() {
	current := list.head

	for current != nil {
		fmt.Printf("%d-->", current.data)
		current = current.next
	}
	fmt.Print("Null")
}

func main() {
	list := LinkedList{}
	list.Insert(10)
	list.Insert(20)

	list.Print()
}
