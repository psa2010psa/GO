package main

import "fmt"

func main() {
	//定义一个数组，[10]int和[5]int是不同类型
	var a [10]int
	var b [5]int
	fmt.Printf("len(a) = %d, len(b) = %d\n", len(a), len(b))
	//结果：len(a) = 10, len(b) = 5

	//注意：定义数组，指定的数组元素个数必须是常量
	//n := 10
	//var c [n]int //err non-constant array bound n

	//操作数组元素，从0开始，到len()-1，不对称元素，这个数字叫下标
	//这个下标，可以是变量或常量
	a[0] = 1
	j := 1
	a[j] = 2 // a[1] = 2
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}
	//打印
	//第一个返回下标，第二个返回元素
	for j, data := range a {
		fmt.Printf("a[%d] = %d\n", j, data)
	}

}
