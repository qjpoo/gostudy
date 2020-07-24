package main

import "fmt"

type A interface {
	test1()
}

type B interface {
	test2()
}

type C interface {
	A
	B
	test3()
}

type Cat struct { // 如果想实现接口c,就要实现接口c中的方法,还要一起实现接口A,B中的方法

}

func (c Cat) test1()  {
	fmt.Println("test1() ...")
}

func (c Cat) test2()  {
	fmt.Println("test2() ...")
}

func (c Cat) test3()  {
	fmt.Println("test3() ...")
}
func main() {
	// 接口的嵌套

	var cat Cat = Cat{}
	cat.test1()
	cat.test2()
	cat.test3()

	fmt.Println("------------------")
	var a A = cat
	a.test1()

	fmt.Println("------------------")
	var b B = cat
	b.test2()

	fmt.Println("------------------")
	var c C = cat
	c.test1()
	c.test2()
	c.test3()

	fmt.Println("----------------")
	// var c2 C = a  这个不行

	fmt.Println("----------------")
	var a2 A = c
	a2.test1()

}
