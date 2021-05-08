package data

import "errors"

// product is private struct we cant get it available outside
type product struct {
	Price int
}

// created interface method to access private outside
type Product interface {
	GetPrice() int
}

// func to set price with 10 % tax and returning
func (p product) GetPrice() int {
	return p.Price + 10
}

// helper function to call original GetPrice()
// only admin can access this module
func Helper(price int, role string) (Product, error) {
	if role != "admin" {
		return nil, errors.New("normal user cant access this module, need to be admin ")
	}
	var p product
	p.Price = price
	return p, nil
}
