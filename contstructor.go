package main

import "fmt"

type Person struct {
	Name string
	City string
}

func main() {
	p := NewPerson("prashant", "pune")
	fmt.Println("object created with constructor", p)
}

func NewPerson(name string, city string) *Person {
	return &Person{
		Name: name,
		City: city,
	}
}
