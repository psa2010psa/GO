package main

import "fmt"

func main() {
	s := "ab"
	//if和{就是条件，条件通常都是关系运算符
	if s == "abc" { //左括号{和if在同一行
		fmt.Println("满足")
	} else { //else后面没有条件
		fmt.Println("不满足")
	}
	//if支持1个初始化语句，初始化语句和判断条件用分号分隔
	if a := 10; a == 10 {
		fmt.Println("a等于10")
	} else {
		fmt.Println("a不等于10")
	}

	a := 1
	if a == 10 {
		fmt.Println("a等于10")
	} else if a > 10 {
		fmt.Println("a大于10")
	} else if a < 10 {
		fmt.Println("a小于10")
	} else {
		fmt.Println("不可能")
	}

}
