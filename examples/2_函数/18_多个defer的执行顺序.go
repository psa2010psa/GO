package main

import "fmt"

func test(x int) {
	result := 100 / x
	fmt.Println("result = ", result)
}

func main() {

	//如果一个函数中有多个defer语句，它们会以LIFO（先进后出）的顺序执行，
	//哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行
	defer fmt.Println("aaaaaa")

	defer fmt.Println("bbbbb")

	//test(0) //会崩  先输出 bbbb  aaaa   不会有cccc
	defer test(0) // 输出 cccc  bbbb aaa  然后输出报错信息

	defer fmt.Println("ccccc")

}
