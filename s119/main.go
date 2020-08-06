package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // 申明一个waitgroup类型的变量,就是一个等待组的对象

func main() {
	// waitGroup  同步等待组
	// Add() 设置等待组中的要执行的子goroutine的数量
	// Wait() 让主的goroutine出山于等待
	// Done(), 让等待组中的counter计数器的值减1 == Add(-1)一样

	//wg.Add(2) //2上goroutes
	//wg.Add(1) //2上goroutes
	wg.Add(3) //2上goroutes
	go A()
	go B()
	fmt.Println("main进入了阻塞状态, 等待wg中的子goroutine结束 ...")
	wg.Wait() // 表示main goroutine进入等待,意思着阻塞
	fmt.Println("main 解除阻塞 ...")
	//wg.Wait()

}

func A()  {
	for i:=1;i<20;i++ {
		fmt.Println("A ... ", i)
	}
	wg.Done() // 给等等组中的count数值减1
}

func B()  {
	for i:=1;i<5;i++ {
		fmt.Println("B ... ", i)
	}
	wg.Done() // 给等等组中的count数值减1
}
