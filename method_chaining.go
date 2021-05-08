package main

import "fmt"

type Employee struct {
	name    string
	city    string
	address string
}

func main() {
	var employee Employee
	employee.Name("prashant").City("pune").Prints()
}

func (e Employee) Name(name string) Employee {
	e.name = name
	return e
}
func (e Employee) City(name string) Employee {
	e.name = name
	return e
}
func (e Employee) Prints() {
	fmt.Println("employee is", e)
}

// employee.Name("Prashant").WorksFor("Persistent").LivesIn("Pune").Build()
