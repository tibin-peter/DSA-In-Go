package main

import (
	"fmt"
)

func reverse(arr []int) {
	start := 0
	end := len(arr) - 1
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Before Reversing:", arr)
	reverse(arr)
	fmt.Println("After Reversing:", arr)
}
