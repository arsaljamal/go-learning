package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asyncReceiver(c chan int) {
	for {
		n := <-c
		fmt.Println(n)
	}
}

func asyncSender(c chan int) {
	for {
		n := rand.Intn(10)
		c <- n
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	c := make(chan int, 10)
	go asyncReceiver(c)
	go asyncSender(c)
	time.Sleep(time.Second * 2)
}