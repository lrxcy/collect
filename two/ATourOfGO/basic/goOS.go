// https://gobyexample.com/signals
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func showHostname() {
	hostname, _ := os.Hostname()
	fmt.Println(hostname)
	log.Fatal(123)

}

func main() {

	// create one chan to print awaiting signal on console
	sigs := make(chan os.Signal, 1)
	// create another chan to receive signal to interrupt original chan
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
