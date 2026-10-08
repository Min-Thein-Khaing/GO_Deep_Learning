package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter wraps a counter and a sync.Mutex to protect shared memory
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// UserCache demonstrates sync.RWMutex (Multiple Readers, Single Writer)
type UserCache struct {
	rw    sync.RWMutex
	store map[string]string
}

func NewUserCache() *UserCache {
	return &UserCache{
		store: make(map[string]string),
	}
}

// Set requires write lock: no other reader or writer can access
func (c *UserCache) Set(key, val string) {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.store[key] = val
}

// Get uses read lock: multiple goroutines can read simultaneously
func (c *UserCache) Get(key string) (string, bool) {
	c.rw.RLock()
	defer c.rw.RUnlock()
	val, ok := c.store[key]
	return val, ok
}

func main() {
	fmt.Println("=== 1. Testing Mutex with Counter ===")
	counter := SafeCounter{}
	var wg sync.WaitGroup

	// 1000 goroutines incrementing the same counter concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()
	// Without Mutex, count would be unpredictable and data race would happen
	// Try checking with: go run -race main.go
	fmt.Printf("Expected: 1000, Actual Counter: %d\n\n", counter.Value())

	fmt.Println("=== 2. Testing RWMutex with Shared Cache ===")
	cache := NewUserCache()
	var cacheWg sync.WaitGroup

	// Concurrent Writer
	cacheWg.Add(1)
	go func() {
		defer cacheWg.Done()
		time.Sleep(50 * time.Millisecond)
		cache.Set("user_1", "Alice")
		fmt.Println("[Writer] Saved user_1: Alice")
	}()

	// Concurrent Readers
	for i := 1; i <= 3; i++ {
		cacheWg.Add(1)
		go func(readerID int) {
			defer cacheWg.Done()
			time.Sleep(100 * time.Millisecond)
			if name, ok := cache.Get("user_1"); ok {
				fmt.Printf("[Reader %d] Found user_1: %s\n", readerID, name)
			}
		}(i)
	}

	cacheWg.Wait()
	fmt.Println("All done safely without data race!")
}
