package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	fmt.Println("case 1 : working")

	if err := doWork(true) ; err != nil {
		log.Fatal(err)
	}
	fmt.Println("case 1 : fail")
if err := doWork(false) ; err != nil {
		log.Fatal(err)
	}
}

func doWork(success bool) error {

	fmt.Println("start : resource acquired")

	defer fmt.Println("clean : resource acquired")

	if !success {
		return errors.New("something went wrong.i am returning early")
	}
	fmt.Println("working doing smthing imp")
	fmt.Println("work :this work is done")
	return nil

}
