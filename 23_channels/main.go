/*
Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.


*/
package main

import "fmt"

func main() {

	//Create a new channel with make(chan val-type). Channels are typed by the values they convey
	messages := make(chan int)

	go func() {
		for j := 1; j <= 9; j++ {
			messages <- j
		}

	}()
	go func() {
		for j := 1; j <= 9; j++ {
			msg := <-messages
			fmt.Println(msg)
		}
	}()

	fmt.Scanln()

}
