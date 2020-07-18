package main

import "fmt"

func f1(){
	fmt.Println( "f1 ...")
}

func f2() [4]int {
	arr1 :=[4]int{1,2,3,4}
	return arr1
}

func f3() *[4]int  {
	arr := [4]int{1,2,3,4}
	return &arr

}
func main() {
	// 函数指针 和 指针函数
	/*
	函数指针: 一个指针,指向了一个函数的指针
		因为go语言中,function,默认看作一个指针,没有 *
	slice map function
	函数也是一种默认的指针类型

	指针函数: 一个函数,该函数的返回值是一个指针
	 */

	var a func()
	a = f1
	a()
	fmt.Printf("%T,%v, %v\n",a, a,&a)
	fmt.Printf("%T,%v\n",f1, f1)

	a1 := f2()
	fmt.Printf("%T, %v, %v\n", a1, a1, &a1)

	a2 := f3()
	fmt.Printf("%T, %v, %v\n", a2, a2, &a2)


}
