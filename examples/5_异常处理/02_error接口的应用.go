package main

import (
	"errors"
	"fmt"
)

func MyDiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("分母不能为0")
	} else {
		result = a / b
	}

	return
}

func main() {
	result, err := MyDiv(10, 0)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("result = ", result)
	}
	//err =  分母不能为0

}
