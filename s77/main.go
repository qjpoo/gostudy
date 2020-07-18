package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func oper(a, b int, fun func(x, y int) int) int {
	fmt.Printf("%T, %T, %T\n", a, b, fun)
	fmt.Println(a, b, fun) // 	10 20 0x49fa60

	return fun(a, b)  // 执行fun
}

func main() {
	// 加减乘除运算
	/* 高阶函数
	高阶函数：  根据go语言的数据类型的特点， 可以将一个函数作为另一个函数的参数
	fun1() fun2()
	将fun1函数作为fun2这个函数的参数
	fun2函数就叫高阶函数，它接收了一个函数作为参数的函数
	fun1函数： 回调函数  回头调用函数
	作为另一个函数的参数的函数叫做回调函数
	*/
	fmt.Println(add(1, 2))

	// add
	sy := oper(10, 20, add)
	fmt.Println(sy)

	// sub
	sk := oper(1, 2, sub)
	fmt.Println(sk)


	fm := func(i, j int) int{
		return i + j
	}

	ret1 := oper(10, 20, fm)
	fmt.Println(ret1)

	// 使用匿名函数
	ret2 := oper(10, 20, func(x, y int)int{ // 这里只是定义了匿名的函数,而执行是在oper函数里面在执行
		return  x + y
	})
	fmt.Println(ret2)
}
