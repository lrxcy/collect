package internal

import (
	"fmt"
	. "github.com/jimweng/dispatcher/tagWorkersPool/utils"
)

/*
	is it really need to use singleton ?
*/

/*
type WorkerPool struct {
	ws controller.WorkerPools
}

var wspool *WorkerPool

var once sync.Once


func GetWorkerPool() *WorkerPool {
	once.Do(func() {
		wspool = &WorkerPool{}
	})
	return wspool
}
*/

// var wspool controller.WorkerPools

func init() {
	// wspool := GetWorkerPool()
	fmt.Println("Starting the dispatcher")
	// wspool.ws = controller.NewDispatcherPool(4, "jim1")
	Wspools = NewDispatcherPool(2, "jim1")
	// wspool.ws.StartDispatcher()
	Wspools.StartDispatcher()
}

func StopDispatcher() {
	// wspool := GetWorkerPool()
	// wspool.ws.StopDispatcher()
	Wspools.StopDispatcher()
}
