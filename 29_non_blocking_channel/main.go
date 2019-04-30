package main

import "fmt"

/*
Basic sends and receives on channels are blocking. However, we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects.
*/

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("Recieved Message", msg)
	default:
		fmt.Println("No Message Recieved")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("Sent Message ", msg)
	default:
		fmt.Println("No Message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("Recieved Message ", msg)
	case sig := <-signals:
		fmt.Println("Recieved Siganl", sig)
	default:
		fmt.Println("No Activity")
	}
}
