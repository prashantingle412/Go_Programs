package singleton

type Singleton interface {
	Addone() int
}
type singletonstr struct {
	count int
}

var instance *singletonstr

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singletonstr)
	}
	return instance
}
func (s *singletonstr) Addone() int {
	s.count++
	return s.count
}
