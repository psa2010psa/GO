package main

import (
	"fmt"
	"runtime"
)

/*
rutime.GOMAXPROCS()用来设置可以并行计算的CPU核数的最大值，并返回之前的值。
*/

func main() {
	n := runtime.GOMAXPROCS(2) //runtime.GOMAXPROCS(1)
	fmt.Println("n = ", n)
	for {
		go fmt.Print(1)

		fmt.Print(0)
	}
}
