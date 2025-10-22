package dsatoolkit

import "fmt"

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) Insert(value int) {
	newNode := &Node{Data: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}
	l.Tail.Next = newNode
	newNode.Prev = l.Tail
	l.Tail = newNode
}

func (l *LinkedList) Delete(value int) {
	current := l.Head
	for current != nil {
		if current.Data == value {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				l.Head = current.Next
			}
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				l.Tail = current.Prev
			}
			return
		}
		current = current.Next
	}
}

func (l *LinkedList) FromArray(arr []int) {
	for _, v := range arr {
		l.Insert(v)
	}
}

func (l *LinkedList) Display() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}
