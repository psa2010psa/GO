package main

import (
	"encoding/json"
	"fmt"
)

//成员变量名首字母必须大写
type IT struct {
	Company  string   `json:"company"` //二次编码
	Subjects []string //`json:"-"`       //此字段不会输出到屏幕
	IsOk     bool     //`json:",string"` //以字符串格式输出
	Price    float64  //`json:",string"`
}

func main() {
	//定义一个结构体变量，同时初始化
	s := IT{"itcast", []string{"Go", "c++", "php", "Python"}, true, 666.666}

	//编码，根据内容生成json文本
	//buf, err := json.Marshal(s)
	buf, err := json.MarshalIndent(s, "", " ") //格式化编码
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf))
	//结果：buf =  {"company":"itcast","Subjects":["Go","c++","php","Python"],"IsOk":true,"Price":666.666}

	/*格式化编码 结果：
	buf =  {
	 "company": "itcast",
	 "Subjects": [
	  "Go",
	  "c++",
	  "php",
	  "Python"
	 ],
	 "IsOk": true,
	 "Price": 666.666
	}
	*/
}
