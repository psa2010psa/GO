package main

import (
	"fmt"
	"os"
)

/*
设备文件：
	屏幕(标准输出设备) fmt.Println()   往标准输入设备写内容
	键盘(标准输入设备) fmt.Scan()      从标准输入设备读取内容
磁盘文件，放在存储设备上的文件：
	1）文本文件 	以记事本打开，看到内容（不是乱码）
	2）二进制文件	已记事本打开，能看到内容（是乱码）

为什么需要文件？ 内存掉电丢失，程序结束，内存中的内容消失
			  文件放在磁盘，程序结束，文件还是存在
*/

func main() {
	//os.Stdout.Close() //关闭后，无法输出
	//fmt.Println("are u ok?") //往标准输出设备(屏幕)写内容
	//标准设备文件，默认已经打开，用户可以直接使用
	//os.Stdout
	os.Stdout.WriteString("are u ok?\n")

	//os.Stdin.Close()  //关闭后无法输入
	var a int
	fmt.Println("请输入……")
	fmt.Scan(&a) //从标准输入设备中读取内容，放在a中
	fmt.Println("a = ", a)
}
