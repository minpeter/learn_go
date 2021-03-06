/*
정수 배열과 정수 N이 주어지면,
N번째로 큰 배열 원소를 찾으시오

ex)
input: [-1, 3, -1, 5, 4], 2
output: 4

input: [2, 4, -2, -3, 8], 1
output: 8

input: [-5, -3, 1], 3
output: -5
*/

package main

import (
	"fmt"
)

func main() {
	var input1 = [...]int{-1, 3, -1, 5, 4}
	input2 := 2

	h := &Heap{}
	for i := 0; i < len(input1); i++ {
		h.Push(input1[i])
		if h.Count() > input2 {
			h.Pop()
		}
	}

	fmt.Println(h.Pop())

}

type Heap struct {
	list []int
}

func (h *Heap) Push(v int) {
	h.list = append(h.list, v)

	idx := len(h.list) - 1

	for idx >= 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 {
			break
		}
		if h.list[idx] < h.list[parentIdx] {
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *Heap) Pop() int {

	if len(h.list) == 0 {
		return 0
	}

	top := h.list[0]
	last := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	if len(h.list) == 0 {
		return top
	}
	h.list[0] = last
	idx := 0
	for idx < len(h.list) {
		swapIdx := -1
		leftIdx := idx*2 + 1
		if leftIdx >= len(h.list) {
			break
		}
		if h.list[leftIdx] < h.list[idx] {
			swapIdx = leftIdx
		}

		rightIdx := idx*2 + 2
		if rightIdx < len(h.list) {
			if h.list[rightIdx] < h.list[idx] {
				if swapIdx < 0 || h.list[swapIdx] > h.list[rightIdx] {
					swapIdx = rightIdx
				}
			}
		}
		if swapIdx < 0 {
			break
		}
		h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
		idx = swapIdx
	}
	return top
}

func (h *Heap) Count() int {
	return len(h.list)
}
