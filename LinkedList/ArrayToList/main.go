package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

func (list *LinkedList) arraytolist(arr []int) {
	for _, v := range arr {
		newNode := &Node{data: v}
		if list.head == nil {
			list.head = newNode
		} else {
			current := list.head
			for current.next != nil {
				current = current.next
			}
			current.next = newNode
		}
	}
}
func (list *LinkedList) Print() {
	current := list.head
	for current != nil {
		fmt.Printf("%d-->", current.data)
		current = current.next
	}
	fmt.Println("NULL")
}

func main() {
	arr := []int{5, 3, 8, 2, 9}
	list := LinkedList{}
	list.arraytolist(arr)
	list.Print()

}
