// URLs provide a uniform way to locate resources,
// and this example would demo how to use 'url'
// https://gobyexample.com/url-parsing
package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// This example url would include a scheme, authentication info, host, path, query params, and query fragment
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parse the URL and ensure there are no errors
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// print out "u"
	fmt.Println("---------u  &  u.Scheme-------------")
	fmt.Println(u)
	fmt.Println("u.Scheme = ", u.Scheme)
	fmt.Println("---------u.User()  &  u.Password()-------------")

	fmt.Println("u.User = ", u.User)
	fmt.Println("u.User.Username() = ", u.User.Username())

	p, _ := u.User.Password()
	fmt.Println("u.User.Password() = ", p)

	// below would print out context detail
	fmt.Println("---------u.Host with net.SplitHostPort-------------")
	// The Host contains both the hostname and the port
	// use SplitHostPort to separate 'host' and 'port'
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("u.Host = ", u.Host)
	fmt.Println("host = ", host)
	fmt.Println("port = ", port)

	fmt.Println("---------u.Path and u.Fragment-------------")
	// print out Path and Fragment
	// Here we can extract the path and the fragment after the '#'
	fmt.Println("u.Path = ", u.Path)
	fmt.Println("u.Fragment = ", u.Fragment)

	fmt.Println("---------u.RawQuery and context-------------")
	// To get query parames in a string of k=v format, use RawQuery
	fmt.Println("u.RawQuery is ", u.RawQuery)

	// Parse a query params into a map
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("url.ParseQuery(u.RawQuery) = ", m)

	// Ths m["k"][0] would be the first value from the parased query params map
	fmt.Println("map[k:[v]] value is ", m["k"][0])

}
