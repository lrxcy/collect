package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// 先宣告一個變數，做計算
	var ops uint64

	// 宣告一個wait group來等待所有的goroutine完成各自的工作
	var wg sync.WaitGroup

	// 啟用50個goroutine來執行`加1的`動作`1000次`
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {

				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)
}
