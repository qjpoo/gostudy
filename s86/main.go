package main

import "fmt"

func main() {
	/* 结构体的嵌套

	 */

	b1 := Book{}
	b1.bookName = "三国演义"
	b1.price = 97.2

	s1 := Student{}
	s1.name = "qujian"
	s1.age = 29
	s1.book = b1 // 值传递
	fmt.Println(s1, b1)
	fmt.Printf("学生姓名: %s, 学生的年纪: %d, 看的书是: %v, 书的价格: %f\n",s1.name, s1.age, s1.book.bookName, s1.book.price)

	fmt.Println("-------------------")
	// 修改了s1.book.bookName对b1是没有影响的
	s1.book.bookName = "美女是怎么练成的"
	fmt.Println(s1)
	fmt.Println(b1)
	fmt.Println("-------------------")

	s2 := Student{name: "chiling", age: 50, book: Book{bookName: "golang是怎么练没的", price: 89.9}}
	fmt.Println(s2)
	fmt.Printf("%T\n", s2)

	s3 := Student{
		name: "qujian",
		age: 33,
		book: Book{
			bookName: "hello, world",
			price: 67.2,
		},
	}
	fmt.Println(s3)

	fmt.Println("---------1----------")
	b4 := Book{bookName: "xxxxooo", price: 87.2}
	s4 := Student2{name: "chiling", age: 50, book: &b4}
	fmt.Println(b4, s4)
	fmt.Println("\t", s4.book)
	fmt.Println("--------1-----------")

	s4.book.bookName = "1xxxooo"
	fmt.Println(b4)
	fmt.Println(s4)
	fmt.Println("\t", s4.book)
}

// 定义书的结构体
type Book struct {
	bookName string
	price    float64
}

// 定义一个学生
type Student struct {
	name string
	age  int
	book Book
}

type Student2 struct {
	name string
	age  int
	book *Book // book的地址
}
