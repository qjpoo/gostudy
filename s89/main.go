package main

import "fmt"

//　定义一个父类
type Person struct {
	name string
	age int
}

// 定义一个子类
type Student struct {
	Person  // 结构休的嵌套, 模拟继承
	school string
}
//  定义一个方法体

func (p Person) eat()  {
	fmt.Println(p.name, "正在吃饭中, 它来自于Person类...")
}

func (s Student) study()  {
	fmt.Println(s.name, "正在学习...")
}

func (s Student) eat()  {
	fmt.Println(s.name, "正在饭中, 它自至于Student类...")
}


func main() {
	/*
	方法的继承
	oop中的继承性:
		如果两个类存在继承关系,其中一个是子类,另一个就是父类
		1. 子类可以直接访问到父类的属性和方法
		2. 子类可以新增自己的属性和方法
		3. 子类可以重写父类可的方法(override,就是将浆糊的已有的方法,重新实现)

	 */

	p1 := Person{"chiling", 60}
	p1.eat()


	// 创建子类类型
	s1 := Student{Person{name:"quinn", age: 34}, "马哥Linux"}
	fmt.Println(s1.name) // 子类可以访问父类的属性
	fmt.Println(s1.school) // 子类可以访问自己新增加的属性
	s1.eat()  // 子类可以访问父类的方法
	s1.study()  // 子类可以新增加方法体
	s1.eat()  // 子类可以重写父类的方法


}
