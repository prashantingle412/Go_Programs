package main

import "fmt"

func greet(c chan string) {
	fmt.Println("hello" + <-c + "!")
}
func main() {
	fmt.Println("main started")
	c := make(chan string)
	go greet(c)
	c <- "prashant"
	fmt.Println("main stopped")
}
