package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	write("test.txt")
	append("test.txt")
	read()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//读取文件
func read() {
	fmt.Println("测试读\r\n取文件")
	b, err := ioutil.ReadFile("test.txt")
	check(err)
	fmt.Println(b)
	str := string(b)
	//str = typeof(str)
	//fmt.Println(typeof(str))
	fmt.Println(len(str))

	fmt.Println(str)
}

func write(fileName string) {
	fmt.Println("测试写入文件")
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile(fileName, d1, 0644)
	check(err)
}

func append(fileName string) {

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)

	data := []byte("再追加一行测试,dallsldl")
	if err != nil {
		return
	}
	defer f.Close()
	n, err := f.Write(data)
	if err == nil && n < len(data) {

	}

}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
