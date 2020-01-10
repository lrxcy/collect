package main

import (
	"fmt"
	"time"
)

var (
	dynamicIssueNo int
	ac             accept
)

func init() {
	dynamicIssueNo = 8
	ac.acceptN = append(ac.acceptN, CountDigits(dynamicIssueNo))
}

type accept struct {
	acceptN []int
}

func (a *accept) Add(i int) {
	a.acceptN = append(a.acceptN, i)
}

func main() {
	go updateLimitSize()
	go print()

	time.Sleep(10 * time.Second)

}

func updateLimitSize() {
	for {
		time.Sleep(2 * time.Second)
		dynamicIssueNo++
	}
}

func print() {
	for {
		time.Sleep(1 * time.Second)
		str := fmt.Sprintf("%d", dynamicIssueNo)
		if !ac.CheckIntContains(len(str)) {
			ac.Add(len(str))
		}
		fmt.Println(len(str), " accept lenght is ", ac)
	}
}

func CountDigits(i int) (count int) {
	for i != 0 {
		i /= 10
		count = count + 1
	}
	return count
}

func (a *accept) CheckIntContains(i int) bool {
	for _, j := range a.acceptN {
		if j == i {
			return true
		}
	}

	return false
}
