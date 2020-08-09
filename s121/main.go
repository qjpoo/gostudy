package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex // 创建一个读写互斥对象
var wg *sync.WaitGroup
func main() {
	// 读写锁
	// 如果数据不变化,就不会有临界资源的问题
	// 锁主要是为了解决临界资源的问题, 关键是写的数据,即是修改的数据.修改的数据就需要同步,其它的goroutine才能感知到
	// sync包有两种锁 sync.Mutex和sync.RWMutex  sync.Mutex 互斥锁是通过计数器功能来实现的
	// sync.RWMutex 读写互斥锁,分别针对读操作和写操作的互斥锁
	// 真真的互斥是,读取和修改,修改和修改之间,读读之间没有互斥操作
	// 任意多个读操作,同一时刻只能有一个写操作,在同一时刻写操作的时候,读操作是不允许的
	// 当一个goroutine在写操作的时候,别的goroutine即不能读也不能写
	// 它和互斥锁最大的不同在于它可以分别针对读操作和写操作

	// RLock 读的锁, 读的锁可以加载多个
	// RUnlock 读的解锁

	// Lock 写的锁
	// Unlock 写锁解锁


	rwMutex = new(sync.RWMutex)
	wg = new(sync.WaitGroup)
/*
	wg.Add(2)

	// 读操作是可以同时进行的
	go readData(1)
	go readData(2)
	wg.Wait()
	fmt.Println("main end ...")
*/
	wg.Add(3)
	go writeData(1)
	go readData(1)
	go writeData(2)
	wg.Wait()
	fmt.Println("main end ...")

}


func writeData(i int)  {
	defer wg.Done()

	fmt.Println(i, "开始写 ...")
	rwMutex.Lock() // 写操作上锁
	fmt.Println(i, "正在写 ...")
	time.Sleep(3 * time.Second)
	rwMutex.Unlock() // 写操作解锁
	fmt.Println("写操作完成")

}



func readData( i int)  {
	defer wg.Done()
	fmt.Println(i, "开始读: read start ...")

	rwMutex.RLock() // 读操作上锁
	fmt.Println(i, "正在读取数据 ...")
	time.Sleep(5 * time.Second)
	rwMutex.RUnlock() // 读解锁
	fmt.Println(i, "读结束 ...")
}