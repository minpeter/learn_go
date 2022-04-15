package main

import "fmt"

type Student struct {
	name string
	age  int
}

func (t *Student) SetName(name string) {
	t.name = name
}

func main() {
	a := &Student{"Tom", 18} //여기서 a는 인스턴트라고 함

	a.SetName("Jerry") // a가 주체가 되는 행동

	fmt.Println(a)
}
