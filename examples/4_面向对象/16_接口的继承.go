package main

import "fmt"

type Humaner interface {
	sayhi()
}

type Personer interface {
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
	var i Personer
	s := &Student{"mike", 666}
	i = s
	i.sayhi()
	i.sing("学生歌")
	/*结果
	Student[mike, 666] sayhi
	Student在唱着 学生歌
	*/
}
