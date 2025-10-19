package main

import "fmt"

func count(arr []int, m int) int {
	var count int
	for _, v := range arr {
		if v == m {
			count++
		}
	}
	return count
}

func main() {
	arr := []int{1, 2, 3, 0, 0, 0}
	m := 0
	r := count(arr, m)

	fmt.Println("Result:", r)
}
