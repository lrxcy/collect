package outputs

import "log"

type OutputPlugin interface {
	Write(string) error
}

type Creator func() OutputPlugin

var OutputPlugins = map[string]Creator{}

func AddOutputPlugin(name string, creator Creator) {
	OutputPlugins[name] = creator
}

type OutputsDelegate struct {
	OutputMaps map[string]Creator
}

func (o *OutputsDelegate) Execute(s string) {
	for i, creator := range o.OutputMaps {
		log.Printf("Going to write item %v\n", i)
		err := creator().Write(s)
		if err != nil {
			log.Println(err)
		}
	}
}

func InitOutputsPlugin() *OutputsDelegate {
	return &OutputsDelegate{
		OutputMaps: OutputPlugins,
	}
}
