package main

import (
	"fmt"
	"runtime"
	"time"
)

func init()  {
	// 设置程序执行的最大CPU核数
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("cpu core: ", n)
}


func main() {
	// runtime包


	// goroot目录
	fmt.Println("goroot: ", runtime.GOROOT())

	// os
	fmt.Println("os: ", runtime.GOOS)

	//cpu cores
	fmt.Println("cpu cores:", runtime.NumCPU())


/*	go func() {
		for i:=1;i<=50;i++ {
			fmt.Println("子gorouteins: ", i)
		}
	}()

	for i:=0;i<6;i++ {
		// 让出时间片
		runtime.Gosched()
		fmt.Println("main goroutines: ", i)
	}*/


	fmt.Println("----------------------")
	/*
	子goroutines start ...
	fun函数 ...
	defer ...
	子goroutines end ...
	 */
	go func() {
		fmt.Println("子goroutines start ...")
		// 调用fun
		fun()
		fmt.Println("子goroutines end ...")
	}()

	time.Sleep(3 * time.Second)
}

func fun()  {
	defer fmt.Println("defer ...")
	//return // 中止函数
	runtime.Goexit() // 会执行defer,但是在main函数中从调用处fun()之后都不会执行了
	fmt.Println("fun函数 ...")
}
