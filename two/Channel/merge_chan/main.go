package main

import (
	"fmt"
)

var intArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

func ping(pings chan<- int, msg int) {
	pings <- msg
}

func pingping(pings chan<- int, msg int) {
	pings <- msg * 2
}

func pong(pings <-chan int, pongs chan<- int) {
	msg := <-pings
	pongs <- msg
}

var counter int

func main() {
	counter = 0
	pings := make(chan int, 5)
	pongs := make(chan int, 5)

	// as sender to send data to channels
	go func() {
		for _, j := range intArray {
			ping(pings, j)
			counter++
		}
	}()

	// as sender to send data to channels
	go func() {
		for _, j := range intArray {
			pingping(pings, j)
			counter++
		}
	}()

	// as receiver to receive data from above channels
	go func() {
		for {
			pong(pings, pongs)
		}
	}()

	for {
		// check whether the channel is empty or not
		// need to use select `block` having default case
		select {
		case x, ok := <-pongs:
			if ok {
				fmt.Println(x)

			} else {
				fmt.Println("channel is going to closed")
			}
		default:
			break
		}
	}

}
