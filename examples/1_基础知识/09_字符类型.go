package main

import "fmt"

func main() {
	var ch byte
	ch = 97
	//fmt.Println("ch = ", ch)

	//格式化输出，%c以字节方式打印，%d以整型方式打印
	fmt.Printf("%c, %d\n", ch, ch)
	ch = 'a' //字符  使用单引号
	fmt.Printf("%c %d\n", ch, ch)

	fmt.Printf("大写：%d, 小写：%d\n", 'A', 'a')
	fmt.Printf("大写转小写：%c\n", 'A'+32)
	fmt.Printf("小写转大写：%c\n", 'a'-32)
	//'\n'以反斜杠开头的字符是转义字符，'\n'代表换行
	fmt.Printf("hello world%c", '\n')
	fmt.Printf("learn go")
}
