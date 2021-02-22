package main

//忽略此包
import _ "fmt" // _操作其实是引入该包，而不是直接使用包里面的函数，而是调用了该包里面的init函数

func main() {

}

/*
//给包起别名
import io "fmt"

func main() {
	io.Println("this is a test")
}
*/

/*
// .操作
import . "fmt"

func main() {
	Println("this is a test")
}
*/

/*
//方法一
//import "fmt" //导入包，必须使用，否则编译不过
//import "os"

//方法二
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("this is a test")
	fmt.Println("os.Args = ", os.Args)
}
*/
