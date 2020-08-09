package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 非缓冲通道: make(chan T) 一次发送, 一次接收, 都是阻塞的

	// 缓冲通道 make(chan T, capacity), 发送时,缓冲区的数据满了,才会阻塞,缓冲区为空时,接收才会阻塞


	cha1 := make(chan int)
	fmt.Println(len(cha1), cap(cha1))  // 0 0

	cha2 := make(chan int, 5)
	fmt.Println(len(cha2), cap(cha2)) // 0, 5
	cha2 <- 100
	fmt.Println(len(cha2), cap(cha2)) // 1, 5
	cha2 <- 200
	cha2 <- 300
	cha2 <- 400
	cha2 <- 500
	fmt.Println(len(cha2), cap(cha2)) // 5, 5

	fmt.Println("---------------------------")
	ch3 := make(chan string, 4)
	go sendData(ch3)
	for {
		v, ok := <- ch3
		if !ok {
			fmt.Println("ch3读完了..", ok)
			break
		}
		fmt.Println("读取的数据是: ", v, ok)
	}
	fmt.Println("main over ...")

}

func sendData(ch chan string)  {
	for i:=1;i<=10;i++{
		ch <- "数据" + strconv.Itoa(i) // 当向ch写入第5个数据时,因为通道满了,会阻塞
		fmt.Printf("子goroutine中写出第%d个数据 ...\n", i)
	}
	close(ch)

}
