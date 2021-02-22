package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func minus(a, b int) int {
	return a - b
}

//函数也是一种数据类型,通过type给一个函数类型起名
//FuncType 它是一个函数类型
type FuncType func(int, int) int //没有函数名字，没有{}  好处：多态思想

func main() {
	var result int
	result = add(1, 2) //传统调用方式
	fmt.Println("result = ", result)

	//声明一个函数类型的变量，变量名叫fTest
	var fTest FuncType
	fTest = add            //是变量就可以负值
	result = fTest(10, 20) //等价于add(10,20)
	fmt.Println("result2 = ", result)

	fTest = minus
	result = fTest(10, 5) //等价于minus(10,5)
	fmt.Println("result3 = ", result)

}
