package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go func() {
		fmt.Println("子goroutine开始执行 ...")
		//time.Sleep(3 * time.Second)
		data := <-ch1 // 从ch1中读取数据
		fmt.Println("data: ", data)
	}()

	time.Sleep(3 * time.Second)
	ch1 <- 1
	fmt.Println("main end ...")
}
