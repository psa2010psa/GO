package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for { //for后面不写任何东西，这个循环条件永远为真，死循环
		i++
		time.Sleep(time.Second) //延迟1s
		if i == 5 {
			continue //跳过本次循环，下一次继续
		}
		if i == 7 {
			break //跳出循环，如果嵌套多个循环，跳出最近的那个内循环
		}

		fmt.Println("i = ", i)
	}
}
