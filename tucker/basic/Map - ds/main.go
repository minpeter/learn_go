package main

import (
	"fmt"

	"github.com/minpeter/learn_go/tucker/basic/dataStruct"
)

func main() {
	fmt.Println("abcde = ", dataStruct.Hash("abcde"))
	fmt.Println("abcde = ", dataStruct.Hash("abcde")) //여러번 호출해도 같은 값
	fmt.Println("abcdf = ", dataStruct.Hash("abcdf")) //입력값이 다른 경우는 다른 값이 나옴
	fmt.Println("tbcde = ", dataStruct.Hash("tbcde"))
	fmt.Println("qwerasdfzxcvttyuighjkvbmn = ", dataStruct.Hash("qwerasdfzxcvttyuighjkvbmn")) //입력값의 길이가 길어도 같은 길이의 출력

	m := dataStruct.CreatMap()
	m.Add("AAA", "010-1111-1111")
	m.Add("BBB", "010-2222-2222")
	m.Add("VUDHFEJKDJ", "010-3333-3333")
	m.Add("CCC", "010-4444-4444")

	fmt.Println("AAA = ", m.Get("AAA"))
	fmt.Println("BBB = ", m.Get("BBB"))
	fmt.Println("VUDHFEJKDJ = ", m.Get("VUDHFEJKDJ"))
	fmt.Println("CCC = ", m.Get("CCC"))
	fmt.Println("DDD (없음) = ", m.Get("DDD"))

}
