package main

import (
	"fmt"
	"time"
)

type user struct {
	name string
}

func main() {
	var a = []int{10, 20, 30, 40, 50}
	for _, v := range a {
		go func() {
			show(v)
		}()
	}
	time.Sleep(5 * time.Second)
}
func show(v int) {
	time.Sleep(5 * time.Millisecond)
	fmt.Println("value recived is", v)
	fmt.Println("")
}
