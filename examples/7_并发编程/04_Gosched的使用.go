package main

/*
runtime.Gosched()用于让出CPU时间片，让出当前goroutine的执行权限，
调度器安排其他等待的任务执行，并在下次某个时间从该位置恢复执行。

这就像跑接力赛，A跑了一会碰到代码runtime.Gosched()就把接力棒交给B了，A歇着了，B继续跑。

*/
import (
	"fmt"
	//"time"
	"runtime"
)

func main() {

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}
	}()
	for i := 0; i < 2; i++ {
		//让出时间片，先让别的协程执行，它执行完，再回来执行次协程
		runtime.Gosched()
		fmt.Println("hello") //如果不让出时间片，执行输出2遍hello，子协程不会执行
	}
	/*结果  输出顺序不固定
	hello
	hello
	go
	go
	go
	go

	*/
}
