package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/jimweng/dispatcher/tagWorkersPool/utils"
)

func Collector(c *gin.Context) {
	doFunc := func(config interface{}) (string, error) {
		log.Println("Doing shift")
		return fmt.Sprintf("Doing shift"), nil
	}

	workTask := WorkRequest{
		Execute: doFunc,
		Result:  make(chan interface{}), // 如果不生成實例，會導致阻塞
	}

	// retrive worker from worker pools
	Wspools.RetriveWorkerQueue() <- workTask

	resp := <-workTask.Result

	c.JSON(http.StatusOK, fmt.Sprintf("%v\n", resp))
}
