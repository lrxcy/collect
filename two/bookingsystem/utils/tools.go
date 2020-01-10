package utils

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func WebResp(c *gin.Context, statusCode, code int, Message string, data interface{}) {
	respMap := map[string]interface{}{
		"Code":    code,
		"Message": Message,
		"Data": map[string]interface{}{
			"content": data,
		},
	}
	c.JSON(statusCode, respMap)
	return
}

// gracefulShutdown: handle the worker connection
func GracefulShutdown(v func()) {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		v()
		fmt.Println(sig)
		done <- true
	}()

	// Log.Info("awaiting signal")
	log.Println("awaiting signal")
	<-done
	// Log.Info("exiting")
	log.Println("exiting")
}
