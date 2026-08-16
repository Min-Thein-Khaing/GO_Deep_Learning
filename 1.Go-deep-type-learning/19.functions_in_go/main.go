package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func sumAndProduct(a int , b int) (int,int){
	sum := a + b 
	product := a*b 
	return sum ,product
}
func main() {
	sum1 := add(4,6)
	fmt.Println(sum1)

	sum2 , product := sumAndProduct(4,6)
	onlySum,_ := sumAndProduct(4,6)

	fmt.Println(sum2,product)
	fmt.Println(onlySum)

}
