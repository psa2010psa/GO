package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc azc a7c aac 888 a9c tac"

	//1)解析规则，它会解析正则表达式，如果成功返回解析器
	//reg1 := regexp.MustCompile(`a.c`)
	//reg1 := regexp.MustCompile(`a[0-9]c`)
	reg1 := regexp.MustCompile(`a\dc`)
	if reg1 == nil { //解析失败，返回nil
		fmt.Println("regexp err")
		return
	}

	//2)根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(buf, -1) //all
	fmt.Println("result1 = ", result1)             //result1 =  [[abc] [azc] [a7c] [aac] [a9c]]

	result1_1 := reg1.FindAllStringSubmatch(buf, 1) //第一个
	fmt.Println("result1_1 = ", result1_1)          //result1_1 =  [[abc]]
}
