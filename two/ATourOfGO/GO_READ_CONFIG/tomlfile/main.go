package main

// refer : https://reurl.cc/Gxlmx

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Title   string
	Owner   ownerInfo
	DB      database `toml:"database"`
	Servers map[string]server
	Clients clients
}

type ownerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

type database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}

func main() {
	var config tomlConfig
	filePath := "/Users/jimweng/go/src/github.com/jimweng/ATourOfGO/GO_READ_CONFIG/tomlfile/test.conf"
	if aa, err := toml.DecodeFile(filePath, &config); err != nil {
		fmt.Printf("%v\n", aa)
		panic(err)
	}
	fmt.Printf("___%v\n", config)
}
