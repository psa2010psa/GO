package main

import "fmt"

func main() {
	str := "abc"
	//通过for循环打印每一个字符
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}
	//迭代打印每一个元素，默认返回2个值：一个是元素的位置，一个是元素本身
	for i, data := range str {
		fmt.Printf("str[%d] = %c\n", i, data)
	}

	for i := range str {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}

	for i, _ := range str {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}

}
