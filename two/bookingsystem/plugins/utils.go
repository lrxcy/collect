package plugins

import (
	"github.com/jimweng/bookingsystem/plugins/inputs"
	"github.com/jimweng/bookingsystem/plugins/outputs"
)

type InputImp interface {
	Execute()
}

type OutputImp interface {
	Execute(string)
}

var (
	Inputs  InputImp
	Outputs OutputImp

	GatherChannel chan string
)

func init() {
	Inputs = inputs.InitInputsPlugin()
	Outputs = outputs.InitOutputsPlugin()

	GatherChannel = inputs.ResultChannel
}
