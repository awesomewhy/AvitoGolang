package main

import "fmt"

type User struct {
	name     string
	age      int
	password string
}

func (u *User) setName(a string) {
	u.name = a
}

func (u User) getName() string {
	return u.name
}

func main() {
	a := User{"ky", 16, "qwe"}
	a.setName("sdf")
	fmt.Println(a.getName())
}
