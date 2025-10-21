package main

import (
	"fmt"
)

// create Node
type Node struct {
	data int
	next *Node
}

// creating Stacklinkedlist
type StackLL struct {
	top *Node
}

// func for add in to the top(push)
func (s *StackLL) push(value int) {
	newNode := &Node{data: value}
	newNode.next = s.top
	s.top = newNode

}

// func for remove at the top node (pop)
func (s *StackLL) pop() {
	if s.top == nil {
		fmt.Println("empty")
		return
	}
	fmt.Println("Removed:", s.top.data)
	s.top = s.top.next
}

// func for find the top node in stack(peek)
func (s *StackLL) peek() {
	if s.top == nil {
		fmt.Println("Stack is empty")
	}
	fmt.Println("Last Added:", s.top.data)
}

// func for display
func (s *StackLL) display() {
	current := s.top
	for current != nil {
		fmt.Printf("%d-->", current.data)
		current = current.next
	}
	fmt.Println("NULL")
}

func main() {
	myStacklist := StackLL{}
	myStacklist.push(10)
	myStacklist.push(20)
	myStacklist.push(30)
	myStacklist.display() // 30,20,10
	myStacklist.peek()    // 30
	myStacklist.pop()     // remove 30
	myStacklist.display()
}
