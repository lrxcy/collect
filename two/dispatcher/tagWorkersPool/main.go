package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/jimweng/dispatcher/tagWorkersPool/internal"
	"github.com/jimweng/dispatcher/tagWorkersPool/router"
)

var (
	HTTPAddr = flag.String("http", "127.0.0.1:9090", "Address to listen for HTTP requests on")
)

func main() {
	flag.Parse()

	route := gin.Default()
	router.ApiRouter(route)

	httpSrv := &http.Server{
		Addr:    *HTTPAddr,
		Handler: route,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("http listen : %v\n", err)
			panic(err)
		}
	}()
	gracefulShutdown()

}

// gracefulShutdown: handle the worker connection
func gracefulShutdown() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		internal.StopDispatcher()
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
