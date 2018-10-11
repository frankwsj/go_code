//传递的参数
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	echo1()
	//echo2()
	echo3()
	//test2()

}

//时间
func echo1() {
	start := time.Now()
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println(s)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func echo2() {
	//var s, sep string
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func test2() {
	//var s, sep string
	//s, sep := "", ""
	for v, arg := range os.Args[0:] {
		fmt.Println(v, arg)
		//s += sep + arg
		//sep = " "
	}
	//fmt.Println(s)
}
