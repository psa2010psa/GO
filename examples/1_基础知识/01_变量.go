package main //必须有一个main包

import "fmt" //导入包含，必须要使用

func main() {
	/*变量  程序运行期间，可以改变的量
	1、声明格式 var 变量名 类型， 变量声明了，必须要使用
	2、只是声明没有初始化的变量，默认值为0
	3、同一个{}里，声明的变量名是唯一的
	4、可以同时声明多个变量 var b, c int
	5、自动推导类型，必须初始化，通过初始化的值确定类型 c := 30

	*/
	var a int

	fmt.Println("a = ", a)
	a = 10 //变量的赋值
	fmt.Println("a = ", a)

	var b int = 10 //初始化  声明变量时，同时赋值（一步到位）
	b = 20         //赋值    赋值前必须先声明，后赋值
	fmt.Println("b = ", b)

	//自动推导类型，先声明变量c,再给c赋值为30
	//自动推导，同一个变量名只能使用一次，用于初始化那次
	c := 30
	fmt.Printf("c type is %T\n", c)
	fmt.Println("c = ", c)

	//c := 40 //前面已经有变量c，不能再新建一个变量c
	c = 40 //可以继续赋值
	fmt.Println("c = ", c)

}
