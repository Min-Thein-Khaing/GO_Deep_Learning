package main

import (
	"fmt"
	"go-modules/internal/greet"
)

func main() {

	testingName := greet.Hello("min thein khaing")
	fmt.Println(testingName)

}
