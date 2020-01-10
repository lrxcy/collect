package main

import (
	"log"
	"net/http"
	"net/http/cgi"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler := new(cgi.Handler)
		handler.Path = "/Users/jimweng/go/src/github.com/jimweng/CGI"
		script := "/Users/jimweng/go/src/github.com/jimweng/CGI/cgi-script/" + r.URL.Path
		log.Println(handler.Path)
		handler.Dir = "/Users/jimweng/go/src/github.com/jimweng/CGI/cgi-script/"
		args := []string{"run", script}
		handler.Args = append(handler.Args, args...)
		handler.Env = append(handler.Env, "GOPATH/Users/jimweng/go")
		handler.Env = append(handler.Env, "GOROOT=/usr/local/Cellar/go/1.13.4/libexec")
		log.Println(handler.Args)

		handler.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":8989", nil))
}
