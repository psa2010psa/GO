package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaa")
}

func testb() {
	//fmt.Println("bbbbbbbb")
	//显式调用panic函数，导致程序中断
	panic("this is a panic test")

}

func testc() {
	fmt.Println("cccccccc")
}

func main() {
	//当panic异常发生时，程序会中断运行
	testa()
	testb()
	testc()

	/*结果
	aaaaaaaa
	panic: this is a panic test
	+错误信息
	*/
}
