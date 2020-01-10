package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker %v started jobs _%v\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker %v finished jobs _%v\n", id, j)
		results <- j * 2
	}
}

func main() {
	// 宣告一個jobs的channel來分派任務給worker
	jobs := make(chan int, 100)
	// 宣告一個results channel來接完成的工作
	results := make(chan int, 100)

	// 啟動worker: 這邊啟動三個worker，每個worker都會在背景執行，聽jobs的工作。
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 塞任務到jobs裡面
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	// 當josb裡面的資料塞完後，關閉jobs這個channel
	close(jobs)

	// 使用 `<-results` 做阻塞，等results拿完5次就跳出迴圈
	for a := 1; a <= 5; a++ {
		<-results
	}
}
