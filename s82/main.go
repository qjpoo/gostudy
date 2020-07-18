package main

import "fmt"

func f1(a int)  {
	fmt.Println("在f1中之前的值为: ", a)
	a = 2
	fmt.Println("在f1中之后值为: ", a)
}

func f2(a *int)  {  // a = &a
	fmt.Println("在f2中之前的值为: ", *a)
	*a = 2
	fmt.Println("在f2中之后值为: ", *a)
}

func f3(a [4]int)  {
	fmt.Println("在f3中之前的值为: ", a)
	a[0] = 2
	fmt.Println("在f3中之后值为: ", a)
}

func f4(a *[4]int)  {  // a = &a
	fmt.Println("在f4中之前的值为: ", (*a)[0])
	(*a)[0] = 2
	fmt.Println("在f4中之后值为: ", (*a)[0])
}
func main() {
	// 指针作为参数
	a := 1
	fmt.Println("a的初始值为: ",a)
	f1(a) // 传值
	fmt.Println("a在调用f1函数之后的值为: ",a)

	fmt.Println("-----------------------")
	fmt.Println("a的初始值为: ",a)
	f2(&a) // 传地址,就是引用传递
	fmt.Println("a在调用f1函数之后的值为: ",a)

	arr := [4]int{1,2,3,4}
	fmt.Println("-----------------------")
	fmt.Println("arr的初始值为: ",arr)
	f3(arr) // 传值
	fmt.Println("arr在调用f3函数之后的值为: ",arr)

	fmt.Println("-----------------------")
	fmt.Println("arr的初始值为: ",arr)
	f4(&arr) // 传地址,就是引用传递
	fmt.Println("arr在调用f4函数之后的值为: ",arr)

}
