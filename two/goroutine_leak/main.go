package main

import (
	"fmt"
	"runtime"
	"time"
)

func workerpool(task <-chan string, result chan<- string, id int) {
	for t := range task {
		fmt.Printf("worker %d execute job %v\n", id, t)
		result <- fmt.Sprintf("worker %d finished job\n", id)
	}

}

func main() {
	go checkGoroutine()
	taskq := make(chan string, 5)
	respq := make(chan string, 5)

	time.Sleep(time.Second * 5)
	// init workers
	for i := 0; i < 3; i++ {
		go workerpool(taskq, respq, i)
	}

	for i := 0; i < 5; i++ {
		taskq <- fmt.Sprintf("---job %d", i)
	}

	for i := 0; i < 5; i++ {
		<-respq
	}
	close(respq)

	// for {
	// 	select {
	// 	case res := <-respq:
	// 		fmt.Println(res)
	// 	default:
	// 		// fmt.Println("error")
	// 	}
	// }

	time.Sleep(time.Second * 10)

}

func checkGoroutine() {
	for {
		n := runtime.NumGoroutine()
		time.Sleep(time.Second * 1)
		fmt.Println(n)
	}
}
