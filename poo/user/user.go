package user

type User struct {
	Name string
	age  int
}

func (u User) Greet() string {
	return "Hola" + u.Name
}

func (u *User) Birthay() {
	u.age++
}
