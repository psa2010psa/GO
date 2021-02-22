package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaa")
}

func testb(x int) {
	var a [10]int
	a[x] = 111 //当x为20的时候，导致数组越界，产生一个panic，导致程序崩溃

}

func testc() {
	fmt.Println("cccccccc")
}

func main() {
	testa()
	testb(20)
	testc()

	/*结果
	aaaaaaaa
	panic: this is a panic test
	+错误信息
	*/
}
