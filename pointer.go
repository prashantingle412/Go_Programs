// package main

// import "fmt"

// func main() {
// 	var name = "prashant"
// 	fmt.Println("before caalled to  change func", name)
// 	change(&name)
// 	fmt.Println("after call to change func, from main", name)
// }
// func change(name *string) {
// 	*name = "ravi"
// 	fmt.Println("value modified in change func ")
// }

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api", Hello)
	http.ListenAndServe(":8080", nil)
}
func Hello(res http.ResponseWriter, req *http.Request) {
	fmt.Println("hello from api")
}
