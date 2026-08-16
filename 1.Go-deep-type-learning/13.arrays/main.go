package main

import (
	"fmt"
)

func main() {
	var marks [3]int

	marks[0] = 80
	marks[1] = 90
	marks[2] = 100

	fmt.Println("first student marks", marks[0])
	fmt.Println("second student marks", marks[1])
	fmt.Println("third student marks", marks[2])

	schoolSubjectMark := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(schoolSubjectMark))
}
