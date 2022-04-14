package main

import "fmt"

//多个返回值
func myfunc01() (int, int, int) {
	return 1, 2, 3
}

//go语言官方推荐写法
func myfunc02() (a int, b int, c int) {
	a, b, c = 111, 222, 333
	return
}
func myfunc03() (a, b, c int) {
	a, b, c = 111, 222, 333
	return
}
func main() {
	//函数调用
	a, b, c := myfunc01()
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	d, e, f := myfunc03()
	fmt.Printf("d = %d, e = %d, f = %d\n", d, e, f)
}
