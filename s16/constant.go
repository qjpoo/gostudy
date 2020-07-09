package main

import "fmt"

func main() {
	/*
		常量的定义
	*/

	fmt.Println(100)
	fmt.Println("hello")

	const PATH string = "https://www.baidu.com"
	const PI = 3.14
	fmt.Println(PATH)

	// 常量是不能修改的

	const C1, C2, C3 = 100, 3.14, "hihi"

	const (
		NUM1 = 0
		NUM2 = 10
		NUM3 = 2
	)
	fmt.Println(NUM1, NUM2, NUM3)

	//常量如果没有定义值的话,就是上一行变量的值
	const (
		NUM4 = 200
		NUM5
		a string = "hi"
		b
		c
	)
	fmt.Println(NUM4, NUM5, a, b, c)

	// 枚举类型,一组相关的数据
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)
}
