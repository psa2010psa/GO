package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "43.14 567 agsdg 1.23 7.  8.99 1sd1jg1 6.66"

	//解析正则表达式 +匹配前一个字符的1次或多次
	reg := regexp.MustCompile(`\d+\.\d+`)
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}

	//提取关键信息
	//result := reg.FindAllString(buf, -1)
	result := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println("result = ", result)
	//结果：result =  [[43.14] [1.23] [8.99] [6.66]]

}
