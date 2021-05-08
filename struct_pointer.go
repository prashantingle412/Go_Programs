package main

import "fmt"

type Student struct {
	name string
	city string
}

func (s *Student) m1() {
	s.name = "ravi"
}
func (s *Student) m2() {
	s.name = "avi"
	s.city = "pune"
}

func main() {
	p := &Student{}
	p.m1()
	p2 := new(Student)
	p2.m2()
	fmt.Println("after call", p, p2)
}
