package main //必须有一个main包

import "fmt"

func main() {
	//变量  程序运行期间，可以改变的量，变量声明需要var
	//常量  程序运行期间，不可以改变的量，常量声明需要const
	const a int = 10
	//a = 20 //err 常量不允许修改
	fmt.Println("a = ", a)

	const b = 20 //没有使用 :=
	fmt.Println("b = ", b)
	fmt.Printf("b type is %T\n", b)
}
