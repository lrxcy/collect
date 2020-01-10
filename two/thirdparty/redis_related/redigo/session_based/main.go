package main

import (
	"log"
	"net/http"

	"github.com/jimweng/thirdparty/redis_related/redigo/session_based/handler"
	"github.com/jimweng/thirdparty/redis_related/redigo/session_based/redispool"
)

func init() {
	redispool.InitCache()
}

func main() {
	http.HandleFunc("/signin", handler.Signin)
	http.HandleFunc("/welcome", handler.Welcome)
	http.HandleFunc("/refresh", handler.Refresh)

	log.Fatal(http.ListenAndServe(":1234", nil))

}
