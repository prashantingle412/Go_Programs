package impl

type User struct {
	Name string
}

func (u User) Find() string {
	return u.Name
}
