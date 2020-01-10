package main

// refer https://www.jianshu.com/p/fc4902159cf5

// type `Structure` func(`input parameters`) `return parameters`
// is like decorator method in python

import "fmt"

// Greeting function types
// 定義Greeting這個資料結構，key是函數(涵括傳入的參數)/value是該函數的返回值
type Greeting func(name string) string

func say(g Greeting, n string) {
	fmt.Println("say is processing ... ")
	fmt.Println(g(n))
}

func english(name string) string {
	fmt.Println("english is processing ... ")
	return "Hello, " + name
}

func main() {

	say(english, "World")

	// equivalent to below ...
	fmt.Printf("\n\nis equivalient to below ... \n\n")

	g := Greeting(english)
	g.say2("World")

	demoBuild(build, "concrete")

}

func (g Greeting) say2(n string) {
	fmt.Println(g(n))
}

func demoBuild(g Greeting, n string) {
	fmt.Println("is Building something")
	fmt.Println(g(n))
}

func build(material string) string {
	return fmt.Sprintf("The building is built with material : %v\n", material)
}
