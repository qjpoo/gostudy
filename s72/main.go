package main

import "fmt"

func f1(x int) int {

	if x == 1 {
		return 1
	}
	sum := f1(x-1) + x
	return  sum
}

func f2( x int) int {

	if x == 1 || x == 2 {
		return 1
	}
	yo := f2(x - 2) + f2(x -1)
	return  yo

}
func main() {

	// 递归函数，自己调用自己
	// 递归函数要有一个出口
	// 求1-5的和

	fmt.Println(f1(5))

	/*
	1 2 3 4 5 6  7  8   9    10
	1 1 2 3 5 8 13  21  34   55
	*/
	fmt.Println(f2(10))
}
