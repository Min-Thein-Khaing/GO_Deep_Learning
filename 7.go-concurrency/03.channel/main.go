package main

import "fmt"

func main() {
	type User struct {
		ID   int
		Name string
	}

	ch := make(chan User)

	go func() {
		ch <- User{ID: 1, Name: "mintheinkhaing"}
	}()

	fmt.Println("main : waiting to receive user")

	u := <-ch
	fmt.Println("main : received user", u)
}
