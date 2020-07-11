package main

import "fmt"

func main() {
	/*
		goto
	*/

	/*	x := 10
		LOOP:
		for x <= 20 {
			if x == 15 {
				x++
				goto LOOP
			}
			fmt.Println(x)
			x++
		}
	*/
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if j == 2 {
				goto breakHere
			}
		}
	}
	return //手动返回,避免执行进入标签
breakHere:
	fmt.Println("end...")
}
