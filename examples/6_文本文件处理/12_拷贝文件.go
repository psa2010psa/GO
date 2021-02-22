package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	list := os.Args //获取命令行参数
	if len(list) != 3 {
		fmt.Println("usage: xxx srcFile dstFile")
		return
	}

	srcFileName := list[1]
	dstFileName := list[2]
	if srcFileName == dstFileName {
		fmt.Println("源文件名字和目的文件不能相同")
		return
	}

	//只读方式打开源文件
	sF, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//新建目的文件
	dF, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}

	//操作完毕，需要关闭文件
	defer sF.Close()
	defer dF.Close()

	//核心处理，从源文件读取内容,往目的文件写，读多少写多少
	buf := make([]byte, 4*1024) //4k大小临时缓冲区
	for {
		n, err3 := sF.Read(buf)
		if err3 != nil {
			if err3 == io.EOF { //文件读取完毕
				break
			}
			fmt.Println("err3 = ", err3)
			return
		}
		//往目的文件写，读多少写多少
		dF.Write(buf[:n])
	}

	/*
		运行 go build 12_拷贝文件.go 生成可执行程序
		然后 运行 12_拷贝文件 源文件.png 目标文件.png
		最终实现拷贝

	*/

}
