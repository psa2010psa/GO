package main

import "fmt"

func main() {
	// Myfunc01(666) //not enough arguments in call to Myfunc01
	Myfunc02()
	Myfunc02(1)
	Myfunc02(1, 2, 3)

	//Myfunc03() //not enough arguments in call to Myfunc03
	Myfunc03(1, 2, 3, 4)
}

func Myfunc01(a int, b int) { //固定参数

}

//...int类型 这样的类型， ...type不定参数类型
//注意：不定参数，一定（只能）放在形参中的最后一个参数
func Myfunc02(args ...int) { //传递的实参可以是0或多个
	fmt.Println("len(args) = ", len(args)) //获取用户传递参数的个数
	for i, data := range args {
		fmt.Printf("args[%d] = %d\n", i, data)
	}
}

//固定参数一定要传参，不定参数根据需求传参
func Myfunc03(a int, args ...int) {

}
