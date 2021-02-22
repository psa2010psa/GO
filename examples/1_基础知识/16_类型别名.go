package main

import "fmt"

func main() {
	//给int64起一个别名叫bigint
	type bigint int64
	var a bigint //等价于 var a int64
	fmt.Printf("a type is %T\n", a)
	type (
		long int64
		char byte
	)
	var b long = 1
	var ch char = 'a'
	fmt.Printf("b = %d\n", b)
	fmt.Printf("ch = %c\n", ch)

}
