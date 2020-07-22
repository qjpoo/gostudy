package main

import "fmt"

// 1. 定义父类
type Person struct {
	name string
	age int
}

// 2. 定义一个子类
type Student struct {
	Person // 匿名字段,Person结构体里面的name,age字段对于Student这个结构体叫提升字段
	// 提升字段可以简写 s3.Person.name == s3.name
	// 提升字段最大的意义是可以简写,所以student对象是可以直接访问到person中的字段
	school string
}
func main() {
	/*
	1. 模拟继承
	type A struct {
		field
	}

	type B struct {
		A  // 匿名字段
	}

	2. 模拟聚合关系
			type C struct {
				c1 string
			}

			type D struct {
	  		    c  C  // 聚合关系
				d string
		}


	 */

	// 创建父类
	p1 := Person{name: "quinn", age: 20}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.age)

	s1 := Student{Person{name: "chiling", age: 20}, "二中"}
	fmt.Println(s1)
	fmt.Println(s1.Person.name, s1.school)

	s2 := Student{Person:Person{name: "xiao",age: 50}, school: "一中"}
	fmt.Println(s2.Person.name, s2.school)

	var s3 Student
	s3.Person.name = "小王"
	s3.Person.age = 10
	s3.school = "洪湖二中"

	var s4 Student
	s4.name = "chiling"
	s4.age = 49
	fmt.Println(s4)
	fmt.Println(s4.name, s4.age)
}

