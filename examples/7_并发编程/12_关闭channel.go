package main

import (
	"fmt"
)

func main() {
	ch := make(chan int) //创建一个无缓存channel

	//新建一个goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //往管道写数据
		}
		//不需要再写数据时，关闭channel
		close(ch)
	}()

	for num := range ch {
		fmt.Println("num = ", num)
	}
	/*结果
	num =  0
	num =  1
	num =  2
	num =  3
	num =  4
	*/
}

func main01() {
	ch := make(chan int) //创建一个无缓存channel

	//新建一个goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //往管道写数据
		}
		//不需要再写数据时，关闭channel
		close(ch)
		//ch <- 666 //关闭channel后无法在发送数据，可以继续向channel接收数据 err: panic: send on closed channel
		//对于nil channel 无论收发都会被阻塞
	}()

	for {
		//如果ok为true，说明管道没有关闭
		if num, ok := <-ch; ok == true {
			fmt.Println("num = ", num)

		} else { //管道关闭
			break
		}
	}
	/*
		num =  0
		num =  1
		num =  2
		num =  3
		num =  4
	*/
}
