package main

import "fmt"

func main() {
	// string
	/*
	string "ABC" 用双引号   ``
	字符串是一些字节的集合
	go中的字符串是一个字节的切片,可以将字节用封装在""来创建字符串,GO中的字符吕是unicode兼容的,并且是utf-8的编码
	字符操作的都是对应的编码值
	字符串就是字符的序列(有下标和索引,每个字节都有固定的位置)
	A   65
	B   66
	a   97
	这些就是字节, byte  utf8

	*/

	s1 := "hello, world ..."
	for k, v := range s1 {
		fmt.Println(k, v)
		fmt.Printf("%c\n",v)
	}

	s2 := "中国"
	s3 := "hello"
	fmt.Println(len(s2)) //每个中文占三个字节
	fmt.Println(len(s3)) //每个英文字符占有一个字节
	fmt.Println(s3[0])  // 104

	s4 := 104
	s5 := 'h'
	fmt.Printf("%c, %c, %c\n", s4, s5, s3[0]) // h, h, h

	// 字符串的遍历
	s6 := "hello, world ..."
	for i:=0;i<len(s6); i++ {
		fmt.Printf("%c ", s6[i]) // 如果有中文的话就是乱码了
	}

	fmt.Println()

	// for range
	s7 := "hello, world"
	for k, v := range s7 {
		fmt.Printf("%d, %c\n", k, v )
	}

	// 字符串是字节的集合
	//s8 := []byte{'A', 'B', 'C'}
	s8 := []byte{100,101,102}
	sa1 := string(s8) //根据字节,转化了字符串
	fmt.Println(sa1)

	sa2 := "helloworld"
	sa3 := []byte(sa2)  //字符串转化了字节
	fmt.Println(sa3)

	// 字符串不能修改
	//sa4 := "helloo"
	//sa4[0] = "K"

	// 在 Go 语言中，字符串的内容是不能修改的，也就是说，你不能用 s[0] 这种方式修改字符串中的 UTF-8 编码，
	//如果你一定要修改，那么你可以将字符串的内容复制到一个可写的缓冲区中，然后再进行修改。这样的缓冲区一般是 []byte 或 []rune。
	//如果要对字符串中的字节进行修改，则转换为 []byte 格式，如果要对字符串中的字符进行修改，则转换为 []rune 格式，转换过程会自动复制数据。
	sn := "hello 中国"
	so := []byte(sn)
	so[0] = 'H'
	fmt.Println(sn)
	fmt.Printf("%c", so[0])

	sy := []rune(sn)
	sy[6] = '美'
	fmt.Println(sn)
	fmt.Printf("%c", sy[6])

}
