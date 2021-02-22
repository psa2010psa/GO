package main

import "fmt"

func main() {
	var p *int
	//p = nil
	fmt.Println("p = ", p)
	//结果 p = <nil>

	// *p = 666 //err,因为p没有合法指向

	var a int
	p = &a
	*p = 666
	fmt.Println("a = ", a)
	//结果：a = 666
}
