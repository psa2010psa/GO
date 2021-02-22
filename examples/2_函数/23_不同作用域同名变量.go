package main

import "fmt"

var a byte //全局变量

func main() {
	var a int //局部变量
	//1、不同作用域，允许定义同名变量
	//2、使用变量的原则：就近原则
	fmt.Printf("%T\n", a) //结果：int

	{
		var a float32
		fmt.Printf("%T\n", a) //结果：float32
	}

	test()
}

func test() {
	fmt.Printf("%T\n", a) //结果：uint8 就是byte
}
