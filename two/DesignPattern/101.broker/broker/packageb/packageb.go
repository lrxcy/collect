package packageb

import (
	"fmt"

	"github.com/jimweng/DesignPattern/101.broker/registry"
)

type Delegate struct{}

func (d *Delegate) Execute() {
	fmt.Println("package B was executed")
}

// another method to add Delegate with registry
func init() {
	registry.Add("packageb", func() registry.Register {
		return &Delegate{}
	})
}
