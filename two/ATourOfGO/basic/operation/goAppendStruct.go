package main

import "fmt"

type tStruct struct {
	str string
	i   int
}

func main() {
	measOfBlockMetricArray := map[string][][]tStruct{}

	a := tStruct{
		str: "str",
		i:   123,
	}

	blockMetrics := []tStruct{a}
	fmt.Println("blockMetrics", blockMetrics)

	measOfBlockMetricArray["jim"] = append(measOfBlockMetricArray["jim"], blockMetrics)
	measOfBlockMetricArray["beck"] = append(measOfBlockMetricArray["beck"], blockMetrics)

	fmt.Println(measOfBlockMetricArray)

	m, exist := measOfBlockMetricArray["jim"]
	fmt.Println("jim content", m)
	fmt.Println("exist", exist)
}
