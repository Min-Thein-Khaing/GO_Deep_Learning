package main

import (
	"fmt"
)

func main() {

	names := []string{"Ali", "Hania", "Ahmed"}
	fmt.Println(names)

	newName := make([]string, 2, 3)
	newName[0] = "Ali"
	newName[1] = "Hania"
	newName = append(newName, "Ahmed")
	newName = append(newName, "Ayesha")

	fmt.Println(newName)
	newName[1] = "Min"
	fmt.Println(newName)

	fmt.Println(cap(newName))

	todos := []string{"do", "workout learn"}

	more := []string{"Coding","Problem Solving"}

	todos = append(todos, more...)

	fmt.Println(todos)
	fmt.Println(cap(todos))
}
