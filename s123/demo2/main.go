package main

import "fmt"

func main() {
	// 通过 range来访问通道
	
	ch1 := make(chan int)
	go sendData(ch1)
	for v :=range ch1{
		fmt.Println("读取数据: ", v)  //从通道中读取数据
	}
	fmt.Println("main over ... ")
}

func sendData(ch1 chan int)  {
	for i:=0;i<10;i++ {
		ch1<- i
	}
	close(ch1)  // 关闭通道
}
