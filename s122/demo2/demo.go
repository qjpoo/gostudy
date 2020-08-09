package main

import "fmt"

func main() {
	// <- a 从通道a中读数据
	// a <- data 向通道a中写数据

	var ch1 chan bool
	ch1 = make(chan bool)

	go func() {
		for i:=0;i<=5;i++ {
			fmt.Println("子的gorotine中 i: ", i)
		}
		// 循环结束后,向通道中写数据,表示结束
		ch1 <- true
		fmt.Println("子的goroutines结束 ...")
	}()

	data := <- ch1  // 阻塞的一直等待ch1中有数据
	fmt.Println("main data: ",data)
	fmt.Println("main end ...")
}
