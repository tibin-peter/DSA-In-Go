package dsatoolkit

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) Display() {
	fmt.Println(s.items)
}
