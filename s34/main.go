package main

import "fmt"

func main() {
	/*
	for 表达式2 {} ==> while
	*/
	//for i := 0; i < 5; i++ {}
	/*
	一般省表达1,把变量定义在外面
	一般省略表达2  条件永远为真,死循环
	省略表达3的话,一般写在for循环体里面

	*/
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println(i) //for 主要就是i的作用域

	// for ;; {} ==> for {}  条件永远为真  for true , while true



}
