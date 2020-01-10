package main

// simple.go
import (
	"fmt"

	"github.com/jimweng/ATourOfGO/VMwareGolang/telegraf_plugin_input/example"
)

// "github.com/influxdata/telegraf"
// "github.com/influxdata/telegraf/plugins/inputs"

type Simple struct {
	Ok bool
}

// func (s *Simple) Description() string {
// 	return "a demo plugin"
// }

// func (s *Simple) SampleConfig() string {
// 	return "ok = true # indicate if everything is fine"
// }

func (s *Simple) Gather(acc example.Accumulator) error {
	if s.Ok {
		fmt.Println("print something")
		acc.AddFields("state", map[string]interface{}{"value": "pretty good"}, nil)
	} else {
		acc.AddFields("state", map[string]interface{}{"value": "not great"}, nil)
	}

	return nil
}

var acc2 example.Accumulator

// func init() {
// 	inputs.Add("simple", func() telegraf.Input { return &Simple{} })
// }
func main() {
	fmt.Println("-------test fields-------")
	s := &Simple{
		Ok: true,
	}
	fmt.Println(s.Ok)
	s.Gather(acc2)
	fmt.Println(acc2)
	fmt.Println("add Accumulator successful")
}
