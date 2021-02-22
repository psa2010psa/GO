package main

import "fmt"

func main() {
	m1 := map[int]string{1: "mike", 2: "yoyo"}
	//赋值，如果已经存在的key值，修改内容
	fmt.Println("m1 = ", m1) //结果：m1 =  map[1:mike 2:yoyo]
	m1[1] = "php"
	fmt.Println("m1 = ", m1) //结果：1 =  map[1:php 2:yoyo]

	//追加赋值，map底层自动扩容，和append类似
	m1[3] = "go"
	fmt.Println("m1 = ", m1) //结果：m1 =  map[1:php 2:yoyo 3:go]
}
