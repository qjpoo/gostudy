package main

import "fmt"

func main() {
	/*
	if语句中的变量作用域
	*/

	if num := 5; num >0 {
		fmt.Println("num是大于0的 ...")
		fmt.Println(num)
	}
	//fmt.Println(num) //num只在If条件语句中才会生效

	num2 := 100
	if num2 >0 {
		fmt.Println("num2 gt 0")
	}
	fmt.Println(num2)

	//num只在if中生效,num2从17行开始到结尾都是生效的,为什么,因为num2定义完成之后有一个回车换行
}
