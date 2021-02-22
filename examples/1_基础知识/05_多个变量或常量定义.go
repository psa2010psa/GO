package main

import "fmt"

func main() {
	var a int
	var b float64
	a, b = 10, 3.14
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	var (
		c int     = 1
		d float64 = 2.0
	)
	c, d = 20, 3.33
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)

	const (
		i int     = 10
		j float64 = 2.22
	)
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
	//自动推导类型
	const (
		p = 90
		q = 2.9
	)
	fmt.Println("p = ", p)
	fmt.Println("q = ", q)

}
