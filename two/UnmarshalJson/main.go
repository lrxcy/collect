package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/jimweng/UnmarshalJson/plugins"
	_ "github.com/jimweng/UnmarshalJson/plugins/all"
)

var ErrInvalidConfig = errors.New("invalid configuration")

func main() {

	config, err := LoadConfig("./agent.json")
	if err != nil {
		panic(err)
	}

	ag, err := NewAgent(config)
	if err != nil {
		panic(err)
	}

	_, err = ag.LoadPlugins()
	if err != nil {
		panic(err)
	}
	if err := ag.crank(); err != nil {
		panic(err)
	}

}

type MapStruct interface{}

type Config struct {
	plugins map[string]MapStruct
}

func (c *Config) Plugins() map[string]MapStruct {
	return c.plugins
}

func (c *Config) Apply(name string, v interface{}) error {
	if _, ok := c.plugins[name]; ok {
		// TODO: 加入json內的設定到每個plugin內
		return nil
	}
	return nil
}

func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var f interface{}
	jsonStruct := f
	if err := json.Unmarshal(data, &jsonStruct); err != nil {
		return nil, err
	}

	c := &Config{
		plugins: make(map[string]MapStruct),
	}

	// need to specify the exact type of jsonStruct
	for name, val := range jsonStruct.(map[string]interface{}) {
		c.plugins[name] = val
	}

	return c, nil
}

type Agent struct {
	Config  *Config
	plugins []plugins.Plugin
}

func NewAgent(config *Config) (*Agent, error) {
	agent := &Agent{Config: config}
	fmt.Printf("%v\n", config)

	return agent, nil
}

func (a *Agent) LoadPlugins() ([]string, error) {
	var names []string
	var err error

	for name, creator := range plugins.Plugins {
		plugin := creator()
		fmt.Println(name)

		err = a.Config.Apply(name, plugin)
		if err != nil {
			return nil, err
		}

		a.plugins = append(a.plugins, plugin)
		names = append(names, name)
	}
	return names, nil
}

func (a *Agent) crank() error {
	for _, plugin := range a.plugins {
		err := plugin.Gather()
		if err != nil {
			return err
		}
	}
	return nil
}
