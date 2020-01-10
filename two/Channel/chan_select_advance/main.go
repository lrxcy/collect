package main

import (
	"fmt"
	"time"
)

var someCondition = true

func main() {

	c1 := make(chan string, 1)
	if someCondition == true {
		fmt.Println("Enter here.")
		go func() {
			time.Sleep(time.Second * 2)
			c1 <- "result 1"
		}()
	}

	select {
	case x, ok := <-c1: //如果someCondition == true 除非這時候剛好得到結果，不然跑不到．
		if ok {
			fmt.Printf("Value %d was read.\n", x)

		} else {
			fmt.Println("Channel closed!") //Channel 被close.
		}
	default:
		fmt.Println("No value ready, moving on.") //Channel 沒有設定過，會馬上離開....
	}
}
