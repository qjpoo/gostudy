package main

import "fmt"

func main() {
	/*
		定向通道

		双向:
			chan T
				cha <- data , 发送数据, 写入到cha通道中
				data <- cha  从cha通道中接收到数据

		单向: 定向
			chan <- T 只支持写
			<- chan T, 只读

	*/

	ch1 := make(chan string)
	done := make(chan bool)
	go sendData(ch1, done)

	data := <-ch1
	fmt.Println("data: ", data)

	<- done
	fmt.Println("main end ...")

}

func sendData(ch1 chan string, done chan bool) {
	ch1 <- "hello" // 发送写数据

	//data := <-ch1 // 从ch1接收数据
	//fmt.Println("从main传来的数据", data)
	done <- true

}
