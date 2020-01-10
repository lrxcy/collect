// http://www.cnblogs.com/Goden/p/4601598.html
package main

import (
	"errors"
	"fmt"
)

// claim a Customerror struct with name1,name2 string
type Customerror struct {
	name1 string
	name2 string
	Err   error
}

func (cerr Customerror) Error() string {
	errorinfo := fmt.Sprintf("infoa : %s , infob : %s , original err info : %s ", cerr.name1, cerr.name2, cerr.Err.Error())
	return errorinfo
}

func main() {

	// method_1. use errors.New to create an
	var err error = errors.New("this is an new error")

	fmt.Println(err.Error())

	// mehtod_2. use fmt.Errorf
	var err2 error
	err2 = fmt.Errorf("%s", "the error test for fmt.Errorf")
	fmt.Println(err2.Error())

	// self-define error as duck type
	var err3 error
	err3 = &Customerror{
		name1: "error infor of name1",
		name2: "error infor of name2",
		Err:   errors.New("self-define error"),
	}

	fmt.Println(err3.Error())
}
