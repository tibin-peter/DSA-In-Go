package main

import (
	"fmt"
)

func find(arr []int) (min, max int) {
	min = arr[0]
	max = arr[0]

	for _, v := range arr {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	return min, max

}

func main() {

	arr := []int{2, 4, 65, 1, 7, 46}
	min, max := find(arr)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
