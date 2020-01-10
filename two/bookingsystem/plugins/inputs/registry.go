package inputs

import (
	"fmt"
)

type InputPlugin interface {
	Gather() error
}

type Creator func() InputPlugin

var InputPlugins = map[string]Creator{}

func AddInputPlugin(name string, creator Creator) {
	InputPlugins[name] = creator
}

type InputsDelegate struct {
	InputMaps map[string]Creator
}

func (i *InputsDelegate) Execute() {
	for i, creator := range i.InputMaps {
		// fmt.Printf("____Going to gather item %v\n", i)
		err := creator().Gather()

		if err != nil {
			// fmt.Println(err)
			ResultChannel <- fmt.Sprintf("error occure while process item %v with msg : %v", i, err)
		}

	}
}

func InitInputsPlugin() *InputsDelegate {
	return &InputsDelegate{
		InputMaps: InputPlugins,
	}
}

// TODO: any better way to distribute this channel to plugins ?
var ResultChannel = make(chan string)
