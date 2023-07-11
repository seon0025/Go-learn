package main

import "fmt"

func main() {
	urls := []string{
		"https://www.google.com/",
	}

	for url := range urls {
		fmt.Println(url)
	}
}
