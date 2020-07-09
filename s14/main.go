package main

import "fmt"

func main() {
	var num1 int
	num1 = 100
	fmt.Printf("num1的类型为: %T, num1的值为: %d, num1的地址为: %p\n", num1, num1, &num1)

	num1 = 200
	fmt.Printf("num1的类型为: %T, num1的值为: %d, num1的地址为: %p\n", num1, num1, &num1)
}
