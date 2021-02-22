package main

import "fmt"

func main() {
	m := map[int]string{1: "mike", 2: "go", 3: "php"}

	//第一个返回值为key，第二个返回值为value，遍历结果是无序的
	for key, value := range m {
		fmt.Printf("%d =====> %s\n", key, value)
	}
	/*结果
	1 =====> mike
	2 =====> go
	3 =====> php
	*/

	//如何判断一个key值是否存在
	//第一个返回值为key所对应的value，第二个返回值为key是否存在的条件，存在ok为true
	value, ok := m[1] //结果 m1 =  mike    value,ok := m[0] 结果：key不存在
	if ok == true {
		fmt.Println("m1 = ", value)
	} else {
		fmt.Println("key不存在")
	}

}
