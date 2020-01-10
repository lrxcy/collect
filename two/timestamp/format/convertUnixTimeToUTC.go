package main

import (
	"fmt"
	"strconv"
	"time"
)

var timeOrigin = "2019-12-09 16:33:41"
var layout = "2019-12-09 08:34:35"

func main() {
	i, err := strconv.ParseInt("1405544146", 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	fmt.Println(tm)
	fmt.Println(tm.Year())
	fmt.Println(int(tm.Month()))
	fmt.Println(tm.Day())
	fmt.Println(tm.Hour())
	fmt.Println(tm.Minute())
	fmt.Println(tm.Second())

	fmt.Println(tm.Format("2006/01/02 15:04:05"))

	fmt.Println(timeOrigin)

	// layout := "01/02/2006 3:04:05 PM"
	// t, err := time.Parse(layout, "02/28/2016 9:03:46 PM")
	t, err := time.Parse("2006-01-15 15:04:05", layout)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Unix())
}

// 2019-12-09 16:33:41
