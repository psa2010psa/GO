package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个有缓存的channel
	ch := make(chan int, 3)

	//len(ch)缓冲区剩余数据个数，cap(ch)缓冲区大小
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch)) //放在子协程里打印也一直是0，0

	//新建协程
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i //往chan写内容
			fmt.Printf("子协程[%d]: len(ch) = %d, cap(ch) = %d\n", i, len(ch), cap(ch))
		}
	}()
	//延时
	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		num := <-ch //读管道中内容，没有内容，阻塞
		fmt.Println("num = ", num)
	}

	/*结果
	每次运行结果不一定
	len(ch) = 0, cap(ch) = 3
	子协程[0]: len(ch) = 1, cap(ch) = 3
	子协程[1]: len(ch) = 2, cap(ch) = 3
	子协程[2]: len(ch) = 3, cap(ch) = 3
	num =  0
	子协程[3]: len(ch) = 3, cap(ch) = 3
	子协程[4]: len(ch) = 3, cap(ch) = 3
	num =  1
	num =  2
	num =  3
	num =  4
	num =  5
	子协程[5]: len(ch) = 0, cap(ch) = 3
	子协程[6]: len(ch) = 0, cap(ch) = 3
	子协程[7]: len(ch) = 1, cap(ch) = 3
	num =  6
	num =  7
	num =  8
	子协程[8]: len(ch) = 2, cap(ch) = 3
	子协程[9]: len(ch) = 0, cap(ch) = 3
	num =  9
	*/
}
