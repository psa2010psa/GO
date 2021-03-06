package main

import "fmt"

//无参有返回值：只有一个返回值
//有返回值的函数需要通过rerun中断函数，通过return返回
func myfunc01() int {
	return 666
}

//给返回值起一个变量名，go推荐写法
func myfunc02() (result int) {
	return 777
}

//给返回值起一个变量名，go推荐写法
func myfunc03() (result int) {
	result = 888
	return
}

func main() {
	//无参数返回值函数调用
	var a int
	a = myfunc01()
	fmt.Println("a = ", a)

	b := myfunc02()
	fmt.Println("b = ", b)

	c := myfunc03()
	fmt.Println("c = ", c)
}
