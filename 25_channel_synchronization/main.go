/*

We can use channels to synchronize execution across goroutines. Here’s an example of using a blocking receive to wait for a goroutine to finish.

*/

package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("Worker is done")

	done <- true

}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
	fmt.Println("Control is back to main")

}
