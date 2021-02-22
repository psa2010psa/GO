package main

import "fmt"

func myfunc01(brr ...int) {
	for _, data := range brr {
		fmt.Println("tmp = ", data)
	}
}

func test01(arr ...int) {
	//全部元素传递给 myfunc01
	//myfunc01(arr...)

	//只想把后2个参数传递给另外一个函数使用
	//myfunc01(arr[2:]...) //从arr[2]（包括本身），把后面的所有元素传递过去

	//只想把前2个参数传递给另外一个函数使用
	myfunc01(arr[:2]...) //从arr[0] ~ arr[2] （不包括arr[2]） 所有元素传递过去
}

func main() {
	test01(1, 2, 3, 4)
}
