// program without go routine with annonymous func
/*
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	for _, v := range a {
		func(i int) {
			fmt.Println("check value is ", i)
		}(v)
	}
}
*/
// program with annonymus with go routines without passing argument and param to annonymous func
package main

import (
	"fmt"
	"time"
)

func main() {
	// with goroutine output changes
	a := []int{1, 2, 3, 4, 5, 6}
	for _, v := range a {
		// if try to pass argument and param will give proper output or will print garbage output
		go func(i int) {
			time.Sleep(1 * time.Second)
			fmt.Println("check value is ", i)
		}(v)
	}
	time.Sleep(5 * time.Second)

}
