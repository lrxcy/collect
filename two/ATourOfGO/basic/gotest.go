package main

import (
	"flag"
	"fmt"
	"strings"
)

var fDebug = flag.Bool("debug", true, "show debug mode help")
var fString = flag.String("name", "Jim", "default name would be Jim")

type test struct {
	Debug bool
}

var test_array []test

func testInit() {
	test_array = []test{
		test{
			Debug: true,
		},
		test{
			Debug: false,
		},
	}
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// go run gotest.go -debug=false
func main() {
	var (
		test  = 1
		test2 = 2
	)

	fmt.Println(test, test2)

	flag.Parse()
	fmt.Println(*fDebug)
	if *fDebug {
		fmt.Println("it's right")
	}

	fmt.Println(*fString)

	list := []string{"jim", "bob", "beck"}
	if StringInSlice(*fString, list) {
		fmt.Println("true")
	}

	testInit()
	fmt.Println(test_array)

	// var testString = "1,2,3,4,5,6"
	var testString = "mem,host=MacPro used_percent=57.654523849487305,total=8589934592i,cached=0i,inactive=3101143040i,slab=0i,active=3521380352i,available_percent=42.345476150512695,available=3637448704i,used=4952485888i,free=536305664i,buffered=0i 1513149940000000000\ndisk,path=/,device=disk1s1,fstype=apfs,mode=rw,host=MacPro inodes_free=9223372036853631618i,inodes_used=1144189i,total=119824367616i,free=71182340096i,used=42605252608i,used_percent=37.44279283491888,inodes_total=9223372036854775807i 1513149940000000000"
	// fmt.Println(testString)
	// var qiReplacer = strings.NewReplacer(",", " ", "=", " ")
	fmt.Println("After split it would be:")
	var splitString = strings.Split(testString, "\n")
	// var splitString = strings.Split(splitString1, "=")
	// splitString := qiReplacer.Replace(testString)
	// fmt.Println(splitString)
	fmt.Println("below would demo the length of splitString and show elements of splitString:")
	// fmt.Println(splitString[0], "\n", splitString[1])
	fmt.Println(len(splitString))
	testString2 := splitString[0]
	splitString2 := strings.Split(testString2, ",")
	fmt.Println(splitString2[0], "\n", splitString2[1])
}
