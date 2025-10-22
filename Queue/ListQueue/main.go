package main

import (
	"fmt"
)

// struct for node
type Node struct {
	data int
	next *Node
	prev *Node
}

// struct for queuelinkedlist
type QueueLL struct {
	head *Node
	tail *Node
}

// func for add an element in to queue
func (q *QueueLL) Enqueue(value int) {
	newNode := &Node{data: value}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
		return
	}
	q.tail.next = newNode
	newNode.prev = q.tail
	q.tail = newNode
}

// func for remove the first added element
func (q *QueueLL) Dequeue() {
	if q.head == nil {
		fmt.Println("empty")
	}
	fmt.Println("Removed :", q.head.data)
	q.head = q.head.next
}

// func for display elements
func (q *QueueLL) Display() {
	if q.head == nil {
		fmt.Println("empty")
		return
	}
	current := q.head
	for current != nil {
		fmt.Printf("%d---->", current.data)
		current = current.next
	}
}

func main() {
	q := QueueLL{}

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	q.Display() // 10,20,30
	q.Dequeue() // 10
	q.Display() //20,30
}
