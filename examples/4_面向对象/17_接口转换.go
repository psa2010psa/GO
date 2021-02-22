package main

import "fmt"

type Humaner interface { //子集
	sayhi()
}

type Personer interface { //超集
	Humaner //匿名字段，继承了sayhi()
	sing(lrc string)
}

type Student struct {
	name string
	id   int
}

//Student实现了sayhi()
func (tmp *Student) sayhi() {
	fmt.Printf("Student[%s, %d] sayhi\n", tmp.name, tmp.id)
}

func (tmp *Student) sing(lrc string) {
	fmt.Println("Student在唱着", lrc)
}

func main() {
	//超集可以转换为子集，反过来不可以
	var iPro Personer //超集
	var i Humaner     //子集

	//iPro = i //err Humaner does not implement Personer (missing sing method)

	iPro = &Student{"mike", 666}
	i = iPro //超级可以转换为子集
	i.sayhi()
	//结果：Student[mike, 666] sayhi
}
