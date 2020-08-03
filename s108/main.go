package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func CopyFile1(srcFile, dstFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(dstFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		return 0, nil
	}
	defer file1.Close()
	defer file2.Close()

	// 读写
	bs := make([]byte, 1024, 1024)
	n := -1
	total := 0
	for {
		n, err = file1.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("copy file finish ...")
			break
		} else if err != nil {
			fmt.Println("报错了...")
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil

}

// 推荐的方法
func CopyFile2(srcFile, dstFile string) (int64, error)  {
	file1, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	file2, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
	return io.Copy(file2, file1)
}

// ioutil ioutil.ReadFile()和ioutil.WriteFile()是一次的读写,如果文件太大,会有内存的溢出
func CopyFile3(srcFile, dstFile string) (int, error)  {
	bs, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return 0, err
	}
	err = ioutil.WriteFile(dstFile, bs, os.ModePerm)
	if err != nil {
		return 0, err
	}
	return len(bs), nil

}




func main() {
	// 复制文件

	srcFile := "src.txt"
	dstFile := "dst.txt"
	//total, err := CopyFile1(srcFile, dstFile)
	//fmt.Println(total, err)

	// 推荐这一种方法
	//total, err := CopyFile2(srcFile, dstFile)
	//fmt.Println(total, err)

	total, err := CopyFile3(srcFile, dstFile)
	fmt.Println(total, err)




/*	pwd, _ := os.Getwd()
	fileInfo, _ := ioutil.ReadDir(pwd)
	for _, v := range fileInfo {
		fmt.Println(v.Name())
	}
*/

}
