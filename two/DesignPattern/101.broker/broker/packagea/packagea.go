package packagea

import (
	"fmt"

	"github.com/jimweng/DesignPattern/101.broker/registry"
)

type Delegate struct{}

func (d *Delegate) Execute() {
	// do something
	fmt.Println("package A was executed")
}

// another method to add Delegate with registry
func init() {
	registry.Add("packagea", func() registry.Register {
		return &Delegate{}
	})
}
