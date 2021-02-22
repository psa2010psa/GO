package main

import "fmt"

//结构体是一种聚合的数据类型，它是由一系列具有相同类型或不同类型的数据构成的数据集合。每个数据称为结构体的成员

//定义一个结构体类型
type Student struct {
	id   int
	name string
	sex  byte //字节类型
	age  int
	addr string
}

func main() {
	//顺序初始化，每个成员必须初始化
	//var s1 Student = Student{1} //err: too few values in Student literal
	var s1 Student = Student{1, "mike", 'm', 18, "bj"}
	fmt.Println("s1 = ", s1)
	//结果： s1 =  {1 mike 109 18 bj}

	//指定成员初始化，没有初始化的成员，自动赋值为0
	s2 := Student{name: "go", addr: "sh"}
	fmt.Println("s2 = ", s2)
	//结果：s2 =  {0 go 0 0 sh}
}
