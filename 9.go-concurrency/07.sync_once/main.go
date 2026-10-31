package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var once sync.Once

	worker := 5

	for i := 1; i <= worker; i++ {
		wg.Add(1)
		go func(id int) {

			defer wg.Done()
			once.Do(setup)

			fmt.Println(id)

		}(i)

	}

	wg.Wait()
	fmt.Println("done 1")

}

func setup() {
	fmt.Println("start")
	time.Sleep(1 * time.Second)
	fmt.Println("done 2")
}
