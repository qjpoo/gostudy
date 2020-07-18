package main

import "fmt"

func f1()  {
}

func f2(x string, y, z int) int {
	return 0
}

func f3(x, y int) (int, string) {
	return 0, "A"
}


func main() {
	// 函数的数据类型
	// 函数的类型   func(参数列表的类型) (函数返回值列表的类型)

	i := 10
	f := 1.0
	s := "xxoo"
	arr := [4]int{1, 2, 3, 4}
	sl := []int{1,2,3,4}
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", arr)
	fmt.Printf("%T\n", sl)
	fmt.Printf("%T\n", f1) //注意这里不能写f1()，不然就是函数的调用了，要写函数名
	fmt.Printf("%T\n", f2)  // func(string, int, int) int
	fmt.Printf("%T\n", f3)  //  func(int, int) (int, string)
}
