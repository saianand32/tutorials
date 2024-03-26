package main

import (
	"fmt"
	"time"
)

// can make a function a goroutine by adding the "go" clause infront of it
// below i had to put the time.Sleep method like a jugaad method as the main func was finishing off before the routine could complete the task
// this needs syncing for which we will use the sync package in the next tutoriial

func main() {
	go greeter(" Hello")
	greeter("World")
}

func greeter(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println(str)
	}
}
