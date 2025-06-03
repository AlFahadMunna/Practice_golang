package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func fetchUrl(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", url, err)
		return
	}
	fmt.Printf("Fetched %s (%d bytes) in %v\n", url, len(body), time.Since(start))
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://google.com",
		"https://example.com",
	}

	var wg sync.WaitGroup
	start := time.Now()
	for _, url := range urls {
		wg.Add(1)             // Increment counter
		go fetchUrl(url, &wg) // Launch goroutine!
	}
	wg.Wait()
	fmt.Printf("Total time: %v\n", time.Since(start))
}
