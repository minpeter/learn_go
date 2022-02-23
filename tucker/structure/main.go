package main

import "fmt"

type Student struct {
	name    string
	class   int
	sungjuk SungJuk
}

type SungJuk struct {
	name  string
	grade string
}

func (s Student) ViewSungJuk() {
	fmt.Println(s.sungjuk)
}

func main() {

	s := Student{name: "홍길동", class: 1}
	s.sungjuk.name = "수학"
	s.sungjuk.grade = "A+"
	s.ViewSungJuk()
}
