package main

import "fmt"

//定义一个结构体类型
type Student struct {
	id   int
	name string
	sex  byte //字节类型
	age  int
	addr string
}

func test01(p *Student) {
	p.id = 666
}

func main() {
	s := Student{1, "mike", 'm', 18, "bj"}

	test01(&s) //地址传递（引用传递），形参可以改变实参
	fmt.Println("main: ", s)
	/*结果
	main:  {666 mike 109 18 bj}
	*/

}
