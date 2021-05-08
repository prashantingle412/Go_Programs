package main

import "fmt"

type CreditCard struct {
	Method string
	Amount string
}
type DebitCard struct {
	Method string
	Amount string
}
type Cash struct {
	Method string
	Amount string
}
type Getter interface {
	Pay() string
}

func (c CreditCard) Pay() string {
	return "you bill is :" + c.Amount + " and paid via" + c.Method
}
func (d DebitCard) Pay() string {
	return "you bill is :" + d.Amount + " and paid via" + d.Method
}
func (c Cash) Pay() string {
	return "you bill is :" + c.Amount + " and paid via" + c.Method
}
func main() {
	var g Getter
	c := CreditCard{Method: "credit card", Amount: "$1000"}
	g = c
	result := g.Pay()
	fmt.Println(result)
}
