package main

import (
	"fmt"
	"sync"
)

// main demonstrates the use of a mutex to safely increment a shared counter variable
// across multiple goroutines. It launches 5 goroutines, each incrementing the counter
// in a thread-safe manner using the incrementCounter function. The final value of the
// counter is printed after all goroutines complete.
func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	counter := 0

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrementCounter(&mutex, &counter)
		}()
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementCounter(mutex *sync.Mutex, counter *int) {
	mutex.Lock()
	defer mutex.Unlock()
	*counter++
}
