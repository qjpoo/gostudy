package main

import "fmt"

func main() {
	/* 结构体的匿名字段

	 */

	// 正常的结构体
	s1 := Student{name: "chiling", age: 49}
	fmt.Println(s1.name, s1.age)

	// 匿名函数
	func() {
		fmt.Println("hello, world")
	}()

	// 匿名结构体
	s2 := struct {
		name string
		age  int
	}{"chi ling", 49}
	fmt.Println(s2.name, s2.age)

	s3 := struct {
		name string
		age  int
	}{
		name: "qujian",
		age: 35,
	}
	fmt.Println(s3.name, s3.age)


	// 调用匿名字段
	w1 := Worker{"qujian", 18}
	fmt.Println(w1.string, w1.int)
}

type Student struct {
	name string
	age  int
}

// 匿名字段
type Worker struct {
	//name string
	//age int
	string
	int
}
