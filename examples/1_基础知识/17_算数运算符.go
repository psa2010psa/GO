package main

import "fmt"

//非0就是真，0就是假 (go语言的bool类型和int不兼容)
//判断一个数，在0~10的范围
func main() {
	var a int = 2
	var b int = 3
	var c int
	c = a + b
	fmt.Println("c = ", c)
	fmt.Println("4 > 3 的结果：", 4 > 3)
	fmt.Println("4 != 的结果：", 4 != 3)
	fmt.Println("!(4> 3 )结果：", !(4 > 3))
	//&& 与，并且，左边右边都是真，结果才为真，其他都是假
	fmt.Println("true && true 结果：", true && true)
	fmt.Println("true && false 结果：", true && false)

	// || ，或者，左边右边都是假，结果才为假，其他都是真
	fmt.Println("true || false 结果：", true || false)
	fmt.Println("false || false 结果：", false || false)

	d := 8
	//fmt.Println("0 <= d <= 10 结果为：", 0 <= d <= 10) 不支持这种写法
	fmt.Println("0 <= d <= 10 结果为：", 0 <= d && d <= 10)

}
