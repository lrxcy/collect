package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 預設一個函數，固定對傳進來的chan送入字串'ping'
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// 預設一個函數，專門打印送進來的chan裡面的字串
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	// 宣告一個chan
	var c chan string = make(chan string)

	// 將這個chan送到ping裡面，讓程式可以固定送字串到chan裡
	go pinger(c)
	// 將chan送到printer，讓chan裡面的值會固定打印出來
	go printer(c)

	// 使用Scanln來搜尋使用者輸入的字串，如果輸入enter就會停止程序
	// var input string
	// fmt.Scanln(&input)

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
