package main

import (
	"flag"
	"fmt"
)

// claim flag with types ; String(), Bool(), Int()
var ip = flag.String("testIP", "127.0.0.1", "message")
var criterion = flag.Bool("debug", false, "wheter it's stdout")

// through IntVar to connect flag to a particular variable
var flagvar int

func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}

// execute with command as below
// go run flagPointer.go -testIP 172.31.17.111 -flagname 9527 -debug true
func main() {
	flag.Parse()
	fmt.Println("ip var use address", ip)
	fmt.Println("IP has value", *ip)

	if !*criterion {
		fmt.Println("use *criterion to log value and ! as a negative note \n and also printt out testIP", *ip)
	}

	fmt.Println("flagname is :", flagvar)

}
