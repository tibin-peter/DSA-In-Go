// Stack
package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}
type Stackll struct {
	top *Node
}

func (s *Stackll) push(value int) {
	newNode := &Node{data: value}
	newNode.next = s.top
	s.top = newNode
}
func (s *Stackll) pop() {
	if s.top == nil {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Printf("Popped value: %d\n", s.top.data)
	s.top = s.top.next
}
func (s *Stackll) peek() {
	if s.top == nil {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Println("Value that top on the Stack:", s.top.data)
}
func (s *Stackll) display() {
	current := s.top
	if current == nil {
		fmt.Println("Stack is empty.")
		return
	}
	for current != nil {
		fmt.Printf("%d ----> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	var value, num, choice int
	s := Stackll{}

	for {
		fmt.Println("Select your choice")
		fmt.Println("1.Add value")
		fmt.Println("2.Show the value in the top of the stack ")
		fmt.Println("3.Delete the last added value")
		fmt.Println("4.Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			{
				fmt.Print("Enter the limit: ")
				fmt.Scan(&num)
				for i := 0; i < num; i++ {
					fmt.Printf("Enter the value %d: ", i+1)
					fmt.Scan(&value)
					s.push(value)
				}

			}
		case 2:
			{
				s.display()

			}
		case 3:
			{
				s.pop()
				fmt.Println("Stack after deletion")
				s.display()

			}
		case 4:
			{
				fmt.Println("Exiting....")
				return

			}
		default:
			{
				fmt.Println("Enter a valid choice")
			}
		}
	}
}
