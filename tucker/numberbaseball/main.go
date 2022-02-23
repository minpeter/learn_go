package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	strikes int
	balls   int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// 무작위 숫자 3개 생성
	numbers := MakeNumbers()

	cnt := 0
	for {
		cnt++
		// 사용자 입력
		inputNumbers := InputNumbers()

		// 사용자 입력과 무작위 숫자를 비교
		result := CompareNumbers(numbers, inputNumbers)

		// 결과 출력
		PrintResult(result)

		// 3s 인가?
		if IsGameEnd(result) {
			// 반복문 종료
			break
		}
	}
	//걸린 회수 출력
	fmt.Printf("%d 회 만에 맞추셨습니다.\n", cnt)
}

func MakeNumbers() [3]int {
	// 0~9 사이에 겹치지 않는 무작위 숫자 3개 반환
	var rst [3]int

	for i := 0; i < 3; i++ {
		for {
			n := rand.Intn(10)
			duplicated := false
			for j := 0; j < i; j++ {
				if rst[j] == n {
					duplicated = true
					break
				}
			}
			if !duplicated {
				rst[i] = n
				break
			}
		}
	}
	return rst
}

func InputNumbers() [3]int {
	// 키보드로 숫자 3개 입력 받아서 배열 반환
	var rst [3]int
	for {
		fmt.Print("숫자 3개를 입력하세요: ")
		var no int
		_, err := fmt.Scanf("%d\n", &no)
		if err != nil {
			fmt.Println("잘못된 입력입니다.")
			continue
		}

		success := true
		idx := 0
		for no > 0 {
			n := no % 10
			no = no / 10

			duplicated := false
			for i := 0; i < idx; i++ {
				if rst[i] == n {
					duplicated = true
					break
				}
			}
			if duplicated {
				fmt.Println("숫자가 중복되었습니다.")
				success = false
				break
			}

			if idx >= 3 {
				fmt.Println("숫자 3개를 입력하셔야 합니다.")
				success = false
				break
			}

			rst[idx] = n
			idx++
		}

		if success && idx < 3 {
			fmt.Println("숫자 3개를 입력하셔야 합니다.")
			success = false
		}

		if !success {
			continue
		}
		break
	}
	rst[0], rst[2] = rst[2], rst[0]
	return rst
}

func CompareNumbers(numbers, inputNumbers [3]int) Result {
	// 두개의 숫자 3개를 비교해서 결과를 반환
	strike := 0
	balls := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if numbers[i] == inputNumbers[j] {
				if i == j {
					strike++
				} else {
					balls++
				}
			}
		}
	}

	return Result{strikes: strike, balls: balls}
}

func PrintResult(result Result) {
	// 결과 출력
	fmt.Println("스트라이크:", result.strikes, "볼:", result.balls)
}

func IsGameEnd(result Result) bool {
	// 게임 종료 여부 반환
	return result.strikes == 3
}
