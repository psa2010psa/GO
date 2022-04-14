package main

import "fmt"

//定义一个结构体类型
type Student struct {
	id   int
	name string
	sex  byte //字节类型
	age  int
	addr string
}

func main() {
	//1、指针有合法指向后，才操作成员
	//先定义一个结构体普通变量
	var s Student
	//再定义一个指针变量，保存s的地址
	var p1 *Student
	p1 = &s

	//通过指针操作成员，p1.id和(*p1).id完全等价，只能使用.运算符
	p1.id = 1
	(*p1).name = "mike"
	p1.sex = 'm' //字符
	p1.age = 18
	p1.addr = "bj"

	fmt.Println("p1 = ", p1) //结果：p1 =  &{1 mike 109 18 bj}

	//2、通过new申请一个结构体
	p2 := new(Student)

	p2.id = 1
	(*p2).name = "go"
	p2.sex = 'm' //字符
	p2.age = 20
	p2.addr = "sh"

	fmt.Println("p2 = ", p2) //结果：p2 =  &{1 go 109 20 sh}

}
