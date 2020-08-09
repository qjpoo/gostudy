package main

import (
	"fmt"
	"time"
)

func main() {
	// 通道的闭关

	ch1 := make(chan int)
	go sendData(ch1)

	// 读取通道中的数据
	for {
		time.Sleep(1*time.Second)
		v, ok := <- ch1
		if !ok {
			fmt.Println("所有的数据读取完毕了 ", ok, v)
			break
		}
		fmt.Println("读取的数据: ", v, ok)
	}
	fmt.Println("main over ...")
}

func sendData(ch1 chan int) {
	for i := 0; i < 10; i++ {
		ch1 <- i // 将i写入到通道中
	}
	close(ch1) // 将通道关闭了
	fmt.Println("ch1通道关闭了 ...")

}
