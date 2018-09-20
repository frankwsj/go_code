package main

import (
	"io"        // io.Copy获取http请求状态用;HttpPost
	"io/ioutil" //处理http获取的数据转换用
	"log"
	"net"
	"net/http" //http 请求用
	"net/url"  //HttpPost3中定义url.Values中使用到
	"os"       // os.Stdout 取io.Copy 返回数据中的status;HttpPost用
	"strings"  //HttpPost2中生成Post数据用
	"time"     //设置超时用
)

var my_url = "http://www.jd.com" //全局变量链接
func main() {
	HttpGet()
	//HttpPost()
	//HttpPost2()
	//HttpPost3()

}

//net/http包中提供了Get/Head/Post/PostForm函数来发出http/https请求
//通过get获取来http请求
func HttpGet() {
	response, err := http.Get(my_url)
	if err != nil { //如果访问不成功,url不存在则会进入改判断
		log.Println(err)
		return
	}
	defer response.Body.Close() //请求完了关闭回复主体
	body, err := ioutil.ReadAll(response.Body)
	log.Println(string(body))
}

//生成client 通过client的Do方法执行
func HttpPost() {
	/*     生成client,参数默认;
	*    这个结构体有四个属性
	*    Transport(RoundTrepper);指定执行的独立、单次http请求的机制
	*    CheckRedirect(func(req *Request, via []*Request)err):指定处理重定向的策略,如果不为nil,客户端在执行重定向之前调用本函数字段.如果CheckRedirect 返回一个错误，本类型中的get方法不会发送请求,如果CheckRedirect为nil,就会采用默认策略:连续请求10次后停止；
	＊    jar(CookieJar):jar指定cookie管理器,若果为nil请求中不会发送cookie,回复中的cookie会被忽略
	＊    TimeOut(time.Duration):指定本类型请求的时间限制，为0表示不设置超时
	*/
	//client := &http.Client{}    这里初始化了一个默认的client，并没有配置一些请求的设置

	//可以通过client中transport的Dail函数,在自定义Dail函数里面设置建立连接超时时长和发送接受数据超时
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*2) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * 2)) //设置发送接受数据超时
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * 2,
		},
	}
	reqest, err := http.NewRequest("POST", my_url, strings.NewReader("name=PostName")) //提交请求;用指定的方法，网址，可选的主体放回一个新的*Request
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		panic(err)
		return
	}
	response, err := client.Do(reqest) //前面预处理一些参数，状态，Do执行发送；处理返回结果;Do:发送请求,
	defer response.Body.Close()
	stdout := os.Stdout                     //将结果定位到标准输出，也可以直接打印出来，或定位到其他地方进行相应处理
	_, err = io.Copy(stdout, response.Body) //将第二个参数拷贝到第一个参数，直到第二参数到达EOF或发生错误，返回拷贝的字节喝遇到的第一个错误.
	status := response.StatusCode           //获取返回状态码，正常是200
	body, err := ioutil.ReadAll(response.Body)
	log.Println(string(body))
	log.Println(status)

}

//HTTP 包中POST 方法
func HttpPost2() {
	//HTTP POST请求
	req, err := http.Post(my_url, "application/x-www-form-urlencoded", strings.NewReader("name=myname")) //这里定义链接和post的数据
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(string(body))
}

//使用HTTP包中PostForm方法；
func HttpPost3() {
	req, err := http.PostForm(my_url, url.Values{"key": {"value"}, "id": {"123"}})
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(body))
}
