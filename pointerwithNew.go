package main

import "fmt"

type User struct {
	name string
}

func main() {
	user1 := new(User)
	user2 := new(User)
	// fmt.Println("user1 and user2 ", user1, user2)
	fmt.Printf("%v %p %v \n", &user1, user1, *user1)
	fmt.Printf("%v %p %v \n", &user2, user2, *user2)
}
