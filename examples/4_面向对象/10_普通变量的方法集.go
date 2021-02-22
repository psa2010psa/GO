package main

import "fmt"

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
	//普通变量
	p := Person{"mike", 'm', 18}

	//内部，先把p转换为&p再调用  (&p).SetInfoPointer()
	p.SetInfoPointer() //结果：SetInfoPointer

	//都是普通变量 不用转
	p.SetInfoValue() //结果：SetInfoValue

}
