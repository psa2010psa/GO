package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个无缓存的channel
	ch := make(chan int, 0)

	//len(ch)缓冲区剩余数据个数，cap(ch)缓冲区大小
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch)) //放在子协程里打印也一直是0，0

	//新建协程
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("子协程 i = %d\n", i)
			ch <- i //往chan写内容
		}
	}()
	//延时
	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-ch //读管道中内容，没有内容，阻塞
		fmt.Println("num = ", num)
	}

	/*结果
		先打印
		len(ch) = 0, cap(ch) = 0
		子协程 i = 0
	    在打印（会阻塞）
		num =  0
		子协程 i = 1
		子协程 i = 2
		num =  1
		num =  2
	*/
}
