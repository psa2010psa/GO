package main

import "fmt"

func main() {
	srcSlice := []int{1, 2}
	dstSlice := []int{6, 6, 6, 6, 6}

	//copy(目标，原)
	copy(dstSlice, srcSlice)
	fmt.Println("dst = ", dstSlice)
	/*结果
	dst =  [1 2 6 6 6]
	*/
}
