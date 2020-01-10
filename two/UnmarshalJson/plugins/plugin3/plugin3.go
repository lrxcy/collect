package plugin3

import (
	"fmt"

	"github.com/jimweng/UnmarshalJson/plugins"
)

type DemoStruct struct {
	Test string
}

func (d *DemoStruct) Gather() error {
	fmt.Println("Enther here3")
	fmt.Printf("%v\n", d.Test)
	return nil
}

func init() {
	plugins.Add("plugin3", func() plugins.Plugin {
		return &DemoStruct{}
	})
}
