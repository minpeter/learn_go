package main

import (
	"fmt"

	"github.com/minpeter/learn_go/nmc/mydict/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	dictionary.Add("first", "First word")
	dictionary.Add("second", "Second word")

	dictionary.Update("first", "First word Update")
	dictionary.Delete("second")
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
