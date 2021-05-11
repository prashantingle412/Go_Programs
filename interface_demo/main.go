package main

import (
	"fmt"
	"interface/fmtf"
	"interface_demo/impl"
	"interface_demo/inter"
)

// interface implemented struct
type service struct {
	getter inter.Getter
	setter inter.Setter
}

func main() {
	var inter1 inter.Getter
	inter1 = impl.User{Name: "prashant"}
	fmt.Println("interface method called", inter1.Find())
	var other fmtf.Finder
	other = impl.User{Name: "product1"}
	fmt.Println("calling other product interface method", other.GetProducts())
	//  implemented interface on struct
	var s service
	u3 := impl.User{Name: "getter interface method"}
	g := impl.User{Name: "setter interface method"}
	s.getter = u3
	s.setter = g
	fmt.Println("getter and setter struct", s)
}
