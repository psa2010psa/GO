package main

import "fmt"

func main() {
	var p1 *int
	p1 = new(int)              //p1为*int类型，指向匿名的int变量
	fmt.Println("*p1 = ", *p1) //*p1 = 0

	p2 := new(int) //p2为*int类型，指向匿名的int变量
	*p2 = 666
	fmt.Println("*p2 = ", *p2) //*p2 = 666

	/*
		我们只需使用new()函数，无需担心其内存的生命周期或怎么样将其删除，
		因为go语言的内存管理系统会帮我们搭理一切
	*/
}
