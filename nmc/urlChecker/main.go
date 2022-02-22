package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url string
	status string
}

func main() {
	
	results := make(map[string]string)
	channel := make(chan requestResult)
	urls := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.youtube.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.instagram.com/",
		"https://www.linkedin.com/",
		"https://www.twitter.com/",
		"https://www.github.com/",
		"https://www.pinterest.com/",
		"https://www.quora.com/",
		"https://www.stackoverflow.com/",
		"https://www.tumblr.com/",
		"https://www.flickr.com/",
	}

	for _, url := range urls {
		go hitURL(url, channel)
	}

	for i:=0; i<len(urls); i++ {
		result := <-channel
		results[result.url] = result.status
	}
	
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, channel chan<- requestResult) { //send only
	fmt.Println("Checking URL:", url)
	status := "OK"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400{
		status = "FAILED"
	}
	channel <- requestResult{url: url, status: status}
}