package main

import "fmt"

func main() {
	// 单向的通道

	ch1 := make(chan int) // 双向,可读可写
	//ch2 := make(chan <- int) // 单向,只能写
	//ch3 := make(<- chan  int) // 单向,只能读

	//ch1 <- 10
	//data := <- ch1
	//ch2 <- 10
	//ch3 <- 100  //报错了,ch3只能读
	// data := <- ch2 ch2只能写,不能读


	// ch1是可读可写的
	go fun1(ch1)
	data := <- ch1
	fmt.Println(data)

	go fun2(ch1)
	ch1 <- 200
	fmt.Println("main end ...")
}

// 函数, 只能操作只写的通道
func fun1(ch chan <- int)  {
	ch <- 100
	fmt.Println("fun1函数 over ...")

}

// 函数,只能读
func fun2(ch <- chan int)  {
	data := <- ch
	fmt.Println("fun2中读取的data: ", data)

}