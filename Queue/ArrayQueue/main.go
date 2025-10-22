package main

import (
	"fmt"
)

// creating a struct for Queue
type Queue struct {
	items []int
}

// func for Enqueue add an element in to end
func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

// func for Dequeue remove the first added element
func (q *Queue) Dequeue() {
	if len(q.items) == 0 {
		fmt.Println("Empty")
		return
	}
	fmt.Println("Dequeued:", q.items[0])
	q.items = q.items[1:]
}

// function for display the queue
func (q *Queue) Display() {
	if q.items == nil {
		fmt.Println("Queue is empty")
	}
	fmt.Println("Queue:", q.items)
}
func main() {
	q := Queue{}

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	q.Display() // [10,20,30]
	q.Dequeue() //10
	q.Display() //[20,30]
}
