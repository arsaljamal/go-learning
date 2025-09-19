package  main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Fake API call
func callAPI(workerID, jobID int) {
	fmt.Printf("Worker %d -> Job %d at %s\n",
		workerID, jobID, time.Now().Format("15:04:05.000"))
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300))) // simulate API latency
}

// Worker function
func worker(id int, jobs <-chan int, wg *sync.WaitGroup, limiter <-chan time.Time) {
	defer wg.Done()
	for job := range jobs {
		<-limiter // wait for rate limiter token
		callAPI(id, job)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numWorkers := 5
	numJobs := 20

	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Rate limiter: 5 requests per second (1 every 200ms)
	limiter := time.Tick(200 * time.Millisecond)

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg, limiter)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for workers
	wg.Wait()
	fmt.Println("All jobs processed ✅")
}
