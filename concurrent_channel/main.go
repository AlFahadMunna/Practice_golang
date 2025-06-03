package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL   string
	Bytes int
	Time  time.Duration
	Error error
}

func fetchURL(url string, ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- Result{URL: url, Error: err}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- Result{URL: url, Error: err}
		return
	}

	ch <- Result{
		URL:   url,
		Bytes: len(body),
		Time:  time.Since(start),
	}
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://google.com",
		"https://example.com",
	}
	var wg sync.WaitGroup
	ch := make(chan Result)
	start := time.Now()

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		if result.Error != nil {
			fmt.Printf("Error fetching %s: %v\n", result.URL, result.Error)
		} else {
			fmt.Printf("Fetched %s (%d bytes) in %v\n", result.URL, result.Bytes, result.Time)
		}
	}

	fmt.Printf("Total time: %v\n", time.Since(start))
}
