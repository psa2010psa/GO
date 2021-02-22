package main

import (
	"fmt"
	"time"
)

/*
并行：指在同一时刻，有多条指令在多个处理器上同时执行

并发：指在同一时刻只能有一条指令执行，但多个进程指令被快速的轮换执行，
使得在宏观上具有多个进程同时执行的效果，但在微观上并不是同时执行的，只是把时间分成若干段，使多个进程快速交替的执行
(时间片轮转)

并行是两个队列同时使用两台咖啡机
并发是两个队列交替使用一台咖啡机
*/

func newTask() {
	for {
		fmt.Println("this is a newTask")
		time.Sleep(time.Second) //延迟1s
	}
}

func main() {
	go newTask() //新建一个协程, 相当于新建一个任务

	for {
		fmt.Println("this is a main goruntine")
		time.Sleep(time.Second) //延迟1s
	}

	/*结果
	this is a main goruntine
	this is a newTask
	上面2个一组，每一秒输出一次
	谁先执行看性能和调度
	*/
}
