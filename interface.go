package main

import (
	"fmt"
)

type User struct {
	name string
}
type reader interface {
	Get()
	GetByID()
}

func main() {
	u := User{name: "prashant"}
	u.Get()
	u.GetByID()
}
func (u User) Get() {
	fmt.Println("get func is called", u.name)
}

func (u User) GetByID() {
	fmt.Println("getByID func is called", u.name)
}
