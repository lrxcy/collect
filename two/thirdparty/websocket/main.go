package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	http.Handle("/ws", http.HandlerFunc(wsHandler))
	http.ListenAndServe(":3000", nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// websocket的Upgrade提供一個Conn: 及其方法 (c *Conn) ...
	c, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}
	go echo(c)
}

func echo(c *websocket.Conn) {
	for {
		//  WriteJson(v interface{}) error : 提供一個將msg寫進channel的方法
		if err := c.WriteJSON("hello world"); err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second)
	}
}
