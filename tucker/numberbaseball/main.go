package main

import (
	"fmt"
)

func main() {
	// 무작위 숫자 3개 생성
	numbers := MakeNumbers()

	cnt := 0
	for {
		// 사용자 입력
		inputNumbers := InputNumbers()

		// 사용자 입력과 무작위 숫자를 비교
		result := CompareNumbers(numbers, inputNumbers)

		// 결과 출력
		PrintResult(result)

		// 3s 인가?
		if IsGameEnd(result) {
			fmt.Println("3s")
			break
		}
	}
	//걸린 회수 출력
	fmt.Printf("%d 회 만에 맞추셨습니다.\n", cnt)
}

func MakeNumbers() [3]int {
	// 0~9 사이에 겹치지 않는 무작위 숫자 3개 반환
	var rst [3]int
	return rst
}

func InputNumbers() [3]int {
	// 키보드로 숫자 3개 입력 받아서 배열 반환
	var rst [3]int
	return rst
}

func CompareNumbers(numbers, inputNumbers [3]int) bool {
	// 두개의 숫자 3개를 비교해서 결과를 반환
	return true
}

func PrintResult(result bool) {
	// 결과 출력
	fmt.Println(result)
}

func IsGameEnd(result bool) bool {
	// 게임 종료 여부 반환
	return true
}
