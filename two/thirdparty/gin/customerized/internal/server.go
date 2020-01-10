package internal

import (
	. "github.com/jimweng/thirdparty/gin/customerized/logger"
	. "github.com/jimweng/thirdparty/gin/customerized/utils"
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

func InitWorkerPools(tag string, n int) {
	// wspool := GetWorkerPool()
	Log.Info("Starting the Workers Pool")
	Wspools = NewDispatcherPool(n, tag)
	// wspool.ws.StartDispatcher()
	Wspools.StartDispatcher()
}

func StopDispatcher() {
	// wspool := GetWorkerPool()
	// wspool.ws.StopDispatcher()
	Wspools.StopDispatcher()
}
