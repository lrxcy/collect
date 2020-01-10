// https://tour.golang.org/concurrency/1
// goroutine 是由go 運行時環境管理的輕量級線程
// go f(x, y, z)開啟ㄧ個新的goroutine執行。
// 其中f, x, y, z是由當前goroutine中定義的，但是在新的goroutine中運行f

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
