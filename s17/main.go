package main

import "fmt"

func main() {
	/*
	数据的基本类型
	布尔:
	数值:
	字符串： string
	*/
	//bool
	var b1 bool
	b1 = true
	fmt.Printf("%T, %t\n", b1, b1)
	var b2 bool
	b2 = true
	fmt.Printf("%T, %t\n", b2, b2)
	b3 := true
	fmt.Printf("%T, %t\n", b3, b3)

	//数值型
	var i1 int8 // -128 - 127
	i1 = 100
	fmt.Println(i1)

	var i2 int
	i2 = 20000
	fmt.Printf("%T, %d\n", i2, i2)

	var i3 int64
	//i3 = i2 //cannot use i2 (type int) as type int64 in assignment
	fmt.Printf("%T, %d\n", i3, i3)

	// byte: uint8   rune: int32  别称
	var i5 uint8
	i5 = 100
	var i6 byte
	i6 = i5
	fmt.Printf("%T, %d\n", i6, i6)

	var i7 = 100
	fmt.Printf("%T, %d\n", i7, i7)

	// float32 float64
	var f1 float32
	f1 = 3.14
	var f2 float64
	f2 = 4.68
	fmt.Printf("%T,%f\n", f1, f1)
	fmt.Printf("%T,%.2f\n", f1, f1)
	fmt.Printf("%T,%f\n", f2, f2)
	fmt.Printf("%T,%.3f\n", f2, f2)
	fmt.Println(f1, f2)

	var f3 = 1.11
	fmt.Printf("%T, %f", f3, f3)
















	/*
	复合数据类型
	*/
}
