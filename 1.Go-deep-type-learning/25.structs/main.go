package main

import (
	"fmt"
)

type Person struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	u1 := Person{
		ID:    1,
		Name:  "min thein khaing",
		Email: "min@gmail.com",
		Age:   28,
	}
	fmt.Println(u1)

	u2 := Person{
		Name: "zaw zaw",
	}

	fmt.Println("partial user", u2)
}
