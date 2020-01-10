package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

var DynamicIssueNoRegister map[string]dynamicIssueNochecker

type dynamicIssueNochecker interface {
	ValidateIssueNoSize(string) bool
}

type DynamicIssueNoSizeCheckter struct {
	acceptIssueNoSize  []int
	currentIssueNoSize int
	frequency          time.Duration
}

func (d *DynamicIssueNoSizeCheckter) checkCurrentIssueNoSize() {
	d.currentIssueNoSize = len(issueNo)
	log.Printf("The current issueNo size is %d", d.currentIssueNoSize)
}

func (d *DynamicIssueNoSizeCheckter) crontabAcceptIssueNoSize() {
	if !d.CheckIntContains(len(issueNo)) {
	}
}

func (d *DynamicIssueNoSizeCheckter) CheckIntContains(i int) bool {
	for _, j := range d.acceptIssueNoSize {
		if j == i {
			return true
		}
	}
	return false
}

// 製作一個動態期號長度的驗證機器
func main() {
	go incrementIssueNo()
	for {
		fmt.Println(issueNo)
		time.Sleep(time.Second * 2)
	}
}

// -------- a lot of dummy data -----
var issueNo string

func incrementIssueNo() string {
	count := 8

	for {
		issueNo = strconv.Itoa(count)
		time.Sleep(time.Second * 1)
		count++
	}
}
