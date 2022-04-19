package main

import "fmt"

func main() {
	//slice로 구현한 stack
	stack := []int{}
	for i := 1; i <= 5; i++ {
		stack = append(stack, i)
	}
	fmt.Println(stack)
	for len(stack) > 0 {
		var last int
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		fmt.Println(last)
	}
	//slice로 구현한 queue
	queue := []int{}
	for i := 1; i <= 5; i++ {
		queue = append(queue, i)
	}
	fmt.Println(queue)
	for len(queue) > 0 {
		var front int
		front, queue = queue[0], queue[1:]
		fmt.Println(front)
	}
}
