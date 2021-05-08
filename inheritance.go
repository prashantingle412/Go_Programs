package main

import "fmt"

type Parent struct {
	name string
}

type Child struct {
	Parent Parent
	city   string
}

func main() {
	p := Parent{name: "bhaurao"}
	c := Child{city: "pune", Parent: p}
	pname := c.Parent.printbase()
	ccity := c.printchild()
	fmt.Println("parent class func called", pname, "city is", ccity)
}

func (p Parent) printbase() string {
	return p.name
}
func (c Child) printchild() string {
	return c.city
}
