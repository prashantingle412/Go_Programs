package main

import "fmt"

type Intr interface {
	m1()
}

type Student struct {
	name string
}

func (s *Student) m1() {
	s.name = "ravi"
	fmt.Println("interface func called", s.name)
}

func main() {
	p := &Student{}
	p.m1()
}

func myfnc(i Intr) {
}
