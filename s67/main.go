package main

import "fmt"

func x(i [5]int) {
	fmt.Printf("函数中数组修改前的值： %v\n", i)  // [100 2 3 4 5]
	i[0] = 100
	fmt.Printf("函数中数组修改后的值： %v\n", i)  // [1 2 3 4 5]

}

func y(i []int) {
	fmt.Printf("函数中数组修改前的值： %v\n", i)  // [100 2 3 4 5]
	i[0] = 100
	fmt.Printf("函数中数组修改后的值： %v\n", i)  // [1 2 3 4 5]

}

func xy(i *int) { // 这里的i 是地址
	fmt.Printf("函数中数组修改前的值： %v\n", *i)  // [100 2 3 4 5]
	*i = 200
	fmt.Printf("函数中数组修改后的值： %v\n", *i)  // [1 2 3 4 5]

}
func main() {
	// 函数的参数传递
	// 值传递
	// 地址传递

	// 值copy 传递的是值的副本
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("数组在函数调用前： %v\n", arr) // [1 2 3 4 5]
	x(arr)
	fmt.Printf("数组在函数调用后： %v\n", arr)  // [1 2 3 4 5]


	fmt.Println("--------------------------------------")
	// 地址传递
	s := []int{1,2,3,4,5}
	fmt.Printf("数组在函数调用前： %v\n", s) // [100 2 3 4 5]
	y(s)
	fmt.Printf("数组在函数调用后： %v\n", s)  // [100 2 3 4 5]


	fmt.Println("--------------------------------------")
	// 地址传递
	xo := 100
	fmt.Printf("数组在函数调用前： %v\n", xo) // [100 2 3 4 5]
	xy(&xo)
	fmt.Printf("数组在函数调用后： %v\n", xo)  // [100 2 3 4 5]



}
