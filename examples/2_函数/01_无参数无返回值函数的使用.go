package main //必须

import "fmt"

//无参数，无返回值函数的定义
func Myfunc() {
	a := 666
	fmt.Println("a = ", a)
}
func main() {
	//无参数无返回值函数的调用：函数名()
	Myfunc()
}
