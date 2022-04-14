package main

import "fmt"

/*
空接口(interface{})不包含任何的方法，正因为如此，所有的类型都实现了空接口，
因此空接口可以存储任意类型的数值。
*/

func main() {
	//空接口万能类型，保存任意类型的值
	var i interface{} = 1
	fmt.Println("i = ", i)

	i = "abc"
	fmt.Println("i = ", i)
	/*结果
	i=  1
	i =  abc
	*/
}
