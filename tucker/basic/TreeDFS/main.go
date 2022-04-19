package main

import (
	"fmt"

	"github.com/minpeter/learn_go/tucker/basic/dataStruct"
)

func main() {
	tree := dataStruct.Tree{}
	val := 1
	tree.AddNode(val)
	val++
	for i := 0; i < 3; i++ {
		tree.Root.AddNode(val)
		val++
	}

	for i := 0; i < len(tree.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Childs[i].AddNode(val)
			val++
		}
	}
	fmt.Println("재귀 호출을 이용한 DFS")
	tree.DFS1()
	fmt.Println()
	fmt.Println("스택을 이용한 DFS")
	tree.DFS2()
}
