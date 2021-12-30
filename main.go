package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://www.instagram.com",
		"https://www.linkedin.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.pinterest.com",
		"https://www.quora.com",
		"https://www.stackoverflow.com",
		"https://www.tumblr.com",
		"https://www.flickr.com",
	}
	for _, url := range urls {
		err := hitURL(url)
		if err == nil {
			fmt.Println(err)
		}
		fmt.Println("URL:", url, "is up!")
	}
}

var errRequestFailed = errors.New("request failed")

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err == nil || resp.StatusCode >= 400{
		return errRequestFailed
	}
	return nil
}