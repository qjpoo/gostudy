package main

import (
	"fmt"
)

func f1() {
	fmt.Println("我是f1函数")
}

func main() {
	// 匿名函数： 没有名字的函数
	// 定义一个匿名函数，直接进行调用，通过小括号， 通常只能使用一次
	// 匿名函数的作用：
	// 1. 将匿名函数作为另一个函数的参数，回调函数
	// 2. 将匿名函数作为另一个函数的返回值， 可以形成闭包结构

	f1()
	a := f1
	a()

	// 匿名函数调用的话，要}后面加一个（）
	func () {
		fmt.Println("我是匿名函数")
	}()

	// 把匿名函数给一个变量，就可以调用多次了
	xxoo :=	func () {
		fmt.Println("我是匿名函数")
	}
	xxoo()

	// 定义带参数的匿名函数
	func (a, b int) {
		fmt.Println(a, b)
	}(1, 2) // 把1， 2传入进来


	// 带返回值 的匿名函数
	ret1 := func (a, b int) int {
		return a + b
	}(10, 20)  // 匿名函数调用了，将执行结果给了ret1
	fmt.Println(ret1)

	ret2 := func (a, b int) int {
		return a + b
	} // 将匿名函数的值，赋值给了ret2   0x49e210
	fmt.Println(ret2)

	fmt.Println(ret2(5, 6))
}
