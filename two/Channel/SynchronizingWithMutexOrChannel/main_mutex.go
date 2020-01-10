package main

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}
var wg sync.WaitGroup

func MutexGenerateMessage(message int) {
	//...

	wg.Add(1)
	go func() {
		defer wg.Done()
		MutexPrintMessage(message)
	}()

	wg.Wait()
}

func MutexPrintMessage(resultMessage int) {
	mutex.Lock()
	defer mutex.Unlock()

	// Internal Logic : Ignore
	// if IsTimeEnabled {
	// 	log.Println(resultMessage)
	// 	return
	// }
	fmt.Println(resultMessage)
}

func main() {
	for i := 0; i < 100; i++ {
		MutexGenerateMessage(i)
	}
}
