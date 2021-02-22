package main

import (
	"fmt"
	"os"
)

func main() {
	//接收用户传递的参数，都是以字符串方式传递
	list := os.Args
	n := len(list)
	fmt.Println("n = ", n) //go build 20_xxx.go 然后 20_xxx a b  打印结果 n=3

	for i, data := range list {
		fmt.Printf("list[%d] = %s\n", i, data)
	}
}
