package main

import "fmt"

func main() {
	results, err := MapReduce([]int{1, 2, 3}, func(x int) (int, error) {
		y := x + 1
		return y, nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(results)

}

type mapper func(int) (int, error)

func MapReduce(d []int, m mapper) ([]int, error) {
	var md []int
	for _, i := range d {
		if v, err := m(i); err != nil {
			return nil, err
		} else {
			md = append(md, v)
		}
	}
	return md, nil
}
