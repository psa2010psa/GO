package main

import "fmt"

func test(a int) {
	if a == 1 {
		fmt.Println("a = ", a)
		return
	}
	test(a - 1)
	fmt.Println("a = ", a)
}

func main() {
	test(3)
	fmt.Println("main")
	// a =1  a=2  a=3   main
}
