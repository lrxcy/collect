package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%s lives in %s.\n", os.Getenv("USER"), os.Getenv("HOME"))

}
