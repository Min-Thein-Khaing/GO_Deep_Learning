package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	go slowWorker(ctx)

	//main wait until ctx is done
	<-ctx.Done()
	fmt.Println("main:content end",ctx.Err())
	fmt.Println("main:exit")
}

func slowWorker(ctx context.Context) {
	select {

	case <-time.After(300 * time.Millisecond):
		return
	case <-ctx.Done():
		fmt.Println("Done")
	}

}
