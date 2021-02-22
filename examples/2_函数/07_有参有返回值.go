package main

import "fmt"

func MaxAndMin(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return
}

func main() {
	max, min := MaxAndMin(10, 20)
	fmt.Printf("max = %d, min = %d\n", max, min)

	a, _ := MaxAndMin(10, 20)
	fmt.Printf("a = %d\n", a)
}
