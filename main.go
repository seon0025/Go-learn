package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [2]string{"nico", "seon"}
	for _, person := range people {
		go isGood(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

func isGood(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + " is good"
}
