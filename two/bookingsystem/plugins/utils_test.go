package plugins

import (
	_ "github.com/jimweng/bookingsystem/plugins/inputs/all"
	_ "github.com/jimweng/bookingsystem/plugins/outputs/all"

	"testing"
)

func TestSomething(t *testing.T) {
	// execute inputs to collector metrics values from plugins
	Inputs.Execute()

	/*
		execute outputs to collector metrics values from plugins
		outputs would not return anything

		it might need another channel to collector errors
	*/
	Outputs.Execute("") // unexpected end of JSON input

}
