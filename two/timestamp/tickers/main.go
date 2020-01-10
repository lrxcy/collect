package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	// 背景執行: 開始執行ticker，每500 milesecond即執行一次
	go func() {
		for {
			select {
			// 如果done沒有值，則不進行return，如果return則結束執行程序;使用done來控制背景線程，間接控制整個主線程
			case <-done:
				return

			// 當ticker有值進來的時候進行打印
			case t := <-ticker.C:
				fmt.Println("Tick at ", t)
			}
		}
	}()

	// 主線程停1600秒後停止ticker，並且傳送bool(true)給done，結束整個線程
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
