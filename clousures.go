package main

import "fmt"

func counter() func() int {
	var i = 0
	return func() int {
		i++
		return i
	}
}
func main() {
	fmt.Println("welcome to clousure ")
	c := counter()
	fmt.Println("closure return", c())
	fmt.Println("closure return", c())
	fmt.Println("closure return ", c())
}
