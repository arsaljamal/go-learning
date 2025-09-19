package main


import (
	"fmt"
	"time"
)

func wokrer(id int, jobs <- chan int, results chan<- int) {
	// Worker function to process jobs from the jobs channel
	//  and send results to the results channel.
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start 3 worker goroutines.
	for w:= 1; w<= 3; w++ {
		go wokrer(w, jobs, results)
	}

	// Send 5 jobs to the jobs channel.
	for j:= 1; j<= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results from the results channel.
	for a:= 1; a<= 5; a++ {
		fmt.Println("Result:", <-results)
	}

	//Imagine you own a pizza shop 🍕. Customers place 5 orders (jobs). 
	//You have 3 chefs (workers). Instead of 1 chef making all pizzas,
	// you divide them across 3 chefs in parallel.
	// The jobs channel is the order queue,
	// results is where the finished pizzas go.
}