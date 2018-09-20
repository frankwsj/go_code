package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	a := Person{}
	a.name = "wsj"
	a.age = 800
	fmt.Println(a)
	fmt.Println("hello, test enviroment~")
}
