package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	s, t := "hello go", "helloo Go"
	isEqual := strings.EqualFold(s, t)
	fmt.Println(isEqual)

	//prefix := "/root/123.txt"
	prefix := "Abc.txt"
	//isTrue := strings.HasPrefix(prefix, "/")
	isTrue := strings.HasPrefix(prefix, "a")
	fmt.Println(isTrue)

	suffix := "/root/123.txt"
	isTrue1 := strings.HasSuffix(suffix, ".txt")
	fmt.Println(isTrue1)

	subStr := "txt"
	isTrue2 := strings.Contains(suffix, subStr)
	fmt.Println(isTrue2)

	r := rune(101)
	fmt.Println(string(r))
	conRun := strings.ContainsRune(s, r)
	fmt.Println(conRun)

	// 子串sep在字符串s中第一次出现的位置，不存在就返回-1
	sep := "0"
	sepIndex := strings.Index(s, sep)
	fmt.Println(sepIndex)

	// 子串sepLast在字符串中最后一次出现的位置， 不存在则返回-1
	sepLast := "o"
	sepLastPo := strings.LastIndex(s, sepLast)
	fmt.Println(sepLastPo)

	// title格式，首字母都是大写
	s1 := "hEllo, quinn"
	t1 := strings.Title(s1)
	fmt.Println(t1) // Hello, Quinn

	// 都变成大写
	t2 := strings.ToTitle(s1)
	fmt.Println(t2) // HELLO, QUINN

	// 大写变成小写
	t4 := strings.ToLower(s1)
	fmt.Println(t4)

	// 重复生成字符串
	t5 := strings.Repeat("xxoo ", 5)
	fmt.Println(t5)

	// // 返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串
	s2 := "go, go, go, to, study"
	s_old, s_new := ",", "#"
	t6 := strings.Replace(s2, s_old, s_new, -1)
	fmt.Println(t6)

	// 去掉字符串前后指定的字符串
	sn, cut := " hello, world   ", " "
	t7 := strings.Trim(sn, cut)
	fmt.Println(t7)

	t8 := strings.TrimLeft("#/bin/bash", "#")
	fmt.Println(t8)

	t9 := strings.TrimPrefix("/root/scriptes/hello.sh", "/root")
	fmt.Println(t9)

	t10 := strings.TrimSpace("  hello, world  ")
	fmt.Println(t10 + "S")

	// 返回将字符串按照空白（unicode.IsSpace确定，可以是一到多个连续的空白字符）分割的多个字符串
	strSplit := strings.Fields("hello world!     go language ...")
	for k, v := range strSplit {
		fmt.Println(k, "====>", v)
	}

	fmt.Println("==========================")
	ssplit := strings.Split("//root/scripts/20200731/mysql.sh", "/")
	for k, v := range ssplit {
		fmt.Println(k, "====>", v)
	}

	//  // 将一系列字符串连接为一个字符串，之间用sep来分隔
	j := strings.Join([]string{"a", "b", "c", "d"}, "@")
	fmt.Println(j)

	// 将s的每一个unicode码值r都替换为mapping(r)，返回这些新码值组成的字符串拷贝。如果mapping返回一个负值，将会丢弃该码值而不会被替换
	map_func := func(r rune) rune {
		switch {
		case r > 'A' && r < 'Z':
			return r + 32
		case r > 'a' && r < 'z':
			return r - 32
		}
		return r
	}
	s = "Hello World!"
	s_map := strings.Map(map_func, s)
	fmt.Println(s_map) // hELLO wORLD!

	// Reader 类型从一个字符串读取数据，实现了io.Reader, io.Seeker等接口。
	newR := strings.NewReader("123456789")
	fmt.Println(reflect.TypeOf(newR), newR)
	fmt.Println(newR.Len())  // 获取未读取的长度
	fmt.Println(newR.Size()) //  获取字符串的长度

	for newR.Len() > 5 {
		b, err := newR.ReadByte() // 读取1 byte
		fmt.Println(string(b), err, newR.Len(), newR.Size())
		/*
		1 <nil> 8 9
		2 <nil> 7 9
		3 <nil> 6 9
		4 <nil> 5 9
		*/
	}


	b_s :=make([]byte, 5)
	n, err := newR.Read(b_s)
	fmt.Println(string(b_s), n, err)  // 56789 5 <nil>
	fmt.Println(newR.Size()) // 9
	fmt.Println(newR.Len())  // 0


	si := "go hello"
	ra := strings.NewReplacer("g", "G", "h", "j")
	fmt.Println(reflect.TypeOf(ra))
	fmt.Println(ra.Replace(si))
}
