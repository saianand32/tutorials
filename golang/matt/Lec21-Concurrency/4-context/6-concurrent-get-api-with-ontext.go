package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

// result struct holds the information for each request: the URL, an error (if any),
// and the latency of the request (time it took to get a response).
type result struct {
	url     string        // URL being requested
	err     error         // error encountered, if any
	latency time.Duration // time taken to get the response
}

// get function performs an HTTP GET request to a given URL and sends the result (including error and latency)
// to a channel. It also uses a context to allow cancellation or timeout of the request.
func get(ctx context.Context, url string, ch chan<- result) {
	// Record the start time for measuring latency
	start := time.Now()

	// Create an HTTP GET request with the given context (so it can be cancelled or timed out)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	// Perform the GET request
	if resp, err := http.DefaultClient.Do(req); err != nil {
		// If there's an error (e.g., timeout, unreachable server), send the error result to the channel
		ch <- result{url, err, 0}
	} else {
		// If the request is successful, calculate the time taken (latency)
		t := time.Since(start).Round(time.Millisecond) // Round to nearest millisecond for readability
		ch <- result{url, nil, t}                      // Send result with the URL, no error, and the latency
		// Close the response body after the request is complete (important for cleanup)
		resp.Body.Close()
	}
}

func main() {
	// Create a channel to communicate results from goroutines
	results := make(chan result)

	// List of URLs to fetch
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
	}

	// Create a context with a 1-second timeout, which will cancel all requests if they take longer than 1 second
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // Ensure cancel() is called to release resources when done

	// Start a goroutine for each URL in the list to fetch the pages concurrently
	for _, url := range list {
		go get(ctx, url, results) // Launch a goroutine for each request
	}

	// Collect the results from the channel (one for each URL)
	for range list {
		r := <-results // Wait for each result to be sent on the channel

		// Print the result: either the error (if any) or the latency
		if r.err != nil {
			// If there was an error (e.g., network issue, timeout), log it
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			// If the request was successful, log the URL and its latency
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}
}
