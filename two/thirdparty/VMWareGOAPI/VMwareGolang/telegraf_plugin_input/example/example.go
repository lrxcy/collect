// // func body() string { return 123 }

// type body interface {
// 	test1() string
// }

// type Test func() body

package main

import (
	"fmt"
)

type Greeting func(name string) string

// func say(g Greeting, n string) {
// 	fmt.Println(g(n))
// }

func (g Greeting) say(n string) {
	fmt.Println(g(n))
}

func english(name string) string {
	return "hello, " + name
}

func main() {
	g := Greeting(english)
	g.say("world")
	// say(english, "world")

}

// type niceToMeetYou func() test

// type test interface {
// 	str() string
// }

// func givenFunc(niceToMeetYou) {

// }
