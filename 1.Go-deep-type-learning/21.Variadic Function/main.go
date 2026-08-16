package main

import (
	"fmt"
)

//numbers = []int

func sum(numbers ...int)int{
	total := 0

	for _,v := range numbers {
		total += v
	}
	return total
}

func main() {
	numbers := []int{1,2,3,4,5}
	fmt.Println(sum(numbers...))

	fmt.Println(sum(1,2,3,4,5))

	res := func(n int) int {
		return n*2
	}

	// res := func(n int) int {
	// 	return n*2
	// }(5) that is same as res(5)
	fmt.Println(res(5))

	//IIFE

	res1 := func(a int,b int)int{
		return a + b
	}(5,55)
	fmt.Println(res1)
}
