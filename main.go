package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request failed")

func main() {
	urls := []string{
		"https://www.google.com/",
		"https://www.naver.com/",
		"https://www.daum.net/",
	}

	for _, url := range urls {
		hitURL(url)

	}
}

func hitURL(url string) error {
	fmt.Println("checking:", url)
	resp, err := http.Get(url)
	if err == nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
