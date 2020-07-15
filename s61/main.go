package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		strconv  字符串和基本类型间的转换
		string convert
	*/

	//fmt.Println("aa" + 100)

	s1 := true
	fmt.Printf("%T, %v", s1, s1)

	s2 := "true"
	fmt.Printf("%T, %v", s2, s2)

	fmt.Println("-----------------------------------")
	// string ==> bool   strconv.ParseBool
	s3 := "true"
	b1, err := strconv.ParseBool(s3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T, %t\n", b1, b1)

	// bool --> string      FormatBool只有一个返回值
	ss1 := strconv.FormatBool(b1)
	fmt.Printf("%T, %s\n", ss1, ss1)

	// string -----> int strcon.ParseInt
	i := "100"
	//in, err := strconv.ParseInt(i, 10, 64)  // 10进制, 64位  输出 100
	in, err := strconv.ParseInt(i, 8, 64)  // 8进制, 64位  输出 64
	//in, err := strconv.ParseInt(i, 2, 64)  // 2进制, 64位  输出 4

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T, %d", in, in)

	// int --> string
	x := int64(12)
	x1 := strconv.FormatInt(x,10)
	fmt.Printf("%T, %s\n", x1, x1)

	// strconv.Itoa()函数的参数是一个整型数字，它可以将数字转换成对应的字符串类型的数字
	xx1 := 10
	xx2 := strconv.Itoa(xx1)
	fmt.Printf("%T, %s\n", xx2, xx2)  // 10

	// string函数的参数若是一个整型数字，它将该整型数字转换成ASCII码值等于该整形数字的字符。
	xi := 100
	xii := string(xi)
	fmt.Printf("%T, %s\n", xii, xii) // 100对就的ascii码是100

	// Atoi
	xa1 := "10"
	xa2, err := strconv.Atoi(xa1)
	fmt.Printf("%T, %d\n", xa2, xa2)

}
