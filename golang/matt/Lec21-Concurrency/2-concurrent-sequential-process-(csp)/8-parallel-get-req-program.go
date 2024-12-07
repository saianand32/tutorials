package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()
	var nilerr error
	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0} // error response
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nilerr, t} // normal response
		resp.Body.Close()
	}
}

func main() {
	results := make(chan result) // channel for results
	list := []string{"https://amazon.com", "https://google.com",
		"https://nytimes.com", "https://wsj.com",
	}
	for _, url := range list {
		go get(url, results) // start a CSP process
	}
	for range list { // read from the channel
		r := <-results
		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}

}
