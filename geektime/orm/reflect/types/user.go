package types

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUserPoint(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

func NewUser(name string, age int) User {
	return User{
		Name: name,
		Age:  age,
	}
}

func (u User) GetAge() int {
	return u.Age
}

func (u *User) ChangeName(NewName string) {
	u.Name = NewName
}

func (u *User) private() {
	fmt.Println("private")
}
