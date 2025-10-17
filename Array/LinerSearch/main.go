package main

import (
	"fmt"
)

func search(arr []int, tar int) bool {
	for _, v := range arr {
		if v == tar {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{1, 4, 54, 7, 3, 7, 8}
	target := 7
	output := search(arr, target)
	fmt.Println("Result:", output)
}
