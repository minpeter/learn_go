package main

import (
	"fmt"

	"github.com/minpeter/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word", "second": "Second word"}
	definition, err := dictionary.Search("fir")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}