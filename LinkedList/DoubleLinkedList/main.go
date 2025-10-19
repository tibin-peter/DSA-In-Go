package main

import (
	"fmt"
)

// create node struct for both next and previous
type Node struct {
	data int
	next *Node
	prev *Node
}

// create struct head and tail for starting and ending
type LinkedList struct {
	head *Node
	tail *Node
}

// func for insert value both head and end
func (list *LinkedList) Insert(value int) {
	newNode := &Node{data: value} // create a newnode with value

	//if nill make the newnode at the head and also the end
	if list.head == nil {
		list.head = newNode
		list.tail = newNode

		return
	}

	list.tail.next = newNode // backword
	newNode.prev = list.tail
	list.tail = newNode
}

// func for printforward
func (list *LinkedList) Printforward() {
	current := list.head

	for current != nil {
		fmt.Printf("%d-->", current.data)
		current = current.next
	}
	fmt.Print("Null\n")
}

// func for Printbackword

func (list *LinkedList) Printbackward() {
	current := list.tail

	for current != nil {
		fmt.Printf("%d-->", current.data)
		current = current.prev
	}
	fmt.Print("Null\n")
}

func main() {
	list := LinkedList{}
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(40)

	list.Printforward()
	list.Printbackward()
}
