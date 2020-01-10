package internal

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jimweng/bookingsystem/models"
	"github.com/jimweng/bookingsystem/plugins"

	. "github.com/jimweng/bookingsystem/logger"
	. "github.com/jimweng/bookingsystem/utils"
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

func InitWorkerPools(tag string, n int) error {
	// wspool := GetWorkerPool()
	Log.Info("Starting the Workers Pool")
	Wspools = NewDispatcherPool(n, tag)
	// wspool.ws.StartDispatcher()
	return Wspools.StartDispatcher()
}

func StopDispatcher() {
	// wspool := GetWorkerPool()
	// wspool.ws.StopDispatcher()
	Wspools.StopDispatcher()
}

func StartAgent() {
	redisIns = models.RetriveRedisAccessModel()

	// TODO: config declare collector agent
	for w := 1; w <= 3; w++ {
		go collectorWorker(w, jobs, results)
	}

	// // TODO: config declare writor agent
	// for w := 1; w <= 3; w++ {
	// 	go writeWorker(w, jobs, results)
	// }

	go generateJob()

	// handle reuslts from worker channel
	go consumeResult()
}

var (
	redisIns models.RedisImplement
	// jobs    = make(chan interface{}, 100)
	jobs = make(chan time.Time, 100)
	// results = make(chan interface{}, 100)
	results = make(chan map[string]string, 100)
)

// result handling: map[string]string{}{"status": "success", "msg": "redis_lock_key"}
func consumeResult() {

	for {

		select {
		case r := <-results:
			switch r["status"] {
			case "success":
				log.Println("Success", r["records"])
				plugins.Outputs.Execute(r["records"])

				redisIns.SetKey(r["msg"], "finish")
				// Log.Info("Success Job and update redis key: ", r["msg"])
			default:
				// Log.Error("error occur with msg ", r["msg"])
				log.Println("Wait")
			}
		}

	}
}

func generateJob() {
	count := 1

	go func() {
		for {
			timeStamp := time.Now()
			jobs <- timeStamp
			time.Sleep(time.Second * 3)

			count++
		}
	}()
}

func writeWorker(id int, jobs <-chan time.Time, results chan<- map[string]string) {
	for j := range jobs {

		// TODO: 不要每次都進來才拿singleton...要移出去...為什麼外面打印的都是nil?
		// redisIns = models.RetriveRedisAccessModel()

		strJob := fmt.Sprintf("%v:%v:%v:%v:%v",
			j.Year(),
			j.Month(),
			j.Day(),
			j.Hour(),
			j.Minute(),
		)

		// check whether token is on redis
		if err := redisIns.SetNonExistedKey(strJob, 10); err != nil {
			// results <- strJob
			results <- map[string]string{"status": "failed", "msg": fmt.Sprintf("%v", err)}
		} else {
			// update redisKey with value
			redisIns.SetExpiredKey(strJob, "start", 10)

			// make a request
			/*
				1. agent receive strings from channel and execute input plugins
				2. agent would collect corresponding inputs plugin and encapsulate it as a task and send to output plugins
			*/

			plugins.Inputs.Execute()
			select {
			case c := <-plugins.GatherChannel:
				results <- map[string]string{"status": "success", "msg": strJob, "records": c}
			}

		}
	}
}

func collectorWorker(id int, jobs <-chan time.Time, results chan<- map[string]string) {
	for j := range jobs {

		// TODO: 不要每次都進來才拿singleton...要移出去...為什麼外面打印的都是nil?
		// redisIns = models.RetriveRedisAccessModel()

		strJob := fmt.Sprintf("%v:%v:%v:%v:%v",
			j.Year(),
			j.Month(),
			j.Day(),
			j.Hour(),
			j.Minute(),
		)

		// check whether token is on redis
		if err := redisIns.SetNonExistedKey(strJob, 10); err != nil {
			// results <- strJob
			results <- map[string]string{"status": "failed", "msg": fmt.Sprintf("%v", err)}
		} else {
			// update redisKey with value
			redisIns.SetExpiredKey(strJob, "start", 10)

			// make a request
			/*
				1. agent receive strings from channel and execute input plugins
				2. agent would collect corresponding inputs plugin and encapsulate it as a task and send to output plugins
			*/

			plugins.Inputs.Execute()
			select {
			case c := <-plugins.GatherChannel:
				results <- map[string]string{"status": "success", "msg": strJob, "records": c}
			}

		}
	}
}

func convertToSenderTime(s string, shift int) string {
	f, _ := strconv.Atoi(s)
	tm := time.Unix(int64(f-shift), 0).UTC()

	return fmt.Sprintf(tm.Format("2006/01/02 15:04:05Z"))
}
