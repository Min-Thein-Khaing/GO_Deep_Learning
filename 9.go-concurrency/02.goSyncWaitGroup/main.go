package main

import (
	"fmt"
	"sync"
	"time"
)


func main(){
	var wg sync.WaitGroup

	wg.Add(3)

	go func (){
		defer wg.Done()	
		fmt.Println("Task 1")
		time.Sleep(250 * time.Millisecond)
		fmt.Println("Task 1 doing now")
	}()

	go func (){
		defer wg.Done()	
		fmt.Println("Task 2")
		time.Sleep(150 * time.Millisecond)
		fmt.Println("Task 2 doing now")
	}()

	go func (){
		defer wg.Done()	
		fmt.Println("Task 3")
		time.Sleep(199 * time.Millisecond)
		fmt.Println("Task 3 doing now")
	}()

	fmt.Println("main : waiting to finshing ")
	wg.Wait()
	fmt.Println("main : all tasks are finished")
}
