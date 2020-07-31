package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error)  {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n >0 {
		return p[:n], nil
	}
	return p, err

}
func main() {
	// io包

	// 从标准输入读取
	//data, err := ReadFrom(os.Stdin, 11)
	//fmt.Println(string(data), err)

	// 从普通文件读取
	f, err := os.Open("./123.txt")
	fmt.Println(f, err)
	data, err := ReadFrom(f, 20)
	fmt.Println(string(data), err)

	// 从字符串读取
	data, err = ReadFrom(strings.NewReader("from hello world "),12)
	fmt.Println(string(data), err)

	// 下面的例子简单的实现将文件中的数据全部读取
/*	fmt.Println("-----------------------")
	fmt.Println(os.Getwd())
	file, err := os.Open("./123.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewReader(os.Stdout)
	by, err := writer.ReadByte()
	fmt.Println(string(by), err)
*/

	fmt.Println("-------------------------")
	stat, err := os.Stat("123.txt")
	//fmt.Println(stat, err)
	p := make([]byte, stat.Size())
	file, _ := os.Open("123.txt")
	file.Read(p)
	log.Println(string(p))

	//  写文件
	fmt.Println("-----------写文件--------------")
	userFile := "test.txt"
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i :=0;i< 10;i++ {
		fout.WriteString("just a test \r\n")
		fout.Write([]byte("hello, world ... \r\n"))
	}

	// 读文件
	fmt.Println("-----------读文件--------------")
	readFile := "test.txt"
	f1, err := os.Open(readFile)
	defer f1.Close()
	if err != nil {
		log.Println(err)
		return
	}
	buf := make([]byte, 1024)
	for {
		n, _ := f1.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}


	var xxoo []int = []int{1,2,3}
	fmt.Println(xxoo)
	xxoo1 := []int{1,2,3}
	fmt.Println(xxoo1)

}
