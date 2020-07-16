package main

import "fmt"

//全局变量, 只有一份,修改了就都修改了.
var name = "quinn"
var k = 500
func f1() {
	name = "xxoo1"
	k = 600
	fmt.Println(name, k)
}

func f2() {
	fmt.Println(name, k)
}
func main() {
	// 函数的作用域
	// 局部 和全局
	// 确定变量的作用域的方法,就是看定义变量的左花括号
	// 局部变量超过了它的范围,就销毁了


	// n 是局部,只能在Main中使用
	n := 100
	fmt.Println(n)

	if i:=1; i<2 {
		n := 200
		fmt.Println(i) // 1
		fmt.Println(n) // 200
	}
	fmt.Println(n) // 100

	// m只能在If语句块中被访问
	if m:=1; m <2 {
		fmt.Println(m)
	}
	// fmt.Println(m) m是不能被访问的

	fmt.Println(name, k) // quinn 500
	f1() // xxoo1 600
	f2()  //xxoo1 600
	fmt.Println(name, k)  // xxoo1 600

}
