package main

import (
	"encoding/json"
	"fmt"
)

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
	m := make(map[string]interface{}, 4)       //创建一个map
	err := json.Unmarshal([]byte(jsonBuf), &m) //第二个参数必须是地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("m = %+v\n", m)
	//结果：m = map[company:itcast isok:true price:666.666 subjects:[Go c++ php Python]]

	// var str string
	// str = m["company"]
	// fmt.Println("str = ", str) //无法转换
	//err: cannot use m["company"] (type interface {}) as type string in assignment: need type assertion

	var str string
	//类型断言 值，它是value类型
	for key, value := range m {
		//fmt.Printf("%v =====> %v\n", key, value)
		switch data := value.(type) {
		case string:
			str = data
			fmt.Printf("map[%s]的值类型为string，value = %s\n", key, str)
		case bool:
			fmt.Printf("map[%s]的值类型为bool，value = %v\n", key, data)
		case float64:
			fmt.Printf("map[%s]的值类型为float64，value = %f\n", key, data)
		case []string: //没有这个类型
			fmt.Printf("map[%s]的值类型为[]string，value = %v\n", key, data)
		case []interface{}:
			fmt.Printf("map[%s]的值类型为[]string，value = %v\n", key, data)
		}
	}

}
