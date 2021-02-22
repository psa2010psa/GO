package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	switch num {
	case 1:
		fmt.Println("这里是1")
		//break //go语言保留了break关键字，跳出switch语句，不写，默认就包含break
		fallthrough //不跳出switch语句，后面的无条件执行
	case 2:
		fmt.Println("这里是2")
		fallthrough
	case 3:
		fmt.Println("这里是3")
	case 4:
		fmt.Println("这里是4")
	default:
		fmt.Println("这里是默认：", num)
	}
}
