package main

import (
	"fmt"
	"time"
)

func main() {
	// Create two channels of int type
	chans := []chan int{
		make(chan int), // Channel 1
		make(chan int), // Channel 2
	}

	// Launch two Goroutines, one for each channel
	// Each Goroutine sends values to its respective channel at different rates
	for i := range chans {
		go func(i int, ch chan<- int) {
			for {
				// Sleep for i seconds before sending a value to the channel
				// Goroutine 1 sends a value every 1 second,
				// Goroutine 2 sends a value every 2 seconds
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i // Send value i to the channel
			}
		}(i+1, chans[i]) // Pass i+1 to represent rate of sending (1s, 2s)
	}

	// Loop to receive messages from channels using select
	// The select statement allows us to wait on multiple channels and
	// receive values from whichever channel becomes ready first.
	for i := 0; i < 12; i++ {
		select {
		// If a message is available in channel 0, it will be received and printed
		case m0 := <-chans[0]:
			fmt.Println("received", m0) // This will be printed when a message is received from the first channel (every 1 second)

		// If a message is available in channel 1, it will be received and printed
		case m1 := <-chans[1]:
			fmt.Println("received", m1) // This will be printed when a message is received from the second channel (every 2 seconds)
		}
	}
}

// output:
// recieved 1
// recieved 2
// recieved 1
// recieved 1
// recieved 2
// recieved 1
// recieved 1
// recieved 2
// recieved 1
// recieved 1
// recieved 1
// recieved 2
