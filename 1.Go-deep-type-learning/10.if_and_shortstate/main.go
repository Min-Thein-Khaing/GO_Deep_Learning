package main

import (
	"fmt"
)

func main() {
	items := 3
	pricePerItems := 49

	if price := items * pricePerItems; price > 100 {
		fmt.Println("High price")
	} else if price > 50 {
		fmt.Println("Medium price")
	} else {
		fmt.Println("Low price")
	}
}
