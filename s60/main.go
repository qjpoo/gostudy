package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
	strings包 对字符串的操作
	*/

	s1 := "hello, world ..."
	// 包括指定的字符串
	fmt.Println(strings.Contains(s1, "h"))
	fmt.Println(strings.Contains(s1, "ko"))
	// 包括任意字符串
	fmt.Println(strings.ContainsAny(s1, "abce"))
	// 子串的个数
	fmt.Println(strings.Count(s1, "o"))

	prefix := "/root/log/2020/nginx.log"
	// 指定字符串开头
	if strings.HasPrefix(prefix, "/") {
		fmt.Println("prefix string is  prefix / ")
	}

	// 指定结尾
	if strings.HasSuffix(prefix, ".log") {
		fmt.Println(" prefix string is suffix .log")
	}

	// 子串第一次出现的次数，没有为-1
	fmt.Println(strings.Index(prefix, "A"))

	// 任意字符出现在prefix中的位置
	fmt.Println(strings.IndexAny(prefix, "/bi."))

	fmt.Println(strings.IndexByte(prefix, 'l'))

	fmt.Println(strings.LastIndex(prefix, "l"))


	// 拼接
	s2 := []string{"hello", "world"}
	fmt.Println(strings.Join(s2, " + "))

	// 切割
	s3 := "hello,world,china"
	s4 := strings.Split(s3, ",")
	for i :=0 ; i < len(s4);i++{
		fmt.Printf("%s ", s4[i])
	}

	// 重复打印某个字符串
	fmt.Println(strings.Repeat("hello", 5))


	s5 := "hello, hello, world"
	// -1 替换所有的
	fmt.Println(strings.Replace(s5,"hello","HELLO", 2))

	fmt.Println(strings.ToUpper(s5))
	fmt.Println(strings.ToLower(s5))

	// 截取子串substring
	// str[start:end] 前闭后开
	l := "abcdefg"
	fmt.Println(l[0:6])







}
