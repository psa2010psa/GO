package main

import "fmt"

//实现1+2+3+4+...+100

func test01() (sum int) {
	for i := 1; i < 101; i++ {
		sum += i
	}
	return
}

func test02(i int) int {
	if i == 1 {
		return 1
	}
	return i + test02(i-1)
}

func test03(i int) int {
	if i == 100 {
		return 100
	}
	return i + test03(i+1)
}

func main() {
	var sum int
	// sum = test01()
	// fmt.Println("sum = ", sum)

	// sum = test02(100)
	// fmt.Println("sum = ", sum)

	sum = test03(1)
	fmt.Println("sum = ", sum)

}
