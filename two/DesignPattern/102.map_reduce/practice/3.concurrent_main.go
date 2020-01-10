package main

import (
	"errors"
	"fmt"
	"io"
	"log"
)

type concurrentProducer interface {
	Next() (interface{}, error)
}

type producerStruct struct {
	id    []interface{}
	index int
}

func (ip *producerStruct) Next() (interface{}, error) {
	if ip.index < len(ip.id) {
		defer func() { ip.index++ }()
		return ip.id[ip.index], nil
	}
	return nil, io.EOF
}

func main() {
	results, err := concurrentPoorErrorMap(&producerStruct{id: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, func(i interface{}) (interface{}, error) {
		if v, ok := i.(int); ok {
			return v, nil
		}
		return nil, errors.New("Invalid type was included in producer")
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(results)
}

func concurrentPoorErrorMap(p producer, m bettermapper) ([]interface{}, error) {
	returnMap := make([]interface{}, 0)
	count := 0
	done := make(chan struct{})

loop:
	for {
		v, err := p.Next()
		switch err {
		case nil:
			count++

			go func(v interface{}) {
				if mv, err := m(v); err != nil {
					panic(err) // 在goroutine中無法返回錯誤...
				} else {
					log.Println(mv) // 因為這邊已經有race condition出現了，所以必須用log來打印出mapper過後的value，否則會有漏數字的可能性
					returnMap = append(returnMap, mv)
					done <- struct{}{}
				}
			}(v)

		case io.EOF:
			break loop

		default:
			continue
		}

	}

	for i := 0; i < count; i++ {
		<-done
	}

	return returnMap, nil
}

// ----- inherent from above example ---

type producer interface {
	Next() (interface{}, error)
}

type bettermapper func(interface{}) (interface{}, error)
