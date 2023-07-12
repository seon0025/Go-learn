package main

import (
	"errors"
	"fmt"
	"time"
)

var errRequestFailed = errors.New("request failed")

func main() {
	go count("nico")
	count("seon")
}

func count(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is good", i)
		time.Sleep(time.Second)
	}
}
