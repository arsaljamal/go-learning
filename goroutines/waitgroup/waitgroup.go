package main

import (
	"fmt"
	"sync"
)

func asyncFunc(n int) {
	fmt.Println(n)
}

func asyncFunc2(n int) {
	fmt.Println(n)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		asyncFunc(0)
		asyncFunc2(1)
		wg.Done()
	}()
	wg.Wait()
}
