package main

import (
	"log"
	"sync"
)

type Singleton struct {
	a int
}

var singleton *Singleton

var once sync.Once

// GetInstance 用於獲取丹利模式對象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}

func main() {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		log.Println("failed")
	}
	ins2.a = 1
	log.Println(ins1.a) // expect output : 1
}
