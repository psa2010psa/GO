package main

import (
	"fmt"
	"time"
)

//验证time.NewTimer()，时间到了，只会响应一次
func main() {
	timer := time.NewTimer(1 * time.Second)
	for {
		<-timer.C
		fmt.Println("时间到")
	}
	/*结果：死锁
	时间到
	fatal error: all goroutines are asleep - deadlock!
	*/
}

func main01() {
	//创建一个定时器，设置时间为2s，2s后往time通道写内容（当前时间）
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间：", time.Now())

	//2s后，往timer.C写数据，有数据后，就可以读取
	t := <-timer.C //channel没有数据前后阻塞
	fmt.Println("t = ", t)
	/*结果
	当前时间： 2021-02-02 17:23:01.346829 +0800 CST m=+0.000276865
	2s后
	t =  2021-02-02 17:23:03.347784 +0800 CST m=+2.001308688
	*/
}
