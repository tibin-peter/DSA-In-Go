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

// func for insert a value at head
func (list *LinkedList) InsertAtHead(value int) {
	newNode := &Node{data: value}
	newNode.next = list.head
	list.head = newNode
}

// func for insert value at tail
func (list *LinkedList) InsertAtTail(value int) {
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

// func for search a value
func (list *LinkedList) Search(value int) bool {
	current := list.head
	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}
	return false
}

// func for delete a value
func (list *LinkedList) DeleteByValue(value int) {
	if list.head == nil {
		fmt.Printf("List is empty")
		return
	}
	if list.head.data == value {
		list.head = list.head.next
		return
	}
	current := list.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}
	if current.next == nil {
		fmt.Println("value not found")
	}
	current.next = current.next.next
}

// func for delete a value using its position
func (list *LinkedList) DeleteByPosition(pos int) {
	if list.head == nil {
		fmt.Println("List is empty.")
		return
	}

	if pos == 0 {
		list.head = list.head.next
		return
	}

	current := list.head
	for i := 0; current != nil && i < pos-1; i++ {
		current = current.next
	}

	if current == nil || current.next == nil {
		fmt.Println("Position out of range.")
		return
	}

	current.next = current.next.next
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
	// creating a verible and assignig the linkedlist
	list := LinkedList{}

	// Insert elements
	list.InsertAtHead(10)
	list.InsertAtTail(20)
	list.InsertAtTail(30)
	list.InsertAtTail(40)
	list.InsertAtHead(5)

	fmt.Println("Linked List:")
	list.Print()

	// Search example
	fmt.Println("Search 30:", list.Search(30))
	fmt.Println("Search 100:", list.Search(100))

	// Delete by value
	list.DeleteByValue(30)
	fmt.Println("After deleting 30:")
	list.Print()

	// Delete by position
	list.DeleteByPosition(2)
	fmt.Println("After deleting position 2:")
	list.Print()
}
