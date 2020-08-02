package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// io的写


	// 1. 打开文件

	// 2. 写出数据

	// 3. 关闭文件

	// file, err := os.Open(filenName)
	filenName := "aa.txt"
	file, err := os.OpenFile(filenName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	//bs := []byte{'A', 'B', 'C', 'D', 'E', 'F'}
	//n, err := file.Write(bs)
	n, err := file.Write([]byte("hello, xlllfkdjf\n"))
	fmt.Println(n)
	HandleError(err)

	writeFile, err := os.OpenFile("writefile.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	HandleError(err)
	_, err = writeFile.WriteString("hello, world ...\r\n")
	HandleError(err)


}

func HandleError(err error)  {
	if err != nil {
		log.Fatal(err)
	}

}