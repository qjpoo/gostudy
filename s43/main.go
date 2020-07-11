package main

import "fmt"

func main() {
	/*
	基本类型: 整数, 浮点 , 布尔, 字符串
	复合类型: array, slice , map, struct, pointer, function, channel

	数组: 存储一组相同数据类型的的数据结构,一旦定义后,大小不能更改

	*/
	//数组的定义
	var arr1 [4]int
	//数组的访问
	arr1[0] = 0
	arr1[1] = 1
	arr1[2] = 2
	arr1[3] = 3
	fmt.Println(arr1[0])
	fmt.Println(arr1[3])

	fmt.Println(len(arr1)) //实际存储的数据量
	fmt.Println(cap(arr1)) //容器中能够存储的最大的数量
	//因为数组是定长的,所以len == cap

	arr1[0] = 100
	fmt.Println(arr1[0])


	//数组的其它创建方式
	//var a [4]int // var a=[4]int
	var b = [4]float64 {1.0, 2.0, 3.14}
	fmt.Println(b)

	var c = [5]int{1:100, 0:200}
	fmt.Println(c)

	var d = [5]string{"hello","A", "B"}
	fmt.Println(d)

	// 自动根据元素的个数来推出数组的大小
	var e = [...]string{"A", "B", "C"}
	fmt.Println(e)
	fmt.Println(len(e))

	var f = [...]string{0:"A", 7:"B"}
	fmt.Println(f)
	fmt.Println(len(f))

}
