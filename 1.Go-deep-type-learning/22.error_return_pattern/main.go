package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}


}

func run() error {

	input := "30"

	level, err := parseLevel(input)
	if err != nil {
		return err
	}
	fmt.Println("selected lvl", level)
	return nil
}

func parseLevel(s string) (int, error) {
	//(value,err)
	//nil error -> success
	//non-nil error -> failure

	//pattern
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("s is must be number")
	}

	if v < 1 || v > 5 {
		return 0, fmt.Errorf("v must be 1 and 5")
	}

	return v, nil
}
