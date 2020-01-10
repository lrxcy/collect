package main

import (
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func startWebserver() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	go http.ListenAndServe(":9999", nil)

}

func startLoadTest() {
	count := 0
	for {
		resp, err := http.Get("http://localhost:9999/")
		if err != nil {
			panic(fmt.Sprintf("Got error: %v", err))
		}
		// 使用io.Copy解決請求後，響應的Body沒有正常反饋的問題
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		log.Printf("Finished GET request #%v", count)
		count += 1
	}

}

func main() {

	// start a webserver in a goroutine
	startWebserver()

	// startLoadTest()
	// use goroutine to make excessive connections in the TIME_WAIT state
	for i := 0; i < 100; i++ {
		go startLoadTest()
	}
	time.Sleep(time.Second * 2400)

}

/*
Another way to end up with excessive connections in the TIME_WAIT state is
to consistently exceed the connnection pool and cause many short-lived connections
to be opened.

Here’s some code which starts up 100 goroutines which are all trying to make requests
concurrently, and each request has a 50 ms delay:

*/
