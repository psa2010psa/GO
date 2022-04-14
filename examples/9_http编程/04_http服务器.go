package main

import (
	"fmt"
	"net/http"
)

//w 给客户端回复数据
//r 读取客户端发送的数据
func HandConn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.Method = ", r.Method)
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Header = ", r.Header)
	fmt.Println("r.Body = ", r.Body)

	w.Write([]byte("hello go")) //给客户端回复数据
}

func main() {
	//注册处理函数，用户连接，自动调用指定的处理函数
	http.HandleFunc("/mike.html", HandConn)

	//监听绑定
	http.ListenAndServe(":8000", nil)

	/*浏览器访问 http://127.0.0.1:8000/mike.html
	  浏览器显示结果 hello go
	*/

	/*服务器端打印结果
	r.Method =  GET
	r.URL =  /mike.html
	r.Header =  map[Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng]]
	r.Body =  {}
	*/
}
