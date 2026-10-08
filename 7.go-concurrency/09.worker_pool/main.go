package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID    int
	Email string
}

type Result struct {
	JobID   int
	Success bool
	Worker  int
}

// worker processes jobs from the jobs channel and outputs to results channel
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("[Worker %d] Started processing job %d for %s\n", id, job.ID, job.Email)
		time.Sleep(200 * time.Millisecond) // Simulate sending email or processing
		fmt.Printf("[Worker %d] Finished job %d\n", id, job.ID)

		results <- Result{
			JobID:   job.ID,
			Success: true,
			Worker:  id,
		}
	}
}

func main() {
	const numJobs = 10
	const numWorkers = 3 // Limit concurrency to 3 workers

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	// 1. Start fixed number of workers (Worker Pool)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// 2. Send jobs to the jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{
			ID:    j,
			Email: fmt.Sprintf("user%d@example.com", j),
		}
	}
	close(jobs) // Close jobs channel so workers know when to exit range loop

	// 3. Wait for all workers to finish in a separate goroutine, then close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// 4. Collect and print all results
	fmt.Println("\n=== Collecting Results ===")
	for res := range results {
		fmt.Printf("Job %d completed successfully by Worker %d\n", res.JobID, res.Worker)
	}

	fmt.Println("\nAll jobs processed via Worker Pool!")
}
