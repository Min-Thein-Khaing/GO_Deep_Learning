package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

// simulateService simulates an external service call that respects context
func simulateService(ctx context.Context, serviceName string, delay time.Duration, shouldFail bool) error {
	select {
	case <-time.After(delay):
		if shouldFail {
			fmt.Printf("[%s] Failed with error!\n", serviceName)
			return errors.New(serviceName + " connection failed")
		}
		fmt.Printf("[%s] Completed successfully\n", serviceName)
		return nil
	case <-ctx.Done():
		fmt.Printf("[%s] Aborted because context was cancelled!\n", serviceName)
		return ctx.Err()
	}
}

func main() {
	fmt.Println("=== 1. All Tasks Succeed ===")
	g1, ctx1 := errgroup.WithContext(context.Background())

	g1.Go(func() error {
		return simulateService(ctx1, "Auth Service", 100*time.Millisecond, false)
	})
	g1.Go(func() error {
		return simulateService(ctx1, "Database Service", 150*time.Millisecond, false)
	})
	g1.Go(func() error {
		return simulateService(ctx1, "Payment Service", 200*time.Millisecond, false)
	})

	if err := g1.Wait(); err != nil {
		fmt.Printf("Group 1 Error: %v\n", err)
	} else {
		fmt.Println("Group 1: All services responded successfully!")
	}

	fmt.Println("\n=== 2. One Task Fails (Auto Cancel Others) ===")
	// When one goroutine returns an error, the group context is automatically cancelled!
	g2, ctx2 := errgroup.WithContext(context.Background())

	g2.Go(func() error {
		// This one fails after 100ms
		return simulateService(ctx2, "Order Service", 100*time.Millisecond, true)
	})
	g2.Go(func() error {
		// This one takes 300ms, will be aborted early when Order Service fails!
		return simulateService(ctx2, "Slow Inventory Service", 300*time.Millisecond, false)
	})
	g2.Go(func() error {
		// This one takes 400ms, will be aborted early too!
		return simulateService(ctx2, "Analytics Service", 400*time.Millisecond, false)
	})

	if err := g2.Wait(); err != nil {
		fmt.Printf("Group 2 Failed with: %v\n", err)
	}

	fmt.Println("\nErrGroup cleanly managed errors and cancelled useless in-flight work!")
}
