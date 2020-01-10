package worker

import "fmt"

var (
	WorkQueue  = make(chan WorkRequest, 100)
	WorkerPool WorkerPoolType
	Workers    []*Worker
)

type WorkRequest struct {
	Execute func(config interface{}) error
}

type Worker struct {
	ID         int
	Work       chan WorkRequest
	WorkerPool chan chan WorkRequest
	QuitChan   chan bool
}

type WorkerPoolType chan chan WorkRequest

func NewWorker(id int, workerQueue chan chan WorkRequest) *Worker {
	worker := &Worker{
		ID:         id,
		Work:       make(chan WorkRequest),
		WorkerPool: workerQueue,
		QuitChan:   make(chan bool),
	}
	Workers = append(Workers, worker)

	return worker
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.Work

			select {
			case work := <-w.Work:
				switch w.ID {
				case 1:
					work.Execute(fmt.Sprintf("Request go to worker %v", w.ID))
				default:
					work.Execute("Request go to another worker")
				}
			case <-w.QuitChan:
				fmt.Printf("worker %d stopping\n", w.ID)
				return
			}
		}
	}()
}

// Stop tells the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}
