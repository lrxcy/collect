package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

/*
	透過宣告key這個值，可以拿取固定資料channel的值
*/
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	var readOps uint64
	var writeOps uint64

	// 創建一個讀取的實例，讓所有goroutine能夠從同一個channel被讀取
	reads := make(chan readOp)
	// 創建一個寫入的實例，讓所有goroutine能夠寫進同一個channel裡面的變數
	writes := make(chan writeOp)

	// 持續讀取...
	/*
		定義迴圈0~100
		持續生成資料結構`readOp`然後放置先準備好的chan reads
		等待被讀取 `<-reads.resp` 後才做資料同步
		同步後，休息一個millisecond再進行下次寫入
	*/
	for r := 0; r < 100; r++ {
		go func() {
			for {
				//
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 持續寫入...
	/*
		定義迴圈0~10
		持續生成資料結構`writeOp`然後放置預先準備好的chan writes
		等待被讀取 `<-write.resp` 後才做資料同步
		同步後，休息一個millisecond再進行下次寫入
	*/
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 此處的goroutine加上for迴圈，不斷的重複讀取與寫入數據
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}

// AddUint64 atomically adds delta to *addr and returns the new value.
// To subtract a signed positive constant value c from x, do AddUint64(&x, ^uint64(c-1)).
// In particular, to decrement x, do AddUint64(&x, ^uint64(0))
