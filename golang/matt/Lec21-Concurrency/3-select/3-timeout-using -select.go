package main

import (
	"log"
	"net/http"
	"time"
)

// Define a struct to hold the result of each HTTP request
type result struct {
	url     string        // URL being requested
	err     error         // Error that occurred, if any
	latency time.Duration // Time taken for the request (if successful)
}

// get is the function that performs the HTTP GET request and sends the result to the channel
func get(url string, ch chan<- result) {
	// Record the start time to calculate latency
	start := time.Now()
	var nilerr error

	// Perform the HTTP GET request
	if resp, err := http.Get(url); err != nil {
		// If there's an error (e.g., network error, invalid URL), send the result with the error
		ch <- result{url, err, 0}
	} else {
		// If the request is successful, calculate the latency and send the result
		t := time.Since(start).Round(time.Millisecond) // Round to nearest millisecond
		ch <- result{url, nilerr, t}                   // Send the successful result
		resp.Body.Close()                              // Always close the response body to free up resources
	}
}

func main() {
	// Create a timeout channel using time.After. This function returns a channel
	// that will receive the current time after the specified duration (3 seconds in this case).
	// The program will be able to check if the timeout occurs before all requests finish.
	stopper := time.After(3 * time.Second)

	// Channel to collect the results of each HTTP request
	results := make(chan result)

	// List of URLs to be tested
	list := []string{
		"https://amazon.com", "https://google.com",
		"https://nytimes.com", "https://wsj.com",
	}

	// Launch a Goroutine for each URL to perform the HTTP GET requests concurrently
	for _, url := range list {
		go get(url, results) // start a CSP (Communicating Sequential Processes) process for each URL
	}

	// Wait for results or timeout using select
	for range list { // We are expecting results equal to the number of URLs
		select {
		// When a result is received, process it
		case r := <-results:
			if r.err != nil {
				// If there's an error (e.g., unable to fetch the URL), log it
				log.Printf("%-20s %s\n", r.url, r.err)
			} else {
				// If the request was successful, log the URL and the latency
				log.Printf("%-20s %s\n", r.url, r.latency)
			}
		// If the timeout is reached, log the timeout and stop the program
		case t := <-stopper:
			log.Fatal("TIMEOUT:", t) // Triggered if 3 seconds elapse without all results
		}
	}
}
