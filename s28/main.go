package main

import "fmt"

func main() {
	/*
		if 嵌套
		if 条件1 {
			代码段1
		}else {
			if 条件2 {
				代码段2
			}else {
				代码段3
			}
		}

		if 条件1 {
				代码段1
		}else if 条件2 {
				代码段2
		}else {
				代码段3
		}
	*/

	sex := "n"
	if sex == "男" {
		fmt.Println("请上男侧所...")
	}else {
		if sex == "女" {
			fmt.Println("请上女侧所...")
		}else {
			fmt.Println("没有侧所可以上...")
		}
	}

	score := 69
	if score >=90 {
		fmt.Println("优秀")

	}else if  score >79 && score <=89 {
		fmt.Println("良好")

	}else if score <=79 && score >=60 {
		fmt.Println("合格")
	}else {
		fmt.Println("不合格")
	}


}
