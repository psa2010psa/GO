package main

import (
	"fmt"
	"time"
)

/*
goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。goroutine奉行通过通信来共享内存，而不是共享内存来通信。

引用类型channel是CSP(通信顺序程序)模式的具体实现，用于多个goroutine通讯。其内部实现了同步，确保并发安全。
*/

//全局变量。创建一个channel
var ch = make(chan int)

//定义一个打印机，参数为字符串，按每个字符打印
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

//Person1执行完后，才能到Person2执行
func Person1() {
	Printer("hello")
	ch <- 666 //给管道写数据
}

func Person2() {
	<-ch //从管道取数据，接受，如果通道没有数据他就会阻塞
	Printer("world")
}

func main() {
	//新建2个协程，代表2个人，2个人同时使用打印机
	go Person1()
	go Person2() //输出：hwoerlllod   资源竞争产生问题

	//特地不让主协程结束，死循环
	for {

	}
}
