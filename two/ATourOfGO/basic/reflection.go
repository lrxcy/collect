// refer to
// https://gist.github.com/drewolson/4771479
// https://studygolang.com/articles/2553

package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

func (f *Foo) reflect() {
	val := reflect.ValueOf(f).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
	}
}

func main() {
	f := &Foo{
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
	}

	f.reflect()
}
