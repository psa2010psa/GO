package main

import "fmt"

func main() {
	var a int
	fmt.Printf("请输入变量a: ")
	//阻塞等待用户的输入
	//fmt.Scanf("%d", &a) //别忘了 &地址
	fmt.Scan(&a)
	fmt.Println("a = ", a)
}
