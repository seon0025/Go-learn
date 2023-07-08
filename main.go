package main

import (
	"fmt"

	"github.com/seon0025/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "greeting"
	error := dictionary.Add(word, definition)
	if error != nil {
		fmt.Println(error)
	}
	hello, error := dictionary.Search(word)
	fmt.Println(hello)
}
