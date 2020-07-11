package main

import "fmt"

func main() {
	/*
	if else
	if 和 { 要在同一行, } 和else 也要在同一行
	*/

	score := 10
	if  score >=60{
		fmt.Println("成绩及格")
	}else {
		fmt.Println("成绩不及格")
	}

	str := "男"
	if str == "男" {
		fmt.Println("上男侧所...")
	}else {
		fmt.Println("上女侧所...")
	}
	fmt.Println("main function end ...")
}
