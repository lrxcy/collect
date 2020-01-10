package plugin1

import (
	"fmt"

	"github.com/jimweng/UnmarshalJson/plugins"
)

type DemoStruct struct {
	Test interface{}
}

func (d *DemoStruct) Gather() error {
	fmt.Println("Enther here1")
	fmt.Printf("%v\n", d.Test)
	return nil
}

func init() {
	plugins.Add("plugin1", func() plugins.Plugin {
		return &DemoStruct{}
	})
}
