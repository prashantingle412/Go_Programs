// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	jobs := make(chan int, 10)
// 	result := make(chan int, 10)

// 	for j := 1; j <= 5; j++ {
// 		go worker(jobs, result)
// 	}
// 	for j := 1; j <= 10; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)
// 	for k := 1; k <= 10; k++ {
// 		fmt.Println("recived jobs", <-result)
// 	}
// 	time.Sleep(1 * time.Second)
// }
// func worker(jobs chan int, result chan int) {
// 	for data := range jobs {
// 		fmt.Println("worker started working on job.....", data)
// 		result <- data
// 	}
// }

package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)

	for j := 1; j <= 10; j++ {
		go worker(jobs)
	}
	// close(jobs)
	time.Sleep(2 * time.Second)
}
func worker(jobs chan int) {
	fmt.Println("called goroutine")
	jobs <- 1
	time.Sleep(100 * time.Second)
	for k := 1; k <= 10; k++ {
		fmt.Println("recived jobs", <-jobs)
	}
}
