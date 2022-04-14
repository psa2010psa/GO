package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Listen err = ", err)
		return
	}

	defer listener.Close()

	//阻塞等待用户的连接
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("Accept err = ", err1)
		return
	}

	defer conn.Close()

	//接收用户的数据
	buf := make([]byte, 1024*4)
	n, err2 := conn.Read(buf)
	if n == 0 {
		fmt.Println("Read err = ", err2)
		return
	}
	fmt.Printf("#%v#\n", string(buf[:n]))
	/*浏览器输入 http://127.0.0.1:8000
	#GET / HTTP/1.1  // 请求行
	Host: 127.0.0.1:8000
	Connection: keep-alive
	Cache-Control: max-age=0
	sec-ch-ua: "Chromium";v="88", "Google Chrome";v="88", ";Not A Brand";v="99"
	sec-ch-ua-mobile: ?0
	Upgrade-Insecure-Requests: 1
	User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36
	Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,
	Sec-Fetch-Site: none
	Sec-Fetch-Mode: navigate
	Sec-Fetch-User: ?1
	Sec-Fetch-Dest: document
	Accept-Encoding: gzip, deflate, br
	Accept-Language: zh-CN,zh;q=0.9,en;q=0.8  //上面都是请求头
											  //还有这个空行，其实下面还有一个包体
	#
	*/
}
