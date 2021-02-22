package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别
	age  int    //年龄
}

func (p Person) SetInfoValue() {
	fmt.Printf("SetInfoValue: %p, %v\n", &p, p)
}

func (p *Person) SetInfoPointer() {
	fmt.Printf("SetInfoPointer: %p, %v\n", p, p)
}

func main() {
	p := Person{"mike", 'm', 18}
	fmt.Printf("main: %p, %v\n", &p, p)

	//方法值  f:= p.SetInfoPointer //隐藏了接收者
	//方法表达式
	f := (*Person).SetInfoPointer
	f(&p) //显式把接收者传递过去 ====>  相当于 p.SetInfoPointer
	//结果：SetInfoPointer: 0xc000096000, &{mike 109 18}

	f2 := (Person).SetInfoValue
	f2(p) //显式把接收者传递过去 ====>  相当于 p.SetInfoValue
	//结果：SetInfoValue: 0xc0000a2000, {mike 109 18}
}
