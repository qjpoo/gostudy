package main

import "fmt"

// 求两个整数的和
func sub(x int, y int) int {
	//func sub(x, y int) int { //可以简写一起
	return x + y
}

func main() {
	// 函数参数使用
	/*
	形式参数：定义函数时，用于接收外部传入的数据
	实际参数： 调用函数时，传给形参的实际数据

	函数的调用：
	A: 函数名必须匹配
	B: 实参与形参必须要一一对应： 顺序， 个数， 类型
	*/

	fmt.Println(sub(10, 20))


}
