package main

import (
	"fmt"
	"time"
)

//主协程退出了，其它子协程也要跟着退出
func main() {

	go func() {
		i := 0
		for {
			i++
			fmt.Println("子协程 i = ", i)
			time.Sleep(time.Second)
		}
	}() //别忘了()

	i := 0
	for {
		i++
		fmt.Println("main i = ", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}

	/*结果
	main i =  1
	子协程 i =  1
	子协程 i =  2
	main i =  2
	子协程 i =  3

	主协程退出了，子协程也要跟着退出
	*/

}
