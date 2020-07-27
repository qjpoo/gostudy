package main

import (
	"fmt"
	"strconv"
)

type myint int


type mystr string

// 定义一个函数类型
type myfun func(int, int)(string)

func fun1() myfun  { // fun1()函数的返回值是myfun类型
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}

// 类型别名
type myint2 = int // 和Int是通用的

// -------------------------------------------
type Person struct {
	name string
}

func (p Person) show() {
	fmt.Println("Person ---->", p.name)
}

type People = Person

// 嵌入两个结构体
type Student struct {
	Person
	People
}








func main() {
	/* type 关键字
	type 类型名 Type

	类型别名
	type 别名 = Type  // 注意有个等号,有等号,就是一回事,只是别名而已
	*/

	var i1 myint
	var i2 = 100 // int
	i1 = 200
	fmt.Println(i1, i2)
	fmt.Printf("%T, %T\n", i1, i2) // main.myint, int
	// i1 = i2 cannot use i2 (type int) as type myint in assignment

	// 在语法上不能通用
	var s1 mystr
	var s2 string
	s1 = "hello"
	s2 = "world"
	fmt.Println(s1, s2)
	fmt.Printf("%T, %T\n", s1, s2)  // main.mystr, string
	// s2 = s1  // cannot use s1 (type mystr) as type string in assignment
	// s1 = s2

	fmt.Println("------------------------------")
	res1 := fun1()
	fmt.Println(res1(10, 20))

	fmt.Println("------------------------------")
	var i3 myint2
	i3 = 1000
	i3 = i2
	fmt.Println(i3)
	fmt.Printf("%T, %T, %T", i3, i2, i1)  //  int, int, main.myint


	fmt.Println("------------------------------")
	var s Student
	s.Person.name = "hello"
	fmt.Println(s.Person.name)
	s.Person.show()

}
