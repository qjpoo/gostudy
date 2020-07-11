package main

import "fmt"

func main() {
	/*
	break  跳出循环
	continue 结束本轮
	*/

	// 1, 2
	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}else {
			fmt.Println(i)
		}
	}
	fmt.Println()

	// 1, 2, 4, 5
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}else {
			fmt.Println(i)
		}
	}
}
