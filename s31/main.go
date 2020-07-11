package main

import "fmt"

func main() {
	/*
	switch
	switch如果省略了变量,那就是作用在true
	*/

	switch {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}

/*
case后面可以根多个数值
switch 变量 {
case 数值1,数值2, ... 数值n:
case 数值2:
}
*/
	month := 5
	days := 0
	year := 2019
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
		fmt.Println(days)
	case 4, 6, 9, 11:
		days = 30
		fmt.Println(days)
	case 2:
		if year % 400 == 0 || year % 4  == 0  && year % 100 != 0 {
			days = 29
			fmt.Println(days)
		}else {
			days = 28
			fmt.Println(days)
		}
	}

	/*
	可以把变量定义在switch中
	switch 定义变量;变量名
	*/
	switch lang := "golang";lang {
	case "golang":
		fmt.Println("golang")
	case "c":
		fmt.Println("c")
	case "java":
		fmt.Println("java")
	}
	//fmt.Println(lang) //这里就不能获得这个lang的变量的值


}
