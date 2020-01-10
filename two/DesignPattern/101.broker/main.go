package main

import (
	"fmt"

	"github.com/jimweng/DesignPattern/101.broker/broker"
	"github.com/jimweng/DesignPattern/101.broker/registry"

	_ "github.com/jimweng/DesignPattern/101.broker/broker/all"
)

var bk = []broker.Broker{broker.NewA(), broker.NewB()}

func main() {
	// call registry to print out result
	for i, j := range registry.PackageManager {
		executor := j()
		fmt.Printf("%v_____\n", i)
		executor.Execute()
		// j().Execute()
	}
}

func init() {
	for _, j := range bk {
		j.Execute()
	}
}
