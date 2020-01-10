package main

import (
	"fmt"
	"net/url"
	"path"
)

type httpClient struct {
	writeURL string
	// config   HTTPConfig
	// client   *http.Client
	url *url.URL
}

func queryURL(u *url.URL, command string) string {
	params := url.Values{}
	params.Set("q", command)

	u.RawQuery = params.Encode()
	p := u.Path
	u.Path = path.Join(p, "query")
	// u.Path = path.Join(p, "db/data/cypher")  //neo4j
	s := u.String()
	u.Path = p
	return s
}

func main() {
	test := "http://172.31.86.190:7474/db/data"
	testurl, _ := url.Parse(test)
	fmt.Println(queryURL(testurl, "1"))
	fmt.Println(testurl)

}
