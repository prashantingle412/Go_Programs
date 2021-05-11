package impl

type User struct {
	Name string
}

func (u User) Find() string {
	return u.Name
}
func (u User) GetProducts() string {
	return u.Name
}
func (u User) Add() string {
	return u.Name
}
