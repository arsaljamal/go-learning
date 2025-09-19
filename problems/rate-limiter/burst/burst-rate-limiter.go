package main

import (
	"fmt"
	"time"
)

// Fake API call
func callAPI(id int) {
	fmt.Println("Calling API for request:", id, "at", time.Now().Format("15:04:05.000"))
}

func main() {
	// Token bucket: allow bursts of up to 10, refill at 5 tokens/sec
	bucket := make(chan struct{}, 10)

	// Refill tokens
	go func() {
		ticker := time.NewTicker(200 * time.Millisecond) // 5/sec
		for range ticker.C {
			select {
			case bucket <- struct{}{}:
			default:
				// bucket full, drop token
			}
		}
	}()

	// Simulate 20 requests
	for i := 1; i <= 20; i++ {
		<-bucket // wait for token
		callAPI(i)
	}
}
