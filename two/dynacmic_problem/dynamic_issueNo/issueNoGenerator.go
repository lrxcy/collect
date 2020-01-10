package main

import (
	"fmt"
	"time"
)

var (
	staticIssueNo1 string
)

func main() {
	backgroundIssueNoIncrement()
	if 11 != len(staticIssueNo1) {
		fmt.Println(len(staticIssueNo1), "____", staticIssueNo1)
		fmt.Println("Invalid IssueNo")
	} else {
		fmt.Println("legal IssueNo")
	}
}

func backgroundIssueNoIncrement() {
	staticIssueNo1 = "2019" + getMonth() + getDay() + getIssueNo()
}

func getMonth() string {
	var monthString string
	m := time.Now().Month()
	if int(m) < 10 {
		monthString = fmt.Sprintf("0%d", int(m))
	} else {
		monthString = fmt.Sprintf("%d", int(m))
	}

	return monthString
}

func getDay() string {
	var dayString string
	d := time.Now().Day()
	if int(d) < 10 {
		dayString = fmt.Sprintf("0%d", int(d))
	} else {
		dayString = fmt.Sprintf("%d", int(d))
	}
	return dayString
}

func getIssueNo() string {
	return "123"
}
