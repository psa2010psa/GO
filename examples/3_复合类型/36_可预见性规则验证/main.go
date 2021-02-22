package main

import (
	"fmt"
	"test"
)

/*
go语言可预见性规则验证
1）如果想使用别的包的函数、结构体类型、结构体成员
   函数名，类型名，结构体成员变量名，首字母必须大写，可见

   如果首字母是小写，只能在同一个包里使用
*/

func main() {
	//包名.函数名
	test.MyFunc() //test.myFunc 会编译错误
	//结果: this is MyFunc ========

	//包名.结构体里类型名
	var s test.Stu // test.stu 会编译错误
	//s.id = 666 //err s.id undefined (cannot refer to unexported field or method id)

	s.Id = 666
	fmt.Println("s = ", s)
	//结果：s = {666}
}
