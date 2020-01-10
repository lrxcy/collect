package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name   string `user name` // 引號裡面的就是tag
	Passwd string `user password`
}

func main() {
	user := &User{"chronos", "pass"}
	s := reflect.TypeOf(user).Elem() //通過反射獲取type定義
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag) //將tag輸出出來
	}
}
