package plugin2

import (
	"fmt"

	"github.com/jimweng/UnmarshalJson/plugins"
)

type DemoStruct struct {
	Test string
}

func (d *DemoStruct) Gather() error {
	fmt.Println("Enther here2")
	fmt.Printf("%v\n", d.Test)
	return nil
}

func init() {
	plugins.Add("plugin2", func() plugins.Plugin {
		return &DemoStruct{}
	})
}
