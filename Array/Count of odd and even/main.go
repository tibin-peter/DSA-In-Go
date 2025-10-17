package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 4, 54, 7, 3, 7, 8}
	var odd int
	var even int
	for _, v := range arr {
		if v%2 == 1 {
			odd++
		} else {
			even++
		}
	}
	fmt.Println("Odd:", odd)
	fmt.Println("Even:", even)
}
