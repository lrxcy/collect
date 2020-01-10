package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		考慮一般單線程，處理一個rate limit的情境
	*/
	// 模擬假設有5個請求近來，會依序放入channel
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 使用Trick每200毫秒就會回傳一個值來做限速
	limiter := time.Tick(200 * time.Millisecond)

	// 藉由上面的limiter來達到限速
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	/*
		有時，可能需要調節速限
		透過burstyLimiter實作一個buffer channel
		允許每次處理最多三個goroutine
	*/
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 因為burstyLimiter最多只能塞入三筆數據，所以請求一次最多同時處理三筆
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// 一樣假設有五個請求進入
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// 同剛剛的拿取方法，只是這次的burstyLimiter是同時三個goroutine在做事，所以速度理應會比之前快三倍
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
