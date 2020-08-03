package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// io效率其实不低,低的是我们频繁的去访问磁盘
	// bufio包, 高效的io读写
	// buffer 缓存, io: input/ output
	// bufio包是对Io的读书在一次的封装
	/*
		将io包下的reader, writer对象进行包装,提高读写的效率
		func (b *Writer) Write(p []byt) (nn int, err err)

		func (b *Writer) WriteByte(c byte) err

		func (b *Write) WriteRuner(r rune) (size int, err error)

		func (b *Writer) WriteString(s string) (int, error)
	*/

	/*
		f, _ := os.Open("123.txt")
		reader := bufio.NewReader(f)
		fmt.Println(reader.Size())
		for {
			str, err := reader.ReadString('\n')
			fmt.Print(str)
			if err == io.EOF{
				fmt.Println("file read finish ...")
				break
			}
		}
	*/

	/*
	//	 os包中的file对象(结构体)它实现了io.read和io.write方法
	fileName := "123.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	 */

	// 创建一个Reader对象
	// 高效读取
	/*
		b1 := bufio.NewReader(file) // 返回的是一个read对象
		p := make([]byte, 1024, 1024)
		n1, err := b1.Read(p) // 把从文件中读取的内容放到byte切片里面
		fmt.Println(n1)
		fmt.Println(err)
		//fmt.Println(p)
		fmt.Println(string(p))
	*/
	/*
	   	// ReadLine() 读一行,比较低效
	   	fmt.Println("--------------------------")
	   	b1 := bufio.NewReader(file) // 返回的是一个read对象
	   	data, flag, err := b1.ReadLine()
	   	fmt.Println(string(data))
	   	fmt.Println(flag)
	   	fmt.Println(err)

	   	data, flag, err = b1.ReadLine()
	   	fmt.Println(string(data))
	   	fmt.Println(flag)
	   	fmt.Println(err)
	   }
	*/

	// ReadString
/*	b1 := bufio.NewReader(file) // 返回的是一个read对象
	s1, err := b1.ReadString('\n')
	fmt.Println(err)
	fmt.Println(s1)
*/

/*
	// ReadString
	b1 := bufio.NewReader(file) // 返回的是一个read对象
	for {
		s1, err :=  b1.ReadString('\n')
		if err == io.EOF {
			fmt.Println("read finish ...")
			break
		}
		fmt.Print(s1)
	}

 */

/*
	// ReadBytes
	b1 := bufio.NewReader(file)
	data, err := b1.ReadBytes('\n')
	fmt.Println(string(data))
	fmt.Println(data)
*/

	/*
	// Scanner 读取输入
	// 如果输入 hello world,只会显示hello,它是以空格
	s2 := ""
	fmt.Scanln(&s2)
	fmt.Println(s2)
	 */

	/*
	r := bufio.NewReader(os.Stdin)
	d, err := r.ReadString('\n')
	fmt.Println(d, err)
	 */


	// 写的操作

	fmt.Println(os.Getwd())
	writeFile := "cc.txt"
	file, err := os.OpenFile(writeFile, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := bufio.NewWriter(file)
	n, err := w.WriteString("heloo, world ..")
	if err != nil {
		fmt.Println(err)
	}
	w.Flush()  // 注意要flush一下,如果写的数据比buff大的话,可以不用flush,因为写的数据大于Buffer了,所以会直接写到文件 里面
	fmt.Println(n)

	// 写到标准输出上
	w1 := bufio.NewWriter(os.Stdout)
	n2, _ := w1.WriteString("he, he, he")
	fmt.Println(n2)
	w1.Flush()

}