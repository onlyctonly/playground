package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	args := os.Args[1:]
	fmt.Println(args)

	limiter := time.Tick(200 * time.Millisecond)

	for v := range requests {
		<-limiter
		fmt.Println(v, time.Now())
	}

	burstLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for v := range burstyRequests {
		<-burstLimiter
		fmt.Println(v, time.Now())
	}

}
