package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"reflect"
)

func main() {
	/*
	file操作
	 */
	pwd, err := os.Getwd()
	fmt.Println(pwd, err)
	fmt.Println(pwd + "\\new.txt")

	fileInfo, err := os.Stat("C:\\Users\\瞿健\\go\\src\\awesomeProject\\new.txt")
	if err != nil {
		fmt.Println(err)
		return
	}else {
		fmt.Println(reflect.TypeOf(fileInfo)) // *os.fileStat
	}

	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.Mode())

	//文件操作
	/*
	相对路径: 相对于当前的工程来说的
	绝对路径
	 */
	isAbs := filepath.IsAbs("C:\\Users\\瞿健\\go\\src\\awesomeProject\\new.txt")
	fmt.Println(isAbs)
	isAbs1 := filepath.IsAbs("new.txt")
	fmt.Println(isAbs1)
	isAbs2, err := filepath.Abs("new.txt")
	fmt.Println(isAbs2)

	fmt.Println(path.Base("/abc/a"))
	fmt.Println(path.Join("/root/abc/test/123.txt", "..")) //取的当前路径的父目录

	//os.mkdir 创建一个不存在的目录,如果存在会报错,只会创建一级
	//os.mkdirall  mkdir -p  递归创建


	// 创建文件,如果文件不存会创建,如果文件存会,置空
	f1, err := os.Create("C:\\Users\\瞿健\\go\\src\\awesomeProject\\123.txt")
	fmt.Println(f1, err)

	f2, err := os.Create("abc.txt") // 相对路径来创建的文件 , 以工程为参考
	fmt.Println(f2, err)

	// 打文件
	f3, err := os.Open("abc.txt") // 只读打开
	fmt.Println(f3, err)

	/*

		O_RDONLY int = syscall.O_RDONLY // open the file read-only.
		O_WRONLY int = syscall.O_WRONLY // open the file write-only.
		O_RDWR   int = syscall.O_RDWR   // open the file read-write.
		// The remaining values may be or'ed in to control behavior.
		O_APPEND int = syscall.O_APPEND // append data to the file when writing.
		O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
		O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
		O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
		O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
	 */
	// 第一个是文件名
	//　第二个是文件打开的方式
	// 第三个是权限,当文件不存在的时候,创建成功了给予的权限

	f4, err := os.OpenFile("abc.txt",os.O_RDONLY|os.O_RDWR, os.ModePerm)
	fmt.Println(f4, err)

	// 关闭文件
	f1.Close()
	f2.Close()
	f3.Close()
	f4.Close()

	// 删除文件 remove可以删除文件和文件夹,但是文件夹要是空
	err1 := os.Remove("abc1.txt")
	if err1 != nil {
		fmt.Println("error ....")
	}else {
		fmt.Println("remove sucessful ...")
	}

	// os.RemoveAll() 相当于rm -rf, 代码删除是不经过回收站的

}
