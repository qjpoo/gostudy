package main

import "fmt"

func main() {
	/*
	变量的定义,它是一种静态语言,即强类型语言
	*/

	// 1.第一种写法
	var num1 int
	num1=10
	fmt.Printf("num1的数值是: %d\n", num1)


	//第二种: 类型推断
	var name="chiling"
	fmt.Printf("name的类型是: %T  name的值是: %s\n", name, name)

	//简短的申明
	sum := 100
	fmt.Printf("sum的类型是: %T  sum的值是: %d\n", sum, sum)


	//多变量申明
	var a, b, c int = 100, 200, 300 //相同的类型
	fmt.Println(a, b, c)

	var x, y, z  = 100, "hello", 300 //不同的类型
	fmt.Println(x, y, z)

	// 集合类型
	var (
		name10 = "xxoo1"
		name20 = 20
		name30 = "xxoo3"
	)
	fmt.Println(name10, name20, name30)
}
