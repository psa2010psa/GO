package main

import "fmt"

func funcc(c int) {
	fmt.Println("c = ", c)
}

func funcb(b int) {
	funcc(b - 1)
	fmt.Println("b = ", b)
}

func funca(a int) {
	funcb(a - 1)
	fmt.Println("a = ", a)
}

func main() {
	//函数调用流程：先调用在返回， 先进后出
	//函数递归，函数调用自己本身，利用此特点
	funca(3)
	fmt.Println("main")
}
