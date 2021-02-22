package main

import (
	"fmt"
	"runtime"
)

/*
runtime.Goexit()将立即终止当前goroutine执行，调度器确保所有已注册 defer延迟调用被执行。
*/

func test() {
	defer fmt.Println("ccccccc")
	//return 终止此函数
	runtime.Goexit() //终止所在的协程
	fmt.Println("ddddddddddddd")
}

func main() {
	//创建新建的协程
	go func() {
		fmt.Println("aaaaaaaaa")
		//调用了别的函数
		test()

		fmt.Println("bbbbbbbbb")
	}()

	//特地写一个死循环，目的不让主协程结束
	for {

	}

	/*结果
	aaaaaaaaa
	ccccccc
	不会输出bbbbbbbbb
	*/

}
