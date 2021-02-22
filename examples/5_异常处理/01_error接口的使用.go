package main

import (
	"errors"
	"fmt"
)

func main() {
	//var err1 error = fmt.Errorf("%s","this is normol err")
	err1 := fmt.Errorf("%s", "this is normol err")
	fmt.Println("err1 = ", err1) //err1 =  this is normol err

	err2 := errors.New("this is normol err2")
	fmt.Println("err2 = ", err2) //err2 =  this is normol err2
}
