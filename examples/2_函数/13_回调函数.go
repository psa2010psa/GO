package main

import "fmt"

type FuncType func(int, int) int

//实现方法
func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

//回调函数：函数有一个参数是函数类型，这个函数就是回调函数
//多态，多种形态，调用同一个接口，不同的表现，可以实现不同表现
//先有想法，后面再实现功能
func Calla(a, b int, fTest FuncType) (result int) {
	fmt.Println("Calla")
	result = fTest(a, b) //这个函数一开始还没有实现，编译不会报错
	//result = add(a, b) //add()必须先定义后，才能调用
	return
}

func main() {
	fmt.Println("main")
	// a := Calla(2, 1, add)
	a := Calla(2, 1, minus)
	fmt.Println("a = ", a)

}
