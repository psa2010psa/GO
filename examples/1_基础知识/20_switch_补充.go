package main

import "fmt"

func main() {
	//支持一个初始化语句， 初始化语句和变量本身，以分号分隔
	switch num := 4; num {
	case 1:
		fmt.Println("这里是1")
	case 2:
		fmt.Println("这里是2")
	case 3, 4, 5:
		fmt.Println("这里是3,4,5")
	case 6:
		fmt.Println("这里是4")
	default:
		fmt.Println("这里是默认：", num)
	}

	score := 85
	switch { //可以没有条件
	case score > 90: //case后面可以放条件
		fmt.Println("优秀")
	case score > 80:
		fmt.Println("良好")
	case score > 70:
		fmt.Println("一般")
	default:
		fmt.Println("其他")
	}
}
