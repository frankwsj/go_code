package main

import "fmt"

func main() {
	for1()
	for2()
	for3()
}

func for1() {
	x := 0
	var sum int
	for x = 1; x < 101; x++ {
		sum = sum + x
		fmt.Println(sum)
	}
}

func for2() {
	var a, b int
	a = 0
	b = 100
	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}
}

func for3() {
	numbers := [6]int{1, 2, 3, 4}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}
