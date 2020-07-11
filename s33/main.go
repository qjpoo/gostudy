package main

import "fmt"

func main() {
	/*
	for 表达式1; 表达式2; 表达式3 {
		循环体
	}
	一般表达1只执行一次,就是初始化的时候
	表达式2条件判断
	执行循环体里面的内容
	在执行表达式3
	在执行表达2的内容
	在执行循环体
	*/

	for i := 0; i < 5; i++ {
		fmt.Println(i, "hello, world ...")
	}
	// fmt.Println(i)  //i是没有定义的
}
