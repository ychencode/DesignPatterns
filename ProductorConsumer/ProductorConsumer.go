package main

import (
	"fmt"
	"time"
)

func productor(channel chan<- string) {
	for {
		channel <- fmt.Sprintf("current time is %v", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second * time.Duration(1))
	}
}

func consumer(channel <-chan string) {
	for {
		message := <- channel
		fmt.Println(message)
	}
}

func main() {
	channel := make(chan string, 5)
	go productor(channel)
	consumer(channel)
}
