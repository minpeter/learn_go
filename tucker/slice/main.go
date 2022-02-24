package main

import "fmt"

func main() {
	a := make([]int, 0, 8) //초기길이 0, 최대길이 8

	fmt.Println("len(a) = ", len(a))
	fmt.Println("cap(a) = ", cap(a))

	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", append(a, 1))
	a = append(a, 1)

	fmt.Println("len(a) = ", len(a))
	fmt.Println("cap(a) = ", cap(a))
	fmt.Printf("%p\n", a)

	t := make([]int, 2, 4)
	t[0] = 1
	t[1] = 2

	y := make([]int, len(t)) //잘못 append하면 동일주소를 가르킬수 있기 때문에
	for i := 0; i < len(t); i++ {
		y[i] = t[i] //이런식으로 append 전에 완전 복사
		//하는게 안전하다 말이야
	}

	y = append(y, 3)

	y[0] = 100 //t에는 죽어도 영향을 끼치지 않아요 :)
	fmt.Printf("t의 주소: %p\n", t)
	fmt.Printf("y의 주소: %p\n", y)

}
