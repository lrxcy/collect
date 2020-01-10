package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() { messages <- "ping" }()
	// go func() { messages <- "no ping" }()  //'<-'表示將值傳到chan裡面，且chan可以被覆寫

	msg := <-messages
	fmt.Println(msg)
}
