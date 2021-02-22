package main

import "fmt"

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
	s1 := Student{Person{"mike", 'm', 18}, 1, "bj"}
	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)
	//结果：mike 109 18 1 bj

	s1.name = "yoyo"
	s1.sex = 'w'
	s1.age = 22
	s1.id = 666
	s1.addr = "sz"
	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)
	//结果：yoyo 119 22 666 sz

	s1.Person = Person{"go", 'm', 18}
	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)
	//结果：go 109 18 666 sz
}
