// package main

// import (
// 	"fmt"
// 	"time"
// )

// // func greet(c chan string) {
// // 	// fmt.Println("hello" + <-c + "!")
// // 	c <- "data pushed to chan from greet func"
// // }
// func main() {
// 	fmt.Println("main started")
// 	c := make(chan int)
// 	go func(ch chan int) {
// 		for i := 1; i <= 10; i++ {
// 			ch <- i
// 		}
// 	}(c)
// 	go func(ch chan int) {
// 		for i := 1; i <= 10; i++ {
// 			v, ok := <-c
// 			if !ok {
// 				break
// 			}
// 			fmt.Println("recived data from ch", v)
// 		}
// 	}(c)
// 	time.Sleep(2 * time.Second)

// 	fmt.Println("main stopped")
// }

