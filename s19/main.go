package main

import "fmt"

func main() {
	var a int8
	a = 10

	var b int16
	b = int16(a)

	fmt.Println(a, b)

	f1 := 3.14
	fmt.Printf("%T,%f\n", f1, f1)
	var c int
	c = int(f1)
	fmt.Println(c)

	//变量是要转行的，而常量是不用的
	var d float64
	d = f1 + 100
	fmt.Println(d)

}
