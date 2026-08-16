package main

import (
	"fmt"
)

func main() {
	views := []int{10,20,45,50,60}

	for index, view := range views {
		fmt.Println(index, view)
	}
}
