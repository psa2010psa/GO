package main

import (
	"net/http"
)

//w 给客户端回复数据
//r 读取客户端发送的数据
func HandConn(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello go")) //给客户端回复数据
}

func main() {
	//注册处理函数，用户连接，自动调用指定的处理函数
	http.HandleFunc("/mike.html", HandConn)

	//监听绑定
	http.ListenAndServe(":8000", nil)

	/*浏览器访问 http://127.0.0.1:8000/mike.html
	结果 hello go
	*/
}
