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
	// Allow 5 requests per second
	limiter := time.Tick(200 * time.Millisecond) // 1000ms / 5 = 200ms per request

	// Simulate 20 requests
	for i := 1; i <= 20; i++ {
		<-limiter // wait before sending
		callAPI(i)
	}
}
