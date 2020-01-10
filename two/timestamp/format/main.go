package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	Date_Whippletree = "2006-01-02"
	Time_Format      = "2006/01/02 15:04:05"
	// Time_Format      = "2006-01-02 15:04:05"
	Date_Format = "20060102"

	start_end_time = "00:00:00-00:00:00"
	frency         = "10m"
)

func main() {
	nowTime := time.Now()

	nowTimeDataStr := nowTime.Format(Date_Whippletree)
	fmt.Printf("The Date_Whippletree format is %v\n", nowTimeDataStr) // The Date_Whippletree format is 2019-10-16

	nowDate := nowTime.Format(Date_Format)
	fmt.Printf("The Date_Format format is %v\n", nowDate) // The Date_Format format is 20191016

	nowTimeUnix := nowTime.Unix()
	fmt.Printf("The nowTimeUnix is %v\n", nowTimeUnix) // The nowTimeUnix is 1571208968

	startTime, endTime := returnStartAndEndTime(start_end_time, "-")
	fmt.Printf("The start and end time are %v\t%v\n", startTime, endTime) // The start and end time are 00:00:00     00:00:00

	// notes: 要先組合成golang可以解析的時間格式才能近一步做parse
	timestr := nowTimeDataStr + " " + startTime
	fmt.Printf("The timer is %v\n", timestr) // 2019-10-16 00:00:00
	timer, err := parseTimeDateStr(timestr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The output timer is %v\n", timer) // The output timer is 2019-10-16 00:00:00 +0000 UTC

	// frencyUnix回傳時間(秒s), (單位數字), (錯誤)
	i, j, err := frencyUnix(frency)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The frencyUnix are %v\t%v\n", i, j) // The frencyUnix are 600  10
}

func returnStartAndEndTime(timeperiod string, symbol string) (string, string) {
	sarr := strings.Split(timeperiod, symbol)
	starttime, endtime := sarr[0], sarr[1]
	return starttime, endtime
}

func parseTimeDateStr(timestr string) (time.Time, error) {
	t, err := time.Parse(Time_Format, timestr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func frencyUnix(timeUStr string) (int64, int, error) {
	timelen := len(timeUStr)
	unit := timeUStr[timelen-1 : timelen]

	numStr := timeUStr[:timelen-1]

	var numUnix int64

	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		return 0, 0, errors.New(fmt.Sprintf("Time num %v to int64 is err : %v", numStr, err))
	}

	switch unit {
	case "s":
		numUnix = num
	case "m":
		numUnix = int64(num * 60)
	case "h":
		numUnix = int64(num * 60 * 60)
	case "d":
		numUnix = int64(num * 60 * 60 * 24)
	default:
		return 0, 0, errors.New(fmt.Sprintf("Time str uni %v is err", unit))
	}

	return numUnix, int(num), nil
}
