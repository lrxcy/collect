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

var myClient *http.Client

func startWebserver() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	go http.ListenAndServe(":9999", nil)

}

func startLoadTest() {
	count := 0
	for {
		resp, err := myClient.Get("http://localhost:9999/")
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

	// Customize the Transport to have larger connection pool
	defaultRoundTripper := http.DefaultTransport
	defaultTransportPointer, ok := defaultRoundTripper.(*http.Transport)
	if !ok {
		panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport"))
	}
	defaultTransport := *defaultTransportPointer // dereference it to get a copy of the struct that the pointer points to
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100

	myClient = &http.Client{Transport: &defaultTransport}

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
描述Golang預設的HTTP Client的行為


By default, the Golang HTTP client will do connectin pooling. Rather than closing a socket
connection after an HTTP request, it will add it to an idle connection pool.

And if we try to make another HTTP request before th idle connection timeout (90 econds by default)
,then it will re-use that existing connetion rather than creating a new one.

This will keep the number total socket connections low, as long as the pool doesn't fill up.
If the pool is full of established socket connections, then it will just create a new socket
connection for HTTP request and use that

example of golang http client example

var DefaultTransport RoundTripper = &Transport{
        ...
  MaxIdleConns:          100,
  IdleConnTimeout:       90 * time.Second,
        ...
}

// DefaultMaxIdleConnsPerHost is the default value of Transport's
// MaxIdleConnsPerHost.
const DefaultMaxIdleConnsPerHost = 2

*/

/*
名詞解釋:


MaxIdleConns: 100
setting sets the size of the connection pool to 100 connections, but with one major caveat
> this is on a per-host basis.

IdleConnTimeout: 90 * time.Second
meaning that after a connection stays in the poll and is unused for 90 seconds,
it will be removed from the pool and closed

defaultMaxIdleConnsPerHos = 2
setting below it. What this means is that even though the ntire connection pool
is set to 100, there is a per-host cap of only 2 connectsions!
*/

/*
In the above xample, there are 100 go routines trying to concurrently make requests to
the same host, but the connection pool an nly hold 2 sockets.

So in the first "round" of the goroutines finishing their http request,
2 of the sockets will remain open in the pool, while the remaining 98 connections
will be closed and end up in the TIME_WAIT state.


Since this is happening in a loop, we will quickly accumulate lots of connections
in the TIME_WAIT state. Eventually, for that particular host at least,
it will run out of ephemeral ports and not be able to open new client connections
for a load testing too,this is bad news.

*/
