package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[0:3:5]
	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s)) //长度 3-0
	fmt.Println("cap(s) = ", cap(s)) //容量 5-0
	/*结果
	s =  [1 2 3]
	len(s) =  3
	cap(s) =  5
	*/

	s = a[1:4:5]
	fmt.Println("s = ", s)           //从下标1开始，取4-1=3个
	fmt.Println("len(s) = ", len(s)) //长度 4-1
	fmt.Println("cap(s) = ", cap(s)) //容量 5-1
	/*结果
	s =  [2 3 4]
	len(s) =  3
	cap(s) =  4

	*/
}
