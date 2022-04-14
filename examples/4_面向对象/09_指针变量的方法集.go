package main

import "fmt"

/*
类型的方法集是指可以被该类型的值调用的所有方法的集合

用实力value和pointer调用方法（含匿名字段）不受方法集约束，编译器总是查找全部方法，并自动转换receiver实参
*/
type Person struct {
	name string //名字
	sex  byte   //性别
	age  int    //年龄
}

func (p Person) SetInfoValue() {
	fmt.Println("SetInfoValue")
}

func (p *Person) SetInfoPointer() {
	fmt.Println("SetInfoPointer")
}

func main() {
	//结构体变量是一个指针变量，它能够调用哪些方法，这些方法就是一个集合，简称方法集
	p := &Person{"mike", 'm', 18}
	//p.SetInfoPointer() //func (p *Person) SetInfoPointer()
	//结果：SetInfoPointer
	(*p).SetInfoPointer() //结果：SetInfoPointer 把(*p)转换成p后再调用，相当于上面

	//内部做的转换，先把指针P,转成*p后再调用
	//相当于 (*p).SetInfoValue()
	p.SetInfoValue()
	//结果：SetInfoValue

}
