package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Recieved JOBS", j)
			} else {
				fmt.Println("Recieved All Jobs")
				done <- true
				return
			}

		}
	}()

	for j := 1; j <= 4; j++ {
		jobs <- j
		fmt.Println("Sent Job", j)
	}
	close(jobs)
	fmt.Println("All Jobs Sent")
	<-done
}
