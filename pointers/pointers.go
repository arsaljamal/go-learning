package main

import "fmt"

func main() {
	var countPtr *int
	count := 3
	countPtr = &count

	fmt.Println("Address ", &countPtr)
	fmt.Println("Value ", *countPtr)

	*countPtr = 43

	fmt.Println("Value ", *countPtr)

	counter := 1
	increment(&counter)
	fmt.Println(counter)

	counterNew := newCounter()
	fmt.Println(*counterNew)
}


func increment(value *int) {
	*value += 1
}

func newCounter() *int {
	var count = 0
	return &count
}