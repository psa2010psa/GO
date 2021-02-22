package main

import "fmt"

//引用传递
func test(m map[int]string) {
	delete(m, 1)
}

func main() {
	m := map[int]string{1: "mike", 2: "go", 3: "php"}
	fmt.Println("m = ", m)

	test(m) //在函数内部删除某个key

	fmt.Println("m = ", m)
	/*结果
	m =  map[1:mike 2:go 3:php]
	m =  map[2:go 3:php]
	*/

}
