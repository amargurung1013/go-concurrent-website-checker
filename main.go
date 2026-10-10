package main

import (
	"fmt"
	"net/http"
	"time"
)

type WebsiteResult struct {
	URL          string
	StatusCode   int
	Reachable    bool
	ResponseTime time.Duration
}

func checkWebsite(url string) WebsiteResult {
	start := time.Now()
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("request failed: ", err)
		return WebsiteResult{
			URL:        url,
			StatusCode: 0,
			Reachable:  false}
	}

	defer response.Body.Close()

	elapsed := time.Since(start)

	reachability := response.StatusCode >= 200 && response.StatusCode <= 299

	return WebsiteResult{
		URL:          url,
		StatusCode:   response.StatusCode,
		Reachable:    reachability,
		ResponseTime: elapsed}
}

func main() {

	websites := []string{
		"https://example.com",
		"https://github.com",
		"https://golang.org",
	}

	results := make(chan WebsiteResult)
	start := time.Now()

	for _, url := range websites {
		go func(url string) {
			result := checkWebsite(url)
			results <- result
		}(url)
	}

	for range websites {
		result := <-results
		fmt.Println(result.URL, result.StatusCode)
	}

	fmt.Println("total time:", time.Since(start))
}
