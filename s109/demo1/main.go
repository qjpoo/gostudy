package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 断点复制

	//bs := []byte{'A',99,'0',0,'B'}
	//fmt.Println(string(bs))
	/*
	seek(offset int64, whence int) (int64, error)
	offset 偏移量
	whence 如何设置
	Seek 将 offset 置为下一个 Read 或 Write 的偏移量 ，
	它的解释取决于 whence：
		0 表示相对于文件的起始处，
		1 表示相对于当前的偏移，
		2 表示相对于其结尾处。
	Seek 返回新的偏移量和一个错误，如果有的话。
	对负数偏移量进行 Seek 会产生错误。
	 */

	fileName := "seek.txt"
	file, err := os.OpenFile(fileName,os.O_RDWR,0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// abcdefg
	bs := []byte{0} // 定义的是一个空的字节
	//fmt.Println(string(bs) + "A")
	file.Read(bs)  // 刚开始,光标在最前面,就是在a字母的前面,在读取一个字节,所以是a
	fmt.Println(string(bs)) // a

	file.Seek(4,io.SeekStart)  // 以文件开头计算,距离文件开头4个字节,此时光标在d的后面
	file.Read(bs)  // 在读一个字节,所以是e了
	fmt.Println(string(bs)) // e

	file.Read(bs)  // 此时光标在e的后面
	fmt.Println(string(bs))   // 在读一个字节,就是  f

	file.Seek(2,io.SeekStart) // 光标相对于文件开头2个偏移量
	file.Read(bs)  // 在读一个
	fmt.Println(string(bs))  // c

	// 当前的光标在C的后面
	file.Seek(2,io.SeekCurrent) // 在移动2个
	file.Read(bs)  // 在读一字节
	fmt.Println(string(bs))  // f

	// 在向文件后面写一些字节
	file.Seek(0, io.SeekEnd) // 光标定位在文件最后
	file.WriteString("hello, world ...")

}
