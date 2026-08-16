package main

import (
	"fmt"
)

func main() {
	views1 := 1000
	views2 := 1000

	viewsTotal := views1 + views2

	likes := 10

	likes++
	likes++

	avgView := viewsTotal / 2

	fmt.Println(avgView, likes, viewsTotal)

	rating1 := 4.5
	rating2 := 5.1

	avgRating := (rating1 + rating2) / 2

	fmt.Println(avgRating)

}