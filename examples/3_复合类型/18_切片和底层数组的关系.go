package main

import "fmt"

func msg(x []int) {
	fmt.Printf("s1 = %v, len(x)=%d, cap(x)=%d \n", x, len(x), cap(x))
}

/*
数组的缺点，初始化大小固定，传参把整个数组拷贝过去

数组的长度在定义之后无法再次修改；数组是值类型，每次传递都将产生一份副本。

切片并不是数组或数组指针，它通过内部指针和相关属性引用数组片段，以实现变长方案。
*/

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := a[2:5] //从a[2]开始，取3个元素
	msg(s1)

	s1[1] = 666
	fmt.Println("s1 = ", s1)
	fmt.Println("a = ", a)
	/*结果
	s1 =  [2 666 4]
	a =  [0 1 2 666 4 5 6 7 8 9]
	*/

	//另外新切片
	s2 := s1[2:7]
	s2[2] = 777
	fmt.Println("s2 = ", s2)
	fmt.Println("s1 = ", s1)
	fmt.Println("a = ", a)
	/*结果
	s2 =  [4 5 777 7 8]
	s1 =  [2 666 4]
	a =  [0 1 2 666 4 5 777 7 8 9]
	*/

}
