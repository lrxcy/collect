package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// create one chan to print awaiting signal on console
	sigs := make(chan os.Signal, 1)
	// create another chan to receive signal to interrupt original chan
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
