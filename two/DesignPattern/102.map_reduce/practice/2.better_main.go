package main

import (
	"errors"
	"fmt"
	"io"
)

type producer interface {
	Next() (interface{}, error)
}

type intStruct struct {
	d     []int
	index int
}

// func (ip *intStruct) Next() (interface{}, error) {
// 	return nil, nil
// }

func (ip *intStruct) Next() (interface{}, error) {
	if ip.index < len(ip.d) {
		defer func() { ip.index++ }()
		return ip.d[ip.index], nil
	}
	return nil, io.EOF
}

type bettermapper func(interface{}) (interface{}, error)

func main() {
	results, err := betterMap(&intStruct{d: []int{1, 2, 3}}, func(i interface{}) (interface{}, error) {
		if v, ok := i.(int); ok {
			return v, nil
		}
		return nil, errors.New("invalid type in mapper input")
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(results)
}

// func betterMap(p producer, m bettermapper) ([]interface{}, error) {
// 	return nil, nil
// }

func betterMap(p producer, m bettermapper) ([]interface{}, error) {
	returnMap := make([]interface{}, 0)
loop:
	for {
		v, err := p.Next()
		switch err {
		case nil:
			if mv, err := m(v); err != nil {
				return nil, err
			} else {
				returnMap = append(returnMap, mv)
			}

		case io.EOF:
			break loop

		default:
			return nil, err
		}
	}

	return returnMap, nil
}

/*
	[How To break for loop in switch case] Refer:
	https://www.golangprograms.com/example-of-switch-case-with-break-in-for-loop.html
*/
