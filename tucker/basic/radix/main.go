package main

import "fmt"

func main() {
	arr := [15]int{0, 5, 4, 9, 3, 8, 2, 7, 1, 6, 9, 2, 1, 3, 2}
	temp := [10]int{} //특정 값이 몇개 있는지 세는 배열  arr의 배열원소의 범위 만큼 생성

	for i := 0; i < len(arr); i++ {
		idx := arr[i]
		temp[idx]++
	}

	idx := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr[idx] = i
			idx++
		}
	}

	fmt.Println(arr)
}
