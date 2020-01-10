package main

import "fmt"

func main() {
	// claim a predefine buffered chan to store at least 4 elements
	ch := make(chan int, 4)
	ch <- 2
	ch <- 4

	// where we close the chan while there is only 2 elements in chan
	close(ch)
	// ch <-6 // panic, send on closed channel

	// after each println, chan is like a queue pop out element
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch) // closed, retruns zero value for element
}
