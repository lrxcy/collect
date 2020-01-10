package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	a := -1.1

	fmt.Println(int(a))

	prefix := "20191212"
	fmt.Println(len(prefix))

	t := time.Now()
	b := fmt.Sprintf("%v", t.Format("20060102150405"))
	c := fmt.Sprintf("%v", t.Unix())
	fmt.Println(len(b), len(c))
	fmt.Println(len("12312312301231231230"))
	s := fmt.Sprintf("%v%v", time.Now().Unix(), time.Now().Unix())
	fmt.Println(len(s))
	fmt.Println(reflect.TypeOf(s))

	printStrings("!", "@", "2", "1")

}

// 86,400

func printStrings(s ...string) {
	for i, j := range s {
		fmt.Println(i, j)
	}
}

func retriveString(s string) string {
	if len(s) < 8 {
		panic("Amount is less than .8 digits")
	}
	return s[:len(s)-8] + "." + s[len(s)-8:]
}
