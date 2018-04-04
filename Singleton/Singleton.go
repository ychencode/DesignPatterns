package main

import (
	"sync"
	"fmt"
)

type Singleton struct {

}

var INSTANCE *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		INSTANCE = &Singleton{}
		fmt.Println("Singleton init.")
	})
	return INSTANCE
}

func (s *Singleton) Say() {
	fmt.Println("Hello from singleton")
}

func main() {
	single := GetInstance()
	single.Say()
	newSingle := GetInstance()
	newSingle.Say()
}
