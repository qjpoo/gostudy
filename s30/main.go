package main

import "fmt"

func main() {
	/*
		switch
		switch可以作用在其他的类型上,不像if只能作用在bool类型上
		case后面的类型要跟num的类型要是一样的
		case是无序的,谁先匹配,谁执行
		case后面的数值是唯一的,不能有重复的,比如case 3,写成二个
		default可有可无
	*/

	num := 5
	switch num {
	case 1:
		fmt.Println("这是第一季度...")
	case 2:
		fmt.Println("这是第二季度...")
	case 3:
		fmt.Println("这是第三季度...")
	case 4:
		fmt.Println("这是第四季度...")
	default:
		fmt.Println("输入错误")
	}

	//计算器
	for {
		num1 := 0
		num2 := 0
		oper := ""
		fmt.Println("input int num1: ")
		fmt.Scanln(&num1)
		fmt.Println("input int num2: ")
		fmt.Scanln(&num2)
		fmt.Println("input oper: '+ - * /'")
		fmt.Scanln(&oper)

		switch oper {
		case "+":
			fmt.Printf("%d = %d + %d\n", num1+num2, num1, num2)
		case "-":
			fmt.Printf("%d = %d - %d\n", num1-num2, num1, num2)
		case "*":
			fmt.Printf("%d = %d * %d\n", num1*num2, num1, num2)
		case "/":
			fmt.Printf("%d = %d / %d\n", num1/num2, num1, num2)
		default:
			fmt.Println("error ...")
		}
	}
}
