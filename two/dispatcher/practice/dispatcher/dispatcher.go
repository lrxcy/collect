package dispatcher

import (
	"fmt"

	"github.com/jimweng/dispatcher/practice/worker"
)

func StartDispatcher(nworkers int) {
	worker.WorkerPool = make(worker.WorkerPoolType, nworkers)

	for i := 0; i < nworkers; i++ {
		fmt.Println("Starting worker", i+1)
		worker.NewWorker(i+1, worker.WorkerPool).Start()
	}

	go func() {
		for {
			select {
			case work := <-worker.WorkQueue:
				fmt.Println("Received work requeust")
				go func() {
					workerPool := <-worker.WorkerPool

					fmt.Println("Dispatching work request")
					workerPool <- work

				}()
			}
		}
	}()
}

func StopWorker() {
	for _, j := range worker.Workers {
		j.Stop()
	}
}
