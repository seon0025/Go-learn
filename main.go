package main

import (
	"fmt"

	"github.com/seon0025/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "firstword"}
	definition, error := dictionary.Search("second")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(definition)
	}
}
