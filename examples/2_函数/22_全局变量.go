package main

import "fmt"

func test() {
	fmt.Println("test a = ", a)
}

//定义在函数外部的变量是全局变量
//全部变量在任何地方都能使用
var a int

func main() {
	a = 100
	fmt.Println("a = ", a)
	test()
}
