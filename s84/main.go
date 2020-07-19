package main

import "fmt"

func main() {
	/* 结构体指针
	结构体是值类型的

	通过new() 生成的指针,并不是一个nil,空指针
	如果是string的话就是空,如果是int的话就是0
	 */
	// 值类型的
	p1 :=Person{"quinn", 35, "男", "湖北"}
	fmt.Printf("%T, %p\n", p1, &p1) // main.Person, 0xc0000d2000

	p2 := p1 // 值拷贝
	p2.name = "chiling"
	fmt.Printf("%T, %p\n", p2, &p2)  // main.Person, 0xc0000d2080

	// 定义结构体指针
	var p3 *Person
	p3 = &p1
	p3.name = "xxoo"  // 或者写成(*p3).name
	fmt.Printf("%T, %p, %v\n", p3, &p3, *p3)
	fmt.Printf("%T, %p, %v\n", p1, &p1, p1)


	// 使用内置的函数new(),go 语言中专门用于创建某种类型的指针的函数
	pp2 := new(Person)  // Persion结构体的指针
	fmt.Printf("%T\n", pp2)
	fmt.Println(pp2)  // &{ 0  }
	pp2.name = "A"
	pp2.age = 20
	pp2.sex = "男"
	pp2.address = "湖北"
	fmt.Println(pp2)

	pp3 := new(int)
	fmt.Printf("%T, %v, %v", pp3, pp3, *pp3)  // *int, 0xc00000a0b8, 0
}


// 定义结构体
type Person struct {
	name    string
	age     int
	sex     string
	address string
}