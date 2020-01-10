package main

import (
	"fmt"
)

func ChanGenerateMessage(message int) {
	// Buffered Channel of type Boolean
	done := make(chan bool, 1)
	go ChanPrintMessage(done, message)

	// Waiting to receive value from channel
	<-done
}

func ChanPrintMessage(done chan bool, message int) {
	defer func() {
		// Sending value to channel
		done <- true
	}()
	// Inside Logic (Dont frighten xD)
	// if IsTimeEnabled {
	// 	log.Println(message)
	// 	return
	// }
	fmt.Println(message)
}

func main() {
	for i := 0; i < 100; i++ {
		ChanGenerateMessage(i)
	}

}
