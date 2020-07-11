package main

import "fmt"

func main() {
	/*
		数组的遍历
	*/


	// 数组的遍历
	var x = [6]int{0, 1, 2, 3, 4, 5}
	for i := 0; i < len(x); i++ {
		x[0] = 100
		fmt.Println(x[i])
	}


	// range
	var str  = [...]string{"A", "B", "C", "D", "E"}
	for k, v := range str {
		fmt.Printf("数组的下标为: %d, 数组的值为: %s\n", k, v)
	}
	for _, v := range str {
		fmt.Printf("数组的值为: %s\n",  v)
	}
}
