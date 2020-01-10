package main

import (
	"flag"

	"github.com/jimweng/memoServer/model"
	"github.com/jimweng/memoServer/router"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	_, err := model.InitDB()
	if err != nil {
		panic(err)
	}
}

var fSample = flag.Bool("sample", false, "show debug mode help")

func main() {

	flag.Parse()
	if *fSample {
		model.InsertSampleData()
	}

	r := router.NewRouter()
	r.Run()
}
