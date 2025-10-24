// Queue

package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type QueueLL struct {
	front *Node
	rear  *Node
}

func (q *QueueLL) enqueue(value int) {
	newNode := &Node{data: value}

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}

}

func (q *QueueLL) dequeue() {
	if q.front == nil {
		fmt.Println("Queue is empty")
		return
	}

	fmt.Printf("Dequeued value: %d\n", q.front.data)
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}
}

func (q *QueueLL) peek() {
	if q.front == nil {
		fmt.Println("Queue is empty.")
		return
	}
	fmt.Printf("Front value: %d\n", q.front.data)
}

func (q *QueueLL) display() {
	if q.front == nil {
		fmt.Println("Queue is empty.")
		return
	}
	current := q.front
	for current != nil {
		fmt.Printf("%d ----> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	var choice, value, num int
	q := QueueLL{}

	for {
		fmt.Println("1. Enqueue values")
		fmt.Println("2. Peek front value")
		fmt.Println("3. Dequeue value")
		fmt.Println("4. Display queue")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter how many values you want to insert: ")
			fmt.Scan(&num)
			for i := 0; i < num; i++ {
				fmt.Printf("Enter value %d: ", i+1)
				fmt.Scan(&value)
				q.enqueue(value)
			}
		case 2:
			q.peek()
		case 3:
			q.dequeue()
		case 4:
			q.display()
		case 5:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid choice. choose currect one.")
		}
	}
}
