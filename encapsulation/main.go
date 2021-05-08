package main

import (
	"encapsulation/data"
	"fmt"
)

func main() {
	usr, err := data.Helper(100, "admin")
	if err != nil {
		panic(err)
	}
	fmt.Println("interface method is called and price is", usr.GetPrice())
	// fmt.Println("price can change as its private", p)

}
