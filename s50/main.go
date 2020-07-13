package main

import "fmt"

func main() {
	/*
	slice是一个引用类型
	*/

	s1 := []int{1, 2, 3}
	fmt.Println(s1)
	fmt.Printf("类型为：%T, 长度为：%d, 容量为：%d\n", s1, len(s1), cap(s1))
	fmt.Printf("%p\n", s1) // s1是引用类型的，不用加& , 它本来存储的就是一个地址 xc00000e380
	fmt.Printf("%p\n", &s1) //如果加了&，表示获取的是本身s1这个地址 0xc0000044a0

	s1 = append(s1, 4, 5, 6)
	fmt.Printf("%p\n", s1) // 0xc00000a2a0


	var x int
	x = 100
	fmt.Printf("%p\n", &x)





}
