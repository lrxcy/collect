package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	startGs := runtime.NumGoroutine()
	fmt.Println("The start goroutine num: ", startGs)

	// 宣告一個chan c1
	c1 := make(chan string)

	// 在下面的goroutine中將'result 1'送入c1
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("%v\n", err)
			}
		}()
		time.Sleep(time.Second * 5)
		c1 <- "result 1"
	}()

	// 使用select來做邏輯判斷，判斷要印出的res
	select {
	case res := <-c1:
		fmt.Println(res)
	// 如果超過1秒就印出timeout 1並且跳出回圈
	case <-time.After(time.Second * 1):
		close(c1)
		res := <-c1
		fmt.Println(res)
		fmt.Println("timeout 1")
	}

	// Capture ending number of goroutines.
	endingGs := runtime.NumGoroutine()
	fmt.Println("The end goroutine num: ", endingGs)

	fmt.Println(endingGs - startGs)

	time.Sleep(time.Second * 10)

	endingGs2 := runtime.NumGoroutine()
	fmt.Println("The second end goroutine num: ", endingGs2)

	fmt.Println(endingGs2 - startGs)
}
