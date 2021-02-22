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

func test01(s Student) {
	s.id = 666
	fmt.Println("test01: ", s)
}

func main() {
	s := Student{1, "mike", 'm', 18, "bj"}

	test01(s) //值传递，形参无法改变实参
	fmt.Println("main: ", s)
	/*结果
	test01:  {666 mike 109 18 bj}
	main:  {1 mike 109 18 bj}
	*/

}
