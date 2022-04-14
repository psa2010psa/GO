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
	//顺序初始化，每个成员必须初始化  别忘了&
	var p1 *Student = &Student{1, "mike", 'm', 18, "bj"}
	fmt.Println("*p1 = ", *p1)
	//结果： *p1 =  {1 mike 109 18 bj}

	//指定成员初始化，没有初始化的成员，自动赋值为0
	p2 := &Student{name: "go", addr: "sh"}
	fmt.Printf("p2 type is %T\n", p2) //结果：p2 type is *main.Student
	fmt.Println("p2 = ", p2)          //结果：p2 =  &{0 go 0 0 sh}  注意&

}
