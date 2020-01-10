package main

import "log"

// AModuleAPI
type AModuleAPI interface {
	TestA() string
}

type AModule struct{}

func (AModule) TestA() string {
	return "TestA"
}

func NewAModule() AModuleAPI {
	return AModule{}
}

// BModuleAPI
type BModuleAPI interface {
	TestB() string
}

type BModule struct{}

func (BModule) TestB() string {
	return "TestB"
}

func NewBModule() BModuleAPI {
	return BModule{}
}

type ModulesAPI interface {
	Test() string
}

type Modules struct {
	AModule
	BModule
}

func (m Modules) Test() string {
	return m.AModule.TestA() + m.BModule.TestB()
}

func NewModuleInterface() ModulesAPI {
	return Modules{}
}

func main() {
	f := NewModuleInterface()
	log.Println(f.Test())
}
