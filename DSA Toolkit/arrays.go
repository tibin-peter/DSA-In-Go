package dsatoolkit

import "fmt"

type Array struct {
	elements []int
}

func NewArray() *Array {
	return &Array{elements: []int{}}
}

func (a *Array) Insert(value int) {
	a.elements = append(a.elements, value)
}

func (a *Array) Delete(index int) {
	if index < 0 || index >= len(a.elements) {
		fmt.Println("Index out of range")
		return
	}
	a.elements = append(a.elements[:index], a.elements[index+1:]...)
}

func (a *Array) Search(value int) int {
	for i, v := range a.elements {
		if v == value {
			return i
		}
	}
	return -1
}

func (a *Array) Reverse() {
	for i, j := 0, len(a.elements)-1; i < j; i, j = i+1, j-1 {
		a.elements[i], a.elements[j] = a.elements[j], a.elements[i]
	}
}

func (a *Array) Display() {
	fmt.Println(a.elements)
}
