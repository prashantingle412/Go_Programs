package main

import (
	"fmt"
)

func main() {
	var list = []int{12, 6, 9, 18, 33, 2, 99}
	for i := 0; i < len(list); i++ {
		swap(list)
	}

	fmt.Println("Hello, playground", list)
}

func swap(list []int) {
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
		}
	}
}
