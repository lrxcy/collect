package main

import "fmt"

type API interface {
	Say(string) string
}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

type hiAPI struct{}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("hi %s\n", name)
}

type helloAPI struct{}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("hello %s\n", name)
}

func main() {
	hiJim := NewAPI(1)
	fmt.Println(hiJim.Say("jim"))

	helloJim := NewAPI(2)
	fmt.Println(helloJim.Say("jim"))
}
