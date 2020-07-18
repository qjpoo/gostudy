package main

import (
	"fmt"
)

func main() {
	/* 指针
	指针是存储另外一个变量的内存地址的变量
	 */

	a := 10
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	// 获取一个变量的地址
	fmt.Printf("%p\n", &a)  // 0xc0000a0090

	// 声明一个指针
	var p *int
	fmt.Println(p) // 空指针 <nil>
	p=&a  // p存储的是变量a的地址
	fmt.Println(*p)  // 取值 * , 获取指针p里面存储变量地址,变量的值  10
	fmt.Println(p)  // 打印出p里面存储的地址0xc0000100a8

	fmt.Printf("%p\n", &p) // p自己的内存地址 0xc000006030

	// 操作变量,修改了变量的值 ,但是变量对应的内存地址是没有修改的
	a = 200
	fmt.Println(a)
	// 获取一个变量的地址
	fmt.Printf("%p\n", &a)  // 0xc0000a0090

	// 操作指针,修改了指针存储对应变量地址的值,也没有变化
	*p = 300
	fmt.Println(a)
	fmt.Printf("%p\n", &p) // p自己的内存地址 0xc000006030
	fmt.Println(*p)  // 取值 * , 获取指针p里面存储变量地址,变量的值  10
	fmt.Println(p)  // 打印出p里面存储的地址0xc0000100a8

	// 指针的指针
	var p2 **int  // p2也是一个指针,不管有多少个*
	fmt.Println(p2)
	fmt.Printf("%T\n", p2) //也是一个指针
	p2 = &p
	fmt.Printf("%T, %T, %T\n", p2, p, a) // **int, *int, int
	fmt.Printf("%T\n", &p2)
	fmt.Printf("%v, %v, %v\n", p2, *p2, **p2) // 0xc000006030, 0xc0000100a8, 300

}
