package main

import (
//	"fmt"
)

var m map[int]int = make(map[int]int)

var t int = 9

func test() []int {
	var t int = 9
	var m map[int]int = make(map[int]int)
	m[1] = 2
	m[2] = 7
	m[3] = 11
	m[4] = 15
	//i := m[1]
	//i, ok := m[1]

	for i, n := range m {
		_, prs := m[n]
		if prs {
			return []int(m[n], i)

		} else {
			m[t-n] = i
		}

	}
	return nil
	//fmt.Println("Key:", i, "Value:", v)

	//fmt.Println("Key:", k, "Value:", m[k])
}

func main() {
	test()
}
