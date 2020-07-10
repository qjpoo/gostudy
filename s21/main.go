package main

import "fmt"

func main() {
	// 关系运算符 >,<,>=,<=, ==, !=
	a := 3
	b := 5
	c := 3
	res1 := a > b
	res2 := b >c
	fmt.Printf("%T, %t\n",res1, res1)
	fmt.Printf("%T, %t\n",res2, res2)

	res3 := a == b
	res4 := b != c
	fmt.Printf("%T, %t\n",res3, res3)
	fmt.Printf("%T, %t\n",res4, res4)
	fmt.Println(a != b, a == c)
}
