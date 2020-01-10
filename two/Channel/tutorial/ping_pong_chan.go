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

	// put `passed message into ping chan`
	ping(pings, "passed message")

	// then pass the string to `pong chan`
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
