package main

import "fmt"

/*
接口定义
接口命名习惯以er结尾
接口只有方法声明，没有实现，没有数据字段
接口可以匿名嵌入其它接口，或嵌入到结构中

接口类型不能将其实例化
我们并不关心对象是什么类型，到底是不是鸭子，只关心行为

接口实现
接口是用来定义行为的类型。这些被定义的行为不由接口直接实现，而是由用户定义的类型实现，一个实现了这些方法的具体类型是这个接口类型的实例


*/

//定义接口类型
type Humaner interface {
	//方法，只有声明，没有实现，由别的类型（自定义类型）实现
	sayhi()
}

type Student struct {
	name string
	id   int
}

//Student实现了此方法
func (tmp *Student) sayhi() {
	fmt.Printf("Student[%s, %d] sayhi\n", tmp.name, tmp.id)
}

type Teacher struct {
	addr  string
	group string
}

//Teacher实现了此方法
func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher[%s, %s] sayhi\n", tmp.addr, tmp.group)
}

type Mystr string

//Mystr实现了此方法
func (tmp *Mystr) sayhi() {
	fmt.Printf("Mystr[%s] sayhi\n", *tmp)
}

//定义一个普通函数，函数的参数为接口类型
//只有一个函数，可以有不同表现，这就是多态
func WhoSayHi(i Humaner) {
	i.sayhi()
}

func main() {
	s := &Student{"mike", 666}
	t := &Teacher{"bj", "go"}
	var str Mystr = "hello mike"

	//调用同一个函数，不同表现，多态，多种形态
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)
	/*结果
	Student[mike, 666] sayhi
	Teacher[bj, go] sayhi
	Mystr[hello mike] sayhi
	*/

	//创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str

	//第一个返回下标，第二个返回下标所对应的值
	for _, i := range x {
		i.sayhi()
	}
	/*结果
	Student[mike, 666] sayhi
	Teacher[bj, go] sayhi
	Mystr[hello mike] sayhi
	*/

}

func main01() {
	//定义接口类型的变量
	var i Humaner

	//只要实现了此接口方法的类型，那么这个类型的变量(接收者)就可以给i赋值
	s := &Student{"mike", 666}
	i = s
	i.sayhi()

	t := &Teacher{"bj", "go"}
	i = t
	i.sayhi()

	var str Mystr = "hello mike"
	i = &str
	i.sayhi()
	/*结果
	Student[mike, 666] sayhi
	Teacher[bj, go] sayhi
	Mystr[hello mike] sayhi
	*/
}
