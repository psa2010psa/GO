package main

import "fmt"

func main() {
	a := 10
	b := "abc"
	c := 'a'
	d := 3.14

	fmt.Printf("%T, %T, %T, %T\n", a, b, c, d)

	//%d 整型格式
	//%s 字符串格式
	//%c 字符格式
	//%f 浮点型格式
	fmt.Printf("%d, %s, %c, %f\n", a, b, c, d)
	//%v自动匹配格式输出, 字符类型不行
	fmt.Printf("%v, %v, %v, %v\n", a, b, c, d)
}
