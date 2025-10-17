package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 5, 3, 65, 23, 6, 4}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum)
}
