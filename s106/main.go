package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// io包(input output) io包的读
	//　接口名一般是以er为后缀,比如Reader


	// 1. 打开文件
	file, err := os.Open("123.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 3. 关闭文件
	defer file.Close()

	// 2. 读取数据
/*
	fmt.Println("------------1------------")
	// 第一次
	n, err :=  file.Read(bs)
	fmt.Println(n)
	fmt.Println(err)
	fmt.Println(bs)
	fmt.Println(string(bs))

	fmt.Println("------------2------------")
	// 第二次
	n, err =  file.Read(bs)
	fmt.Println(n)
	fmt.Println(err)
	fmt.Println(bs)
	fmt.Println(string(bs))

	fmt.Println("------------3------------")
	// 第三次
	n, err =  file.Read(bs)
	fmt.Println(n)  // 2
	fmt.Println(err)  // nil
	fmt.Println(bs)  // [57 48 55 56]
	fmt.Println(string(bs)) // 9078


	fmt.Println("------------4------------")
	// 第四次
	n, err =  file.Read(bs)
	fmt.Println(n) // 0
	fmt.Println(err) // EOF
	fmt.Println(bs) // [57 48 55 56]
	fmt.Println(string(bs))  // 9078
*/
	bs := make([]byte, 1024)
	for {
		n, err := file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("文件读完")
			break
		}
		fmt.Println(string(bs[:n]))
	}


}
