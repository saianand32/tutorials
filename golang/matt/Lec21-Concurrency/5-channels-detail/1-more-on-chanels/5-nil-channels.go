package main

import "fmt"

func main() {
	var ch chan int // Nil channel (uninitialized)

	// Attempting to send to a nil channel will block forever
	go func() {
		ch <- 1 // This will block forever because the channel is nil
	}()

	// Attempting to receive from a nil channel will block forever
	go func() {
		val := <-ch // This will block forever as well
		fmt.Println(val)
	}()

	// Sleep to give goroutines time to block (this is just to show the effect)
	// In real-world code, you should use synchronization tools (e.g., WaitGroup)
	select {}
}
