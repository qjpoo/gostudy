package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 断点续传
	/*
		文件的传递: 文件的复制
	*/

	srcFile := "C:\\Users\\瞿健\\go\\src\\leg.jpg"
	// abc := path.Ext(srcFile) // .jpg
	// abc := filepath.Base(srcFile) // leg.jpg
	//abc := strings.LastIndex(srcFile, "\\")
	//fmt.Println(string(srcFile[abc+1:]))
	// 取得文件名
	dstFile := srcFile[strings.LastIndex(srcFile, "\\")+1:]
	//fmt.Println(dstFile)
	tmpFile := dstFile + "tmp.txt"
	//fmt.Println(tmpFile)

	file1, err := os.Open(srcFile)
	HandleErr(err)
	file2, err := os.OpenFile(dstFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	HandleErr(err)
	file3, err := os.OpenFile(tmpFile, os.O_CREATE|os.O_RDWR, 0777)
	HandleErr(err)

	defer file1.Close()
	defer file2.Close()

	// 先读取临时文件中的数据,在seek
	file3.Seek(0,io.SeekStart)
	bs := make([]byte,100, 100)
	n1, err := file3.Read(bs)
	//HandleErr(err)
	// 把byte转成string
	countStr := string(bs[:n1])
	// 把string转为十进制
	count, err := strconv.ParseInt(countStr,10,64)
	// HandleErr(err)
	fmt.Println(count)

	// 设置读,写的位置
	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)
	data := make([]byte, 1024,1024)
	n2 := -1 // 读取的数据量
	n3 := -1 // 写出的数据量
	total := int(count)

	// 复制文件
	for {
		n2, err = file1.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("copy file finish ...", total)
			file3.Close()
			os.Remove(tmpFile)
			break
		}
		n3, err = file2.Write(data[:n2])
		total += n3

		// 将复制的总量, 存储到临时的文件中去
		file3.Seek(0,io.SeekStart)
		file3.WriteString(strconv.Itoa(total)) // 把int 转化为字符串
		fmt.Println("total: ", total)


		// 模拟断电
		if total > 10240 {
			panic("断电了 ...")
		}
	}


}

func HandleErr(err error)  {
	if err !=nil {
		//log.Fatal(err)
		fmt.Println(err)
	}
}
