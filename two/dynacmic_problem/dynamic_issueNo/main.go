package main

import (
	"errors"
	"fmt"
	"time"
)

const dynamicIssueNo = 0

var dynamicIssueNoMap map[string]checker

type checker interface {
	validIssueNo(string) bool
}

func validateIssueNoSize(issueNoSize int, issueNo string, lotteryCode string) bool {
	switch issueNoSize {
	case dynamicIssueNo:
		return dynamicIssueNoMap[lotteryCode].validIssueNo(issueNo)
	default:
		return len(issueNo) == issueNoSize
	}
}

type DynamicIssueNoChecker struct {
	acceptSize []int // fequently resize the array to make sure keep it works well
	stableSize int
	freqency   time.Duration
}

func NewDynamicIssueNoChecker(lotteryCode string) *DynamicIssueNoChecker {
	return &DynamicIssueNoChecker{
		acceptSize: make([]int, 0),
		freqency:   time.Second,
	}
}

// if the acceptSize is on
func (d *DynamicIssueNoChecker) allTheMapSame() bool {
	m := make(map[int]int, 0)

	for i, j := range d.acceptSize {
		m[j] = i
	}

	return len(m) == 1
}

func (d *DynamicIssueNoChecker) ValidatedIssueNo(issueNo int) bool {
	if d.allTheMapSame() {
		return d.stableSize == issueNo
	} else {
		for _, j := range d.acceptSize {
			if j == issueNo {
				return true
			}
		}
		return false
	}
}

func (d *DynamicIssueNoChecker) retriveStableIssueNoSzie() int {
	return d.stableSize
}

func (d *DynamicIssueNoChecker) crontabCheck(issueNo int) {
	for {
		if d.stableSize > 1 {
			for _, j := range d.acceptSize {
				// callback acceptSize and descide whether to resize the array
				if j != d.stableSize {
					fmt.Println(j)
				}
			}
		}

		time.Sleep(d.freqency)
	}
}

func main() {
	var a interface{} = 1
	if ar, ok := a.([]int); ok {
		for _, j := range ar {
			fmt.Println(j)
		}
	}

	fmt.Println("test")
	if !validateIssueNoSize(2, "2") {
		panic(errors.New("issueNo len is invalid"))
	}
}
