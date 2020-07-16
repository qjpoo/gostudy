package main

import "fmt"

func x(a, b int) (m, n int) {
	m = a * b
	n = 2 * ( a + b)
	return m, n
}
func main() {
	// 函数的多返回值


	// 计算一个长方形的面积和周长
	v1, v2 := x(1, 2)
	fmt.Printf("面积为：%d, 周长为： %d", v1, v2)
}
