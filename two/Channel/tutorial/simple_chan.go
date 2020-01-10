package main

import "fmt"

func main() {
	// Create a channel
	message := make(chan string)

	// Init a goroutine
	go func() {
		// Send some data into the channel.
		message <- "Hello from channel"
	}()

	// Receive the data from the channel
	// hereby, no need to add sync.WaitGroup cause <- give process a block
	msg := <-message
	fmt.Println(msg)
}
