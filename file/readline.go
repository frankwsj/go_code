package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
)

func main() {
	Read()
	//Test()

}

//*按行读取文件
func Read() {
	//s := strings.NewReader("ABC\nDEF\r\nGHI\nJKL")
	s, _ := os.Open("test.txt")
	//s = "test.txt"
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords) //按词切分
	//bs.Split(bufio.ScanLines)    //按行切分
	//bs.Split(bufio.ScanBytes)

	var ss []string
	for bs.Scan() {
		//fmt.Printf("%s %v\n", bs.Bytes(), bs.Text())
		fmt.Println(bs.Text())
		lb := append(ss, bs.Text())
		//ll := append(l, bs.Text())
		fmt.Println(lb)
	}

	defer s.Close()
	return
}

/*
func Test() {
	b := Read()
	fmt.Println(b)

	//	for _, value := range Read("test.txt") {
	//		fmt.Println(value)
	//	}
}*/
