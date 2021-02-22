package main

import "fmt"

func main() {
	var str1 string //声明变量
	str1 = "abc"    //双引号
	fmt.Println("str1 = ", str1)
	//自动推导
	str2 := "mike"
	fmt.Printf("str2 type is %T\n", str2)

	//
	fmt.Println("len(str2) = ", len(str2))

}
