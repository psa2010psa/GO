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
	name   string //和Person同名了
}

func main() {
	var s Student
	//默认规则（就近原则），如果能在本作用域找到此成员，就操作此成员
	//                  如果没有找到，找到继承的字段
	s.name = "mike" //操作的Student的name，还是Person的name？ 结论为Student的name
	s.sex = 'm'
	s.age = 18
	s.addr = "bj"
	fmt.Printf("s = %+v\n", s)
	//结果：s = {Person:{name: sex:109 age:18} id:0 addr:bj name:mike}

	//显式调用
	s.Person.name = "yoyo"
	fmt.Printf("s = %+v\n", s)
	//结果：s = {Person:{name:yoyo sex:109 age:18} id:0 addr:bj name:mike}
}
