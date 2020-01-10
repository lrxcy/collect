package main

import "time"
import "fmt"

func main() {

	c1 := make(chan string, 1)

	someCondition := true
	if someCondition == true {
		go func() {
			time.Sleep(time.Second * 4)
			c1 <- "result 1"
		}()
	}

	c2 := make(chan string, 1)
	if someCondition == true {
		go func() {
			//故意延遲九秒，所以這個不會順利結束．
			time.Sleep(time.Second * 9)
			c2 <- "result 1"
		}()
	}

	doneCount := 0
	allDone := 2
	timeCount := 0

	// 別忘了... select 滿足任何一個都會離開，所以要有個for在外面讓他不停跑
	for doneCount < allDone && timeCount < 5 {

		select {
		//檢查C1
		case x, ok := <-c1:
			if ok {
				fmt.Printf("C1 in valus is %s.\n", x)
				doneCount++
			} else {
				fmt.Println("Channel closed!") //Channel 被close.
			}
		//檢查C2
		case x, _ := <-c2:
			fmt.Printf("C2 in valus is %s.\n", x)
			doneCount++

			//另外準備一個離開條件，當五秒會離開...
		case <-time.After(time.Second * 1):
			fmt.Println("tick..")
			timeCount++

		}
	}
}
