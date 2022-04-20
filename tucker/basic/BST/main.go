package main

import (
	"fmt"

	"github.com/minpeter/learn_go/tucker/basic/dataStruct"
)

func main() {
	tree := dataStruct.NewBinaryTree(5)
	tree.Root.AddNode(3)
	tree.Root.AddNode(2)
	tree.Root.AddNode(4)
	tree.Root.AddNode(8)
	tree.Root.AddNode(7)
	tree.Root.AddNode(6)
	tree.Root.AddNode(10)
	tree.Root.AddNode(9)

	tree.Print()

	searchV := 11
	if found, cnt := tree.Search(searchV); found {
		fmt.Println("found:", searchV, "cnt:", cnt)
	} else {
		fmt.Println("not found cnt: ", cnt)
	}
}
