package main

import "fmt"
func main() {
	//算术运算符
	a := 10
	b := 3
	sum := a + b
	fmt.Printf("%d = %d + %d\n", sum, a, b)

	sub := a - b
	fmt.Printf("%d = %d - %d\n", sub, a, b)

	mul := a * b
	fmt.Printf("%d = %d * %d\n", mul, a, b)

	c := a / b
	fmt.Printf("%d = %d / %d\n", c, a, b)

	mo := a % b
	fmt.Printf("%d = %d / %d\n", mo, a, b)

	d := 1
	d++
	fmt.Println(d)

	e := 1
	e--
	fmt.Println(e)

	//f := 100
	//sum2 := f + f++ + --f
}
