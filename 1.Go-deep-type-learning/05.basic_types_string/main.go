package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "min thein "
	lastName := "khaing"

	fullName := firstName + " " + lastName

	fmt.Println(strings.ToUpper(fullName))
	fmt.Println(strings.ToLower(fullName))
	fmt.Println(strings.TrimSpace(fullName))
	fmt.Println(strings.Contains(fullName, "khaing"))
	fmt.Println(strings.Split(fullName, " "))
	fmt.Println(strings.Fields(fullName))
	fmt.Println(len(fullName))
	fmt.Println(strings.Replace(fullName, "min", "max", 1))
	fmt.Println(strings.HasPrefix(fullName, "min"))
	fmt.Println(strings.HasSuffix(fullName, "khaing"))
	fmt.Println(strings.Count(fullName, "a"))
	fmt.Println(strings.Repeat(fullName, 3))
	fmt.Println(strings.Index(fullName, "khaing"))
	fmt.Println(strings.LastIndex(fullName, "khaing"))
	
	

}
