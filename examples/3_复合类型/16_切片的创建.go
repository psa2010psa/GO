package main

import "fmt"

func main() {
	//自动推导类型
	s1 := []int{1, 2, 3, 4}
	fmt.Println("s1 = ", s1)

	//借助make函数，格式 make(切片类型,长度,容量)
	s2 := make([]int, 5, 10)
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))
	//结果：len = 5, cap = 10

	//没有指定容量，容量和长度一样
	s3 := make([]int, 5)
	fmt.Printf("len = %d, cap = %d\n", len(s3), cap(s3))
	//结果：en = 5, cap = 5
}

func main01() {
	//切片和数组的区别
	//数组[]里面的长度是固定的一个常量，数组不能修改长度，len和cap永远都是5
	a := [5]int{}
	fmt.Printf("len = %d, cap = %d\n", len(a), cap(a))

	//切片，[]里面为空，或者为... 切片的长度或容量可以不固定
	s := []int{}
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))

	s = append(s, 11) //给切片末尾追加一个成员
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
	/*结果
	len = 5, cap = 5
	len = 0, cap = 0
	len = 1, cap = 1

	*/
}
