package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//  Burst Limter
	burstLimter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstLimter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstLimter <- t
		}
	}()

	burstRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstRequests <- i
	}
	close(burstRequests)

	for req := range burstRequests {
		<-burstLimter
		fmt.Println("request", req, time.Now())
	}
}
