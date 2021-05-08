package main

import (
	"fmt"
	"interface_demo/impl"
	"interface_demo/inter"
)

func main() {
	var inter1 inter.Getter
	inter1 = impl.User{Name: "prashant"}
	fmt.Println("interface method called", inter1.Find())
}
