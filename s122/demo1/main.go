package main

import "fmt"

func main() {
	// channel 通道
	// 通道是一个引用型的数据
	// data := <- a 先写<-,在写通道名. 从一个通道中读取数据在赋值给data
	// a <- data   先写通道名,后写<- ,把data数据写到通道a中


	// 申明一个通道
	var a chan  int
	fmt.Printf("%T, %v\n", a, a)  // chan int, <nil>
	if a == nil {
		// 创建一个通道
		a = make(chan int)
		fmt.Printf("%T, %v\n", a, a)  // chan int, 0xc00003e060
	}

	// 证明他传递的是内存地址, 是地址传递
	test1(a)


	var ch1 chan bool
	ch1 = make(chan bool)
	fmt.Println(ch1)
	

}
func test1(x chan int)  {
	fmt.Printf("%T, %v\n", x, x)  // chan int, 0xc00003e060
}

