package main

import "fmt"

type Circle struct {
	height int
}
type Rect struct {
	height, width int
}
type calculator interface {
	calculate() int
}

func main() {
	// r := Rect{height: 10, width: 10}
	c := Circle{height: 100}
	fmt.Println("area is", print(c))
}
func (c Circle) calculate() int {
	return c.height * 2
}

func (r Rect) calculate() int {
	return r.height * r.width
}
func print(c calculator) int {
	return c.calculate()
}
