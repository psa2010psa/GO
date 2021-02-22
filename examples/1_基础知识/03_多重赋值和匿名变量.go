package main

import "fmt"

func test() (a, b, c int) {
	return 1, 2, 3
}

func main() {
	a, b := 10, 20
	var tmp int
	tmp = a
	a = b
	b = tmp
	fmt.Printf("a = %d, b= %d\n", a, b)

	//交换2个变量的值
	i, j := 30, 40
	i, j = j, i
	fmt.Printf("i = %d, j= %d\n", i, j)

	//_匿名变量，丢弃数据不处理，_匿名变量配合函数返回值使用，才有优势
	_, tmp = i, j
	fmt.Println("tmp2 = ", tmp)

	var e, f, g int
	e, f, g = test()
	fmt.Printf("%d, %d, %d\n", e, f, g)

	_, f, _ = test()
	fmt.Printf("f = %d\n", f)

}
