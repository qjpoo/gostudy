package main

import "fmt"

func main() {
	// for的嵌套

	//打印第一行5个星星,打印五行
	for i := 1; i <= 5; i++ {
		for j := 1; j <=5; j++ {
			fmt.Printf("*\t")
		}
		fmt.Println()
	}


	//乘法口诀
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d = %d * %d\t", i * j, j, i)
		}
		fmt.Println()
	}
}
