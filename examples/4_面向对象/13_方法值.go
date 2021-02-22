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

	p.SetInfoPointer() //传统调用方式
	/*结果 地址没变
	main: 0xc000096000, {mike 109 18}
	SetInfoPointer: 0xc000096000, &{mike 109 18}
	*/

	//保存函数入口地址
	pFunc := p.SetInfoPointer //这个就是方法值，调用函数时，无需再传递接收者，隐藏了接收者
	pFunc()                   //等价于 p.SetInfoPointer()
	//结果：SetInfoPointer: 0xc0000a0000, &{mike 109 18}

	vFunc := p.SetInfoValue
	vFunc() //相当于 p.SetInfoValue()  地址变了，只是值拷贝
	//结果：SetInfoValue: 0xc0000ac020, {mike 109 18}
}
