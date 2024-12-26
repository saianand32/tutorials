package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

// result struct holds the information of each HTTP request's outcome.
// It contains the URL, any error encountered, and the latency (time it took for the request to complete).
type result struct {
	url     string        // The URL that was requested
	err     error         // Error encountered, if any (e.g., timeout, network failure)
	latency time.Duration // Time taken to get the response, measured in milliseconds
}

// get function sends an HTTP GET request to the provided URL using the given context.
// If the request is successful, it sends the result (including latency) to the channel.
// If there is an error, it sends the error to the channel instead.
func get(ctx context.Context, url string, ch chan<- result) {
	// Record the time before making the HTTP request to calculate latency
	start := time.Now()

	// Create an HTTP GET request with the provided context.
	// The context allows cancellation and timeout management.
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	// Execute the request using the default HTTP client
	if resp, err := http.DefaultClient.Do(req); err != nil {
		// If an error occurred (e.g., network failure, timeout), send the error result
		ch <- result{url, err, 0} // We don't have latency in case of an error
	} else {
		// If the request was successful, calculate the latency (round to the nearest millisecond)
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t} // Send the result containing URL and the latency
		// Close the response body to free resources after reading
		resp.Body.Close()
	}
}

func main() {
	// Create a channel to collect the results from the goroutines
	results := make(chan result)

	// List of URLs to send HTTP GET requests to
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
	}

	// Create a cancellable context. This context can be cancelled manually.
	// The context will be passed to each HTTP request, allowing us to cancel them at any point.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure that the context is cancelled when the program exits

	// Launch a goroutine for each URL in the list to perform the GET request concurrently
	for _, url := range list {
		go get(ctx, url, results) // start a goroutine for each request
	}

	// Wait for the first result to come in from the channel
	// This select block will either receive a result or handle context cancellation
	select {
	case r := <-results: // The first result comes in here
		// Log the result based on whether there was an error or a valid latency value
		if r.err != nil {
			// If there's an error, print the URL and the error message
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			// If the request was successful, print the URL and the latency
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
		// After receiving the first result, cancel the context to stop the remaining requests
		cancel()

	// This case will be triggered if the context is cancelled before receiving any result
	case <-ctx.Done():
		// This indicates that the context was cancelled, which might happen if a timeout occurs
		// or if you manually cancel it (like calling cancel())
		log.Printf("Parent context canceled..[if exists]")
	}

	// After the first result is processed, no more requests will be handled because the context has been cancelled
}
