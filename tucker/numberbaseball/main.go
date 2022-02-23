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
	fmt.Print("%d회 만에 맞추셨습니다.\n", cnt)
}
