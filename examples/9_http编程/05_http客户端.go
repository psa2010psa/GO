package main

import (
	"fmt"
	"net/http"
)

func main() {
	//resp, err := http.Get("http://www.baidu.com")
	resp, err := http.Get("http://127.0.0.1:8000/mike.html")
	if err != nil {
		fmt.Println("http.Get err = ", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Status = ", resp.Status)
	fmt.Println("StatusCode = ", resp.StatusCode)
	fmt.Println("Header = ", resp.Header)
	//fmt.Println("Body = ", resp.Body)

	buf := make([]byte, 1024*4)
	var tmp string

	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("read err = ", err)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println("tmp = ", tmp)

	/*开启04_http服务器.go
	结果
	localhost:9_http协议 pengchen$ go run 05_http客户端.go
	Status =  200 OK
	StatusCode =  200
	Header =  map[Content-Length:[8] Content-Type:[text/plain; charset=utf-8] Date:[Thu, 04 Mar 2021 07:11:32 GMT]]
	read err =  EOF
	tmp =  hello go
	*/
}
