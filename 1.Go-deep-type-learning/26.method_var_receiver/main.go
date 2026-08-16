package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	u1 := User{
		Name: "min thein khaing",
		Age:  28,
	}

	fmt.Println(u1.intro())

}

//val receiver means this method receivers a copy of the user

func (u User) intro() string {
	return fmt.Sprintf("Hi,I'am %s", u.Name)
}
