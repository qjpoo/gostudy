package main

import "fmt"

func main() {
	/* 数组指针和指针数组
	数组指针: 首先是一个指针, 一个数组的地址
	*[4]Type *[4]int
	指针数组: 首先是一个数组, 存储的数据类型是指针
	[4]*Type   [4]*int

	*[4]float64   指针, 一个存储了5个浮点类型的数组的指针
	*[3]string

	[3]*string  数组, 存储了3个字符串的指针数组
	[5]*float64

	*[5]*float64 指针, 一个数组的指针,存储了5个float类型的数据的指针地址的数组的指针
	*[3]*string

	**[4]string 指针, 存储了4个字符串数据的数组的指针的指针
	**[4]*string



	*/

	// 创建一个普通的数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Printf("%T, %v\n", arr1, arr1)

	// 创建一个指针, 存储该数组的地址  数组指针
	var p1 *[4]int
	p1 = &arr1
	fmt.Printf("%T, %v\n", p1, p1) // *[4]int, &[1 2 3 4]
	fmt.Printf("%p\n", p1)         // 0xc00000e440
	fmt.Printf("%p\n", &p1)        //0xc000006030

	// 操作数组
	fmt.Println(*p1) // [1 2 3 4]
	(*p1)[0] = 100
	fmt.Println(arr1[0]) // 100

	p1[0] = 200          // p1[0]  就是等于 (*p1)[0] p1[0]是简写
	fmt.Println(arr1[0]) // 200

	// 指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}
	arr3 := [4]*int{&a, &b, &c, &d}
	fmt.Println(arr2)  // [1 2 3 4]
	fmt.Println(arr3)  // [0xc000010110 0xc000010118 0xc000010120 0xc000010128]
	arr2[0] = 100
	fmt.Println(arr2) // [100 2 3 4]
	fmt.Println(a) // 1

	*arr3[0] = 200
	fmt.Println(a) // 200

	fmt.Println("-------------------")
	b = 1000
	fmt.Println(arr2)  // [100 2 3 4]
	fmt.Println(arr3) // [0xc000010110 0xc000010118 0xc000010120 0xc000010128]

	// 200
	//1000
	//3
	//4
	for i:=0;i<len(arr3);i++{
		fmt.Println(*arr3[i])
	}

}
