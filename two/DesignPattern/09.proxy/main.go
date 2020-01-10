package main

import "log"

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	var res string

	// 在調用真實對象之前的工作 e.g. 檢查緩存，判斷權限，實例化真實對象... 等
	res += "pre:"

	// 調用真實對象
	res += p.real.Do()

	// 調用之後的操作，如緩存結果，對結果進行處理... 等
	res += ":after"

	return res
}

func main() {
	var sub Subject
	sub = &Proxy{}

	res := sub.Do()
	log.Println(res)
}
