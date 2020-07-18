package main

import "fmt"

func f1(a, b int) {
	fmt.Println(a, b)

}

func f2(a, b int) int {
	return a + b

}
func main() {

	// 函数的本质
	/*
		f1和f1()的区别
		f1返回的是一个以f1为名称的内存地址： 0x12313132
		f1()是返回函数执行结果，函数调用
	*/
	//m := make(map[int]map[string]int)
	fmt.Printf("%T\n", f1)
	fmt.Println(f1) // 0x49fa60
	f1(1, 2)

	// 函数类型的变量
	a := f1 // 将f1的地址给了a,这里不能写成 f1()
	fmt.Printf("%T\n", a)

	var b func(int, int)
	b = f1 // 将f1的地址给了b
	fmt.Printf("%T\n", b)
	b(1, 2)

	ret1 := f2 //将f2的内存地址赋值给了ret1
	ret2 := f2(1, 2) // 将f2函数执行结果return后的结果，给了ret2
	fmt.Printf("%T, %v\n", ret1, ret1) // 0x49fb30
	fmt.Printf("%T, %v\n", ret2, ret2) // 3
	fmt.Println(ret1(10, 20))
}
