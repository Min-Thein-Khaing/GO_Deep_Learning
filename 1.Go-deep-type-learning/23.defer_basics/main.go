package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	fmt.Println("case 1 : working")

	if err := doWork(true); err != nil {
		log.Println("Error:", err) // log.Fatal အစား log.Println သုံးထားပါသည်
	}

	fmt.Println("\ncase 2 : fail")

	if err := doWork(false); err != nil {
		log.Println("Error:", err) // Program ဆက်လက် Run နိုင်မည် ဖြစ်သည်
	}
}

func doWork(success bool) error {

	fmt.Println("start : resource acquired")

	defer fmt.Println("clean : resource acquired")

	if !success {
		return errors.New("something went wrong. i am returning early")
	}
	fmt.Println("working doing smthing imp")
	fmt.Println("work : this work is done")
	return nil

}