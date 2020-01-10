// GO document
// https://blog.golang.org/defer-panic-and-recover
// GO example
// http://blog.csdn.net/chenbaoke/article/details/41966827
// https://gobyexample.com/panic
// GO defer
// https://xiaozhou.net/something-about-defer-2014-05-25.html

// golang中没有try... catch...，所以當程式中遇到panic，如果不進行recover便會導致程式停止
// 解決辦法 ： 使用defer將延遲處理的recover進行恢復
package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
