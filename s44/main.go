package main

import "fmt"

func main() {

	var x int //创建变量,是没有开辟内存空间的,当被赋值了,就开辟内存空间
	x = 100
	fmt.Println(&x)

	var arr = [...]int {0, 1, 2, 3, 4}
	fmt.Println(&arr)
	fmt.Println(&arr[0])
	fmt.Println(&arr[1])
	fmt.Println(&arr[2])
	fmt.Println(&arr[3])
	fmt.Println(&arr[4])
}

