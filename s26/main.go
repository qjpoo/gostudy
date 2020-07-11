package main

import "fmt"

func main() {
	/*
	if流程语句
	*/

	var i int
	i = 16
	if i > 10 {
		fmt.Println("i gt 10")
	}

	score := 10
	if score >= 60 {
		fmt.Println("成绩及格...")
	}else {
		fmt.Println("成绩不及格...")
	}
	fmt.Println("main end ...")
}
