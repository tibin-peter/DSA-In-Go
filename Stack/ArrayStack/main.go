package main

import (
	"fmt"
)

// creating a struct called stack contain a array
type Stack struct {
	item []int
}

// func for push at the top(push)
func (s *Stack) Push(value int) {
	s.item = append(s.item, value)
}

// remove at the top of the stck array(pop)
func (s *Stack) Pop() {
	if len(s.item) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Println("Poped:", s.item[len(s.item)-1])
	s.item = s.item[:len(s.item)-1]
}

// func to see the top of the stack(peek)
func (s *Stack) Peek() {
	if len(s.item) == 0 {
		fmt.Println("Stack is empty")
	}
	fmt.Println("Top:", s.item[len(s.item)-1])
}

// func for display the array
func (s *Stack) Display() {
	fmt.Println("Stack:", s.item)
}

func main() {
	myStack := Stack{}
	myStack.Push(10)
	myStack.Push(20)
	myStack.Push(30)
	myStack.Push(40)

	fmt.Println(" The stack After Push")
	myStack.Display()

	fmt.Println(" The stack After Pop")
	myStack.Pop()
	myStack.Display()

	myStack.Peek()
}
