package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaa")
}

func testb(x int) {
	//设置recover
	defer func() {
		//recover() //可以打印panic的错误信息
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}() //别忘了(),调用此匿名函数

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
	runtime error: index out of range
	cccccccc
	*/
}
