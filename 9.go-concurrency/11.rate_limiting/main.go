package main

import (
	"fmt"
	"time"
)

func main() {
	// ==========================================
	// 1. Basic Rate Limiting (Fixed Rate)
	// ==========================================
	fmt.Println("=== 1. Basic Rate Limiting (1 request every 200ms) ===")
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Ticker emits a tick every 200ms
	limiter := time.NewTicker(200 * time.Millisecond)
	defer limiter.Stop()

	for req := range requests {
		<-limiter.C // Blocks until ticker sends tick
		fmt.Printf("Processed request %d at %v\n", req, time.Now().Format("15:04:05.000"))
	}

	// ==========================================
	// 2. Bursty Rate Limiting (Token Bucket)
	// ==========================================
	fmt.Println("\n=== 2. Bursty Rate Limiting (Burst of 3, then 1 per 200ms) ===")
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// Buffered channel acting as a Token Bucket with capacity of 3
	burstyLimiter := make(chan time.Time, 3)

	// Fill the bucket initially with 3 tokens (allows an initial burst of 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Add new token to the bucket every 200ms in the background
	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()
		for t := range ticker.C {
			select {
			case burstyLimiter <- t:
			default: // Bucket full, discard token
			}
		}
	}()

	for req := range burstyRequests {
		<-burstyLimiter // Takes 1 token from bucket
		fmt.Printf("Processed bursty request %d at %v\n", req, time.Now().Format("15:04:05.000"))
	}

	fmt.Println("\nRate limiting finished successfully!")
}
