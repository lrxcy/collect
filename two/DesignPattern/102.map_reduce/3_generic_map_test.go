package main

import (
	"errors"
	"io"
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type intProducer struct {
	index int
	data  []int
}

type outputConsumer2 struct{}

func TestGenericMap(t *testing.T) {
	strarr, err := GenericMap(&intProducer{data: []int{10, 11, 12}}, outputConsumer2{}, func(ele interface{}) (interface{}, error) {
		if i, ok := ele.(int); ok {
			return strconv.FormatInt(int64(i), 16), nil
		}
		return nil, errors.New("mapper: not an int")
	})
	assert.Nil(t, err)
	assert.Equal(t, []string{"a", "b", "c"}, strarr)
}

// // GenericMap 為預計的映射map reduce 函數
// func GenericMap(sp *intProducer, oc outputConsumer2, f func(interface{}) (interface{}, error)) ([]int64, error) {
// 	return nil, nil
// }

/*
	對 `intProducer` 以及 `outputConsumer2` 做一層抽象
*/

/*
	GenericMap 同前面的BetterMap，迭代producer裡面的data後，再依序取值。
	將取出的值，依序送到consumer裏打印後在append到results後返回
*/
func GenericMap(p genericProducer, c genericConsumer, m genericMapper) ([]string, error) {
	results := make([]string, 0)
	for {
		next, err := p.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		datum, err := m(next)
		if err != nil {
			return nil, err
		}
		if err := c.Send(datum); err != nil {
			return nil, err
		}

		if i, ok := datum.(string); ok {
			results = append(results, i)
		} else {
			return nil, errors.New("Invalid data type is included.")
		}
	}
	return results, nil
}

type genericProducer interface {
	Next() (interface{}, error)
}

func (ip *intProducer) Next() (interface{}, error) {
	if ip.index < len(ip.data) {
		defer func() { ip.index++ }()
		return ip.data[ip.index], nil
	}
	return nil, io.EOF
}

type genericConsumer interface {
	Send(interface{}) error
}

func (c outputConsumer2) Send(ele interface{}) error {
	log.Printf("output consumer2 : %v\n", ele)
	return nil
}

type genericMapper func(interface{}) (interface{}, error)
