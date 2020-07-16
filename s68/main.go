package main

import "fmt"

func x(a int) (sum int) { // 指明了要返回的值是sum
	sum += a
	return sum
}
func main() {
	// 函数返回值: 一个函数定义上有返回值 ，那么函数中必须要使用return语句，将结果返回

	fmt.Print(x(100))

}
