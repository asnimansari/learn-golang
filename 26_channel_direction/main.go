/*
when a channel is passed as argument we can mention whether its used for sending or recieving
this improves type safety of program
*/
package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "pased msge")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
