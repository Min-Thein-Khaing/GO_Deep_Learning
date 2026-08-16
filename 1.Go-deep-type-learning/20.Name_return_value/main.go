package main

import (
	"fmt"
)


func divided (a int , b int) (jhon int,sagam int){
	jhon = a/b 
	sagam = a%b
	return 
}

func main() {
	jhon ,sagam := divided(10,3)
	fmt.Println(jhon ,sagam)
}
