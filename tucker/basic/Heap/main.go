package main

import (
	"fmt"

	"github.com/minpeter/learn_go/tucker/basic/dataStruct"
)

func main() {
	h := &dataStruct.Heap{}

	fmt.Printf("\n===Start Push===\n")

	h.Push(2)
	fmt.Print("Push:", 2, "  currt:")
	h.Print()
	h.Push(6)
	fmt.Print("Push:", 6, "  currt:")
	h.Print()
	h.Push(9)
	fmt.Print("Push:", 9, "  currt:")
	h.Print()
	h.Push(7)
	fmt.Print("Push:", 7, "  currt:")
	h.Print()
	h.Push(8)
	fmt.Print("Push:", 8, "  currt:")
	h.Print()

	fmt.Printf("\n===Start Pop===\n")

	fmt.Print("Pop:", h.Pop(), "  currt:")
	h.Print()
	fmt.Print("Pop:", h.Pop(), "  currt:")
	h.Print()
	fmt.Print("Pop:", h.Pop(), "  currt:")
	h.Print()
	fmt.Print("Pop:", h.Pop(), "  currt:")
	h.Print()
}
