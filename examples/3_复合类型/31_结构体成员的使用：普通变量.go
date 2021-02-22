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
	//定义一个结构体普通变量
	var s Student

	//操作成员，需要使用点(.)运算符
	s.id = 1
	s.name = "mike"
	s.sex = 'm' //字符
	s.age = 18
	s.addr = "bj"

	fmt.Println("s = ", s) //结果：s =  {1 mike 109 18 bj}

}
