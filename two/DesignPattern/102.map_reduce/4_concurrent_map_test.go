package main

import (
	"errors"
	"io"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcurrentMap(t *testing.T) {
	results2 := outputConsumer2{}
	returnArr, err := ConcurrentMapPoorErrorHandling(&intProducer{data: []int{1, 2, 3, 4, 5, 6, 7}}, &results2, func(x interface{}) (interface{}, error) {
		if i, ok := x.(int); ok {
			return strconv.FormatInt(int64(i), 2), nil
		}
		return nil, errors.New("lambda: not an int")
	})
	assert.Nil(t, err)
	// assert.Equal(t, []interface{}{"111", "10", "101", "100", "11", "110", "1"}, returnArr)
	assert.NotEmpty(t, returnArr)
}

func ConcurrentMapPoorErrorHandling(p genericProducer, c genericConsumer, mapper genericMapper) ([]string, error) {
	// var results []interface{}
	results := make([]string, 0)
	count := 0
	// empty struct{} is a type in Go. You can also redefine the type as:
	// type DoneSignal struct{}
	done := make(chan struct{})
	for {
		next, err := p.Next()
		if err != nil {
			if err == io.EOF {
				break // There is no more elements in the producer.
			}
			return nil, err // There is an error in the producer. Shut down the mapping.
		}

		/*
			使用 count 這個計數器來告知後面要從 `done` 這個 chan 裡面取幾次值
		*/
		count++

		// 使用 goroutine 將 `next` 平行放入 c.Send(ele) 裡面
		go func(next interface{}) {
			ele, err := mapper(next)
			if err != nil {
				panic(err)
			}
			err = c.Send(ele)
			if err != nil {
				panic(err)
			}

			results = append(results, ele.(string))

			// if i, ok := ele.(string); ok {
			// 	results = append(results, i)
			// }

			done <- struct{}{}
		}(next)
	}

	// for loop 依序迭代將done拉出，同時用作堵塞，避免程序結束
	for i := 0; i < count; i++ {
		<-done
	}

	return results, nil
}
