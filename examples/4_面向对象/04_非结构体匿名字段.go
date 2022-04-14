package main

import "fmt"

type mystr string //自定义类型，给一个类型改名

type Person struct {
	name string //名字
	sex  byte   //性别
	age  int    //年龄
}

type Student struct {
	Person //结果体匿名字段
	int    //基础类型的匿名字段
	mystr
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 666, "bj"}
	fmt.Printf("s = %+v\n", s)
	//结果：s = {Person:{name:mike sex:109 age:18} int:666 mystr:bj}

	fmt.Println(s.name, s.sex, s.age, s.int, s.mystr)
	//结果：mike 109 18 666 bj

	fmt.Println(s.Person, s.int, s.mystr)
	//结果：{mike 109 18} 666 bj

	s.Person = Person{"go", 'm', 22}
	fmt.Println(s.Person, s.int, s.mystr)
	//结果：{go 109 22} 666 bj

}
