package my_package

type User struct {
	Name string
}

func New(name string) User{
	return User{Name: name}
}
