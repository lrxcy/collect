package main

import (
	"fmt"
	"log"
)

func main() {
	retry := 5
	for i := 0; i < retry; i++ {
		log.Println(i)

		err := returnError(i)
		if err == nil {
			retry = 0
		}
	}
}

// refer from :
// A simple retry machanism
// https://stackoverflow.com/questions/50676817/does-the-http-request-automatically-retry?rq=1

func returnError(i int) error {
	if i == 2 {
		return nil
	}
	return fmt.Errorf("This is error")

}
