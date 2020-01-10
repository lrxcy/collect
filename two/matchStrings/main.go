package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var rgex = regexp.MustCompile(`\d+[s,m,h]{1}`)
var rgunit = regexp.MustCompile(`[s,m,h]{1}`)

func main() {
	hrMinSec := []string{"1h", "59m,", "40s"}
	for _, j := range hrMinSec {
		// rs := rgex.FindString(j)
		rs := rgunit.FindString(j)
		// fmt.Println(rs)

		switch rs {
		case "h":
			unit := strings.Split(j, rs)[0]
			// func strconv.ParseInt(s string, base int, bitSize int) (i int64, err error)
			timeStamp, err := strconv.ParseInt(unit, 10, 64)
			if err != nil {
				panic(err)
			}
			fmt.Println(timeStamp * 3600)

		case "m":
			unit := strings.Split(j, rs)[0]
			// func strconv.ParseInt(s string, base int, bitSize int) (i int64, err error)
			timeStamp, err := strconv.ParseInt(unit, 10, 64)
			if err != nil {
				panic(err)
			}
			fmt.Println(timeStamp * 60)

		case "s":
			unit := strings.Split(j, rs)[0]
			// func strconv.ParseInt(s string, base int, bitSize int) (i int64, err error)
			timeStamp, err := strconv.ParseInt(unit, 10, 64)
			if err != nil {
				panic(err)
			}
			fmt.Println(timeStamp)
		}

	}
}
