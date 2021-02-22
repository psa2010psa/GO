package main

import "fmt"

/*
尽管go语言中没有封装、继承、多态这些概念，但同样通过别的方式实现这些特性：
封装：通过方法实现
继承：通过匿名字段实现
多态：通过接口实现
*/

type Person struct {
	name string //名字
	sex  byte   //性别
	age  int    //年龄
}

type Student struct {
	Person //只有类型，没有名字，这就是匿名字段，继承了Person成员
	id     int
	addr   string
}

func main() {
	//顺序初始化
	var s1 Student = Student{Person{"mike", 'm', 18}, 1, "bj"}
	fmt.Println("s1 = ", s1)
	//结果：s1 =  {{mike 109 18} 1 bj}

	//自动推导类型
	s2 := Student{Person{"mike", 'm', 18}, 1, "bj"}
	//%+v，显示更详细
	fmt.Printf("s2 = %+v\n", s2)
	//结果：s2 = {Person:{name:mike sex:109 age:18} id:1 addr:bj}

	//指定成员初始化，没有初始化的成员自动赋值为0
	s3 := Student{id: 1}
	fmt.Printf("s3 = %+v\n", s3)
	//结果：s3 = {Person:{name: sex:0 age:0} id:1 addr:}

	s4 := Student{Person: Person{name: "go"}, id: 2}
	fmt.Printf("s4 = %+v\n", s4)
	//结果：s4 = {Person:{name:go sex:0 age:0} id:2 addr:}
}
