package main

import "fmt"

func main() {
	/*
		结构体: 是由一系列具有相同类型或者不同类型的数据构成的数据集合
		结构体成员是由一系列的成员变量构成,这些成员变量也被称为字段

	*/
	//方法一
	var p1 Person
	fmt.Println(p1)

	p1.name = "quinn"
	p1.age = 30
	p1.sex = "男"
	p1.address = "湖北"
	fmt.Printf("%s, %d, %s, %s", p1.name, p1.age, p1.sex, p1.address)

	// 方法二
	p2 := Person{}
	p2.name = "xiaoming"
	p2.age = 20
	p2.sex = "男"
	p2.address = "湖南"
	fmt.Printf("%s, %d, %s, %s", p1.name, p2.age, p2.sex, p2.address)

	// 方法三
	p3 := Person{name: "chiling", age: 50, sex: "女", address: "香港"}
	fmt.Println(p3)

	// 方法四
	p4 := Person{"xiaohong", 30, "女", "北京"}
	fmt.Println(p4)

	p5 := Person{
		name:    "xiaohong",
		age:     35,
		sex:     "女",
		address: "陕西",
	}
	fmt.Println(p5)
}

// 定义结构体
type Person struct {
	name    string
	age     int
	sex     string
	address string
}
