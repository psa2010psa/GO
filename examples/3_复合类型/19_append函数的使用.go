package main

import "fmt"

func msg(x []int) {
	fmt.Printf("s = %v, len=%d, cap=%d \n", x, len(x), cap(x))
}

func main() {
	s1 := []int{}
	msg(s1)
	//结果：s1 = [], len=0, cap=0

	//在原切片的末尾添加元素
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	msg(s1)
	//结果：s1 = [1 2 3], len=3, cap=4

	s2 := []int{1, 2, 3}

	s2 = append(s2, 5)
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	msg(s2)
	//结果：s = [1 2 3 5 5 5], len=6, cap=6

}
