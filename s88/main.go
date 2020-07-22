package main

import "fmt"

// 定义一个工人的结构体
type Worker struct {
	name string
	age int
	sex string
}

// 定义行为方法
func (w Worker) work() {
	fmt.Println(w.name, "正在努力的工作 ...")
}

func (p *Worker) rest()  {
	fmt.Println(p.name, "正在休息中 ...") // (*p).name

}

// 方法体的名字可以相同,但是类型要不一样
type Cat struct {
	name string
	age int
}

func (p *Worker) printinfo()  {
	fmt.Printf("worker info:　name: %v, age:　%v,　sex: %v\n",p.name, p.age, p.sex)
}

func (p *Cat) printinfo()  {
	fmt.Printf("cat info:　name: %v, age:　%v\n",p.name, p.age)
}

func main() {
	/*
	方法: 方法也是一个函数,它是一个包含了接收者的函数,接受者可以是命名类型或者是结构体类型的一个值或者是一个指针

	函数:
	func funcname(para list) (return lsit) {}

	方法  要指定一个接收者类型,相当于说就是这个接收者调用这个方法
	func (t type) methodName(para list) (return list) {}

	对比函数:
		A: 意义
			方法: 某个类别的行为功能 , 需要指定接收者调用
			函数: 一段独立的功能代码,可以直接调用
		B: 语法:
			方法: 方法名可以相同, 只要接受者不同
			函数: 命名不能冲突
	
	*/

	w1 := Worker{name: "chiling", age: 20, sex: "女"}
	w1.work()  // w = w1,把w1的值 copy一份给了w

	w2 :=&Worker{name: "quinn", age: 20, sex: "男"}  //地址
	fmt.Printf("%T\n", w2)
	fmt.Println(w2)
	w2.work()
	w2.rest()  // p = w2 传递的是指针地址
	w1.rest()  // p = w1  也是指针地址

	c1 := Cat{name:"cat", age: 4}
	c1.printinfo()

	w1.printinfo()
}
