package main

import "fmt"

/*
Go语言中的map(映射、字典)是一种内置的数据结构，
它是一个无序的key-value对的集合
比如以身份证号作为唯一键来标识一个人的信息

map[keyType]valueType

在一个map里所有的键都是唯一的，而且必须是支持==和!=操作符的类型，
切片、函数以及包含切片的结构类型这些类型由于具有引用语义，不能作为映射的键，使用这些类型会造成编译错误
*/

func main() {
	//定义一个变量，类型为map[int]string
	var m1 map[int]string
	fmt.Println("m1 = ", m1)
	//对于map只有len，没有cap
	fmt.Println("len = ", len(m1))

	m2 := make(map[int]string)
	fmt.Println("m2 = ", m2)
	fmt.Println("len = ", len(m2))

	//可以通过make创建，可以指定长度，只是指定了容量，但是里面却是一个数据没有
	m3 := make(map[int]string, 2)
	fmt.Println("m3 = ", m3)
	fmt.Println("len = ", len(m3))
	/*
		结果
		m1 =  map[]
		len =  0
		m2 =  map[]
		len =  0
		m3 =  map[]
		len =  0
	*/
	m3[1] = "mike"
	m3[2] = "go"
	m3[3] = "php"
	fmt.Println("m3 = ", m3)
	fmt.Println("len = ", len(m3))
	/*
		结果
		m3 =  map[1:mike 2:go 3:php]
		len =  3
	*/

	//初始化  下面是常用方法
	//键值是唯一的
	m4 := map[int]string{1: "mike", 2: "go", 3: "php"}
	fmt.Println("m4 = ", m4)
	/*
		结果
		m4 =  map[1:mike 2:go 3:php]
	*/

}
