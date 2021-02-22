package main

import (
	"fmt"
	"strings"
)

func main() {
	//Contains 包含
	//"hellogo"中是否包含 “hello”,包含返回true，不包含返回false
	fmt.Println(strings.Contains("hellogo", "hello")) //true
	fmt.Println(strings.Contains("hellogo", "abc"))   //false

	//Join 组合
	s := []string{"abc", "hello", "mike", "go"}
	buf := strings.Join(s, "|")
	fmt.Println("buf = ", buf) //buf =  abc|hello|mike|go

	//Index, 查找子串的位置
	fmt.Println(strings.Index("abcdhello", "hello")) // 4
	fmt.Println(strings.Index("abcdhello", "Hello")) // -1

	//Repeat 重复(几次)
	buf2 := strings.Repeat("go", 3)
	fmt.Println("buf2 = ", buf2) //buf2 =  gogogo

	//Split 以指定的分隔符拆分
	buf = "hell@abc@go@mike"
	s2 := strings.Split(buf, "@")
	fmt.Println("s2 = ", s2) //s2 =  [hell abc go mike]

	//Trim 去掉两头的字符
	fmt.Println(strings.Trim("###are u ok?####", "#")) //are u ok?

	//Fields 去掉空格，把元素放入切片中
	s3 := strings.Fields("hello go world")
	fmt.Println("s3 = ", s3) //切片  s3 =  [hello go world]

	s4 := strings.Replace("helloabc abc", "abc", "go", 1)
	fmt.Println("s4 = ", s4)

}
