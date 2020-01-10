package main

import (
	"fmt"

	"github.com/naoina/toml"
	"github.com/naoina/toml/ast"
)

type Config struct {
	URL       string
	Username  string
	Password  string
	Database  string
	UserAgent string
	Tags      map[string]string

	plugins map[string]*ast.Table
}

func (c *Config) Plugins() map[string]*ast.Table {
	return c.plugins
}

func (c *Config) Apply(name string, v interface{}) error {
	if tbl, ok := c.plugins[name]; ok {
		return toml.UnmarshalTable(tbl, v)
	}

	return nil
}

func main() {
	var testVar Config
	fmt.Println(testVar)
	fmt.Println("test")
}
