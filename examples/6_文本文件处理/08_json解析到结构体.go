package main

import (
	"encoding/json"
	"fmt"
)

//成员变量名首字母必须大写
type IT struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func main() {
	jsonBuf := `
	{
		"company": "itcast",
		"subjects": [
		 "Go",
		 "c++",
		 "php",
		 "Python"
		],
		"isok": true,
		"price": 666.666
	   }
`
	var tmp IT                                   //定义一个结构体
	err := json.Unmarshal([]byte(jsonBuf), &tmp) //第二个参数必须是地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("tmp = %+v\n", tmp)
	//结果：tmp = {Company:itcast Subjects:[Go c++ php Python] IsOk:true Price:666.666}

	type IT2 struct {
		Subjects []string `json:"subjects"`
	}
	var tmp2 IT2
	err = json.Unmarshal([]byte(jsonBuf), &tmp2) //第二个参数必须是地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("tmp2 = %+v\n", tmp2)
	//结果：tmp2 = {Subjects:[Go c++ php Python]}
}
