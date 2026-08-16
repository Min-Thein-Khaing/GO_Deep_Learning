package main

import (
	
	"fmt"
	
)

//that is my logic

// func main() {
// 	score := 10
// 	fmt.Println("before address is",&score)
// 	fmt.Println("before value is",score)

// 	addScore(score)
// 	// fmt.Println("after",score)

	
// }

// func addScore(s int) {
// 	p := &s
// 	fmt.Println("value is",*p)

// 	*p += 10
// 	fmt.Println("new value is",*p)
// }



//that is other logic
func main() {
	score := 10
	fmt.Println("before value is",score)

	addScore(&score)
	fmt.Println("after",score)

	
}

func addScore(s *int) {
	*s += 10
}