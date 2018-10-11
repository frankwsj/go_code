package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	get("http://www.baidu.com")
}

func get(my_url string) {
	i, err := http.Get(my_url)
	if err != nil {
		return
	}
	fmt.Println(i)
	defer i.Body.Close()
	body, err := ioutil.ReadAll(i.Body)
	/*if err != nil {
		return
	}*/
	fmt.Println(body)
	log.Println(string(body))

	//defer i.Body.close()
}

/*
func HttpGet() {
	my_url := "http://jd.com"
	response, err := http.Get(my_url)
	if err != nil { //如果访问不成功,url不存在则会进入改判断
		log.Println(err)
		return
	}
	defer response.Body.Close() //请求完了关闭回复主体
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(body)
	log.Println(string(body))
}
*/
