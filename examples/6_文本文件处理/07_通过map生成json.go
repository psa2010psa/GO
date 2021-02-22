package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//创建一个map
	m := make(map[string]interface{}, 4)
	m["company"] = "itcast"
	m["subjects"] = []string{"Go", "c++", "php", "Python"}
	m["isok"] = true
	m["price"] = 666.666
	//编码成json
	buf, err := json.Marshal(m)
	//buf, err := json.MarshalIndent(m, "", " ") //格式化编码
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf))
	//结果：buf =  {"company":"itcast","isok":true,"price":666.666,"subjects":["Go","c++","php","Python"]}

	/*格式化编码 结果：
	buf =  {
	 "company": "itcast",
	 "isok": true,
	 "price": 666.666,
	 "subjects": [
	  "Go",
	  "c++",
	  "php",
	  "Python"
	 ]
	}
	*/
}
