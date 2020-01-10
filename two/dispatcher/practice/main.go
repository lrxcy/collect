package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jimweng/dispatcher/practice/collector"
	"github.com/jimweng/dispatcher/practice/dispatcher"
)

var (
	workerNum = flag.Int("worker", 4, "showthe number of workers.")
	httpAddr  = flag.String("addr", "127.0.0.1:8001", "specific the http address.")
)

func main() {
	flag.Parse()
	// Start the dispatcher.
	fmt.Println("Starting the dispatcher")
	dispatcher.StartDispatcher(*workerNum)

	// Register our collector as an HTTP handler function.
	fmt.Println("Registering the collector")
	http.HandleFunc("/work", collector.Collector)

	// Start the HTTP server!
	go ServerRun(*httpAddr)

	// create one chan to print awaiting signal on console
	sigs := make(chan os.Signal, 1)

	// create another chan to receive signal to interrupt original chan
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Println(sig)

		dispatcher.StopWorker()
		time.Sleep(time.Duration(2) * time.Second)

		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

}

func ServerRun(addr string) {
	fmt.Println("HTTP server listening on", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println(err.Error())
	}
}
