package main

import "fmt"

func main() {
	//数组是值类型的,传递的是副本, 相当于copy
	// 数组的类型 [size]type
	// 值传递: int, float, string, bool, array
	//引用传递: slice, map
	num := 10
	fmt.Printf("%T\n", num)
	num1 := num //值传递
	fmt.Println(num1, num)
	num1 = 20
	fmt.Println(num1, num)

	fmt.Println("=======================")
	arr1 := [3]int{0, 1, 2}
	arr2 := [3]float64{0.1, 1.2, 2.3}
	arr3 := [3]string{"A", "B", "C"}
	fmt.Printf("%T\n", arr1) // [3]int
	fmt.Printf("%T\n", arr2) // [3]float64
	fmt.Printf("%T\n", arr3) // [3]string

	arrx := arr1
	fmt.Println(arrx, arr1)
	arrx[0] = 100 //数组也是一个值传递
	fmt.Println(arrx, arr1)

	if arrx == arr1 {
		fmt.Println("true")
	}
	fmt.Println(arr1 == arrx ) //长度和类型要一致,才能比较大小
}
