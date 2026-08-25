package main

import (
	"fmt"
	"time"
)


func main() {
	start := time.Now()

	
	go func ()  {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Hello from goroutine:first ", time.Since(start))
	}()


	go func ()  {
		time.Sleep(150 * time.Millisecond)
		fmt.Println("Hello from goroutine:second", time.Since(start))
	}()

	fmt.Println("main : started two go routine",time.Since(start))
	fmt.Println("main : doing step 1",time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("main : doing step 2",time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("main : doing step 3",time.Since(start))

	time.Sleep(500 * time.Millisecond)

	fmt.Println("main : exit",time.Since(start))

}