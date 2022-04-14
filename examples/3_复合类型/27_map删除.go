package main

import "fmt"

func main() {
	m := map[int]string{1: "mike", 2: "go", 3: "php"}
	fmt.Println("m = ", m)

	delete(m, 1) //删除key为1的内容
	fmt.Println("m = ", m)

	/*结果
	m =  map[1:mike 2:go 3
	m =  map[2:go 3:php]
	*/
}
