package main //必须

import "fmt"

//函数名首字母小写即为private，大写即为public

// 有参数无返回值函数的定义，普通参数列表
// 定义函数时，在函数名后面（）定义的参数叫形参
// 参数传递， 只能由实参传递给形参，不能反过来，单向传递
func Myfunc01(a int) {
	// a = 111
	fmt.Println("a = ", a)
}

func Myfunc02(a int, b int) {
	// a = 111
	fmt.Println("a +b = ", a+b)
}

// 形参简便写法,a、b都是int
func Myfunc03(a, b int, c string, d, e float64) {
}

func main() {
	//有参数无返回值函数调用：函数名(所需参数)
	Myfunc01(666)

	Myfunc02(3, 2)
}
