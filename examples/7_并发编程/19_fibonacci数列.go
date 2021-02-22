package main

import (
	"fmt"
)

/*
Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。
*/

//fibonacci 1 1 2 3 5 8

func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		//监听channel数据的流动
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag = ", flag)
			return
		}
	}
}

func main() {
	ch := make(chan int)    //数字通信
	quit := make(chan bool) //程序是否结束

	//消费者，从channel读取内容
	//新建协程
	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch
			fmt.Println("num = ", num)
		}
		//可以停止
		quit <- true
	}()
	//生产者，生产数字，写入channel
	fibonacci(ch, quit)
	/*结果
	num =  1
	num =  1
	num =  2
	num =  3
	num =  5
	num =  8
	num =  13
	num =  21
	flag =  true
	*/
}
