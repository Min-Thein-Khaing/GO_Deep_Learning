package main

import (
	"fmt"
)

func main() {
	points := map[string]int{
		"a": 10,
		"b": 0,
	}

	fmt.Println("a", points["a"])
	fmt.Println("b", points["b"])
	fmt.Println("c", points["c"])

	valB, okB := points["b"]
	fmt.Println(valB, okB)

	valC, okC := points["c"]
	fmt.Println(valC, okC)

	if val, ok := points["c"]; ok {
		fmt.Println("c exists", val)
	} else {
		fmt.Println("c does not exist")
	}

	prices := map[string]int{
		"xyz": 500,
		"def": 1800,
	}

	if val, ok := prices["def"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("def does not exist")
	}

	total := 0
	for item, price := range prices {
		fmt.Println(item, price)
		total += price
	}
	fmt.Println(total)
}
