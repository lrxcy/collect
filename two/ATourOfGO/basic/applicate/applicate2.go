package main

import (
	"fmt"

	"github.com/jimweng/GoStructPractice/myfirstgo/plugins"
	"github.com/naoina/toml"
	"github.com/naoina/toml/ast"
	"github.com/vektra/cypress"
)

type Metrics interface {
	Receive(*cypress.Message) error
}

type Config struct {
	URL       string
	Username  string
	Password  string
	Database  string
	UserAgent string
	Tags      map[string]string

	plugins map[string]*ast.Table
}

type Agent struct {
	Debug bool
	HTTP  string

	Config *Config

	plugins []plugins.Plugin
	metrics Metrics

	eachInternal []func()
}

func (c *Config) Apply(name string, v interface{}) error {
	if tbl, ok := c.plugins[name]; ok {
		return toml.UnmarshalTable(tbl, v)
	}

	return nil
}

func main() {

	fmt.Println("hi")
}
