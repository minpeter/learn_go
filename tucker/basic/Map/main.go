package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["AAA"] = "010-1111-1111"
	var ok bool
	_, ok = m["AAA"]
	fmt.Print(ok, " ")
	_, ok = m["BBB"]
	fmt.Println(ok)
	delete(m, "AAA")
	_, ok = m["AAA"]
	fmt.Print(ok, " ")
	_, ok = m["BBB"]
	fmt.Println(ok)

	m["CCC"] = "010-1111-1111"
	m["BBB"] = "010-2222-2222"
	m["AAA"] = "010-3333-3333"

	for k, v := range m {
		fmt.Println(k, v)
	}
}
